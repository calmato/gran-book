// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_information_apiv1_pb = require('../proto/information_apiv1_pb.js');

function serialize_proto_CreateInquiryRequest(arg) {
  if (!(arg instanceof proto_information_apiv1_pb.CreateInquiryRequest)) {
    throw new Error('Expected argument of type proto.CreateInquiryRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_CreateInquiryRequest(buffer_arg) {
  return proto_information_apiv1_pb.CreateInquiryRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_HelloRequest(arg) {
  if (!(arg instanceof proto_information_apiv1_pb.HelloRequest)) {
    throw new Error('Expected argument of type proto.HelloRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_HelloRequest(buffer_arg) {
  return proto_information_apiv1_pb.HelloRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_HelloResponse(arg) {
  if (!(arg instanceof proto_information_apiv1_pb.HelloResponse)) {
    throw new Error('Expected argument of type proto.HelloResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_HelloResponse(buffer_arg) {
  return proto_information_apiv1_pb.HelloResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_InquiryResponse(arg) {
  if (!(arg instanceof proto_information_apiv1_pb.InquiryResponse)) {
    throw new Error('Expected argument of type proto.InquiryResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_InquiryResponse(buffer_arg) {
  return proto_information_apiv1_pb.InquiryResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var NotificationServiceService = exports.NotificationServiceService = {
  reply: {
    path: '/proto.NotificationService/Reply',
    requestStream: false,
    responseStream: false,
    requestType: proto_information_apiv1_pb.HelloRequest,
    responseType: proto_information_apiv1_pb.HelloResponse,
    requestSerialize: serialize_proto_HelloRequest,
    requestDeserialize: deserialize_proto_HelloRequest,
    responseSerialize: serialize_proto_HelloResponse,
    responseDeserialize: deserialize_proto_HelloResponse,
  },
};

exports.NotificationServiceClient = grpc.makeGenericClientConstructor(NotificationServiceService);
var InquiryServiceService = exports.InquiryServiceService = {
  createInquiry: {
    path: '/proto.InquiryService/CreateInquiry',
    requestStream: false,
    responseStream: false,
    requestType: proto_information_apiv1_pb.CreateInquiryRequest,
    responseType: proto_information_apiv1_pb.InquiryResponse,
    requestSerialize: serialize_proto_CreateInquiryRequest,
    requestDeserialize: deserialize_proto_CreateInquiryRequest,
    responseSerialize: serialize_proto_InquiryResponse,
    responseDeserialize: deserialize_proto_InquiryResponse,
  },
};

exports.InquiryServiceClient = grpc.makeGenericClientConstructor(InquiryServiceService);
