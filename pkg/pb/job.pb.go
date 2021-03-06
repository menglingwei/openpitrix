// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

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

type CreateJobRequest struct {
	ClusterId            *wrappers.StringValue `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	AppId                *wrappers.StringValue `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	VersionId            *wrappers.StringValue `protobuf:"bytes,3,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	JobAction            *wrappers.StringValue `protobuf:"bytes,4,opt,name=job_action,json=jobAction,proto3" json:"job_action,omitempty"`
	Provider             *wrappers.StringValue `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
	Directive            *wrappers.StringValue `protobuf:"bytes,6,opt,name=directive,proto3" json:"directive,omitempty"`
	RuntimeId            *wrappers.StringValue `protobuf:"bytes,7,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateJobRequest) Reset()         { *m = CreateJobRequest{} }
func (m *CreateJobRequest) String() string { return proto.CompactTextString(m) }
func (*CreateJobRequest) ProtoMessage()    {}
func (*CreateJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{0}
}

func (m *CreateJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJobRequest.Unmarshal(m, b)
}
func (m *CreateJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJobRequest.Marshal(b, m, deterministic)
}
func (m *CreateJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJobRequest.Merge(m, src)
}
func (m *CreateJobRequest) XXX_Size() int {
	return xxx_messageInfo_CreateJobRequest.Size(m)
}
func (m *CreateJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJobRequest proto.InternalMessageInfo

func (m *CreateJobRequest) GetClusterId() *wrappers.StringValue {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *CreateJobRequest) GetAppId() *wrappers.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *CreateJobRequest) GetVersionId() *wrappers.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *CreateJobRequest) GetJobAction() *wrappers.StringValue {
	if m != nil {
		return m.JobAction
	}
	return nil
}

func (m *CreateJobRequest) GetProvider() *wrappers.StringValue {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *CreateJobRequest) GetDirective() *wrappers.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *CreateJobRequest) GetRuntimeId() *wrappers.StringValue {
	if m != nil {
		return m.RuntimeId
	}
	return nil
}

type CreateJobResponse struct {
	JobId                *wrappers.StringValue `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	ClusterId            *wrappers.StringValue `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	AppId                *wrappers.StringValue `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	VersionId            *wrappers.StringValue `protobuf:"bytes,4,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	RuntimeId            *wrappers.StringValue `protobuf:"bytes,5,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateJobResponse) Reset()         { *m = CreateJobResponse{} }
func (m *CreateJobResponse) String() string { return proto.CompactTextString(m) }
func (*CreateJobResponse) ProtoMessage()    {}
func (*CreateJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{1}
}

func (m *CreateJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJobResponse.Unmarshal(m, b)
}
func (m *CreateJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJobResponse.Marshal(b, m, deterministic)
}
func (m *CreateJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJobResponse.Merge(m, src)
}
func (m *CreateJobResponse) XXX_Size() int {
	return xxx_messageInfo_CreateJobResponse.Size(m)
}
func (m *CreateJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJobResponse proto.InternalMessageInfo

func (m *CreateJobResponse) GetJobId() *wrappers.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *CreateJobResponse) GetClusterId() *wrappers.StringValue {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *CreateJobResponse) GetAppId() *wrappers.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *CreateJobResponse) GetVersionId() *wrappers.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *CreateJobResponse) GetRuntimeId() *wrappers.StringValue {
	if m != nil {
		return m.RuntimeId
	}
	return nil
}

