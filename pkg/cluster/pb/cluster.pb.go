// Copyright 2020 The jackal Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: cluster.proto

package pb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	stravaganza "github.com/jackal-xmpp/stravaganza/v2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// StreamErrorReason is an enumerated type that describes stream error reason.
type StreamErrorReason int32

const (
	StreamErrorReason_STREAM_ERROR_REASON_INVALID_XML              StreamErrorReason = 0  // Invalid XML.
	StreamErrorReason_STREAM_ERROR_REASON_INVALID_NAMESPACE        StreamErrorReason = 1  // Invalid namespace.
	StreamErrorReason_STREAM_ERROR_REASON_HOST_UNKNOWN             StreamErrorReason = 2  // Host unknown.
	StreamErrorReason_STREAM_ERROR_REASON_CONFLICT                 StreamErrorReason = 3  // Conflict.
	StreamErrorReason_STREAM_ERROR_REASON_INVALID_FROM             StreamErrorReason = 4  // Invalid from.
	StreamErrorReason_STREAM_ERROR_REASON_POLICY_VIOLATION         StreamErrorReason = 5  // Policy violation.
	StreamErrorReason_STREAM_ERROR_REASON_REMOTE_CONNECTION_FAILED StreamErrorReason = 6  // Remote connection failed.
	StreamErrorReason_STREAM_ERROR_REASON_CONNECTION_TIMEOUT       StreamErrorReason = 7  // Connection timeout.
	StreamErrorReason_STREAM_ERROR_REASON_UNSUPPORTED_STANZA_TYPE  StreamErrorReason = 8  // Unsupported stanza type.
	StreamErrorReason_STREAM_ERROR_REASON_UNSUPPORTED_VERSION      StreamErrorReason = 9  // Unsupported version.
	StreamErrorReason_STREAM_ERROR_REASON_NOT_AUTHORIZED           StreamErrorReason = 10 // Not authorized.
	StreamErrorReason_STREAM_ERROR_REASON_RESOURCE_CONSTRAINT      StreamErrorReason = 11 // Resource constraint.
	StreamErrorReason_STREAM_ERROR_REASON_SYSTEM_SHUTDOWN          StreamErrorReason = 12 // System shutdown.
	StreamErrorReason_STREAM_ERROR_REASON_UNDEFINED_CONDITION      StreamErrorReason = 13 // Undefined condition.
	StreamErrorReason_STREAM_ERROR_REASON_INTERNAL_SERVER_ERROR    StreamErrorReason = 14 // Server error.
)

// Enum value maps for StreamErrorReason.
var (
	StreamErrorReason_name = map[int32]string{
		0:  "STREAM_ERROR_REASON_INVALID_XML",
		1:  "STREAM_ERROR_REASON_INVALID_NAMESPACE",
		2:  "STREAM_ERROR_REASON_HOST_UNKNOWN",
		3:  "STREAM_ERROR_REASON_CONFLICT",
		4:  "STREAM_ERROR_REASON_INVALID_FROM",
		5:  "STREAM_ERROR_REASON_POLICY_VIOLATION",
		6:  "STREAM_ERROR_REASON_REMOTE_CONNECTION_FAILED",
		7:  "STREAM_ERROR_REASON_CONNECTION_TIMEOUT",
		8:  "STREAM_ERROR_REASON_UNSUPPORTED_STANZA_TYPE",
		9:  "STREAM_ERROR_REASON_UNSUPPORTED_VERSION",
		10: "STREAM_ERROR_REASON_NOT_AUTHORIZED",
		11: "STREAM_ERROR_REASON_RESOURCE_CONSTRAINT",
		12: "STREAM_ERROR_REASON_SYSTEM_SHUTDOWN",
		13: "STREAM_ERROR_REASON_UNDEFINED_CONDITION",
		14: "STREAM_ERROR_REASON_INTERNAL_SERVER_ERROR",
	}
	StreamErrorReason_value = map[string]int32{
		"STREAM_ERROR_REASON_INVALID_XML":              0,
		"STREAM_ERROR_REASON_INVALID_NAMESPACE":        1,
		"STREAM_ERROR_REASON_HOST_UNKNOWN":             2,
		"STREAM_ERROR_REASON_CONFLICT":                 3,
		"STREAM_ERROR_REASON_INVALID_FROM":             4,
		"STREAM_ERROR_REASON_POLICY_VIOLATION":         5,
		"STREAM_ERROR_REASON_REMOTE_CONNECTION_FAILED": 6,
		"STREAM_ERROR_REASON_CONNECTION_TIMEOUT":       7,
		"STREAM_ERROR_REASON_UNSUPPORTED_STANZA_TYPE":  8,
		"STREAM_ERROR_REASON_UNSUPPORTED_VERSION":      9,
		"STREAM_ERROR_REASON_NOT_AUTHORIZED":           10,
		"STREAM_ERROR_REASON_RESOURCE_CONSTRAINT":      11,
		"STREAM_ERROR_REASON_SYSTEM_SHUTDOWN":          12,
		"STREAM_ERROR_REASON_UNDEFINED_CONDITION":      13,
		"STREAM_ERROR_REASON_INTERNAL_SERVER_ERROR":    14,
	}
)

