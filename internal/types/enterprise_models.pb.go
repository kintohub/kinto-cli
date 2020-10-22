// Code generated by protoc-gen-go. DO NOT EDIT.
// source: enterprise_models.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Subscription state != Stripe's states. We will not need extra information such as cancelled, etc.
// This is simple paying, free, and past due states that are used for our systems and translated from
// stripe's statuses or future billing statuses that may drive our payments system.
type Account_SubscriptionState int32

const (
	Account_NOT_SET Account_SubscriptionState = 0
	// Free user or cancelled subscription. This is the default state
	Account_FREE Account_SubscriptionState = 1
	// Has an active and valid credit card on the subscription
	Account_ACTIVE Account_SubscriptionState = 2
	// Active, but the payment was not made and is past due. (Invalid CC, blocked our payment, etc)
	Account_PAST_DUE Account_SubscriptionState = 3
	// Payment cancelled by us/them, need to lock the account for manual actions
	Account_CANCELLED Account_SubscriptionState = 4
)

var Account_SubscriptionState_name = map[int32]string{
	0: "NOT_SET",
	1: "FREE",
	2: "ACTIVE",
	3: "PAST_DUE",
	4: "CANCELLED",
}

var Account_SubscriptionState_value = map[string]int32{
	"NOT_SET":   0,
	"FREE":      1,
	"ACTIVE":    2,
	"PAST_DUE":  3,
	"CANCELLED": 4,
}

func (x Account_SubscriptionState) String() string {
	return proto.EnumName(Account_SubscriptionState_name, int32(x))
}

func (Account_SubscriptionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_532dd47d300bf0f4, []int{1, 0}
}

type ValidationRequest struct {
	Email                string               `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Token                string               `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ExpireAt             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expireAt,proto3" json:"expireAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ValidationRequest) Reset()         { *m = ValidationRequest{} }
func (m *ValidationRequest) String() string { return proto.CompactTextString(m) }
func (*ValidationRequest) ProtoMessage()    {}
func (*ValidationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_532dd47d300bf0f4, []int{0}
}

func (m *ValidationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationRequest.Unmarshal(m, b)
}
func (m *ValidationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationRequest.Marshal(b, m, deterministic)
}
func (m *ValidationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationRequest.Merge(m, src)
}
func (m *ValidationRequest) XXX_Size() int {
	return xxx_messageInfo_ValidationRequest.Size(m)
}
func (m *ValidationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationRequest proto.InternalMessageInfo

func (m *ValidationRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ValidationRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ValidationRequest) GetExpireAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireAt
	}
	return nil
}

type Account struct {
	Id                             string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                          string                    `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password                       string                    `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	CreatedAt                      *timestamp.Timestamp      `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	LastSeenAt                     *timestamp.Timestamp      `protobuf:"bytes,5,opt,name=lastSeenAt,proto3" json:"lastSeenAt,omitempty"`
	ResetPassRequest               *ValidationRequest        `protobuf:"bytes,6,opt,name=resetPassRequest,proto3" json:"resetPassRequest,omitempty"`
	ValidateEmailRequest           *ValidationRequest        `protobuf:"bytes,7,opt,name=validateEmailRequest,proto3" json:"validateEmailRequest,omitempty"`
	IsEmailValidated               bool                      `protobuf:"varint,8,opt,name=isEmailValidated,proto3" json:"isEmailValidated,omitempty"`
	MaxAllowedEnvironments         int32                     `protobuf:"varint,9,opt,name=maxAllowedEnvironments,proto3" json:"maxAllowedEnvironments,omitempty"`
	StripeCustomerId               string                    `protobuf:"bytes,10,opt,name=stripeCustomerId,proto3" json:"stripeCustomerId,omitempty"`
	StripeSubscriptionId           string                    `protobuf:"bytes,11,opt,name=stripeSubscriptionId,proto3" json:"stripeSubscriptionId,omitempty"`
	StripeSubscriptionHash         string                    `protobuf:"bytes,12,opt,name=stripeSubscriptionHash,proto3" json:"stripeSubscriptionHash,omitempty"`
	StripePriceSubscriptionItemIds map[string]string         `protobuf:"bytes,13,rep,name=stripePriceSubscriptionItemIds,proto3" json:"stripePriceSubscriptionItemIds,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SubscriptionState              Account_SubscriptionState `protobuf:"varint,14,opt,name=subscriptionState,proto3,enum=Account_SubscriptionState" json:"subscriptionState,omitempty"`
	DisplayName                    string                    `protobuf:"bytes,15,opt,name=displayName,proto3" json:"displayName,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}                  `json:"-"`
	XXX_unrecognized               []byte                    `json:"-"`
	XXX_sizecache                  int32                     `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_532dd47d300bf0f4, []int{1}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Account) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Account) GetLastSeenAt() *timestamp.Timestamp {
	if m != nil {
		return m.LastSeenAt
	}
	return nil
}

