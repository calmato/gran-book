// package: proto
// file: proto/book_apiv1.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as proto_book_apiv1_pb from "../proto/book_apiv1_pb";

interface IBookServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createBook: IBookServiceService_ICreateBook;
    createAndUpdateBooks: IBookServiceService_ICreateAndUpdateBooks;
    registerBookshelf: IBookServiceService_IRegisterBookshelf;
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
interface IBookServiceService_ICreateAndUpdateBooks extends grpc.MethodDefinition<proto_book_apiv1_pb.CreateAndUpdateBooksRequest, proto_book_apiv1_pb.BookListResponse> {
    path: "/proto.BookService/CreateAndUpdateBooks";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.CreateAndUpdateBooksRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.CreateAndUpdateBooksRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookListResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookListResponse>;
}
interface IBookServiceService_IRegisterBookshelf extends grpc.MethodDefinition<proto_book_apiv1_pb.RegisterBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse> {
    path: "/proto.BookService/RegisterBookshelf";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.RegisterBookshelfRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.RegisterBookshelfRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookshelfResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookshelfResponse>;
}

export const BookServiceService: IBookServiceService;

export interface IBookServiceServer extends grpc.UntypedServiceImplementation {
    createBook: grpc.handleUnaryCall<proto_book_apiv1_pb.CreateBookRequest, proto_book_apiv1_pb.BookResponse>;
    createAndUpdateBooks: grpc.handleUnaryCall<proto_book_apiv1_pb.CreateAndUpdateBooksRequest, proto_book_apiv1_pb.BookListResponse>;
    registerBookshelf: grpc.handleUnaryCall<proto_book_apiv1_pb.RegisterBookshelfRequest, proto_book_apiv1_pb.BookshelfResponse>;
}

export interface IBookServiceClient {
    createBook(request: proto_book_apiv1_pb.CreateBookRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    createBook(request: proto_book_apiv1_pb.CreateBookRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    createBook(request: proto_book_apiv1_pb.CreateBookRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    createAndUpdateBooks(request: proto_book_apiv1_pb.CreateAndUpdateBooksRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    createAndUpdateBooks(request: proto_book_apiv1_pb.CreateAndUpdateBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    createAndUpdateBooks(request: proto_book_apiv1_pb.CreateAndUpdateBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    registerBookshelf(request: proto_book_apiv1_pb.RegisterBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    registerBookshelf(request: proto_book_apiv1_pb.RegisterBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    registerBookshelf(request: proto_book_apiv1_pb.RegisterBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
}

export class BookServiceClient extends grpc.Client implements IBookServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public createBook(request: proto_book_apiv1_pb.CreateBookRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public createBook(request: proto_book_apiv1_pb.CreateBookRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public createBook(request: proto_book_apiv1_pb.CreateBookRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public createAndUpdateBooks(request: proto_book_apiv1_pb.CreateAndUpdateBooksRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    public createAndUpdateBooks(request: proto_book_apiv1_pb.CreateAndUpdateBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    public createAndUpdateBooks(request: proto_book_apiv1_pb.CreateAndUpdateBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    public registerBookshelf(request: proto_book_apiv1_pb.RegisterBookshelfRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public registerBookshelf(request: proto_book_apiv1_pb.RegisterBookshelfRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
    public registerBookshelf(request: proto_book_apiv1_pb.RegisterBookshelfRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookshelfResponse) => void): grpc.ClientUnaryCall;
}
