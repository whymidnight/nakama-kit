// Copyright 2024 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//*
// The realtime protocol for Nakama server.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: peer.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Peer_Call_FullMethodName   = "/nakama.peer.Peer/Call"
	Peer_Stream_FullMethodName = "/nakama.peer.Peer/Stream"
)

// PeerClient is the client API for Peer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeerClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseWriter, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (Peer_StreamClient, error)
}

type peerClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerClient(cc grpc.ClientConnInterface) PeerClient {
	return &peerClient{cc}
}

func (c *peerClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseWriter, error) {
	out := new(ResponseWriter)
	err := c.cc.Invoke(ctx, Peer_Call_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Peer_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Peer_ServiceDesc.Streams[0], Peer_Stream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &peerStreamClient{stream}
	return x, nil
}

type Peer_StreamClient interface {
	Send(*Request) error
	Recv() (*ResponseWriter, error)
	grpc.ClientStream
}

type peerStreamClient struct {
	grpc.ClientStream
}

func (x *peerStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *peerStreamClient) Recv() (*ResponseWriter, error) {
	m := new(ResponseWriter)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PeerServer is the server API for Peer service.
// All implementations must embed UnimplementedPeerServer
// for forward compatibility
type PeerServer interface {
	Call(context.Context, *Request) (*ResponseWriter, error)
	Stream(Peer_StreamServer) error
	mustEmbedUnimplementedPeerServer()
}

// UnimplementedPeerServer must be embedded to have forward compatible implementations.
type UnimplementedPeerServer struct {
}

func (UnimplementedPeerServer) Call(context.Context, *Request) (*ResponseWriter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedPeerServer) Stream(Peer_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedPeerServer) mustEmbedUnimplementedPeerServer() {}

// UnsafePeerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeerServer will
// result in compilation errors.
type UnsafePeerServer interface {
	mustEmbedUnimplementedPeerServer()
}

func RegisterPeerServer(s grpc.ServiceRegistrar, srv PeerServer) {
	s.RegisterService(&Peer_ServiceDesc, srv)
}

func _Peer_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Peer_Call_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Peer_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PeerServer).Stream(&peerStreamServer{stream})
}

type Peer_StreamServer interface {
	Send(*ResponseWriter) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type peerStreamServer struct {
	grpc.ServerStream
}

func (x *peerStreamServer) Send(m *ResponseWriter) error {
	return x.ServerStream.SendMsg(m)
}

func (x *peerStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Peer_ServiceDesc is the grpc.ServiceDesc for Peer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Peer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nakama.peer.Peer",
	HandlerType: (*PeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Peer_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Peer_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "peer.proto",
}
