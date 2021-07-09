// package: proto
// file: proto/information_apiv1.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as proto_information_apiv1_pb from "../proto/information_apiv1_pb";

interface INotificationServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    list: INotificationServiceService_IList;
    get: INotificationServiceService_IGet;
    create: INotificationServiceService_ICreate;
    update: INotificationServiceService_IUpdate;
    delete: INotificationServiceService_IDelete;
}

interface INotificationServiceService_IList extends grpc.MethodDefinition<proto_information_apiv1_pb.EmptyNotification, proto_information_apiv1_pb.NotificationResponse> {
    path: "/proto.NotificationService/List";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_information_apiv1_pb.EmptyNotification>;
    requestDeserialize: grpc.deserialize<proto_information_apiv1_pb.EmptyNotification>;
    responseSerialize: grpc.serialize<proto_information_apiv1_pb.NotificationResponse>;
    responseDeserialize: grpc.deserialize<proto_information_apiv1_pb.NotificationResponse>;
}
interface INotificationServiceService_IGet extends grpc.MethodDefinition<proto_information_apiv1_pb.GetNotificationRequest, proto_information_apiv1_pb.NotificationResponse> {
    path: "/proto.NotificationService/Get";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_information_apiv1_pb.GetNotificationRequest>;
    requestDeserialize: grpc.deserialize<proto_information_apiv1_pb.GetNotificationRequest>;
    responseSerialize: grpc.serialize<proto_information_apiv1_pb.NotificationResponse>;
    responseDeserialize: grpc.deserialize<proto_information_apiv1_pb.NotificationResponse>;
}
interface INotificationServiceService_ICreate extends grpc.MethodDefinition<proto_information_apiv1_pb.CreateNotificationRequest, proto_information_apiv1_pb.NotificationResponse> {
    path: "/proto.NotificationService/Create";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_information_apiv1_pb.CreateNotificationRequest>;
    requestDeserialize: grpc.deserialize<proto_information_apiv1_pb.CreateNotificationRequest>;
    responseSerialize: grpc.serialize<proto_information_apiv1_pb.NotificationResponse>;
    responseDeserialize: grpc.deserialize<proto_information_apiv1_pb.NotificationResponse>;
}
interface INotificationServiceService_IUpdate extends grpc.MethodDefinition<proto_information_apiv1_pb.UpdateNotificationRequest, proto_information_apiv1_pb.NotificationResponse> {
    path: "/proto.NotificationService/Update";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_information_apiv1_pb.UpdateNotificationRequest>;
    requestDeserialize: grpc.deserialize<proto_information_apiv1_pb.UpdateNotificationRequest>;
    responseSerialize: grpc.serialize<proto_information_apiv1_pb.NotificationResponse>;
    responseDeserialize: grpc.deserialize<proto_information_apiv1_pb.NotificationResponse>;
}
interface INotificationServiceService_IDelete extends grpc.MethodDefinition<proto_information_apiv1_pb.DeleteNotificationRequest, proto_information_apiv1_pb.NotificationResponse> {
    path: "/proto.NotificationService/Delete";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_information_apiv1_pb.DeleteNotificationRequest>;
    requestDeserialize: grpc.deserialize<proto_information_apiv1_pb.DeleteNotificationRequest>;
    responseSerialize: grpc.serialize<proto_information_apiv1_pb.NotificationResponse>;
    responseDeserialize: grpc.deserialize<proto_information_apiv1_pb.NotificationResponse>;
}

export const NotificationServiceService: INotificationServiceService;

export interface INotificationServiceServer extends grpc.UntypedServiceImplementation {
    list: grpc.handleUnaryCall<proto_information_apiv1_pb.EmptyNotification, proto_information_apiv1_pb.NotificationResponse>;
    get: grpc.handleUnaryCall<proto_information_apiv1_pb.GetNotificationRequest, proto_information_apiv1_pb.NotificationResponse>;
    create: grpc.handleUnaryCall<proto_information_apiv1_pb.CreateNotificationRequest, proto_information_apiv1_pb.NotificationResponse>;
    update: grpc.handleUnaryCall<proto_information_apiv1_pb.UpdateNotificationRequest, proto_information_apiv1_pb.NotificationResponse>;
    delete: grpc.handleUnaryCall<proto_information_apiv1_pb.DeleteNotificationRequest, proto_information_apiv1_pb.NotificationResponse>;
}

export interface INotificationServiceClient {
    list(request: proto_information_apiv1_pb.EmptyNotification, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    list(request: proto_information_apiv1_pb.EmptyNotification, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    list(request: proto_information_apiv1_pb.EmptyNotification, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    get(request: proto_information_apiv1_pb.GetNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    get(request: proto_information_apiv1_pb.GetNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    get(request: proto_information_apiv1_pb.GetNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    create(request: proto_information_apiv1_pb.CreateNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    create(request: proto_information_apiv1_pb.CreateNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    create(request: proto_information_apiv1_pb.CreateNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    update(request: proto_information_apiv1_pb.UpdateNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    update(request: proto_information_apiv1_pb.UpdateNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    update(request: proto_information_apiv1_pb.UpdateNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    delete(request: proto_information_apiv1_pb.DeleteNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    delete(request: proto_information_apiv1_pb.DeleteNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    delete(request: proto_information_apiv1_pb.DeleteNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
}

export class NotificationServiceClient extends grpc.Client implements INotificationServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public list(request: proto_information_apiv1_pb.EmptyNotification, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public list(request: proto_information_apiv1_pb.EmptyNotification, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public list(request: proto_information_apiv1_pb.EmptyNotification, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public get(request: proto_information_apiv1_pb.GetNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public get(request: proto_information_apiv1_pb.GetNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public get(request: proto_information_apiv1_pb.GetNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public create(request: proto_information_apiv1_pb.CreateNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public create(request: proto_information_apiv1_pb.CreateNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public create(request: proto_information_apiv1_pb.CreateNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public update(request: proto_information_apiv1_pb.UpdateNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public update(request: proto_information_apiv1_pb.UpdateNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public update(request: proto_information_apiv1_pb.UpdateNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public delete(request: proto_information_apiv1_pb.DeleteNotificationRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public delete(request: proto_information_apiv1_pb.DeleteNotificationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
    public delete(request: proto_information_apiv1_pb.DeleteNotificationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.NotificationResponse) => void): grpc.ClientUnaryCall;
}
