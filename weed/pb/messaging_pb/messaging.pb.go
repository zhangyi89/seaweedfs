// Code generated by protoc-gen-go.
// source: messaging.proto
// DO NOT EDIT!

/*
Package messaging_pb is a generated protocol buffer package.

It is generated from these files:
	messaging.proto

It has these top-level messages:
	SubscriberMessage
	Message
	BrokerMessage
	PublishRequest
	PublishResponse
	DeleteTopicRequest
	DeleteTopicResponse
	ConfigureTopicRequest
	ConfigureTopicResponse
	GetTopicConfigurationRequest
	GetTopicConfigurationResponse
	FindBrokerRequest
	FindBrokerResponse
	TopicConfiguration
*/
package messaging_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type SubscriberMessage_InitMessage_StartPosition int32

const (
	SubscriberMessage_InitMessage_LATEST    SubscriberMessage_InitMessage_StartPosition = 0
	SubscriberMessage_InitMessage_EARLIEST  SubscriberMessage_InitMessage_StartPosition = 1
	SubscriberMessage_InitMessage_TIMESTAMP SubscriberMessage_InitMessage_StartPosition = 2
)

var SubscriberMessage_InitMessage_StartPosition_name = map[int32]string{
	0: "LATEST",
	1: "EARLIEST",
	2: "TIMESTAMP",
}
var SubscriberMessage_InitMessage_StartPosition_value = map[string]int32{
	"LATEST":    0,
	"EARLIEST":  1,
	"TIMESTAMP": 2,
}

func (x SubscriberMessage_InitMessage_StartPosition) String() string {
	return proto.EnumName(SubscriberMessage_InitMessage_StartPosition_name, int32(x))
}
func (SubscriberMessage_InitMessage_StartPosition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

type TopicConfiguration_Partitioning int32

const (
	TopicConfiguration_NonNullKeyHash TopicConfiguration_Partitioning = 0
	TopicConfiguration_KeyHash        TopicConfiguration_Partitioning = 1
	TopicConfiguration_RoundRobin     TopicConfiguration_Partitioning = 2
)

var TopicConfiguration_Partitioning_name = map[int32]string{
	0: "NonNullKeyHash",
	1: "KeyHash",
	2: "RoundRobin",
}
var TopicConfiguration_Partitioning_value = map[string]int32{
	"NonNullKeyHash": 0,
	"KeyHash":        1,
	"RoundRobin":     2,
}

func (x TopicConfiguration_Partitioning) String() string {
	return proto.EnumName(TopicConfiguration_Partitioning_name, int32(x))
}
func (TopicConfiguration_Partitioning) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{13, 0}
}

type SubscriberMessage struct {
	Init    *SubscriberMessage_InitMessage `protobuf:"bytes,1,opt,name=init" json:"init,omitempty"`
	Ack     *SubscriberMessage_AckMessage  `protobuf:"bytes,2,opt,name=ack" json:"ack,omitempty"`
	IsClose bool                           `protobuf:"varint,3,opt,name=is_close,json=isClose" json:"is_close,omitempty"`
}

func (m *SubscriberMessage) Reset()                    { *m = SubscriberMessage{} }
func (m *SubscriberMessage) String() string            { return proto.CompactTextString(m) }
func (*SubscriberMessage) ProtoMessage()               {}
func (*SubscriberMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SubscriberMessage) GetInit() *SubscriberMessage_InitMessage {
	if m != nil {
		return m.Init
	}
	return nil
}

func (m *SubscriberMessage) GetAck() *SubscriberMessage_AckMessage {
	if m != nil {
		return m.Ack
	}
	return nil
}

func (m *SubscriberMessage) GetIsClose() bool {
	if m != nil {
		return m.IsClose
	}
	return false
}

type SubscriberMessage_InitMessage struct {
	Namespace     string                                      `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Topic         string                                      `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Partition     int32                                       `protobuf:"varint,3,opt,name=partition" json:"partition,omitempty"`
	StartPosition SubscriberMessage_InitMessage_StartPosition `protobuf:"varint,4,opt,name=startPosition,enum=messaging_pb.SubscriberMessage_InitMessage_StartPosition" json:"startPosition,omitempty"`
	TimestampNs   int64                                       `protobuf:"varint,5,opt,name=timestampNs" json:"timestampNs,omitempty"`
	SubscriberId  string                                      `protobuf:"bytes,6,opt,name=subscriber_id,json=subscriberId" json:"subscriber_id,omitempty"`
}

