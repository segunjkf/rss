// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: google/api/apikeys/v2/resources.proto

package apikeyspb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The representation of a key managed by the API Keys API.
type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the key.
	// The `name` has the form:
	// `projects/<PROJECT_NUMBER>/locations/global/keys/<KEY_ID>`.
	// For example:
	// `projects/123456867718/locations/global/keys/b7ff1f9f-8275-410a-94dd-3855ee9b5dd2`
	//
	// NOTE: Key is a global resource; hence the only supported value for
	// location is `global`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Unique id in UUID4 format.
	Uid string `protobuf:"bytes,5,opt,name=uid,proto3" json:"uid,omitempty"`
	// Human-readable display name of this key that you can modify.
	// The maximum length is 63 characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. An encrypted and signed value held by this key.
	// This field can be accessed only through the `GetKeyString` method.
	KeyString string `protobuf:"bytes,3,opt,name=key_string,json=keyString,proto3" json:"key_string,omitempty"`
	// Output only. A timestamp identifying the time this key was originally
	// created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. A timestamp identifying the time this key was last
	// updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. A timestamp when this key was deleted. If the resource is not deleted,
	// this must be empty.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// Annotations is an unstructured key-value map stored with a policy that
	// may be set by external tools to store and retrieve arbitrary metadata.
	// They are not queryable and should be preserved when modifying objects.
	Annotations map[string]string `protobuf:"bytes,8,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Key restrictions.
	Restrictions *Restrictions `protobuf:"bytes,9,opt,name=restrictions,proto3" json:"restrictions,omitempty"`
	// Output only. A checksum computed by the server based on the current value of the Key
	// resource. This may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	// See https://google.aip.dev/154.
	Etag string `protobuf:"bytes,11,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_google_api_apikeys_v2_resources_proto_rawDescGZIP(), []int{0}
}

func (x *Key) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Key) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Key) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Key) GetKeyString() string {
	if x != nil {
		return x.KeyString
	}
	return ""
}

func (x *Key) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Key) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Key) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *Key) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Key) GetRestrictions() *Restrictions {
	if x != nil {
		return x.Restrictions
	}
	return nil
}

func (x *Key) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

// Describes the restrictions on the key.
type Restrictions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The websites, IP addresses, Android apps, or iOS apps (the clients) that
	// are allowed to use the key. You can specify only one type of client
	// restrictions per key.
	//
	// Types that are assignable to ClientRestrictions:
	//	*Restrictions_BrowserKeyRestrictions
	//	*Restrictions_ServerKeyRestrictions
	//	*Restrictions_AndroidKeyRestrictions
	//	*Restrictions_IosKeyRestrictions
	ClientRestrictions isRestrictions_ClientRestrictions `protobuf_oneof:"client_restrictions"`
	// A restriction for a specific service and optionally one or
	// more specific methods. Requests are allowed if they
	// match any of these restrictions. If no restrictions are
	// specified, all targets are allowed.
	ApiTargets []*ApiTarget `protobuf:"bytes,5,rep,name=api_targets,json=apiTargets,proto3" json:"api_targets,omitempty"`
}

func (x *Restrictions) Reset() {
	*x = Restrictions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restrictions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restrictions) ProtoMessage() {}

func (x *Restrictions) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restrictions.ProtoReflect.Descriptor instead.
func (*Restrictions) Descriptor() ([]byte, []int) {
	return file_google_api_apikeys_v2_resources_proto_rawDescGZIP(), []int{1}
}

func (m *Restrictions) GetClientRestrictions() isRestrictions_ClientRestrictions {
	if m != nil {
		return m.ClientRestrictions
	}
	return nil
}

func (x *Restrictions) GetBrowserKeyRestrictions() *BrowserKeyRestrictions {
	if x, ok := x.GetClientRestrictions().(*Restrictions_BrowserKeyRestrictions); ok {
		return x.BrowserKeyRestrictions
	}
	return nil
}

