// Code generated by protoc-gen-go. DO NOT EDIT.
// source: weather.proto

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
	return fileDescriptor_231dcd72b885f4be, []int{0}
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
	return fileDescriptor_231dcd72b885f4be, []int{1}
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

func init() {
	proto.RegisterType((*LocationAndDate)(nil), "services.LocationAndDate")
	proto.RegisterType((*Weather)(nil), "services.Weather")
}

func init() { proto.RegisterFile("weather.proto", fileDescriptor_231dcd72b885f4be) }

var fileDescriptor_231dcd72b885f4be = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4f, 0x4d, 0x2c,
	0xc9, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0x2d, 0x56, 0x72, 0xe4, 0xe2, 0xf7, 0xc9, 0x4f, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x73,
	0xcc, 0x4b, 0x71, 0x49, 0x2c, 0x49, 0x15, 0x92, 0xe2, 0xe2, 0xc8, 0x81, 0x0a, 0x49, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x42, 0x42, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0xa9, 0x12, 0x4c,
	0x60, 0x71, 0x30, 0x5b, 0x29, 0x9a, 0x8b, 0x3d, 0x1c, 0x62, 0x3a, 0xa9, 0x5a, 0x85, 0x14, 0xb8,
	0xb8, 0x4b, 0x52, 0x73, 0x0b, 0x52, 0x8b, 0x12, 0x4b, 0x4a, 0x8b, 0x52, 0x25, 0x98, 0x15, 0x18,
	0x35, 0x58, 0x83, 0x90, 0x85, 0x8c, 0xfc, 0xb8, 0xf8, 0xa0, 0x86, 0x07, 0x43, 0x9c, 0x2c, 0x64,
	0xc3, 0xc5, 0xe5, 0x9e, 0x5a, 0x02, 0xb3, 0x51, 0x52, 0x0f, 0xe6, 0x15, 0x3d, 0x34, 0x7f, 0x48,
	0x09, 0x22, 0xa4, 0xa0, 0xaa, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x01, 0x60, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x95, 0x07, 0xf7, 0xd5, 0x11, 0x01, 0x00, 0x00,
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
	Metadata: "weather.proto",
}