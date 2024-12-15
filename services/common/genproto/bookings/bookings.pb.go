// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: bookings.proto

package bookings

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ride struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RideId      int32  `protobuf:"varint,1,opt,name=ride_id,json=rideId,proto3" json:"ride_id,omitempty"`
	Source      string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Destination string `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	Distance    int32  `protobuf:"varint,4,opt,name=distance,proto3" json:"distance,omitempty"`
	Cost        int32  `protobuf:"varint,5,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *Ride) Reset() {
	*x = Ride{}
	mi := &file_bookings_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ride) ProtoMessage() {}

func (x *Ride) ProtoReflect() protoreflect.Message {
	mi := &file_bookings_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ride.ProtoReflect.Descriptor instead.
func (*Ride) Descriptor() ([]byte, []int) {
	return file_bookings_proto_rawDescGZIP(), []int{0}
}

func (x *Ride) GetRideId() int32 {
	if x != nil {
		return x.RideId
	}
	return 0
}

func (x *Ride) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Ride) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *Ride) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *Ride) GetCost() int32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId int32                  `protobuf:"varint,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	RideId    int32                  `protobuf:"varint,2,opt,name=ride_id,json=rideId,proto3" json:"ride_id,omitempty"`
	UserId    int32                  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RideTime  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=ride_time,json=rideTime,proto3" json:"ride_time,omitempty"`
}

// Error implements error.
func (x *Booking) Error() string {
	panic("unimplemented")
}

func (x *Booking) Reset() {
	*x = Booking{}
	mi := &file_bookings_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_bookings_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_bookings_proto_rawDescGZIP(), []int{1}
}

func (x *Booking) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *Booking) GetRideId() int32 {
	if x != nil {
		return x.RideId
	}
	return 0
}

func (x *Booking) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Booking) GetRideTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RideTime
	}
	return nil
}

type CreateBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Ride   *Ride `protobuf:"bytes,2,opt,name=ride,proto3" json:"ride,omitempty"`
}

func (x *CreateBookingRequest) Reset() {
	*x = CreateBookingRequest{}
	mi := &file_bookings_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookingRequest) ProtoMessage() {}

func (x *CreateBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookings_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookingRequest.ProtoReflect.Descriptor instead.
func (*CreateBookingRequest) Descriptor() ([]byte, []int) {
	return file_bookings_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBookingRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateBookingRequest) GetRide() *Ride {
	if x != nil {
		return x.Ride
	}
	return nil
}

type CreateBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Booking *Booking `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
}

func (x *CreateBookingResponse) Reset() {
	*x = CreateBookingResponse{}
	mi := &file_bookings_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookingResponse) ProtoMessage() {}

func (x *CreateBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookings_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookingResponse.ProtoReflect.Descriptor instead.
func (*CreateBookingResponse) Descriptor() ([]byte, []int) {
	return file_bookings_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBookingResponse) GetBooking() *Booking {
	if x != nil {
		return x.Booking
	}
	return nil
}

type GetBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId int32 `protobuf:"varint,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
}

func (x *GetBookingRequest) Reset() {
	*x = GetBookingRequest{}
	mi := &file_bookings_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookingRequest) ProtoMessage() {}