type Job struct {
	JobId                *wrappers.StringValue `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	ClusterId            *wrappers.StringValue `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	AppId                *wrappers.StringValue `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	VersionId            *wrappers.StringValue `protobuf:"bytes,4,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	JobAction            *wrappers.StringValue `protobuf:"bytes,5,opt,name=job_action,json=jobAction,proto3" json:"job_action,omitempty"`
	Status               *wrappers.StringValue `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	ErrorCode            *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	Directive            *wrappers.StringValue `protobuf:"bytes,8,opt,name=directive,proto3" json:"directive,omitempty"`
	Executor             *wrappers.StringValue `protobuf:"bytes,9,opt,name=executor,proto3" json:"executor,omitempty"`
	TaskCount            *wrappers.UInt32Value `protobuf:"bytes,10,opt,name=task_count,json=taskCount,proto3" json:"task_count,omitempty"`
	Owner                *wrappers.StringValue `protobuf:"bytes,11,opt,name=owner,proto3" json:"owner,omitempty"`
	Provider             *wrappers.StringValue `protobuf:"bytes,12,opt,name=provider,proto3" json:"provider,omitempty"`
	RuntimeId            *wrappers.StringValue `protobuf:"bytes,13,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	CreateTime           *timestamp.Timestamp  `protobuf:"bytes,14,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	StatusTime           *timestamp.Timestamp  `protobuf:"bytes,15,opt,name=status_time,json=statusTime,proto3" json:"status_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{2}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetJobId() *wrappers.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *Job) GetClusterId() *wrappers.StringValue {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *Job) GetAppId() *wrappers.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *Job) GetVersionId() *wrappers.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *Job) GetJobAction() *wrappers.StringValue {
	if m != nil {
		return m.JobAction
	}
	return nil
}

func (m *Job) GetStatus() *wrappers.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Job) GetErrorCode() *wrappers.UInt32Value {
	if m != nil {
		return m.ErrorCode
	}
	return nil
}

func (m *Job) GetDirective() *wrappers.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *Job) GetExecutor() *wrappers.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *Job) GetTaskCount() *wrappers.UInt32Value {
	if m != nil {
		return m.TaskCount
	}
	return nil
}

func (m *Job) GetOwner() *wrappers.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Job) GetProvider() *wrappers.StringValue {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *Job) GetRuntimeId() *wrappers.StringValue {
	if m != nil {
		return m.RuntimeId
	}
	return nil
}

func (m *Job) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Job) GetStatusTime() *timestamp.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type DescribeJobsRequest struct {
	JobId     []string              `protobuf:"bytes,1,rep,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	ClusterId *wrappers.StringValue `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	AppId     *wrappers.StringValue `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	VersionId *wrappers.StringValue `protobuf:"bytes,4,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	Executor  *wrappers.StringValue `protobuf:"bytes,5,opt,name=executor,proto3" json:"executor,omitempty"`
	Provider  *wrappers.StringValue `protobuf:"bytes,6,opt,name=provider,proto3" json:"provider,omitempty"`
	RuntimeId *wrappers.StringValue `protobuf:"bytes,7,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	Status    []string              `protobuf:"bytes,8,rep,name=status,proto3" json:"status,omitempty"`
	// default is 20, max value is 200
	Limit uint32 `protobuf:"varint,9,opt,name=limit,proto3" json:"limit,omitempty"`
	// default is 0
	Offset               uint32                `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	SearchWord           *wrappers.StringValue `protobuf:"bytes,11,opt,name=search_word,json=searchWord,proto3" json:"search_word,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DescribeJobsRequest) Reset()         { *m = DescribeJobsRequest{} }
func (m *DescribeJobsRequest) String() string { return proto.CompactTextString(m) }
func (*DescribeJobsRequest) ProtoMessage()    {}
func (*DescribeJobsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{3}
}

func (m *DescribeJobsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeJobsRequest.Unmarshal(m, b)
}
func (m *DescribeJobsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeJobsRequest.Marshal(b, m, deterministic)
}
func (m *DescribeJobsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeJobsRequest.Merge(m, src)
}
func (m *DescribeJobsRequest) XXX_Size() int {
	return xxx_messageInfo_DescribeJobsRequest.Size(m)
}
func (m *DescribeJobsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeJobsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeJobsRequest proto.InternalMessageInfo

func (m *DescribeJobsRequest) GetJobId() []string {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *DescribeJobsRequest) GetClusterId() *wrappers.StringValue {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *DescribeJobsRequest) GetAppId() *wrappers.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *DescribeJobsRequest) GetVersionId() *wrappers.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *DescribeJobsRequest) GetExecutor() *wrappers.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *DescribeJobsRequest) GetProvider() *wrappers.StringValue {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *DescribeJobsRequest) GetRuntimeId() *wrappers.StringValue {
	if m != nil {
		return m.RuntimeId
	}
	return nil
}

func (m *DescribeJobsRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeJobsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeJobsRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeJobsRequest) GetSearchWord() *wrappers.StringValue {
	if m != nil {
		return m.SearchWord
	}
	return nil
}

type DescribeJobsResponse struct {
	TotalCount           uint32   `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	JobSet               []*Job   `protobuf:"bytes,2,rep,name=job_set,json=jobSet,proto3" json:"job_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeJobsResponse) Reset()         { *m = DescribeJobsResponse{} }
func (m *DescribeJobsResponse) String() string { return proto.CompactTextString(m) }
func (*DescribeJobsResponse) ProtoMessage()    {}
func (*DescribeJobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{4}
}

func (m *DescribeJobsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeJobsResponse.Unmarshal(m, b)
}
func (m *DescribeJobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeJobsResponse.Marshal(b, m, deterministic)
}
func (m *DescribeJobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeJobsResponse.Merge(m, src)
}
func (m *DescribeJobsResponse) XXX_Size() int {
	return xxx_messageInfo_DescribeJobsResponse.Size(m)
}
func (m *DescribeJobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeJobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeJobsResponse proto.InternalMessageInfo

func (m *DescribeJobsResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeJobsResponse) GetJobSet() []*Job {
	if m != nil {
		return m.JobSet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateJobRequest)(nil), "openpitrix.CreateJobRequest")
	proto.RegisterType((*CreateJobResponse)(nil), "openpitrix.CreateJobResponse")
	proto.RegisterType((*Job)(nil), "openpitrix.Job")
	proto.RegisterType((*DescribeJobsRequest)(nil), "openpitrix.DescribeJobsRequest")
	proto.RegisterType((*DescribeJobsResponse)(nil), "openpitrix.DescribeJobsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JobManagerClient is the client API for JobManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobManagerClient interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error)
	DescribeJobs(ctx context.Context, in *DescribeJobsRequest, opts ...grpc.CallOption) (*DescribeJobsResponse, error)
}

type jobManagerClient struct {
	cc *grpc.ClientConn
}

func NewJobManagerClient(cc *grpc.ClientConn) JobManagerClient {
	return &jobManagerClient{cc}
}

func (c *jobManagerClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error) {
	out := new(CreateJobResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.JobManager/CreateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobManagerClient) DescribeJobs(ctx context.Context, in *DescribeJobsRequest, opts ...grpc.CallOption) (*DescribeJobsResponse, error) {
	out := new(DescribeJobsResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.JobManager/DescribeJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobManagerServer is the server API for JobManager service.
type JobManagerServer interface {
	CreateJob(context.Context, *CreateJobRequest) (*CreateJobResponse, error)
	DescribeJobs(context.Context, *DescribeJobsRequest) (*DescribeJobsResponse, error)
}

func RegisterJobManagerServer(s *grpc.Server, srv JobManagerServer) {
	s.RegisterService(&_JobManager_serviceDesc, srv)
}

func _JobManager_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobManagerServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.JobManager/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobManagerServer).CreateJob(ctx, req.(*CreateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobManager_DescribeJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobManagerServer).DescribeJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.JobManager/DescribeJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobManagerServer).DescribeJobs(ctx, req.(*DescribeJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.JobManager",
	HandlerType: (*JobManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _JobManager_CreateJob_Handler,
		},
		{
			MethodName: "DescribeJobs",
			Handler:    _JobManager_DescribeJobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}

func init() { proto.RegisterFile("job.proto", fileDescriptor_f32c477d91a04ead) }

var fileDescriptor_f32c477d91a04ead = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x56, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x56, 0x7e, 0x9c, 0x26, 0x27, 0xed, 0xed, 0xbd, 0xbe, 0xbd, 0x57, 0xbe, 0xb9, 0x85, 0x46,
	0x59, 0x65, 0x41, 0x13, 0x91, 0xb2, 0xa8, 0xa8, 0x58, 0x94, 0xb2, 0x20, 0x95, 0xd8, 0xa4, 0xfc,
	0x48, 0x6c, 0xa2, 0xb1, 0x7d, 0x92, 0x4e, 0x48, 0x7d, 0x86, 0x99, 0x71, 0xd2, 0x2d, 0xbc, 0x01,
	0xf0, 0x24, 0x48, 0xbc, 0x09, 0x3b, 0xd6, 0xbc, 0x00, 0x6f, 0x80, 0xc6, 0x76, 0x82, 0xd3, 0x52,
	0x61, 0x97, 0x15, 0x62, 0x15, 0x79, 0xce, 0xf7, 0xe5, 0x1c, 0x9f, 0xf3, 0x7d, 0x73, 0x0c, 0xb5,
	0x09, 0xb9, 0x1d, 0x21, 0x49, 0x93, 0x0d, 0x24, 0x30, 0x10, 0x5c, 0x4b, 0x7e, 0xde, 0xb8, 0x39,
	0x26, 0x1a, 0x4f, 0xb1, 0x1b, 0x45, 0xdc, 0x70, 0xd4, 0x9d, 0x4b, 0x26, 0x04, 0x4a, 0x15, 0x63,
	0x1b, 0x3b, 0x17, 0xe3, 0x9a, 0x9f, 0xa1, 0xd2, 0xec, 0x4c, 0x24, 0x80, 0xed, 0x04, 0xc0, 0x04,
	0xef, 0xb2, 0x20, 0x20, 0xcd, 0x34, 0xa7, 0x60, 0x41, 0xbf, 0x15, 0xfd, 0x78, 0xbb, 0x63, 0x0c,
	0x76, 0xd5, 0x9c, 0x8d, 0xc7, 0x28, 0xbb, 0x24, 0x22, 0xc4, 0x65, 0x74, 0xeb, 0x7d, 0x09, 0xfe,
	0x3c, 0x92, 0xc8, 0x34, 0x1e, 0x93, 0x3b, 0xc0, 0x97, 0x21, 0x2a, 0x6d, 0x1f, 0x00, 0x78, 0xd3,
	0x50, 0x69, 0x94, 0x43, 0xee, 0x3b, 0x85, 0x66, 0xa1, 0x5d, 0xef, 0x6d, 0x77, 0xe2, 0xac, 0x9d,
	0x45, 0x59, 0x9d, 0x13, 0x2d, 0x79, 0x30, 0x7e, 0xca, 0xa6, 0x21, 0x0e, 0x6a, 0x09, 0xbe, 0xef,
	0xdb, 0x7b, 0x50, 0x61, 0x42, 0x18, 0x62, 0x31, 0x03, 0xd1, 0x62, 0x42, 0xf4, 0x7d, 0x93, 0x71,
	0x86, 0x52, 0x71, 0x0a, 0x0c, 0xb1, 0x94, 0x25, 0x63, 0x82, 0x8f, 0xc9, 0x13, 0x72, 0x87, 0xcc,
	0x33, 0x2f, 0xe6, 0x94, 0xb3, 0x90, 0x27, 0xe4, 0x1e, 0x46, 0x70, 0x7b, 0x1f, 0xaa, 0x42, 0xd2,
	0x8c, 0xfb, 0x28, 0x1d, 0x2b, 0x03, 0x75, 0x89, 0xb6, 0xef, 0x42, 0xcd, 0xe7, 0x12, 0x3d, 0xcd,
	0x67, 0xe8, 0x54, 0xb2, 0x64, 0x5d, 0xc2, 0x4d, 0xc9, 0x32, 0x0c, 0xcc, 0x60, 0xcd, 0xfb, 0xae,
	0x65, 0x21, 0x27, 0xf8, 0xbe, 0xdf, 0xfa, 0x50, 0x84, 0xbf, 0x52, 0x33, 0x53, 0x82, 0x02, 0x85,
	0xa6, 0xef, 0xa6, 0x0b, 0x19, 0x07, 0x66, 0x4d, 0xc8, 0x8d, 0x5b, 0x97, 0x9a, 0x74, 0xf1, 0xba,
	0x93, 0x2e, 0x5d, 0x77, 0xd2, 0xe5, 0xdc, 0x93, 0x4e, 0xb5, 0xcd, 0xca, 0xd7, 0xb6, 0x2f, 0x15,
	0x28, 0x1d, 0x93, 0xfb, 0xbb, 0x34, 0x2a, 0x65, 0x09, 0x2b, 0x9f, 0x25, 0xee, 0x40, 0x45, 0x69,
	0xa6, 0x43, 0x95, 0x49, 0xd5, 0x09, 0xd6, 0xa4, 0x44, 0x29, 0x49, 0x0e, 0x3d, 0xf2, 0xf1, 0x4a,
	0x49, 0x3f, 0xe9, 0x07, 0x7a, 0xaf, 0x97, 0xa4, 0x8c, 0xf0, 0x47, 0xe4, 0xe3, 0xaa, 0x97, 0xaa,
	0xf9, 0xbc, 0xb4, 0x0f, 0x55, 0x3c, 0x47, 0x2f, 0xd4, 0x24, 0x9d, 0x5a, 0x16, 0x07, 0x2f, 0xd0,
	0xa6, 0x64, 0xcd, 0xd4, 0x8b, 0xa1, 0x47, 0x61, 0xa0, 0x1d, 0xc8, 0x52, 0xb2, 0xc1, 0x1f, 0x19,
	0xb8, 0xdd, 0x03, 0x8b, 0xe6, 0x01, 0x4a, 0xa7, 0x9e, 0x65, 0xa6, 0x11, 0x74, 0xe5, 0xb2, 0x59,
	0xcf, 0x75, 0xd9, 0xac, 0x2a, 0x7f, 0x23, 0x97, 0xf2, 0xed, 0x03, 0xa8, 0x7b, 0xd1, 0x7d, 0x31,
	0x34, 0x07, 0xce, 0x1f, 0x11, 0xbb, 0x71, 0x89, 0xfd, 0x78, 0xb1, 0x67, 0x06, 0x10, 0xc3, 0xcd,
	0x81, 0x21, 0xc7, 0x13, 0x8e, 0xc9, 0x9b, 0x3f, 0x26, 0xc7, 0x70, 0x73, 0xd0, 0x7a, 0x53, 0x86,
	0xbf, 0x1f, 0xa0, 0xf2, 0x24, 0x77, 0xcd, 0x65, 0xa5, 0x16, 0x1b, 0xe6, 0x9f, 0x94, 0x07, 0x4b,
	0xed, 0xda, 0x2f, 0xea, 0xb2, 0xb4, 0xf2, 0xac, 0x5c, 0xca, 0x4b, 0x0b, 0xa1, 0xf2, 0x13, 0x42,
	0xc8, 0xb7, 0x39, 0xec, 0x7f, 0x97, 0xce, 0xae, 0x46, 0x6d, 0x5f, 0x78, 0x77, 0x0b, 0xac, 0x29,
	0x3f, 0xe3, 0x3a, 0xf2, 0xcf, 0xc6, 0x20, 0x7e, 0x30, 0x68, 0x1a, 0x8d, 0x14, 0xc6, 0xd6, 0xd8,
	0x18, 0x24, 0x4f, 0xf6, 0x3d, 0xa8, 0x2b, 0x64, 0xd2, 0x3b, 0x1d, 0xce, 0x49, 0xfa, 0x99, 0xf4,
	0x0f, 0x31, 0xe1, 0x19, 0x49, 0xbf, 0xc5, 0x60, 0x6b, 0x55, 0x12, 0xc9, 0x02, 0xdb, 0x81, 0xba,
	0x26, 0xcd, 0xa6, 0x89, 0x1d, 0x0b, 0x51, 0x4e, 0x88, 0x8e, 0x62, 0xc7, 0xb5, 0x61, 0xcd, 0x88,
	0xc6, 0x14, 0x54, 0x6c, 0x96, 0xda, 0xf5, 0xde, 0x66, 0xe7, 0xdb, 0x67, 0x55, 0xc7, 0xec, 0x42,
	0x23, 0xaa, 0x13, 0xd4, 0xbd, 0x4f, 0x05, 0x80, 0x63, 0x72, 0x1f, 0xb1, 0x80, 0x8d, 0x51, 0xda,
	0x0f, 0xa1, 0xb6, 0xdc, 0x97, 0xf6, 0x76, 0x9a, 0x74, 0xf1, 0xd3, 0xa7, 0x71, 0xe3, 0x8a, 0x68,
	0x52, 0xe3, 0xab, 0x02, 0xac, 0xa7, 0x8b, 0xb7, 0x77, 0xd2, 0xf8, 0xef, 0x28, 0xbd, 0xd1, 0xbc,
	0x1a, 0x10, 0xff, 0x67, 0xab, 0xf3, 0xf6, 0xf0, 0x7f, 0xfb, 0x3f, 0x3f, 0x09, 0x35, 0x27, 0xe4,
	0xaa, 0xe6, 0x9c, 0xeb, 0xd3, 0xe6, 0x88, 0x4f, 0x35, 0xca, 0xd7, 0x1f, 0x3f, 0xbf, 0x2b, 0x82,
	0x5d, 0xed, 0xce, 0x6e, 0x77, 0x4d, 0xec, 0x7e, 0xf9, 0x79, 0x51, 0xb8, 0x6e, 0x25, 0xea, 0xf3,
	0xde, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x7f, 0xb7, 0xc4, 0x65, 0x0a, 0x00, 0x00,
}
