// Code generated by protoc-gen-go.
// source: session.proto
// DO NOT EDIT!

/*
Package mnemosynerpc is a generated protocol buffer package.

It is generated from these files:
	session.proto

It has these top-level messages:
	Session
	GetRequest
	GetResponse
	ContextResponse
	ListRequest
	ListResponse
	ExistsRequest
	ExistsResponse
	StartRequest
	StartResponse
	AbandonRequest
	AbandonResponse
	SetValueRequest
	SetValueResponse
	DeleteRequest
	DeleteResponse
*/
package mnemosynerpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Session struct {
	AccessToken   string                     `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	SubjectId     string                     `protobuf:"bytes,2,opt,name=subject_id,json=subjectId" json:"subject_id,omitempty"`
	SubjectClient string                     `protobuf:"bytes,3,opt,name=subject_client,json=subjectClient" json:"subject_client,omitempty"`
	Bag           map[string]string          `protobuf:"bytes,4,rep,name=bag" json:"bag,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ExpireAt      *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=expire_at,json=expireAt" json:"expire_at,omitempty"`
	RefreshToken  string                     `protobuf:"bytes,6,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Session) GetBag() map[string]string {
	if m != nil {
		return m.Bag
	}
	return nil
}

func (m *Session) GetExpireAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireAt
	}
	return nil
}

type GetRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type ContextResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *ContextResponse) Reset()                    { *m = ContextResponse{} }
func (m *ContextResponse) String() string            { return proto.CompactTextString(m) }
func (*ContextResponse) ProtoMessage()               {}
func (*ContextResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ContextResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type ListRequest struct {
	Offset       int64                      `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit        int64                      `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	ExpireAtFrom *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=expire_at_from,json=expireAtFrom" json:"expire_at_from,omitempty"`
	ExpireAtTo   *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=expire_at_to,json=expireAtTo" json:"expire_at_to,omitempty"`
	RefreshToken string                     `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListRequest) GetExpireAtFrom() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireAtFrom
	}
	return nil
}

func (m *ListRequest) GetExpireAtTo() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireAtTo
	}
	return nil
}

type ListResponse struct {
	Sessions []*Session `protobuf:"bytes,1,rep,name=sessions" json:"sessions,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListResponse) GetSessions() []*Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

type ExistsRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
}

func (m *ExistsRequest) Reset()                    { *m = ExistsRequest{} }
func (m *ExistsRequest) String() string            { return proto.CompactTextString(m) }
func (*ExistsRequest) ProtoMessage()               {}
func (*ExistsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ExistsResponse struct {
	Exists bool `protobuf:"varint,1,opt,name=exists" json:"exists,omitempty"`
}

func (m *ExistsResponse) Reset()                    { *m = ExistsResponse{} }
func (m *ExistsResponse) String() string            { return proto.CompactTextString(m) }
func (*ExistsResponse) ProtoMessage()               {}
func (*ExistsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type StartRequest struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *StartRequest) Reset()                    { *m = StartRequest{} }
func (m *StartRequest) String() string            { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()               {}
func (*StartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StartRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type StartResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *StartResponse) Reset()                    { *m = StartResponse{} }
func (m *StartResponse) String() string            { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()               {}
func (*StartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StartResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type AbandonRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
}

func (m *AbandonRequest) Reset()                    { *m = AbandonRequest{} }
func (m *AbandonRequest) String() string            { return proto.CompactTextString(m) }
func (*AbandonRequest) ProtoMessage()               {}
func (*AbandonRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type AbandonResponse struct {
	Abandoned bool `protobuf:"varint,1,opt,name=abandoned" json:"abandoned,omitempty"`
}

func (m *AbandonResponse) Reset()                    { *m = AbandonResponse{} }
func (m *AbandonResponse) String() string            { return proto.CompactTextString(m) }
func (*AbandonResponse) ProtoMessage()               {}
func (*AbandonResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type SetValueRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	Key         string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value       string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *SetValueRequest) Reset()                    { *m = SetValueRequest{} }
func (m *SetValueRequest) String() string            { return proto.CompactTextString(m) }
func (*SetValueRequest) ProtoMessage()               {}
func (*SetValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type SetValueResponse struct {
	Bag map[string]string `protobuf:"bytes,1,rep,name=bag" json:"bag,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SetValueResponse) Reset()                    { *m = SetValueResponse{} }
func (m *SetValueResponse) String() string            { return proto.CompactTextString(m) }
func (*SetValueResponse) ProtoMessage()               {}
func (*SetValueResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SetValueResponse) GetBag() map[string]string {
	if m != nil {
		return m.Bag
	}
	return nil
}

type DeleteRequest struct {
	// Types that are valid to be assigned to DeleteBy:
	//	*DeleteRequest_AccessToken
	//	*DeleteRequest_RefreshToken
	DeleteBy     isDeleteRequest_DeleteBy   `protobuf_oneof:"delete_by"`
	ExpireAtFrom *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=expire_at_from,json=expireAtFrom" json:"expire_at_from,omitempty"`
	ExpireAtTo   *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=expire_at_to,json=expireAtTo" json:"expire_at_to,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type isDeleteRequest_DeleteBy interface {
	isDeleteRequest_DeleteBy()
}

type DeleteRequest_AccessToken struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,oneof"`
}
type DeleteRequest_RefreshToken struct {
	RefreshToken string `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,oneof"`
}

func (*DeleteRequest_AccessToken) isDeleteRequest_DeleteBy()  {}
func (*DeleteRequest_RefreshToken) isDeleteRequest_DeleteBy() {}

func (m *DeleteRequest) GetDeleteBy() isDeleteRequest_DeleteBy {
	if m != nil {
		return m.DeleteBy
	}
	return nil
}

func (m *DeleteRequest) GetAccessToken() string {
	if x, ok := m.GetDeleteBy().(*DeleteRequest_AccessToken); ok {
		return x.AccessToken
	}
	return ""
}

func (m *DeleteRequest) GetRefreshToken() string {
	if x, ok := m.GetDeleteBy().(*DeleteRequest_RefreshToken); ok {
		return x.RefreshToken
	}
	return ""
}

func (m *DeleteRequest) GetExpireAtFrom() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireAtFrom
	}
	return nil
}

func (m *DeleteRequest) GetExpireAtTo() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpireAtTo
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DeleteRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DeleteRequest_OneofMarshaler, _DeleteRequest_OneofUnmarshaler, _DeleteRequest_OneofSizer, []interface{}{
		(*DeleteRequest_AccessToken)(nil),
		(*DeleteRequest_RefreshToken)(nil),
	}
}

func _DeleteRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DeleteRequest)
	// delete_by
	switch x := m.DeleteBy.(type) {
	case *DeleteRequest_AccessToken:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AccessToken)
	case *DeleteRequest_RefreshToken:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.RefreshToken)
	case nil:
	default:
		return fmt.Errorf("DeleteRequest.DeleteBy has unexpected type %T", x)
	}
	return nil
}

