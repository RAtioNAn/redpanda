//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	metav1 "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminAPI) DeepCopyInto(out *AdminAPI) {
	*out = *in
	in.External.DeepCopyInto(&out.External)
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminAPI.
func (in *AdminAPI) DeepCopy() *AdminAPI {
	if in == nil {
		return nil
	}
	out := new(AdminAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminAPITLS) DeepCopyInto(out *AdminAPITLS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminAPITLS.
func (in *AdminAPITLS) DeepCopy() *AdminAPITLS {
	if in == nil {
		return nil
	}
	out := new(AdminAPITLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStorageConfig) DeepCopyInto(out *CloudStorageConfig) {
	*out = *in
	out.SecretKeyRef = in.SecretKeyRef
	if in.CacheStorage != nil {
		in, out := &in.CacheStorage, &out.CacheStorage
		*out = new(StorageSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStorageConfig.
func (in *CloudStorageConfig) DeepCopy() *CloudStorageConfig {
	if in == nil {
		return nil
	}
	out := new(CloudStorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PDBConfig)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.Sidecars.DeepCopyInto(&out.Sidecars)
	in.Configuration.DeepCopyInto(&out.Configuration)
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Storage.DeepCopyInto(&out.Storage)
	in.CloudStorage.DeepCopyInto(&out.CloudStorage)
	if in.Superusers != nil {
		in, out := &in.Superusers, &out.Superusers
		*out = make([]Superuser, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalConfiguration != nil {
		in, out := &in.AdditionalConfiguration, &out.AdditionalConfiguration
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RestartConfig != nil {
		in, out := &in.RestartConfig, &out.RestartConfig
		*out = new(RestartConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.Nodes.DeepCopyInto(&out.Nodes)
	if in.DecommissioningNode != nil {
		in, out := &in.DecommissioningNode, &out.DecommissioningNode
		*out = new(int32)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Connect) DeepCopyInto(out *Connect) {
	*out = *in
	if in.ConnectTimeout != nil {
		in, out := &in.ConnectTimeout, &out.ConnectTimeout
		*out = new(apismetav1.Duration)
		**out = **in
	}
	if in.ReadTimeout != nil {
		in, out := &in.ReadTimeout, &out.ReadTimeout
		*out = new(apismetav1.Duration)
		**out = **in
	}
	if in.RequestTimeout != nil {
		in, out := &in.RequestTimeout, &out.RequestTimeout
		*out = new(apismetav1.Duration)
		**out = **in
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]ConnectCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connect.
func (in *Connect) DeepCopy() *Connect {
	if in == nil {
		return nil
	}
	out := new(Connect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectCluster) DeepCopyInto(out *ConnectCluster) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(ConnectClusterTLS)
		(*in).DeepCopyInto(*out)
	}
	if in.BasicAuthRef != nil {
		in, out := &in.BasicAuthRef, &out.BasicAuthRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.TokenRef != nil {
		in, out := &in.TokenRef, &out.TokenRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectCluster.
func (in *ConnectCluster) DeepCopy() *ConnectCluster {
	if in == nil {
		return nil
	}
	out := new(ConnectCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectClusterTLS) DeepCopyInto(out *ConnectClusterTLS) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectClusterTLS.
func (in *ConnectClusterTLS) DeepCopy() *ConnectClusterTLS {
	if in == nil {
		return nil
	}
	out := new(ConnectClusterTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Connectivity) DeepCopyInto(out *Connectivity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connectivity.
func (in *Connectivity) DeepCopy() *Connectivity {
	if in == nil {
		return nil
	}
	out := new(Connectivity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Console) DeepCopyInto(out *Console) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Console.
func (in *Console) DeepCopy() *Console {
	if in == nil {
		return nil
	}
	out := new(Console)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Console) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleList) DeepCopyInto(out *ConsoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Console, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleList.
func (in *ConsoleList) DeepCopy() *ConsoleList {
	if in == nil {
		return nil
	}
	out := new(ConsoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleSpec) DeepCopyInto(out *ConsoleSpec) {
	*out = *in
	in.Server.DeepCopyInto(&out.Server)
	out.SchemaRegistry = in.SchemaRegistry
	out.ClusterKeyRef = in.ClusterKeyRef
	out.Deployment = in.Deployment
	in.Connect.DeepCopyInto(&out.Connect)
	if in.Enterprise != nil {
		in, out := &in.Enterprise, &out.Enterprise
		*out = new(Enterprise)
		**out = **in
	}
	if in.License != nil {
		in, out := &in.License, &out.License
		*out = new(SecretKeyRef)
		**out = **in
	}
	if in.Login != nil {
		in, out := &in.Login, &out.Login
		*out = new(EnterpriseLogin)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleSpec.
func (in *ConsoleSpec) DeepCopy() *ConsoleSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleStatus) DeepCopyInto(out *ConsoleStatus) {
	*out = *in
	if in.ConfigMapRef != nil {
		in, out := &in.ConfigMapRef, &out.ConfigMapRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.Connectivity != nil {
		in, out := &in.Connectivity, &out.Connectivity
		*out = new(Connectivity)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleStatus.
func (in *ConsoleStatus) DeepCopy() *ConsoleStatus {
	if in == nil {
		return nil
	}
	out := new(ConsoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Deployment) DeepCopyInto(out *Deployment) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Deployment.
func (in *Deployment) DeepCopy() *Deployment {
	if in == nil {
		return nil
	}
	out := new(Deployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Enterprise) DeepCopyInto(out *Enterprise) {
	*out = *in
	out.RBAC = in.RBAC
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Enterprise.
func (in *Enterprise) DeepCopy() *Enterprise {
	if in == nil {
		return nil
	}
	out := new(Enterprise)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseLogin) DeepCopyInto(out *EnterpriseLogin) {
	*out = *in
	out.JWTSecret = in.JWTSecret
	if in.Google != nil {
		in, out := &in.Google, &out.Google
		*out = new(EnterpriseLoginGoogle)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseLogin.
func (in *EnterpriseLogin) DeepCopy() *EnterpriseLogin {
	if in == nil {
		return nil
	}
	out := new(EnterpriseLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseLoginGoogle) DeepCopyInto(out *EnterpriseLoginGoogle) {
	*out = *in
	out.ClientCredentials = in.ClientCredentials
	if in.Directory != nil {
		in, out := &in.Directory, &out.Directory
		*out = new(EnterpriseLoginGoogleDirectory)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseLoginGoogle.
func (in *EnterpriseLoginGoogle) DeepCopy() *EnterpriseLoginGoogle {
	if in == nil {
		return nil
	}
	out := new(EnterpriseLoginGoogle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseLoginGoogleDirectory) DeepCopyInto(out *EnterpriseLoginGoogleDirectory) {
	*out = *in
	out.ServiceAccountRef = in.ServiceAccountRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseLoginGoogleDirectory.
func (in *EnterpriseLoginGoogleDirectory) DeepCopy() *EnterpriseLoginGoogleDirectory {
	if in == nil {
		return nil
	}
	out := new(EnterpriseLoginGoogleDirectory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseRBAC) DeepCopyInto(out *EnterpriseRBAC) {
	*out = *in
	out.RoleBindingsRef = in.RoleBindingsRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseRBAC.
func (in *EnterpriseRBAC) DeepCopy() *EnterpriseRBAC {
	if in == nil {
		return nil
	}
	out := new(EnterpriseRBAC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalConnectivityConfig) DeepCopyInto(out *ExternalConnectivityConfig) {
	*out = *in
	if in.Bootstrap != nil {
		in, out := &in.Bootstrap, &out.Bootstrap
		*out = new(LoadBalancerConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalConnectivityConfig.
func (in *ExternalConnectivityConfig) DeepCopy() *ExternalConnectivityConfig {
	if in == nil {
		return nil
	}
	out := new(ExternalConnectivityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaAPI) DeepCopyInto(out *KafkaAPI) {
	*out = *in
	in.External.DeepCopyInto(&out.External)
	in.TLS.DeepCopyInto(&out.TLS)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaAPI.
func (in *KafkaAPI) DeepCopy() *KafkaAPI {
	if in == nil {
		return nil
	}
	out := new(KafkaAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaAPITLS) DeepCopyInto(out *KafkaAPITLS) {
	*out = *in
	if in.IssuerRef != nil {
		in, out := &in.IssuerRef, &out.IssuerRef
		*out = new(metav1.ObjectReference)
		**out = **in
	}
	if in.NodeSecretRef != nil {
		in, out := &in.NodeSecretRef, &out.NodeSecretRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaAPITLS.
func (in *KafkaAPITLS) DeepCopy() *KafkaAPITLS {
	if in == nil {
		return nil
	}
	out := new(KafkaAPITLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerWithName) DeepCopyInto(out *ListenerWithName) {
	*out = *in
	in.KafkaAPI.DeepCopyInto(&out.KafkaAPI)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerWithName.
func (in *ListenerWithName) DeepCopy() *ListenerWithName {
	if in == nil {
		return nil
	}
	out := new(ListenerWithName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerConfig) DeepCopyInto(out *LoadBalancerConfig) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerConfig.
func (in *LoadBalancerConfig) DeepCopy() *LoadBalancerConfig {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerStatus) DeepCopyInto(out *LoadBalancerStatus) {
	*out = *in
	in.LoadBalancerStatus.DeepCopyInto(&out.LoadBalancerStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerStatus.
func (in *LoadBalancerStatus) DeepCopy() *LoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodesList) DeepCopyInto(out *NodesList) {
	*out = *in
	if in.Internal != nil {
		in, out := &in.Internal, &out.Internal
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalBootstrap != nil {
		in, out := &in.ExternalBootstrap, &out.ExternalBootstrap
		*out = new(LoadBalancerStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalAdmin != nil {
		in, out := &in.ExternalAdmin, &out.ExternalAdmin
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalPandaproxy != nil {
		in, out := &in.ExternalPandaproxy, &out.ExternalPandaproxy
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PandaproxyIngress != nil {
		in, out := &in.PandaproxyIngress, &out.PandaproxyIngress
		*out = new(string)
		**out = **in
	}
	if in.SchemaRegistry != nil {
		in, out := &in.SchemaRegistry, &out.SchemaRegistry
		*out = new(SchemaRegistryStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodesList.
func (in *NodesList) DeepCopy() *NodesList {
	if in == nil {
		return nil
	}
	out := new(NodesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PDBConfig) DeepCopyInto(out *PDBConfig) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PDBConfig.
func (in *PDBConfig) DeepCopy() *PDBConfig {
	if in == nil {
		return nil
	}
	out := new(PDBConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PandaproxyAPI) DeepCopyInto(out *PandaproxyAPI) {
	*out = *in
	in.External.DeepCopyInto(&out.External)
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PandaproxyAPI.
func (in *PandaproxyAPI) DeepCopy() *PandaproxyAPI {
	if in == nil {
		return nil
	}
	out := new(PandaproxyAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PandaproxyAPITLS) DeepCopyInto(out *PandaproxyAPITLS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PandaproxyAPITLS.
func (in *PandaproxyAPITLS) DeepCopy() *PandaproxyAPITLS {
	if in == nil {
		return nil
	}
	out := new(PandaproxyAPITLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedpandaConfig) DeepCopyInto(out *RedpandaConfig) {
	*out = *in
	out.RPCServer = in.RPCServer
	if in.KafkaAPI != nil {
		in, out := &in.KafkaAPI, &out.KafkaAPI
		*out = make([]KafkaAPI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdminAPI != nil {
		in, out := &in.AdminAPI, &out.AdminAPI
		*out = make([]AdminAPI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PandaproxyAPI != nil {
		in, out := &in.PandaproxyAPI, &out.PandaproxyAPI
		*out = make([]PandaproxyAPI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SchemaRegistry != nil {
		in, out := &in.SchemaRegistry, &out.SchemaRegistry
		*out = new(SchemaRegistryAPI)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedpandaConfig.
func (in *RedpandaConfig) DeepCopy() *RedpandaConfig {
	if in == nil {
		return nil
	}
	out := new(RedpandaConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedpandaResourceRequirements) DeepCopyInto(out *RedpandaResourceRequirements) {
	*out = *in
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.Redpanda != nil {
		in, out := &in.Redpanda, &out.Redpanda
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedpandaResourceRequirements.
func (in *RedpandaResourceRequirements) DeepCopy() *RedpandaResourceRequirements {
	if in == nil {
		return nil
	}
	out := new(RedpandaResourceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestartConfig) DeepCopyInto(out *RestartConfig) {
	*out = *in
	if in.DisableMaintenanceModeHooks != nil {
		in, out := &in.DisableMaintenanceModeHooks, &out.DisableMaintenanceModeHooks
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestartConfig.
func (in *RestartConfig) DeepCopy() *RestartConfig {
	if in == nil {
		return nil
	}
	out := new(RestartConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Schema) DeepCopyInto(out *Schema) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Schema.
func (in *Schema) DeepCopy() *Schema {
	if in == nil {
		return nil
	}
	out := new(Schema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaRegistryAPI) DeepCopyInto(out *SchemaRegistryAPI) {
	*out = *in
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalConnectivityConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(SchemaRegistryAPITLS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaRegistryAPI.
func (in *SchemaRegistryAPI) DeepCopy() *SchemaRegistryAPI {
	if in == nil {
		return nil
	}
	out := new(SchemaRegistryAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaRegistryAPITLS) DeepCopyInto(out *SchemaRegistryAPITLS) {
	*out = *in
	if in.IssuerRef != nil {
		in, out := &in.IssuerRef, &out.IssuerRef
		*out = new(metav1.ObjectReference)
		**out = **in
	}
	if in.NodeSecretRef != nil {
		in, out := &in.NodeSecretRef, &out.NodeSecretRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaRegistryAPITLS.
func (in *SchemaRegistryAPITLS) DeepCopy() *SchemaRegistryAPITLS {
	if in == nil {
		return nil
	}
	out := new(SchemaRegistryAPITLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaRegistryStatus) DeepCopyInto(out *SchemaRegistryStatus) {
	*out = *in
	if in.ExternalNodeIPs != nil {
		in, out := &in.ExternalNodeIPs, &out.ExternalNodeIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaRegistryStatus.
func (in *SchemaRegistryStatus) DeepCopy() *SchemaRegistryStatus {
	if in == nil {
		return nil
	}
	out := new(SchemaRegistryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyRef) DeepCopyInto(out *SecretKeyRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyRef.
func (in *SecretKeyRef) DeepCopy() *SecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(SecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	if in.ServerGracefulShutdownTimeout != nil {
		in, out := &in.ServerGracefulShutdownTimeout, &out.ServerGracefulShutdownTimeout
		*out = new(apismetav1.Duration)
		**out = **in
	}
	if in.HTTPServerReadTimeout != nil {
		in, out := &in.HTTPServerReadTimeout, &out.HTTPServerReadTimeout
		*out = new(apismetav1.Duration)
		**out = **in
	}
	if in.HTTPServerWriteTimeout != nil {
		in, out := &in.HTTPServerWriteTimeout, &out.HTTPServerWriteTimeout
		*out = new(apismetav1.Duration)
		**out = **in
	}
	if in.HTTPServerIdleTimeout != nil {
		in, out := &in.HTTPServerIdleTimeout, &out.HTTPServerIdleTimeout
		*out = new(apismetav1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sidecar) DeepCopyInto(out *Sidecar) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sidecar.
func (in *Sidecar) DeepCopy() *Sidecar {
	if in == nil {
		return nil
	}
	out := new(Sidecar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sidecars) DeepCopyInto(out *Sidecars) {
	*out = *in
	if in.RpkStatus != nil {
		in, out := &in.RpkStatus, &out.RpkStatus
		*out = new(Sidecar)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sidecars.
func (in *Sidecars) DeepCopy() *Sidecars {
	if in == nil {
		return nil
	}
	out := new(Sidecars)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SocketAddress) DeepCopyInto(out *SocketAddress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SocketAddress.
func (in *SocketAddress) DeepCopy() *SocketAddress {
	if in == nil {
		return nil
	}
	out := new(SocketAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	out.Capacity = in.Capacity.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Superuser) DeepCopyInto(out *Superuser) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Superuser.
func (in *Superuser) DeepCopy() *Superuser {
	if in == nil {
		return nil
	}
	out := new(Superuser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	if in.IssuerRef != nil {
		in, out := &in.IssuerRef, &out.IssuerRef
		*out = new(metav1.ObjectReference)
		**out = **in
	}
	if in.NodeSecretRef != nil {
		in, out := &in.NodeSecretRef, &out.NodeSecretRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}