func (x *Restrictions) GetServerKeyRestrictions() *ServerKeyRestrictions {
	if x, ok := x.GetClientRestrictions().(*Restrictions_ServerKeyRestrictions); ok {
		return x.ServerKeyRestrictions
	}
	return nil
}

func (x *Restrictions) GetAndroidKeyRestrictions() *AndroidKeyRestrictions {
	if x, ok := x.GetClientRestrictions().(*Restrictions_AndroidKeyRestrictions); ok {
		return x.AndroidKeyRestrictions
	}
	return nil
}

func (x *Restrictions) GetIosKeyRestrictions() *IosKeyRestrictions {
	if x, ok := x.GetClientRestrictions().(*Restrictions_IosKeyRestrictions); ok {
		return x.IosKeyRestrictions
	}
	return nil
}

func (x *Restrictions) GetApiTargets() []*ApiTarget {
	if x != nil {
		return x.ApiTargets
	}
	return nil
}

type isRestrictions_ClientRestrictions interface {
	isRestrictions_ClientRestrictions()
}

type Restrictions_BrowserKeyRestrictions struct {
	// The HTTP referrers (websites) that are allowed to use the key.
	BrowserKeyRestrictions *BrowserKeyRestrictions `protobuf:"bytes,1,opt,name=browser_key_restrictions,json=browserKeyRestrictions,proto3,oneof"`
}

type Restrictions_ServerKeyRestrictions struct {
	// The IP addresses of callers that are allowed to use the key.
	ServerKeyRestrictions *ServerKeyRestrictions `protobuf:"bytes,2,opt,name=server_key_restrictions,json=serverKeyRestrictions,proto3,oneof"`
}

type Restrictions_AndroidKeyRestrictions struct {
	// The Android apps that are allowed to use the key.
	AndroidKeyRestrictions *AndroidKeyRestrictions `protobuf:"bytes,3,opt,name=android_key_restrictions,json=androidKeyRestrictions,proto3,oneof"`
}

type Restrictions_IosKeyRestrictions struct {
	// The iOS apps that are allowed to use the key.
	IosKeyRestrictions *IosKeyRestrictions `protobuf:"bytes,4,opt,name=ios_key_restrictions,json=iosKeyRestrictions,proto3,oneof"`
}

func (*Restrictions_BrowserKeyRestrictions) isRestrictions_ClientRestrictions() {}

func (*Restrictions_ServerKeyRestrictions) isRestrictions_ClientRestrictions() {}

func (*Restrictions_AndroidKeyRestrictions) isRestrictions_ClientRestrictions() {}

func (*Restrictions_IosKeyRestrictions) isRestrictions_ClientRestrictions() {}

// The HTTP referrers (websites) that are allowed to use the key.
type BrowserKeyRestrictions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of regular expressions for the referrer URLs that are allowed
	// to make API calls with this key.
	AllowedReferrers []string `protobuf:"bytes,1,rep,name=allowed_referrers,json=allowedReferrers,proto3" json:"allowed_referrers,omitempty"`
}

func (x *BrowserKeyRestrictions) Reset() {
	*x = BrowserKeyRestrictions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrowserKeyRestrictions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrowserKeyRestrictions) ProtoMessage() {}

func (x *BrowserKeyRestrictions) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrowserKeyRestrictions.ProtoReflect.Descriptor instead.
func (*BrowserKeyRestrictions) Descriptor() ([]byte, []int) {
	return file_google_api_apikeys_v2_resources_proto_rawDescGZIP(), []int{2}
}

func (x *BrowserKeyRestrictions) GetAllowedReferrers() []string {
	if x != nil {
		return x.AllowedReferrers
	}
	return nil
}

// The IP addresses of callers that are allowed to use the key.
type ServerKeyRestrictions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of the caller IP addresses that are allowed to make API calls
	// with this key.
	AllowedIps []string `protobuf:"bytes,1,rep,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`
}

func (x *ServerKeyRestrictions) Reset() {
	*x = ServerKeyRestrictions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerKeyRestrictions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerKeyRestrictions) ProtoMessage() {}

