//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2021 Weaveworks or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MPL-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerFileSource) DeepCopyInto(out *ContainerFileSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerFileSource.
func (in *ContainerFileSource) DeepCopy() *ContainerFileSource {
	if in == nil {
		return nil
	}
	out := new(ContainerFileSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalLoadBalancer) DeepCopyInto(out *ExternalLoadBalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalLoadBalancer.
func (in *ExternalLoadBalancer) DeepCopy() *ExternalLoadBalancer {
	if in == nil {
		return nil
	}
	out := new(ExternalLoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalLoadBalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalLoadBalancerEndpoint) DeepCopyInto(out *ExternalLoadBalancerEndpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalLoadBalancerEndpoint.
func (in *ExternalLoadBalancerEndpoint) DeepCopy() *ExternalLoadBalancerEndpoint {
	if in == nil {
		return nil
	}
	out := new(ExternalLoadBalancerEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalLoadBalancerList) DeepCopyInto(out *ExternalLoadBalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MicrovmCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalLoadBalancerList.
func (in *ExternalLoadBalancerList) DeepCopy() *ExternalLoadBalancerList {
	if in == nil {
		return nil
	}
	out := new(ExternalLoadBalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalLoadBalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalLoadBalancerSpec) DeepCopyInto(out *ExternalLoadBalancerSpec) {
	*out = *in
	out.Endpoint = in.Endpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalLoadBalancerSpec.
func (in *ExternalLoadBalancerSpec) DeepCopy() *ExternalLoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalLoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalLoadBalancerStatus) DeepCopyInto(out *ExternalLoadBalancerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalLoadBalancerStatus.
func (in *ExternalLoadBalancerStatus) DeepCopy() *ExternalLoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalLoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmCluster) DeepCopyInto(out *MicrovmCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmCluster.
func (in *MicrovmCluster) DeepCopy() *MicrovmCluster {
	if in == nil {
		return nil
	}
	out := new(MicrovmCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrovmCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmClusterList) DeepCopyInto(out *MicrovmClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MicrovmCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmClusterList.
func (in *MicrovmClusterList) DeepCopy() *MicrovmClusterList {
	if in == nil {
		return nil
	}
	out := new(MicrovmClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrovmClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmClusterSpec) DeepCopyInto(out *MicrovmClusterSpec) {
	*out = *in
	in.Placement.DeepCopyInto(&out.Placement)
	if in.EndpointRef != nil {
		in, out := &in.EndpointRef, &out.EndpointRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmClusterSpec.
func (in *MicrovmClusterSpec) DeepCopy() *MicrovmClusterSpec {
	if in == nil {
		return nil
	}
	out := new(MicrovmClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmClusterStatus) DeepCopyInto(out *MicrovmClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(v1beta1.FailureDomains, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmClusterStatus.
func (in *MicrovmClusterStatus) DeepCopy() *MicrovmClusterStatus {
	if in == nil {
		return nil
	}
	out := new(MicrovmClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmHost) DeepCopyInto(out *MicrovmHost) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmHost.
func (in *MicrovmHost) DeepCopy() *MicrovmHost {
	if in == nil {
		return nil
	}
	out := new(MicrovmHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmMachine) DeepCopyInto(out *MicrovmMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmMachine.
func (in *MicrovmMachine) DeepCopy() *MicrovmMachine {
	if in == nil {
		return nil
	}
	out := new(MicrovmMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrovmMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmMachineList) DeepCopyInto(out *MicrovmMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MicrovmMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmMachineList.
func (in *MicrovmMachineList) DeepCopy() *MicrovmMachineList {
	if in == nil {
		return nil
	}
	out := new(MicrovmMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrovmMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmMachineSpec) DeepCopyInto(out *MicrovmMachineSpec) {
	*out = *in
	in.MicrovmSpec.DeepCopyInto(&out.MicrovmSpec)
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmMachineSpec.
func (in *MicrovmMachineSpec) DeepCopy() *MicrovmMachineSpec {
	if in == nil {
		return nil
	}
	out := new(MicrovmMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmMachineStatus) DeepCopyInto(out *MicrovmMachineStatus) {
	*out = *in
	if in.VMState != nil {
		in, out := &in.VMState, &out.VMState
		*out = new(VMState)
		**out = **in
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1beta1.MachineAddress, len(*in))
		copy(*out, *in)
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmMachineStatus.
func (in *MicrovmMachineStatus) DeepCopy() *MicrovmMachineStatus {
	if in == nil {
		return nil
	}
	out := new(MicrovmMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmMachineTemplate) DeepCopyInto(out *MicrovmMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmMachineTemplate.
func (in *MicrovmMachineTemplate) DeepCopy() *MicrovmMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(MicrovmMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrovmMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmMachineTemplateList) DeepCopyInto(out *MicrovmMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MicrovmMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmMachineTemplateList.
func (in *MicrovmMachineTemplateList) DeepCopy() *MicrovmMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(MicrovmMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicrovmMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmMachineTemplateResource) DeepCopyInto(out *MicrovmMachineTemplateResource) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmMachineTemplateResource.
func (in *MicrovmMachineTemplateResource) DeepCopy() *MicrovmMachineTemplateResource {
	if in == nil {
		return nil
	}
	out := new(MicrovmMachineTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmMachineTemplateSpec) DeepCopyInto(out *MicrovmMachineTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmMachineTemplateSpec.
func (in *MicrovmMachineTemplateSpec) DeepCopy() *MicrovmMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(MicrovmMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrovmSpec) DeepCopyInto(out *MicrovmSpec) {
	*out = *in
	out.RootVolume = in.RootVolume
	if in.AdditionalVolumes != nil {
		in, out := &in.AdditionalVolumes, &out.AdditionalVolumes
		*out = make([]Volume, len(*in))
		copy(*out, *in)
	}
	out.Kernel = in.Kernel
	if in.Initrd != nil {
		in, out := &in.Initrd, &out.Initrd
		*out = new(ContainerFileSource)
		**out = **in
	}
	if in.NetworkInterfaces != nil {
		in, out := &in.NetworkInterfaces, &out.NetworkInterfaces
		*out = make([]NetworkInterface, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrovmSpec.
func (in *MicrovmSpec) DeepCopy() *MicrovmSpec {
	if in == nil {
		return nil
	}
	out := new(MicrovmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterface) DeepCopyInto(out *NetworkInterface) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterface.
func (in *NetworkInterface) DeepCopy() *NetworkInterface {
	if in == nil {
		return nil
	}
	out := new(NetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Placement) DeepCopyInto(out *Placement) {
	*out = *in
	if in.StaticPool != nil {
		in, out := &in.StaticPool, &out.StaticPool
		*out = new(StaticPoolPlacement)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Placement.
func (in *Placement) DeepCopy() *Placement {
	if in == nil {
		return nil
	}
	out := new(Placement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticPoolPlacement) DeepCopyInto(out *StaticPoolPlacement) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]MicrovmHost, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticPoolPlacement.
func (in *StaticPoolPlacement) DeepCopy() *StaticPoolPlacement {
	if in == nil {
		return nil
	}
	out := new(StaticPoolPlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}
