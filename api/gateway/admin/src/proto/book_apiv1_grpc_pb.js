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

function serialize_proto_BookshelfListResponse(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.BookshelfListResponse)) {
    throw new Error('Expected argument of type proto.BookshelfListResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_BookshelfListResponse(buffer_arg) {
  return proto_book_apiv1_pb.BookshelfListResponse.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_proto_CreateBookRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.CreateBookRequest)) {
    throw new Error('Expected argument of type proto.CreateBookRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_CreateBookRequest(buffer_arg) {
  return proto_book_apiv1_pb.CreateBookRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_DeleteBookRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.DeleteBookRequest)) {
    throw new Error('Expected argument of type proto.DeleteBookRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_DeleteBookRequest(buffer_arg) {
  return proto_book_apiv1_pb.DeleteBookRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_DeleteBookshelfRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.DeleteBookshelfRequest)) {
    throw new Error('Expected argument of type proto.DeleteBookshelfRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_DeleteBookshelfRequest(buffer_arg) {
  return proto_book_apiv1_pb.DeleteBookshelfRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_EmptyBook(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.EmptyBook)) {
    throw new Error('Expected argument of type proto.EmptyBook');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_EmptyBook(buffer_arg) {
  return proto_book_apiv1_pb.EmptyBook.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_GetBookRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.GetBookRequest)) {
    throw new Error('Expected argument of type proto.GetBookRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_GetBookRequest(buffer_arg) {
  return proto_book_apiv1_pb.GetBookRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_GetBookshelfRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.GetBookshelfRequest)) {
    throw new Error('Expected argument of type proto.GetBookshelfRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_GetBookshelfRequest(buffer_arg) {
  return proto_book_apiv1_pb.GetBookshelfRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_ListBookshelfRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.ListBookshelfRequest)) {
    throw new Error('Expected argument of type proto.ListBookshelfRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_ListBookshelfRequest(buffer_arg) {
  return proto_book_apiv1_pb.ListBookshelfRequest.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_proto_UpdateBookRequest(arg) {
  if (!(arg instanceof proto_book_apiv1_pb.UpdateBookRequest)) {
    throw new Error('Expected argument of type proto.UpdateBookRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UpdateBookRequest(buffer_arg) {
  return proto_book_apiv1_pb.UpdateBookRequest.deserializeBinary(new Uint8Array(buffer_arg));
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
  listBookshelf: {
    path: '/proto.BookService/ListBookshelf',
    requestStream: false,
    responseStream: false,
    requestType: proto_book_apiv1_pb.ListBookshelfRequest,
    responseType: proto_book_apiv1_pb.BookshelfListResponse,
    requestSerialize: serialize_proto_ListBookshelfRequest,
    requestDeserialize: deserialize_proto_ListBookshelfRequest,
    responseSerialize: serialize_proto_BookshelfListResponse,
    responseDeserialize: deserialize_proto_BookshelfListResponse,
  },
  getBook: {
    path: '/proto.BookService/GetBook',
    requestStream: false,
    responseStream: false,
    requestType: proto_book_apiv1_pb.GetBookRequest,
    responseType: proto_book_apiv1_pb.BookResponse,
    requestSerialize: serialize_proto_GetBookRequest,
    requestDeserialize: deserialize_proto_GetBookRequest,
    responseSerialize: serialize_proto_BookResponse,
    responseDeserialize: deserialize_proto_BookResponse,
  },
  getBookshelf: {
    path: '/proto.BookService/GetBookshelf',
    requestStream: false,
    responseStream: false,
    requestType: proto_book_apiv1_pb.GetBookshelfRequest,
    responseType: proto_book_apiv1_pb.BookshelfResponse,
    requestSerialize: serialize_proto_GetBookshelfRequest,
    requestDeserialize: deserialize_proto_GetBookshelfRequest,
    responseSerialize: serialize_proto_BookshelfResponse,
    responseDeserialize: deserialize_proto_BookshelfResponse,
  },
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
  updateBook: {
    path: '/proto.BookService/UpdateBook',
    requestStream: false,
    responseStream: false,
    requestType: proto_book_apiv1_pb.UpdateBookRequest,
    responseType: proto_book_apiv1_pb.BookResponse,
    requestSerialize: serialize_proto_UpdateBookRequest,
    requestDeserialize: deserialize_proto_UpdateBookRequest,
    responseSerialize: serialize_proto_BookResponse,
    responseDeserialize: deserialize_proto_BookResponse,
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
  deleteBook: {
    path: '/proto.BookService/DeleteBook',
    requestStream: false,
    responseStream: false,
    requestType: proto_book_apiv1_pb.DeleteBookRequest,
    responseType: proto_book_apiv1_pb.EmptyBook,
    requestSerialize: serialize_proto_DeleteBookRequest,
    requestDeserialize: deserialize_proto_DeleteBookRequest,
    responseSerialize: serialize_proto_EmptyBook,
    responseDeserialize: deserialize_proto_EmptyBook,
  },
  deleteBookshelf: {
    path: '/proto.BookService/DeleteBookshelf',
    requestStream: false,
    responseStream: false,
    requestType: proto_book_apiv1_pb.DeleteBookshelfRequest,
    responseType: proto_book_apiv1_pb.EmptyBook,
    requestSerialize: serialize_proto_DeleteBookshelfRequest,
    requestDeserialize: deserialize_proto_DeleteBookshelfRequest,
    responseSerialize: serialize_proto_EmptyBook,
    responseDeserialize: deserialize_proto_EmptyBook,
  },
};

exports.BookServiceClient = grpc.makeGenericClientConstructor(BookServiceService);
