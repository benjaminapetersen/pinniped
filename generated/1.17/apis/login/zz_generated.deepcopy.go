// +build !ignore_autogenerated

// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by deepcopy-gen. DO NOT EDIT.

package login

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCredential) DeepCopyInto(out *ClusterCredential) {
	*out = *in
	in.ExpirationTimestamp.DeepCopyInto(&out.ExpirationTimestamp)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCredential.
func (in *ClusterCredential) DeepCopy() *ClusterCredential {
	if in == nil {
		return nil
	}
	out := new(ClusterCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenCredentialRequest) DeepCopyInto(out *TokenCredentialRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenCredentialRequest.
func (in *TokenCredentialRequest) DeepCopy() *TokenCredentialRequest {
	if in == nil {
		return nil
	}
	out := new(TokenCredentialRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TokenCredentialRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenCredentialRequestList) DeepCopyInto(out *TokenCredentialRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TokenCredentialRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenCredentialRequestList.
func (in *TokenCredentialRequestList) DeepCopy() *TokenCredentialRequestList {
	if in == nil {
		return nil
	}
	out := new(TokenCredentialRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TokenCredentialRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenCredentialRequestSpec) DeepCopyInto(out *TokenCredentialRequestSpec) {
	*out = *in
	in.IdentityProvider.DeepCopyInto(&out.IdentityProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenCredentialRequestSpec.
func (in *TokenCredentialRequestSpec) DeepCopy() *TokenCredentialRequestSpec {
	if in == nil {
		return nil
	}
	out := new(TokenCredentialRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenCredentialRequestStatus) DeepCopyInto(out *TokenCredentialRequestStatus) {
	*out = *in
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(ClusterCredential)
		(*in).DeepCopyInto(*out)
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenCredentialRequestStatus.
func (in *TokenCredentialRequestStatus) DeepCopy() *TokenCredentialRequestStatus {
	if in == nil {
		return nil
	}
	out := new(TokenCredentialRequestStatus)
	in.DeepCopyInto(out)
	return out
}
