// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ProductService.proto

package ProductService

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddProductRequest_Classfication int32

const (
	AddProductRequest_FRUIT      AddProductRequest_Classfication = 0
	AddProductRequest_MEAT       AddProductRequest_Classfication = 1
	AddProductRequest_STAPLE     AddProductRequest_Classfication = 2
	AddProductRequest_TOILETRIES AddProductRequest_Classfication = 3
	AddProductRequest_DRESS      AddProductRequest_Classfication = 4
)

var AddProductRequest_Classfication_name = map[int32]string{
	0: "FRUIT",
	1: "MEAT",
	2: "STAPLE",
	3: "TOILETRIES",
	4: "DRESS",
}

var AddProductRequest_Classfication_value = map[string]int32{
	"FRUIT":      0,
	"MEAT":       1,
	"STAPLE":     2,
	"TOILETRIES": 3,
	"DRESS":      4,
}

func (x AddProductRequest_Classfication) String() string {
	return proto.EnumName(AddProductRequest_Classfication_name, int32(x))
}

func (AddProductRequest_Classfication) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a39275549ecd65e1, []int{0, 0}
}

// 请求/响应结构体定义
// 添加产品message
type AddProductRequest struct {
	ProductName          string                          `protobuf:"bytes,1,opt,name=productName,proto3" json:"productName,omitempty"`
	Classification       AddProductRequest_Classfication `protobuf:"varint,2,opt,name=classification,proto3,enum=AddProductRequest_Classfication" json:"classification,omitempty"`
	ManufacturerId       string                          `protobuf:"bytes,3,opt,name=manufacturerId,proto3" json:"manufacturerId,omitempty"`
	Weight               float64                         `protobuf:"fixed64,4,opt,name=weight,proto3" json:"weight,omitempty"`
	ProductionDate       int64                           `protobuf:"varint,5,opt,name=productionDate,proto3" json:"productionDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *AddProductRequest) Reset()         { *m = AddProductRequest{} }
func (m *AddProductRequest) String() string { return proto.CompactTextString(m) }
func (*AddProductRequest) ProtoMessage()    {}
func (*AddProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39275549ecd65e1, []int{0}
}

func (m *AddProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProductRequest.Unmarshal(m, b)
}
func (m *AddProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProductRequest.Marshal(b, m, deterministic)
}
func (m *AddProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProductRequest.Merge(m, src)
}
func (m *AddProductRequest) XXX_Size() int {
	return xxx_messageInfo_AddProductRequest.Size(m)
}
func (m *AddProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddProductRequest proto.InternalMessageInfo

func (m *AddProductRequest) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *AddProductRequest) GetClassification() AddProductRequest_Classfication {
	if m != nil {
		return m.Classification
	}
	return AddProductRequest_FRUIT
}

func (m *AddProductRequest) GetManufacturerId() string {
	if m != nil {
		return m.ManufacturerId
	}
	return ""
}

func (m *AddProductRequest) GetWeight() float64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *AddProductRequest) GetProductionDate() int64 {
	if m != nil {
		return m.ProductionDate
	}
	return 0
}

// 添加产品，服务端响应message
type AddProductResponse struct {
	ProductId            string   `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddProductResponse) Reset()         { *m = AddProductResponse{} }
func (m *AddProductResponse) String() string { return proto.CompactTextString(m) }
func (*AddProductResponse) ProtoMessage()    {}
func (*AddProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39275549ecd65e1, []int{1}
}

