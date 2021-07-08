// package: proto
// file: proto/information_apiv1.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as proto_information_apiv1_pb from "../proto/information_apiv1_pb";

interface INotificationServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listNotification: INotificationServiceService_IListNotification;
    getNotification: INotificationServiceService_IGetNotification;
    createNotification: INotificationServiceService_ICreateNotification;
    updateNotification: INotificationServiceService_IUpdateNotification;
    deleteNotification: INotificationServiceService_IDeleteNotification;
}

interface INotificationServiceService_IListNotification extends grpc.MethodDefinition<proto_information_apiv1_pb.EmptyNotification, proto_information_apiv1_pb.NotificationResponse> {
    path: "/proto.NotificationService/ListNotification";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_information_apiv1_pb.EmptyNotification>;
    requestDeserialize: grpc.deserialize<proto_information_apiv1_pb.EmptyNotification>;
    responseSerialize: grpc.serialize<proto_information_apiv1_pb.NotificationResponse>;
    responseDeserialize: grpc.deserialize<proto_information_apiv1_pb.NotificationResponse>;
}
interface INotificationServiceService_IGetNotification extends grpc.MethodDefinition<proto_information_apiv1_pb.GetNotificationRequest, proto_information_apiv1_pb.NotificationResponse> {
    path: "/proto.NotificationService/GetNotification";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_information_apiv1_pb.GetNotificationRequest>;
    requestDeserialize: grpc.deserialize<proto_information_apiv1_pb.GetNotificationRequest>;
    responseSerialize: grpc.serialize<proto_information_apiv1_pb.NotificationResponse>;
    responseDeserialize: grpc.deserialize<proto_information_apiv1_pb.NotificationResponse>;
}
interface INotificationServiceService_ICreateNotification extends grpc.MethodDefinition<proto_information_apiv1_pb.CreateNotificationRequest, proto_information_apiv1_pb.NotificationResponse> {
    path: "/proto.NotificationService/CreateNotification";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_information_apiv1_pb.CreateNotificationRequest>;
    requestDeserialize: grpc.deserialize<proto_information_apiv1_pb.CreateNotificationRequest>;
    responseSerialize: grpc.serialize<proto_information_apiv1_pb.NotificationResponse>;
    responseDeserialize: grpc.deserialize<proto_information_apiv1_pb.NotificationResponse>;
}
interface INotificationServiceService_IUpdateNotification extends grpc.MethodDefinition<proto_information_apiv1_pb.UpdateNotificationRequest, proto_information_apiv1_pb.NotificationResponse> {
    path: "/proto.NotificationService/UpdateNotification";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_information_apiv1_pb.UpdateNotificationRequest>;
    requestDeserialize: grpc.deserialize<proto_information_apiv1_pb.UpdateNotificationRequest>;
    responseSerialize: grpc.serialize<proto_information_apiv1_pb.NotificationResponse>;
    responseDeserialize: grpc.deserialize<proto_information_apiv1_pb.NotificationResponse>;
}
interface INotificationServiceService_IDeleteNotification extends grpc.MethodDefinition<proto_information_apiv1_pb.DeleteNotificationRequest, proto_information_apiv1_pb.NotificationResponse> {
    path: "/proto.NotificationService/DeleteNotification";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_information_apiv1_pb.DeleteNotificationRequest>;
    requestDeserialize: grpc.deserialize<proto_information_apiv1_pb.DeleteNotificationRequest>;
    responseSerialize: grpc.serialize<proto_information_apiv1_pb.NotificationResponse>;
    responseDeserialize: grpc.deserialize<proto_information_apiv1_pb.NotificationResponse>;
}

export const NotificationServiceService: INotificationServiceService;

export interface INotificationServiceServer extends grpc.UntypedServiceImplementation {
    listNotification: grpc.handleUnaryCall<proto_information_apiv1_pb.EmptyNotification, proto_information_apiv1_pb.NotificationResponse>;
    getNotification: grpc.handleUnaryCall<proto_information_apiv1_pb.GetNotificationRequest, proto_information_apiv1_pb.NotificationResponse>;
    createNotification: grpc.handleUnaryCall<proto_information_apiv1_pb.CreateNotificationRequest, proto_information_apiv1_pb.NotificationResponse>;
    updateNotification: grpc.handleUnaryCall<proto_information_apiv1_pb.UpdateNotificationRequest, proto_information_apiv1_pb.NotificationResponse>;
    deleteNotification: grpc.handleUnaryCall<proto_information_apiv1_pb.DeleteNotificationRequest, proto_information_apiv1_pb.NotificationResponse>;
}

