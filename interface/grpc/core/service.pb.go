// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/mesg-foundation/core/interface/grpc/core/service.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// This is the definition of a MESG Service.
type Service struct {
	Id                   string                 `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Tasks                map[string]*Task       `protobuf:"bytes,5,rep,name=tasks,proto3" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Events               map[string]*Event      `protobuf:"bytes,6,rep,name=events,proto3" json:"events,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Dependencies         map[string]*Dependency `protobuf:"bytes,7,rep,name=dependencies,proto3" json:"dependencies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Configuration        *Dependency            `protobuf:"bytes,8,opt,name=configuration,proto3" json:"configuration,omitempty"`
	Repository           string                 `protobuf:"bytes,9,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_4f7f9adec6a93c01, []int{0}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (dst *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(dst, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service) GetTasks() map[string]*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *Service) GetEvents() map[string]*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Service) GetDependencies() map[string]*Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *Service) GetConfiguration() *Dependency {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *Service) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

// Events are emitted by the service whenever the service wants.
// TODO(ilgooz) remove key, serviceName fields when Event type crafted manually.
type Event struct {
	Key                  string                `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string                `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ServiceName          string                `protobuf:"bytes,5,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Data                 map[string]*Parameter `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_4f7f9adec6a93c01, []int{1}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Event) GetData() map[string]*Parameter {
	if m != nil {
		return m.Data
	}
	return nil
}

// A task is a function that requires inputs and returns output.
// TODO(ilgooz) remove key, serviceName fields when Task type crafted manually.
type Task struct {
	Key                  string                `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string                `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ServiceName          string                `protobuf:"bytes,9,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Inputs               map[string]*Parameter `protobuf:"bytes,6,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Outputs              map[string]*Output    `protobuf:"bytes,7,rep,name=outputs,proto3" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_4f7f9adec6a93c01, []int{2}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Task) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Task) GetInputs() map[string]*Parameter {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Task) GetOutputs() map[string]*Output {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// A output is the data a task must return.
// TODO(ilgooz) remove key, taskKey, serviceName fields when Output type crafted manually.
type Output struct {
	Key                  string                `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string                `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	TaskKey              string                `protobuf:"bytes,5,opt,name=taskKey,proto3" json:"taskKey,omitempty"`
	ServiceName          string                `protobuf:"bytes,6,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Data                 map[string]*Parameter `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_4f7f9adec6a93c01, []int{3}
}
func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (dst *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(dst, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Output) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Output) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Output) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *Output) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Output) GetData() map[string]*Parameter {
	if m != nil {
		return m.Data
	}
	return nil
}

