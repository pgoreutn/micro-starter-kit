// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/echo/proto/echo/echo.proto

package go_micro_srv_echo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Echo service

type EchoService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Echo_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Echo_PingPongService, error)
}

type echoService struct {
	c    client.Client
	name string
}

func NewEchoService(name string, c client.Client) EchoService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.echo"
	}
	return &echoService{
		c:    c,
		name: name,
	}
}

func (c *echoService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Echo.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Echo_StreamService, error) {
	req := c.c.NewRequest(c.name, "Echo.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &echoServiceStream{stream}, nil
}

type Echo_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type echoServiceStream struct {
	stream client.Stream
}

func (x *echoServiceStream) Close() error {
	return x.stream.Close()
}

func (x *echoServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *echoServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *echoServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoService) PingPong(ctx context.Context, opts ...client.CallOption) (Echo_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Echo.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &echoServicePingPong{stream}, nil
}

type Echo_PingPongService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type echoServicePingPong struct {
	stream client.Stream
}

func (x *echoServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *echoServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *echoServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *echoServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *echoServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Echo service

type EchoHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Echo_StreamStream) error
	PingPong(context.Context, Echo_PingPongStream) error
}

func RegisterEchoHandler(s server.Server, hdlr EchoHandler, opts ...server.HandlerOption) error {
	type echo interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Echo struct {
		echo
	}
	h := &echoHandler{hdlr}
	return s.Handle(s.NewHandler(&Echo{h}, opts...))
}

type echoHandler struct {
	EchoHandler
}

func (h *echoHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.EchoHandler.Call(ctx, in, out)
}

func (h *echoHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.EchoHandler.Stream(ctx, m, &echoStreamStream{stream})
}

type Echo_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type echoStreamStream struct {
	stream server.Stream
}

func (x *echoStreamStream) Close() error {
	return x.stream.Close()
}

func (x *echoStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *echoStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *echoStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *echoHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.EchoHandler.PingPong(ctx, &echoPingPongStream{stream})
}

type Echo_PingPongStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type echoPingPongStream struct {
	stream server.Stream
}

func (x *echoPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *echoPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *echoPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *echoPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *echoPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