func (m *SubscriberMessage_InitMessage) Reset()         { *m = SubscriberMessage_InitMessage{} }
func (m *SubscriberMessage_InitMessage) String() string { return proto.CompactTextString(m) }
func (*SubscriberMessage_InitMessage) ProtoMessage()    {}
func (*SubscriberMessage_InitMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

func (m *SubscriberMessage_InitMessage) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SubscriberMessage_InitMessage) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SubscriberMessage_InitMessage) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *SubscriberMessage_InitMessage) GetStartPosition() SubscriberMessage_InitMessage_StartPosition {
	if m != nil {
		return m.StartPosition
	}
	return SubscriberMessage_InitMessage_LATEST
}

func (m *SubscriberMessage_InitMessage) GetTimestampNs() int64 {
	if m != nil {
		return m.TimestampNs
	}
	return 0
}

func (m *SubscriberMessage_InitMessage) GetSubscriberId() string {
	if m != nil {
		return m.SubscriberId
	}
	return ""
}

type SubscriberMessage_AckMessage struct {
	MessageId int64 `protobuf:"varint,1,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
}

func (m *SubscriberMessage_AckMessage) Reset()                    { *m = SubscriberMessage_AckMessage{} }
func (m *SubscriberMessage_AckMessage) String() string            { return proto.CompactTextString(m) }
func (*SubscriberMessage_AckMessage) ProtoMessage()               {}
func (*SubscriberMessage_AckMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *SubscriberMessage_AckMessage) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type Message struct {
	EventTimeNs int64             `protobuf:"varint,1,opt,name=event_time_ns,json=eventTimeNs" json:"event_time_ns,omitempty"`
	Key         []byte            `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value       []byte            `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Headers     map[string][]byte `protobuf:"bytes,4,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IsClose     bool              `protobuf:"varint,5,opt,name=is_close,json=isClose" json:"is_close,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Message) GetEventTimeNs() int64 {
	if m != nil {
		return m.EventTimeNs
	}
	return 0
}

func (m *Message) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Message) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Message) GetHeaders() map[string][]byte {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Message) GetIsClose() bool {
	if m != nil {
		return m.IsClose
	}
	return false
}

type BrokerMessage struct {
	Data *Message `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *BrokerMessage) Reset()                    { *m = BrokerMessage{} }
func (m *BrokerMessage) String() string            { return proto.CompactTextString(m) }
func (*BrokerMessage) ProtoMessage()               {}
func (*BrokerMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BrokerMessage) GetData() *Message {
	if m != nil {
		return m.Data
	}
	return nil
}

type PublishRequest struct {
	Init *PublishRequest_InitMessage `protobuf:"bytes,1,opt,name=init" json:"init,omitempty"`
	Data *Message                    `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *PublishRequest) Reset()                    { *m = PublishRequest{} }
func (m *PublishRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()               {}
func (*PublishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PublishRequest) GetInit() *PublishRequest_InitMessage {
	if m != nil {
		return m.Init
	}
	return nil
}

func (m *PublishRequest) GetData() *Message {
	if m != nil {
		return m.Data
	}
	return nil
}

type PublishRequest_InitMessage struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Topic     string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Partition int32  `protobuf:"varint,3,opt,name=partition" json:"partition,omitempty"`
}

func (m *PublishRequest_InitMessage) Reset()                    { *m = PublishRequest_InitMessage{} }
func (m *PublishRequest_InitMessage) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest_InitMessage) ProtoMessage()               {}
func (*PublishRequest_InitMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *PublishRequest_InitMessage) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PublishRequest_InitMessage) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishRequest_InitMessage) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

