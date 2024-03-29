// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.3
// source: task.proto

package pb

import (
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

//Режим торгов инструмента
type TaskStatusCode int32

const (
	TaskStatusCode_TASK_STATUS_OK         TaskStatusCode = 0
	TaskStatusCode_TASK_STATUS_NOT_FOUND  TaskStatusCode = 1
	TaskStatusCode_TASK_STATUS_RESTRICTED TaskStatusCode = 2
)

// Enum value maps for TaskStatusCode.
var (
	TaskStatusCode_name = map[int32]string{
		0: "TASK_STATUS_OK",
		1: "TASK_STATUS_NOT_FOUND",
		2: "TASK_STATUS_RESTRICTED",
	}
	TaskStatusCode_value = map[string]int32{
		"TASK_STATUS_OK":         0,
		"TASK_STATUS_NOT_FOUND":  1,
		"TASK_STATUS_RESTRICTED": 2,
	}
)

func (x TaskStatusCode) Enum() *TaskStatusCode {
	p := new(TaskStatusCode)
	*p = x
	return p
}

func (x TaskStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_task_proto_enumTypes[0].Descriptor()
}

func (TaskStatusCode) Type() protoreflect.EnumType {
	return &file_task_proto_enumTypes[0]
}

func (x TaskStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatusCode.Descriptor instead.
func (TaskStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

// TaskDefinition is task definition as received from server
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        TaskStatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=TaskStatusCode" json:"code,omitempty"`
	Name        string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Interpreter string         `protobuf:"bytes,3,opt,name=interpreter,proto3" json:"interpreter,omitempty"`
	Source      []byte         `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	CreatedAt   uint64         `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   uint64         `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetCode() TaskStatusCode {
	if x != nil {
		return x.Code
	}
	return TaskStatusCode_TASK_STATUS_OK
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetInterpreter() string {
	if x != nil {
		return x.Interpreter
	}
	return ""
}

func (x *Task) GetSource() []byte {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Task) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Task) GetUpdatedAt() uint64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_task_proto protoreflect.FileDescriptor

var file_task_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a,
	0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x23, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x2a, 0x5b, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x41,
	0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x66, 0x6c, 0x69, 0x63, 0x2e, 0x72, 0x75, 0x2f,
	0x76, 0x6f, 0x64, 0x6f, 0x6c, 0x61, 0x7a, 0x30, 0x39, 0x35, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_proto_rawDescOnce sync.Once
	file_task_proto_rawDescData = file_task_proto_rawDesc
)

func file_task_proto_rawDescGZIP() []byte {
	file_task_proto_rawDescOnce.Do(func() {
		file_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_proto_rawDescData)
	})
	return file_task_proto_rawDescData
}

var file_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_task_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_task_proto_goTypes = []interface{}{
	(TaskStatusCode)(0), // 0: TaskStatusCode
	(*Task)(nil),        // 1: Task
}
var file_task_proto_depIdxs = []int32{
	0, // 0: Task.code:type_name -> TaskStatusCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_task_proto_init() }
func file_task_proto_init() {
	if File_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
			RawDescriptor: file_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_task_proto_goTypes,
		DependencyIndexes: file_task_proto_depIdxs,
		EnumInfos:         file_task_proto_enumTypes,
		MessageInfos:      file_task_proto_msgTypes,
	}.Build()
	File_task_proto = out.File
	file_task_proto_rawDesc = nil
	file_task_proto_goTypes = nil
	file_task_proto_depIdxs = nil
}
