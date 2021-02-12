// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_user_apiv1_pb = require('../proto/user_apiv1_pb.js');

function serialize_proto_AuthResponse(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.AuthResponse)) {
    throw new Error('Expected argument of type proto.AuthResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_AuthResponse(buffer_arg) {
  return proto_user_apiv1_pb.AuthResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_CreateAuthRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.CreateAuthRequest)) {
    throw new Error('Expected argument of type proto.CreateAuthRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_CreateAuthRequest(buffer_arg) {
  return proto_user_apiv1_pb.CreateAuthRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_EmptyUser(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.EmptyUser)) {
    throw new Error('Expected argument of type proto.EmptyUser');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_EmptyUser(buffer_arg) {
  return proto_user_apiv1_pb.EmptyUser.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateAuthAddressRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UpdateAuthAddressRequest)) {
    throw new Error('Expected argument of type proto.UpdateAuthAddressRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UpdateAuthAddressRequest(buffer_arg) {
  return proto_user_apiv1_pb.UpdateAuthAddressRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateAuthEmailRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UpdateAuthEmailRequest)) {
    throw new Error('Expected argument of type proto.UpdateAuthEmailRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UpdateAuthEmailRequest(buffer_arg) {
  return proto_user_apiv1_pb.UpdateAuthEmailRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateAuthPasswordRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UpdateAuthPasswordRequest)) {
    throw new Error('Expected argument of type proto.UpdateAuthPasswordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UpdateAuthPasswordRequest(buffer_arg) {
  return proto_user_apiv1_pb.UpdateAuthPasswordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateAuthProfileRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UpdateAuthProfileRequest)) {
    throw new Error('Expected argument of type proto.UpdateAuthProfileRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UpdateAuthProfileRequest(buffer_arg) {
  return proto_user_apiv1_pb.UpdateAuthProfileRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var AuthServiceService = exports.AuthServiceService = {
  getAuth: {
    path: '/proto.AuthService/GetAuth',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.EmptyUser,
    responseType: proto_user_apiv1_pb.AuthResponse,
    requestSerialize: serialize_proto_EmptyUser,
    requestDeserialize: deserialize_proto_EmptyUser,
    responseSerialize: serialize_proto_AuthResponse,
    responseDeserialize: deserialize_proto_AuthResponse,
  },
  createAuth: {
    path: '/proto.AuthService/CreateAuth',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.CreateAuthRequest,
    responseType: proto_user_apiv1_pb.AuthResponse,
    requestSerialize: serialize_proto_CreateAuthRequest,
    requestDeserialize: deserialize_proto_CreateAuthRequest,
    responseSerialize: serialize_proto_AuthResponse,
    responseDeserialize: deserialize_proto_AuthResponse,
  },
  updateAuthEmail: {
    path: '/proto.AuthService/UpdateAuthEmail',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.UpdateAuthEmailRequest,
    responseType: proto_user_apiv1_pb.AuthResponse,
    requestSerialize: serialize_proto_UpdateAuthEmailRequest,
    requestDeserialize: deserialize_proto_UpdateAuthEmailRequest,
    responseSerialize: serialize_proto_AuthResponse,
    responseDeserialize: deserialize_proto_AuthResponse,
  },
  updateAuthPassword: {
    path: '/proto.AuthService/UpdateAuthPassword',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.UpdateAuthPasswordRequest,
    responseType: proto_user_apiv1_pb.AuthResponse,
    requestSerialize: serialize_proto_UpdateAuthPasswordRequest,
    requestDeserialize: deserialize_proto_UpdateAuthPasswordRequest,
    responseSerialize: serialize_proto_AuthResponse,
    responseDeserialize: deserialize_proto_AuthResponse,
  },
  updateAuthProfile: {
    path: '/proto.AuthService/UpdateAuthProfile',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.UpdateAuthProfileRequest,
    responseType: proto_user_apiv1_pb.AuthResponse,
    requestSerialize: serialize_proto_UpdateAuthProfileRequest,
    requestDeserialize: deserialize_proto_UpdateAuthProfileRequest,
    responseSerialize: serialize_proto_AuthResponse,
    responseDeserialize: deserialize_proto_AuthResponse,
  },
  updateAuthAddress: {
    path: '/proto.AuthService/UpdateAuthAddress',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.UpdateAuthAddressRequest,
    responseType: proto_user_apiv1_pb.AuthResponse,
    requestSerialize: serialize_proto_UpdateAuthAddressRequest,
    requestDeserialize: deserialize_proto_UpdateAuthAddressRequest,
    responseSerialize: serialize_proto_AuthResponse,
    responseDeserialize: deserialize_proto_AuthResponse,
  },
};

exports.AuthServiceClient = grpc.makeGenericClientConstructor(AuthServiceService);