func (m *Account) GetResetPassRequest() *ValidationRequest {
	if m != nil {
		return m.ResetPassRequest
	}
	return nil
}

func (m *Account) GetValidateEmailRequest() *ValidationRequest {
	if m != nil {
		return m.ValidateEmailRequest
	}
	return nil
}

func (m *Account) GetIsEmailValidated() bool {
	if m != nil {
		return m.IsEmailValidated
	}
	return false
}

func (m *Account) GetMaxAllowedEnvironments() int32 {
	if m != nil {
		return m.MaxAllowedEnvironments
	}
	return 0
}

func (m *Account) GetStripeCustomerId() string {
	if m != nil {
		return m.StripeCustomerId
	}
	return ""
}

func (m *Account) GetStripeSubscriptionId() string {
	if m != nil {
		return m.StripeSubscriptionId
	}
	return ""
}

func (m *Account) GetStripeSubscriptionHash() string {
	if m != nil {
		return m.StripeSubscriptionHash
	}
	return ""
}

func (m *Account) GetStripePriceSubscriptionItemIds() map[string]string {
	if m != nil {
		return m.StripePriceSubscriptionItemIds
	}
	return nil
}

func (m *Account) GetSubscriptionState() Account_SubscriptionState {
	if m != nil {
		return m.SubscriptionState
	}
	return Account_NOT_SET
}

func (m *Account) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type Session struct {
	Id                   string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId            string                    `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
	CreatedAt            *timestamp.Timestamp      `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ExpiresAt            *timestamp.Timestamp      `protobuf:"bytes,4,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
	SubscriptionState    Account_SubscriptionState `protobuf:"varint,5,opt,name=subscriptionState,proto3,enum=Account_SubscriptionState" json:"subscriptionState,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_532dd47d300bf0f4, []int{2}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Session) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Session) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Session) GetExpiresAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

func (m *Session) GetSubscriptionState() Account_SubscriptionState {
	if m != nil {
		return m.SubscriptionState
	}
	return Account_NOT_SET
}

type Me struct {
	Id                   string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IsEmailValidated     bool                      `protobuf:"varint,2,opt,name=isEmailValidated,proto3" json:"isEmailValidated,omitempty"`
	Email                string                    `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	SubscriptionState    Account_SubscriptionState `protobuf:"varint,4,opt,name=subscriptionState,proto3,enum=Account_SubscriptionState" json:"subscriptionState,omitempty"`
	DisplayName          string                    `protobuf:"bytes,5,opt,name=displayName,proto3" json:"displayName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Me) Reset()         { *m = Me{} }
func (m *Me) String() string { return proto.CompactTextString(m) }
func (*Me) ProtoMessage()    {}
func (*Me) Descriptor() ([]byte, []int) {
	return fileDescriptor_532dd47d300bf0f4, []int{3}
}

func (m *Me) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Me.Unmarshal(m, b)
}
func (m *Me) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Me.Marshal(b, m, deterministic)
}
func (m *Me) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Me.Merge(m, src)
}
func (m *Me) XXX_Size() int {
	return xxx_messageInfo_Me.Size(m)
}
func (m *Me) XXX_DiscardUnknown() {
	xxx_messageInfo_Me.DiscardUnknown(m)
}

var xxx_messageInfo_Me proto.InternalMessageInfo

func (m *Me) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Me) GetIsEmailValidated() bool {
	if m != nil {
		return m.IsEmailValidated
	}
	return false
}

func (m *Me) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Me) GetSubscriptionState() Account_SubscriptionState {
	if m != nil {
		return m.SubscriptionState
	}
	return Account_NOT_SET
}

func (m *Me) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type Member struct {
	AccountId            string                `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	Role                 AccessPermission_Role `protobuf:"varint,2,opt,name=role,proto3,enum=AccessPermission_Role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_532dd47d300bf0f4, []int{4}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Member) GetRole() AccessPermission_Role {
	if m != nil {
		return m.Role
	}
	return AccessPermission_NOT_SET
}

