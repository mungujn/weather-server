// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// The message containing the location and date
type LocationAndDate struct {
	Location             string   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Date                 string   `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocationAndDate) Reset()         { *m = LocationAndDate{} }
func (m *LocationAndDate) String() string { return proto.CompactTextString(m) }
func (*LocationAndDate) ProtoMessage()    {}
func (*LocationAndDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{0}
}

func (m *LocationAndDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationAndDate.Unmarshal(m, b)
}
func (m *LocationAndDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationAndDate.Marshal(b, m, deterministic)
}
func (m *LocationAndDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationAndDate.Merge(m, src)
}
func (m *LocationAndDate) XXX_Size() int {
	return xxx_messageInfo_LocationAndDate.Size(m)
}
func (m *LocationAndDate) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationAndDate.DiscardUnknown(m)
}

var xxx_messageInfo_LocationAndDate proto.InternalMessageInfo

func (m *LocationAndDate) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *LocationAndDate) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

// The response message containing the weather
type Weather struct {
	Location             string   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Date                 string   `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Temperature          int32    `protobuf:"varint,3,opt,name=temperature,proto3" json:"temperature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Weather) Reset()         { *m = Weather{} }
func (m *Weather) String() string { return proto.CompactTextString(m) }
func (*Weather) ProtoMessage()    {}
func (*Weather) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{1}
}

func (m *Weather) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Weather.Unmarshal(m, b)
}
func (m *Weather) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Weather.Marshal(b, m, deterministic)
}
func (m *Weather) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Weather.Merge(m, src)
}
func (m *Weather) XXX_Size() int {
	return xxx_messageInfo_Weather.Size(m)
}
func (m *Weather) XXX_DiscardUnknown() {
	xxx_messageInfo_Weather.DiscardUnknown(m)
}

var xxx_messageInfo_Weather proto.InternalMessageInfo

func (m *Weather) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Weather) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Weather) GetTemperature() int32 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

// Data
type Data struct {
	Location             string            `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Values               map[string]string `protobuf:"bytes,5,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{2}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Data) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

// Write result
type Result struct {
	Status               bool     `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{3}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Result) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*LocationAndDate)(nil), "services.LocationAndDate")
	proto.RegisterType((*Weather)(nil), "services.Weather")
	proto.RegisterType((*Data)(nil), "services.Data")
	proto.RegisterMapType((map[string]string)(nil), "services.Data.ValuesEntry")
	proto.RegisterType((*Result)(nil), "services.Result")
}

func init() { proto.RegisterFile("services.proto", fileDescriptor_8e16ccb8c5307b32) }