export interface INotificationServiceClient {
    listNotification(request: proto_information_apiv1_pb.EmptyNotification, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    listNotification(request: proto_information_apiv1_pb.EmptyNotification, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    listNotification(request: proto_information_apiv1_pb.EmptyNotification, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    getNotification(request: proto_information_apiv1_pb.GetNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    getNotification(request: proto_information_apiv1_pb.GetNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    getNotification(request: proto_information_apiv1_pb.GetNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    createNotification(request: proto_information_apiv1_pb.CreateNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    createNotification(request: proto_information_apiv1_pb.CreateNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    createNotification(request: proto_information_apiv1_pb.CreateNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    updateNotification(request: proto_information_apiv1_pb.UpdateNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    updateNotification(request: proto_information_apiv1_pb.UpdateNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    updateNotification(request: proto_information_apiv1_pb.UpdateNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    deleteNotification(request: proto_information_apiv1_pb.DeleteNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    deleteNotification(request: proto_information_apiv1_pb.DeleteNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    deleteNotification(request: proto_information_apiv1_pb.DeleteNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
}

export class NotificationServiceClient extends grpc.Client implements INotificationServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public listNotification(request: proto_information_apiv1_pb.EmptyNotification, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public listNotification(request: proto_information_apiv1_pb.EmptyNotification, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public listNotification(request: proto_information_apiv1_pb.EmptyNotification, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public getNotification(request: proto_information_apiv1_pb.GetNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public getNotification(request: proto_information_apiv1_pb.GetNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public getNotification(request: proto_information_apiv1_pb.GetNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public createNotification(request: proto_information_apiv1_pb.CreateNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public createNotification(request: proto_information_apiv1_pb.CreateNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public createNotification(request: proto_information_apiv1_pb.CreateNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public updateNotification(request: proto_information_apiv1_pb.UpdateNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public updateNotification(request: proto_information_apiv1_pb.UpdateNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public updateNotification(request: proto_information_apiv1_pb.UpdateNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public deleteNotification(request: proto_information_apiv1_pb.DeleteNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public deleteNotification(request: proto_information_apiv1_pb.DeleteNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public deleteNotification(request: proto_information_apiv1_pb.DeleteNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
}

interface IInquiryServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createInquiry: IInquiryServiceService_ICreateInquiry;
}

interface IInquiryServiceService_ICreateInquiry extends grpc.MethodDefinition<proto_information_apiv1_pb.CreateInquiryRequest, proto_information_apiv1_pb.InquiryResponse> {
    path: "/proto.InquiryService/CreateInquiry";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_information_apiv1_pb.CreateInquiryRequest>;
    requestDeserialize: grpc.deserialize<proto_information_apiv1_pb.CreateInquiryRequest>;
    responseSerialize: grpc.serialize<proto_information_apiv1_pb.InquiryResponse>;
    responseDeserialize: grpc.deserialize<proto_information_apiv1_pb.InquiryResponse>;
}

export const InquiryServiceService: IInquiryServiceService;

export interface IInquiryServiceServer extends grpc.UntypedServiceImplementation {
    createInquiry: grpc.handleUnaryCall<proto_information_apiv1_pb.CreateInquiryRequest, proto_information_apiv1_pb.InquiryResponse>;
}

export interface IInquiryServiceClient {
    createInquiry(request: proto_information_apiv1_pb.CreateInquiryRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.InquiryResponse) => void): grpc.ClientUnaryCall;
    createInquiry(request: proto_information_apiv1_pb.CreateInquiryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.InquiryResponse) => void): grpc.ClientUnaryCall;
    createInquiry(request: proto_information_apiv1_pb.CreateInquiryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.InquiryResponse) => void): grpc.ClientUnaryCall;
}

export class InquiryServiceClient extends grpc.Client implements IInquiryServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public createInquiry(request: proto_information_apiv1_pb.CreateInquiryRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.InquiryResponse) => void): grpc.ClientUnaryCall;
    public createInquiry(request: proto_information_apiv1_pb.CreateInquiryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.InquiryResponse) => void): grpc.ClientUnaryCall;
    public createInquiry(request: proto_information_apiv1_pb.CreateInquiryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.InquiryResponse) => void): grpc.ClientUnaryCall;
}