func (x *ServerKeyRestrictions) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerKeyRestrictions.ProtoReflect.Descriptor instead.
func (*ServerKeyRestrictions) Descriptor() ([]byte, []int) {
	return file_google_api_apikeys_v2_resources_proto_rawDescGZIP(), []int{3}
}

func (x *ServerKeyRestrictions) GetAllowedIps() []string {
	if x != nil {
		return x.AllowedIps
	}
	return nil
}

// The Android apps that are allowed to use the key.
type AndroidKeyRestrictions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of Android applications that are allowed to make API calls with
	// this key.
	AllowedApplications []*AndroidApplication `protobuf:"bytes,1,rep,name=allowed_applications,json=allowedApplications,proto3" json:"allowed_applications,omitempty"`
}

func (x *AndroidKeyRestrictions) Reset() {
	*x = AndroidKeyRestrictions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AndroidKeyRestrictions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AndroidKeyRestrictions) ProtoMessage() {}

func (x *AndroidKeyRestrictions) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AndroidKeyRestrictions.ProtoReflect.Descriptor instead.
func (*AndroidKeyRestrictions) Descriptor() ([]byte, []int) {
	return file_google_api_apikeys_v2_resources_proto_rawDescGZIP(), []int{4}
}

func (x *AndroidKeyRestrictions) GetAllowedApplications() []*AndroidApplication {
	if x != nil {
		return x.AllowedApplications
	}
	return nil
}

