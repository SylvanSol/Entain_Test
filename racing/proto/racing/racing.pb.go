// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: racing.proto

package racing

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListRacesRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Filter        *ListRacesRequestFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRacesRequest) Reset() {
	*x = ListRacesRequest{}
	mi := &file_racing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesRequest) ProtoMessage() {}

func (x *ListRacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_racing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesRequest.ProtoReflect.Descriptor instead.
func (*ListRacesRequest) Descriptor() ([]byte, []int) {
	return file_racing_proto_rawDescGZIP(), []int{0}
}

func (x *ListRacesRequest) GetFilter() *ListRacesRequestFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

// Response to ListRaces call.
type ListRacesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Races         []*Race                `protobuf:"bytes,1,rep,name=races,proto3" json:"races,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRacesResponse) Reset() {
	*x = ListRacesResponse{}
	mi := &file_racing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesResponse) ProtoMessage() {}

func (x *ListRacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_racing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesResponse.ProtoReflect.Descriptor instead.
func (*ListRacesResponse) Descriptor() ([]byte, []int) {
	return file_racing_proto_rawDescGZIP(), []int{1}
}

func (x *ListRacesResponse) GetRaces() []*Race {
	if x != nil {
		return x.Races
	}
	return nil
}

// Filter for listing races.
type ListRacesRequestFilter struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	MeetingIds []int64                `protobuf:"varint,1,rep,packed,name=meeting_ids,json=meetingIds,proto3" json:"meeting_ids,omitempty"`
	// Add Visibility Filter
	OnlyVisible   bool `protobuf:"varint,2,opt,name=only_visible,json=onlyVisible,proto3" json:"only_visible,omitempty"` // If true only returns races where visible = true
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRacesRequestFilter) Reset() {
	*x = ListRacesRequestFilter{}
	mi := &file_racing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRacesRequestFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesRequestFilter) ProtoMessage() {}

func (x *ListRacesRequestFilter) ProtoReflect() protoreflect.Message {
	mi := &file_racing_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesRequestFilter.ProtoReflect.Descriptor instead.
func (*ListRacesRequestFilter) Descriptor() ([]byte, []int) {
	return file_racing_proto_rawDescGZIP(), []int{2}
}

func (x *ListRacesRequestFilter) GetMeetingIds() []int64 {
	if x != nil {
		return x.MeetingIds
	}
	return nil
}

func (x *ListRacesRequestFilter) GetOnlyVisible() bool {
	if x != nil {
		return x.OnlyVisible
	}
	return false
}

// A race resource.
type Race struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID represents a unique identifier for the race.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// MeetingID represents a unique identifier for the races meeting.
	MeetingId int64 `protobuf:"varint,2,opt,name=meeting_id,json=meetingId,proto3" json:"meeting_id,omitempty"`
	// Name is the official name given to the race.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Number represents the number of the race.
	Number int64 `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	// Visible represents whether or not the race is visible.
	Visible bool `protobuf:"varint,5,opt,name=visible,proto3" json:"visible,omitempty"`
	// AdvertisedStartTime is the time the race is advertised to run.
	AdvertisedStartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=advertised_start_time,json=advertisedStartTime,proto3" json:"advertised_start_time,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *Race) Reset() {
	*x = Race{}
	mi := &file_racing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Race) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Race) ProtoMessage() {}

func (x *Race) ProtoReflect() protoreflect.Message {
	mi := &file_racing_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Race.ProtoReflect.Descriptor instead.
func (*Race) Descriptor() ([]byte, []int) {
	return file_racing_proto_rawDescGZIP(), []int{3}
}

func (x *Race) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Race) GetMeetingId() int64 {
	if x != nil {
		return x.MeetingId
	}
	return 0
}

func (x *Race) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Race) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Race) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Race) GetAdvertisedStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AdvertisedStartTime
	}
	return nil
}

var File_racing_proto protoreflect.FileDescriptor

const file_racing_proto_rawDesc = "" +
	"\n" +
	"\fracing.proto\x12\x06racing\x1a\x1fgoogle/protobuf/timestamp.proto\"J\n" +
	"\x10ListRacesRequest\x126\n" +
	"\x06filter\x18\x01 \x01(\v2\x1e.racing.ListRacesRequestFilterR\x06filter\"7\n" +
	"\x11ListRacesResponse\x12\"\n" +
	"\x05races\x18\x01 \x03(\v2\f.racing.RaceR\x05races\"\\\n" +
	"\x16ListRacesRequestFilter\x12\x1f\n" +
	"\vmeeting_ids\x18\x01 \x03(\x03R\n" +
	"meetingIds\x12!\n" +
	"\fonly_visible\x18\x02 \x01(\bR\vonlyVisible\"\xcb\x01\n" +
	"\x04Race\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1d\n" +
	"\n" +
	"meeting_id\x18\x02 \x01(\x03R\tmeetingId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x16\n" +
	"\x06number\x18\x04 \x01(\x03R\x06number\x12\x18\n" +
	"\avisible\x18\x05 \x01(\bR\avisible\x12N\n" +
	"\x15advertised_start_time\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\x13advertisedStartTime2L\n" +
	"\x06Racing\x12B\n" +
	"\tListRaces\x12\x18.racing.ListRacesRequest\x1a\x19.racing.ListRacesResponse\"\x00B\tZ\a/racingb\x06proto3"

var (
	file_racing_proto_rawDescOnce sync.Once
	file_racing_proto_rawDescData []byte
)

func file_racing_proto_rawDescGZIP() []byte {
	file_racing_proto_rawDescOnce.Do(func() {
		file_racing_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_racing_proto_rawDesc), len(file_racing_proto_rawDesc)))
	})
	return file_racing_proto_rawDescData
}

var file_racing_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_racing_proto_goTypes = []any{
	(*ListRacesRequest)(nil),       // 0: racing.ListRacesRequest
	(*ListRacesResponse)(nil),      // 1: racing.ListRacesResponse
	(*ListRacesRequestFilter)(nil), // 2: racing.ListRacesRequestFilter
	(*Race)(nil),                   // 3: racing.Race
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
}
var file_racing_proto_depIdxs = []int32{
	2, // 0: racing.ListRacesRequest.filter:type_name -> racing.ListRacesRequestFilter
	3, // 1: racing.ListRacesResponse.races:type_name -> racing.Race
	4, // 2: racing.Race.advertised_start_time:type_name -> google.protobuf.Timestamp
	0, // 3: racing.Racing.ListRaces:input_type -> racing.ListRacesRequest
	1, // 4: racing.Racing.ListRaces:output_type -> racing.ListRacesResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_racing_proto_init() }
func file_racing_proto_init() {
	if File_racing_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_racing_proto_rawDesc), len(file_racing_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_racing_proto_goTypes,
		DependencyIndexes: file_racing_proto_depIdxs,
		MessageInfos:      file_racing_proto_msgTypes,
	}.Build()
	File_racing_proto = out.File
	file_racing_proto_goTypes = nil
	file_racing_proto_depIdxs = nil
}
