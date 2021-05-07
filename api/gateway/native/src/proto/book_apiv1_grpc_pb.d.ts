// package: proto
// file: proto/book_apiv1.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as proto_book_apiv1_pb from "../proto/book_apiv1_pb";

interface IBookServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listBookByBookIds: IBookServiceService_IListBookByBookIds;
    listBookshelf: IBookServiceService_IListBookshelf;
    listBookReview: IBookServiceService_IListBookReview;
    listUserReview: IBookServiceService_IListUserReview;
    getBook: IBookServiceService_IGetBook;
    getBookByIsbn: IBookServiceService_IGetBookByIsbn;
    getBookshelf: IBookServiceService_IGetBookshelf;
    getReview: IBookServiceService_IGetReview;
    createBook: IBookServiceService_ICreateBook;
    updateBook: IBookServiceService_IUpdateBook;
    readBookshelf: IBookServiceService_IReadBookshelf;
    readingBookshelf: IBookServiceService_IReadingBookshelf;
    stackBookshelf: IBookServiceService_IStackBookshelf;
    wantBookshelf: IBookServiceService_IWantBookshelf;
    releaseBookshelf: IBookServiceService_IReleaseBookshelf;
    deleteBook: IBookServiceService_IDeleteBook;
    deleteBookshelf: IBookServiceService_IDeleteBookshelf;
}

interface IBookServiceService_IListBookByBookIds extends grpc.MethodDefinition<proto_book_apiv1_pb.ListBookByBookIdsRequest, proto_book_apiv1_pb.BookListResponse> {
    path: "/proto.BookService/ListBookByBookIds";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.ListBookByBookIdsRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.ListBookByBookIdsRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookListResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookListResponse>;
}
interface IBookServiceService_IListBookshelf extends grpc.MethodDefinition<proto_book_apiv1_pb.ListBookshelfRequest, proto_book_apiv1_pb.BookshelfListResponse> {
    path: "/proto.BookService/ListBookshelf";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.ListBookshelfRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.ListBookshelfRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookshelfListResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookshelfListResponse>;
}
interface IBookServiceService_IListBookReview extends grpc.MethodDefinition<proto_book_apiv1_pb.ListBookReviewRequest, proto_book_apiv1_pb.ReviewListResponse> {
    path: "/proto.BookService/ListBookReview";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.ListBookReviewRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.ListBookReviewRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.ReviewListResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.ReviewListResponse>;
}
interface IBookServiceService_IListUserReview extends grpc.MethodDefinition<proto_book_apiv1_pb.ListUserReviewRequest, proto_book_apiv1_pb.ReviewListResponse> {
    path: "/proto.BookService/ListUserReview";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.ListUserReviewRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.ListUserReviewRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.ReviewListResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.ReviewListResponse>;
}
interface IBookServiceService_IGetBook extends grpc.MethodDefinition<proto_book_apiv1_pb.GetBookRequest, proto_book_apiv1_pb.BookResponse> {
    path: "/proto.BookService/GetBook";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.GetBookRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.GetBookRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookResponse>;
}
interface IBookServiceService_IGetBookByIsbn extends grpc.MethodDefinition<proto_book_apiv1_pb.GetBookByIsbnRequest, proto_book_apiv1_pb.BookResponse> {
    path: "/proto.BookService/GetBookByIsbn";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.GetBookByIsbnRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.GetBookByIsbnRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookResponse>;
}
interface IBookServiceService_IGetBookshelf extends grpc.MethodDefinition<proto_book_apiv1_pb.GetBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse> {
    path: "/proto.BookService/GetBookshelf";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.GetBookshelfRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.GetBookshelfRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookshelfResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookshelfResponse>;
}
interface IBookServiceService_IGetReview extends grpc.MethodDefinition<proto_book_apiv1_pb.GetReviewRequest, proto_book_apiv1_pb.ReviewResponse> {
    path: "/proto.BookService/GetReview";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.GetReviewRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.GetReviewRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.ReviewResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.ReviewResponse>;
}
interface IBookServiceService_ICreateBook extends grpc.MethodDefinition<proto_book_apiv1_pb.CreateBookRequest, proto_book_apiv1_pb.BookResponse> {
    path: "/proto.BookService/CreateBook";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.CreateBookRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.CreateBookRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookResponse>;
}
interface IBookServiceService_IUpdateBook extends grpc.MethodDefinition<proto_book_apiv1_pb.UpdateBookRequest, proto_book_apiv1_pb.BookResponse> {
    path: "/proto.BookService/UpdateBook";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.UpdateBookRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.UpdateBookRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookResponse>;
}
interface IBookServiceService_IReadBookshelf extends grpc.MethodDefinition<proto_book_apiv1_pb.ReadBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse> {
    path: "/proto.BookService/ReadBookshelf";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.ReadBookshelfRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.ReadBookshelfRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookshelfResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookshelfResponse>;
}
interface IBookServiceService_IReadingBookshelf extends grpc.MethodDefinition<proto_book_apiv1_pb.ReadingBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse> {
    path: "/proto.BookService/ReadingBookshelf";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.ReadingBookshelfRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.ReadingBookshelfRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookshelfResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookshelfResponse>;
}
interface IBookServiceService_IStackBookshelf extends grpc.MethodDefinition<proto_book_apiv1_pb.StackBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse> {
    path: "/proto.BookService/StackBookshelf";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.StackBookshelfRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.StackBookshelfRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookshelfResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookshelfResponse>;
}
interface IBookServiceService_IWantBookshelf extends grpc.MethodDefinition<proto_book_apiv1_pb.WantBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse> {
    path: "/proto.BookService/WantBookshelf";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.WantBookshelfRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.WantBookshelfRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookshelfResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookshelfResponse>;
}
interface IBookServiceService_IReleaseBookshelf extends grpc.MethodDefinition<proto_book_apiv1_pb.ReleaseBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse> {
    path: "/proto.BookService/ReleaseBookshelf";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.ReleaseBookshelfRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.ReleaseBookshelfRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookshelfResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookshelfResponse>;
}
interface IBookServiceService_IDeleteBook extends grpc.MethodDefinition<proto_book_apiv1_pb.DeleteBookRequest, proto_book_apiv1_pb.EmptyBook> {
    path: "/proto.BookService/DeleteBook";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.DeleteBookRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.DeleteBookRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.EmptyBook>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.EmptyBook>;
}
interface IBookServiceService_IDeleteBookshelf extends grpc.MethodDefinition<proto_book_apiv1_pb.DeleteBookshelfRequest, proto_book_apiv1_pb.EmptyBook> {
    path: "/proto.BookService/DeleteBookshelf";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.DeleteBookshelfRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.DeleteBookshelfRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.EmptyBook>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.EmptyBook>;
}