func (x StreamErrorReason) Enum() *StreamErrorReason {
	p := new(StreamErrorReason)
	*p = x
	return p
}

func (x StreamErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_cluster_proto_enumTypes[0].Descriptor()
}

func (StreamErrorReason) Type() protoreflect.EnumType {
	return &file_cluster_proto_enumTypes[0]
}

func (x StreamErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamErrorReason.Descriptor instead.
func (StreamErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{0}
}

// LocalRouteRequest is the parameter message for LocalRouter Route rpc.
type LocalRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// username name of the user to which the stanza is routed to.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// resource is the registered resource to which the stanza is routed to.
	Resource string `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	// stanza contains the routed XMPP stanza.
	Stanza *stravaganza.PBElement `protobuf:"bytes,3,opt,name=stanza,proto3" json:"stanza,omitempty"`
}

func (x *LocalRouteRequest) Reset() {
	*x = LocalRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalRouteRequest) ProtoMessage() {}

func (x *LocalRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalRouteRequest.ProtoReflect.Descriptor instead.
func (*LocalRouteRequest) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *LocalRouteRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LocalRouteRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *LocalRouteRequest) GetStanza() *stravaganza.PBElement {
	if x != nil {
		return x.Stanza
	}
	return nil
}

// LocalRouteResponse is the response returned by LocalRouter Route rpc.
type LocalRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LocalRouteResponse) Reset() {
	*x = LocalRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalRouteResponse) ProtoMessage() {}

func (x *LocalRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalRouteResponse.ProtoReflect.Descriptor instead.
func (*LocalRouteResponse) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{1}
}

// LocalDisconnectRequest is the parameter message for LocalRouter Disconnect rpc.
type LocalDisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// username is the JID node value for the session we wish to disconnect.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// resource is the JID resourcepart value for the session we wish to disconnect.
	Resource string `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	// StreamError is disconnection stream error.
	StreamError *StreamError `protobuf:"bytes,3,opt,name=stream_error,json=streamError,proto3" json:"stream_error,omitempty"`
}

func (x *LocalDisconnectRequest) Reset() {
	*x = LocalDisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalDisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalDisconnectRequest) ProtoMessage() {}

func (x *LocalDisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalDisconnectRequest.ProtoReflect.Descriptor instead.
func (*LocalDisconnectRequest) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *LocalDisconnectRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LocalDisconnectRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *LocalDisconnectRequest) GetStreamError() *StreamError {
	if x != nil {
		return x.StreamError
	}
	return nil
}

// LocalDisconnectResponse is the response returned by LocalRouter Disconnect rpc.
type LocalDisconnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LocalDisconnectResponse) Reset() {
	*x = LocalDisconnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalDisconnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalDisconnectResponse) ProtoMessage() {}

func (x *LocalDisconnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalDisconnectResponse.ProtoReflect.Descriptor instead.
func (*LocalDisconnectResponse) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{3}
}

// ComponentRouteRequest is the parameter message for ComponentRouter Route rpc.
type ComponentRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stanza contains the routed XMPP stanza.
	Stanza *stravaganza.PBElement `protobuf:"bytes,1,opt,name=stanza,proto3" json:"stanza,omitempty"`
}