var fileDescriptor_8e16ccb8c5307b32 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x3d, 0x6f, 0xe2, 0x40,
	0x10, 0x65, 0xc1, 0x2c, 0xdc, 0x20, 0x01, 0xb7, 0x3a, 0x9d, 0xf6, 0x5c, 0x59, 0x5b, 0xb9, 0x40,
	0x14, 0xbe, 0x26, 0x41, 0x69, 0x50, 0x88, 0xd2, 0x44, 0x29, 0x1c, 0x25, 0x29, 0x52, 0x2d, 0x78,
	0x94, 0xa0, 0x18, 0x8c, 0x76, 0xc7, 0x48, 0xfc, 0x8c, 0xfc, 0xa8, 0xfc, 0xaf, 0xc8, 0x66, 0xcd,
	0x57, 0x1a, 0xd2, 0xed, 0x7b, 0xf3, 0xe6, 0xe9, 0xcd, 0xb3, 0xa1, 0x6b, 0xd1, 0xac, 0xe7, 0x33,
	0xb4, 0xc3, 0x95, 0xc9, 0x28, 0x13, 0xed, 0x0a, 0xab, 0x31, 0xf4, 0xee, 0xb2, 0x99, 0xa6, 0x79,
	0xb6, 0x1c, 0x2f, 0x93, 0x89, 0x26, 0x14, 0x3e, 0xb4, 0x53, 0x47, 0x49, 0x16, 0xb0, 0xf0, 0x57,
	0xbc, 0xc3, 0x42, 0x80, 0x97, 0x68, 0x42, 0x59, 0x2f, 0xf9, 0xf2, 0xad, 0x5e, 0xa0, 0xf5, 0x8c,
	0x9a, 0xde, 0xd0, 0xfc, 0x74, 0x55, 0x04, 0xd0, 0x21, 0x5c, 0xac, 0xd0, 0x68, 0xca, 0x0d, 0xca,
	0x46, 0xc0, 0xc2, 0x66, 0x7c, 0x48, 0xa9, 0x0f, 0x06, 0xde, 0x44, 0x93, 0x3e, 0xb2, 0xf6, 0x4e,
	0xac, 0x23, 0xe0, 0x6b, 0x9d, 0xe6, 0x68, 0x65, 0x33, 0x68, 0x84, 0x9d, 0xc8, 0x1f, 0xee, 0xee,
	0x2d, 0x76, 0x87, 0x4f, 0xe5, 0xf0, 0x66, 0x49, 0x66, 0x13, 0x3b, 0xa5, 0x7f, 0x09, 0x9d, 0x03,
	0x5a, 0xf4, 0xa1, 0xf1, 0x8e, 0x1b, 0x17, 0xba, 0x78, 0x8a, 0x3f, 0xd0, 0x2c, 0xa5, 0x2e, 0xf0,
	0x16, 0x8c, 0xea, 0x17, 0x4c, 0x8d, 0x80, 0xc7, 0x68, 0xf3, 0x94, 0xc4, 0x5f, 0xe0, 0x96, 0x34,
	0xe5, 0x56, 0xf2, 0x80, 0x85, 0xed, 0xd8, 0x21, 0x21, 0xa1, 0xb5, 0x40, 0x6b, 0xf5, 0x2b, 0xca,
	0x56, 0xb9, 0x5d, 0xc1, 0xe8, 0x1e, 0xba, 0xae, 0xac, 0x87, 0x6d, 0x44, 0x71, 0x05, 0x70, 0x8b,
	0x54, 0x35, 0xf8, 0x6f, 0x1f, 0xfd, 0xe4, 0xbb, 0xf8, 0xbf, 0xf7, 0x23, 0xa7, 0x56, 0xb5, 0xe8,
	0x93, 0x41, 0xaf, 0xb8, 0x71, 0xaa, 0x2d, 0x56, 0x8e, 0x03, 0xe0, 0xd7, 0x06, 0x8b, 0x7e, 0xbb,
	0xc7, 0x45, 0xf8, 0xfd, 0x3d, 0xde, 0x5e, 0xa0, 0x6a, 0x22, 0x04, 0x2f, 0x46, 0x9d, 0x7c, 0xd3,
	0x9e, 0x60, 0x55, 0x2b, 0x7c, 0x1f, 0x57, 0xc9, 0xb9, 0xbe, 0x03, 0xe0, 0x13, 0x4c, 0xf1, 0x3c,
	0xf5, 0x94, 0x97, 0x3f, 0xe6, 0xff, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x3c, 0x77, 0xe4,
	0xaa, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WeatherServiceClient is the client API for WeatherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WeatherServiceClient interface {
	// Retrieve the weather
	GetWeather(ctx context.Context, in *LocationAndDate, opts ...grpc.CallOption) (*Weather, error)
}

type weatherServiceClient struct {
	cc *grpc.ClientConn
}

func NewWeatherServiceClient(cc *grpc.ClientConn) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) GetWeather(ctx context.Context, in *LocationAndDate, opts ...grpc.CallOption) (*Weather, error) {
	out := new(Weather)
	err := c.cc.Invoke(ctx, "/services.WeatherService/GetWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServiceServer is the server API for WeatherService service.
type WeatherServiceServer interface {
	// Retrieve the weather
	GetWeather(context.Context, *LocationAndDate) (*Weather, error)
}

func RegisterWeatherServiceServer(s *grpc.Server, srv WeatherServiceServer) {
	s.RegisterService(&_WeatherService_serviceDesc, srv)
}

func _WeatherService_GetWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationAndDate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.WeatherService/GetWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetWeather(ctx, req.(*LocationAndDate))
	}
	return interceptor(ctx, in, info, handler)
}

var _WeatherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWeather",
			Handler:    _WeatherService_GetWeather_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

// DatabaseServiceClient is the client API for DatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseServiceClient interface {
	// create data
	Create(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Result, error)
	// read data
	Read(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error)
	// update data
	Update(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Result, error)
	// delete data
	Delete(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Result, error)
}

type databaseServiceClient struct {
	cc *grpc.ClientConn
}

func NewDatabaseServiceClient(cc *grpc.ClientConn) DatabaseServiceClient {
	return &databaseServiceClient{cc}
}

func (c *databaseServiceClient) Create(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/services.DatabaseService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Read(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/services.DatabaseService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Update(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/services.DatabaseService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Delete(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/services.DatabaseService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServiceServer is the server API for DatabaseService service.
type DatabaseServiceServer interface {
	// create data
	Create(context.Context, *Data) (*Result, error)
	// read data
	Read(context.Context, *Data) (*Data, error)
	// update data
	Update(context.Context, *Data) (*Result, error)
	// delete data
	Delete(context.Context, *Data) (*Result, error)
}

func RegisterDatabaseServiceServer(s *grpc.Server, srv DatabaseServiceServer) {
	s.RegisterService(&_DatabaseService_serviceDesc, srv)
}

func _DatabaseService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.DatabaseService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Create(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.DatabaseService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Read(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.DatabaseService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Update(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.DatabaseService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Delete(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatabaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.DatabaseService",
	HandlerType: (*DatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DatabaseService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _DatabaseService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DatabaseService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DatabaseService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
