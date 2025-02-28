// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: runme/config/v1alpha1/config.proto

package configv1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type Config_FilterType int32

const (
	Config_FILTER_TYPE_UNSPECIFIED Config_FilterType = 0
	Config_FILTER_TYPE_BLOCK       Config_FilterType = 1
	Config_FILTER_TYPE_DOCUMENT    Config_FilterType = 2
)

// Enum value maps for Config_FilterType.
var (
	Config_FilterType_name = map[int32]string{
		0: "FILTER_TYPE_UNSPECIFIED",
		1: "FILTER_TYPE_BLOCK",
		2: "FILTER_TYPE_DOCUMENT",
	}
	Config_FilterType_value = map[string]int32{
		"FILTER_TYPE_UNSPECIFIED": 0,
		"FILTER_TYPE_BLOCK":       1,
		"FILTER_TYPE_DOCUMENT":    2,
	}
)

func (x Config_FilterType) Enum() *Config_FilterType {
	p := new(Config_FilterType)
	*p = x
	return p
}

func (x Config_FilterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Config_FilterType) Descriptor() protoreflect.EnumDescriptor {
	return file_runme_config_v1alpha1_config_proto_enumTypes[0].Descriptor()
}

func (Config_FilterType) Type() protoreflect.EnumType {
	return &file_runme_config_v1alpha1_config_proto_enumTypes[0]
}

func (x Config_FilterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Config_FilterType.Descriptor instead.
func (Config_FilterType) EnumDescriptor() ([]byte, []int) {
	return file_runme_config_v1alpha1_config_proto_rawDescGZIP(), []int{0, 0}
}

// Config describes the configuration of the runme tools, including CLI, server, and clients like VS Code extension.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// source is a source of Markdown files to look into.
	//
	// Types that are assignable to Source:
	//
	//	*Config_Project_
	//	*Config_Filename
	Source isConfig_Source `protobuf_oneof:"source"`
	// env is the environment variables configuration.
	Env *Config_Env `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
	// filters is a list of filters to apply.
	// Filters can be applied to documents or
	// individual code blocks.
	Filters []*Config_Filter `protobuf:"bytes,5,rep,name=filters,proto3" json:"filters,omitempty"`
	// log contains the log configuration.
	Log *Config_Log `protobuf:"bytes,7,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_config_v1alpha1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_runme_config_v1alpha1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_runme_config_v1alpha1_config_proto_rawDescGZIP(), []int{0}
}

func (m *Config) GetSource() isConfig_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *Config) GetProject() *Config_Project {
	if x, ok := x.GetSource().(*Config_Project_); ok {
		return x.Project
	}
	return nil
}

func (x *Config) GetFilename() string {
	if x, ok := x.GetSource().(*Config_Filename); ok {
		return x.Filename
	}
	return ""
}

func (x *Config) GetEnv() *Config_Env {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *Config) GetFilters() []*Config_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *Config) GetLog() *Config_Log {
	if x != nil {
		return x.Log
	}
	return nil
}

type isConfig_Source interface {
	isConfig_Source()
}

type Config_Project_ struct {
	// project indicates a dir-based source typically including multiple Markdown files.
	Project *Config_Project `protobuf:"bytes,1,opt,name=project,proto3,oneof"`
}

type Config_Filename struct {
	// filename indicates a single Markdown file.
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3,oneof"`
}

func (*Config_Project_) isConfig_Source() {}

func (*Config_Filename) isConfig_Source() {}

type Config_Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// dir is the directory to look for Markdown files.
	Dir string `protobuf:"bytes,1,opt,name=dir,proto3" json:"dir,omitempty"`
	// find_repo_upward indicates whether to find the nearest Git repository upward.
	// This is useful to, for example, recognize .gitignore files.
	FindRepoUpward bool `protobuf:"varint,2,opt,name=find_repo_upward,json=findRepoUpward,proto3" json:"find_repo_upward,omitempty"`
	// ignore_paths is a list of paths to ignore relative to dir.
	IgnorePaths []string `protobuf:"bytes,3,rep,name=ignore_paths,json=ignore,proto3" json:"ignore_paths,omitempty"`
	// disable_gitignore indicates whether to disable the .gitignore file.
	DisableGitignore bool `protobuf:"varint,4,opt,name=disable_gitignore,json=disableGitignore,proto3" json:"disable_gitignore,omitempty"`
}

