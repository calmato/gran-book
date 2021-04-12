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
    uploadAuthThumbnail: IAuthServiceService_IUploadAuthThumbnail;
    deleteAuth: IAuthServiceService_IDeleteAuth;
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
interface IAuthServiceService_IUploadAuthThumbnail extends grpc.MethodDefinition<proto_user_apiv1_pb.UploadAuthThumbnailRequest, proto_user_apiv1_pb.AuthThumbnailResponse> {
    path: "/proto.AuthService/UploadAuthThumbnail";
    requestStream: true;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.UploadAuthThumbnailRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.UploadAuthThumbnailRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AuthThumbnailResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AuthThumbnailResponse>;
}
interface IAuthServiceService_IDeleteAuth extends grpc.MethodDefinition<proto_user_apiv1_pb.EmptyUser, proto_user_apiv1_pb.EmptyUser> {
    path: "/proto.AuthService/DeleteAuth";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.EmptyUser>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.EmptyUser>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.EmptyUser>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.EmptyUser>;
}

export const AuthServiceService: IAuthServiceService;

export interface IAuthServiceServer extends grpc.UntypedServiceImplementation {
    getAuth: grpc.handleUnaryCall<proto_user_apiv1_pb.EmptyUser, proto_user_apiv1_pb.AuthResponse>;
    createAuth: grpc.handleUnaryCall<proto_user_apiv1_pb.CreateAuthRequest, proto_user_apiv1_pb.AuthResponse>;
    updateAuthEmail: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAuthEmailRequest, proto_user_apiv1_pb.AuthResponse>;
    updateAuthPassword: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAuthPasswordRequest, proto_user_apiv1_pb.AuthResponse>;
    updateAuthProfile: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAuthProfileRequest, proto_user_apiv1_pb.AuthResponse>;
    updateAuthAddress: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAuthAddressRequest, proto_user_apiv1_pb.AuthResponse>;
    uploadAuthThumbnail: handleClientStreamingCall<proto_user_apiv1_pb.UploadAuthThumbnailRequest, proto_user_apiv1_pb.AuthThumbnailResponse>;
    deleteAuth: grpc.handleUnaryCall<proto_user_apiv1_pb.EmptyUser, proto_user_apiv1_pb.EmptyUser>;
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
    uploadAuthThumbnail(callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAuthThumbnailRequest>;
    uploadAuthThumbnail(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAuthThumbnailRequest>;
    uploadAuthThumbnail(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAuthThumbnailRequest>;
    uploadAuthThumbnail(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAuthThumbnailRequest>;
    deleteAuth(request: proto_user_apiv1_pb.EmptyUser, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.EmptyUser) => void): grpc.ClientUnaryCall;
    deleteAuth(request: proto_user_apiv1_pb.EmptyUser, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.EmptyUser) => void): grpc.ClientUnaryCall;
    deleteAuth(request: proto_user_apiv1_pb.EmptyUser, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.EmptyUser) => void): grpc.ClientUnaryCall;
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
    public uploadAuthThumbnail(callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAuthThumbnailRequest>;
    public uploadAuthThumbnail(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAuthThumbnailRequest>;
    public uploadAuthThumbnail(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAuthThumbnailRequest>;
    public uploadAuthThumbnail(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAuthThumbnailRequest>;
    public deleteAuth(request: proto_user_apiv1_pb.EmptyUser, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.EmptyUser) => void): grpc.ClientUnaryCall;
    public deleteAuth(request: proto_user_apiv1_pb.EmptyUser, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.EmptyUser) => void): grpc.ClientUnaryCall;
    public deleteAuth(request: proto_user_apiv1_pb.EmptyUser, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.EmptyUser) => void): grpc.ClientUnaryCall;
}

interface IAdminServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listAdmin: IAdminServiceService_IListAdmin;
    searchAdmin: IAdminServiceService_ISearchAdmin;
    getAdmin: IAdminServiceService_IGetAdmin;
    createAdmin: IAdminServiceService_ICreateAdmin;
    updateAdminRole: IAdminServiceService_IUpdateAdminRole;
    updateAdminPassword: IAdminServiceService_IUpdateAdminPassword;
    updateAdminProfile: IAdminServiceService_IUpdateAdminProfile;
    uploadAdminThumbnail: IAdminServiceService_IUploadAdminThumbnail;
}

interface IAdminServiceService_IListAdmin extends grpc.MethodDefinition<proto_user_apiv1_pb.ListAdminRequest, proto_user_apiv1_pb.AdminListResponse> {
    path: "/proto.AdminService/ListAdmin";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.ListAdminRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.ListAdminRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AdminListResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AdminListResponse>;
}
interface IAdminServiceService_ISearchAdmin extends grpc.MethodDefinition<proto_user_apiv1_pb.SearchAdminRequest, proto_user_apiv1_pb.AdminListResponse> {
    path: "/proto.AdminService/SearchAdmin";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.SearchAdminRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.SearchAdminRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AdminListResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AdminListResponse>;
}
interface IAdminServiceService_IGetAdmin extends grpc.MethodDefinition<proto_user_apiv1_pb.GetAdminRequest, proto_user_apiv1_pb.AdminResponse> {
    path: "/proto.AdminService/GetAdmin";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.GetAdminRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.GetAdminRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AdminResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AdminResponse>;
}
interface IAdminServiceService_ICreateAdmin extends grpc.MethodDefinition<proto_user_apiv1_pb.CreateAdminRequest, proto_user_apiv1_pb.AdminResponse> {
    path: "/proto.AdminService/CreateAdmin";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.CreateAdminRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.CreateAdminRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AdminResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AdminResponse>;
}
interface IAdminServiceService_IUpdateAdminRole extends grpc.MethodDefinition<proto_user_apiv1_pb.UpdateAdminRoleRequest, proto_user_apiv1_pb.AdminResponse> {
    path: "/proto.AdminService/UpdateAdminRole";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.UpdateAdminRoleRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.UpdateAdminRoleRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AdminResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AdminResponse>;
}
interface IAdminServiceService_IUpdateAdminPassword extends grpc.MethodDefinition<proto_user_apiv1_pb.UpdateAdminPasswordRequest, proto_user_apiv1_pb.AdminResponse> {
    path: "/proto.AdminService/UpdateAdminPassword";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.UpdateAdminPasswordRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.UpdateAdminPasswordRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AdminResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AdminResponse>;
}
interface IAdminServiceService_IUpdateAdminProfile extends grpc.MethodDefinition<proto_user_apiv1_pb.UpdateAdminProfileRequest, proto_user_apiv1_pb.AdminResponse> {
    path: "/proto.AdminService/UpdateAdminProfile";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.UpdateAdminProfileRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.UpdateAdminProfileRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AdminResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AdminResponse>;
}
interface IAdminServiceService_IUploadAdminThumbnail extends grpc.MethodDefinition<proto_user_apiv1_pb.UploadAdminThumbnailRequest, proto_user_apiv1_pb.AdminThumbnailResponse> {
    path: "/proto.AdminService/UploadAdminThumbnail";
    requestStream: true;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.UploadAdminThumbnailRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.UploadAdminThumbnailRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.AdminThumbnailResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.AdminThumbnailResponse>;
}

export const AdminServiceService: IAdminServiceService;

export interface IAdminServiceServer extends grpc.UntypedServiceImplementation {
    listAdmin: grpc.handleUnaryCall<proto_user_apiv1_pb.ListAdminRequest, proto_user_apiv1_pb.AdminListResponse>;
    searchAdmin: grpc.handleUnaryCall<proto_user_apiv1_pb.SearchAdminRequest, proto_user_apiv1_pb.AdminListResponse>;
    getAdmin: grpc.handleUnaryCall<proto_user_apiv1_pb.GetAdminRequest, proto_user_apiv1_pb.AdminResponse>;
    createAdmin: grpc.handleUnaryCall<proto_user_apiv1_pb.CreateAdminRequest, proto_user_apiv1_pb.AdminResponse>;
    updateAdminRole: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAdminRoleRequest, proto_user_apiv1_pb.AdminResponse>;
    updateAdminPassword: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAdminPasswordRequest, proto_user_apiv1_pb.AdminResponse>;
    updateAdminProfile: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAdminProfileRequest, proto_user_apiv1_pb.AdminResponse>;
    uploadAdminThumbnail: handleClientStreamingCall<proto_user_apiv1_pb.UploadAdminThumbnailRequest, proto_user_apiv1_pb.AdminThumbnailResponse>;
}

export interface IAdminServiceClient {
    listAdmin(request: proto_user_apiv1_pb.ListAdminRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminListResponse) => void): grpc.ClientUnaryCall;
    listAdmin(request: proto_user_apiv1_pb.ListAdminRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminListResponse) => void): grpc.ClientUnaryCall;
    listAdmin(request: proto_user_apiv1_pb.ListAdminRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminListResponse) => void): grpc.ClientUnaryCall;
    searchAdmin(request: proto_user_apiv1_pb.SearchAdminRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminListResponse) => void): grpc.ClientUnaryCall;
    searchAdmin(request: proto_user_apiv1_pb.SearchAdminRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminListResponse) => void): grpc.ClientUnaryCall;
    searchAdmin(request: proto_user_apiv1_pb.SearchAdminRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminListResponse) => void): grpc.ClientUnaryCall;
    getAdmin(request: proto_user_apiv1_pb.GetAdminRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    getAdmin(request: proto_user_apiv1_pb.GetAdminRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    getAdmin(request: proto_user_apiv1_pb.GetAdminRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    createAdmin(request: proto_user_apiv1_pb.CreateAdminRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    createAdmin(request: proto_user_apiv1_pb.CreateAdminRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    createAdmin(request: proto_user_apiv1_pb.CreateAdminRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    updateAdminRole(request: proto_user_apiv1_pb.UpdateAdminRoleRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    updateAdminRole(request: proto_user_apiv1_pb.UpdateAdminRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    updateAdminRole(request: proto_user_apiv1_pb.UpdateAdminRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    updateAdminPassword(request: proto_user_apiv1_pb.UpdateAdminPasswordRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    updateAdminPassword(request: proto_user_apiv1_pb.UpdateAdminPasswordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    updateAdminPassword(request: proto_user_apiv1_pb.UpdateAdminPasswordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    updateAdminProfile(request: proto_user_apiv1_pb.UpdateAdminProfileRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    updateAdminProfile(request: proto_user_apiv1_pb.UpdateAdminProfileRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    updateAdminProfile(request: proto_user_apiv1_pb.UpdateAdminProfileRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    uploadAdminThumbnail(callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAdminThumbnailRequest>;
    uploadAdminThumbnail(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAdminThumbnailRequest>;
    uploadAdminThumbnail(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAdminThumbnailRequest>;
    uploadAdminThumbnail(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAdminThumbnailRequest>;
}

export class AdminServiceClient extends grpc.Client implements IAdminServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public listAdmin(request: proto_user_apiv1_pb.ListAdminRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminListResponse) => void): grpc.ClientUnaryCall;
    public listAdmin(request: proto_user_apiv1_pb.ListAdminRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminListResponse) => void): grpc.ClientUnaryCall;
    public listAdmin(request: proto_user_apiv1_pb.ListAdminRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminListResponse) => void): grpc.ClientUnaryCall;
    public searchAdmin(request: proto_user_apiv1_pb.SearchAdminRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminListResponse) => void): grpc.ClientUnaryCall;
    public searchAdmin(request: proto_user_apiv1_pb.SearchAdminRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminListResponse) => void): grpc.ClientUnaryCall;
    public searchAdmin(request: proto_user_apiv1_pb.SearchAdminRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminListResponse) => void): grpc.ClientUnaryCall;
    public getAdmin(request: proto_user_apiv1_pb.GetAdminRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public getAdmin(request: proto_user_apiv1_pb.GetAdminRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public getAdmin(request: proto_user_apiv1_pb.GetAdminRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public createAdmin(request: proto_user_apiv1_pb.CreateAdminRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public createAdmin(request: proto_user_apiv1_pb.CreateAdminRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public createAdmin(request: proto_user_apiv1_pb.CreateAdminRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public updateAdminRole(request: proto_user_apiv1_pb.UpdateAdminRoleRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public updateAdminRole(request: proto_user_apiv1_pb.UpdateAdminRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public updateAdminRole(request: proto_user_apiv1_pb.UpdateAdminRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public updateAdminPassword(request: proto_user_apiv1_pb.UpdateAdminPasswordRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public updateAdminPassword(request: proto_user_apiv1_pb.UpdateAdminPasswordRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public updateAdminPassword(request: proto_user_apiv1_pb.UpdateAdminPasswordRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public updateAdminProfile(request: proto_user_apiv1_pb.UpdateAdminProfileRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public updateAdminProfile(request: proto_user_apiv1_pb.UpdateAdminProfileRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public updateAdminProfile(request: proto_user_apiv1_pb.UpdateAdminProfileRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public uploadAdminThumbnail(callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAdminThumbnailRequest>;
    public uploadAdminThumbnail(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAdminThumbnailRequest>;
    public uploadAdminThumbnail(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAdminThumbnailRequest>;
    public uploadAdminThumbnail(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminThumbnailResponse) => void): grpc.ClientWritableStream<proto_user_apiv1_pb.UploadAdminThumbnailRequest>;
}

interface IUserServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listFollow: IUserServiceService_IListFollow;
    listFollower: IUserServiceService_IListFollower;
    getUserProfile: IUserServiceService_IGetUserProfile;
    registerFollow: IUserServiceService_IRegisterFollow;
    unregisterFollow: IUserServiceService_IUnregisterFollow;
}

interface IUserServiceService_IListFollow extends grpc.MethodDefinition<proto_user_apiv1_pb.ListFollowRequest, proto_user_apiv1_pb.FollowListResponse> {
    path: "/proto.UserService/ListFollow";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.ListFollowRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.ListFollowRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.FollowListResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.FollowListResponse>;
}
interface IUserServiceService_IListFollower extends grpc.MethodDefinition<proto_user_apiv1_pb.ListFollowerRequest, proto_user_apiv1_pb.FollowerListResponse> {
    path: "/proto.UserService/ListFollower";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.ListFollowerRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.ListFollowerRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.FollowerListResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.FollowerListResponse>;
}
interface IUserServiceService_IGetUserProfile extends grpc.MethodDefinition<proto_user_apiv1_pb.GetUserProfileRequest, proto_user_apiv1_pb.UserProfileResponse> {
    path: "/proto.UserService/GetUserProfile";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.GetUserProfileRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.GetUserProfileRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.UserProfileResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.UserProfileResponse>;
}
interface IUserServiceService_IRegisterFollow extends grpc.MethodDefinition<proto_user_apiv1_pb.RegisterFollowRequest, proto_user_apiv1_pb.UserProfileResponse> {
    path: "/proto.UserService/RegisterFollow";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.RegisterFollowRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.RegisterFollowRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.UserProfileResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.UserProfileResponse>;
}
interface IUserServiceService_IUnregisterFollow extends grpc.MethodDefinition<proto_user_apiv1_pb.UnregisterFollowRequest, proto_user_apiv1_pb.UserProfileResponse> {
    path: "/proto.UserService/UnregisterFollow";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.UnregisterFollowRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.UnregisterFollowRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.UserProfileResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.UserProfileResponse>;
}

export const UserServiceService: IUserServiceService;

export interface IUserServiceServer extends grpc.UntypedServiceImplementation {
    listFollow: grpc.handleUnaryCall<proto_user_apiv1_pb.ListFollowRequest, proto_user_apiv1_pb.FollowListResponse>;
    listFollower: grpc.handleUnaryCall<proto_user_apiv1_pb.ListFollowerRequest, proto_user_apiv1_pb.FollowerListResponse>;
    getUserProfile: grpc.handleUnaryCall<proto_user_apiv1_pb.GetUserProfileRequest, proto_user_apiv1_pb.UserProfileResponse>;
    registerFollow: grpc.handleUnaryCall<proto_user_apiv1_pb.RegisterFollowRequest, proto_user_apiv1_pb.UserProfileResponse>;
    unregisterFollow: grpc.handleUnaryCall<proto_user_apiv1_pb.UnregisterFollowRequest, proto_user_apiv1_pb.UserProfileResponse>;
}

export interface IUserServiceClient {
    listFollow(request: proto_user_apiv1_pb.ListFollowRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowListResponse) => void): grpc.ClientUnaryCall;
    listFollow(request: proto_user_apiv1_pb.ListFollowRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowListResponse) => void): grpc.ClientUnaryCall;
    listFollow(request: proto_user_apiv1_pb.ListFollowRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowListResponse) => void): grpc.ClientUnaryCall;
    listFollower(request: proto_user_apiv1_pb.ListFollowerRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowerListResponse) => void): grpc.ClientUnaryCall;
    listFollower(request: proto_user_apiv1_pb.ListFollowerRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowerListResponse) => void): grpc.ClientUnaryCall;
    listFollower(request: proto_user_apiv1_pb.ListFollowerRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowerListResponse) => void): grpc.ClientUnaryCall;
    getUserProfile(request: proto_user_apiv1_pb.GetUserProfileRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    getUserProfile(request: proto_user_apiv1_pb.GetUserProfileRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    getUserProfile(request: proto_user_apiv1_pb.GetUserProfileRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    registerFollow(request: proto_user_apiv1_pb.RegisterFollowRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    registerFollow(request: proto_user_apiv1_pb.RegisterFollowRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    registerFollow(request: proto_user_apiv1_pb.RegisterFollowRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    unregisterFollow(request: proto_user_apiv1_pb.UnregisterFollowRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    unregisterFollow(request: proto_user_apiv1_pb.UnregisterFollowRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    unregisterFollow(request: proto_user_apiv1_pb.UnregisterFollowRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
}

export class UserServiceClient extends grpc.Client implements IUserServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public listFollow(request: proto_user_apiv1_pb.ListFollowRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowListResponse) => void): grpc.ClientUnaryCall;
    public listFollow(request: proto_user_apiv1_pb.ListFollowRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowListResponse) => void): grpc.ClientUnaryCall;
    public listFollow(request: proto_user_apiv1_pb.ListFollowRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowListResponse) => void): grpc.ClientUnaryCall;
    public listFollower(request: proto_user_apiv1_pb.ListFollowerRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowerListResponse) => void): grpc.ClientUnaryCall;
    public listFollower(request: proto_user_apiv1_pb.ListFollowerRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowerListResponse) => void): grpc.ClientUnaryCall;
    public listFollower(request: proto_user_apiv1_pb.ListFollowerRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowerListResponse) => void): grpc.ClientUnaryCall;
    public getUserProfile(request: proto_user_apiv1_pb.GetUserProfileRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    public getUserProfile(request: proto_user_apiv1_pb.GetUserProfileRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    public getUserProfile(request: proto_user_apiv1_pb.GetUserProfileRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    public registerFollow(request: proto_user_apiv1_pb.RegisterFollowRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    public registerFollow(request: proto_user_apiv1_pb.RegisterFollowRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    public registerFollow(request: proto_user_apiv1_pb.RegisterFollowRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    public unregisterFollow(request: proto_user_apiv1_pb.UnregisterFollowRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    public unregisterFollow(request: proto_user_apiv1_pb.UnregisterFollowRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
    public unregisterFollow(request: proto_user_apiv1_pb.UnregisterFollowRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserProfileResponse) => void): grpc.ClientUnaryCall;
}
