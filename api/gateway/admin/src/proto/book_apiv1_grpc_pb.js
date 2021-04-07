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

function serialize_proto_BookshelfResponse(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.BookshelfResponse)) {
    throw new Error('Expected argument of type proto.BookshelfResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_BookshelfResponse(buffer_arg) {
  return proto_book_apiv1_pb.BookshelfResponse.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_proto_ReadBookshelfRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.ReadBookshelfRequest)) {
    throw new Error('Expected argument of type proto.ReadBookshelfRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_ReadBookshelfRequest(buffer_arg) {
  return proto_book_apiv1_pb.ReadBookshelfRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_ReadingBookshelfRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.ReadingBookshelfRequest)) {
    throw new Error('Expected argument of type proto.ReadingBookshelfRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_ReadingBookshelfRequest(buffer_arg) {
  return proto_book_apiv1_pb.ReadingBookshelfRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_ReleaseBookshelfRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.ReleaseBookshelfRequest)) {
    throw new Error('Expected argument of type proto.ReleaseBookshelfRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_ReleaseBookshelfRequest(buffer_arg) {
  return proto_book_apiv1_pb.ReleaseBookshelfRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_StackBookshelfRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.StackBookshelfRequest)) {
    throw new Error('Expected argument of type proto.StackBookshelfRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_StackBookshelfRequest(buffer_arg) {
  return proto_book_apiv1_pb.StackBookshelfRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_WantBookshelfRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.WantBookshelfRequest)) {
    throw new Error('Expected argument of type proto.WantBookshelfRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_WantBookshelfRequest(buffer_arg) {
  return proto_book_apiv1_pb.WantBookshelfRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var BookServiceService = exports.BookServiceService = {
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
  readBookshelf: {
    path: '/proto.BookService/ReadBookshelf',
    requestStream: false,
    responseStream: false,
    requestType: proto_book_apiv1_pb.ReadBookshelfRequest,
    responseType: proto_book_apiv1_pb.BookshelfResponse,
    requestSerialize: serialize_proto_ReadBookshelfRequest,
    requestDeserialize: deserialize_proto_ReadBookshelfRequest,
    responseSerialize: serialize_proto_BookshelfResponse,
    responseDeserialize: deserialize_proto_BookshelfResponse,
  },
  readingBookshelf: {
    path: '/proto.BookService/ReadingBookshelf',
    requestStream: false,
    responseStream: false,
    requestType: proto_book_apiv1_pb.ReadingBookshelfRequest,
    responseType: proto_book_apiv1_pb.BookshelfResponse,
    requestSerialize: serialize_proto_ReadingBookshelfRequest,
    requestDeserialize: deserialize_proto_ReadingBookshelfRequest,
    responseSerialize: serialize_proto_BookshelfResponse,
    responseDeserialize: deserialize_proto_BookshelfResponse,
  },
  stackBookshelf: {
    path: '/proto.BookService/StackBookshelf',
    requestStream: false,
    responseStream: false,
    requestType: proto_book_apiv1_pb.StackBookshelfRequest,
    responseType: proto_book_apiv1_pb.BookshelfResponse,
    requestSerialize: serialize_proto_StackBookshelfRequest,
    requestDeserialize: deserialize_proto_StackBookshelfRequest,
    responseSerialize: serialize_proto_BookshelfResponse,
    responseDeserialize: deserialize_proto_BookshelfResponse,
  },
  wantBookshelf: {
    path: '/proto.BookService/WantBookshelf',
    requestStream: false,
    responseStream: false,
    requestType: proto_book_apiv1_pb.WantBookshelfRequest,
    responseType: proto_book_apiv1_pb.BookshelfResponse,
    requestSerialize: serialize_proto_WantBookshelfRequest,
    requestDeserialize: deserialize_proto_WantBookshelfRequest,
    responseSerialize: serialize_proto_BookshelfResponse,
    responseDeserialize: deserialize_proto_BookshelfResponse,
  },
  releaseBookshelf: {
    path: '/proto.BookService/ReleaseBookshelf',
    requestStream: false,
    responseStream: false,
    requestType: proto_book_apiv1_pb.ReleaseBookshelfRequest,
    responseType: proto_book_apiv1_pb.BookshelfResponse,
    requestSerialize: serialize_proto_ReleaseBookshelfRequest,
    requestDeserialize: deserialize_proto_ReleaseBookshelfRequest,
    responseSerialize: serialize_proto_BookshelfResponse,
    responseDeserialize: deserialize_proto_BookshelfResponse,
  },
};

exports.BookServiceClient = grpc.makeGenericClientConstructor(BookServiceService);