func (m *AddProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProductResponse.Unmarshal(m, b)
}
func (m *AddProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProductResponse.Marshal(b, m, deterministic)
}
func (m *AddProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProductResponse.Merge(m, src)
}
func (m *AddProductResponse) XXX_Size() int {
	return xxx_messageInfo_AddProductResponse.Size(m)
}
func (m *AddProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddProductResponse proto.InternalMessageInfo

func (m *AddProductResponse) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *AddProductResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// 删除产品message
type DeleteProductRequest struct {
	ProductId            string   `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProductRequest) Reset()         { *m = DeleteProductRequest{} }
func (m *DeleteProductRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProductRequest) ProtoMessage()    {}
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39275549ecd65e1, []int{2}
}

func (m *DeleteProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProductRequest.Unmarshal(m, b)
}
func (m *DeleteProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProductRequest.Marshal(b, m, deterministic)
}
func (m *DeleteProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProductRequest.Merge(m, src)
}
func (m *DeleteProductRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProductRequest.Size(m)
}
func (m *DeleteProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProductRequest proto.InternalMessageInfo

func (m *DeleteProductRequest) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

type QueryProductRequest struct {
	ProductId            string   `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryProductRequest) Reset()         { *m = QueryProductRequest{} }
func (m *QueryProductRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProductRequest) ProtoMessage()    {}
func (*QueryProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39275549ecd65e1, []int{3}
}

func (m *QueryProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryProductRequest.Unmarshal(m, b)
}
func (m *QueryProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryProductRequest.Marshal(b, m, deterministic)
}
func (m *QueryProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProductRequest.Merge(m, src)
}
func (m *QueryProductRequest) XXX_Size() int {
	return xxx_messageInfo_QueryProductRequest.Size(m)
}
func (m *QueryProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProductRequest proto.InternalMessageInfo

func (m *QueryProductRequest) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

// 单产品详情message
type ProductInfoResponse struct {
	ProductName          string   `protobuf:"bytes,1,opt,name=productName,proto3" json:"productName,omitempty"`
	ProductId            string   `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty"`
	ManufacturerId       string   `protobuf:"bytes,3,opt,name=manufacturerId,proto3" json:"manufacturerId,omitempty"`
	Weight               float64  `protobuf:"fixed64,4,opt,name=weight,proto3" json:"weight,omitempty"`
	ProductionDate       int64    `protobuf:"varint,5,opt,name=productionDate,proto3" json:"productionDate,omitempty"`
	ImportDate           int64    `protobuf:"varint,6,opt,name=importDate,proto3" json:"importDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductInfoResponse) Reset()         { *m = ProductInfoResponse{} }
func (m *ProductInfoResponse) String() string { return proto.CompactTextString(m) }
func (*ProductInfoResponse) ProtoMessage()    {}
func (*ProductInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39275549ecd65e1, []int{4}
}

func (m *ProductInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductInfoResponse.Unmarshal(m, b)
}
func (m *ProductInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductInfoResponse.Marshal(b, m, deterministic)
}
func (m *ProductInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductInfoResponse.Merge(m, src)
}
func (m *ProductInfoResponse) XXX_Size() int {
	return xxx_messageInfo_ProductInfoResponse.Size(m)
}
func (m *ProductInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductInfoResponse proto.InternalMessageInfo

func (m *ProductInfoResponse) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *ProductInfoResponse) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *ProductInfoResponse) GetManufacturerId() string {
	if m != nil {
		return m.ManufacturerId
	}
	return ""
}

func (m *ProductInfoResponse) GetWeight() float64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *ProductInfoResponse) GetProductionDate() int64 {
	if m != nil {
		return m.ProductionDate
	}
	return 0
}

func (m *ProductInfoResponse) GetImportDate() int64 {
	if m != nil {
		return m.ImportDate
	}
	return 0
}

type ProductsInfoResponse struct {
	Infos                []*ProductInfoResponse `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ProductsInfoResponse) Reset()         { *m = ProductsInfoResponse{} }
func (m *ProductsInfoResponse) String() string { return proto.CompactTextString(m) }
func (*ProductsInfoResponse) ProtoMessage()    {}
func (*ProductsInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39275549ecd65e1, []int{5}
}

func (m *ProductsInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductsInfoResponse.Unmarshal(m, b)
}
func (m *ProductsInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductsInfoResponse.Marshal(b, m, deterministic)
}
func (m *ProductsInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductsInfoResponse.Merge(m, src)
}
func (m *ProductsInfoResponse) XXX_Size() int {
	return xxx_messageInfo_ProductsInfoResponse.Size(m)
}
func (m *ProductsInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductsInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductsInfoResponse proto.InternalMessageInfo

func (m *ProductsInfoResponse) GetInfos() []*ProductInfoResponse {
	if m != nil {
		return m.Infos
	}
	return nil
}

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39275549ecd65e1, []int{6}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39275549ecd65e1, []int{7}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("AddProductRequest_Classfication", AddProductRequest_Classfication_name, AddProductRequest_Classfication_value)
	proto.RegisterType((*AddProductRequest)(nil), "AddProductRequest")
	proto.RegisterType((*AddProductResponse)(nil), "AddProductResponse")
	proto.RegisterType((*DeleteProductRequest)(nil), "DeleteProductRequest")
	proto.RegisterType((*QueryProductRequest)(nil), "QueryProductRequest")
	proto.RegisterType((*ProductInfoResponse)(nil), "ProductInfoResponse")
	proto.RegisterType((*ProductsInfoResponse)(nil), "ProductsInfoResponse")
	proto.RegisterType((*EmptyRequest)(nil), "EmptyRequest")
	proto.RegisterType((*EmptyResponse)(nil), "EmptyResponse")
}

func init() { proto.RegisterFile("ProductService.proto", fileDescriptor_a39275549ecd65e1) }

var fileDescriptor_a39275549ecd65e1 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x5d, 0x8f, 0x93, 0x40,
	0x14, 0x05, 0xfa, 0xa1, 0xbd, 0x6b, 0x91, 0xbd, 0xed, 0x1a, 0xd2, 0x18, 0x43, 0xe6, 0xc1, 0x10,
	0x1f, 0xe6, 0xa1, 0x6b, 0xa2, 0x89, 0x2f, 0x56, 0x8b, 0x91, 0x58, 0x75, 0x1d, 0xf0, 0x07, 0x20,
	0x4c, 0x57, 0x92, 0xe5, 0x43, 0x66, 0xd0, 0xec, 0xab, 0xbf, 0xc9, 0xbf, 0xe3, 0x7f, 0x31, 0xa5,
	0xec, 0x2e, 0x50, 0xa2, 0xf1, 0x65, 0x1f, 0xe7, 0x70, 0xee, 0x9d, 0x7b, 0xce, 0x3d, 0x03, 0xcc,
	0xcf, 0x8a, 0x2c, 0x2a, 0x43, 0xe9, 0xf1, 0xe2, 0x7b, 0x1c, 0x72, 0x9a, 0x17, 0x99, 0xcc, 0xc8,
	0x2f, 0x0d, 0x8e, 0x57, 0x51, 0x54, 0x7f, 0x63, 0xfc, 0x5b, 0xc9, 0x85, 0x44, 0x0b, 0x8e, 0xf2,
	0x3d, 0xf2, 0x21, 0x48, 0xb8, 0xa9, 0x5a, 0xaa, 0x3d, 0x61, 0x4d, 0x08, 0xdf, 0x82, 0x1e, 0x5e,
	0x04, 0x42, 0xc4, 0xdb, 0x38, 0x0c, 0x64, 0x9c, 0xa5, 0xa6, 0x66, 0xa9, 0xb6, 0xbe, 0xb4, 0xe8,
	0x41, 0x37, 0xfa, 0x7a, 0x47, 0xbc, 0xe2, 0xb1, 0x4e, 0x1d, 0x3e, 0x06, 0x3d, 0x09, 0xd2, 0x72,
	0x1b, 0x84, 0xb2, 0x2c, 0x78, 0xe1, 0x46, 0xe6, 0xa0, 0xba, 0xae, 0x83, 0xe2, 0x03, 0x18, 0xff,
	0xe0, 0xf1, 0xf9, 0x57, 0x69, 0x0e, 0x2d, 0xd5, 0x56, 0x59, 0x7d, 0xda, 0xd5, 0xd7, 0x83, 0xc5,
	0x59, 0xba, 0x0e, 0x24, 0x37, 0x47, 0x96, 0x6a, 0x0f, 0x58, 0x07, 0x25, 0xef, 0x60, 0xda, 0x1a,
	0x04, 0x27, 0x30, 0x7a, 0xc3, 0x3e, 0xbb, 0xbe, 0xa1, 0xe0, 0x5d, 0x18, 0xbe, 0x77, 0x56, 0xbe,
	0xa1, 0x22, 0xc0, 0xd8, 0xf3, 0x57, 0x67, 0x1b, 0xc7, 0xd0, 0x50, 0x07, 0xf0, 0x3f, 0xba, 0x1b,
	0xc7, 0x67, 0xae, 0xe3, 0x19, 0x83, 0x5d, 0xc1, 0x9a, 0x39, 0x9e, 0x67, 0x0c, 0xc9, 0x06, 0xb0,
	0xa9, 0x53, 0xe4, 0x59, 0x2a, 0x38, 0x3e, 0x84, 0x49, 0x7d, 0xa9, 0x1b, 0xd5, 0xa6, 0xdd, 0x00,
	0x68, 0xc2, 0x9d, 0x84, 0x0b, 0x11, 0x9c, 0xf3, 0xca, 0xab, 0x09, 0xbb, 0x3a, 0x92, 0xa7, 0x30,
	0x5f, 0xf3, 0x0b, 0x2e, 0x79, 0x67, 0x0d, 0x7f, 0xed, 0x47, 0x4e, 0x61, 0xf6, 0xa9, 0xe4, 0xc5,
	0xe5, 0x7f, 0x15, 0xfd, 0x56, 0x61, 0x56, 0x17, 0xb8, 0xe9, 0x36, 0xbb, 0x1e, 0xfd, 0xdf, 0x1b,
	0x6f, 0xf5, 0xd5, 0xba, 0xe2, 0x6e, 0x69, 0x8b, 0xf8, 0x08, 0x20, 0x4e, 0xf2, 0xac, 0x90, 0x15,
	0x67, 0x5c, 0x71, 0x1a, 0x08, 0x79, 0x75, 0x9d, 0x73, 0xd1, 0xd2, 0xf7, 0x04, 0x46, 0x71, 0xba,
	0xcd, 0x84, 0xa9, 0x5a, 0x03, 0xfb, 0x68, 0x39, 0xa7, 0x3d, 0x26, 0xb0, 0x3d, 0x85, 0xe8, 0x70,
	0xcf, 0x49, 0x72, 0x79, 0x59, 0x3b, 0x4a, 0xee, 0xc3, 0xb4, 0x3e, 0xef, 0x79, 0xcb, 0x9f, 0x1a,
	0xe8, 0xed, 0xd7, 0x84, 0xcf, 0x00, 0x6e, 0x02, 0x81, 0x78, 0xf8, 0x0a, 0x16, 0x33, 0x7a, 0x98,
	0x18, 0xa2, 0xe0, 0x73, 0x98, 0xb6, 0x76, 0x8f, 0x27, 0xb4, 0x2f, 0x0b, 0x0b, 0x9d, 0xb6, 0x66,
	0x20, 0x0a, 0xbe, 0x04, 0xa3, 0xb9, 0xff, 0x9d, 0x12, 0x9c, 0xd3, 0x9e, 0x48, 0x2c, 0x7a, 0xd5,
	0x12, 0x05, 0x5f, 0xc0, 0x71, 0x93, 0x5e, 0x39, 0x86, 0x53, 0xda, 0x14, 0xbf, 0x38, 0xa1, 0x7d,
	0x7e, 0x12, 0xe5, 0xcb, 0xb8, 0xfa, 0x81, 0x9c, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xb9,
	0xd6, 0x42, 0x58, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	// 添加产品
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error)
	// 删除产品
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// 根据产品Id查询产品详情
	QueryProductInfo(ctx context.Context, in *QueryProductRequest, opts ...grpc.CallOption) (*ProductInfoResponse, error)
	// 查询所有产品详情
	QueryProductsInfo(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ProductsInfoResponse, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error) {
	out := new(AddProductResponse)
	err := c.cc.Invoke(ctx, "/ProductService/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/ProductService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) QueryProductInfo(ctx context.Context, in *QueryProductRequest, opts ...grpc.CallOption) (*ProductInfoResponse, error) {
	out := new(ProductInfoResponse)
	err := c.cc.Invoke(ctx, "/ProductService/QueryProductInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) QueryProductsInfo(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ProductsInfoResponse, error) {
	out := new(ProductsInfoResponse)
	err := c.cc.Invoke(ctx, "/ProductService/QueryProductsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	// 添加产品
	AddProduct(context.Context, *AddProductRequest) (*AddProductResponse, error)
	// 删除产品
	DeleteProduct(context.Context, *DeleteProductRequest) (*EmptyResponse, error)
	// 根据产品Id查询产品详情
	QueryProductInfo(context.Context, *QueryProductRequest) (*ProductInfoResponse, error)
	// 查询所有产品详情
	QueryProductsInfo(context.Context, *EmptyRequest) (*ProductsInfoResponse, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) AddProduct(ctx context.Context, req *AddProductRequest) (*AddProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (*UnimplementedProductServiceServer) DeleteProduct(ctx context.Context, req *DeleteProductRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (*UnimplementedProductServiceServer) QueryProductInfo(ctx context.Context, req *QueryProductRequest) (*ProductInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProductInfo not implemented")
}
func (*UnimplementedProductServiceServer) QueryProductsInfo(ctx context.Context, req *EmptyRequest) (*ProductsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProductsInfo not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_QueryProductInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).QueryProductInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/QueryProductInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).QueryProductInfo(ctx, req.(*QueryProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_QueryProductsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).QueryProductsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/QueryProductsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).QueryProductsInfo(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProduct",
			Handler:    _ProductService_AddProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "QueryProductInfo",
			Handler:    _ProductService_QueryProductInfo_Handler,
		},
		{
			MethodName: "QueryProductsInfo",
			Handler:    _ProductService_QueryProductsInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ProductService.proto",
}