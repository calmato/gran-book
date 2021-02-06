// package: proto
// file: proto/user_apiv1.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as proto_user_apiv1_pb from "../proto/user_apiv1_pb";

interface IAuthServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getAuth: IAuthServiceService_IGetAuth;
    createAuth: IAuthServiceService_ICreateAuth;
    updateAuthEmail: IAuthServiceService_IUpdateAuthEmail;
    updateAuthPassword: IAuthServiceService_IUpdateAuthPassword;
    updateAuthProfile: IAuthServiceService_IUpdateAuthProfile;
    updateAuthAddress: IAuthServiceService_IUpdateAuthAddress;
}

interface IAuthServiceService_IGetAuth extends grpc.MethodDefinition<proto_user_apiv1_pb.EmptyUser, proto_user_apiv1_pb.AuthResponse> {
    path: "/proto.AuthService/GetAuth";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.EmptyUser>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.EmptyUser>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AuthResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AuthResponse>;
}
interface IAuthServiceService_ICreateAuth extends grpc.MethodDefinition<proto_user_apiv1_pb.CreateAuthRequest, proto_user_apiv1_pb.AuthResponse> {
    path: "/proto.AuthService/CreateAuth";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.CreateAuthRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.CreateAuthRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AuthResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AuthResponse>;
}
interface IAuthServiceService_IUpdateAuthEmail extends grpc.MethodDefinition<proto_user_apiv1_pb.UpdateAuthEmailRequest, proto_user_apiv1_pb.AuthResponse> {
    path: "/proto.AuthService/UpdateAuthEmail";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.UpdateAuthEmailRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.UpdateAuthEmailRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AuthResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AuthResponse>;
}
interface IAuthServiceService_IUpdateAuthPassword extends grpc.MethodDefinition<proto_user_apiv1_pb.UpdateAuthPasswordRequest, proto_user_apiv1_pb.AuthResponse> {
    path: "/proto.AuthService/UpdateAuthPassword";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.UpdateAuthPasswordRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.UpdateAuthPasswordRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AuthResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AuthResponse>;
}
interface IAuthServiceService_IUpdateAuthProfile extends grpc.MethodDefinition<proto_user_apiv1_pb.UpdateAuthProfileRequest, proto_user_apiv1_pb.AuthResponse> {
    path: "/proto.AuthService/UpdateAuthProfile";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.UpdateAuthProfileRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.UpdateAuthProfileRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AuthResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AuthResponse>;
}
interface IAuthServiceService_IUpdateAuthAddress extends grpc.MethodDefinition<proto_user_apiv1_pb.UpdateAuthAddressRequest, proto_user_apiv1_pb.AuthResponse> {
    path: "/proto.AuthService/UpdateAuthAddress";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.UpdateAuthAddressRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.UpdateAuthAddressRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AuthResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AuthResponse>;
}

export const AuthServiceService: IAuthServiceService;

export interface IAuthServiceServer extends grpc.UntypedServiceImplementation {
    getAuth: grpc.handleUnaryCall<proto_user_apiv1_pb.EmptyUser, proto_user_apiv1_pb.AuthResponse>;
    createAuth: grpc.handleUnaryCall<proto_user_apiv1_pb.CreateAuthRequest, proto_user_apiv1_pb.AuthResponse>;
    updateAuthEmail: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAuthEmailRequest, proto_user_apiv1_pb.AuthResponse>;
    updateAuthPassword: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAuthPasswordRequest, proto_user_apiv1_pb.AuthResponse>;
    updateAuthProfile: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAuthProfileRequest, proto_user_apiv1_pb.AuthResponse>;
    updateAuthAddress: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAuthAddressRequest, proto_user_apiv1_pb.AuthResponse>;
}

export interface IAuthServiceClient {
    getAuth(request: proto_user_apiv1_pb.EmptyUser, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    getAuth(request: proto_user_apiv1_pb.EmptyUser, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    getAuth(request: proto_user_apiv1_pb.EmptyUser, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    createAuth(request: proto_user_apiv1_pb.CreateAuthRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    createAuth(request: proto_user_apiv1_pb.CreateAuthRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    createAuth(request: proto_user_apiv1_pb.CreateAuthRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    updateAuthEmail(request: proto_user_apiv1_pb.UpdateAuthEmailRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    updateAuthEmail(request: proto_user_apiv1_pb.UpdateAuthEmailRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    updateAuthEmail(request: proto_user_apiv1_pb.UpdateAuthEmailRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    updateAuthPassword(request: proto_user_apiv1_pb.UpdateAuthPasswordRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    updateAuthPassword(request: proto_user_apiv1_pb.UpdateAuthPasswordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    updateAuthPassword(request: proto_user_apiv1_pb.UpdateAuthPasswordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    updateAuthProfile(request: proto_user_apiv1_pb.UpdateAuthProfileRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    updateAuthProfile(request: proto_user_apiv1_pb.UpdateAuthProfileRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    updateAuthProfile(request: proto_user_apiv1_pb.UpdateAuthProfileRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    updateAuthAddress(request: proto_user_apiv1_pb.UpdateAuthAddressRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    updateAuthAddress(request: proto_user_apiv1_pb.UpdateAuthAddressRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    updateAuthAddress(request: proto_user_apiv1_pb.UpdateAuthAddressRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
}

export class AuthServiceClient extends grpc.Client implements IAuthServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public getAuth(request: proto_user_apiv1_pb.EmptyUser, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public getAuth(request: proto_user_apiv1_pb.EmptyUser, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public getAuth(request: proto_user_apiv1_pb.EmptyUser, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public createAuth(request: proto_user_apiv1_pb.CreateAuthRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public createAuth(request: proto_user_apiv1_pb.CreateAuthRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public createAuth(request: proto_user_apiv1_pb.CreateAuthRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public updateAuthEmail(request: proto_user_apiv1_pb.UpdateAuthEmailRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public updateAuthEmail(request: proto_user_apiv1_pb.UpdateAuthEmailRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public updateAuthEmail(request: proto_user_apiv1_pb.UpdateAuthEmailRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public updateAuthPassword(request: proto_user_apiv1_pb.UpdateAuthPasswordRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public updateAuthPassword(request: proto_user_apiv1_pb.UpdateAuthPasswordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public updateAuthPassword(request: proto_user_apiv1_pb.UpdateAuthPasswordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public updateAuthProfile(request: proto_user_apiv1_pb.UpdateAuthProfileRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public updateAuthProfile(request: proto_user_apiv1_pb.UpdateAuthProfileRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public updateAuthProfile(request: proto_user_apiv1_pb.UpdateAuthProfileRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public updateAuthAddress(request: proto_user_apiv1_pb.UpdateAuthAddressRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public updateAuthAddress(request: proto_user_apiv1_pb.UpdateAuthAddressRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public updateAuthAddress(request: proto_user_apiv1_pb.UpdateAuthAddressRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
}
