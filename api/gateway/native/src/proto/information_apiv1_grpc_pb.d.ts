// package: proto
// file: proto/information_apiv1.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as proto_information_apiv1_pb from "../proto/information_apiv1_pb";

interface INotificationServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    reply: INotificationServiceService_IReply;
}

interface INotificationServiceService_IReply extends grpc.MethodDefinition<proto_information_apiv1_pb.HelloRequest, proto_information_apiv1_pb.HelloResponse> {
    path: "/proto.NotificationService/Reply";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_information_apiv1_pb.HelloRequest>;
    requestDeserialize: grpc.deserialize<proto_information_apiv1_pb.HelloRequest>;
    responseSerialize: grpc.serialize<proto_information_apiv1_pb.HelloResponse>;
    responseDeserialize: grpc.deserialize<proto_information_apiv1_pb.HelloResponse>;
}

export const NotificationServiceService: INotificationServiceService;

export interface INotificationServiceServer extends grpc.UntypedServiceImplementation {
    reply: grpc.handleUnaryCall<proto_information_apiv1_pb.HelloRequest, proto_information_apiv1_pb.HelloResponse>;
}

export interface INotificationServiceClient {
    reply(request: proto_information_apiv1_pb.HelloRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.HelloResponse) => void): grpc.ClientUnaryCall;
    reply(request: proto_information_apiv1_pb.HelloRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.HelloResponse) => void): grpc.ClientUnaryCall;
    reply(request: proto_information_apiv1_pb.HelloRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.HelloResponse) => void): grpc.ClientUnaryCall;
}

export class NotificationServiceClient extends grpc.Client implements INotificationServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public reply(request: proto_information_apiv1_pb.HelloRequest, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.HelloResponse) => void): grpc.ClientUnaryCall;
    public reply(request: proto_information_apiv1_pb.HelloRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.HelloResponse) => void): grpc.ClientUnaryCall;
    public reply(request: proto_information_apiv1_pb.HelloRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_information_apiv1_pb.HelloResponse) => void): grpc.ClientUnaryCall;
}