type PublishResponse struct {
	Config   *PublishResponse_ConfigMessage   `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	Redirect *PublishResponse_RedirectMessage `protobuf:"bytes,2,opt,name=redirect" json:"redirect,omitempty"`
	IsClosed bool                             `protobuf:"varint,3,opt,name=is_closed,json=isClosed" json:"is_closed,omitempty"`
}

func (m *PublishResponse) Reset()                    { *m = PublishResponse{} }
func (m *PublishResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()               {}
func (*PublishResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PublishResponse) GetConfig() *PublishResponse_ConfigMessage {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *PublishResponse) GetRedirect() *PublishResponse_RedirectMessage {
	if m != nil {
		return m.Redirect
	}
	return nil
}

func (m *PublishResponse) GetIsClosed() bool {
	if m != nil {
		return m.IsClosed
	}
	return false
}

type PublishResponse_ConfigMessage struct {
	PartitionCount int32 `protobuf:"varint,1,opt,name=partition_count,json=partitionCount" json:"partition_count,omitempty"`
}

func (m *PublishResponse_ConfigMessage) Reset()         { *m = PublishResponse_ConfigMessage{} }
func (m *PublishResponse_ConfigMessage) String() string { return proto.CompactTextString(m) }
func (*PublishResponse_ConfigMessage) ProtoMessage()    {}
func (*PublishResponse_ConfigMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

func (m *PublishResponse_ConfigMessage) GetPartitionCount() int32 {
	if m != nil {
		return m.PartitionCount
	}
	return 0
}

type PublishResponse_RedirectMessage struct {
	NewBroker string `protobuf:"bytes,1,opt,name=new_broker,json=newBroker" json:"new_broker,omitempty"`
}

func (m *PublishResponse_RedirectMessage) Reset()         { *m = PublishResponse_RedirectMessage{} }
func (m *PublishResponse_RedirectMessage) String() string { return proto.CompactTextString(m) }
func (*PublishResponse_RedirectMessage) ProtoMessage()    {}
func (*PublishResponse_RedirectMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 1}
}

func (m *PublishResponse_RedirectMessage) GetNewBroker() string {
	if m != nil {
		return m.NewBroker
	}
	return ""
}

type DeleteTopicRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Topic     string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
}

func (m *DeleteTopicRequest) Reset()                    { *m = DeleteTopicRequest{} }
func (m *DeleteTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTopicRequest) ProtoMessage()               {}
func (*DeleteTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteTopicRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeleteTopicRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type DeleteTopicResponse struct {
}

func (m *DeleteTopicResponse) Reset()                    { *m = DeleteTopicResponse{} }
func (m *DeleteTopicResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteTopicResponse) ProtoMessage()               {}
func (*DeleteTopicResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ConfigureTopicRequest struct {
	Namespace     string              `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Topic         string              `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Configuration *TopicConfiguration `protobuf:"bytes,3,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *ConfigureTopicRequest) Reset()                    { *m = ConfigureTopicRequest{} }
func (m *ConfigureTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigureTopicRequest) ProtoMessage()               {}
func (*ConfigureTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ConfigureTopicRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ConfigureTopicRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ConfigureTopicRequest) GetConfiguration() *TopicConfiguration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

type ConfigureTopicResponse struct {
}

func (m *ConfigureTopicResponse) Reset()                    { *m = ConfigureTopicResponse{} }
func (m *ConfigureTopicResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigureTopicResponse) ProtoMessage()               {}
func (*ConfigureTopicResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type GetTopicConfigurationRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Topic     string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
}

func (m *GetTopicConfigurationRequest) Reset()                    { *m = GetTopicConfigurationRequest{} }
func (m *GetTopicConfigurationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTopicConfigurationRequest) ProtoMessage()               {}
func (*GetTopicConfigurationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetTopicConfigurationRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetTopicConfigurationRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type GetTopicConfigurationResponse struct {
	Configuration *TopicConfiguration `protobuf:"bytes,1,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *GetTopicConfigurationResponse) Reset()                    { *m = GetTopicConfigurationResponse{} }
func (m *GetTopicConfigurationResponse) String() string            { return proto.CompactTextString(m) }
func (*GetTopicConfigurationResponse) ProtoMessage()               {}
func (*GetTopicConfigurationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetTopicConfigurationResponse) GetConfiguration() *TopicConfiguration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

type FindBrokerRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Topic     string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Parition  int32  `protobuf:"varint,3,opt,name=parition" json:"parition,omitempty"`
}

func (m *FindBrokerRequest) Reset()                    { *m = FindBrokerRequest{} }
func (m *FindBrokerRequest) String() string            { return proto.CompactTextString(m) }
func (*FindBrokerRequest) ProtoMessage()               {}
func (*FindBrokerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *FindBrokerRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *FindBrokerRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *FindBrokerRequest) GetParition() int32 {
	if m != nil {
		return m.Parition
	}
	return 0
}

type FindBrokerResponse struct {
	Broker string `protobuf:"bytes,1,opt,name=broker" json:"broker,omitempty"`
}

func (m *FindBrokerResponse) Reset()                    { *m = FindBrokerResponse{} }
func (m *FindBrokerResponse) String() string            { return proto.CompactTextString(m) }
func (*FindBrokerResponse) ProtoMessage()               {}
func (*FindBrokerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *FindBrokerResponse) GetBroker() string {
	if m != nil {
		return m.Broker
	}
	return ""
}

type TopicConfiguration struct {
	PartitionCount int32                           `protobuf:"varint,1,opt,name=partition_count,json=partitionCount" json:"partition_count,omitempty"`
	Collection     string                          `protobuf:"bytes,2,opt,name=collection" json:"collection,omitempty"`
	Replication    string                          `protobuf:"bytes,3,opt,name=replication" json:"replication,omitempty"`
	IsTransient    bool                            `protobuf:"varint,4,opt,name=is_transient,json=isTransient" json:"is_transient,omitempty"`
	Partitoning    TopicConfiguration_Partitioning `protobuf:"varint,5,opt,name=partitoning,enum=messaging_pb.TopicConfiguration_Partitioning" json:"partitoning,omitempty"`
}

func (m *TopicConfiguration) Reset()                    { *m = TopicConfiguration{} }
func (m *TopicConfiguration) String() string            { return proto.CompactTextString(m) }
func (*TopicConfiguration) ProtoMessage()               {}
func (*TopicConfiguration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *TopicConfiguration) GetPartitionCount() int32 {
	if m != nil {
		return m.PartitionCount
	}
	return 0
}

func (m *TopicConfiguration) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *TopicConfiguration) GetReplication() string {
	if m != nil {
		return m.Replication
	}
	return ""
}

func (m *TopicConfiguration) GetIsTransient() bool {
	if m != nil {
		return m.IsTransient
	}
	return false
}

func (m *TopicConfiguration) GetPartitoning() TopicConfiguration_Partitioning {
	if m != nil {
		return m.Partitoning
	}
	return TopicConfiguration_NonNullKeyHash
}

func init() {
	proto.RegisterType((*SubscriberMessage)(nil), "messaging_pb.SubscriberMessage")
	proto.RegisterType((*SubscriberMessage_InitMessage)(nil), "messaging_pb.SubscriberMessage.InitMessage")
	proto.RegisterType((*SubscriberMessage_AckMessage)(nil), "messaging_pb.SubscriberMessage.AckMessage")
	proto.RegisterType((*Message)(nil), "messaging_pb.Message")
	proto.RegisterType((*BrokerMessage)(nil), "messaging_pb.BrokerMessage")
	proto.RegisterType((*PublishRequest)(nil), "messaging_pb.PublishRequest")
	proto.RegisterType((*PublishRequest_InitMessage)(nil), "messaging_pb.PublishRequest.InitMessage")
	proto.RegisterType((*PublishResponse)(nil), "messaging_pb.PublishResponse")
	proto.RegisterType((*PublishResponse_ConfigMessage)(nil), "messaging_pb.PublishResponse.ConfigMessage")
	proto.RegisterType((*PublishResponse_RedirectMessage)(nil), "messaging_pb.PublishResponse.RedirectMessage")
	proto.RegisterType((*DeleteTopicRequest)(nil), "messaging_pb.DeleteTopicRequest")
	proto.RegisterType((*DeleteTopicResponse)(nil), "messaging_pb.DeleteTopicResponse")
	proto.RegisterType((*ConfigureTopicRequest)(nil), "messaging_pb.ConfigureTopicRequest")
	proto.RegisterType((*ConfigureTopicResponse)(nil), "messaging_pb.ConfigureTopicResponse")
	proto.RegisterType((*GetTopicConfigurationRequest)(nil), "messaging_pb.GetTopicConfigurationRequest")
	proto.RegisterType((*GetTopicConfigurationResponse)(nil), "messaging_pb.GetTopicConfigurationResponse")
	proto.RegisterType((*FindBrokerRequest)(nil), "messaging_pb.FindBrokerRequest")
	proto.RegisterType((*FindBrokerResponse)(nil), "messaging_pb.FindBrokerResponse")
	proto.RegisterType((*TopicConfiguration)(nil), "messaging_pb.TopicConfiguration")
	proto.RegisterEnum("messaging_pb.SubscriberMessage_InitMessage_StartPosition", SubscriberMessage_InitMessage_StartPosition_name, SubscriberMessage_InitMessage_StartPosition_value)
	proto.RegisterEnum("messaging_pb.TopicConfiguration_Partitioning", TopicConfiguration_Partitioning_name, TopicConfiguration_Partitioning_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SeaweedMessaging service

type SeaweedMessagingClient interface {
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_SubscribeClient, error)
	Publish(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_PublishClient, error)
	DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*DeleteTopicResponse, error)
	ConfigureTopic(ctx context.Context, in *ConfigureTopicRequest, opts ...grpc.CallOption) (*ConfigureTopicResponse, error)
	GetTopicConfiguration(ctx context.Context, in *GetTopicConfigurationRequest, opts ...grpc.CallOption) (*GetTopicConfigurationResponse, error)
	FindBroker(ctx context.Context, in *FindBrokerRequest, opts ...grpc.CallOption) (*FindBrokerResponse, error)
}

type seaweedMessagingClient struct {
	cc *grpc.ClientConn
}

func NewSeaweedMessagingClient(cc *grpc.ClientConn) SeaweedMessagingClient {
	return &seaweedMessagingClient{cc}
}

func (c *seaweedMessagingClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_SubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SeaweedMessaging_serviceDesc.Streams[0], c.cc, "/messaging_pb.SeaweedMessaging/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedMessagingSubscribeClient{stream}
	return x, nil
}

type SeaweedMessaging_SubscribeClient interface {
	Send(*SubscriberMessage) error
	Recv() (*BrokerMessage, error)
	grpc.ClientStream
}

type seaweedMessagingSubscribeClient struct {
	grpc.ClientStream
}

func (x *seaweedMessagingSubscribeClient) Send(m *SubscriberMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *seaweedMessagingSubscribeClient) Recv() (*BrokerMessage, error) {
	m := new(BrokerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedMessagingClient) Publish(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_PublishClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SeaweedMessaging_serviceDesc.Streams[1], c.cc, "/messaging_pb.SeaweedMessaging/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedMessagingPublishClient{stream}
	return x, nil
}

type SeaweedMessaging_PublishClient interface {
	Send(*PublishRequest) error
	Recv() (*PublishResponse, error)
	grpc.ClientStream
}

type seaweedMessagingPublishClient struct {
	grpc.ClientStream
}

func (x *seaweedMessagingPublishClient) Send(m *PublishRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *seaweedMessagingPublishClient) Recv() (*PublishResponse, error) {
	m := new(PublishResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedMessagingClient) DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*DeleteTopicResponse, error) {
	out := new(DeleteTopicResponse)
	err := grpc.Invoke(ctx, "/messaging_pb.SeaweedMessaging/DeleteTopic", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) ConfigureTopic(ctx context.Context, in *ConfigureTopicRequest, opts ...grpc.CallOption) (*ConfigureTopicResponse, error) {
	out := new(ConfigureTopicResponse)
	err := grpc.Invoke(ctx, "/messaging_pb.SeaweedMessaging/ConfigureTopic", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) GetTopicConfiguration(ctx context.Context, in *GetTopicConfigurationRequest, opts ...grpc.CallOption) (*GetTopicConfigurationResponse, error) {
	out := new(GetTopicConfigurationResponse)
	err := grpc.Invoke(ctx, "/messaging_pb.SeaweedMessaging/GetTopicConfiguration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) FindBroker(ctx context.Context, in *FindBrokerRequest, opts ...grpc.CallOption) (*FindBrokerResponse, error) {
	out := new(FindBrokerResponse)
	err := grpc.Invoke(ctx, "/messaging_pb.SeaweedMessaging/FindBroker", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SeaweedMessaging service

type SeaweedMessagingServer interface {
	Subscribe(SeaweedMessaging_SubscribeServer) error
	Publish(SeaweedMessaging_PublishServer) error
	DeleteTopic(context.Context, *DeleteTopicRequest) (*DeleteTopicResponse, error)
	ConfigureTopic(context.Context, *ConfigureTopicRequest) (*ConfigureTopicResponse, error)
	GetTopicConfiguration(context.Context, *GetTopicConfigurationRequest) (*GetTopicConfigurationResponse, error)
	FindBroker(context.Context, *FindBrokerRequest) (*FindBrokerResponse, error)
}

func RegisterSeaweedMessagingServer(s *grpc.Server, srv SeaweedMessagingServer) {
	s.RegisterService(&_SeaweedMessaging_serviceDesc, srv)
}

func _SeaweedMessaging_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SeaweedMessagingServer).Subscribe(&seaweedMessagingSubscribeServer{stream})
}

type SeaweedMessaging_SubscribeServer interface {
	Send(*BrokerMessage) error
	Recv() (*SubscriberMessage, error)
	grpc.ServerStream
}

type seaweedMessagingSubscribeServer struct {
	grpc.ServerStream
}

func (x *seaweedMessagingSubscribeServer) Send(m *BrokerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *seaweedMessagingSubscribeServer) Recv() (*SubscriberMessage, error) {
	m := new(SubscriberMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SeaweedMessaging_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SeaweedMessagingServer).Publish(&seaweedMessagingPublishServer{stream})
}

type SeaweedMessaging_PublishServer interface {
	Send(*PublishResponse) error
	Recv() (*PublishRequest, error)
	grpc.ServerStream
}

type seaweedMessagingPublishServer struct {
	grpc.ServerStream
}

func (x *seaweedMessagingPublishServer) Send(m *PublishResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *seaweedMessagingPublishServer) Recv() (*PublishRequest, error) {
	m := new(PublishRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SeaweedMessaging_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging_pb.SeaweedMessaging/DeleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).DeleteTopic(ctx, req.(*DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_ConfigureTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).ConfigureTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging_pb.SeaweedMessaging/ConfigureTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).ConfigureTopic(ctx, req.(*ConfigureTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_GetTopicConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).GetTopicConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging_pb.SeaweedMessaging/GetTopicConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).GetTopicConfiguration(ctx, req.(*GetTopicConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_FindBroker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBrokerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).FindBroker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging_pb.SeaweedMessaging/FindBroker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).FindBroker(ctx, req.(*FindBrokerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SeaweedMessaging_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messaging_pb.SeaweedMessaging",
	HandlerType: (*SeaweedMessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteTopic",
			Handler:    _SeaweedMessaging_DeleteTopic_Handler,
		},
		{
			MethodName: "ConfigureTopic",
			Handler:    _SeaweedMessaging_ConfigureTopic_Handler,
		},
		{
			MethodName: "GetTopicConfiguration",
			Handler:    _SeaweedMessaging_GetTopicConfiguration_Handler,
		},
		{
			MethodName: "FindBroker",
			Handler:    _SeaweedMessaging_FindBroker_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _SeaweedMessaging_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Publish",
			Handler:       _SeaweedMessaging_Publish_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "messaging.proto",
}

func init() { proto.RegisterFile("messaging.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1002 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0xe3, 0x54,
	0x10, 0xae, 0xdd, 0xfc, 0x8e, 0x93, 0x34, 0x3b, 0xd0, 0x55, 0xf0, 0xb6, 0x90, 0xf5, 0x22, 0x08,
	0x14, 0xa2, 0x2a, 0xdc, 0x94, 0x6a, 0xa5, 0x55, 0x1b, 0xba, 0x34, 0xa2, 0xed, 0x86, 0x93, 0xdc,
	0x22, 0xcb, 0xb1, 0xcf, 0xa6, 0x47, 0x75, 0x8e, 0x8d, 0x8f, 0xb3, 0x55, 0x9f, 0x83, 0x7b, 0x1e,
	0x00, 0x89, 0x3b, 0x5e, 0x80, 0xd7, 0xe0, 0x21, 0x78, 0x06, 0xe4, 0xdf, 0xd8, 0x49, 0x36, 0x5d,
	0xb6, 0xda, 0xbb, 0x9c, 0xc9, 0x37, 0x33, 0xdf, 0x99, 0xf9, 0x66, 0x8e, 0x61, 0x67, 0x46, 0x85,
	0x30, 0xa6, 0x8c, 0x4f, 0xbb, 0xae, 0xe7, 0xf8, 0x0e, 0xd6, 0x52, 0x83, 0xee, 0x4e, 0xb4, 0xdf,
	0x0b, 0xf0, 0x68, 0x34, 0x9f, 0x08, 0xd3, 0x63, 0x13, 0xea, 0x5d, 0x86, 0x7f, 0x51, 0x7c, 0x01,
	0x05, 0xc6, 0x99, 0xdf, 0x92, 0xda, 0x52, 0x47, 0xe9, 0x1d, 0x74, 0xb3, 0x2e, 0xdd, 0x15, 0x78,
	0x77, 0xc0, 0x99, 0x1f, 0xff, 0x26, 0xa1, 0x23, 0x3e, 0x87, 0x6d, 0xc3, 0xbc, 0x69, 0xc9, 0xa1,
	0xff, 0xd7, 0xf7, 0xf9, 0x9f, 0x98, 0x37, 0x89, 0x7b, 0xe0, 0x86, 0x9f, 0x40, 0x85, 0x09, 0xdd,
	0xb4, 0x1d, 0x41, 0x5b, 0xdb, 0x6d, 0xa9, 0x53, 0x21, 0x65, 0x26, 0xfa, 0xc1, 0x51, 0xfd, 0x5b,
	0x06, 0x25, 0x93, 0x0e, 0xf7, 0xa0, 0xca, 0x8d, 0x19, 0x15, 0xae, 0x61, 0xd2, 0x90, 0x6e, 0x95,
	0x2c, 0x0c, 0xf8, 0x31, 0x14, 0x7d, 0xc7, 0x65, 0x66, 0x48, 0xa4, 0x4a, 0xa2, 0x43, 0xe0, 0xe3,
	0x1a, 0x9e, 0xcf, 0x7c, 0xe6, 0xf0, 0x30, 0x7e, 0x91, 0x2c, 0x0c, 0xa8, 0x43, 0x5d, 0xf8, 0x86,
	0xe7, 0x0f, 0x1d, 0x11, 0x21, 0x0a, 0x6d, 0xa9, 0xd3, 0xe8, 0x7d, 0xff, 0x3f, 0x8a, 0xd0, 0x1d,
	0x65, 0x03, 0x90, 0x7c, 0x3c, 0x6c, 0x83, 0xe2, 0xb3, 0x19, 0x15, 0xbe, 0x31, 0x73, 0xaf, 0x44,
	0xab, 0xd8, 0x96, 0x3a, 0xdb, 0x24, 0x6b, 0xc2, 0x67, 0x50, 0x17, 0x69, 0x7c, 0x9d, 0x59, 0xad,
	0x52, 0x48, 0xbf, 0xb6, 0x30, 0x0e, 0x2c, 0xed, 0x08, 0xea, 0xb9, 0x34, 0x08, 0x50, 0xba, 0x38,
	0x19, 0x9f, 0x8d, 0xc6, 0xcd, 0x2d, 0xac, 0x41, 0xe5, 0xec, 0x84, 0x5c, 0x0c, 0x82, 0x93, 0x84,
	0x75, 0xa8, 0x8e, 0x07, 0x97, 0x67, 0xa3, 0xf1, 0xc9, 0xe5, 0xb0, 0x29, 0xab, 0x07, 0x00, 0x8b,
	0x8a, 0xe3, 0x3e, 0x40, 0x74, 0x33, 0x1a, 0x64, 0x92, 0x42, 0x36, 0xd5, 0xd8, 0x32, 0xb0, 0xb4,
	0x7f, 0x25, 0x28, 0x27, 0xd0, 0x2f, 0xa0, 0x4e, 0xdf, 0x50, 0xee, 0xeb, 0x01, 0x59, 0x9d, 0x8b,
	0x08, 0x7d, 0x2a, 0x1f, 0x4a, 0x44, 0x09, 0xff, 0x18, 0xb3, 0x19, 0xbd, 0x12, 0xd8, 0x84, 0xed,
	0x1b, 0x7a, 0x17, 0x16, 0xbd, 0x46, 0x82, 0x9f, 0x41, 0x23, 0xde, 0x18, 0xf6, 0x3c, 0x6a, 0x67,
	0x8d, 0x44, 0x07, 0x7c, 0x0e, 0xe5, 0x6b, 0x6a, 0x58, 0xd4, 0x13, 0xad, 0x42, 0x7b, 0xbb, 0xa3,
	0xf4, 0xb4, 0x7c, 0x91, 0x93, 0x72, 0x9e, 0x47, 0xa0, 0x33, 0xee, 0x7b, 0x77, 0x24, 0x71, 0xc9,
	0xa9, 0xa4, 0x98, 0x57, 0xc9, 0x31, 0xd4, 0xb2, 0x3e, 0x09, 0xa1, 0x48, 0x1f, 0x79, 0x42, 0x72,
	0x86, 0xd0, 0xb1, 0x7c, 0x24, 0x69, 0xc7, 0x50, 0x3f, 0xf5, 0x9c, 0x9b, 0xc5, 0x30, 0x7c, 0x05,
	0x05, 0xcb, 0xf0, 0x8d, 0x78, 0x18, 0x76, 0xd7, 0x52, 0x24, 0x21, 0x44, 0xfb, 0x47, 0x82, 0xc6,
	0x70, 0x3e, 0xb1, 0x99, 0xb8, 0x26, 0xf4, 0xd7, 0x39, 0x15, 0xc1, 0x24, 0x64, 0x47, 0xa9, 0x93,
	0xf7, 0xce, 0x63, 0xd7, 0xcc, 0x51, 0x92, 0x5b, 0xbe, 0x37, 0xb7, 0xaa, 0x7f, 0xe0, 0xc1, 0xd0,
	0xfe, 0x90, 0x61, 0x27, 0x25, 0x2c, 0x5c, 0x87, 0x0b, 0x8a, 0x7d, 0x28, 0x99, 0x0e, 0x7f, 0xcd,
	0xa6, 0xeb, 0x57, 0xc5, 0x12, 0xbc, 0xdb, 0x0f, 0xb1, 0x09, 0xef, 0xd8, 0x15, 0x07, 0x50, 0xf1,
	0xa8, 0xc5, 0x3c, 0x6a, 0xfa, 0xf1, 0x45, 0xbf, 0xdd, 0x1c, 0x86, 0xc4, 0xe8, 0x24, 0x50, 0xea,
	0x8e, 0x4f, 0xa0, 0x9a, 0x68, 0xc2, 0x8a, 0x57, 0x47, 0x25, 0x16, 0x85, 0xa5, 0x1e, 0x41, 0x3d,
	0x47, 0x00, 0xbf, 0x84, 0x9d, 0xf4, 0x7a, 0xba, 0xe9, 0xcc, 0x79, 0xd4, 0xa6, 0x22, 0x69, 0xa4,
	0xe6, 0x7e, 0x60, 0x55, 0x0f, 0x61, 0x67, 0x29, 0x67, 0x30, 0x36, 0x9c, 0xde, 0xea, 0x93, 0x50,
	0x2a, 0x69, 0x81, 0xe9, 0x6d, 0xa4, 0x1d, 0xed, 0x1c, 0xf0, 0x07, 0x6a, 0x53, 0x9f, 0x8e, 0x83,
	0xca, 0x26, 0x62, 0x78, 0x8f, 0xa6, 0x68, 0xbb, 0xf0, 0x51, 0x2e, 0x52, 0x54, 0x03, 0xed, 0x37,
	0x09, 0x76, 0xa3, 0xdb, 0xcc, 0xbd, 0x07, 0x27, 0xc1, 0x97, 0x50, 0x37, 0xe3, 0x60, 0x46, 0xda,
	0x7d, 0xa5, 0xd7, 0xce, 0xf7, 0x21, 0x4c, 0xd3, 0xcf, 0xe2, 0x48, 0xde, 0x4d, 0x6b, 0xc1, 0xe3,
	0x65, 0x52, 0x31, 0x5f, 0x02, 0x7b, 0x3f, 0x52, 0x7f, 0x4d, 0x84, 0x07, 0x94, 0x66, 0x0a, 0xfb,
	0x6f, 0x89, 0x19, 0xcb, 0x73, 0xe5, 0x5a, 0xd2, 0xfb, 0x5d, 0xcb, 0x84, 0x47, 0x2f, 0x19, 0xb7,
	0xa2, 0xde, 0x3e, 0xa4, 0xce, 0x2a, 0x54, 0x5c, 0xc3, 0xcb, 0x0e, 0x58, 0x7a, 0xd6, 0xbe, 0x01,
	0xcc, 0x26, 0x89, 0xaf, 0xf0, 0x18, 0x4a, 0x39, 0x8d, 0xc5, 0x27, 0xed, 0x2f, 0x19, 0x70, 0x95,
	0xf8, 0x3b, 0x4b, 0x1a, 0x3f, 0x05, 0x30, 0x1d, 0xdb, 0xa6, 0x66, 0xc8, 0x25, 0x22, 0x99, 0xb1,
	0x04, 0xaf, 0x94, 0x47, 0x5d, 0x9b, 0x99, 0x0b, 0x3d, 0x54, 0x49, 0xd6, 0x84, 0x4f, 0xa1, 0xc6,
	0x84, 0xee, 0x7b, 0x06, 0x17, 0x8c, 0x72, 0x3f, 0x7c, 0x27, 0x2b, 0x44, 0x61, 0x62, 0x9c, 0x98,
	0xf0, 0x15, 0x28, 0x51, 0x5a, 0x87, 0x33, 0x3e, 0x0d, 0xb7, 0x74, 0x63, 0x79, 0xb8, 0x57, 0x2f,
	0xd1, 0x1d, 0x26, 0x54, 0x19, 0x9f, 0x92, 0x6c, 0x04, 0xed, 0x05, 0xd4, 0xb2, 0x7f, 0x22, 0x42,
	0xe3, 0xca, 0xe1, 0x57, 0x73, 0xdb, 0xfe, 0x89, 0xde, 0x9d, 0x1b, 0xe2, 0xba, 0xb9, 0x85, 0x0a,
	0x94, 0x93, 0x83, 0x84, 0x0d, 0x00, 0xe2, 0xcc, 0xb9, 0x45, 0x9c, 0x09, 0xe3, 0x4d, 0xb9, 0xf7,
	0x67, 0x01, 0x9a, 0x23, 0x6a, 0xdc, 0x52, 0x6a, 0x5d, 0x26, 0x2c, 0xf0, 0x15, 0x54, 0xd3, 0xf7,
	0x1c, 0x3f, 0xbb, 0xe7, 0xa1, 0x57, 0x9f, 0xe4, 0x01, 0xb9, 0xc7, 0x42, 0xdb, 0xea, 0x48, 0x87,
	0x12, 0x5e, 0x40, 0x39, 0xde, 0x59, 0xb8, 0xb7, 0x69, 0xe3, 0xab, 0xfb, 0x1b, 0x17, 0x5d, 0x1c,
	0x6d, 0x0c, 0x4a, 0x66, 0x03, 0xe0, 0x92, 0x7a, 0x57, 0xd7, 0x8c, 0xfa, 0x74, 0x03, 0x22, 0x89,
	0x8c, 0xbf, 0x40, 0x23, 0x3f, 0xaa, 0xf8, 0x2c, 0xef, 0xb6, 0x76, 0xbb, 0xa8, 0x9f, 0x6f, 0x06,
	0xa5, 0xe1, 0x3d, 0xd8, 0x5d, 0x3b, 0x9b, 0xb8, 0xf4, 0x35, 0xb8, 0x69, 0x29, 0xa8, 0x07, 0xef,
	0x84, 0x4d, 0x73, 0xfe, 0x0c, 0xb0, 0x98, 0xa0, 0xe5, 0x46, 0xae, 0x0c, 0xb0, 0xda, 0x7e, 0x3b,
	0x20, 0x09, 0x79, 0xaa, 0x41, 0x53, 0x44, 0x72, 0x79, 0x2d, 0xba, 0xa6, 0x1d, 0xa8, 0xfa, 0xb4,
	0x91, 0x2a, 0x67, 0x18, 0x7c, 0x51, 0x4f, 0x4a, 0xe1, 0x87, 0xf5, 0x77, 0xff, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x7f, 0x62, 0xba, 0x48, 0x6b, 0x0b, 0x00, 0x00,
}
