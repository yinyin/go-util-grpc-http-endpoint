package grpchttpendpoint

import (
	"context"
	"errors"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ErrNotImplementedInWSGHE indicate the function is not implemented in WebSocket gRPC HTTP endpoint.
var ErrNotImplementedInWSGHE = errors.New("not implemented in WSGHE")

// WSGHEServerStream provides a server-side gRPC stream interface based on WebSocket.
// This implementation fulfills grpc.ServerStream interface.
// Most function is not implemented as it is not applicable.
type WSGHEServerStream struct {
	Ctx  context.Context
	Conn *websocket.Conn

	MarshalOpts   *protojson.MarshalOptions
	UnmarshalOpts *protojson.UnmarshalOptions
}

// SetHeader is not implemented.
func (ss *WSGHEServerStream) SetHeader(metadata.MD) error {
	return ErrNotImplementedInWSGHE
}

// SendHeader is not implemented.
func (ss *WSGHEServerStream) SendHeader(metadata.MD) error {
	return ErrNotImplementedInWSGHE
}

// SetTrailer is not implemented.
func (ss *WSGHEServerStream) SetTrailer(metadata.MD) {}

// Context returns the context for this stream.
func (ss *WSGHEServerStream) Context() context.Context {
	return ss.Ctx
}

// SendMsg wrap around SendProtoMessage().
// If given message is not a protobuf message, it will return ErrNotImplementedInWSGHE.
func (ss *WSGHEServerStream) SendMsg(m any) error {
	pbMessage, ok := m.(proto.Message)
	if !ok {
		return ErrNotImplementedInWSGHE
	}
	return ss.SendProtoMessage(pbMessage)
}

// RecvMsg wrap around RecvProtoMessage().
// If given message is not a protobuf message, it will return ErrNotImplementedInWSGHE.
func (ss *WSGHEServerStream) RecvMsg(m any) error {
	pbMessage, ok := m.(proto.Message)
	if !ok {
		return ErrNotImplementedInWSGHE
	}
	return ss.RecvProtoMessage(pbMessage)
}

// SendProtoMessage send the given protobuf message to the client.
func (ss *WSGHEServerStream) SendProtoMessage(m proto.Message) error {
	b, err := ss.MarshalOpts.Marshal(m)
	if nil != err {
		return err
	}
	return ss.Conn.WriteMessage(websocket.TextMessage, b)
}

// RecvProtoMessage receive the protobuf message from the client.
func (ss *WSGHEServerStream) RecvProtoMessage(m proto.Message) error {
	_, b, err := ss.Conn.ReadMessage()
	if nil != err {
		return err
	}
	return ss.UnmarshalOpts.Unmarshal(b, m)
}