// Internal Model Only!!! MUST NOT RETURN THIS TO THE PUBLIC
type Cluster struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DisplayName          string               `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
	HostName             string               `protobuf:"bytes,3,opt,name=hostName,proto3" json:"hostName,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	LastHealthCheck      *timestamp.Timestamp `protobuf:"bytes,5,opt,name=lastHealthCheck,proto3" json:"lastHealthCheck,omitempty"`
	ClientSecret         string               `protobuf:"bytes,6,opt,name=clientSecret,proto3" json:"clientSecret,omitempty"`
	AccessTokenSecretKey []byte               `protobuf:"bytes,7,opt,name=accessTokenSecretKey,proto3" json:"accessTokenSecretKey,omitempty"`
	WebHostName          string               `protobuf:"bytes,8,opt,name=webHostName,proto3" json:"webHostName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_532dd47d300bf0f4, []int{5}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Cluster) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Cluster) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *Cluster) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Cluster) GetLastHealthCheck() *timestamp.Timestamp {
	if m != nil {
		return m.LastHealthCheck
	}
	return nil
}

func (m *Cluster) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *Cluster) GetAccessTokenSecretKey() []byte {
	if m != nil {
		return m.AccessTokenSecretKey
	}
	return nil
}

func (m *Cluster) GetWebHostName() string {
	if m != nil {
		return m.WebHostName
	}
	return ""
}

type ClusterEnvironment struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClusterId            string               `protobuf:"bytes,2,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Members              []*Member            `protobuf:"bytes,5,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClusterEnvironment) Reset()         { *m = ClusterEnvironment{} }
func (m *ClusterEnvironment) String() string { return proto.CompactTextString(m) }
func (*ClusterEnvironment) ProtoMessage()    {}
func (*ClusterEnvironment) Descriptor() ([]byte, []int) {
	return fileDescriptor_532dd47d300bf0f4, []int{6}
}

func (m *ClusterEnvironment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterEnvironment.Unmarshal(m, b)
}
func (m *ClusterEnvironment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterEnvironment.Marshal(b, m, deterministic)
}
func (m *ClusterEnvironment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterEnvironment.Merge(m, src)
}
func (m *ClusterEnvironment) XXX_Size() int {
	return xxx_messageInfo_ClusterEnvironment.Size(m)
}
func (m *ClusterEnvironment) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterEnvironment.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterEnvironment proto.InternalMessageInfo

