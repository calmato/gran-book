// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_book_apiv1_pb = require('../proto/book_apiv1_pb.js');

function serialize_proto_BookResponse(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.BookResponse)) {
    throw new Error('Expected argument of type proto.BookResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_BookResponse(buffer_arg) {
  return proto_book_apiv1_pb.BookResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_CreateBookRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.CreateBookRequest)) {
    throw new Error('Expected argument of type proto.CreateBookRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_CreateBookRequest(buffer_arg) {
  return proto_book_apiv1_pb.CreateBookRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var BookServiceService = exports.BookServiceService = {
  createBook: {
    path: '/proto.BookService/CreateBook',
    requestStream: false,
    responseStream: false,
    requestType: proto_book_apiv1_pb.CreateBookRequest,
    responseType: proto_book_apiv1_pb.BookResponse,
    requestSerialize: serialize_proto_CreateBookRequest,
    requestDeserialize: deserialize_proto_CreateBookRequest,
    responseSerialize: serialize_proto_BookResponse,
    responseDeserialize: deserialize_proto_BookResponse,
  },
};

exports.BookServiceClient = grpc.makeGenericClientConstructor(BookServiceService);
