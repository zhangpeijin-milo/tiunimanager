// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: db.proto

package db

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_db_proto protoreflect.FileDescriptor

var file_db_proto_rawDesc = []byte{
	0x0a, 0x08, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x64, 0x62, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x64, 0x62, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x64, 0x62, 0x5f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x64, 0x62, 0x5f, 0x74,
	0x69, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x64, 0x62, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbf, 0x13, 0x0a, 0x0d, 0x54, 0x69, 0x45,
	0x4d, 0x44, 0x42, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x46, 0x69,
	0x6e, 0x64, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x44, 0x42, 0x46, 0x69, 0x6e,
	0x64, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x44, 0x42, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x44, 0x42, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x44, 0x42,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x13, 0x2e, 0x44, 0x42, 0x53, 0x61, 0x76, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x44, 0x42, 0x53, 0x61, 0x76, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x46,
	0x69, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x13, 0x2e, 0x44, 0x42, 0x46, 0x69, 0x6e,
	0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x44, 0x42, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x42, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x44,
	0x42, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x44, 0x42, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x44, 0x42, 0x41,
	0x64, 0x64, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x44, 0x42, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x48, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x49, 0x6e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x44, 0x42, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x73, 0x74,
	0x73, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x44, 0x42, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x49, 0x6e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x44, 0x42, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x44, 0x42, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x73, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x44,
	0x42, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x49, 0x6e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x44, 0x42, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x4c, 0x69, 0x73,
	0x74, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x44, 0x42, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x44, 0x42, 0x4c,
	0x69, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x16, 0x2e, 0x44, 0x42, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x44, 0x42, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x42, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x48, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x17, 0x2e, 0x44, 0x42, 0x50, 0x72, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x48,
	0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x44, 0x42,
	0x50, 0x72, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x4c, 0x6f, 0x63, 0x6b, 0x48, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x13, 0x2e, 0x44, 0x42, 0x4c, 0x6f, 0x63, 0x6b, 0x48, 0x6f, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x44, 0x42, 0x4c, 0x6f, 0x63, 0x6b,
	0x48, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x1a, 0x2e, 0x44, 0x42, 0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x44, 0x42, 0x47, 0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x44, 0x42,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x44, 0x42, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x17, 0x2e, 0x44, 0x42, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x44, 0x42, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x54, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x2e, 0x44, 0x42, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x44, 0x42, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x75, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1a, 0x2e, 0x44, 0x42, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x44, 0x42, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x75, 0x70, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b,
	0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x44, 0x42,
	0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x44, 0x42, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x44, 0x42, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x44, 0x42, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x2e, 0x44,
	0x42, 0x53, 0x61, 0x76, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x44, 0x42, 0x53, 0x61, 0x76,
	0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x2e, 0x44, 0x42,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x44, 0x42, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1b, 0x2e,
	0x44, 0x42, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x44, 0x42, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e,
	0x44, 0x42, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x44, 0x42, 0x53,
	0x61, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x14, 0x53, 0x61, 0x76, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x18, 0x2e, 0x44, 0x42, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x44, 0x42, 0x53,
	0x61, 0x76, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x1e, 0x2e, 0x44, 0x42, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x44, 0x42, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x75, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x75, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x75, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x75, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x75, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x75, 0x70, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x46, 0x69,
	0x6e, 0x64, 0x54, 0x69, 0x75, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x12, 0x18,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x69, 0x75, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54,
	0x69, 0x75, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x54, 0x69, 0x75, 0x70, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x79, 0x42, 0x69, 0x7a, 0x49, 0x44, 0x12,
	0x20, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x75, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x79, 0x42, 0x69, 0x7a, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x75, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x79, 0x42, 0x69, 0x7a, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c,
	0x6f, 0x77, 0x12, 0x14, 0x2e, 0x44, 0x42, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x44, 0x42, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e,
	0x44, 0x42, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x44, 0x42, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x2e, 0x44, 0x42, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x44, 0x42, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x44, 0x42, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x44, 0x42, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x33, 0x0a, 0x08, 0x4c, 0x6f, 0x61, 0x64, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x12, 0x2e, 0x44,
	0x42, 0x4c, 0x6f, 0x61, 0x64, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x44, 0x42, 0x4c, 0x6f, 0x61, 0x64, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x4c, 0x6f, 0x61, 0x64, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x12, 0x2e, 0x44, 0x42, 0x4c, 0x6f, 0x61, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44, 0x42, 0x4c, 0x6f, 0x61, 0x64, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f,
	0x3b, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_db_proto_goTypes = []interface{}{
	(*DBFindTenantRequest)(nil),              // 0: DBFindTenantRequest
	(*DBFindAccountRequest)(nil),             // 1: DBFindAccountRequest
	(*DBSaveTokenRequest)(nil),               // 2: DBSaveTokenRequest
	(*DBFindTokenRequest)(nil),               // 3: DBFindTokenRequest
	(*DBFindRolesByPermissionRequest)(nil),   // 4: DBFindRolesByPermissionRequest
	(*DBAddHostRequest)(nil),                 // 5: DBAddHostRequest
	(*DBAddHostsInBatchRequest)(nil),         // 6: DBAddHostsInBatchRequest
	(*DBRemoveHostRequest)(nil),              // 7: DBRemoveHostRequest
	(*DBRemoveHostsInBatchRequest)(nil),      // 8: DBRemoveHostsInBatchRequest
	(*DBListHostsRequest)(nil),               // 9: DBListHostsRequest
	(*DBCheckDetailsRequest)(nil),            // 10: DBCheckDetailsRequest
	(*DBPreAllocHostsRequest)(nil),           // 11: DBPreAllocHostsRequest
	(*DBLockHostsRequest)(nil),               // 12: DBLockHostsRequest
	(*DBGetFailureDomainRequest)(nil),        // 13: DBGetFailureDomainRequest
	(*DBCreateClusterRequest)(nil),           // 14: DBCreateClusterRequest
	(*DBDeleteClusterRequest)(nil),           // 15: DBDeleteClusterRequest
	(*DBUpdateClusterStatusRequest)(nil),     // 16: DBUpdateClusterStatusRequest
	(*DBUpdateTiupConfigRequest)(nil),        // 17: DBUpdateTiupConfigRequest
	(*DBLoadClusterRequest)(nil),             // 18: DBLoadClusterRequest
	(*DBListClusterRequest)(nil),             // 19: DBListClusterRequest
	(*DBSaveBackupRecordRequest)(nil),        // 20: DBSaveBackupRecordRequest
	(*DBDeleteBackupRecordRequest)(nil),      // 21: DBDeleteBackupRecordRequest
	(*DBListBackupRecordsRequest)(nil),       // 22: DBListBackupRecordsRequest
	(*DBSaveRecoverRecordRequest)(nil),       // 23: DBSaveRecoverRecordRequest
	(*DBSaveParametersRequest)(nil),          // 24: DBSaveParametersRequest
	(*DBGetCurrentParametersRequest)(nil),    // 25: DBGetCurrentParametersRequest
	(*CreateTiupTaskRequest)(nil),            // 26: CreateTiupTaskRequest
	(*UpdateTiupTaskRequest)(nil),            // 27: UpdateTiupTaskRequest
	(*FindTiupTaskByIDRequest)(nil),          // 28: FindTiupTaskByIDRequest
	(*GetTiupTaskStatusByBizIDRequest)(nil),  // 29: GetTiupTaskStatusByBizIDRequest
	(*DBCreateFlowRequest)(nil),              // 30: DBCreateFlowRequest
	(*DBCreateTaskRequest)(nil),              // 31: DBCreateTaskRequest
	(*DBUpdateFlowRequest)(nil),              // 32: DBUpdateFlowRequest
	(*DBUpdateTaskRequest)(nil),              // 33: DBUpdateTaskRequest
	(*DBLoadFlowRequest)(nil),                // 34: DBLoadFlowRequest
	(*DBLoadTaskRequest)(nil),                // 35: DBLoadTaskRequest
	(*DBFindTenantResponse)(nil),             // 36: DBFindTenantResponse
	(*DBFindAccountResponse)(nil),            // 37: DBFindAccountResponse
	(*DBSaveTokenResponse)(nil),              // 38: DBSaveTokenResponse
	(*DBFindTokenResponse)(nil),              // 39: DBFindTokenResponse
	(*DBFindRolesByPermissionResponse)(nil),  // 40: DBFindRolesByPermissionResponse
	(*DBAddHostResponse)(nil),                // 41: DBAddHostResponse
	(*DBAddHostsInBatchResponse)(nil),        // 42: DBAddHostsInBatchResponse
	(*DBRemoveHostResponse)(nil),             // 43: DBRemoveHostResponse
	(*DBRemoveHostsInBatchResponse)(nil),     // 44: DBRemoveHostsInBatchResponse
	(*DBListHostsResponse)(nil),              // 45: DBListHostsResponse
	(*DBCheckDetailsResponse)(nil),           // 46: DBCheckDetailsResponse
	(*DBPreAllocHostsResponse)(nil),          // 47: DBPreAllocHostsResponse
	(*DBLockHostsResponse)(nil),              // 48: DBLockHostsResponse
	(*DBGetFailureDomainResponse)(nil),       // 49: DBGetFailureDomainResponse
	(*DBCreateClusterResponse)(nil),          // 50: DBCreateClusterResponse
	(*DBDeleteClusterResponse)(nil),          // 51: DBDeleteClusterResponse
	(*DBUpdateClusterStatusResponse)(nil),    // 52: DBUpdateClusterStatusResponse
	(*DBUpdateTiupConfigResponse)(nil),       // 53: DBUpdateTiupConfigResponse
	(*DBLoadClusterResponse)(nil),            // 54: DBLoadClusterResponse
	(*DBListClusterResponse)(nil),            // 55: DBListClusterResponse
	(*DBSaveBackupRecordResponse)(nil),       // 56: DBSaveBackupRecordResponse
	(*DBDeleteBackupRecordResponse)(nil),     // 57: DBDeleteBackupRecordResponse
	(*DBListBackupRecordsResponse)(nil),      // 58: DBListBackupRecordsResponse
	(*DBSaveRecoverRecordResponse)(nil),      // 59: DBSaveRecoverRecordResponse
	(*DBSaveParametersResponse)(nil),         // 60: DBSaveParametersResponse
	(*DBGetCurrentParametersResponse)(nil),   // 61: DBGetCurrentParametersResponse
	(*CreateTiupTaskResponse)(nil),           // 62: CreateTiupTaskResponse
	(*UpdateTiupTaskResponse)(nil),           // 63: UpdateTiupTaskResponse
	(*FindTiupTaskByIDResponse)(nil),         // 64: FindTiupTaskByIDResponse
	(*GetTiupTaskStatusByBizIDResponse)(nil), // 65: GetTiupTaskStatusByBizIDResponse
	(*DBCreateFlowResponse)(nil),             // 66: DBCreateFlowResponse
	(*DBCreateTaskResponse)(nil),             // 67: DBCreateTaskResponse
	(*DBUpdateFlowResponse)(nil),             // 68: DBUpdateFlowResponse
	(*DBUpdateTaskResponse)(nil),             // 69: DBUpdateTaskResponse
	(*DBLoadFlowResponse)(nil),               // 70: DBLoadFlowResponse
	(*DBLoadTaskResponse)(nil),               // 71: DBLoadTaskResponse
}
var file_db_proto_depIdxs = []int32{
	0,  // 0: TiEMDBService.FindTenant:input_type -> DBFindTenantRequest
	1,  // 1: TiEMDBService.FindAccount:input_type -> DBFindAccountRequest
	2,  // 2: TiEMDBService.SaveToken:input_type -> DBSaveTokenRequest
	3,  // 3: TiEMDBService.FindToken:input_type -> DBFindTokenRequest
	4,  // 4: TiEMDBService.FindRolesByPermission:input_type -> DBFindRolesByPermissionRequest
	5,  // 5: TiEMDBService.AddHost:input_type -> DBAddHostRequest
	6,  // 6: TiEMDBService.AddHostsInBatch:input_type -> DBAddHostsInBatchRequest
	7,  // 7: TiEMDBService.RemoveHost:input_type -> DBRemoveHostRequest
	8,  // 8: TiEMDBService.RemoveHostsInBatch:input_type -> DBRemoveHostsInBatchRequest
	9,  // 9: TiEMDBService.ListHost:input_type -> DBListHostsRequest
	10, // 10: TiEMDBService.CheckDetails:input_type -> DBCheckDetailsRequest
	11, // 11: TiEMDBService.PreAllocHosts:input_type -> DBPreAllocHostsRequest
	12, // 12: TiEMDBService.LockHosts:input_type -> DBLockHostsRequest
	13, // 13: TiEMDBService.GetFailureDomain:input_type -> DBGetFailureDomainRequest
	14, // 14: TiEMDBService.CreateCluster:input_type -> DBCreateClusterRequest
	15, // 15: TiEMDBService.DeleteCluster:input_type -> DBDeleteClusterRequest
	16, // 16: TiEMDBService.UpdateClusterStatus:input_type -> DBUpdateClusterStatusRequest
	17, // 17: TiEMDBService.UpdateClusterTiupConfig:input_type -> DBUpdateTiupConfigRequest
	18, // 18: TiEMDBService.LoadCluster:input_type -> DBLoadClusterRequest
	19, // 19: TiEMDBService.ListCluster:input_type -> DBListClusterRequest
	20, // 20: TiEMDBService.SaveBackupRecord:input_type -> DBSaveBackupRecordRequest
	21, // 21: TiEMDBService.DeleteBackupRecord:input_type -> DBDeleteBackupRecordRequest
	22, // 22: TiEMDBService.ListBackupRecords:input_type -> DBListBackupRecordsRequest
	23, // 23: TiEMDBService.SaveRecoverRecord:input_type -> DBSaveRecoverRecordRequest
	24, // 24: TiEMDBService.SaveParametersRecord:input_type -> DBSaveParametersRequest
	25, // 25: TiEMDBService.GetCurrentParametersRecord:input_type -> DBGetCurrentParametersRequest
	26, // 26: TiEMDBService.CreateTiupTask:input_type -> CreateTiupTaskRequest
	27, // 27: TiEMDBService.UpdateTiupTask:input_type -> UpdateTiupTaskRequest
	28, // 28: TiEMDBService.FindTiupTaskByID:input_type -> FindTiupTaskByIDRequest
	29, // 29: TiEMDBService.GetTiupTaskStatusByBizID:input_type -> GetTiupTaskStatusByBizIDRequest
	30, // 30: TiEMDBService.CreateFlow:input_type -> DBCreateFlowRequest
	31, // 31: TiEMDBService.CreateTask:input_type -> DBCreateTaskRequest
	32, // 32: TiEMDBService.UpdateFlow:input_type -> DBUpdateFlowRequest
	33, // 33: TiEMDBService.UpdateTask:input_type -> DBUpdateTaskRequest
	34, // 34: TiEMDBService.LoadFlow:input_type -> DBLoadFlowRequest
	35, // 35: TiEMDBService.LoadTask:input_type -> DBLoadTaskRequest
	36, // 36: TiEMDBService.FindTenant:output_type -> DBFindTenantResponse
	37, // 37: TiEMDBService.FindAccount:output_type -> DBFindAccountResponse
	38, // 38: TiEMDBService.SaveToken:output_type -> DBSaveTokenResponse
	39, // 39: TiEMDBService.FindToken:output_type -> DBFindTokenResponse
	40, // 40: TiEMDBService.FindRolesByPermission:output_type -> DBFindRolesByPermissionResponse
	41, // 41: TiEMDBService.AddHost:output_type -> DBAddHostResponse
	42, // 42: TiEMDBService.AddHostsInBatch:output_type -> DBAddHostsInBatchResponse
	43, // 43: TiEMDBService.RemoveHost:output_type -> DBRemoveHostResponse
	44, // 44: TiEMDBService.RemoveHostsInBatch:output_type -> DBRemoveHostsInBatchResponse
	45, // 45: TiEMDBService.ListHost:output_type -> DBListHostsResponse
	46, // 46: TiEMDBService.CheckDetails:output_type -> DBCheckDetailsResponse
	47, // 47: TiEMDBService.PreAllocHosts:output_type -> DBPreAllocHostsResponse
	48, // 48: TiEMDBService.LockHosts:output_type -> DBLockHostsResponse
	49, // 49: TiEMDBService.GetFailureDomain:output_type -> DBGetFailureDomainResponse
	50, // 50: TiEMDBService.CreateCluster:output_type -> DBCreateClusterResponse
	51, // 51: TiEMDBService.DeleteCluster:output_type -> DBDeleteClusterResponse
	52, // 52: TiEMDBService.UpdateClusterStatus:output_type -> DBUpdateClusterStatusResponse
	53, // 53: TiEMDBService.UpdateClusterTiupConfig:output_type -> DBUpdateTiupConfigResponse
	54, // 54: TiEMDBService.LoadCluster:output_type -> DBLoadClusterResponse
	55, // 55: TiEMDBService.ListCluster:output_type -> DBListClusterResponse
	56, // 56: TiEMDBService.SaveBackupRecord:output_type -> DBSaveBackupRecordResponse
	57, // 57: TiEMDBService.DeleteBackupRecord:output_type -> DBDeleteBackupRecordResponse
	58, // 58: TiEMDBService.ListBackupRecords:output_type -> DBListBackupRecordsResponse
	59, // 59: TiEMDBService.SaveRecoverRecord:output_type -> DBSaveRecoverRecordResponse
	60, // 60: TiEMDBService.SaveParametersRecord:output_type -> DBSaveParametersResponse
	61, // 61: TiEMDBService.GetCurrentParametersRecord:output_type -> DBGetCurrentParametersResponse
	62, // 62: TiEMDBService.CreateTiupTask:output_type -> CreateTiupTaskResponse
	63, // 63: TiEMDBService.UpdateTiupTask:output_type -> UpdateTiupTaskResponse
	64, // 64: TiEMDBService.FindTiupTaskByID:output_type -> FindTiupTaskByIDResponse
	65, // 65: TiEMDBService.GetTiupTaskStatusByBizID:output_type -> GetTiupTaskStatusByBizIDResponse
	66, // 66: TiEMDBService.CreateFlow:output_type -> DBCreateFlowResponse
	67, // 67: TiEMDBService.CreateTask:output_type -> DBCreateTaskResponse
	68, // 68: TiEMDBService.UpdateFlow:output_type -> DBUpdateFlowResponse
	69, // 69: TiEMDBService.UpdateTask:output_type -> DBUpdateTaskResponse
	70, // 70: TiEMDBService.LoadFlow:output_type -> DBLoadFlowResponse
	71, // 71: TiEMDBService.LoadTask:output_type -> DBLoadTaskResponse
	36, // [36:72] is the sub-list for method output_type
	0,  // [0:36] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_db_proto_init() }
func file_db_proto_init() {
	if File_db_proto != nil {
		return
	}
	file_db_auth_proto_init()
	file_db_host_proto_init()
	file_db_cluster_proto_init()
	file_db_tiup_proto_init()
	file_db_task_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_db_proto_goTypes,
		DependencyIndexes: file_db_proto_depIdxs,
	}.Build()
	File_db_proto = out.File
	file_db_proto_rawDesc = nil
	file_db_proto_goTypes = nil
	file_db_proto_depIdxs = nil
}