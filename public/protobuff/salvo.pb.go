// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: salvo.proto

package protobuff

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start        *timestamp.Timestamp `protobuf:"bytes,1,opt,name=Start,proto3" json:"Start,omitempty"`
	End          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=End,proto3" json:"End,omitempty"`
	Success      bool                 `protobuf:"varint,3,opt,name=Success,proto3" json:"Success,omitempty"`
	StatusCode   int64                `protobuf:"varint,5,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	ResponseBody string               `protobuf:"bytes,6,opt,name=ResponseBody,proto3" json:"ResponseBody,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_salvo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_salvo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_salvo_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetStart() *timestamp.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Result) GetEnd() *timestamp.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *Result) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Result) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Result) GetResponseBody() string {
	if x != nil {
		return x.ResponseBody
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCount int32     `protobuf:"varint,7,opt,name=ErrorCount,proto3" json:"ErrorCount,omitempty"`
	Result     []*Result `protobuf:"bytes,8,rep,name=Result,proto3" json:"Result,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_salvo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_salvo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_salvo_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetErrorCount() int32 {
	if x != nil {
		return x.ErrorCount
	}
	return 0
}

func (x *Response) GetResult() []*Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path         string  `protobuf:"bytes,9,opt,name=Path,proto3" json:"Path,omitempty"`
	Users        int64   `protobuf:"varint,10,opt,name=Users,proto3" json:"Users,omitempty"`
	Timeout      int64   `protobuf:"varint,11,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	SuccessCodes []int64 `protobuf:"varint,12,rep,packed,name=SuccessCodes,proto3" json:"SuccessCodes,omitempty"`
	Rate         float32 `protobuf:"fixed32,13,opt,name=Rate,proto3" json:"Rate,omitempty"`
	Time         int64   `protobuf:"varint,14,opt,name=Time,proto3" json:"Time,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_salvo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_salvo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_salvo_proto_rawDescGZIP(), []int{2}
}

func (x *Options) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Options) GetUsers() int64 {
	if x != nil {
		return x.Users
	}
	return 0
}

func (x *Options) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Options) GetSuccessCodes() []int64 {
	if x != nil {
		return x.SuccessCodes
	}
	return nil
}

func (x *Options) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *Options) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_salvo_proto protoreflect.FileDescriptor

var file_salvo_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x76, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6,
	0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x45,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x45, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x6f, 0x64, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x4b, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0c, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x61, 0x74, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x52, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65,
	0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_salvo_proto_rawDescOnce sync.Once
	file_salvo_proto_rawDescData = file_salvo_proto_rawDesc
)

func file_salvo_proto_rawDescGZIP() []byte {
	file_salvo_proto_rawDescOnce.Do(func() {
		file_salvo_proto_rawDescData = protoimpl.X.CompressGZIP(file_salvo_proto_rawDescData)
	})
	return file_salvo_proto_rawDescData
}

var file_salvo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_salvo_proto_goTypes = []interface{}{
	(*Result)(nil),              // 0: Result
	(*Response)(nil),            // 1: Response
	(*Options)(nil),             // 2: Options
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_salvo_proto_depIdxs = []int32{
	3, // 0: Result.Start:type_name -> google.protobuf.Timestamp
	3, // 1: Result.End:type_name -> google.protobuf.Timestamp
	0, // 2: Response.Result:type_name -> Result
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_salvo_proto_init() }
func file_salvo_proto_init() {
	if File_salvo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_salvo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_salvo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_salvo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
			RawDescriptor: file_salvo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_salvo_proto_goTypes,
		DependencyIndexes: file_salvo_proto_depIdxs,
		MessageInfos:      file_salvo_proto_msgTypes,
	}.Build()
	File_salvo_proto = out.File
	file_salvo_proto_rawDesc = nil
	file_salvo_proto_goTypes = nil
	file_salvo_proto_depIdxs = nil
}
