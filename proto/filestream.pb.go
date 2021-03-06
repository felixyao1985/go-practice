// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/filestream.proto

/*
	Package protos is a generated protocol buffer package.

	It is generated from these files:
		protos/filestream.proto

	It has these top-level messages:
		FileChunk
		FileUploadAck
*/
package test

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FileUploadStatusCode int32

const (
	FileUploadStatusCode_Unknown FileUploadStatusCode = 0
	FileUploadStatusCode_Ok      FileUploadStatusCode = 1
	FileUploadStatusCode_Failed  FileUploadStatusCode = 2
)

var FileUploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}
var FileUploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x FileUploadStatusCode) String() string {
	return proto.EnumName(FileUploadStatusCode_name, int32(x))
}
func (FileUploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorFilestream, []int{0}
}

type FileChunk struct {
	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (m *FileChunk) Reset()                    { *m = FileChunk{} }
func (m *FileChunk) String() string            { return proto.CompactTextString(m) }
func (*FileChunk) ProtoMessage()               {}
func (*FileChunk) Descriptor() ([]byte, []int) { return fileDescriptorFilestream, []int{0} }

func (m *FileChunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type FileUploadAck struct {
	Message string               `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code    FileUploadStatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=protos.FileUploadStatusCode" json:"Code,omitempty"`
}

func (m *FileUploadAck) Reset()                    { *m = FileUploadAck{} }
func (m *FileUploadAck) String() string            { return proto.CompactTextString(m) }
func (*FileUploadAck) ProtoMessage()               {}
func (*FileUploadAck) Descriptor() ([]byte, []int) { return fileDescriptorFilestream, []int{1} }

func (m *FileUploadAck) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *FileUploadAck) GetCode() FileUploadStatusCode {
	if m != nil {
		return m.Code
	}
	return FileUploadStatusCode_Unknown
}

func init() {
	proto.RegisterType((*FileChunk)(nil), "protos.FileChunk")
	proto.RegisterType((*FileUploadAck)(nil), "protos.FileUploadAck")
	proto.RegisterEnum("protos.FileUploadStatusCode", FileUploadStatusCode_name, FileUploadStatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GRPCStreamUploadService service

type GRPCStreamUploadServiceClient interface {
	SendFile(ctx context.Context, opts ...grpc.CallOption) (GRPCStreamUploadService_SendFileClient, error)
}

type gRPCStreamUploadServiceClient struct {
	cc *grpc.ClientConn
}

func NewGRPCStreamUploadServiceClient(cc *grpc.ClientConn) GRPCStreamUploadServiceClient {
	return &gRPCStreamUploadServiceClient{cc}
}

func (c *gRPCStreamUploadServiceClient) SendFile(ctx context.Context, opts ...grpc.CallOption) (GRPCStreamUploadService_SendFileClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GRPCStreamUploadService_serviceDesc.Streams[0], c.cc, "/protos.gRPCStreamUploadService/SendFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCStreamUploadServiceSendFileClient{stream}
	return x, nil
}

type GRPCStreamUploadService_SendFileClient interface {
	Send(*FileChunk) error
	CloseAndRecv() (*FileUploadAck, error)
	grpc.ClientStream
}

type gRPCStreamUploadServiceSendFileClient struct {
	grpc.ClientStream
}

func (x *gRPCStreamUploadServiceSendFileClient) Send(m *FileChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gRPCStreamUploadServiceSendFileClient) CloseAndRecv() (*FileUploadAck, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileUploadAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GRPCStreamUploadService service

type GRPCStreamUploadServiceServer interface {
	SendFile(GRPCStreamUploadService_SendFileServer) error
}

func RegisterGRPCStreamUploadServiceServer(s *grpc.Server, srv GRPCStreamUploadServiceServer) {
	s.RegisterService(&_GRPCStreamUploadService_serviceDesc, srv)
}

func _GRPCStreamUploadService_SendFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GRPCStreamUploadServiceServer).SendFile(&gRPCStreamUploadServiceSendFileServer{stream})
}

type GRPCStreamUploadService_SendFileServer interface {
	SendAndClose(*FileUploadAck) error
	Recv() (*FileChunk, error)
	grpc.ServerStream
}

type gRPCStreamUploadServiceSendFileServer struct {
	grpc.ServerStream
}

func (x *gRPCStreamUploadServiceSendFileServer) SendAndClose(m *FileUploadAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gRPCStreamUploadServiceSendFileServer) Recv() (*FileChunk, error) {
	m := new(FileChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GRPCStreamUploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.gRPCStreamUploadService",
	HandlerType: (*GRPCStreamUploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendFile",
			Handler:       _GRPCStreamUploadService_SendFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protos/filestream.proto",
}

func (m *FileChunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileChunk) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFilestream(dAtA, i, uint64(len(m.Content)))
		i += copy(dAtA[i:], m.Content)
	}
	return i, nil
}

func (m *FileUploadAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileUploadAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFilestream(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.Code != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintFilestream(dAtA, i, uint64(m.Code))
	}
	return i, nil
}

func encodeVarintFilestream(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FileChunk) Size() (n int) {
	var l int
	_ = l
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovFilestream(uint64(l))
	}
	return n
}

func (m *FileUploadAck) Size() (n int) {
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovFilestream(uint64(l))
	}
	if m.Code != 0 {
		n += 1 + sovFilestream(uint64(m.Code))
	}
	return n
}

