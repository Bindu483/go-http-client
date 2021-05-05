// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: proto/cluster/cluster.proto

package cluster

import (
	proto "github.com/golang/protobuf/proto"
	deploy_app "github.com/reynencourt/rc-common-lib/v2/proto/deploy_app"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ProviderType int32

const (
	ProviderType_ProviderType_Unknown ProviderType = 0
	ProviderType_AWS                  ProviderType = 1
	ProviderType_Azure                ProviderType = 3
	ProviderType_OnPrem               ProviderType = 2
	ProviderType_GCP                  ProviderType = 4
)

// Enum value maps for ProviderType.
var (
	ProviderType_name = map[int32]string{
		0: "ProviderType_Unknown",
		1: "AWS",
		3: "Azure",
		2: "OnPrem",
		4: "GCP",
	}
	ProviderType_value = map[string]int32{
		"ProviderType_Unknown": 0,
		"AWS":                  1,
		"Azure":                3,
		"OnPrem":               2,
		"GCP":                  4,
	}
)

func (x ProviderType) Enum() *ProviderType {
	p := new(ProviderType)
	*p = x
	return p
}

func (x ProviderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProviderType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_cluster_cluster_proto_enumTypes[0].Descriptor()
}

func (ProviderType) Type() protoreflect.EnumType {
	return &file_proto_cluster_cluster_proto_enumTypes[0]
}

func (x ProviderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProviderType.Descriptor instead.
func (ProviderType) EnumDescriptor() ([]byte, []int) {
	return file_proto_cluster_cluster_proto_rawDescGZIP(), []int{0}
}

type ClusterStatus int32

const (
	ClusterStatus_ClusterStatus_Unknown   ClusterStatus = 0
	ClusterStatus_Initialising            ClusterStatus = 1
	ClusterStatus_CreateClusterInProgress ClusterStatus = 2
	ClusterStatus_Success                 ClusterStatus = 3
	ClusterStatus_CreateClusterFailed     ClusterStatus = 4
	ClusterStatus_DeleteClusterFailed     ClusterStatus = 5
	ClusterStatus_DeleteClusterInProgress ClusterStatus = 6
	ClusterStatus_NodeDeletionInProgress  ClusterStatus = 7
	ClusterStatus_NodeAdditionInProgress  ClusterStatus = 8
)

// Enum value maps for ClusterStatus.
var (
	ClusterStatus_name = map[int32]string{
		0: "ClusterStatus_Unknown",
		1: "Initialising",
		2: "CreateClusterInProgress",
		3: "Success",
		4: "CreateClusterFailed",
		5: "DeleteClusterFailed",
		6: "DeleteClusterInProgress",
		7: "NodeDeletionInProgress",
		8: "NodeAdditionInProgress",
	}
	ClusterStatus_value = map[string]int32{
		"ClusterStatus_Unknown":   0,
		"Initialising":            1,
		"CreateClusterInProgress": 2,
		"Success":                 3,
		"CreateClusterFailed":     4,
		"DeleteClusterFailed":     5,
		"DeleteClusterInProgress": 6,
		"NodeDeletionInProgress":  7,
		"NodeAdditionInProgress":  8,
	}
)

func (x ClusterStatus) Enum() *ClusterStatus {
	p := new(ClusterStatus)
	*p = x
	return p
}

func (x ClusterStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_cluster_cluster_proto_enumTypes[1].Descriptor()
}

func (ClusterStatus) Type() protoreflect.EnumType {
	return &file_proto_cluster_cluster_proto_enumTypes[1]
}

func (x ClusterStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterStatus.Descriptor instead.
func (ClusterStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_cluster_cluster_proto_rawDescGZIP(), []int{1}
}

type NodeType int32

const (
	NodeType_NodeType_Unknown NodeType = 0
	NodeType_Master           NodeType = 1
	NodeType_Etcd             NodeType = 2
	NodeType_Worker           NodeType = 3
)

// Enum value maps for NodeType.
var (
	NodeType_name = map[int32]string{
		0: "NodeType_Unknown",
		1: "Master",
		2: "Etcd",
		3: "Worker",
	}
	NodeType_value = map[string]int32{
		"NodeType_Unknown": 0,
		"Master":           1,
		"Etcd":             2,
		"Worker":           3,
	}
)

func (x NodeType) Enum() *NodeType {
	p := new(NodeType)
	*p = x
	return p
}

func (x NodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_cluster_cluster_proto_enumTypes[2].Descriptor()
}

func (NodeType) Type() protoreflect.EnumType {
	return &file_proto_cluster_cluster_proto_enumTypes[2]
}

func (x NodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeType.Descriptor instead.
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return file_proto_cluster_cluster_proto_rawDescGZIP(), []int{2}
}

type NodeStatus int32

const (
	NodeStatus_NodeStatus_Unknown     NodeStatus = 0
	NodeStatus_NodeStatusInitialising NodeStatus = 1
	NodeStatus_Ready                  NodeStatus = 3
	NodeStatus_NodeStatusFailed       NodeStatus = 4
	NodeStatus_NodeStatusNotReady     NodeStatus = 5
	NodeStatus_AddNodeInProgress      NodeStatus = 6
	NodeStatus_AddNodeFailed          NodeStatus = 7
	NodeStatus_DeleteNodeInProgress   NodeStatus = 8
	NodeStatus_DeleteNodeFailed       NodeStatus = 9
	NodeStatus_NetworkUnavailable     NodeStatus = 10
)

// Enum value maps for NodeStatus.
var (
	NodeStatus_name = map[int32]string{
		0:  "NodeStatus_Unknown",
		1:  "NodeStatusInitialising",
		3:  "Ready",
		4:  "NodeStatusFailed",
		5:  "NodeStatusNotReady",
		6:  "AddNodeInProgress",
		7:  "AddNodeFailed",
		8:  "DeleteNodeInProgress",
		9:  "DeleteNodeFailed",
		10: "NetworkUnavailable",
	}
	NodeStatus_value = map[string]int32{
		"NodeStatus_Unknown":     0,
		"NodeStatusInitialising": 1,
		"Ready":                  3,
		"NodeStatusFailed":       4,
		"NodeStatusNotReady":     5,
		"AddNodeInProgress":      6,
		"AddNodeFailed":          7,
		"DeleteNodeInProgress":   8,
		"DeleteNodeFailed":       9,
		"NetworkUnavailable":     10,
	}
)

func (x NodeStatus) Enum() *NodeStatus {
	p := new(NodeStatus)
	*p = x
	return p
}

func (x NodeStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_cluster_cluster_proto_enumTypes[3].Descriptor()
}

func (NodeStatus) Type() protoreflect.EnumType {
	return &file_proto_cluster_cluster_proto_enumTypes[3]
}

func (x NodeStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeStatus.Descriptor instead.
func (NodeStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_cluster_cluster_proto_rawDescGZIP(), []int{3}
}

type OsType int32

const (
	OsType_OsType_Unknown OsType = 0
	OsType_Ubuntu1804     OsType = 1
	OsType_Centos7        OsType = 2
	OsType_Centos8        OsType = 3
	OsType_Rhel7          OsType = 4
	OsType_Rhel8          OsType = 5
)

// Enum value maps for OsType.
var (
	OsType_name = map[int32]string{
		0: "OsType_Unknown",
		1: "Ubuntu1804",
		2: "Centos7",
		3: "Centos8",
		4: "Rhel7",
		5: "Rhel8",
	}
	OsType_value = map[string]int32{
		"OsType_Unknown": 0,
		"Ubuntu1804":     1,
		"Centos7":        2,
		"Centos8":        3,
		"Rhel7":          4,
		"Rhel8":          5,
	}
)

func (x OsType) Enum() *OsType {
	p := new(OsType)
	*p = x
	return p
}

func (x OsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OsType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_cluster_cluster_proto_enumTypes[4].Descriptor()
}

func (OsType) Type() protoreflect.EnumType {
	return &file_proto_cluster_cluster_proto_enumTypes[4]
}

func (x OsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OsType.Descriptor instead.
func (OsType) EnumDescriptor() ([]byte, []int) {
	return file_proto_cluster_cluster_proto_rawDescGZIP(), []int{4}
}

type ClusterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId     string                      `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CreatedBy     string                      `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Name          string                      `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Status        ClusterStatus               `protobuf:"varint,5,opt,name=status,proto3,enum=cluster.ClusterStatus" json:"status,omitempty"`
	Nodes         []*Node                     `protobuf:"bytes,6,rep,name=nodes,proto3" json:"nodes,omitempty"`
	ProviderType  ProviderType                `protobuf:"varint,7,opt,name=provider_type,json=providerType,proto3,enum=cluster.ProviderType" json:"provider_type,omitempty"`
	CreationDate  string                      `protobuf:"bytes,8,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	UpdatedDate   string                      `protobuf:"bytes,9,opt,name=updated_date,json=updatedDate,proto3" json:"updated_date,omitempty"`
	Template      string                      `protobuf:"bytes,10,opt,name=template,proto3" json:"template,omitempty"`
	PodCidr       string                      `protobuf:"bytes,11,opt,name=pod_cidr,json=podCidr,proto3" json:"pod_cidr,omitempty"`
	PostHooks     map[string]bool             `protobuf:"bytes,12,rep,name=post_hooks,json=postHooks,proto3" json:"post_hooks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ClusterConfig map[string]string           `protobuf:"bytes,13,rep,name=cluster_config,json=clusterConfig,proto3" json:"cluster_config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DeployFalco   bool                        `protobuf:"varint,14,opt,name=deploy_falco,json=deployFalco,proto3" json:"deploy_falco,omitempty"`
	NodePrefix    int32                       `protobuf:"varint,15,opt,name=node_prefix,json=nodePrefix,proto3" json:"node_prefix,omitempty"`
	Deployments   map[string]*DeploymentsInfo `protobuf:"bytes,16,rep,name=deployments,proto3" json:"deployments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ClusterInfo) Reset() {
	*x = ClusterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterInfo) ProtoMessage() {}

func (x *ClusterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterInfo.ProtoReflect.Descriptor instead.
func (*ClusterInfo) Descriptor() ([]byte, []int) {
	return file_proto_cluster_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClusterInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ClusterInfo) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *ClusterInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterInfo) GetStatus() ClusterStatus {
	if x != nil {
		return x.Status
	}
	return ClusterStatus_ClusterStatus_Unknown
}

func (x *ClusterInfo) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *ClusterInfo) GetProviderType() ProviderType {
	if x != nil {
		return x.ProviderType
	}
	return ProviderType_ProviderType_Unknown
}

func (x *ClusterInfo) GetCreationDate() string {
	if x != nil {
		return x.CreationDate
	}
	return ""
}

func (x *ClusterInfo) GetUpdatedDate() string {
	if x != nil {
		return x.UpdatedDate
	}
	return ""
}

func (x *ClusterInfo) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

func (x *ClusterInfo) GetPodCidr() string {
	if x != nil {
		return x.PodCidr
	}
	return ""
}

func (x *ClusterInfo) GetPostHooks() map[string]bool {
	if x != nil {
		return x.PostHooks
	}
	return nil
}

func (x *ClusterInfo) GetClusterConfig() map[string]string {
	if x != nil {
		return x.ClusterConfig
	}
	return nil
}

func (x *ClusterInfo) GetDeployFalco() bool {
	if x != nil {
		return x.DeployFalco
	}
	return false
}

func (x *ClusterInfo) GetNodePrefix() int32 {
	if x != nil {
		return x.NodePrefix
	}
	return 0
}

func (x *ClusterInfo) GetDeployments() map[string]*DeploymentsInfo {
	if x != nil {
		return x.Deployments
	}
	return nil
}

type DeploymentsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId string                      `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	SolutionId   string                      `protobuf:"bytes,2,opt,name=solution_id,json=solutionId,proto3" json:"solution_id,omitempty"`
	Version      string                      `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	SolutionIcon string                      `protobuf:"bytes,4,opt,name=solution_icon,json=solutionIcon,proto3" json:"solution_icon,omitempty"`
	Status       deploy_app.DeploymentStatus `protobuf:"varint,5,opt,name=status,proto3,enum=deploy_app.DeploymentStatus" json:"status,omitempty"`
}

func (x *DeploymentsInfo) Reset() {
	*x = DeploymentsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentsInfo) ProtoMessage() {}

func (x *DeploymentsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentsInfo.ProtoReflect.Descriptor instead.
func (*DeploymentsInfo) Descriptor() ([]byte, []int) {
	return file_proto_cluster_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *DeploymentsInfo) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *DeploymentsInfo) GetSolutionId() string {
	if x != nil {
		return x.SolutionId
	}
	return ""
}

func (x *DeploymentsInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *DeploymentsInfo) GetSolutionIcon() string {
	if x != nil {
		return x.SolutionIcon
	}
	return ""
}

func (x *DeploymentsInfo) GetStatus() deploy_app.DeploymentStatus {
	if x != nil {
		return x.Status
	}
	return deploy_app.DeploymentStatus_Unknown
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip           string     `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Number       int32      `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Type         NodeType   `protobuf:"varint,3,opt,name=type,proto3,enum=cluster.NodeType" json:"type,omitempty"`
	HostName     string     `protobuf:"bytes,4,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	OsType       OsType     `protobuf:"varint,5,opt,name=os_type,json=osType,proto3,enum=cluster.OsType" json:"os_type,omitempty"`
	Status       NodeStatus `protobuf:"varint,6,opt,name=status,proto3,enum=cluster.NodeStatus" json:"status,omitempty"`
	User         string     `protobuf:"bytes,7,opt,name=user,proto3" json:"user,omitempty"`
	CreationTime string     `protobuf:"bytes,8,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_proto_cluster_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *Node) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Node) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Node) GetType() NodeType {
	if x != nil {
		return x.Type
	}
	return NodeType_NodeType_Unknown
}

func (x *Node) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *Node) GetOsType() OsType {
	if x != nil {
		return x.OsType
	}
	return OsType_OsType_Unknown
}

func (x *Node) GetStatus() NodeStatus {
	if x != nil {
		return x.Status
	}
	return NodeStatus_NodeStatus_Unknown
}

func (x *Node) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Node) GetCreationTime() string {
	if x != nil {
		return x.CreationTime
	}
	return ""
}

var File_proto_cluster_cluster_proto protoreflect.FileDescriptor

var file_proto_cluster_cluster_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f,
	0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x06, 0x0a, 0x0b, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x12, 0x3a, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x43, 0x69, 0x64, 0x72, 0x12, 0x42, 0x0a, 0x0a, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x12,
	0x4e, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x46, 0x61, 0x6c,
	0x63, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x47, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x3c, 0x0a, 0x0e,
	0x50, 0x6f, 0x73, 0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x40, 0x0a, 0x12, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x58, 0x0a, 0x10,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xcc, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x63, 0x6f, 0x6e, 0x12,
	0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x82, 0x02, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x6f, 0x73,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6f, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x51, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x57, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x6e, 0x50, 0x72,
	0x65, 0x6d, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x43, 0x50, 0x10, 0x04, 0x2a, 0xed, 0x01,
	0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x19, 0x0a, 0x15, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x04, 0x12,
	0x17, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10,
	0x07, 0x12, 0x1a, 0x0a, 0x16, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x08, 0x2a, 0x42, 0x0a,
	0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x6f, 0x64,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x45,
	0x74, 0x63, 0x64, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x10,
	0x03, 0x2a, 0xeb, 0x01, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x73, 0x69,
	0x6e, 0x67, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x03, 0x12,
	0x14, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x05, 0x12, 0x15, 0x0a,
	0x11, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x46,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10,
	0x08, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x46,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x0a, 0x2a,
	0x5c, 0x0a, 0x06, 0x4f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x55, 0x62, 0x75, 0x6e, 0x74, 0x75, 0x31, 0x38, 0x30, 0x34, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x37, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x65,
	0x6e, 0x74, 0x6f, 0x73, 0x38, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x68, 0x65, 0x6c, 0x37,
	0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x68, 0x65, 0x6c, 0x38, 0x10, 0x05, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x79, 0x6e,
	0x65, 0x6e, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x2f, 0x72, 0x63, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2d, 0x6c, 0x69, 0x62, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cluster_cluster_proto_rawDescOnce sync.Once
	file_proto_cluster_cluster_proto_rawDescData = file_proto_cluster_cluster_proto_rawDesc
)

func file_proto_cluster_cluster_proto_rawDescGZIP() []byte {
	file_proto_cluster_cluster_proto_rawDescOnce.Do(func() {
		file_proto_cluster_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cluster_cluster_proto_rawDescData)
	})
	return file_proto_cluster_cluster_proto_rawDescData
}

var file_proto_cluster_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_proto_cluster_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_cluster_cluster_proto_goTypes = []interface{}{
	(ProviderType)(0),                // 0: cluster.ProviderType
	(ClusterStatus)(0),               // 1: cluster.ClusterStatus
	(NodeType)(0),                    // 2: cluster.NodeType
	(NodeStatus)(0),                  // 3: cluster.NodeStatus
	(OsType)(0),                      // 4: cluster.OsType
	(*ClusterInfo)(nil),              // 5: cluster.ClusterInfo
	(*DeploymentsInfo)(nil),          // 6: cluster.DeploymentsInfo
	(*Node)(nil),                     // 7: cluster.Node
	nil,                              // 8: cluster.ClusterInfo.PostHooksEntry
	nil,                              // 9: cluster.ClusterInfo.ClusterConfigEntry
	nil,                              // 10: cluster.ClusterInfo.DeploymentsEntry
	(deploy_app.DeploymentStatus)(0), // 11: deploy_app.DeploymentStatus
}
var file_proto_cluster_cluster_proto_depIdxs = []int32{
	1,  // 0: cluster.ClusterInfo.status:type_name -> cluster.ClusterStatus
	7,  // 1: cluster.ClusterInfo.nodes:type_name -> cluster.Node
	0,  // 2: cluster.ClusterInfo.provider_type:type_name -> cluster.ProviderType
	8,  // 3: cluster.ClusterInfo.post_hooks:type_name -> cluster.ClusterInfo.PostHooksEntry
	9,  // 4: cluster.ClusterInfo.cluster_config:type_name -> cluster.ClusterInfo.ClusterConfigEntry
	10, // 5: cluster.ClusterInfo.deployments:type_name -> cluster.ClusterInfo.DeploymentsEntry
	11, // 6: cluster.DeploymentsInfo.status:type_name -> deploy_app.DeploymentStatus
	2,  // 7: cluster.Node.type:type_name -> cluster.NodeType
	4,  // 8: cluster.Node.os_type:type_name -> cluster.OsType
	3,  // 9: cluster.Node.status:type_name -> cluster.NodeStatus
	6,  // 10: cluster.ClusterInfo.DeploymentsEntry.value:type_name -> cluster.DeploymentsInfo
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_proto_cluster_cluster_proto_init() }
func file_proto_cluster_cluster_proto_init() {
	if File_proto_cluster_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cluster_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cluster_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentsInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cluster_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_cluster_cluster_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_cluster_cluster_proto_goTypes,
		DependencyIndexes: file_proto_cluster_cluster_proto_depIdxs,
		EnumInfos:         file_proto_cluster_cluster_proto_enumTypes,
		MessageInfos:      file_proto_cluster_cluster_proto_msgTypes,
	}.Build()
	File_proto_cluster_cluster_proto = out.File
	file_proto_cluster_cluster_proto_rawDesc = nil
	file_proto_cluster_cluster_proto_goTypes = nil
	file_proto_cluster_cluster_proto_depIdxs = nil
}