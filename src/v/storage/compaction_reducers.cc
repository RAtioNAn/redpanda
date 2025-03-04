// Copyright 2020 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

#include "storage/compaction_reducers.h"

#include "compression/compression.h"
#include "model/record.h"
#include "model/record_batch_types.h"
#include "model/record_utils.h"
#include "random/generators.h"
#include "storage/index_state.h"
#include "storage/logger.h"
#include "storage/parser_utils.h"
#include "storage/segment_appender_utils.h"
#include "storage/segment_utils.h"
#include "vlog.h"

#include <seastar/core/future.hh>

#include <absl/algorithm/container.h>
#include <absl/container/flat_hash_map.h>
#include <boost/range/irange.hpp>

#include <algorithm>
#include <exception>

namespace storage::internal {

ss::future<ss::stop_iteration>
compaction_key_reducer::operator()(compacted_index::entry&& e) {
    using stop_t = ss::stop_iteration;
    const model::offset o = e.offset + model::offset(e.delta);

    auto it = _indices.find(e.key);
    if (it != _indices.end()) {
        if (o > it->second.offset) {
            // cannot be std::max() because _natural_index must be preserved
            it->second.offset = o;
            it->second.natural_index = _natural_index;
        }
    } else {
        // not found - insert
        auto const key_size = e.key.size();
        auto const expected_size = [this, key_size] {
            return idx_mem_usage() + _keys_mem_usage + key_size;
        };

        // if index allocates to much memory remove some entries.
        while (expected_size() >= _max_mem && !_indices.empty()) {
            /**
             * Evict first entry, we use hash function that guarante good
             * randomness so evicting first entry is actually evicting a
             * pseudo random elemnent
             */
            auto mit = _indices.begin();
            _keys_mem_usage -= mit->first.size();

            // write the entry again - we ran out of scratch space
            _inverted.add(mit->second.natural_index);
            _indices.erase(mit);
        }
        _keys_mem_usage += e.key.size();
        // 2. do the insertion
        _indices.emplace(std::move(e.key), value_type(o, _natural_index));
    }

    ++_natural_index; // MOST important
    return ss::make_ready_future<stop_t>(stop_t::no);
}
Roaring compaction_key_reducer::end_of_stream() {
    // TODO: optimization - detect if the index does not need compaction
    // by linear scan of natural_index from 0-N with no gaps.
    for (auto& e : _indices) {
        _inverted.add(e.second.natural_index);
    }
    _inverted.shrinkToFit();
    return std::move(_inverted);
}

ss::future<ss::stop_iteration>
index_copy_reducer::operator()(compacted_index::entry&& e) {
    using stop_t = ss::stop_iteration;
    return _writer->append(std::move(e)).then([] {
        return ss::make_ready_future<stop_t>(stop_t::no);
    });
}

ss::future<ss::stop_iteration>
index_filtered_copy_reducer::operator()(compacted_index::entry&& e) {
    using stop_t = ss::stop_iteration;
    const bool should_add = _bm.contains(_natural_index);
    ++_natural_index;
    if (should_add) {
        return _writer->index(e.key, e.offset, e.delta)
          .then([k = std::move(e.key)] {
              return ss::make_ready_future<stop_t>(stop_t::no);
          });
    }
    return ss::make_ready_future<stop_t>(stop_t::no);
}

ss::future<ss::stop_iteration>
compacted_offset_list_reducer::operator()(compacted_index::entry&& e) {
    using stop_t = ss::stop_iteration;
    const model::offset o = e.offset + model::offset(e.delta);
    _list.add(o);
    return ss::make_ready_future<stop_t>(stop_t::no);
}

std::optional<model::record_batch>
copy_data_segment_reducer::filter(model::record_batch&& batch) {
    // do not compact raft configuration and archival metadata as they shift
    // offset translation
    if (!is_compactible(batch)) {
        return std::move(batch);
    }

    // 1. compute which records to keep
    const auto base = batch.base_offset();
    std::vector<int32_t> offset_deltas;
    offset_deltas.reserve(batch.record_count());
    batch.for_each_record([this, base, &offset_deltas](const model::record& r) {
        if (should_keep(base, r.offset_delta())) {
            offset_deltas.push_back(r.offset_delta());
        }
    });

    // 2. no record to keep
    if (offset_deltas.empty()) {
        return std::nullopt;
    }

    // 3. keep all records
    if (offset_deltas.size() == static_cast<size_t>(batch.record_count())) {
        return std::move(batch);
    }

    // 4. filter
    iobuf ret;
    int32_t rec_count = 0;
    std::optional<int64_t> first_timestamp_delta;
    int64_t last_timestamp_delta;
    batch.for_each_record([&rec_count,
                           &first_timestamp_delta,
                           &last_timestamp_delta,
                           &ret,
                           &offset_deltas](model::record record) {
        // contains the key
        if (std::count(
              offset_deltas.begin(),
              offset_deltas.end(),
              record.offset_delta())) {
            /*
             * TODO when we further optimize lazy record materialization ot
             * make use of views we can avoid this re-encoding by copying or
             * sharing the view. either way, we were building
             * record batch with the uncompressed records so they were being
             * re-encoded.
             */
            if (!first_timestamp_delta) {
                first_timestamp_delta = record.timestamp_delta();
            }
            last_timestamp_delta = record.timestamp_delta();
            model::append_record_to_buffer(ret, record);
            ++rec_count;
        }
    });
    // From: DefaultRecordBatch.java
    // On Compaction: Unlike the older message formats, magic v2 and above
    // preserves the first and last offset/sequence numbers from the
    // original batch when the log is cleaned. This is required in order to
    // be able to restore the producer'size() state when the log is reloaded. If
    // we did not retain the last sequence number, then following a
    // partition leader failure, once the new leader has rebuilt the
    // producer state from the log, the next sequence expected number would
    // no longer be in sync with what was written by the client. This would
    // cause an unexpected OutOfOrderSequence error, which is typically
    // fatal. The base sequence number must be preserved for duplicate
    // checking: the broker checks incoming Produce requests for duplicates
    // by verifying that the first and last sequence numbers of the incoming
    // batch match the last from that producer.
    //
    if (rec_count == 0) {
        // TODO:agallego - implement
        //
        // Note that if all of the records in a batch are removed during
        // compaction, the broker may still retain an empty batch header in
        // order to preserve the producer sequence information as described
        // above. These empty batches are retained only until either a new
        // sequence number is written by the corresponding producer or the
        // producerId is expired from lack of activity.
        return std::nullopt;
    }

    // There is no similar need to preserve the timestamp from the original
    // batch after compaction. The FirstTimestamp field therefore always
    // reflects the timestamp of the first record in the batch. If the batch is
    // empty, the FirstTimestamp will be set to -1 (NO_TIMESTAMP).
    //
    // Similarly, the MaxTimestamp field reflects the maximum timestamp of the
    // current records if the timestamp type is CREATE_TIME. For
    // LOG_APPEND_TIME, on the other hand, the MaxTimestamp field reflects the
    // timestamp set by the broker and is preserved after compaction.
    // Additionally, the MaxTimestamp of an empty batch always retains the
    // previous value prior to becoming empty.
    //
    const auto& oldh = batch.header();
    const auto first_time = model::timestamp(
      oldh.first_timestamp() + first_timestamp_delta.value());
    auto last_time = oldh.max_timestamp;
    if (oldh.attrs.timestamp_type() == model::timestamp_type::create_time) {
        last_time = model::timestamp(first_time() + last_timestamp_delta);
    }
    auto h = oldh;
    h.first_timestamp = first_time;
    h.max_timestamp = last_time;
    h.record_count = rec_count;
    reset_size_checksum_metadata(h, ret);
    auto new_batch = model::record_batch(
      h, std::move(ret), model::record_batch::tag_ctor_ng{});
    return new_batch;
}
ss::future<ss::stop_iteration> copy_data_segment_reducer::do_compaction(
  model::compression original, model::record_batch&& b) {
    using stop_t = ss::stop_iteration;
    auto to_copy = filter(std::move(b));
    if (to_copy == std::nullopt) {
        return ss::make_ready_future<stop_t>(stop_t::no);
    }
    return compress_batch(original, std::move(to_copy.value()))
      .then([this](model::record_batch&& b) {
          return ss::do_with(std::move(b), [this](model::record_batch& batch) {
              auto const start_offset = _appender->file_byte_offset();
              auto const header_size = batch.header().size_bytes;
              _acc += header_size;
              if (_idx.maybe_index(
                    _acc,
                    32_KiB,
                    start_offset,
                    batch.base_offset(),
                    batch.last_offset(),
                    batch.header().first_timestamp,
                    batch.header().max_timestamp)) {
                  _acc = 0;
              }
              return storage::write(*_appender, batch)
                .then([this, start_offset, header_size] {
                    vassert(
                      _appender->file_byte_offset()
                        == start_offset + header_size,
                      "Size must be deterministic. Expected:{} == {}",
                      _appender->file_byte_offset(),
                      start_offset + header_size);
                });
          });
      })
      .then([] { return ss::make_ready_future<stop_t>(stop_t::no); });
}

ss::future<ss::stop_iteration>
copy_data_segment_reducer::operator()(model::record_batch&& b) {
    const auto comp = b.header().attrs.compression();
    if (!b.compressed()) {
        return do_compaction(comp, std::move(b));
    }
    return decompress_batch(std::move(b))
      .then([comp, this](model::record_batch&& b) {
          return do_compaction(comp, std::move(b));
      });
}

ss::future<ss::stop_iteration>
index_rebuilder_reducer::operator()(model::record_batch&& b) {
    using stop_t = ss::stop_iteration;
    auto f = ss::now();
    if (!b.compressed()) {
        f = do_index(std::move(b));
    } else {
        f = internal::decompress_batch(std::move(b))
              .then([this](model::record_batch&& b) {
                  return do_index(std::move(b));
              });
    }
    return f.then([] { return ss::make_ready_future<stop_t>(stop_t::no); });
}

ss::future<> index_rebuilder_reducer::do_index(model::record_batch&& b) {
    return ss::do_with(std::move(b), [this](model::record_batch& b) {
        return model::for_each_record(
          b,
          [this, bt = b.header().type, o = b.base_offset()](model::record& r) {
              return _w->index(bt, r.key(), o, r.offset_delta());
          });
    });
}
} // namespace storage::internal