func sovFilestream(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFilestream(x uint64) (n int) {
	return sovFilestream(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FileChunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFilestream
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FileChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilestream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFilestream
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content[:0], dAtA[iNdEx:postIndex]...)
			if m.Content == nil {
				m.Content = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFilestream(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFilestream
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FileUploadAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFilestream
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FileUploadAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileUploadAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilestream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFilestream
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilestream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (FileUploadStatusCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFilestream(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFilestream
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFilestream(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFilestream
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFilestream
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFilestream
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFilestream
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFilestream
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFilestream(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFilestream = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFilestream   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protos/filestream.proto", fileDescriptorFilestream) }

var fileDescriptorFilestream = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0xcb, 0xcc, 0x49, 0x2d, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x03, 0x8b,
	0x08, 0xb1, 0x41, 0x24, 0x94, 0x54, 0xb9, 0x38, 0xdd, 0x32, 0x73, 0x52, 0x9d, 0x33, 0x4a, 0xf3,
	0xb2, 0x85, 0x24, 0xb8, 0xd8, 0x9d, 0xf3, 0xf3, 0x4a, 0x52, 0xf3, 0x4a, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x78, 0x82, 0x60, 0x5c, 0xa5, 0x68, 0x2e, 0x5e, 0x90, 0xb2, 0xd0, 0x82, 0x9c, 0xfc, 0xc4,
	0x14, 0xc7, 0x64, 0xb0, 0x52, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0xb0, 0x52, 0xce, 0x20,
	0x18, 0x57, 0xc8, 0x80, 0x8b, 0xc5, 0x39, 0x3f, 0x25, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0xcf,
	0x48, 0x06, 0x62, 0x5f, 0xb1, 0x1e, 0x42, 0x7b, 0x70, 0x49, 0x62, 0x49, 0x69, 0x31, 0x48, 0x4d,
	0x10, 0x58, 0xa5, 0x96, 0x39, 0x97, 0x08, 0x36, 0x59, 0x21, 0x6e, 0x2e, 0xf6, 0xd0, 0xbc, 0xec,
	0xbc, 0xfc, 0xf2, 0x3c, 0x01, 0x06, 0x21, 0x36, 0x2e, 0x26, 0xff, 0x6c, 0x01, 0x46, 0x21, 0x2e,
	0x2e, 0x36, 0xb7, 0xc4, 0xcc, 0x9c, 0xd4, 0x14, 0x01, 0x26, 0xa3, 0x60, 0x2e, 0xf1, 0xf4, 0xa0,
	0x00, 0xe7, 0x60, 0xb0, 0xc7, 0xa0, 0xda, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x2c, 0xb8,
	0x38, 0x82, 0x53, 0xf3, 0x52, 0x40, 0xe6, 0x0a, 0x09, 0x22, 0xbb, 0x01, 0xec, 0x53, 0x29, 0x51,
	0x4c, 0x67, 0x39, 0x26, 0x67, 0x2b, 0x31, 0x68, 0x30, 0x3a, 0x09, 0x9c, 0x78, 0x24, 0xc7, 0x78,
	0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x24, 0x41, 0xc2, 0xca,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x97, 0xfd, 0xb6, 0x4d, 0x01, 0x00, 0x00,
}