func (m *ClusterEnvironment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ClusterEnvironment) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ClusterEnvironment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterEnvironment) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ClusterEnvironment) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type Email struct {
	TemplateId           string            `protobuf:"bytes,1,opt,name=templateId,proto3" json:"templateId,omitempty"`
	Subject              string            `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Email                string            `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Variables            map[string]string `protobuf:"bytes,4,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Email) Reset()         { *m = Email{} }
func (m *Email) String() string { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()    {}
func (*Email) Descriptor() ([]byte, []int) {
	return fileDescriptor_532dd47d300bf0f4, []int{7}
}

func (m *Email) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Email.Unmarshal(m, b)
}
func (m *Email) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Email.Marshal(b, m, deterministic)
}
func (m *Email) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Email.Merge(m, src)
}
func (m *Email) XXX_Size() int {
	return xxx_messageInfo_Email.Size(m)
}
func (m *Email) XXX_DiscardUnknown() {
	xxx_messageInfo_Email.DiscardUnknown(m)
}

var xxx_messageInfo_Email proto.InternalMessageInfo

func (m *Email) GetTemplateId() string {
	if m != nil {
		return m.TemplateId
	}
	return ""
}

func (m *Email) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Email) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Email) GetVariables() map[string]string {
	if m != nil {
		return m.Variables
	}
	return nil
}

func init() {
	proto.RegisterEnum("Account_SubscriptionState", Account_SubscriptionState_name, Account_SubscriptionState_value)
	proto.RegisterType((*ValidationRequest)(nil), "ValidationRequest")
	proto.RegisterType((*Account)(nil), "Account")
	proto.RegisterMapType((map[string]string)(nil), "Account.StripePriceSubscriptionItemIdsEntry")
	proto.RegisterType((*Session)(nil), "Session")
	proto.RegisterType((*Me)(nil), "Me")
	proto.RegisterType((*Member)(nil), "Member")
	proto.RegisterType((*Cluster)(nil), "Cluster")
	proto.RegisterType((*ClusterEnvironment)(nil), "ClusterEnvironment")
	proto.RegisterType((*Email)(nil), "Email")
	proto.RegisterMapType((map[string]string)(nil), "Email.VariablesEntry")
}

func init() {
	proto.RegisterFile("enterprise_models.proto", fileDescriptor_532dd47d300bf0f4)
}

var fileDescriptor_532dd47d300bf0f4 = []byte{
	// 925 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x2d, 0x29, 0xc9, 0x92, 0x46, 0x8e, 0x22, 0x2f, 0xdc, 0x94, 0x10, 0x8a, 0x54, 0x65, 0x2f,
	0x42, 0x90, 0x52, 0x80, 0x02, 0x04, 0x46, 0x50, 0x14, 0x50, 0x65, 0x06, 0x16, 0x9a, 0xb8, 0x2e,
	0xa5, 0xfa, 0xd0, 0x8b, 0xb1, 0x22, 0xa7, 0xd6, 0x56, 0xfc, 0xea, 0xee, 0xd2, 0x8e, 0x7e, 0x57,
	0xff, 0x41, 0x8f, 0xbd, 0x14, 0xfd, 0x31, 0xbd, 0x17, 0x5c, 0x52, 0x5f, 0xa6, 0x54, 0x3b, 0x4d,
	0x6e, 0xdc, 0x37, 0x6f, 0x77, 0xde, 0xce, 0xce, 0x3c, 0x09, 0x3e, 0xc3, 0x50, 0x22, 0x8f, 0x39,
	0x13, 0x78, 0x15, 0x44, 0x1e, 0xfa, 0xc2, 0x8a, 0x79, 0x24, 0xa3, 0xf6, 0x17, 0xd7, 0x51, 0x74,
	0xed, 0x63, 0x4f, 0xad, 0xa6, 0xc9, 0x2f, 0x3d, 0xc9, 0x02, 0x14, 0x92, 0x06, 0x71, 0x4e, 0x68,
	0xcd, 0xe7, 0xee, 0xd6, 0x16, 0xf3, 0x16, 0x8e, 0x2e, 0xa9, 0xcf, 0x3c, 0x2a, 0x59, 0x14, 0x3a,
	0xf8, 0x5b, 0x82, 0x42, 0x92, 0x63, 0xa8, 0x60, 0x40, 0x99, 0x6f, 0x68, 0x1d, 0xad, 0x5b, 0x77,
	0xb2, 0x45, 0x8a, 0xca, 0x68, 0x8e, 0xa1, 0xa1, 0x67, 0xa8, 0x5a, 0x90, 0x97, 0x50, 0xc3, 0x77,
	0x31, 0xe3, 0x38, 0x90, 0x46, 0xa9, 0xa3, 0x75, 0x1b, 0xfd, 0xb6, 0x95, 0xc9, 0xb0, 0x96, 0x32,
	0xac, 0xc9, 0x52, 0x86, 0xb3, 0xe2, 0x9a, 0x7f, 0x55, 0xa1, 0x3a, 0x70, 0xdd, 0x28, 0x09, 0x25,
	0x69, 0x82, 0xce, 0xbc, 0x3c, 0x99, 0xce, 0xbc, 0x75, 0x7e, 0x7d, 0x33, 0x7f, 0x1b, 0x6a, 0x31,
	0x15, 0xe2, 0x36, 0xe2, 0x9e, 0xca, 0x54, 0x77, 0x56, 0x6b, 0x72, 0x02, 0x75, 0x97, 0x23, 0x95,
	0xe8, 0x0d, 0xa4, 0x51, 0xbe, 0x57, 0xc6, 0x9a, 0x4c, 0x5e, 0x01, 0xf8, 0x54, 0xc8, 0x31, 0x62,
	0x38, 0x90, 0x46, 0xe5, 0xde, 0xad, 0x1b, 0x6c, 0xf2, 0x2d, 0xb4, 0x38, 0x0a, 0x94, 0x17, 0x54,
	0x88, 0xbc, 0x76, 0xc6, 0x81, 0x3a, 0x81, 0x58, 0x85, 0xaa, 0x3a, 0x05, 0x2e, 0x79, 0x0d, 0xc7,
	0x37, 0x19, 0x0d, 0xed, 0xf4, 0x8a, 0xcb, 0x33, 0xaa, 0x7b, 0xcf, 0xd8, 0xc9, 0x27, 0xcf, 0xa0,
	0xc5, 0x84, 0x42, 0xf2, 0x1d, 0xe8, 0x19, 0xb5, 0x8e, 0xd6, 0xad, 0x39, 0x05, 0x9c, 0xbc, 0x84,
	0x27, 0x01, 0x7d, 0x37, 0xf0, 0xfd, 0xe8, 0x16, 0x3d, 0x3b, 0xbc, 0x61, 0x3c, 0x0a, 0x03, 0x0c,
	0xa5, 0x30, 0xea, 0x1d, 0xad, 0x5b, 0x71, 0xf6, 0x44, 0xd3, 0x1c, 0x42, 0x72, 0x16, 0xe3, 0x30,
	0x11, 0x32, 0x0a, 0x90, 0x8f, 0x3c, 0x03, 0xd4, 0x2b, 0x14, 0x70, 0xd2, 0x87, 0xe3, 0x0c, 0x1b,
	0x27, 0x53, 0xe1, 0x72, 0x16, 0xa7, 0x57, 0x18, 0x79, 0x46, 0x43, 0xf1, 0x77, 0xc6, 0x52, 0x5d,
	0x45, 0xfc, 0x8c, 0x8a, 0x99, 0x71, 0xa8, 0x76, 0xed, 0x89, 0x12, 0x09, 0x4f, 0xb3, 0xc8, 0x05,
	0x67, 0xee, 0xf6, 0xa1, 0x12, 0x83, 0x91, 0x27, 0x8c, 0x47, 0x9d, 0x52, 0xb7, 0xd1, 0x7f, 0x6e,
	0xe5, 0xdd, 0x66, 0x8d, 0xff, 0x93, 0x6e, 0x87, 0x92, 0x2f, 0x9c, 0x7b, 0xce, 0x24, 0x67, 0x70,
	0x24, 0x36, 0xe0, 0xb1, 0xa4, 0x12, 0x8d, 0x66, 0x47, 0xeb, 0x36, 0xfb, 0xed, 0x75, 0xa2, 0xbb,
	0x0c, 0xa7, 0xb8, 0x89, 0x74, 0xa0, 0xe1, 0x31, 0x11, 0xfb, 0x74, 0x71, 0x4e, 0x03, 0x34, 0x1e,
	0xab, 0xcb, 0x6e, 0x42, 0xed, 0x1f, 0xe1, 0xab, 0x07, 0x48, 0x26, 0x2d, 0x28, 0xcd, 0x71, 0x91,
	0x4f, 0x51, 0xfa, 0x99, 0x8e, 0xd1, 0x0d, 0xf5, 0x13, 0x5c, 0x8e, 0x91, 0x5a, 0xbc, 0xd2, 0x4f,
	0x34, 0x73, 0x0c, 0x47, 0x05, 0x71, 0xa4, 0x01, 0xd5, 0xf3, 0x1f, 0x26, 0x57, 0x63, 0x7b, 0xd2,
	0xfa, 0x84, 0xd4, 0xa0, 0xfc, 0xda, 0xb1, 0xed, 0x96, 0x46, 0x00, 0x0e, 0x06, 0xc3, 0xc9, 0xe8,
	0xd2, 0x6e, 0xe9, 0xe4, 0x10, 0x6a, 0x17, 0x83, 0xf1, 0xe4, 0xea, 0xf4, 0x27, 0xbb, 0x55, 0x22,
	0x8f, 0xa0, 0x3e, 0x1c, 0x9c, 0x0f, 0xed, 0x37, 0x6f, 0xec, 0xd3, 0x56, 0xd9, 0xfc, 0x47, 0x83,
	0xea, 0x18, 0x85, 0x60, 0x51, 0x58, 0x98, 0xe8, 0xcf, 0xa1, 0x4e, 0xb3, 0xaa, 0x8c, 0xbc, 0x5c,
	0xce, 0x1a, 0xd8, 0x9e, 0xde, 0xd2, 0xfb, 0x4c, 0xef, 0x09, 0xd4, 0x33, 0x47, 0x11, 0x0f, 0x9b,
	0xfb, 0x15, 0x79, 0xf7, 0x0b, 0x56, 0xfe, 0xc7, 0x0b, 0x9a, 0x7f, 0x68, 0xa0, 0xbf, 0xc5, 0xc2,
	0x95, 0x77, 0x0d, 0xa5, 0xbe, 0x67, 0x28, 0x57, 0x86, 0x57, 0xda, 0x34, 0xbc, 0x9d, 0x12, 0xcb,
	0x1f, 0xa1, 0xc9, 0x2a, 0x85, 0x26, 0x33, 0x1d, 0x38, 0x78, 0x8b, 0xc1, 0x14, 0xf9, 0xf6, 0x53,
	0x69, 0x77, 0x9f, 0xea, 0x19, 0x94, 0x79, 0xe4, 0x67, 0x2d, 0xd5, 0xec, 0x3f, 0x49, 0x65, 0xa0,
	0x10, 0x17, 0xc8, 0x03, 0xa6, 0x5e, 0xde, 0x72, 0x22, 0x1f, 0x1d, 0xc5, 0x31, 0xff, 0xd6, 0xa1,
	0x3a, 0xf4, 0x13, 0x21, 0x91, 0x17, 0xaa, 0x73, 0x47, 0x91, 0x5e, 0x50, 0x94, 0xda, 0xfd, 0x2c,
	0x12, 0x52, 0x85, 0x73, 0xbb, 0x5f, 0xae, 0x3f, 0xc0, 0xee, 0x4f, 0xe1, 0x71, 0x6a, 0xe0, 0x67,
	0x48, 0x7d, 0x39, 0x1b, 0xce, 0xd0, 0x9d, 0x3f, 0xc0, 0xf3, 0xef, 0x6e, 0x21, 0x26, 0x1c, 0xba,
	0x3e, 0xc3, 0x50, 0x8e, 0xd1, 0xe5, 0x98, 0x99, 0x7e, 0xdd, 0xd9, 0xc2, 0x52, 0x13, 0xa4, 0xaa,
	0x38, 0x93, 0xf4, 0x77, 0x32, 0x03, 0xbf, 0xc7, 0x85, 0x32, 0xf7, 0x43, 0x67, 0x67, 0x2c, 0xad,
	0xca, 0x2d, 0x4e, 0xcf, 0x96, 0xd7, 0xae, 0x65, 0x55, 0xd9, 0x80, 0xcc, 0xdf, 0x35, 0x20, 0x79,
	0x4d, 0x37, 0xec, 0x79, 0xd7, 0xbc, 0xb9, 0x19, 0x6b, 0x3d, 0x6f, 0x2b, 0x80, 0x10, 0x28, 0x87,
	0xeb, 0xb2, 0xaa, 0xef, 0x0f, 0x28, 0xe9, 0x97, 0x50, 0x0d, 0x54, 0xeb, 0x08, 0xa3, 0xa2, 0xac,
	0xb6, 0x6a, 0x65, 0xad, 0xe4, 0x2c, 0x71, 0xf3, 0x4f, 0x0d, 0x2a, 0xaa, 0xe5, 0xc9, 0x53, 0x00,
	0x89, 0x41, 0xec, 0x53, 0x89, 0xab, 0xf6, 0xda, 0x40, 0x88, 0x01, 0x55, 0x91, 0x4c, 0x7f, 0x45,
	0x57, 0xe6, 0xb2, 0x97, 0xcb, 0x3d, 0x33, 0xf2, 0x02, 0xea, 0x37, 0x94, 0x33, 0x3a, 0xf5, 0x51,
	0x18, 0x65, 0x95, 0xfe, 0x53, 0x4b, 0xa5, 0xb2, 0x2e, 0x97, 0x78, 0x66, 0xe9, 0x6b, 0x5e, 0xfb,
	0x1b, 0x68, 0x6e, 0x07, 0xdf, 0xc7, 0x3c, 0xbf, 0xb3, 0x7e, 0x7e, 0x7e, 0xcd, 0xe4, 0x2c, 0x99,
	0x5a, 0x6e, 0x14, 0xf4, 0xe6, 0x2c, 0x94, 0xd1, 0x2c, 0x99, 0x66, 0x1f, 0x5f, 0xbb, 0x3e, 0xeb,
	0xb1, 0xf4, 0xef, 0x59, 0x48, 0xfd, 0x9e, 0x5c, 0xc4, 0x28, 0xa6, 0x07, 0xaa, 0x7c, 0x2f, 0xfe,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x17, 0x01, 0x39, 0xb7, 0x09, 0x00, 0x00,
}
