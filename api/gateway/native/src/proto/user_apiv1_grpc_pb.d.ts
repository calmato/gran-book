// package: proto
// file: proto/user_apiv1.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
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
    registerAuthDevice: IAuthServiceService_IRegisterAuthDevice;
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
interface IAuthServiceService_IRegisterAuthDevice extends grpc.MethodDefinition<proto_user_apiv1_pb.RegisterAuthDeviceRequest, proto_user_apiv1_pb.AuthResponse> {
    path: "/proto.AuthService/RegisterAuthDevice";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.RegisterAuthDeviceRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.RegisterAuthDeviceRequest>;
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
    uploadAuthThumbnail: grpc.handleClientStreamingCall<proto_user_apiv1_pb.UploadAuthThumbnailRequest, proto_user_apiv1_pb.AuthThumbnailResponse>;
    deleteAuth: grpc.handleUnaryCall<proto_user_apiv1_pb.EmptyUser, proto_user_apiv1_pb.EmptyUser>;
    registerAuthDevice: grpc.handleUnaryCall<proto_user_apiv1_pb.RegisterAuthDeviceRequest, proto_user_apiv1_pb.AuthResponse>;
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
    registerAuthDevice(request: proto_user_apiv1_pb.RegisterAuthDeviceRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    registerAuthDevice(request: proto_user_apiv1_pb.RegisterAuthDeviceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    registerAuthDevice(request: proto_user_apiv1_pb.RegisterAuthDeviceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
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
    public registerAuthDevice(request: proto_user_apiv1_pb.RegisterAuthDeviceRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public registerAuthDevice(request: proto_user_apiv1_pb.RegisterAuthDeviceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public registerAuthDevice(request: proto_user_apiv1_pb.RegisterAuthDeviceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AuthResponse) => void): grpc.ClientUnaryCall;
}

interface IAdminServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listAdmin: IAdminServiceService_IListAdmin;
    searchAdmin: IAdminServiceService_ISearchAdmin;
    getAdmin: IAdminServiceService_IGetAdmin;
    createAdmin: IAdminServiceService_ICreateAdmin;
    updateAdminContact: IAdminServiceService_IUpdateAdminContact;
    updateAdminPassword: IAdminServiceService_IUpdateAdminPassword;
    updateAdminProfile: IAdminServiceService_IUpdateAdminProfile;
    uploadAdminThumbnail: IAdminServiceService_IUploadAdminThumbnail;
    deleteAdmin: IAdminServiceService_IDeleteAdmin;
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
interface IAdminServiceService_IUpdateAdminContact extends grpc.MethodDefinition<proto_user_apiv1_pb.UpdateAdminContactRequest, proto_user_apiv1_pb.AdminResponse> {
    path: "/proto.AdminService/UpdateAdminContact";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.UpdateAdminContactRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.UpdateAdminContactRequest>;
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
interface IAdminServiceService_IDeleteAdmin extends grpc.MethodDefinition<proto_user_apiv1_pb.DeleteAdminRequest, proto_user_apiv1_pb.EmptyUser> {
    path: "/proto.AdminService/DeleteAdmin";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.DeleteAdminRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.DeleteAdminRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.EmptyUser>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.EmptyUser>;
}

export const AdminServiceService: IAdminServiceService;

export interface IAdminServiceServer extends grpc.UntypedServiceImplementation {
    listAdmin: grpc.handleUnaryCall<proto_user_apiv1_pb.ListAdminRequest, proto_user_apiv1_pb.AdminListResponse>;
    searchAdmin: grpc.handleUnaryCall<proto_user_apiv1_pb.SearchAdminRequest, proto_user_apiv1_pb.AdminListResponse>;
    getAdmin: grpc.handleUnaryCall<proto_user_apiv1_pb.GetAdminRequest, proto_user_apiv1_pb.AdminResponse>;
    createAdmin: grpc.handleUnaryCall<proto_user_apiv1_pb.CreateAdminRequest, proto_user_apiv1_pb.AdminResponse>;
    updateAdminContact: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAdminContactRequest, proto_user_apiv1_pb.AdminResponse>;
    updateAdminPassword: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAdminPasswordRequest, proto_user_apiv1_pb.AdminResponse>;
    updateAdminProfile: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAdminProfileRequest, proto_user_apiv1_pb.AdminResponse>;
    uploadAdminThumbnail: grpc.handleClientStreamingCall<proto_user_apiv1_pb.UploadAdminThumbnailRequest, proto_user_apiv1_pb.AdminThumbnailResponse>;
    deleteAdmin: grpc.handleUnaryCall<proto_user_apiv1_pb.DeleteAdminRequest, proto_user_apiv1_pb.EmptyUser>;
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
    updateAdminContact(request: proto_user_apiv1_pb.UpdateAdminContactRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    updateAdminContact(request: proto_user_apiv1_pb.UpdateAdminContactRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    updateAdminContact(request: proto_user_apiv1_pb.UpdateAdminContactRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
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
    deleteAdmin(request: proto_user_apiv1_pb.DeleteAdminRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.EmptyUser) => void): grpc.ClientUnaryCall;
    deleteAdmin(request: proto_user_apiv1_pb.DeleteAdminRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.EmptyUser) => void): grpc.ClientUnaryCall;
    deleteAdmin(request: proto_user_apiv1_pb.DeleteAdminRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.EmptyUser) => void): grpc.ClientUnaryCall;
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
    public updateAdminContact(request: proto_user_apiv1_pb.UpdateAdminContactRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public updateAdminContact(request: proto_user_apiv1_pb.UpdateAdminContactRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
    public updateAdminContact(request: proto_user_apiv1_pb.UpdateAdminContactRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.AdminResponse) => void): grpc.ClientUnaryCall;
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
    public deleteAdmin(request: proto_user_apiv1_pb.DeleteAdminRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.EmptyUser) => void): grpc.ClientUnaryCall;
    public deleteAdmin(request: proto_user_apiv1_pb.DeleteAdminRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.EmptyUser) => void): grpc.ClientUnaryCall;
    public deleteAdmin(request: proto_user_apiv1_pb.DeleteAdminRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.EmptyUser) => void): grpc.ClientUnaryCall;
}

interface IUserServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listUser: IUserServiceService_IListUser;
    listUserByUserIds: IUserServiceService_IListUserByUserIds;
    listFollow: IUserServiceService_IListFollow;
    listFollower: IUserServiceService_IListFollower;
    searchUser: IUserServiceService_ISearchUser;
    getUser: IUserServiceService_IGetUser;
    getUserProfile: IUserServiceService_IGetUserProfile;
    registerFollow: IUserServiceService_IRegisterFollow;
    unregisterFollow: IUserServiceService_IUnregisterFollow;
}

interface IUserServiceService_IListUser extends grpc.MethodDefinition<proto_user_apiv1_pb.ListUserRequest, proto_user_apiv1_pb.UserListResponse> {
    path: "/proto.UserService/ListUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.ListUserRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.ListUserRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.UserListResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.UserListResponse>;
}
interface IUserServiceService_IListUserByUserIds extends grpc.MethodDefinition<proto_user_apiv1_pb.ListUserByUserIdsRequest, proto_user_apiv1_pb.UserListResponse> {
    path: "/proto.UserService/ListUserByUserIds";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.ListUserByUserIdsRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.ListUserByUserIdsRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.UserListResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.UserListResponse>;
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
interface IUserServiceService_ISearchUser extends grpc.MethodDefinition<proto_user_apiv1_pb.SearchUserRequest, proto_user_apiv1_pb.UserListResponse> {
    path: "/proto.UserService/SearchUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.SearchUserRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.SearchUserRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.UserListResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.UserListResponse>;
}
interface IUserServiceService_IGetUser extends grpc.MethodDefinition<proto_user_apiv1_pb.GetUserRequest, proto_user_apiv1_pb.UserResponse> {
    path: "/proto.UserService/GetUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.GetUserRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.GetUserRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.UserResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.UserResponse>;
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
    listUser: grpc.handleUnaryCall<proto_user_apiv1_pb.ListUserRequest, proto_user_apiv1_pb.UserListResponse>;
    listUserByUserIds: grpc.handleUnaryCall<proto_user_apiv1_pb.ListUserByUserIdsRequest, proto_user_apiv1_pb.UserListResponse>;
    listFollow: grpc.handleUnaryCall<proto_user_apiv1_pb.ListFollowRequest, proto_user_apiv1_pb.FollowListResponse>;
    listFollower: grpc.handleUnaryCall<proto_user_apiv1_pb.ListFollowerRequest, proto_user_apiv1_pb.FollowerListResponse>;
    searchUser: grpc.handleUnaryCall<proto_user_apiv1_pb.SearchUserRequest, proto_user_apiv1_pb.UserListResponse>;
    getUser: grpc.handleUnaryCall<proto_user_apiv1_pb.GetUserRequest, proto_user_apiv1_pb.UserResponse>;
    getUserProfile: grpc.handleUnaryCall<proto_user_apiv1_pb.GetUserProfileRequest, proto_user_apiv1_pb.UserProfileResponse>;
    registerFollow: grpc.handleUnaryCall<proto_user_apiv1_pb.RegisterFollowRequest, proto_user_apiv1_pb.UserProfileResponse>;
    unregisterFollow: grpc.handleUnaryCall<proto_user_apiv1_pb.UnregisterFollowRequest, proto_user_apiv1_pb.UserProfileResponse>;
}

export interface IUserServiceClient {
    listUser(request: proto_user_apiv1_pb.ListUserRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    listUser(request: proto_user_apiv1_pb.ListUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    listUser(request: proto_user_apiv1_pb.ListUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    listUserByUserIds(request: proto_user_apiv1_pb.ListUserByUserIdsRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    listUserByUserIds(request: proto_user_apiv1_pb.ListUserByUserIdsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    listUserByUserIds(request: proto_user_apiv1_pb.ListUserByUserIdsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    listFollow(request: proto_user_apiv1_pb.ListFollowRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowListResponse) => void): grpc.ClientUnaryCall;
    listFollow(request: proto_user_apiv1_pb.ListFollowRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowListResponse) => void): grpc.ClientUnaryCall;
    listFollow(request: proto_user_apiv1_pb.ListFollowRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowListResponse) => void): grpc.ClientUnaryCall;
    listFollower(request: proto_user_apiv1_pb.ListFollowerRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowerListResponse) => void): grpc.ClientUnaryCall;
    listFollower(request: proto_user_apiv1_pb.ListFollowerRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowerListResponse) => void): grpc.ClientUnaryCall;
    listFollower(request: proto_user_apiv1_pb.ListFollowerRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowerListResponse) => void): grpc.ClientUnaryCall;
    searchUser(request: proto_user_apiv1_pb.SearchUserRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    searchUser(request: proto_user_apiv1_pb.SearchUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    searchUser(request: proto_user_apiv1_pb.SearchUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    getUser(request: proto_user_apiv1_pb.GetUserRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserResponse) => void): grpc.ClientUnaryCall;
    getUser(request: proto_user_apiv1_pb.GetUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserResponse) => void): grpc.ClientUnaryCall;
    getUser(request: proto_user_apiv1_pb.GetUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserResponse) => void): grpc.ClientUnaryCall;
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
    public listUser(request: proto_user_apiv1_pb.ListUserRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    public listUser(request: proto_user_apiv1_pb.ListUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    public listUser(request: proto_user_apiv1_pb.ListUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    public listUserByUserIds(request: proto_user_apiv1_pb.ListUserByUserIdsRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    public listUserByUserIds(request: proto_user_apiv1_pb.ListUserByUserIdsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    public listUserByUserIds(request: proto_user_apiv1_pb.ListUserByUserIdsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    public listFollow(request: proto_user_apiv1_pb.ListFollowRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowListResponse) => void): grpc.ClientUnaryCall;
    public listFollow(request: proto_user_apiv1_pb.ListFollowRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowListResponse) => void): grpc.ClientUnaryCall;
    public listFollow(request: proto_user_apiv1_pb.ListFollowRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowListResponse) => void): grpc.ClientUnaryCall;
    public listFollower(request: proto_user_apiv1_pb.ListFollowerRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowerListResponse) => void): grpc.ClientUnaryCall;
    public listFollower(request: proto_user_apiv1_pb.ListFollowerRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowerListResponse) => void): grpc.ClientUnaryCall;
    public listFollower(request: proto_user_apiv1_pb.ListFollowerRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.FollowerListResponse) => void): grpc.ClientUnaryCall;
    public searchUser(request: proto_user_apiv1_pb.SearchUserRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    public searchUser(request: proto_user_apiv1_pb.SearchUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    public searchUser(request: proto_user_apiv1_pb.SearchUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserListResponse) => void): grpc.ClientUnaryCall;
    public getUser(request: proto_user_apiv1_pb.GetUserRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public getUser(request: proto_user_apiv1_pb.GetUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public getUser(request: proto_user_apiv1_pb.GetUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.UserResponse) => void): grpc.ClientUnaryCall;
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

interface IChatServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listRoom: IChatServiceService_IListRoom;
    createRoom: IChatServiceService_ICreateRoom;
    createMessage: IChatServiceService_ICreateMessage;
    uploadImage: IChatServiceService_IUploadImage;
}

interface IChatServiceService_IListRoom extends grpc.MethodDefinition<proto_user_apiv1_pb.ListChatRoomRequest, proto_user_apiv1_pb.ChatRoomListResponse> {
    path: "/proto.ChatService/ListRoom";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.ListChatRoomRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.ListChatRoomRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.ChatRoomListResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.ChatRoomListResponse>;
}
interface IChatServiceService_ICreateRoom extends grpc.MethodDefinition<proto_user_apiv1_pb.CreateChatRoomRequest, proto_user_apiv1_pb.ChatRoomResponse> {
    path: "/proto.ChatService/CreateRoom";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.CreateChatRoomRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.CreateChatRoomRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.ChatRoomResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.ChatRoomResponse>;
}
interface IChatServiceService_ICreateMessage extends grpc.MethodDefinition<proto_user_apiv1_pb.CreateChatMessageRequest, proto_user_apiv1_pb.ChatMessageResponse> {
    path: "/proto.ChatService/CreateMessage";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.CreateChatMessageRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.CreateChatMessageRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.ChatMessageResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.ChatMessageResponse>;
}
interface IChatServiceService_IUploadImage extends grpc.MethodDefinition<proto_user_apiv1_pb.UploadChatImageRequest, proto_user_apiv1_pb.ChatMessageResponse> {
    path: "/proto.ChatService/UploadImage";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_user_apiv1_pb.UploadChatImageRequest>;
    requestDeserialize: grpc.deserialize<proto_user_apiv1_pb.UploadChatImageRequest>;
    responseSerialize: grpc.serialize<proto_user_apiv1_pb.ChatMessageResponse>;
    responseDeserialize: grpc.deserialize<proto_user_apiv1_pb.ChatMessageResponse>;
}

export const ChatServiceService: IChatServiceService;

export interface IChatServiceServer extends grpc.UntypedServiceImplementation {
    listRoom: grpc.handleUnaryCall<proto_user_apiv1_pb.ListChatRoomRequest, proto_user_apiv1_pb.ChatRoomListResponse>;
    createRoom: grpc.handleUnaryCall<proto_user_apiv1_pb.CreateChatRoomRequest, proto_user_apiv1_pb.ChatRoomResponse>;
    createMessage: grpc.handleUnaryCall<proto_user_apiv1_pb.CreateChatMessageRequest, proto_user_apiv1_pb.ChatMessageResponse>;
    uploadImage: grpc.handleUnaryCall<proto_user_apiv1_pb.UploadChatImageRequest, proto_user_apiv1_pb.ChatMessageResponse>;
}

export interface IChatServiceClient {
    listRoom(request: proto_user_apiv1_pb.ListChatRoomRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatRoomListResponse) => void): grpc.ClientUnaryCall;
    listRoom(request: proto_user_apiv1_pb.ListChatRoomRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatRoomListResponse) => void): grpc.ClientUnaryCall;
    listRoom(request: proto_user_apiv1_pb.ListChatRoomRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatRoomListResponse) => void): grpc.ClientUnaryCall;
    createRoom(request: proto_user_apiv1_pb.CreateChatRoomRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatRoomResponse) => void): grpc.ClientUnaryCall;
    createRoom(request: proto_user_apiv1_pb.CreateChatRoomRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatRoomResponse) => void): grpc.ClientUnaryCall;
    createRoom(request: proto_user_apiv1_pb.CreateChatRoomRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatRoomResponse) => void): grpc.ClientUnaryCall;
    createMessage(request: proto_user_apiv1_pb.CreateChatMessageRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatMessageResponse) => void): grpc.ClientUnaryCall;
    createMessage(request: proto_user_apiv1_pb.CreateChatMessageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatMessageResponse) => void): grpc.ClientUnaryCall;
    createMessage(request: proto_user_apiv1_pb.CreateChatMessageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatMessageResponse) => void): grpc.ClientUnaryCall;
    uploadImage(request: proto_user_apiv1_pb.UploadChatImageRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatMessageResponse) => void): grpc.ClientUnaryCall;
    uploadImage(request: proto_user_apiv1_pb.UploadChatImageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatMessageResponse) => void): grpc.ClientUnaryCall;
    uploadImage(request: proto_user_apiv1_pb.UploadChatImageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatMessageResponse) => void): grpc.ClientUnaryCall;
}