// A parameter is the definition of a specific value.
type Parameter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Optional             bool     `protobuf:"varint,4,opt,name=optional,proto3" json:"optional,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_4f7f9adec6a93c01, []int{4}
}
func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (dst *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(dst, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Parameter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Parameter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Parameter) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

// A dependency is a configuration of an other Docker container that runs separately from the service.
type Dependency struct {
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Volumes              []string `protobuf:"bytes,2,rep,name=volumes,proto3" json:"volumes,omitempty"`
	Volumesfrom          []string `protobuf:"bytes,3,rep,name=volumesfrom,proto3" json:"volumesfrom,omitempty"`
	Ports                []string `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty"`
	Command              string   `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dependency) Reset()         { *m = Dependency{} }
func (m *Dependency) String() string { return proto.CompactTextString(m) }
func (*Dependency) ProtoMessage()    {}
func (*Dependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_4f7f9adec6a93c01, []int{5}
}
func (m *Dependency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dependency.Unmarshal(m, b)
}
func (m *Dependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dependency.Marshal(b, m, deterministic)
}
func (dst *Dependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dependency.Merge(dst, src)
}
func (m *Dependency) XXX_Size() int {
	return xxx_messageInfo_Dependency.Size(m)
}
func (m *Dependency) XXX_DiscardUnknown() {
	xxx_messageInfo_Dependency.DiscardUnknown(m)
}

var xxx_messageInfo_Dependency proto.InternalMessageInfo

func (m *Dependency) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Dependency) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *Dependency) GetVolumesfrom() []string {
	if m != nil {
		return m.Volumesfrom
	}
	return nil
}

func (m *Dependency) GetPorts() []string {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Dependency) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func init() {
	proto.RegisterType((*Service)(nil), "core.Service")
	proto.RegisterMapType((map[string]*Dependency)(nil), "core.Service.DependenciesEntry")
	proto.RegisterMapType((map[string]*Event)(nil), "core.Service.EventsEntry")
	proto.RegisterMapType((map[string]*Task)(nil), "core.Service.TasksEntry")
	proto.RegisterType((*Event)(nil), "core.Event")
	proto.RegisterMapType((map[string]*Parameter)(nil), "core.Event.DataEntry")
	proto.RegisterType((*Task)(nil), "core.Task")
	proto.RegisterMapType((map[string]*Parameter)(nil), "core.Task.InputsEntry")
	proto.RegisterMapType((map[string]*Output)(nil), "core.Task.OutputsEntry")
	proto.RegisterType((*Output)(nil), "core.Output")
	proto.RegisterMapType((map[string]*Parameter)(nil), "core.Output.DataEntry")
	proto.RegisterType((*Parameter)(nil), "core.Parameter")
	proto.RegisterType((*Dependency)(nil), "core.Dependency")
}

func init() {
	proto.RegisterFile("github.com/mesg-foundation/core/interface/grpc/core/service.proto", fileDescriptor_service_4f7f9adec6a93c01)
}

var fileDescriptor_service_4f7f9adec6a93c01 = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x8e, 0xd3, 0x3c,
	0x10, 0x55, 0xd3, 0x34, 0x6d, 0x26, 0xfb, 0x7d, 0x2c, 0x16, 0x2c, 0x26, 0x17, 0x10, 0x2a, 0x81,
	0x16, 0x24, 0x52, 0xed, 0x22, 0x21, 0xc4, 0x1d, 0x62, 0x17, 0x2d, 0x20, 0xf1, 0x13, 0x78, 0x01,
	0x6f, 0xe2, 0x16, 0x6b, 0x9b, 0x38, 0x38, 0x4e, 0xa5, 0xdc, 0xf2, 0x00, 0x3c, 0x0b, 0x4f, 0x85,
	0xc4, 0x5b, 0x20, 0xdb, 0x49, 0xeb, 0xd0, 0x5e, 0x2d, 0xcb, 0x9d, 0x7d, 0x66, 0xce, 0x99, 0xcc,
	0x99, 0x71, 0x0b, 0x2f, 0x16, 0x4c, 0x7e, 0xa9, 0xcf, 0xe3, 0x94, 0xe7, 0xb3, 0x9c, 0x56, 0x8b,
	0xc7, 0x73, 0x5e, 0x17, 0x19, 0x91, 0x8c, 0x17, 0xb3, 0x94, 0x0b, 0x3a, 0x63, 0x85, 0xa4, 0x62,
	0x4e, 0x52, 0x3a, 0x5b, 0x88, 0x32, 0x35, 0x58, 0x45, 0xc5, 0x8a, 0xa5, 0x34, 0x2e, 0x05, 0x97,
	0x1c, 0xb9, 0x0a, 0x9b, 0xfe, 0x70, 0x61, 0xfc, 0xc9, 0xe0, 0xe8, 0x7f, 0x70, 0x58, 0x86, 0x21,
	0x1a, 0x1c, 0xfa, 0x89, 0xc3, 0x32, 0x84, 0xc0, 0x2d, 0x48, 0x4e, 0xf1, 0x40, 0x23, 0xfa, 0x8c,
	0x22, 0x08, 0x32, 0x5a, 0xa5, 0x82, 0x95, 0xaa, 0x16, 0x76, 0x74, 0xc8, 0x86, 0x50, 0x0c, 0x23,
	0x49, 0xaa, 0x8b, 0x0a, 0x8f, 0xa2, 0xe1, 0x61, 0x70, 0x8c, 0x63, 0x55, 0x27, 0x6e, 0x6b, 0xc4,
	0x9f, 0x55, 0xe8, 0xb4, 0x90, 0xa2, 0x49, 0x4c, 0x1a, 0x3a, 0x02, 0x8f, 0xae, 0x68, 0x21, 0x2b,
	0xec, 0x69, 0xc2, 0xed, 0x3e, 0xe1, 0x54, 0xc7, 0x0c, 0xa3, 0x4d, 0x44, 0x2f, 0x61, 0x2f, 0xa3,
	0x25, 0x2d, 0x32, 0x5a, 0xa4, 0x8c, 0x56, 0x78, 0xac, 0x89, 0x77, 0xfb, 0xc4, 0x13, 0x2b, 0xc3,
	0xd0, 0x7b, 0x24, 0xf4, 0x14, 0xfe, 0x4b, 0x79, 0x31, 0x67, 0x8b, 0x5a, 0x68, 0xdf, 0xf0, 0x24,
	0x1a, 0x1c, 0x06, 0xc7, 0xfb, 0x46, 0x65, 0xcd, 0x6e, 0x92, 0x7e, 0x1a, 0xba, 0x03, 0x20, 0x68,
	0xc9, 0x2b, 0x26, 0xb9, 0x68, 0xb0, 0xaf, 0x0d, 0xb0, 0x90, 0xf0, 0x04, 0x60, 0xd3, 0x24, 0xda,
	0x87, 0xe1, 0x05, 0x6d, 0x5a, 0x0b, 0xd5, 0x11, 0x45, 0x30, 0x5a, 0x91, 0x65, 0x4d, 0xb5, 0x77,
	0xc1, 0x31, 0x98, 0x7a, 0x8a, 0x92, 0x98, 0xc0, 0x73, 0xe7, 0xd9, 0x20, 0x7c, 0x05, 0x81, 0xd5,
	0xf9, 0x0e, 0x99, 0x7b, 0x7d, 0x99, 0xc0, 0xc8, 0x68, 0x8e, 0xad, 0xf3, 0x11, 0xae, 0x6f, 0x19,
	0xb1, 0x43, 0xed, 0x41, 0x5f, 0x6d, 0xdb, 0x84, 0x8d, 0xe4, 0xf4, 0xe7, 0x00, 0x46, 0xba, 0x4e,
	0xa7, 0xe3, 0x6e, 0x74, 0x2e, 0xb7, 0x32, 0x11, 0x04, 0xed, 0x6e, 0xbe, 0x53, 0xe4, 0x91, 0xc9,
	0xb0, 0x20, 0xf4, 0x10, 0xdc, 0x8c, 0x48, 0x82, 0x87, 0x7a, 0xd2, 0x37, 0xad, 0x66, 0xe3, 0x13,
	0x22, 0x89, 0x99, 0xaf, 0x4e, 0x09, 0xcf, 0xc0, 0x5f, 0x43, 0x3b, 0x3a, 0xbd, 0xdf, 0xef, 0xf4,
	0x9a, 0x91, 0xfa, 0x40, 0x04, 0xc9, 0xa9, 0xa4, 0xc2, 0x6e, 0xf4, 0x97, 0x03, 0xae, 0x9a, 0x4b,
	0xa7, 0x32, 0xb9, 0xe2, 0x3e, 0xfd, 0xed, 0x3e, 0x63, 0xf0, 0x58, 0x51, 0xd6, 0xeb, 0xc7, 0x70,
	0xb0, 0xd9, 0x8e, 0xf8, 0xb5, 0x0e, 0xb4, 0x2f, 0xc1, 0x64, 0xa1, 0x23, 0x18, 0xf3, 0x5a, 0x6a,
	0x82, 0x79, 0x04, 0xb7, 0x2c, 0xc2, 0x7b, 0x13, 0x31, 0x8c, 0x2e, 0x2f, 0x7c, 0x03, 0x81, 0xa5,
	0xf4, 0x57, 0x0e, 0x85, 0x67, 0xb0, 0x67, 0x17, 0xd9, 0x21, 0x36, 0xed, 0x8b, 0xed, 0x19, 0x31,
	0x43, 0xb2, 0xbd, 0xfe, 0xe6, 0x80, 0x67, 0xd0, 0x2b, 0xdb, 0x2a, 0x0c, 0x63, 0xf5, 0x0b, 0xf3,
	0x96, 0x36, 0xed, 0x46, 0x75, 0xd7, 0x3f, 0xe7, 0xe0, 0x6d, 0xcf, 0xe1, 0x51, 0x6f, 0xdf, 0x0e,
	0xec, 0xaf, 0xfe, 0x87, 0x0b, 0xf7, 0x15, 0xfc, 0x35, 0x7e, 0xc9, 0xa6, 0x11, 0xb8, 0xb2, 0x29,
	0x29, 0x1e, 0x1a, 0x96, 0x3a, 0xa3, 0x10, 0x26, 0x5c, 0x47, 0xc9, 0x52, 0xbb, 0x3a, 0x49, 0xd6,
	0xf7, 0xe9, 0xf7, 0x01, 0xc0, 0xe6, 0x99, 0xa3, 0x1b, 0x30, 0x62, 0x39, 0x59, 0x74, 0x55, 0xcd,
	0x45, 0x39, 0xb9, 0xe2, 0xcb, 0x3a, 0xa7, 0x15, 0x76, 0xa2, 0xa1, 0x72, 0xb2, 0xbd, 0xaa, 0x0f,
	0x6a, 0x8f, 0x73, 0xc1, 0x73, 0x6d, 0x97, 0x9f, 0xd8, 0x90, 0x52, 0x2c, 0xb9, 0x90, 0x15, 0x76,
	0x75, 0xcc, 0x5c, 0x94, 0x62, 0xca, 0xf3, 0x9c, 0x14, 0x59, 0x37, 0x9b, 0xf6, 0x7a, 0xee, 0xe9,
	0x7f, 0xa7, 0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x65, 0xaa, 0xd2, 0xd5, 0xe2, 0x06, 0x00,
	0x00,
}
