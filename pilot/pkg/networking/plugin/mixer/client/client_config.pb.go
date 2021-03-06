// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mixer/v1/config/client/client_config.proto

// Describes the configuration state for the Mixer client library that's built into Envoy.

package client

import (
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// Describes the policy.
type NetworkFailPolicy_FailPolicy int32

const (
	// If network connection fails, request is allowed and delivered to the
	// service.
	NetworkFailPolicy_FAIL_OPEN NetworkFailPolicy_FailPolicy = 0
	// If network connection fails, request is rejected.
	NetworkFailPolicy_FAIL_CLOSE NetworkFailPolicy_FailPolicy = 1
)

var NetworkFailPolicy_FailPolicy_name = map[int32]string{
	0: "FAIL_OPEN",
	1: "FAIL_CLOSE",
}

var NetworkFailPolicy_FailPolicy_value = map[string]int32{
	"FAIL_OPEN":  0,
	"FAIL_CLOSE": 1,
}

func (x NetworkFailPolicy_FailPolicy) String() string {
	return proto.EnumName(NetworkFailPolicy_FailPolicy_name, int32(x))
}

func (NetworkFailPolicy_FailPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_27bf0dec365e2f6f, []int{0, 0}
}

// Specifies the behavior when the client is unable to connect to Mixer.
type NetworkFailPolicy struct {
	// Specifies the behavior when the client is unable to connect to Mixer.
	Policy NetworkFailPolicy_FailPolicy `protobuf:"varint,1,opt,name=policy,proto3,enum=istio.mixer.v1.config.client.NetworkFailPolicy_FailPolicy" json:"policy,omitempty"`
	// Max retries on transport error.
	MaxRetry uint32 `protobuf:"varint,2,opt,name=max_retry,json=maxRetry,proto3" json:"max_retry,omitempty"`
	// Base time to wait between retries.  Will be adjusted by exponential
	// backoff and jitter.
	BaseRetryWait *duration.Duration `protobuf:"bytes,3,opt,name=base_retry_wait,json=baseRetryWait,proto3" json:"base_retry_wait,omitempty"`
	// Max time to wait between retries.
	MaxRetryWait         *duration.Duration `protobuf:"bytes,4,opt,name=max_retry_wait,json=maxRetryWait,proto3" json:"max_retry_wait,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NetworkFailPolicy) Reset()         { *m = NetworkFailPolicy{} }
func (m *NetworkFailPolicy) String() string { return proto.CompactTextString(m) }
func (*NetworkFailPolicy) ProtoMessage()    {}
func (*NetworkFailPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_27bf0dec365e2f6f, []int{0}
}

func (m *NetworkFailPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkFailPolicy.Unmarshal(m, b)
}
func (m *NetworkFailPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkFailPolicy.Marshal(b, m, deterministic)
}
func (m *NetworkFailPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkFailPolicy.Merge(m, src)
}
func (m *NetworkFailPolicy) XXX_Size() int {
	return xxx_messageInfo_NetworkFailPolicy.Size(m)
}
func (m *NetworkFailPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkFailPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkFailPolicy proto.InternalMessageInfo

func (m *NetworkFailPolicy) GetPolicy() NetworkFailPolicy_FailPolicy {
	if m != nil {
		return m.Policy
	}
	return NetworkFailPolicy_FAIL_OPEN
}

func (m *NetworkFailPolicy) GetMaxRetry() uint32 {
	if m != nil {
		return m.MaxRetry
	}
	return 0
}

func (m *NetworkFailPolicy) GetBaseRetryWait() *duration.Duration {
	if m != nil {
		return m.BaseRetryWait
	}
	return nil
}

func (m *NetworkFailPolicy) GetMaxRetryWait() *duration.Duration {
	if m != nil {
		return m.MaxRetryWait
	}
	return nil
}

// Defines the per-service client configuration.
type ServiceConfig struct {
	// If true, do not call Mixer Check.
	DisableCheckCalls bool `protobuf:"varint,1,opt,name=disable_check_calls,json=disableCheckCalls,proto3" json:"disable_check_calls,omitempty"`
	// If true, do not call Mixer Report.
	DisableReportCalls bool `protobuf:"varint,2,opt,name=disable_report_calls,json=disableReportCalls,proto3" json:"disable_report_calls,omitempty"`
	// Send these attributes to Mixer in both Check and Report. This
	// typically includes the "destination.service" attribute.
	// In case of a per-route override, per-route attributes take precedence
	// over the attributes supplied in the client configuration.
	MixerAttributes *Attributes `protobuf:"bytes,3,opt,name=mixer_attributes,json=mixerAttributes,proto3" json:"mixer_attributes,omitempty"`
	// HTTP API specifications to generate API attributes.
	HttpApiSpec []*HTTPAPISpec `protobuf:"bytes,4,rep,name=http_api_spec,json=httpApiSpec,proto3" json:"http_api_spec,omitempty"`
	// Quota specifications to generate quota requirements.
	QuotaSpec []*QuotaSpec `protobuf:"bytes,5,rep,name=quota_spec,json=quotaSpec,proto3" json:"quota_spec,omitempty"`
	// Specifies the behavior when the client is unable to connect to Mixer.
	// This is the service-level policy. It overrides
	// [mesh-level
	// policy][istio.mixer.v1.config.client.TransportConfig.network_fail_policy].
	NetworkFailPolicy *NetworkFailPolicy `protobuf:"bytes,7,opt,name=network_fail_policy,json=networkFailPolicy,proto3" json:"network_fail_policy,omitempty"`
	// Default attributes to forward to upstream. This typically
	// includes the "source.ip" and "source.uid" attributes.
	// In case of a per-route override, per-route attributes take precedence
	// over the attributes supplied in the client configuration.
	//
	// Forwarded attributes take precedence over the static Mixer attributes,
	// except in cases where there is clear configuration to ignore forwarded
	// attributes. Gateways, for instance, should never use forwarded attributes.
	//
	// The full order of application is as follows:
	// 1. static Mixer attributes from the filter config;
	// 2. static Mixer attributes from the route config;
	// 3. forwarded attributes from the source filter config (if any and not ignored);
	// 4. forwarded attributes from the source route config (if any and not ignored);
	// 5. derived attributes from the request metadata.
	ForwardAttributes    *Attributes `protobuf:"bytes,8,opt,name=forward_attributes,json=forwardAttributes,proto3" json:"forward_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ServiceConfig) Reset()         { *m = ServiceConfig{} }
func (m *ServiceConfig) String() string { return proto.CompactTextString(m) }
func (*ServiceConfig) ProtoMessage()    {}
func (*ServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_27bf0dec365e2f6f, []int{1}
}

func (m *ServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceConfig.Unmarshal(m, b)
}
func (m *ServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceConfig.Marshal(b, m, deterministic)
}
func (m *ServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceConfig.Merge(m, src)
}
func (m *ServiceConfig) XXX_Size() int {
	return xxx_messageInfo_ServiceConfig.Size(m)
}
func (m *ServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceConfig proto.InternalMessageInfo

func (m *ServiceConfig) GetDisableCheckCalls() bool {
	if m != nil {
		return m.DisableCheckCalls
	}
	return false
}

func (m *ServiceConfig) GetDisableReportCalls() bool {
	if m != nil {
		return m.DisableReportCalls
	}
	return false
}

func (m *ServiceConfig) GetMixerAttributes() *Attributes {
	if m != nil {
		return m.MixerAttributes
	}
	return nil
}

func (m *ServiceConfig) GetHttpApiSpec() []*HTTPAPISpec {
	if m != nil {
		return m.HttpApiSpec
	}
	return nil
}

func (m *ServiceConfig) GetQuotaSpec() []*QuotaSpec {
	if m != nil {
		return m.QuotaSpec
	}
	return nil
}

func (m *ServiceConfig) GetNetworkFailPolicy() *NetworkFailPolicy {
	if m != nil {
		return m.NetworkFailPolicy
	}
	return nil
}

func (m *ServiceConfig) GetForwardAttributes() *Attributes {
	if m != nil {
		return m.ForwardAttributes
	}
	return nil
}

// Defines the transport config on how to call Mixer.
type TransportConfig struct {
	// The flag to disable check cache.
	DisableCheckCache bool `protobuf:"varint,1,opt,name=disable_check_cache,json=disableCheckCache,proto3" json:"disable_check_cache,omitempty"`
	// The flag to disable quota cache.
	DisableQuotaCache bool `protobuf:"varint,2,opt,name=disable_quota_cache,json=disableQuotaCache,proto3" json:"disable_quota_cache,omitempty"`
	// The flag to disable report batch.
	DisableReportBatch bool `protobuf:"varint,3,opt,name=disable_report_batch,json=disableReportBatch,proto3" json:"disable_report_batch,omitempty"`
	// Specifies the behavior when the client is unable to connect to Mixer.
	// This is the mesh level policy. The default value for policy is FAIL_OPEN.
	NetworkFailPolicy *NetworkFailPolicy `protobuf:"bytes,4,opt,name=network_fail_policy,json=networkFailPolicy,proto3" json:"network_fail_policy,omitempty"`
	// Specify refresh interval to write Mixer client statistics to Envoy share
	// memory. If not specified, the interval is 10 seconds.
	StatsUpdateInterval *duration.Duration `protobuf:"bytes,5,opt,name=stats_update_interval,json=statsUpdateInterval,proto3" json:"stats_update_interval,omitempty"`
	// Name of the cluster that will forward check calls to a pool of mixer
	// servers. Defaults to "mixer_server". By using different names for
	// checkCluster and reportCluster, it is possible to have one set of
	// Mixer servers handle check calls, while another set of Mixer servers
	// handle report calls.
	//
	// NOTE: Any value other than the default "mixer_server" will require the
	// Istio Grafana dashboards to be reconfigured to use the new name.
	CheckCluster string `protobuf:"bytes,6,opt,name=check_cluster,json=checkCluster,proto3" json:"check_cluster,omitempty"`
	// Name of the cluster that will forward report calls to a pool of mixer
	// servers. Defaults to "mixer_server". By using different names for
	// checkCluster and reportCluster, it is possible to have one set of
	// Mixer servers handle check calls, while another set of Mixer servers
	// handle report calls.
	//
	// NOTE: Any value other than the default "mixer_server" will require the
	// Istio Grafana dashboards to be reconfigured to use the new name.
	ReportCluster string `protobuf:"bytes,7,opt,name=report_cluster,json=reportCluster,proto3" json:"report_cluster,omitempty"`
	// Default attributes to forward to Mixer upstream. This typically
	// includes the "source.ip" and "source.uid" attributes. These
	// attributes are consumed by the proxy in front of mixer.
	AttributesForMixerProxy *Attributes `protobuf:"bytes,8,opt,name=attributes_for_mixer_proxy,json=attributesForMixerProxy,proto3" json:"attributes_for_mixer_proxy,omitempty"`
	// When disable_report_batch is false, this value specifies the maximum number
	// of requests that are batched in report. If left unspecified, the default value
	// of report_batch_max_entries == 0 will use the hardcoded defaults of
	// istio::mixerclient::ReportOptions.
	ReportBatchMaxEntries uint32 `protobuf:"varint,9,opt,name=report_batch_max_entries,json=reportBatchMaxEntries,proto3" json:"report_batch_max_entries,omitempty"`
	// When disable_report_batch is false, this value specifies the maximum elapsed
	// time a batched report will be sent after a user request is processed. If left
	// unspecified, the default report_batch_max_time == 0 will use the hardcoded
	// defaults of istio::mixerclient::ReportOptions.
	ReportBatchMaxTime   *duration.Duration `protobuf:"bytes,10,opt,name=report_batch_max_time,json=reportBatchMaxTime,proto3" json:"report_batch_max_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TransportConfig) Reset()         { *m = TransportConfig{} }
func (m *TransportConfig) String() string { return proto.CompactTextString(m) }
func (*TransportConfig) ProtoMessage()    {}
func (*TransportConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_27bf0dec365e2f6f, []int{2}
}

func (m *TransportConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransportConfig.Unmarshal(m, b)
}
func (m *TransportConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransportConfig.Marshal(b, m, deterministic)
}
func (m *TransportConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransportConfig.Merge(m, src)
}
func (m *TransportConfig) XXX_Size() int {
	return xxx_messageInfo_TransportConfig.Size(m)
}
func (m *TransportConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TransportConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TransportConfig proto.InternalMessageInfo

func (m *TransportConfig) GetDisableCheckCache() bool {
	if m != nil {
		return m.DisableCheckCache
	}
	return false
}

func (m *TransportConfig) GetDisableQuotaCache() bool {
	if m != nil {
		return m.DisableQuotaCache
	}
	return false
}

func (m *TransportConfig) GetDisableReportBatch() bool {
	if m != nil {
		return m.DisableReportBatch
	}
	return false
}

func (m *TransportConfig) GetNetworkFailPolicy() *NetworkFailPolicy {
	if m != nil {
		return m.NetworkFailPolicy
	}
	return nil
}

func (m *TransportConfig) GetStatsUpdateInterval() *duration.Duration {
	if m != nil {
		return m.StatsUpdateInterval
	}
	return nil
}

func (m *TransportConfig) GetCheckCluster() string {
	if m != nil {
		return m.CheckCluster
	}
	return ""
}

func (m *TransportConfig) GetReportCluster() string {
	if m != nil {
		return m.ReportCluster
	}
	return ""
}

func (m *TransportConfig) GetAttributesForMixerProxy() *Attributes {
	if m != nil {
		return m.AttributesForMixerProxy
	}
	return nil
}

func (m *TransportConfig) GetReportBatchMaxEntries() uint32 {
	if m != nil {
		return m.ReportBatchMaxEntries
	}
	return 0
}

func (m *TransportConfig) GetReportBatchMaxTime() *duration.Duration {
	if m != nil {
		return m.ReportBatchMaxTime
	}
	return nil
}

// Defines the client config for HTTP.
type HttpClientConfig struct {
	// The transport config.
	Transport *TransportConfig `protobuf:"bytes,1,opt,name=transport,proto3" json:"transport,omitempty"`
	// Map of control configuration indexed by destination.service. This
	// is used to support per-service configuration for cases where a
	// mixerclient serves multiple services.
	ServiceConfigs map[string]*ServiceConfig `protobuf:"bytes,2,rep,name=service_configs,json=serviceConfigs,proto3" json:"service_configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Default destination service name if none was specified in the
	// client request.
	DefaultDestinationService string `protobuf:"bytes,3,opt,name=default_destination_service,json=defaultDestinationService,proto3" json:"default_destination_service,omitempty"`
	// Default attributes to send to Mixer in both Check and
	// Report. This typically includes "destination.ip" and
	// "destination.uid" attributes.
	MixerAttributes *Attributes `protobuf:"bytes,4,opt,name=mixer_attributes,json=mixerAttributes,proto3" json:"mixer_attributes,omitempty"`
	// Default attributes to forward to upstream. This typically
	// includes the "source.ip" and "source.uid" attributes.
	ForwardAttributes *Attributes `protobuf:"bytes,5,opt,name=forward_attributes,json=forwardAttributes,proto3" json:"forward_attributes,omitempty"`
	// Whether or not to use attributes forwarded in the request headers to
	// create the attribute bag to send to mixer. For intra-mesh traffic,
	// this should be set to "false". For ingress/egress gateways, this
	// should be set to "true".
	IgnoreForwardedAttributes bool     `protobuf:"varint,6,opt,name=ignore_forwarded_attributes,json=ignoreForwardedAttributes,proto3" json:"ignore_forwarded_attributes,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *HttpClientConfig) Reset()         { *m = HttpClientConfig{} }
func (m *HttpClientConfig) String() string { return proto.CompactTextString(m) }
func (*HttpClientConfig) ProtoMessage()    {}
func (*HttpClientConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_27bf0dec365e2f6f, []int{3}
}

func (m *HttpClientConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpClientConfig.Unmarshal(m, b)
}
func (m *HttpClientConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpClientConfig.Marshal(b, m, deterministic)
}
func (m *HttpClientConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpClientConfig.Merge(m, src)
}
func (m *HttpClientConfig) XXX_Size() int {
	return xxx_messageInfo_HttpClientConfig.Size(m)
}
func (m *HttpClientConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpClientConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HttpClientConfig proto.InternalMessageInfo

func (m *HttpClientConfig) GetTransport() *TransportConfig {
	if m != nil {
		return m.Transport
	}
	return nil
}

func (m *HttpClientConfig) GetServiceConfigs() map[string]*ServiceConfig {
	if m != nil {
		return m.ServiceConfigs
	}
	return nil
}

func (m *HttpClientConfig) GetDefaultDestinationService() string {
	if m != nil {
		return m.DefaultDestinationService
	}
	return ""
}

func (m *HttpClientConfig) GetMixerAttributes() *Attributes {
	if m != nil {
		return m.MixerAttributes
	}
	return nil
}

func (m *HttpClientConfig) GetForwardAttributes() *Attributes {
	if m != nil {
		return m.ForwardAttributes
	}
	return nil
}

func (m *HttpClientConfig) GetIgnoreForwardedAttributes() bool {
	if m != nil {
		return m.IgnoreForwardedAttributes
	}
	return false
}

// Defines the client config for TCP.
type TcpClientConfig struct {
	// The transport config.
	Transport *TransportConfig `protobuf:"bytes,1,opt,name=transport,proto3" json:"transport,omitempty"`
	// Default attributes to send to Mixer in both Check and
	// Report. This typically includes "destination.ip" and
	// "destination.uid" attributes.
	MixerAttributes *Attributes `protobuf:"bytes,2,opt,name=mixer_attributes,json=mixerAttributes,proto3" json:"mixer_attributes,omitempty"`
	// If set to true, disables Mixer check calls.
	DisableCheckCalls bool `protobuf:"varint,3,opt,name=disable_check_calls,json=disableCheckCalls,proto3" json:"disable_check_calls,omitempty"`
	// If set to true, disables Mixer check calls.
	DisableReportCalls bool `protobuf:"varint,4,opt,name=disable_report_calls,json=disableReportCalls,proto3" json:"disable_report_calls,omitempty"`
	// Quota specifications to generate quota requirements.
	// It applies on the new TCP connections.
	ConnectionQuotaSpec *QuotaSpec `protobuf:"bytes,5,opt,name=connection_quota_spec,json=connectionQuotaSpec,proto3" json:"connection_quota_spec,omitempty"`
	// Specify report interval to send periodical reports for long TCP
	// connections. If not specified, the interval is 10 seconds. This interval
	// should not be less than 1 second, otherwise it will be reset to 1 second.
	ReportInterval       *duration.Duration `protobuf:"bytes,6,opt,name=report_interval,json=reportInterval,proto3" json:"report_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TcpClientConfig) Reset()         { *m = TcpClientConfig{} }
func (m *TcpClientConfig) String() string { return proto.CompactTextString(m) }
func (*TcpClientConfig) ProtoMessage()    {}
func (*TcpClientConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_27bf0dec365e2f6f, []int{4}
}

func (m *TcpClientConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpClientConfig.Unmarshal(m, b)
}
func (m *TcpClientConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpClientConfig.Marshal(b, m, deterministic)
}
func (m *TcpClientConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpClientConfig.Merge(m, src)
}
func (m *TcpClientConfig) XXX_Size() int {
	return xxx_messageInfo_TcpClientConfig.Size(m)
}
func (m *TcpClientConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpClientConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TcpClientConfig proto.InternalMessageInfo

func (m *TcpClientConfig) GetTransport() *TransportConfig {
	if m != nil {
		return m.Transport
	}
	return nil
}

func (m *TcpClientConfig) GetMixerAttributes() *Attributes {
	if m != nil {
		return m.MixerAttributes
	}
	return nil
}

func (m *TcpClientConfig) GetDisableCheckCalls() bool {
	if m != nil {
		return m.DisableCheckCalls
	}
	return false
}

func (m *TcpClientConfig) GetDisableReportCalls() bool {
	if m != nil {
		return m.DisableReportCalls
	}
	return false
}

func (m *TcpClientConfig) GetConnectionQuotaSpec() *QuotaSpec {
	if m != nil {
		return m.ConnectionQuotaSpec
	}
	return nil
}

func (m *TcpClientConfig) GetReportInterval() *duration.Duration {
	if m != nil {
		return m.ReportInterval
	}
	return nil
}

func init() {
	proto.RegisterEnum("istio.mixer.v1.config.client.NetworkFailPolicy_FailPolicy", NetworkFailPolicy_FailPolicy_name, NetworkFailPolicy_FailPolicy_value)
	proto.RegisterType((*NetworkFailPolicy)(nil), "istio.mixer.v1.config.client.NetworkFailPolicy")
	proto.RegisterType((*ServiceConfig)(nil), "istio.mixer.v1.config.client.ServiceConfig")
	proto.RegisterType((*TransportConfig)(nil), "istio.mixer.v1.config.client.TransportConfig")
	proto.RegisterType((*HttpClientConfig)(nil), "istio.mixer.v1.config.client.HttpClientConfig")
	proto.RegisterMapType((map[string]*ServiceConfig)(nil), "istio.mixer.v1.config.client.HttpClientConfig.ServiceConfigsEntry")
	proto.RegisterType((*TcpClientConfig)(nil), "istio.mixer.v1.config.client.TcpClientConfig")
}

func init() {
	proto.RegisterFile("mixer/v1/config/client/client_config.proto", fileDescriptor_27bf0dec365e2f6f)
}

var fileDescriptor_27bf0dec365e2f6f = []byte{
	// 973 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0x1a, 0x47,
	0x14, 0x36, 0x3f, 0x26, 0x61, 0x08, 0x3f, 0x1e, 0x62, 0x15, 0x93, 0x0a, 0x21, 0xa2, 0xa8, 0xb4,
	0x51, 0x97, 0x86, 0xaa, 0x6a, 0x95, 0x8b, 0x56, 0x98, 0x18, 0xc5, 0xaa, 0x9d, 0xd0, 0x35, 0x55,
	0xa4, 0xf6, 0x62, 0x34, 0x2c, 0x03, 0x8c, 0xbc, 0xec, 0x6c, 0x66, 0x07, 0x8c, 0xdf, 0x28, 0x2f,
	0xd0, 0x77, 0xe8, 0x55, 0xaf, 0x7b, 0x9b, 0x3c, 0x41, 0xef, 0x7b, 0xd1, 0x6a, 0x7e, 0x96, 0x35,
	0x36, 0x3f, 0xa5, 0x6a, 0xaf, 0x98, 0x3d, 0xe7, 0xfb, 0x0e, 0x3b, 0xdf, 0xf9, 0xe6, 0xcc, 0x82,
	0xcf, 0x26, 0x74, 0x4e, 0x78, 0x63, 0xf6, 0xac, 0xe1, 0x30, 0x6f, 0x48, 0x47, 0x0d, 0xc7, 0xa5,
	0xc4, 0x13, 0xe6, 0x07, 0xe9, 0xa0, 0xe5, 0x73, 0x26, 0x18, 0xfc, 0x98, 0x06, 0x82, 0x32, 0x4b,
	0x31, 0xac, 0xd9, 0x33, 0xcb, 0x24, 0x35, 0xb4, 0xfc, 0x70, 0xc4, 0x46, 0x4c, 0x01, 0x1b, 0x72,
	0xa5, 0x39, 0xe5, 0xca, 0x88, 0xb1, 0x91, 0x4b, 0x1a, 0xea, 0xa9, 0x3f, 0x1d, 0x36, 0x06, 0x53,
	0x8e, 0x05, 0x65, 0x9e, 0xc9, 0x1f, 0x2d, 0xfe, 0x1f, 0x0b, 0xc1, 0x69, 0x7f, 0x2a, 0x48, 0x60,
	0x52, 0x4f, 0xd6, 0xbc, 0x1a, 0xf6, 0x29, 0x0a, 0x7c, 0xe2, 0x18, 0x58, 0x6d, 0x0d, 0xec, 0xed,
	0x94, 0x09, 0xac, 0x31, 0xb5, 0x5f, 0xe2, 0xe0, 0xe0, 0x15, 0x11, 0x57, 0x8c, 0x5f, 0x76, 0x30,
	0x75, 0xbb, 0xcc, 0xa5, 0xce, 0x35, 0xb4, 0x41, 0xca, 0x57, 0xab, 0x52, 0xac, 0x1a, 0xab, 0xe7,
	0x9a, 0xcf, 0xad, 0x4d, 0x1b, 0xb4, 0xee, 0x14, 0xb0, 0xa2, 0xa5, 0x6d, 0x2a, 0xc1, 0x47, 0x20,
	0x3d, 0xc1, 0x73, 0xc4, 0x89, 0xe0, 0xd7, 0xa5, 0x78, 0x35, 0x56, 0xcf, 0xda, 0xf7, 0x27, 0x78,
	0x6e, 0xcb, 0x67, 0xd8, 0x02, 0xf9, 0x3e, 0x0e, 0x88, 0xce, 0xa2, 0x2b, 0x4c, 0x45, 0x29, 0x51,
	0x8d, 0xd5, 0x33, 0xcd, 0x23, 0x4b, 0xcb, 0x64, 0x85, 0x32, 0x59, 0x2f, 0x8c, 0x4c, 0x76, 0x56,
	0x32, 0x14, 0xfd, 0x0d, 0xa6, 0x02, 0x7e, 0x07, 0x72, 0x8b, 0xfa, 0xba, 0x42, 0x72, 0x5b, 0x85,
	0x07, 0xe1, 0xff, 0xcb, 0x02, 0xb5, 0xa7, 0x00, 0xdc, 0x90, 0x20, 0x0b, 0xd2, 0x9d, 0xd6, 0xe9,
	0x19, 0x7a, 0xdd, 0x3d, 0x79, 0x55, 0xd8, 0x83, 0x39, 0x00, 0xd4, 0x63, 0xfb, 0xec, 0xf5, 0xc5,
	0x49, 0x21, 0x56, 0xfb, 0x33, 0x01, 0xb2, 0x17, 0x84, 0xcf, 0xa8, 0x43, 0xda, 0x4a, 0x0b, 0x68,
	0x81, 0xe2, 0x80, 0x06, 0xb8, 0xef, 0x12, 0xe4, 0x8c, 0x89, 0x73, 0x89, 0x1c, 0xec, 0xba, 0x81,
	0x12, 0xf0, 0xbe, 0x7d, 0x60, 0x52, 0x6d, 0x99, 0x69, 0xcb, 0x04, 0xfc, 0x02, 0x3c, 0x0c, 0xf1,
	0x9c, 0xf8, 0x8c, 0x0b, 0x43, 0x88, 0x2b, 0x02, 0x34, 0x39, 0x5b, 0xa5, 0x34, 0xe3, 0x04, 0x14,
	0x54, 0x03, 0x50, 0x64, 0x08, 0xa3, 0x52, 0xf9, 0x76, 0x7f, 0x5a, 0x0b, 0x84, 0x9d, 0x57, 0xc1,
	0x28, 0x00, 0xcf, 0x41, 0x76, 0x2c, 0x84, 0x8f, 0x42, 0xb7, 0x94, 0x92, 0xd5, 0x44, 0x3d, 0xd3,
	0xfc, 0x74, 0x73, 0x8f, 0x5f, 0xf6, 0x7a, 0xdd, 0x56, 0xf7, 0xf4, 0xc2, 0x27, 0x8e, 0x9d, 0x91,
	0xfc, 0x96, 0x4f, 0xe5, 0x03, 0xec, 0x00, 0xa0, 0x0c, 0xa5, 0x6b, 0xed, 0xab, 0x5a, 0x9f, 0x6c,
	0xae, 0xf5, 0x83, 0xc4, 0xab, 0x4a, 0xe9, 0xb7, 0xe1, 0x12, 0x22, 0x50, 0xf4, 0xb4, 0x8f, 0xd0,
	0x10, 0x53, 0x17, 0x19, 0x03, 0xde, 0x53, 0x1b, 0x6c, 0xec, 0x68, 0x40, 0xfb, 0xc0, 0xbb, 0x63,
	0xea, 0x53, 0x00, 0x87, 0x8c, 0x5f, 0x61, 0x3e, 0xb8, 0x29, 0xe0, 0xfd, 0xad, 0x02, 0x1e, 0x18,
	0x56, 0x14, 0xaa, 0xfd, 0x95, 0x04, 0xf9, 0x1e, 0xc7, 0x5e, 0xa0, 0x9a, 0xb3, 0xb6, 0xff, 0xce,
	0x98, 0xac, 0xee, 0xbf, 0x33, 0x26, 0x37, 0xf1, 0x5a, 0x3f, 0x8d, 0x8f, 0x2f, 0xe1, 0x95, 0x52,
	0x1a, 0x7f, 0xd7, 0x2f, 0x7d, 0x2c, 0x9c, 0xb1, 0x72, 0xc0, 0x6d, 0xbf, 0x1c, 0xcb, 0xcc, 0x3a,
	0x45, 0x93, 0xff, 0x99, 0xa2, 0xe7, 0xe0, 0x30, 0x10, 0x58, 0x04, 0x68, 0xea, 0x0f, 0xb0, 0x20,
	0x88, 0x7a, 0x82, 0xf0, 0x19, 0x76, 0x4b, 0xfb, 0xdb, 0x4e, 0x5e, 0x51, 0xf1, 0x7e, 0x54, 0xb4,
	0x53, 0xc3, 0x82, 0x8f, 0x41, 0xd6, 0x28, 0xe7, 0x4e, 0x03, 0x41, 0x78, 0x29, 0x55, 0x8d, 0xd5,
	0xd3, 0xf6, 0x03, 0x15, 0x6c, 0xeb, 0x18, 0x7c, 0x02, 0x72, 0xe1, 0x71, 0x31, 0xa8, 0x7b, 0x0a,
	0x95, 0xd5, 0xd1, 0x10, 0xf6, 0x06, 0x94, 0xa3, 0x26, 0xa3, 0x21, 0xe3, 0x48, 0x1f, 0x1d, 0x9f,
	0xb3, 0xf9, 0xf5, 0x3f, 0x68, 0xfa, 0x47, 0x11, 0xbb, 0xc3, 0xf8, 0xb9, 0x44, 0x74, 0x25, 0x15,
	0x7e, 0x0d, 0x4a, 0x37, 0xe5, 0x47, 0x72, 0xe6, 0x10, 0x4f, 0x70, 0x4a, 0x82, 0x52, 0x5a, 0x4d,
	0xb5, 0x43, 0x1e, 0xf5, 0xe0, 0x1c, 0xcf, 0x4f, 0x74, 0x12, 0x9e, 0x81, 0xc3, 0x3b, 0x44, 0x41,
	0x27, 0xa4, 0x04, 0xb6, 0x89, 0x05, 0x97, 0x0b, 0xf6, 0xe8, 0x84, 0xd4, 0x7e, 0x4b, 0x82, 0xc2,
	0x4b, 0x21, 0xfc, 0xb6, 0x6a, 0x97, 0xb1, 0xe0, 0xf7, 0x20, 0x2d, 0x42, 0x57, 0x2a, 0xe3, 0x65,
	0x9a, 0x9f, 0x6f, 0x6e, 0xf3, 0x2d, 0x13, 0xdb, 0x11, 0x1f, 0x5e, 0x82, 0x7c, 0xa0, 0x07, 0x9c,
	0xb9, 0xeb, 0xe4, 0x68, 0x92, 0x87, 0xfb, 0x78, 0xcb, 0xa0, 0xb8, 0xf5, 0x56, 0xd6, 0xd2, 0x98,
	0x0c, 0xa4, 0x1c, 0xd7, 0x76, 0x2e, 0x58, 0x0a, 0xc2, 0x6f, 0xc1, 0xa3, 0x01, 0x19, 0xe2, 0xa9,
	0x2b, 0xd0, 0x80, 0x04, 0x82, 0x7a, 0x6a, 0xe7, 0xc8, 0xa0, 0x94, 0xc7, 0xd3, 0xf6, 0x91, 0x81,
	0xbc, 0x88, 0x10, 0xa6, 0xf6, 0xca, 0xd1, 0x98, 0xdc, 0x7d, 0x34, 0xae, 0x1e, 0x11, 0xfb, 0xff,
	0x62, 0x44, 0xc8, 0x1d, 0xd1, 0x91, 0xc7, 0x38, 0x41, 0x26, 0x47, 0x96, 0x6a, 0xa6, 0xd4, 0xa9,
	0x3d, 0xd2, 0x90, 0x4e, 0x88, 0x88, 0xf8, 0x65, 0x0f, 0x14, 0x57, 0x08, 0x07, 0x0b, 0x20, 0x71,
	0x49, 0xf4, 0xb5, 0x9c, 0xb6, 0xe5, 0x12, 0xb6, 0xc0, 0xfe, 0x0c, 0xbb, 0x53, 0x3d, 0x39, 0x32,
	0xcd, 0xa7, 0x9b, 0xbb, 0xb3, 0x54, 0xd3, 0xd6, 0xcc, 0xe7, 0xf1, 0x6f, 0x62, 0xb5, 0x77, 0x09,
	0x90, 0xef, 0x39, 0xff, 0xa3, 0x9f, 0x56, 0xb5, 0x28, 0xbe, 0x7b, 0x8b, 0xd6, 0x5c, 0xb3, 0x89,
	0x5d, 0xaf, 0xd9, 0xe4, 0xda, 0x6b, 0xf6, 0x67, 0x70, 0xe8, 0x30, 0xcf, 0x23, 0x8e, 0xb2, 0xe0,
	0xd2, 0xdd, 0x16, 0xdb, 0xe5, 0x6e, 0x2b, 0x46, 0x55, 0x16, 0x41, 0x78, 0x0c, 0xf2, 0xe6, 0x35,
	0x16, 0xc3, 0x32, 0xb5, 0xed, 0xfc, 0x9b, 0x81, 0x17, 0xce, 0xc9, 0xe3, 0xaf, 0x7e, 0x7d, 0x5f,
	0xd9, 0xfb, 0xe3, 0x7d, 0x65, 0xef, 0xdd, 0x87, 0xca, 0xde, 0xef, 0x1f, 0x2a, 0xb1, 0x9f, 0x1e,
	0xeb, 0x57, 0xa2, 0x4c, 0x7e, 0x02, 0x36, 0x56, 0x7f, 0xf6, 0xf5, 0x53, 0xaa, 0xf2, 0x97, 0x7f,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x78, 0xc0, 0xb5, 0x46, 0xd9, 0x0a, 0x00, 0x00,
}
