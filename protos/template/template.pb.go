// Code generated by protoc-gen-go. DO NOT EDIT.
// source: template/template.proto

package template

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_dca67df6b60706ce, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type WorkflowTemplate struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data                 []byte               `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WorkflowTemplate) Reset()         { *m = WorkflowTemplate{} }
func (m *WorkflowTemplate) String() string { return proto.CompactTextString(m) }
func (*WorkflowTemplate) ProtoMessage()    {}
func (*WorkflowTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_dca67df6b60706ce, []int{1}
}

func (m *WorkflowTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowTemplate.Unmarshal(m, b)
}
func (m *WorkflowTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowTemplate.Marshal(b, m, deterministic)
}
func (m *WorkflowTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTemplate.Merge(m, src)
}
func (m *WorkflowTemplate) XXX_Size() int {
	return xxx_messageInfo_WorkflowTemplate.Size(m)
}
func (m *WorkflowTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTemplate proto.InternalMessageInfo

func (m *WorkflowTemplate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WorkflowTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowTemplate) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *WorkflowTemplate) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *WorkflowTemplate) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *WorkflowTemplate) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type CreateResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dca67df6b60706ce, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dca67df6b60706ce, []int{3}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "github.com.tinkerbell.tink.protos.template.Empty")
	proto.RegisterType((*WorkflowTemplate)(nil), "github.com.tinkerbell.tink.protos.template.WorkflowTemplate")
	proto.RegisterType((*CreateResponse)(nil), "github.com.tinkerbell.tink.protos.template.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "github.com.tinkerbell.tink.protos.template.GetRequest")
}

func init() {
	proto.RegisterFile("template/template.proto", fileDescriptor_dca67df6b60706ce)
}

var fileDescriptor_dca67df6b60706ce = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xd9, 0xbe, 0x6d, 0xdf, 0x76, 0xaa, 0x51, 0xf6, 0x62, 0x08, 0x1e, 0x42, 0x4e, 0xc5,
	0xc3, 0xaa, 0x15, 0x44, 0xc4, 0x8b, 0xff, 0xe8, 0xc5, 0x53, 0xa8, 0x08, 0xde, 0x92, 0x66, 0x5a,
	0x43, 0x93, 0x6e, 0xcc, 0x4e, 0x10, 0xf1, 0xa4, 0xde, 0xfc, 0x1a, 0x7e, 0x49, 0x8f, 0xd2, 0x6d,
	0x93, 0x42, 0x3d, 0xd4, 0x10, 0x6f, 0xc3, 0xec, 0xf3, 0x1b, 0x9e, 0x67, 0x76, 0x60, 0x87, 0x30,
	0x4e, 0x22, 0x8f, 0x70, 0x3f, 0x2f, 0x44, 0x92, 0x4a, 0x92, 0x7c, 0x6f, 0x1c, 0xd2, 0x43, 0xe6,
	0x8b, 0xa1, 0x8c, 0x05, 0x85, 0xd3, 0x09, 0xa6, 0x3e, 0x46, 0x91, 0x2e, 0xe7, 0x0a, 0x25, 0x72,
	0xc2, 0xda, 0xa2, 0x30, 0x46, 0x45, 0x5e, 0x9c, 0xcc, 0x9f, 0x9c, 0xff, 0xd0, 0xb8, 0x8e, 0x13,
	0x7a, 0x76, 0xbe, 0x18, 0x6c, 0xdf, 0xc9, 0x74, 0x32, 0x8a, 0xe4, 0xd3, 0x60, 0x21, 0xe7, 0x06,
	0xd4, 0xc2, 0xc0, 0x64, 0x36, 0xeb, 0xb6, 0xdd, 0x5a, 0x18, 0x70, 0x0e, 0xf5, 0xa9, 0x17, 0xa3,
	0x59, 0xd3, 0x1d, 0x5d, 0xcf, 0x7a, 0x81, 0x47, 0x9e, 0xf9, 0xcf, 0x66, 0xdd, 0x0d, 0x57, 0xd7,
	0xfc, 0x04, 0xda, 0xc3, 0x14, 0x3d, 0xc2, 0xe0, 0x9c, 0xcc, 0xba, 0xcd, 0xba, 0x9d, 0x9e, 0x25,
	0xc6, 0x52, 0x8e, 0xa3, 0x85, 0x69, 0x3f, 0x1b, 0x89, 0x41, 0x6e, 0xc5, 0x5d, 0x8a, 0x67, 0x64,
	0x96, 0x04, 0x0b, 0xb2, 0xb1, 0x9e, 0x2c, 0xc4, 0x33, 0x32, 0xc0, 0x08, 0xe7, 0x64, 0x73, 0x3d,
	0x59, 0x88, 0x1d, 0x1b, 0x8c, 0x4b, 0x6d, 0xc0, 0x45, 0x95, 0xc8, 0xa9, 0xfa, 0x91, 0xdb, 0xd9,
	0x05, 0xe8, 0x23, 0xb9, 0xf8, 0x98, 0xa1, 0xa2, 0xd5, 0xd7, 0xde, 0x67, 0x03, 0x5a, 0xc5, 0xca,
	0x3e, 0x58, 0x3e, 0xad, 0x68, 0x9d, 0x89, 0xdf, 0xff, 0x90, 0x58, 0xfd, 0x03, 0xeb, 0xb4, 0x0c,
	0xbd, 0x92, 0xe3, 0x9d, 0x41, 0xa7, 0x8f, 0x54, 0x38, 0x39, 0x2e, 0x33, 0x6b, 0x99, 0xd8, 0xaa,
	0x94, 0x80, 0xbf, 0x80, 0x71, 0xa5, 0x97, 0x5d, 0xd9, 0xc7, 0x61, 0x19, 0x4e, 0xdf, 0x35, 0x7f,
	0x63, 0xb0, 0x79, 0x13, 0xaa, 0x62, 0x07, 0x8a, 0x97, 0x1f, 0x52, 0x2d, 0xff, 0x01, 0xe3, 0xaf,
	0x0c, 0x8c, 0x5b, 0x7d, 0xa9, 0x7f, 0x74, 0x14, 0xe5, 0x33, 0x5c, 0xc0, 0x7d, 0x2b, 0xef, 0xf8,
	0x4d, 0x2d, 0x39, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x72, 0xba, 0x4c, 0x77, 0x54, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TemplateClient is the client API for Template service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TemplateClient interface {
	CreateTemplate(ctx context.Context, in *WorkflowTemplate, opts ...grpc.CallOption) (*CreateResponse, error)
	GetTemplate(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error)
	DeleteTemplate(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Empty, error)
	ListTemplates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Template_ListTemplatesClient, error)
	UpdateTemplate(ctx context.Context, in *WorkflowTemplate, opts ...grpc.CallOption) (*Empty, error)
}

type templateClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateClient(cc grpc.ClientConnInterface) TemplateClient {
	return &templateClient{cc}
}

func (c *templateClient) CreateTemplate(ctx context.Context, in *WorkflowTemplate, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.template.Template/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) GetTemplate(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error) {
	out := new(WorkflowTemplate)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.template.Template/GetTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) DeleteTemplate(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.template.Template/DeleteTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) ListTemplates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Template_ListTemplatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Template_serviceDesc.Streams[0], "/github.com.tinkerbell.tink.protos.template.Template/ListTemplates", opts...)
	if err != nil {
		return nil, err
	}
	x := &templateListTemplatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Template_ListTemplatesClient interface {
	Recv() (*WorkflowTemplate, error)
	grpc.ClientStream
}

type templateListTemplatesClient struct {
	grpc.ClientStream
}

func (x *templateListTemplatesClient) Recv() (*WorkflowTemplate, error) {
	m := new(WorkflowTemplate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *templateClient) UpdateTemplate(ctx context.Context, in *WorkflowTemplate, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.template.Template/UpdateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServer is the server API for Template service.
type TemplateServer interface {
	CreateTemplate(context.Context, *WorkflowTemplate) (*CreateResponse, error)
	GetTemplate(context.Context, *GetRequest) (*WorkflowTemplate, error)
	DeleteTemplate(context.Context, *GetRequest) (*Empty, error)
	ListTemplates(*Empty, Template_ListTemplatesServer) error
	UpdateTemplate(context.Context, *WorkflowTemplate) (*Empty, error)
}

// UnimplementedTemplateServer can be embedded to have forward compatible implementations.
type UnimplementedTemplateServer struct {
}

func (*UnimplementedTemplateServer) CreateTemplate(ctx context.Context, req *WorkflowTemplate) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}
func (*UnimplementedTemplateServer) GetTemplate(ctx context.Context, req *GetRequest) (*WorkflowTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplate not implemented")
}
func (*UnimplementedTemplateServer) DeleteTemplate(ctx context.Context, req *GetRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplate not implemented")
}
func (*UnimplementedTemplateServer) ListTemplates(req *Empty, srv Template_ListTemplatesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTemplates not implemented")
}
func (*UnimplementedTemplateServer) UpdateTemplate(ctx context.Context, req *WorkflowTemplate) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTemplate not implemented")
}

func RegisterTemplateServer(s *grpc.Server, srv TemplateServer) {
	s.RegisterService(&_Template_serviceDesc, srv)
}

func _Template_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.template.Template/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).CreateTemplate(ctx, req.(*WorkflowTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.template.Template/GetTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).GetTemplate(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_DeleteTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).DeleteTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.template.Template/DeleteTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).DeleteTemplate(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_ListTemplates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TemplateServer).ListTemplates(m, &templateListTemplatesServer{stream})
}

type Template_ListTemplatesServer interface {
	Send(*WorkflowTemplate) error
	grpc.ServerStream
}

type templateListTemplatesServer struct {
	grpc.ServerStream
}

func (x *templateListTemplatesServer) Send(m *WorkflowTemplate) error {
	return x.ServerStream.SendMsg(m)
}

func _Template_UpdateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).UpdateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.template.Template/UpdateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).UpdateTemplate(ctx, req.(*WorkflowTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

var _Template_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.tinkerbell.tink.protos.template.Template",
	HandlerType: (*TemplateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTemplate",
			Handler:    _Template_CreateTemplate_Handler,
		},
		{
			MethodName: "GetTemplate",
			Handler:    _Template_GetTemplate_Handler,
		},
		{
			MethodName: "DeleteTemplate",
			Handler:    _Template_DeleteTemplate_Handler,
		},
		{
			MethodName: "UpdateTemplate",
			Handler:    _Template_UpdateTemplate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTemplates",
			Handler:       _Template_ListTemplates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "template/template.proto",
}
