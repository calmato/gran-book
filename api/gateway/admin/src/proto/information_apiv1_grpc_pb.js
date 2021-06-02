// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_information_apiv1_pb = require('../proto/information_apiv1_pb.js');

function serialize_proto_CreateNotificationRequest(arg) {
  if (!(arg instanceof proto_information_apiv1_pb.CreateNotificationRequest)) {
    throw new Error('Expected argument of type proto.CreateNotificationRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_CreateNotificationRequest(buffer_arg) {
  return proto_information_apiv1_pb.CreateNotificationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_DeleteNotificationRequest(arg) {
  if (!(arg instanceof proto_information_apiv1_pb.DeleteNotificationRequest)) {
    throw new Error('Expected argument of type proto.DeleteNotificationRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_DeleteNotificationRequest(buffer_arg) {
  return proto_information_apiv1_pb.DeleteNotificationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_EmptyNotification(arg) {
  if (!(arg instanceof proto_information_apiv1_pb.EmptyNotification)) {
    throw new Error('Expected argument of type proto.EmptyNotification');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_EmptyNotification(buffer_arg) {
  return proto_information_apiv1_pb.EmptyNotification.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_GetNotificationRequest(arg) {
  if (!(arg instanceof proto_information_apiv1_pb.GetNotificationRequest)) {
    throw new Error('Expected argument of type proto.GetNotificationRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_GetNotificationRequest(buffer_arg) {
  return proto_information_apiv1_pb.GetNotificationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_NotificationResponse(arg) {
  if (!(arg instanceof proto_information_apiv1_pb.NotificationResponse)) {
    throw new Error('Expected argument of type proto.NotificationResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_NotificationResponse(buffer_arg) {
  return proto_information_apiv1_pb.NotificationResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateNotificationRequest(arg) {
  if (!(arg instanceof proto_information_apiv1_pb.UpdateNotificationRequest)) {
    throw new Error('Expected argument of type proto.UpdateNotificationRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UpdateNotificationRequest(buffer_arg) {
  return proto_information_apiv1_pb.UpdateNotificationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var NotificationServiceService = exports.NotificationServiceService = {
  listNotification: {
    path: '/proto.NotificationService/ListNotification',
    requestStream: false,
    responseStream: false,
    requestType: proto_information_apiv1_pb.EmptyNotification,
    responseType: proto_information_apiv1_pb.NotificationResponse,
    requestSerialize: serialize_proto_EmptyNotification,
    requestDeserialize: deserialize_proto_EmptyNotification,
    responseSerialize: serialize_proto_NotificationResponse,
    responseDeserialize: deserialize_proto_NotificationResponse,
  },
  getNotification: {
    path: '/proto.NotificationService/GetNotification',
    requestStream: false,
    responseStream: false,
    requestType: proto_information_apiv1_pb.GetNotificationRequest,
    responseType: proto_information_apiv1_pb.NotificationResponse,
    requestSerialize: serialize_proto_GetNotificationRequest,
    requestDeserialize: deserialize_proto_GetNotificationRequest,
    responseSerialize: serialize_proto_NotificationResponse,
    responseDeserialize: deserialize_proto_NotificationResponse,
  },
  createNotification: {
    path: '/proto.NotificationService/CreateNotification',
    requestStream: false,
    responseStream: false,
    requestType: proto_information_apiv1_pb.CreateNotificationRequest,
    responseType: proto_information_apiv1_pb.NotificationResponse,
    requestSerialize: serialize_proto_CreateNotificationRequest,
    requestDeserialize: deserialize_proto_CreateNotificationRequest,
    responseSerialize: serialize_proto_NotificationResponse,
    responseDeserialize: deserialize_proto_NotificationResponse,
  },
  updateNotification: {
    path: '/proto.NotificationService/UpdateNotification',
    requestStream: false,
    responseStream: false,
    requestType: proto_information_apiv1_pb.UpdateNotificationRequest,
    responseType: proto_information_apiv1_pb.NotificationResponse,
    requestSerialize: serialize_proto_UpdateNotificationRequest,
    requestDeserialize: deserialize_proto_UpdateNotificationRequest,
    responseSerialize: serialize_proto_NotificationResponse,
    responseDeserialize: deserialize_proto_NotificationResponse,
  },
  deleteNotification: {
    path: '/proto.NotificationService/DeleteNotification',
    requestStream: false,
    responseStream: false,
    requestType: proto_information_apiv1_pb.DeleteNotificationRequest,
    responseType: proto_information_apiv1_pb.NotificationResponse,
    requestSerialize: serialize_proto_DeleteNotificationRequest,
    requestDeserialize: deserialize_proto_DeleteNotificationRequest,
    responseSerialize: serialize_proto_NotificationResponse,
    responseDeserialize: deserialize_proto_NotificationResponse,
  },
};

exports.NotificationServiceClient = grpc.makeGenericClientConstructor(NotificationServiceService);
