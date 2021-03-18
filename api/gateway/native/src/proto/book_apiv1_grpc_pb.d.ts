// package: proto
// file: proto/book_apiv1.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as proto_book_apiv1_pb from "../proto/book_apiv1_pb";

interface IBookServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createBook: IBookServiceService_ICreateBook;
    createMultipleBooks: IBookServiceService_ICreateMultipleBooks;
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
interface IBookServiceService_ICreateMultipleBooks extends grpc.MethodDefinition<proto_book_apiv1_pb.CreateMultipleBooksRequest, proto_book_apiv1_pb.BookListResponse> {
    path: "/proto.BookService/CreateMultipleBooks";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_book_apiv1_pb.CreateMultipleBooksRequest>;
    requestDeserialize: grpc.deserialize<proto_book_apiv1_pb.CreateMultipleBooksRequest>;
    responseSerialize: grpc.serialize<proto_book_apiv1_pb.BookListResponse>;
    responseDeserialize: grpc.deserialize<proto_book_apiv1_pb.BookListResponse>;
}

export const BookServiceService: IBookServiceService;

export interface IBookServiceServer extends grpc.UntypedServiceImplementation {
    createBook: grpc.handleUnaryCall<proto_book_apiv1_pb.CreateBookRequest, proto_book_apiv1_pb.BookResponse>;
    createMultipleBooks: grpc.handleUnaryCall<proto_book_apiv1_pb.CreateMultipleBooksRequest, proto_book_apiv1_pb.BookListResponse>;
}

export interface IBookServiceClient {
    createBook(request: proto_book_apiv1_pb.CreateBookRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    createBook(request: proto_book_apiv1_pb.CreateBookRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    createBook(request: proto_book_apiv1_pb.CreateBookRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    createMultipleBooks(request: proto_book_apiv1_pb.CreateMultipleBooksRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    createMultipleBooks(request: proto_book_apiv1_pb.CreateMultipleBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    createMultipleBooks(request: proto_book_apiv1_pb.CreateMultipleBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
}

export class BookServiceClient extends grpc.Client implements IBookServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public createBook(request: proto_book_apiv1_pb.CreateBookRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public createBook(request: proto_book_apiv1_pb.CreateBookRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public createBook(request: proto_book_apiv1_pb.CreateBookRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public createMultipleBooks(request: proto_book_apiv1_pb.CreateMultipleBooksRequest, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    public createMultipleBooks(request: proto_book_apiv1_pb.CreateMultipleBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
    public createMultipleBooks(request: proto_book_apiv1_pb.CreateMultipleBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_book_apiv1_pb.BookListResponse) => void): grpc.ClientUnaryCall;
}