export const BookServiceService: IBookServiceService;

export interface IBookServiceServer extends grpc.UntypedServiceImplementation {
    listBookByBookIds: grpc.handleUnaryCall<proto_book_apiv1_pb.ListBookByBookIdsRequest, proto_book_apiv1_pb.BookListResponse>;
    listBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.ListBookshelfRequest, proto_book_apiv1_pb.BookshelfListResponse>;
    listBookReview: grpc.handleUnaryCall<proto_book_apiv1_pb.ListBookReviewRequest, proto_book_apiv1_pb.ReviewListResponse>;
    listUserReview: grpc.handleUnaryCall<proto_book_apiv1_pb.ListUserReviewRequest, proto_book_apiv1_pb.ReviewListResponse>;
    getBook: grpc.handleUnaryCall<proto_book_apiv1_pb.GetBookRequest, proto_book_apiv1_pb.BookResponse>;
    getBookByIsbn: grpc.handleUnaryCall<proto_book_apiv1_pb.GetBookByIsbnRequest, proto_book_apiv1_pb.BookResponse>;
    getBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.GetBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse>;
    getReview: grpc.handleUnaryCall<proto_book_apiv1_pb.GetReviewRequest, proto_book_apiv1_pb.ReviewResponse>;
    createBook: grpc.handleUnaryCall<proto_book_apiv1_pb.CreateBookRequest, proto_book_apiv1_pb.BookResponse>;
    updateBook: grpc.handleUnaryCall<proto_book_apiv1_pb.UpdateBookRequest, proto_book_apiv1_pb.BookResponse>;
    readBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.ReadBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse>;
    readingBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.ReadingBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse>;
    stackBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.StackBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse>;
    wantBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.WantBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse>;
    releaseBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.ReleaseBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse>;
    deleteBook: grpc.handleUnaryCall<proto_book_apiv1_pb.DeleteBookRequest, proto_book_apiv1_pb.EmptyBook>;
    deleteBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.DeleteBookshelfRequest, proto_book_apiv1_pb.EmptyBook>;
}