func (x *GetBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookings_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookingRequest.ProtoReflect.Descriptor instead.
func (*GetBookingRequest) Descriptor() ([]byte, []int) {
	return file_bookings_proto_rawDescGZIP(), []int{4}
}

func (x *GetBookingRequest) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

type GetBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Booking *Booking `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
}

func (x *GetBookingResponse) Reset() {
	*x = GetBookingResponse{}
	mi := &file_bookings_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookingResponse) ProtoMessage() {}

func (x *GetBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookings_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookingResponse.ProtoReflect.Descriptor instead.
func (*GetBookingResponse) Descriptor() ([]byte, []int) {
	return file_bookings_proto_rawDescGZIP(), []int{5}
}

func (x *GetBookingResponse) GetBooking() *Booking {
	if x != nil {
		return x.Booking
	}
	return nil
}

type RideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RideId int32 `protobuf:"varint,1,opt,name=ride_id,json=rideId,proto3" json:"ride_id,omitempty"`
	Ride   *Ride `protobuf:"bytes,2,opt,name=ride,proto3" json:"ride,omitempty"`
}

func (x *RideRequest) Reset() {
	*x = RideRequest{}
	mi := &file_bookings_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RideRequest) ProtoMessage() {}

func (x *RideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookings_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RideRequest.ProtoReflect.Descriptor instead.
func (*RideRequest) Descriptor() ([]byte, []int) {
	return file_bookings_proto_rawDescGZIP(), []int{6}
}

func (x *RideRequest) GetRideId() int32 {
	if x != nil {
		return x.RideId
	}
	return 0
}

func (x *RideRequest) GetRide() *Ride {
	if x != nil {
		return x.Ride
	}
	return nil
}

type RideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RideResponse) Reset() {
	*x = RideResponse{}
	mi := &file_bookings_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RideResponse) ProtoMessage() {}

func (x *RideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookings_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RideResponse.ProtoReflect.Descriptor instead.
func (*RideResponse) Descriptor() ([]byte, []int) {
	return file_bookings_proto_rawDescGZIP(), []int{7}
}

func (x *RideResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_bookings_proto protoreflect.FileDescriptor

var file_bookings_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x89, 0x01, 0x0a, 0x04, 0x52, 0x69, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x69,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x69, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0x93, 0x01,
	0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x69, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x69, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x69,
	0x64, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x72, 0x69, 0x64, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x72, 0x69, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x52, 0x04, 0x72, 0x69, 0x64, 0x65, 0x22,
	0x3b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x32, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64,
	0x22, 0x38, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x41, 0x0a, 0x0b, 0x52, 0x69,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x69, 0x64,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x69, 0x64, 0x65,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x72, 0x69, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x52, 0x04, 0x72, 0x69, 0x64, 0x65, 0x22, 0x28, 0x0a,
	0x0c, 0x52, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb8, 0x01, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x69, 0x64, 0x65, 0x12, 0x0c, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x61, 0x62, 0x65, 0x65, 0x68, 0x61, 0x2d, 0x4d, 0x75, 0x64, 0x61, 0x73, 0x73, 0x69,
	0x72, 0x2f, 0x52, 0x69, 0x64, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bookings_proto_rawDescOnce sync.Once
	file_bookings_proto_rawDescData = file_bookings_proto_rawDesc
)

func file_bookings_proto_rawDescGZIP() []byte {
	file_bookings_proto_rawDescOnce.Do(func() {
		file_bookings_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookings_proto_rawDescData)
	})
	return file_bookings_proto_rawDescData
}

var file_bookings_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_bookings_proto_goTypes = []any{
	(*Ride)(nil),                  // 0: Ride
	(*Booking)(nil),               // 1: Booking
	(*CreateBookingRequest)(nil),  // 2: CreateBookingRequest
	(*CreateBookingResponse)(nil), // 3: CreateBookingResponse
	(*GetBookingRequest)(nil),     // 4: GetBookingRequest
	(*GetBookingResponse)(nil),    // 5: GetBookingResponse
	(*RideRequest)(nil),           // 6: RideRequest
	(*RideResponse)(nil),          // 7: RideResponse
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_bookings_proto_depIdxs = []int32{
	8, // 0: Booking.ride_time:type_name -> google.protobuf.Timestamp
	0, // 1: CreateBookingRequest.ride:type_name -> Ride
	1, // 2: CreateBookingResponse.booking:type_name -> Booking
	1, // 3: GetBookingResponse.booking:type_name -> Booking
	0, // 4: RideRequest.ride:type_name -> Ride
	4, // 5: BookingService.GetBooking:input_type -> GetBookingRequest
	2, // 6: BookingService.CreateBooking:input_type -> CreateBookingRequest
	6, // 7: BookingService.UpdateRide:input_type -> RideRequest
	5, // 8: BookingService.GetBooking:output_type -> GetBookingResponse
	3, // 9: BookingService.CreateBooking:output_type -> CreateBookingResponse
	7, // 10: BookingService.UpdateRide:output_type -> RideResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_bookings_proto_init() }
func file_bookings_proto_init() {
	if File_bookings_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bookings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bookings_proto_goTypes,
		DependencyIndexes: file_bookings_proto_depIdxs,
		MessageInfos:      file_bookings_proto_msgTypes,
	}.Build()
	File_bookings_proto = out.File
	file_bookings_proto_rawDesc = nil
	file_bookings_proto_goTypes = nil
	file_bookings_proto_depIdxs = nil
}
