// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_book_apiv1_pb = require('../proto/book_apiv1_pb.js');

function serialize_proto_BookListResponse(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.BookListResponse)) {
    throw new Error('Expected argument of type proto.BookListResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_BookListResponse(buffer_arg) {
  return proto_book_apiv1_pb.BookListResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_BookResponse(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.BookResponse)) {
    throw new Error('Expected argument of type proto.BookResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_BookResponse(buffer_arg) {
  return proto_book_apiv1_pb.BookResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_CreateAndUpdateBooksRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.CreateAndUpdateBooksRequest)) {
    throw new Error('Expected argument of type proto.CreateAndUpdateBooksRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_CreateAndUpdateBooksRequest(buffer_arg) {
  return proto_book_apiv1_pb.CreateAndUpdateBooksRequest.deserializeBinary(new Uint8Array(buffer_arg));
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
  createAndUpdateBooks: {
    path: '/proto.BookService/CreateAndUpdateBooks',
    requestStream: false,
    responseStream: false,
    requestType: proto_book_apiv1_pb.CreateAndUpdateBooksRequest,
    responseType: proto_book_apiv1_pb.BookListResponse,
    requestSerialize: serialize_proto_CreateAndUpdateBooksRequest,
    requestDeserialize: deserialize_proto_CreateAndUpdateBooksRequest,
    responseSerialize: serialize_proto_BookListResponse,
    responseDeserialize: deserialize_proto_BookListResponse,
  },
};

exports.BookServiceClient = grpc.makeGenericClientConstructor(BookServiceService);
