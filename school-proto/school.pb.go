// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: school.proto

package school_proto

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

// Запрос на авторизацию
type SchoolLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// email от edu.21-school.ru формата nickname@student.21-school.ru
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// Пароль пользователя от edu.21-school.ru
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SchoolLoginRequest) Reset() {
	*x = SchoolLoginRequest{}
	mi := &file_school_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SchoolLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchoolLoginRequest) ProtoMessage() {}

func (x *SchoolLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchoolLoginRequest.ProtoReflect.Descriptor instead.
func (*SchoolLoginRequest) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{0}
}

func (x *SchoolLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SchoolLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// Ответ на запрос авторизации
type SchoolLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Авторизационный токен JWT
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *SchoolLoginResponse) Reset() {
	*x = SchoolLoginResponse{}
	mi := &file_school_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SchoolLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchoolLoginResponse) ProtoMessage() {}

func (x *SchoolLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchoolLoginResponse.ProtoReflect.Descriptor instead.
func (*SchoolLoginResponse) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{1}
}

func (x *SchoolLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// Запрос на получение всех кампусов
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_school_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{2}
}

// Ответ на запрос с данными кампуса
type CampusesOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Uuid кампуса
	CampusUuid string `protobuf:"bytes,1,opt,name=campus_uuid,json=campusUuid,proto3" json:"campus_uuid,omitempty"`
	// Сокращенное название кампуса
	ShortName string `protobuf:"bytes,2,opt,name=shortName,proto3" json:"shortName,omitempty"`
	// Полное название кампуса
	FullName string `protobuf:"bytes,3,opt,name=fullName,proto3" json:"fullName,omitempty"`
}

func (x *CampusesOut) Reset() {
	*x = CampusesOut{}
	mi := &file_school_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CampusesOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampusesOut) ProtoMessage() {}

func (x *CampusesOut) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampusesOut.ProtoReflect.Descriptor instead.
func (*CampusesOut) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{3}
}

func (x *CampusesOut) GetCampusUuid() string {
	if x != nil {
		return x.CampusUuid
	}
	return ""
}

func (x *CampusesOut) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *CampusesOut) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

// Запрос на получение всех трайбов кампуса
type CampusUuidIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Uuid кампуса, который призходит от community сервиса
	CampusUuid string `protobuf:"bytes,1,opt,name=campus_uuid,json=campusUuid,proto3" json:"campus_uuid,omitempty"`
}

func (x *CampusUuidIn) Reset() {
	*x = CampusUuidIn{}
	mi := &file_school_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CampusUuidIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampusUuidIn) ProtoMessage() {}

func (x *CampusUuidIn) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampusUuidIn.ProtoReflect.Descriptor instead.
func (*CampusUuidIn) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{4}
}

func (x *CampusUuidIn) GetCampusUuid() string {
	if x != nil {
		return x.CampusUuid
	}
	return ""
}

// Сообщение для выходных данных
type Tribe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id  тайба
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Название трайба
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tribe) Reset() {
	*x = Tribe{}
	mi := &file_school_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tribe) ProtoMessage() {}

func (x *Tribe) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tribe.ProtoReflect.Descriptor instead.
func (*Tribe) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{5}
}

func (x *Tribe) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tribe) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Ответ на запрос на получение всех трайбов кампуса
type TribesOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Список трайбов
	Tribes []*Tribe `protobuf:"bytes,1,rep,name=tribes,proto3" json:"tribes,omitempty"`
}

func (x *TribesOut) Reset() {
	*x = TribesOut{}
	mi := &file_school_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TribesOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TribesOut) ProtoMessage() {}

func (x *TribesOut) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TribesOut.ProtoReflect.Descriptor instead.
func (*TribesOut) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{6}
}

func (x *TribesOut) GetTribes() []*Tribe {
	if x != nil {
		return x.Tribes
	}
	return nil
}

var File_school_proto protoreflect.FileDescriptor

var file_school_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46,
	0x0a, 0x12, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2b, 0x0a, 0x13, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x68, 0x0a, 0x0b,
	0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x61, 0x6d, 0x70, 0x75, 0x73, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x0c, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x55, 0x75, 0x69, 0x64, 0x49, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6d,
	0x70, 0x75, 0x73, 0x55, 0x75, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x05, 0x54, 0x72, 0x69, 0x62, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x09, 0x54, 0x72, 0x69, 0x62, 0x65, 0x73, 0x4f, 0x75,
	0x74, 0x12, 0x1e, 0x0a, 0x06, 0x74, 0x72, 0x69, 0x62, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x06, 0x2e, 0x54, 0x72, 0x69, 0x62, 0x65, 0x52, 0x06, 0x74, 0x72, 0x69, 0x62, 0x65,
	0x73, 0x32, 0xa2, 0x01, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x53,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x65, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0c, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x62, 0x65, 0x73, 0x42, 0x79, 0x43,
	0x61, 0x6d, 0x70, 0x75, 0x73, 0x55, 0x75, 0x69, 0x64, 0x12, 0x0d, 0x2e, 0x43, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x55, 0x75, 0x69, 0x64, 0x49, 0x6e, 0x1a, 0x0a, 0x2e, 0x54, 0x72, 0x69, 0x62, 0x65,
	0x73, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_school_proto_rawDescOnce sync.Once
	file_school_proto_rawDescData = file_school_proto_rawDesc
)

func file_school_proto_rawDescGZIP() []byte {
	file_school_proto_rawDescOnce.Do(func() {
		file_school_proto_rawDescData = protoimpl.X.CompressGZIP(file_school_proto_rawDescData)
	})
	return file_school_proto_rawDescData
}

var file_school_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_school_proto_goTypes = []any{
	(*SchoolLoginRequest)(nil),  // 0: SchoolLoginRequest
	(*SchoolLoginResponse)(nil), // 1: SchoolLoginResponse
	(*Empty)(nil),               // 2: Empty
	(*CampusesOut)(nil),         // 3: CampusesOut
	(*CampusUuidIn)(nil),        // 4: CampusUuidIn
	(*Tribe)(nil),               // 5: Tribe
	(*TribesOut)(nil),           // 6: TribesOut
}
var file_school_proto_depIdxs = []int32{
	5, // 0: TribesOut.tribes:type_name -> Tribe
	0, // 1: SchoolService.Login:input_type -> SchoolLoginRequest
	2, // 2: SchoolService.GetCampuses:input_type -> Empty
	4, // 3: SchoolService.GetTribesByCampusUuid:input_type -> CampusUuidIn
	1, // 4: SchoolService.Login:output_type -> SchoolLoginResponse
	3, // 5: SchoolService.GetCampuses:output_type -> CampusesOut
	6, // 6: SchoolService.GetTribesByCampusUuid:output_type -> TribesOut
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_school_proto_init() }
func file_school_proto_init() {
	if File_school_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_school_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_school_proto_goTypes,
		DependencyIndexes: file_school_proto_depIdxs,
		MessageInfos:      file_school_proto_msgTypes,
	}.Build()
	File_school_proto = out.File
	file_school_proto_rawDesc = nil
	file_school_proto_goTypes = nil
	file_school_proto_depIdxs = nil
}