// Identifier of an Android application for key use.
type AndroidApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The SHA1 fingerprint of the application. For example, both sha1 formats are
	// acceptable : DA:39:A3:EE:5E:6B:4B:0D:32:55:BF:EF:95:60:18:90:AF:D8:07:09 or
	// DA39A3EE5E6B4B0D3255BFEF95601890AFD80709.
	// Output format is the latter.
	Sha1Fingerprint string `protobuf:"bytes,1,opt,name=sha1_fingerprint,json=sha1Fingerprint,proto3" json:"sha1_fingerprint,omitempty"`
	// The package name of the application.
	PackageName string `protobuf:"bytes,2,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
}

func (x *AndroidApplication) Reset() {
	*x = AndroidApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AndroidApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AndroidApplication) ProtoMessage() {}

func (x *AndroidApplication) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AndroidApplication.ProtoReflect.Descriptor instead.
func (*AndroidApplication) Descriptor() ([]byte, []int) {
	return file_google_api_apikeys_v2_resources_proto_rawDescGZIP(), []int{5}
}

func (x *AndroidApplication) GetSha1Fingerprint() string {
	if x != nil {
		return x.Sha1Fingerprint
	}
	return ""
}

func (x *AndroidApplication) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

// The iOS apps that are allowed to use the key.
type IosKeyRestrictions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of bundle IDs that are allowed when making API calls with this key.
	AllowedBundleIds []string `protobuf:"bytes,1,rep,name=allowed_bundle_ids,json=allowedBundleIds,proto3" json:"allowed_bundle_ids,omitempty"`
}

func (x *IosKeyRestrictions) Reset() {
	*x = IosKeyRestrictions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IosKeyRestrictions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IosKeyRestrictions) ProtoMessage() {}

func (x *IosKeyRestrictions) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IosKeyRestrictions.ProtoReflect.Descriptor instead.
func (*IosKeyRestrictions) Descriptor() ([]byte, []int) {
	return file_google_api_apikeys_v2_resources_proto_rawDescGZIP(), []int{6}
}

func (x *IosKeyRestrictions) GetAllowedBundleIds() []string {
	if x != nil {
		return x.AllowedBundleIds
	}
	return nil
}

// A restriction for a specific service and optionally one or multiple
// specific methods. Both fields are case insensitive.
type ApiTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The service for this restriction. It should be the canonical
	// service name, for example: `translate.googleapis.com`.
	// You can use [`gcloud services list`](/sdk/gcloud/reference/services/list)
	// to get a list of services that are enabled in the project.
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// Optional. List of one or more methods that can be called.
	// If empty, all methods for the service are allowed. A wildcard
	// (*) can be used as the last symbol.
	// Valid examples:
	//   `google.cloud.translate.v2.TranslateService.GetSupportedLanguage`
	//   `TranslateText`
	//   `Get*`
	//   `translate.googleapis.com.Get*`
	Methods []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *ApiTarget) Reset() {
	*x = ApiTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiTarget) ProtoMessage() {}

func (x *ApiTarget) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_apikeys_v2_resources_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiTarget.ProtoReflect.Descriptor instead.
func (*ApiTarget) Descriptor() ([]byte, []int) {
	return file_google_api_apikeys_v2_resources_proto_rawDescGZIP(), []int{7}
}

func (x *ApiTarget) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ApiTarget) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

var File_google_api_apikeys_v2_resources_proto protoreflect.FileDescriptor

var file_google_api_apikeys_v2_resources_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x6b, 0x65, 0x79, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x05, 0x0a, 0x03,
	0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x09, 0x6b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x4d, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4b, 0x65,
	0x79, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x47, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x65, 0x74, 0x61,
	0x67, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x3a, 0x61, 0xea, 0x41, 0x5e, 0x0a, 0x1a, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b,
	0x65, 0x79, 0x12, 0x32, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6b, 0x65, 0x79, 0x73,
	0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x7d, 0x2a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x32, 0x03, 0x6b, 0x65,
	0x79, 0x52, 0x01, 0x01, 0x22, 0x85, 0x04, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x69, 0x0a, 0x18, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72,
	0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x16, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65,
	0x72, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x66, 0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x72,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48,
	0x00, 0x52, 0x15, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x69, 0x0a, 0x18, 0x61, 0x6e, 0x64, 0x72,
	0x6f, 0x69, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x16, 0x61, 0x6e, 0x64,
	0x72, 0x6f, 0x69, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x5d, 0x0a, 0x14, 0x69, 0x6f, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x72,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6f, 0x73, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x12,
	0x69, 0x6f, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x41, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x41, 0x70, 0x69, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x42, 0x15, 0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x45, 0x0a, 0x16,
	0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x72, 0x73, 0x22, 0x38, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x70, 0x73, 0x22, 0x76, 0x0a,
	0x16, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5c, 0x0a, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x6e,
	0x64, 0x72, 0x6f, 0x69, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x62, 0x0a, 0x12, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x73,
	0x68, 0x61, 0x31, 0x5f, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x68, 0x61, 0x31, 0x46, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x12, 0x49, 0x6f, 0x73,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x2c, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x22, 0x44, 0x0a,
	0x09, 0x41, 0x70, 0x69, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x73, 0x42, 0xb5, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76,
	0x32, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x70, 0x62,
	0x3b, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x70, 0x62, 0xaa, 0x02, 0x17, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x73, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x5c, 0x56, 0x32, 0xea, 0x02,
	0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_api_apikeys_v2_resources_proto_rawDescOnce sync.Once
	file_google_api_apikeys_v2_resources_proto_rawDescData = file_google_api_apikeys_v2_resources_proto_rawDesc
)

func file_google_api_apikeys_v2_resources_proto_rawDescGZIP() []byte {
	file_google_api_apikeys_v2_resources_proto_rawDescOnce.Do(func() {
		file_google_api_apikeys_v2_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_apikeys_v2_resources_proto_rawDescData)
	})
	return file_google_api_apikeys_v2_resources_proto_rawDescData
}

var file_google_api_apikeys_v2_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_google_api_apikeys_v2_resources_proto_goTypes = []interface{}{
	(*Key)(nil),                    // 0: google.api.apikeys.v2.Key
	(*Restrictions)(nil),           // 1: google.api.apikeys.v2.Restrictions
	(*BrowserKeyRestrictions)(nil), // 2: google.api.apikeys.v2.BrowserKeyRestrictions
	(*ServerKeyRestrictions)(nil),  // 3: google.api.apikeys.v2.ServerKeyRestrictions
	(*AndroidKeyRestrictions)(nil), // 4: google.api.apikeys.v2.AndroidKeyRestrictions
	(*AndroidApplication)(nil),     // 5: google.api.apikeys.v2.AndroidApplication
	(*IosKeyRestrictions)(nil),     // 6: google.api.apikeys.v2.IosKeyRestrictions
	(*ApiTarget)(nil),              // 7: google.api.apikeys.v2.ApiTarget
	nil,                            // 8: google.api.apikeys.v2.Key.AnnotationsEntry
	(*timestamppb.Timestamp)(nil),  // 9: google.protobuf.Timestamp
}
var file_google_api_apikeys_v2_resources_proto_depIdxs = []int32{
	9,  // 0: google.api.apikeys.v2.Key.create_time:type_name -> google.protobuf.Timestamp
	9,  // 1: google.api.apikeys.v2.Key.update_time:type_name -> google.protobuf.Timestamp
	9,  // 2: google.api.apikeys.v2.Key.delete_time:type_name -> google.protobuf.Timestamp
	8,  // 3: google.api.apikeys.v2.Key.annotations:type_name -> google.api.apikeys.v2.Key.AnnotationsEntry
	1,  // 4: google.api.apikeys.v2.Key.restrictions:type_name -> google.api.apikeys.v2.Restrictions
	2,  // 5: google.api.apikeys.v2.Restrictions.browser_key_restrictions:type_name -> google.api.apikeys.v2.BrowserKeyRestrictions
	3,  // 6: google.api.apikeys.v2.Restrictions.server_key_restrictions:type_name -> google.api.apikeys.v2.ServerKeyRestrictions
	4,  // 7: google.api.apikeys.v2.Restrictions.android_key_restrictions:type_name -> google.api.apikeys.v2.AndroidKeyRestrictions
	6,  // 8: google.api.apikeys.v2.Restrictions.ios_key_restrictions:type_name -> google.api.apikeys.v2.IosKeyRestrictions
	7,  // 9: google.api.apikeys.v2.Restrictions.api_targets:type_name -> google.api.apikeys.v2.ApiTarget
	5,  // 10: google.api.apikeys.v2.AndroidKeyRestrictions.allowed_applications:type_name -> google.api.apikeys.v2.AndroidApplication
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_google_api_apikeys_v2_resources_proto_init() }
func file_google_api_apikeys_v2_resources_proto_init() {
	if File_google_api_apikeys_v2_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_apikeys_v2_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_google_api_apikeys_v2_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restrictions); i {
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
		file_google_api_apikeys_v2_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrowserKeyRestrictions); i {
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
		file_google_api_apikeys_v2_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerKeyRestrictions); i {
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
		file_google_api_apikeys_v2_resources_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AndroidKeyRestrictions); i {
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
		file_google_api_apikeys_v2_resources_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AndroidApplication); i {
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
		file_google_api_apikeys_v2_resources_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IosKeyRestrictions); i {
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
		file_google_api_apikeys_v2_resources_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiTarget); i {
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
	file_google_api_apikeys_v2_resources_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Restrictions_BrowserKeyRestrictions)(nil),
		(*Restrictions_ServerKeyRestrictions)(nil),
		(*Restrictions_AndroidKeyRestrictions)(nil),
		(*Restrictions_IosKeyRestrictions)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_api_apikeys_v2_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_apikeys_v2_resources_proto_goTypes,
		DependencyIndexes: file_google_api_apikeys_v2_resources_proto_depIdxs,
		MessageInfos:      file_google_api_apikeys_v2_resources_proto_msgTypes,
	}.Build()
	File_google_api_apikeys_v2_resources_proto = out.File
	file_google_api_apikeys_v2_resources_proto_rawDesc = nil
	file_google_api_apikeys_v2_resources_proto_goTypes = nil
	file_google_api_apikeys_v2_resources_proto_depIdxs = nil
}