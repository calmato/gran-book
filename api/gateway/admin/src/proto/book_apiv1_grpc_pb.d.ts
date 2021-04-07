// package: proto
// file: proto/book_apiv1.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as proto_book_apiv1_pb from "../proto/book_apiv1_pb";

interface IBookServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createAndUpdateBooks: IBookServiceService_ICreateAndUpdateBooks;
    readBookshelf: IBookServiceService_IReadBookshelf;
    readingBookshelf: IBookServiceService_IReadingBookshelf;
    stackBookshelf: IBookServiceService_IStackBookshelf;
    wantBookshelf: IBookServiceService_IWantBookshelf;
    releaseBookshelf: IBookServiceService_IReleaseBookshelf;
}

interface IBookServiceService_ICreateAndUpdateBooks extends grpc.MethodDefinition<proto_book_apiv1_pb.CreateAndUpdateBooksRequest, proto_book_apiv1_pb.BookListResponse> {
    path: "/proto.BookService/CreateAndUpdateBooks";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.CreateAndUpdateBooksRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.CreateAndUpdateBooksRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookListResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookListResponse>;
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

export const BookServiceService: IBookServiceService;

export interface IBookServiceServer extends grpc.UntypedServiceImplementation {
    createAndUpdateBooks: grpc.handleUnaryCall<proto_book_apiv1_pb.CreateAndUpdateBooksRequest, proto_book_apiv1_pb.BookListResponse>;
    readBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.ReadBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse>;
    readingBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.ReadingBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse>;
    stackBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.StackBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse>;
    wantBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.WantBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse>;
    releaseBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.ReleaseBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse>;
}

export interface IBookServiceClient {
    createAndUpdateBooks(request: proto_book_apiv1_pb.CreateAndUpdateBooksRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    createAndUpdateBooks(request: proto_book_apiv1_pb.CreateAndUpdateBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    createAndUpdateBooks(request: proto_book_apiv1_pb.CreateAndUpdateBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
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
}

export class BookServiceClient extends grpc.Client implements IBookServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public createAndUpdateBooks(request: proto_book_apiv1_pb.CreateAndUpdateBooksRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    public createAndUpdateBooks(request: proto_book_apiv1_pb.CreateAndUpdateBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    public createAndUpdateBooks(request: proto_book_apiv1_pb.CreateAndUpdateBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
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
}
