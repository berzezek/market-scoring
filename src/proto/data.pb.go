// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: src/proto/data.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SellerId string `protobuf:"bytes,1,opt,name=sellerId,proto3" json:"sellerId,omitempty"`
}

func (x *DataRequest) Reset() {
	*x = DataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRequest) ProtoMessage() {}

func (x *DataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRequest.ProtoReflect.Descriptor instead.
func (*DataRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_data_proto_rawDescGZIP(), []int{0}
}

func (x *DataRequest) GetSellerId() string {
	if x != nil {
		return x.SellerId
	}
	return ""
}

type DataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message          string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Success          bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	ActiveProducts   int32                  `protobuf:"varint,3,opt,name=active_products,json=activeProducts,proto3" json:"active_products,omitempty"`
	RegistrationDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=registration_date,json=registrationDate,proto3" json:"registration_date,omitempty"`
	Turnover         float64                `protobuf:"fixed64,5,opt,name=turnover,proto3" json:"turnover,omitempty"`
	SalesLastMonth   int32                  `protobuf:"varint,6,opt,name=sales_last_month,json=salesLastMonth,proto3" json:"sales_last_month,omitempty"`
}

func (x *DataResponse) Reset() {
	*x = DataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataResponse) ProtoMessage() {}

func (x *DataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataResponse.ProtoReflect.Descriptor instead.
func (*DataResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_data_proto_rawDescGZIP(), []int{1}
}

func (x *DataResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DataResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DataResponse) GetActiveProducts() int32 {
	if x != nil {
		return x.ActiveProducts
	}
	return 0
}

func (x *DataResponse) GetRegistrationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.RegistrationDate
	}
	return nil
}

func (x *DataResponse) GetTurnover() float64 {
	if x != nil {
		return x.Turnover
	}
	return 0
}

func (x *DataResponse) GetSalesLastMonth() int32 {
	if x != nil {
		return x.SalesLastMonth
	}
	return 0
}

var File_src_proto_data_proto protoreflect.FileDescriptor

var file_src_proto_data_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a,
	0x0b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x22, 0xfa, 0x01, 0x0a, 0x0c, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x27, 0x0a,
	0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x47, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x75, 0x72, 0x6e, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x74, 0x75, 0x72, 0x6e, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x73,
	0x61, 0x6c, 0x65, 0x73, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x4c, 0x61, 0x73, 0x74,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x32, 0x43, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x73,
	0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_proto_data_proto_rawDescOnce sync.Once
	file_src_proto_data_proto_rawDescData = file_src_proto_data_proto_rawDesc
)

func file_src_proto_data_proto_rawDescGZIP() []byte {
	file_src_proto_data_proto_rawDescOnce.Do(func() {
		file_src_proto_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_data_proto_rawDescData)
	})
	return file_src_proto_data_proto_rawDescData
}

var file_src_proto_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_proto_data_proto_goTypes = []interface{}{
	(*DataRequest)(nil),           // 0: data.DataRequest
	(*DataResponse)(nil),          // 1: data.DataResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_src_proto_data_proto_depIdxs = []int32{
	2, // 0: data.DataResponse.registration_date:type_name -> google.protobuf.Timestamp
	0, // 1: data.DataService.ProcessData:input_type -> data.DataRequest
	1, // 2: data.DataService.ProcessData:output_type -> data.DataResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_src_proto_data_proto_init() }
func file_src_proto_data_proto_init() {
	if File_src_proto_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRequest); i {
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
		file_src_proto_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataResponse); i {
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
			RawDescriptor: file_src_proto_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_data_proto_goTypes,
		DependencyIndexes: file_src_proto_data_proto_depIdxs,
		MessageInfos:      file_src_proto_data_proto_msgTypes,
	}.Build()
	File_src_proto_data_proto = out.File
	file_src_proto_data_proto_rawDesc = nil
	file_src_proto_data_proto_goTypes = nil
	file_src_proto_data_proto_depIdxs = nil
}