export class ChatServiceClient extends grpc.Client implements IChatServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public listRoom(request: proto_user_apiv1_pb.ListChatRoomRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatRoomListResponse) => void): grpc.ClientUnaryCall;
    public listRoom(request: proto_user_apiv1_pb.ListChatRoomRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatRoomListResponse) => void): grpc.ClientUnaryCall;
    public listRoom(request: proto_user_apiv1_pb.ListChatRoomRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatRoomListResponse) => void): grpc.ClientUnaryCall;
    public createRoom(request: proto_user_apiv1_pb.CreateChatRoomRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatRoomResponse) => void): grpc.ClientUnaryCall;
    public createRoom(request: proto_user_apiv1_pb.CreateChatRoomRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatRoomResponse) => void): grpc.ClientUnaryCall;
    public createRoom(request: proto_user_apiv1_pb.CreateChatRoomRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatRoomResponse) => void): grpc.ClientUnaryCall;
    public createMessage(request: proto_user_apiv1_pb.CreateChatMessageRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatMessageResponse) => void): grpc.ClientUnaryCall;
    public createMessage(request: proto_user_apiv1_pb.CreateChatMessageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatMessageResponse) => void): grpc.ClientUnaryCall;
    public createMessage(request: proto_user_apiv1_pb.CreateChatMessageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatMessageResponse) => void): grpc.ClientUnaryCall;
    public uploadImage(request: proto_user_apiv1_pb.UploadChatImageRequest, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatMessageResponse) => void): grpc.ClientUnaryCall;
    public uploadImage(request: proto_user_apiv1_pb.UploadChatImageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatMessageResponse) => void): grpc.ClientUnaryCall;
    public uploadImage(request: proto_user_apiv1_pb.UploadChatImageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_user_apiv1_pb.ChatMessageResponse) => void): grpc.ClientUnaryCall;
}