export interface IBookServiceClient {
    listBookByBookIds(request: proto_book_apiv1_pb.ListBookByBookIdsRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    listBookByBookIds(request: proto_book_apiv1_pb.ListBookByBookIdsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    listBookByBookIds(request: proto_book_apiv1_pb.ListBookByBookIdsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    listBookshelf(request: proto_book_apiv1_pb.ListBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfListResponse) => void): grpc.ClientUnaryCall;
    listBookshelf(request: proto_book_apiv1_pb.ListBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfListResponse) => void): grpc.ClientUnaryCall;
    listBookshelf(request: proto_book_apiv1_pb.ListBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfListResponse) => void): grpc.ClientUnaryCall;
    listBookReview(request: proto_book_apiv1_pb.ListBookReviewRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewListResponse) => void): grpc.ClientUnaryCall;
    listBookReview(request: proto_book_apiv1_pb.ListBookReviewRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewListResponse) => void): grpc.ClientUnaryCall;
    listBookReview(request: proto_book_apiv1_pb.ListBookReviewRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewListResponse) => void): grpc.ClientUnaryCall;
    listUserReview(request: proto_book_apiv1_pb.ListUserReviewRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewListResponse) => void): grpc.ClientUnaryCall;
    listUserReview(request: proto_book_apiv1_pb.ListUserReviewRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewListResponse) => void): grpc.ClientUnaryCall;
    listUserReview(request: proto_book_apiv1_pb.ListUserReviewRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewListResponse) => void): grpc.ClientUnaryCall;
    getBook(request: proto_book_apiv1_pb.GetBookRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    getBook(request: proto_book_apiv1_pb.GetBookRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    getBook(request: proto_book_apiv1_pb.GetBookRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    getBookByIsbn(request: proto_book_apiv1_pb.GetBookByIsbnRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    getBookByIsbn(request: proto_book_apiv1_pb.GetBookByIsbnRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    getBookByIsbn(request: proto_book_apiv1_pb.GetBookByIsbnRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    getBookshelf(request: proto_book_apiv1_pb.GetBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    getBookshelf(request: proto_book_apiv1_pb.GetBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    getBookshelf(request: proto_book_apiv1_pb.GetBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    getReview(request: proto_book_apiv1_pb.GetReviewRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewResponse) => void): grpc.ClientUnaryCall;
    getReview(request: proto_book_apiv1_pb.GetReviewRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewResponse) => void): grpc.ClientUnaryCall;
    getReview(request: proto_book_apiv1_pb.GetReviewRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewResponse) => void): grpc.ClientUnaryCall;
    createBook(request: proto_book_apiv1_pb.CreateBookRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    createBook(request: proto_book_apiv1_pb.CreateBookRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    createBook(request: proto_book_apiv1_pb.CreateBookRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    updateBook(request: proto_book_apiv1_pb.UpdateBookRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    updateBook(request: proto_book_apiv1_pb.UpdateBookRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    updateBook(request: proto_book_apiv1_pb.UpdateBookRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    readBookshelf(request: proto_book_apiv1_pb.ReadBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    readBookshelf(request: proto_book_apiv1_pb.ReadBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    readBookshelf(request: proto_book_apiv1_pb.ReadBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    readingBookshelf(request: proto_book_apiv1_pb.ReadingBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    readingBookshelf(request: proto_book_apiv1_pb.ReadingBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    readingBookshelf(request: proto_book_apiv1_pb.ReadingBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    stackBookshelf(request: proto_book_apiv1_pb.StackBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    stackBookshelf(request: proto_book_apiv1_pb.StackBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    stackBookshelf(request: proto_book_apiv1_pb.StackBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    wantBookshelf(request: proto_book_apiv1_pb.WantBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    wantBookshelf(request: proto_book_apiv1_pb.WantBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    wantBookshelf(request: proto_book_apiv1_pb.WantBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    releaseBookshelf(request: proto_book_apiv1_pb.ReleaseBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    releaseBookshelf(request: proto_book_apiv1_pb.ReleaseBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    releaseBookshelf(request: proto_book_apiv1_pb.ReleaseBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    deleteBook(request: proto_book_apiv1_pb.DeleteBookRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.EmptyBook) => void): grpc.ClientUnaryCall;
    deleteBook(request: proto_book_apiv1_pb.DeleteBookRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.EmptyBook) => void): grpc.ClientUnaryCall;
    deleteBook(request: proto_book_apiv1_pb.DeleteBookRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.EmptyBook) => void): grpc.ClientUnaryCall;
    deleteBookshelf(request: proto_book_apiv1_pb.DeleteBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.EmptyBook) => void): grpc.ClientUnaryCall;
    deleteBookshelf(request: proto_book_apiv1_pb.DeleteBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.EmptyBook) => void): grpc.ClientUnaryCall;
    deleteBookshelf(request: proto_book_apiv1_pb.DeleteBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.EmptyBook) => void): grpc.ClientUnaryCall;
}

export class BookServiceClient extends grpc.Client implements IBookServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public listBookByBookIds(request: proto_book_apiv1_pb.ListBookByBookIdsRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    public listBookByBookIds(request: proto_book_apiv1_pb.ListBookByBookIdsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    public listBookByBookIds(request: proto_book_apiv1_pb.ListBookByBookIdsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    public listBookshelf(request: proto_book_apiv1_pb.ListBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfListResponse) => void): grpc.ClientUnaryCall;
    public listBookshelf(request: proto_book_apiv1_pb.ListBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfListResponse) => void): grpc.ClientUnaryCall;
    public listBookshelf(request: proto_book_apiv1_pb.ListBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfListResponse) => void): grpc.ClientUnaryCall;
    public listBookReview(request: proto_book_apiv1_pb.ListBookReviewRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewListResponse) => void): grpc.ClientUnaryCall;
    public listBookReview(request: proto_book_apiv1_pb.ListBookReviewRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewListResponse) => void): grpc.ClientUnaryCall;
    public listBookReview(request: proto_book_apiv1_pb.ListBookReviewRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewListResponse) => void): grpc.ClientUnaryCall;
    public listUserReview(request: proto_book_apiv1_pb.ListUserReviewRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewListResponse) => void): grpc.ClientUnaryCall;
    public listUserReview(request: proto_book_apiv1_pb.ListUserReviewRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewListResponse) => void): grpc.ClientUnaryCall;
    public listUserReview(request: proto_book_apiv1_pb.ListUserReviewRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewListResponse) => void): grpc.ClientUnaryCall;
    public getBook(request: proto_book_apiv1_pb.GetBookRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public getBook(request: proto_book_apiv1_pb.GetBookRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public getBook(request: proto_book_apiv1_pb.GetBookRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public getBookByIsbn(request: proto_book_apiv1_pb.GetBookByIsbnRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public getBookByIsbn(request: proto_book_apiv1_pb.GetBookByIsbnRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public getBookByIsbn(request: proto_book_apiv1_pb.GetBookByIsbnRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public getBookshelf(request: proto_book_apiv1_pb.GetBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public getBookshelf(request: proto_book_apiv1_pb.GetBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public getBookshelf(request: proto_book_apiv1_pb.GetBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public getReview(request: proto_book_apiv1_pb.GetReviewRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewResponse) => void): grpc.ClientUnaryCall;
    public getReview(request: proto_book_apiv1_pb.GetReviewRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewResponse) => void): grpc.ClientUnaryCall;
    public getReview(request: proto_book_apiv1_pb.GetReviewRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.ReviewResponse) => void): grpc.ClientUnaryCall;
    public createBook(request: proto_book_apiv1_pb.CreateBookRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public createBook(request: proto_book_apiv1_pb.CreateBookRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public createBook(request: proto_book_apiv1_pb.CreateBookRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public updateBook(request: proto_book_apiv1_pb.UpdateBookRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public updateBook(request: proto_book_apiv1_pb.UpdateBookRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public updateBook(request: proto_book_apiv1_pb.UpdateBookRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public readBookshelf(request: proto_book_apiv1_pb.ReadBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public readBookshelf(request: proto_book_apiv1_pb.ReadBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public readBookshelf(request: proto_book_apiv1_pb.ReadBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public readingBookshelf(request: proto_book_apiv1_pb.ReadingBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public readingBookshelf(request: proto_book_apiv1_pb.ReadingBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public readingBookshelf(request: proto_book_apiv1_pb.ReadingBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public stackBookshelf(request: proto_book_apiv1_pb.StackBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public stackBookshelf(request: proto_book_apiv1_pb.StackBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public stackBookshelf(request: proto_book_apiv1_pb.StackBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public wantBookshelf(request: proto_book_apiv1_pb.WantBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public wantBookshelf(request: proto_book_apiv1_pb.WantBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public wantBookshelf(request: proto_book_apiv1_pb.WantBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public releaseBookshelf(request: proto_book_apiv1_pb.ReleaseBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public releaseBookshelf(request: proto_book_apiv1_pb.ReleaseBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public releaseBookshelf(request: proto_book_apiv1_pb.ReleaseBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public deleteBook(request: proto_book_apiv1_pb.DeleteBookRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.EmptyBook) => void): grpc.ClientUnaryCall;
    public deleteBook(request: proto_book_apiv1_pb.DeleteBookRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.EmptyBook) => void): grpc.ClientUnaryCall;
    public deleteBook(request: proto_book_apiv1_pb.DeleteBookRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.EmptyBook) => void): grpc.ClientUnaryCall;
    public deleteBookshelf(request: proto_book_apiv1_pb.DeleteBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.EmptyBook) => void): grpc.ClientUnaryCall;
    public deleteBookshelf(request: proto_book_apiv1_pb.DeleteBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.EmptyBook) => void): grpc.ClientUnaryCall;
    public deleteBookshelf(request: proto_book_apiv1_pb.DeleteBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.EmptyBook) => void): grpc.ClientUnaryCall;
}
