// Code generated by protoc-gen-go.
// source: bazil.org/bazil/peer/wire/peer.proto
// DO NOT EDIT!

/*
Package wire is a generated protocol buffer package.

It is generated from these files:
	bazil.org/bazil/peer/wire/peer.proto

It has these top-level messages:
	PingRequest
	PingResponse
	ObjectPutRequest
	ObjectPutResponse
	ObjectGetRequest
	ObjectGetResponse
*/
package wire

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type PingRequest struct {
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}

type PingResponse struct {
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}

type ObjectPutRequest struct {
	// Only set in the first streamed message.
	Key  []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ObjectPutRequest) Reset()         { *m = ObjectPutRequest{} }
func (m *ObjectPutRequest) String() string { return proto.CompactTextString(m) }
func (*ObjectPutRequest) ProtoMessage()    {}

type ObjectPutResponse struct {
}

func (m *ObjectPutResponse) Reset()         { *m = ObjectPutResponse{} }
func (m *ObjectPutResponse) String() string { return proto.CompactTextString(m) }
func (*ObjectPutResponse) ProtoMessage()    {}

type ObjectGetRequest struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *ObjectGetRequest) Reset()         { *m = ObjectGetRequest{} }
func (m *ObjectGetRequest) String() string { return proto.CompactTextString(m) }
func (*ObjectGetRequest) ProtoMessage()    {}

type ObjectGetResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ObjectGetResponse) Reset()         { *m = ObjectGetResponse{} }
func (m *ObjectGetResponse) String() string { return proto.CompactTextString(m) }
func (*ObjectGetResponse) ProtoMessage()    {}

func init() {
}

// Client API for Peer service

type PeerClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	ObjectPut(ctx context.Context, opts ...grpc.CallOption) (Peer_ObjectPutClient, error)
	ObjectGet(ctx context.Context, in *ObjectGetRequest, opts ...grpc.CallOption) (Peer_ObjectGetClient, error)
}

type peerClient struct {
	cc *grpc.ClientConn
}

func NewPeerClient(cc *grpc.ClientConn) PeerClient {
	return &peerClient{cc}
}

func (c *peerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/bazil.peer.Peer/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerClient) ObjectPut(ctx context.Context, opts ...grpc.CallOption) (Peer_ObjectPutClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Peer_serviceDesc.Streams[0], c.cc, "/bazil.peer.Peer/ObjectPut", opts...)
	if err != nil {
		return nil, err
	}
	x := &peerObjectPutClient{stream}
	return x, nil
}

type Peer_ObjectPutClient interface {
	Send(*ObjectPutRequest) error
	CloseAndRecv() (*ObjectPutResponse, error)
	grpc.ClientStream
}

type peerObjectPutClient struct {
	grpc.ClientStream
}

func (x *peerObjectPutClient) Send(m *ObjectPutRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *peerObjectPutClient) CloseAndRecv() (*ObjectPutResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ObjectPutResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *peerClient) ObjectGet(ctx context.Context, in *ObjectGetRequest, opts ...grpc.CallOption) (Peer_ObjectGetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Peer_serviceDesc.Streams[1], c.cc, "/bazil.peer.Peer/ObjectGet", opts...)
	if err != nil {
		return nil, err
	}
	x := &peerObjectGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Peer_ObjectGetClient interface {
	Recv() (*ObjectGetResponse, error)
	grpc.ClientStream
}

type peerObjectGetClient struct {
	grpc.ClientStream
}

func (x *peerObjectGetClient) Recv() (*ObjectGetResponse, error) {
	m := new(ObjectGetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Peer service

type PeerServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	ObjectPut(Peer_ObjectPutServer) error
	ObjectGet(*ObjectGetRequest, Peer_ObjectGetServer) error
}

func RegisterPeerServer(s *grpc.Server, srv PeerServer) {
	s.RegisterService(&_Peer_serviceDesc, srv)
}

func _Peer_Ping_Handler(srv interface{}, ctx context.Context, buf []byte) (interface{}, error) {
	in := new(PingRequest)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(PeerServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Peer_ObjectPut_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PeerServer).ObjectPut(&peerObjectPutServer{stream})
}

type Peer_ObjectPutServer interface {
	SendAndClose(*ObjectPutResponse) error
	Recv() (*ObjectPutRequest, error)
	grpc.ServerStream
}

type peerObjectPutServer struct {
	grpc.ServerStream
}

func (x *peerObjectPutServer) SendAndClose(m *ObjectPutResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *peerObjectPutServer) Recv() (*ObjectPutRequest, error) {
	m := new(ObjectPutRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Peer_ObjectGet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ObjectGetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PeerServer).ObjectGet(m, &peerObjectGetServer{stream})
}

type Peer_ObjectGetServer interface {
	Send(*ObjectGetResponse) error
	grpc.ServerStream
}

type peerObjectGetServer struct {
	grpc.ServerStream
}

func (x *peerObjectGetServer) Send(m *ObjectGetResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Peer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bazil.peer.Peer",
	HandlerType: (*PeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Peer_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ObjectPut",
			Handler:       _Peer_ObjectPut_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ObjectGet",
			Handler:       _Peer_ObjectGet_Handler,
			ServerStreams: true,
		},
	},
}