func (x *Config_Project) Reset() {
	*x = Config_Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_config_v1alpha1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_Project) ProtoMessage() {}

func (x *Config_Project) ProtoReflect() protoreflect.Message {
	mi := &file_runme_config_v1alpha1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_Project.ProtoReflect.Descriptor instead.
func (*Config_Project) Descriptor() ([]byte, []int) {
	return file_runme_config_v1alpha1_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Config_Project) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

func (x *Config_Project) GetFindRepoUpward() bool {
	if x != nil {
		return x.FindRepoUpward
	}
	return false
}

func (x *Config_Project) GetIgnorePaths() []string {
	if x != nil {
		return x.IgnorePaths
	}
	return nil
}

func (x *Config_Project) GetDisableGitignore() bool {
	if x != nil {
		return x.DisableGitignore
	}
	return false
}

type Config_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type is the type of the filter.
	Type Config_FilterType `protobuf:"varint,1,opt,name=type,proto3,enum=runme.config.v1alpha1.Config_FilterType" json:"type,omitempty"`
	// condition is the filter program to execute for each document or block,
	// depending on the filter type.
	//
	// The condition should be a valid Expr expression and it should return a boolean value.
	// You can read more about the Expr syntax on https://expr-lang.org/.
	Condition string `protobuf:"bytes,2,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *Config_Filter) Reset() {
	*x = Config_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_config_v1alpha1_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_Filter) ProtoMessage() {}

func (x *Config_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_runme_config_v1alpha1_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_Filter.ProtoReflect.Descriptor instead.
func (*Config_Filter) Descriptor() ([]byte, []int) {
	return file_runme_config_v1alpha1_config_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Config_Filter) GetType() Config_FilterType {
	if x != nil {
		return x.Type
	}
	return Config_FILTER_TYPE_UNSPECIFIED
}

func (x *Config_Filter) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

type Config_Env struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// use_system_env indicates whether to use the system environment variables.
	UseSystemEnv bool `protobuf:"varint,1,opt,name=use_system_env,json=useSystemEnv,proto3" json:"use_system_env,omitempty"`
	// sources is a list of files with env.
	Sources []string `protobuf:"bytes,2,rep,name=sources,proto3" json:"sources,omitempty"`
}

func (x *Config_Env) Reset() {
	*x = Config_Env{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_config_v1alpha1_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_Env) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_Env) ProtoMessage() {}

func (x *Config_Env) ProtoReflect() protoreflect.Message {
	mi := &file_runme_config_v1alpha1_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_Env.ProtoReflect.Descriptor instead.
func (*Config_Env) Descriptor() ([]byte, []int) {
	return file_runme_config_v1alpha1_config_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Config_Env) GetUseSystemEnv() bool {
	if x != nil {
		return x.UseSystemEnv
	}
	return false
}

func (x *Config_Env) GetSources() []string {
	if x != nil {
		return x.Sources
	}
	return nil
}

type Config_Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// enabled indicates whether to enable logging.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// path is the path to the log output file.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// verbose is the verbosity level of the log.
	Verbose bool `protobuf:"varint,3,opt,name=verbose,proto3" json:"verbose,omitempty"`
}

func (x *Config_Log) Reset() {
	*x = Config_Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_config_v1alpha1_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_Log) ProtoMessage() {}

func (x *Config_Log) ProtoReflect() protoreflect.Message {
	mi := &file_runme_config_v1alpha1_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_Log.ProtoReflect.Descriptor instead.
func (*Config_Log) Descriptor() ([]byte, []int) {
	return file_runme_config_v1alpha1_config_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Config_Log) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Config_Log) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Config_Log) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

var File_runme_config_v1alpha1_config_proto protoreflect.FileDescriptor

var file_runme_config_v1alpha1_config_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x06, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x45, 0x6e, 0x76, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x3e, 0x0a, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x75, 0x6e,
	0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x33, 0x0a, 0x03, 0x6c, 0x6f, 0x67,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x1a, 0x90,
	0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x72, 0x12, 0x28, 0x0a, 0x10,
	0x66, 0x69, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x75, 0x70, 0x77, 0x61, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x66, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6f,
	0x55, 0x70, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x0c, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x67,
	0x6e, 0x6f, 0x72, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x67, 0x69, 0x74, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x10, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x69, 0x74, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x1a, 0x7a, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x72, 0x75, 0x6e, 0x6d,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x08, 0xba, 0x48, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18,
	0x80, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x45, 0x0a,
	0x03, 0x45, 0x6e, 0x76, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x65, 0x6e, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x6e, 0x76, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x1a, 0x4d, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x65, 0x22, 0x5a, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15,
	0x0a, 0x11, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4c,
	0x4f, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x43, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x42,
	0x0f, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01,
	0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_runme_config_v1alpha1_config_proto_rawDescOnce sync.Once
	file_runme_config_v1alpha1_config_proto_rawDescData = file_runme_config_v1alpha1_config_proto_rawDesc
)

func file_runme_config_v1alpha1_config_proto_rawDescGZIP() []byte {
	file_runme_config_v1alpha1_config_proto_rawDescOnce.Do(func() {
		file_runme_config_v1alpha1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_runme_config_v1alpha1_config_proto_rawDescData)
	})
	return file_runme_config_v1alpha1_config_proto_rawDescData
}

var file_runme_config_v1alpha1_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_runme_config_v1alpha1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_runme_config_v1alpha1_config_proto_goTypes = []interface{}{
	(Config_FilterType)(0), // 0: runme.config.v1alpha1.Config.FilterType
	(*Config)(nil),         // 1: runme.config.v1alpha1.Config
	(*Config_Project)(nil), // 2: runme.config.v1alpha1.Config.Project
	(*Config_Filter)(nil),  // 3: runme.config.v1alpha1.Config.Filter
	(*Config_Env)(nil),     // 4: runme.config.v1alpha1.Config.Env
	(*Config_Log)(nil),     // 5: runme.config.v1alpha1.Config.Log
}
var file_runme_config_v1alpha1_config_proto_depIdxs = []int32{
	2, // 0: runme.config.v1alpha1.Config.project:type_name -> runme.config.v1alpha1.Config.Project
	4, // 1: runme.config.v1alpha1.Config.env:type_name -> runme.config.v1alpha1.Config.Env
	3, // 2: runme.config.v1alpha1.Config.filters:type_name -> runme.config.v1alpha1.Config.Filter
	5, // 3: runme.config.v1alpha1.Config.log:type_name -> runme.config.v1alpha1.Config.Log
	0, // 4: runme.config.v1alpha1.Config.Filter.type:type_name -> runme.config.v1alpha1.Config.FilterType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_runme_config_v1alpha1_config_proto_init() }
func file_runme_config_v1alpha1_config_proto_init() {
	if File_runme_config_v1alpha1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runme_config_v1alpha1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_runme_config_v1alpha1_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_Project); i {
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
		file_runme_config_v1alpha1_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_Filter); i {
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
		file_runme_config_v1alpha1_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_Env); i {
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
		file_runme_config_v1alpha1_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_Log); i {
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
	file_runme_config_v1alpha1_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Config_Project_)(nil),
		(*Config_Filename)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_runme_config_v1alpha1_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_runme_config_v1alpha1_config_proto_goTypes,
		DependencyIndexes: file_runme_config_v1alpha1_config_proto_depIdxs,
		EnumInfos:         file_runme_config_v1alpha1_config_proto_enumTypes,
		MessageInfos:      file_runme_config_v1alpha1_config_proto_msgTypes,
	}.Build()
	File_runme_config_v1alpha1_config_proto = out.File
	file_runme_config_v1alpha1_config_proto_rawDesc = nil
	file_runme_config_v1alpha1_config_proto_goTypes = nil
	file_runme_config_v1alpha1_config_proto_depIdxs = nil
}