func _DeleteRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DeleteRequest)
	switch tag {
	case 1: // delete_by.access_token
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.DeleteBy = &DeleteRequest_AccessToken{x}
		return true, err
	case 4: // delete_by.refresh_token
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.DeleteBy = &DeleteRequest_RefreshToken{x}
		return true, err
	default:
		return false, nil
	}
}

func _DeleteRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DeleteRequest)
	// delete_by
	switch x := m.DeleteBy.(type) {
	case *DeleteRequest_AccessToken:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AccessToken)))
		n += len(x.AccessToken)
	case *DeleteRequest_RefreshToken:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RefreshToken)))
		n += len(x.RefreshToken)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DeleteResponse struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func init() {
	proto.RegisterType((*Session)(nil), "mnemosynerpc.Session")
	proto.RegisterType((*GetRequest)(nil), "mnemosynerpc.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "mnemosynerpc.GetResponse")
	proto.RegisterType((*ContextResponse)(nil), "mnemosynerpc.ContextResponse")
	proto.RegisterType((*ListRequest)(nil), "mnemosynerpc.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "mnemosynerpc.ListResponse")
	proto.RegisterType((*ExistsRequest)(nil), "mnemosynerpc.ExistsRequest")
	proto.RegisterType((*ExistsResponse)(nil), "mnemosynerpc.ExistsResponse")
	proto.RegisterType((*StartRequest)(nil), "mnemosynerpc.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "mnemosynerpc.StartResponse")
	proto.RegisterType((*AbandonRequest)(nil), "mnemosynerpc.AbandonRequest")
	proto.RegisterType((*AbandonResponse)(nil), "mnemosynerpc.AbandonResponse")
	proto.RegisterType((*SetValueRequest)(nil), "mnemosynerpc.SetValueRequest")
	proto.RegisterType((*SetValueResponse)(nil), "mnemosynerpc.SetValueResponse")
	proto.RegisterType((*DeleteRequest)(nil), "mnemosynerpc.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "mnemosynerpc.DeleteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SessionManager service

type SessionManagerClient interface {
	Context(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ContextResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error)
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Abandon(ctx context.Context, in *AbandonRequest, opts ...grpc.CallOption) (*AbandonResponse, error)
	SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type sessionManagerClient struct {
	cc *grpc.ClientConn
}

func NewSessionManagerClient(cc *grpc.ClientConn) SessionManagerClient {
	return &sessionManagerClient{cc}
}

func (c *sessionManagerClient) Context(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ContextResponse, error) {
	out := new(ContextResponse)
	err := grpc.Invoke(ctx, "/mnemosynerpc.SessionManager/Context", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/mnemosynerpc.SessionManager/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/mnemosynerpc.SessionManager/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error) {
	out := new(ExistsResponse)
	err := grpc.Invoke(ctx, "/mnemosynerpc.SessionManager/Exists", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := grpc.Invoke(ctx, "/mnemosynerpc.SessionManager/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) Abandon(ctx context.Context, in *AbandonRequest, opts ...grpc.CallOption) (*AbandonResponse, error) {
	out := new(AbandonResponse)
	err := grpc.Invoke(ctx, "/mnemosynerpc.SessionManager/Abandon", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueResponse, error) {
	out := new(SetValueResponse)
	err := grpc.Invoke(ctx, "/mnemosynerpc.SessionManager/SetValue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/mnemosynerpc.SessionManager/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SessionManager service

type SessionManagerServer interface {
	Context(context.Context, *google_protobuf1.Empty) (*ContextResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Exists(context.Context, *ExistsRequest) (*ExistsResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Abandon(context.Context, *AbandonRequest) (*AbandonResponse, error)
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterSessionManagerServer(s *grpc.Server, srv SessionManagerServer) {
	s.RegisterService(&_SessionManager_serviceDesc, srv)
}

func _SessionManager_Context_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).Context(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mnemosynerpc.SessionManager/Context",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).Context(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mnemosynerpc.SessionManager/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mnemosynerpc.SessionManager/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).Exists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mnemosynerpc.SessionManager/Exists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).Exists(ctx, req.(*ExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mnemosynerpc.SessionManager/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_Abandon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbandonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).Abandon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mnemosynerpc.SessionManager/Abandon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).Abandon(ctx, req.(*AbandonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mnemosynerpc.SessionManager/SetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).SetValue(ctx, req.(*SetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mnemosynerpc.SessionManager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mnemosynerpc.SessionManager",
	HandlerType: (*SessionManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Context",
			Handler:    _SessionManager_Context_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SessionManager_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SessionManager_List_Handler,
		},
		{
			MethodName: "Exists",
			Handler:    _SessionManager_Exists_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _SessionManager_Start_Handler,
		},
		{
			MethodName: "Abandon",
			Handler:    _SessionManager_Abandon_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _SessionManager_SetValue_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SessionManager_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}

func init() { proto.RegisterFile("session.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 793 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x15, 0x45, 0x5d, 0x47, 0x17, 0x1b, 0x0b, 0xd7, 0x60, 0xe9, 0x0b, 0x5c, 0x1a, 0x6e, 0xf5,
	0x44, 0xb5, 0x32, 0xd0, 0xab, 0x51, 0xdb, 0x72, 0x55, 0xbb, 0x68, 0xf2, 0x42, 0x1b, 0x79, 0x0a,
	0x20, 0x50, 0xd2, 0x8a, 0x66, 0x2c, 0x72, 0x19, 0xee, 0x2a, 0xb0, 0x3e, 0x20, 0x3f, 0x15, 0xe4,
	0x63, 0xf2, 0x98, 0xcf, 0x08, 0xb4, 0xbb, 0xa4, 0x44, 0x9a, 0xbe, 0xc8, 0x7e, 0x11, 0xb8, 0xb3,
	0x67, 0x66, 0xf6, 0xcc, 0xcc, 0x19, 0x41, 0x83, 0x62, 0x4a, 0x5d, 0xe2, 0x9b, 0x41, 0x48, 0x18,
	0x41, 0x75, 0xcf, 0xc7, 0x1e, 0xa1, 0x33, 0x1f, 0x87, 0xc1, 0x50, 0xff, 0xcb, 0x71, 0xd9, 0xf5,
	0x74, 0x60, 0x0e, 0x89, 0xd7, 0x76, 0xc8, 0xc4, 0xf6, 0x9d, 0x36, 0x87, 0x0d, 0xa6, 0xe3, 0x76,
	0xc0, 0x66, 0x01, 0xa6, 0x6d, 0xe6, 0x7a, 0x98, 0x32, 0xdb, 0x0b, 0x16, 0x5f, 0x22, 0x94, 0x7e,
	0xf8, 0xb8, 0x33, 0xf6, 0x02, 0x36, 0x13, 0xbf, 0xc2, 0xc9, 0xf8, 0x94, 0x87, 0xf2, 0xa5, 0x78,
	0x11, 0xfa, 0x01, 0xea, 0xf6, 0x70, 0x88, 0x29, 0xed, 0x33, 0x72, 0x83, 0x7d, 0x4d, 0xd9, 0x53,
	0x5a, 0x55, 0xab, 0x26, 0x6c, 0x57, 0x73, 0x13, 0xda, 0x01, 0xa0, 0xd3, 0xc1, 0x3b, 0x3c, 0x64,
	0x7d, 0x77, 0xa4, 0xe5, 0x39, 0xa0, 0x2a, 0x2d, 0xff, 0x8d, 0xd0, 0x01, 0x34, 0xa3, 0xeb, 0xe1,
	0xc4, 0xc5, 0x3e, 0xd3, 0x54, 0x0e, 0x69, 0x48, 0xeb, 0x19, 0x37, 0xa2, 0x9f, 0x41, 0x1d, 0xd8,
	0x8e, 0x56, 0xd8, 0x53, 0x5b, 0xb5, 0xce, 0xae, 0xb9, 0x5c, 0x02, 0x53, 0x3e, 0xc6, 0xec, 0xda,
	0x4e, 0xcf, 0x67, 0xe1, 0xcc, 0x9a, 0x43, 0xd1, 0x6f, 0x50, 0xc5, 0xb7, 0x81, 0x1b, 0xe2, 0xbe,
	0xcd, 0xb4, 0xe2, 0x9e, 0xd2, 0xaa, 0x75, 0x74, 0xd3, 0x21, 0xc4, 0x99, 0x60, 0x33, 0x22, 0x69,
	0x5e, 0x45, 0x05, 0xb1, 0x2a, 0x02, 0x7c, 0xca, 0xd0, 0x3e, 0x34, 0x42, 0x3c, 0x0e, 0x31, 0xbd,
	0x96, 0xa4, 0x4a, 0xfc, 0x41, 0x75, 0x69, 0xe4, 0xac, 0xf4, 0x5f, 0xa1, 0x12, 0xa5, 0x43, 0xeb,
	0xa0, 0xde, 0xe0, 0x99, 0xe4, 0x3e, 0xff, 0x44, 0x1b, 0x50, 0xfc, 0x60, 0x4f, 0xa6, 0x58, 0xd2,
	0x15, 0x87, 0x3f, 0xf3, 0xbf, 0x2b, 0x46, 0x1b, 0xe0, 0x1c, 0x33, 0x0b, 0xbf, 0x9f, 0x62, 0xca,
	0x9e, 0x50, 0x3e, 0xe3, 0x6f, 0xa8, 0x71, 0x07, 0x1a, 0x10, 0x9f, 0x62, 0xd4, 0x86, 0xb2, 0x9c,
	0x06, 0x0e, 0xae, 0x75, 0xbe, 0xcb, 0xac, 0x85, 0x15, 0xa1, 0x8c, 0x2e, 0xac, 0x9d, 0x11, 0x9f,
	0xe1, 0xdb, 0x17, 0xc4, 0xf8, 0xa2, 0x40, 0xed, 0x95, 0x4b, 0xe3, 0x67, 0x6f, 0x42, 0x89, 0x8c,
	0xc7, 0x14, 0x33, 0xee, 0xaf, 0x5a, 0xf2, 0x34, 0xa7, 0x3d, 0x71, 0x3d, 0x97, 0x71, 0xda, 0xaa,
	0x25, 0x0e, 0xe8, 0x04, 0x9a, 0x71, 0x23, 0xfa, 0xe3, 0x90, 0x78, 0xbc, 0xc3, 0x0f, 0x77, 0xa3,
	0x1e, 0x75, 0xe3, 0xdf, 0x90, 0x78, 0xe8, 0x08, 0xea, 0x8b, 0x08, 0x8c, 0x68, 0x85, 0x47, 0xfd,
	0x21, 0xf2, 0xbf, 0x22, 0x77, 0xfb, 0x59, 0xbc, 0xdb, 0x4f, 0xe3, 0x14, 0xea, 0x82, 0xa1, 0xac,
	0xd1, 0x2f, 0x50, 0x91, 0xec, 0xa9, 0xa6, 0xf0, 0xa1, 0xbb, 0xa7, 0x48, 0x31, 0xcc, 0xe8, 0x40,
	0xa3, 0x77, 0xeb, 0x52, 0x46, 0x57, 0xe8, 0x6e, 0x0b, 0x9a, 0x91, 0x8f, 0x4c, 0xbc, 0x09, 0x25,
	0xcc, 0x2d, 0x1c, 0x5e, 0xb1, 0xe4, 0xc9, 0x38, 0x86, 0xfa, 0x25, 0xb3, 0xc3, 0xb8, 0x07, 0x2b,
	0x37, 0xf1, 0x04, 0x1a, 0x32, 0xc0, 0x73, 0xc7, 0xe0, 0x10, 0x9a, 0xa7, 0x03, 0xdb, 0x1f, 0x11,
	0x7f, 0x05, 0x86, 0x6d, 0x58, 0x8b, 0x9d, 0x64, 0xe2, 0x6d, 0xa8, 0xda, 0xc2, 0x84, 0x47, 0x92,
	0xe5, 0xc2, 0x60, 0xbc, 0x85, 0xb5, 0x4b, 0xcc, 0xde, 0xcc, 0x15, 0xf3, 0xf4, 0x34, 0x91, 0x06,
	0xf3, 0x19, 0x1a, 0x54, 0x97, 0x34, 0x68, 0x7c, 0x54, 0x60, 0x7d, 0x11, 0x5e, 0x3e, 0xe8, 0x0f,
	0xb1, 0x5c, 0x44, 0x9f, 0x7f, 0x4a, 0x57, 0x21, 0x09, 0x4e, 0x6e, 0x99, 0x67, 0xef, 0x81, 0xaf,
	0x0a, 0x34, 0xfe, 0xc1, 0x13, 0xcc, 0x62, 0x92, 0xfb, 0x59, 0x24, 0x2f, 0x72, 0x49, 0x9a, 0x07,
	0xe9, 0x59, 0x2e, 0x48, 0x54, 0x62, 0x9a, 0x33, 0x24, 0x97, 0x7f, 0xa1, 0xe4, 0xd4, 0x55, 0x24,
	0xd7, 0xad, 0x41, 0x75, 0xc4, 0xc9, 0xf5, 0x07, 0x33, 0xe3, 0x47, 0x68, 0x46, 0x4c, 0x65, 0xbd,
	0x37, 0xa0, 0x38, 0x24, 0x53, 0x3f, 0x5a, 0x1f, 0xe2, 0xd0, 0xf9, 0x5c, 0x80, 0xa6, 0x9c, 0xb9,
	0xd7, 0xb6, 0x6f, 0x3b, 0x38, 0x44, 0x5d, 0x28, 0xcb, 0xe5, 0x85, 0x36, 0xef, 0xa4, 0xee, 0xcd,
	0xff, 0x93, 0xf4, 0x9d, 0x64, 0xbb, 0x52, 0xbb, 0xce, 0xc8, 0xa1, 0x23, 0x50, 0xcf, 0x31, 0x43,
	0x5a, 0x12, 0xb7, 0x58, 0xc2, 0xfa, 0xf7, 0x19, 0x37, 0xb1, 0xf7, 0x31, 0x14, 0xe6, 0x7b, 0x01,
	0xa5, 0x40, 0x4b, 0xdb, 0x50, 0xd7, 0xb3, 0xae, 0xe2, 0x00, 0x3d, 0x28, 0x09, 0x85, 0xa3, 0xad,
	0x24, 0x2e, 0xb1, 0x2b, 0xf4, 0xed, 0xec, 0xcb, 0x38, 0x4c, 0x17, 0x8a, 0x5c, 0xbd, 0x28, 0x95,
	0x6d, 0x79, 0x27, 0xe8, 0x5b, 0x99, 0x77, 0x71, 0x8c, 0x0b, 0x28, 0x4b, 0x29, 0xa2, 0x54, 0xba,
	0xa4, 0xac, 0xd3, 0x35, 0x4d, 0xe9, 0xd7, 0xc8, 0xa1, 0xff, 0xa1, 0x12, 0xe9, 0x02, 0xed, 0xdc,
	0xa7, 0x17, 0x11, 0x6b, 0xf7, 0x61, 0x39, 0x89, 0x0a, 0x89, 0xf9, 0x48, 0x57, 0x28, 0xa1, 0x8f,
	0x74, 0x85, 0x92, 0x23, 0x65, 0xe4, 0x06, 0x25, 0x3e, 0x18, 0x87, 0xdf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xa1, 0xb5, 0x45, 0xa1, 0x2e, 0x09, 0x00, 0x00,
}