func (x *ComponentRouteRequest) Reset() {
	*x = ComponentRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentRouteRequest) ProtoMessage() {}

func (x *ComponentRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentRouteRequest.ProtoReflect.Descriptor instead.
func (*ComponentRouteRequest) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *ComponentRouteRequest) GetStanza() *stravaganza.PBElement {
	if x != nil {
		return x.Stanza
	}
	return nil
}

// ComponentRouteResponse is the response returned by ComponentRouter Route rpc.
type ComponentRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ComponentRouteResponse) Reset() {
	*x = ComponentRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentRouteResponse) ProtoMessage() {}

func (x *ComponentRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentRouteResponse.ProtoReflect.Descriptor instead.
func (*ComponentRouteResponse) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{5}
}

// StreamError represents a stream disconnection error reason.
type StreamError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason             StreamErrorReason      `protobuf:"varint,1,opt,name=reason,proto3,enum=cluster.v1.StreamErrorReason" json:"reason,omitempty"`
	Lang               string                 `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Text               string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	ApplicationElement *stravaganza.PBElement `protobuf:"bytes,4,opt,name=applicationElement,proto3" json:"applicationElement,omitempty"`
}

func (x *StreamError) Reset() {
	*x = StreamError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamError) ProtoMessage() {}

func (x *StreamError) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamError.ProtoReflect.Descriptor instead.
func (*StreamError) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{6}
}

func (x *StreamError) GetReason() StreamErrorReason {
	if x != nil {
		return x.Reason
	}
	return StreamErrorReason_STREAM_ERROR_REASON_INVALID_XML
}

func (x *StreamError) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *StreamError) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *StreamError) GetApplicationElement() *stravaganza.PBElement {
	if x != nil {
		return x.ApplicationElement
	}
	return nil
}

var File_cluster_proto protoreflect.FileDescriptor

var file_cluster_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x63, 0x6b, 0x61, 0x6c, 0x2d, 0x78,
	0x6d, 0x70, 0x70, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x67, 0x61, 0x6e, 0x7a, 0x61, 0x2f,
	0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x67, 0x61, 0x6e, 0x7a, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7b, 0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x6e, 0x7a, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x67, 0x61, 0x6e, 0x7a, 0x61, 0x2e, 0x50, 0x42, 0x45,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x74, 0x61, 0x6e, 0x7a, 0x61, 0x22, 0x14,
	0x0a, 0x12, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x16, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x19, 0x0a, 0x17, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47,
	0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x6e, 0x7a,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61,
	0x67, 0x61, 0x6e, 0x7a, 0x61, 0x2e, 0x50, 0x42, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x6e, 0x7a, 0x61, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xb4, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x46, 0x0a, 0x12, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73,
	0x74, 0x72, 0x61, 0x76, 0x61, 0x67, 0x61, 0x6e, 0x7a, 0x61, 0x2e, 0x50, 0x42, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x12, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2a, 0x91, 0x05, 0x0a, 0x11, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x23,
	0x0a, 0x1f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x58, 0x4d,
	0x4c, 0x10, 0x00, 0x12, 0x29, 0x0a, 0x25, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x01, 0x12, 0x24,
	0x0a, 0x20, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x46,
	0x4c, 0x49, 0x43, 0x54, 0x10, 0x03, 0x12, 0x24, 0x0a, 0x20, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x10, 0x04, 0x12, 0x28, 0x0a, 0x24,
	0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x56, 0x49, 0x4f, 0x4c, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x30, 0x0a, 0x2c, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x52, 0x45,
	0x4d, 0x4f, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x2a, 0x0a, 0x26, 0x53, 0x54, 0x52, 0x45,
	0x41, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f,
	0x55, 0x54, 0x10, 0x07, 0x12, 0x2f, 0x0a, 0x2b, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x55,
	0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x4e, 0x5a, 0x41, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x10, 0x08, 0x12, 0x2b, 0x0a, 0x27, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x10, 0x09, 0x12, 0x26, 0x0a, 0x22, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x55,
	0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x2b, 0x0a, 0x27, 0x53, 0x54,
	0x52, 0x45, 0x41, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f,
	0x4e, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54,
	0x52, 0x41, 0x49, 0x4e, 0x54, 0x10, 0x0b, 0x12, 0x27, 0x0a, 0x23, 0x53, 0x54, 0x52, 0x45, 0x41,
	0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53,
	0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x53, 0x48, 0x55, 0x54, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x0c,
	0x12, 0x2b, 0x0a, 0x27, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x2d, 0x0a,
	0x29, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0e, 0x32, 0xac, 0x01, 0x0a,
	0x0b, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x05,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x12, 0x22, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x61, 0x0a, 0x0f, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x4e,
	0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cluster_proto_rawDescOnce sync.Once
	file_cluster_proto_rawDescData = file_cluster_proto_rawDesc
)

func file_cluster_proto_rawDescGZIP() []byte {
	file_cluster_proto_rawDescOnce.Do(func() {
		file_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_cluster_proto_rawDescData)
	})
	return file_cluster_proto_rawDescData
}

var file_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cluster_proto_goTypes = []interface{}{
	(StreamErrorReason)(0),          // 0: cluster.v1.StreamErrorReason
	(*LocalRouteRequest)(nil),       // 1: cluster.v1.LocalRouteRequest
	(*LocalRouteResponse)(nil),      // 2: cluster.v1.LocalRouteResponse
	(*LocalDisconnectRequest)(nil),  // 3: cluster.v1.LocalDisconnectRequest
	(*LocalDisconnectResponse)(nil), // 4: cluster.v1.LocalDisconnectResponse
	(*ComponentRouteRequest)(nil),   // 5: cluster.v1.ComponentRouteRequest
	(*ComponentRouteResponse)(nil),  // 6: cluster.v1.ComponentRouteResponse
	(*StreamError)(nil),             // 7: cluster.v1.StreamError
	(*stravaganza.PBElement)(nil),   // 8: stravaganza.PBElement
}
var file_cluster_proto_depIdxs = []int32{
	8, // 0: cluster.v1.LocalRouteRequest.stanza:type_name -> stravaganza.PBElement
	7, // 1: cluster.v1.LocalDisconnectRequest.stream_error:type_name -> cluster.v1.StreamError
	8, // 2: cluster.v1.ComponentRouteRequest.stanza:type_name -> stravaganza.PBElement
	0, // 3: cluster.v1.StreamError.reason:type_name -> cluster.v1.StreamErrorReason
	8, // 4: cluster.v1.StreamError.applicationElement:type_name -> stravaganza.PBElement
	1, // 5: cluster.v1.LocalRouter.Route:input_type -> cluster.v1.LocalRouteRequest
	3, // 6: cluster.v1.LocalRouter.Disconnect:input_type -> cluster.v1.LocalDisconnectRequest
	5, // 7: cluster.v1.ComponentRouter.Route:input_type -> cluster.v1.ComponentRouteRequest
	2, // 8: cluster.v1.LocalRouter.Route:output_type -> cluster.v1.LocalRouteResponse
	4, // 9: cluster.v1.LocalRouter.Disconnect:output_type -> cluster.v1.LocalDisconnectResponse
	6, // 10: cluster.v1.ComponentRouter.Route:output_type -> cluster.v1.ComponentRouteResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cluster_proto_init() }
func file_cluster_proto_init() {
	if File_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalRouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalRouteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalDisconnectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalDisconnectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentRouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentRouteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cluster_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cluster_proto_goTypes,
		DependencyIndexes: file_cluster_proto_depIdxs,
		EnumInfos:         file_cluster_proto_enumTypes,
		MessageInfos:      file_cluster_proto_msgTypes,
	}.Build()
	File_cluster_proto = out.File
	file_cluster_proto_rawDesc = nil
	file_cluster_proto_goTypes = nil
	file_cluster_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LocalRouterClient is the client API for LocalRouter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LocalRouterClient interface {
	// Route method routes a stanza to a local available resource.
	Route(ctx context.Context, in *LocalRouteRequest, opts ...grpc.CallOption) (*LocalRouteResponse, error)
	// Disconnect performs disconnection over a local available resource.
	Disconnect(ctx context.Context, in *LocalDisconnectRequest, opts ...grpc.CallOption) (*LocalDisconnectResponse, error)
}

type localRouterClient struct {
	cc *grpc.ClientConn
}

func NewLocalRouterClient(cc *grpc.ClientConn) LocalRouterClient {
	return &localRouterClient{cc}
}

func (c *localRouterClient) Route(ctx context.Context, in *LocalRouteRequest, opts ...grpc.CallOption) (*LocalRouteResponse, error) {
	out := new(LocalRouteResponse)
	err := c.cc.Invoke(ctx, "/cluster.v1.LocalRouter/Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localRouterClient) Disconnect(ctx context.Context, in *LocalDisconnectRequest, opts ...grpc.CallOption) (*LocalDisconnectResponse, error) {
	out := new(LocalDisconnectResponse)
	err := c.cc.Invoke(ctx, "/cluster.v1.LocalRouter/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocalRouterServer is the server API for LocalRouter service.
type LocalRouterServer interface {
	// Route method routes a stanza to a local available resource.
	Route(context.Context, *LocalRouteRequest) (*LocalRouteResponse, error)
	// Disconnect performs disconnection over a local available resource.
	Disconnect(context.Context, *LocalDisconnectRequest) (*LocalDisconnectResponse, error)
}

// UnimplementedLocalRouterServer can be embedded to have forward compatible implementations.
type UnimplementedLocalRouterServer struct {
}

func (*UnimplementedLocalRouterServer) Route(context.Context, *LocalRouteRequest) (*LocalRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (*UnimplementedLocalRouterServer) Disconnect(context.Context, *LocalDisconnectRequest) (*LocalDisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}

func RegisterLocalRouterServer(s *grpc.Server, srv LocalRouterServer) {
	s.RegisterService(&_LocalRouter_serviceDesc, srv)
}

func _LocalRouter_Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocalRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalRouterServer).Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.v1.LocalRouter/Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalRouterServer).Route(ctx, req.(*LocalRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocalRouter_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocalDisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalRouterServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.v1.LocalRouter/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalRouterServer).Disconnect(ctx, req.(*LocalDisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LocalRouter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cluster.v1.LocalRouter",
	HandlerType: (*LocalRouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Route",
			Handler:    _LocalRouter_Route_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _LocalRouter_Disconnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster.proto",
}

// ComponentRouterClient is the client API for ComponentRouter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComponentRouterClient interface {
	// Route method routes a stanza to a local registered component.
	Route(ctx context.Context, in *ComponentRouteRequest, opts ...grpc.CallOption) (*ComponentRouteResponse, error)
}

type componentRouterClient struct {
	cc *grpc.ClientConn
}

func NewComponentRouterClient(cc *grpc.ClientConn) ComponentRouterClient {
	return &componentRouterClient{cc}
}

func (c *componentRouterClient) Route(ctx context.Context, in *ComponentRouteRequest, opts ...grpc.CallOption) (*ComponentRouteResponse, error) {
	out := new(ComponentRouteResponse)
	err := c.cc.Invoke(ctx, "/cluster.v1.ComponentRouter/Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComponentRouterServer is the server API for ComponentRouter service.
type ComponentRouterServer interface {
	// Route method routes a stanza to a local registered component.
	Route(context.Context, *ComponentRouteRequest) (*ComponentRouteResponse, error)
}

// UnimplementedComponentRouterServer can be embedded to have forward compatible implementations.
type UnimplementedComponentRouterServer struct {
}

func (*UnimplementedComponentRouterServer) Route(context.Context, *ComponentRouteRequest) (*ComponentRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Route not implemented")
}

func RegisterComponentRouterServer(s *grpc.Server, srv ComponentRouterServer) {
	s.RegisterService(&_ComponentRouter_serviceDesc, srv)
}

func _ComponentRouter_Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComponentRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentRouterServer).Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.v1.ComponentRouter/Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentRouterServer).Route(ctx, req.(*ComponentRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ComponentRouter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cluster.v1.ComponentRouter",
	HandlerType: (*ComponentRouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Route",
			Handler:    _ComponentRouter_Route_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster.proto",
}