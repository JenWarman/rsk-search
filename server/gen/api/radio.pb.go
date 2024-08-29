// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: radio.proto

package api

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RadioState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentEpisode     *CurrentRadioEpisode `protobuf:"bytes,1,opt,name=current_episode,json=currentEpisode,proto3" json:"current_episode,omitempty"`
	CurrentTimestampMs int32                `protobuf:"varint,3,opt,name=current_timestamp_ms,json=currentTimestampMs,proto3" json:"current_timestamp_ms,omitempty"`
}

func (x *RadioState) Reset() {
	*x = RadioState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_radio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioState) ProtoMessage() {}

func (x *RadioState) ProtoReflect() protoreflect.Message {
	mi := &file_radio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioState.ProtoReflect.Descriptor instead.
func (*RadioState) Descriptor() ([]byte, []int) {
	return file_radio_proto_rawDescGZIP(), []int{0}
}

func (x *RadioState) GetCurrentEpisode() *CurrentRadioEpisode {
	if x != nil {
		return x.CurrentEpisode
	}
	return nil
}

func (x *RadioState) GetCurrentTimestampMs() int32 {
	if x != nil {
		return x.CurrentTimestampMs
	}
	return 0
}

type CurrentRadioEpisode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId   string `protobuf:"bytes,1,opt,name=shortId,proto3" json:"shortId,omitempty"`
	StartedAt string `protobuf:"bytes,2,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
}

func (x *CurrentRadioEpisode) Reset() {
	*x = CurrentRadioEpisode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_radio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentRadioEpisode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentRadioEpisode) ProtoMessage() {}

func (x *CurrentRadioEpisode) ProtoReflect() protoreflect.Message {
	mi := &file_radio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentRadioEpisode.ProtoReflect.Descriptor instead.
func (*CurrentRadioEpisode) Descriptor() ([]byte, []int) {
	return file_radio_proto_rawDescGZIP(), []int{1}
}

func (x *CurrentRadioEpisode) GetShortId() string {
	if x != nil {
		return x.ShortId
	}
	return ""
}

func (x *CurrentRadioEpisode) GetStartedAt() string {
	if x != nil {
		return x.StartedAt
	}
	return ""
}

type PutRadioStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentEpisode     *CurrentRadioEpisode `protobuf:"bytes,1,opt,name=current_episode,json=currentEpisode,proto3" json:"current_episode,omitempty"`
	CurrentTimestampMs int32                `protobuf:"varint,3,opt,name=current_timestamp_ms,json=currentTimestampMs,proto3" json:"current_timestamp_ms,omitempty"`
}

func (x *PutRadioStateRequest) Reset() {
	*x = PutRadioStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_radio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRadioStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRadioStateRequest) ProtoMessage() {}

func (x *PutRadioStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_radio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRadioStateRequest.ProtoReflect.Descriptor instead.
func (*PutRadioStateRequest) Descriptor() ([]byte, []int) {
	return file_radio_proto_rawDescGZIP(), []int{2}
}

func (x *PutRadioStateRequest) GetCurrentEpisode() *CurrentRadioEpisode {
	if x != nil {
		return x.CurrentEpisode
	}
	return nil
}

func (x *PutRadioStateRequest) GetCurrentTimestampMs() int32 {
	if x != nil {
		return x.CurrentTimestampMs
	}
	return 0
}

type NextRadioEpisode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId string `protobuf:"bytes,1,opt,name=shortId,proto3" json:"shortId,omitempty"`
}

func (x *NextRadioEpisode) Reset() {
	*x = NextRadioEpisode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_radio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextRadioEpisode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextRadioEpisode) ProtoMessage() {}

func (x *NextRadioEpisode) ProtoReflect() protoreflect.Message {
	mi := &file_radio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextRadioEpisode.ProtoReflect.Descriptor instead.
func (*NextRadioEpisode) Descriptor() ([]byte, []int) {
	return file_radio_proto_rawDescGZIP(), []int{3}
}

func (x *NextRadioEpisode) GetShortId() string {
	if x != nil {
		return x.ShortId
	}
	return ""
}

var File_radio_proto protoreflect.FileDescriptor

var file_radio_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72,
	0x73, 0x6b, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01,
	0x0a, 0x0a, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x41, 0x0a, 0x0f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x73, 0x6b, 0x2e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x52,
	0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x12,
	0x30, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d,
	0x73, 0x22, 0x4e, 0x0a, 0x13, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x64, 0x69,
	0x6f, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x8b, 0x01, 0x0a, 0x14, 0x50, 0x75, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x73, 0x6b, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x52, 0x0e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a,
	0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x73, 0x22,
	0x2c, 0x0a, 0x10, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x45, 0x70, 0x69, 0x73,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x32, 0xb7, 0x03,
	0x0a, 0x0c, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89,
	0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x72, 0x73, 0x6b, 0x2e, 0x52,
	0x61, 0x64, 0x69, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x4f, 0x92, 0x41, 0x34, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x20, 0x47, 0x65, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x20,
	0x74, 0x6f, 0x20, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x2a, 0x08, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x61, 0x64, 0x69, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x72, 0x73, 0x6b, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x61,
	0x64, 0x69, 0x6f, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x22, 0x4a, 0x92, 0x41, 0x30, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1d, 0x47, 0x65, 0x74, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x6e, 0x65, 0x78, 0x74, 0x20, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x20, 0x74, 0x6f,
	0x20, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x2a, 0x07, 0x67, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x61, 0x64, 0x69,
	0x6f, 0x2f, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x8e, 0x01, 0x0a, 0x0d, 0x50, 0x75, 0x74, 0x52, 0x61,
	0x64, 0x69, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x72, 0x73, 0x6b, 0x2e, 0x50,
	0x75, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4a, 0x92, 0x41, 0x2c,
	0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x2a, 0x08, 0x70, 0x75, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x3a, 0x01, 0x2a, 0x1a, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x61, 0x64, 0x69,
	0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x88, 0x01, 0x92, 0x41, 0x57, 0x12, 0x05, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x01, 0x72, 0x4b, 0x0a, 0x32, 0x52, 0x61, 0x64, 0x69, 0x6f,
	0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x73, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x72, 0x61, 0x64, 0x69, 0x6f, 0x20, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x12, 0x15, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x6d, 0x70, 0x74, 0x6f, 0x6e,
	0x2e, 0x63, 0x6f, 0x6d, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x61, 0x72, 0x6d, 0x61, 0x6e, 0x73, 0x2f, 0x72, 0x73, 0x6b, 0x2d, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_radio_proto_rawDescOnce sync.Once
	file_radio_proto_rawDescData = file_radio_proto_rawDesc
)

func file_radio_proto_rawDescGZIP() []byte {
	file_radio_proto_rawDescOnce.Do(func() {
		file_radio_proto_rawDescData = protoimpl.X.CompressGZIP(file_radio_proto_rawDescData)
	})
	return file_radio_proto_rawDescData
}

var file_radio_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_radio_proto_goTypes = []interface{}{
	(*RadioState)(nil),           // 0: rsk.RadioState
	(*CurrentRadioEpisode)(nil),  // 1: rsk.CurrentRadioEpisode
	(*PutRadioStateRequest)(nil), // 2: rsk.PutRadioStateRequest
	(*NextRadioEpisode)(nil),     // 3: rsk.NextRadioEpisode
	(*emptypb.Empty)(nil),        // 4: google.protobuf.Empty
}
var file_radio_proto_depIdxs = []int32{
	1, // 0: rsk.RadioState.current_episode:type_name -> rsk.CurrentRadioEpisode
	1, // 1: rsk.PutRadioStateRequest.current_episode:type_name -> rsk.CurrentRadioEpisode
	4, // 2: rsk.RadioService.GetRadioState:input_type -> google.protobuf.Empty
	4, // 3: rsk.RadioService.GetRadioNext:input_type -> google.protobuf.Empty
	2, // 4: rsk.RadioService.PutRadioState:input_type -> rsk.PutRadioStateRequest
	0, // 5: rsk.RadioService.GetRadioState:output_type -> rsk.RadioState
	3, // 6: rsk.RadioService.GetRadioNext:output_type -> rsk.NextRadioEpisode
	4, // 7: rsk.RadioService.PutRadioState:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_radio_proto_init() }
func file_radio_proto_init() {
	if File_radio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_radio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioState); i {
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
		file_radio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentRadioEpisode); i {
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
		file_radio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRadioStateRequest); i {
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
		file_radio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextRadioEpisode); i {
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
			RawDescriptor: file_radio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_radio_proto_goTypes,
		DependencyIndexes: file_radio_proto_depIdxs,
		MessageInfos:      file_radio_proto_msgTypes,
	}.Build()
	File_radio_proto = out.File
	file_radio_proto_rawDesc = nil
	file_radio_proto_goTypes = nil
	file_radio_proto_depIdxs = nil
}
