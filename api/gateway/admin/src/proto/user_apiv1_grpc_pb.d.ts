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

interface IAdminServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listAdmin: IAdminServiceService_IListAdmin;
    searchAdmin: IAdminServiceService_ISearchAdmin;
    getAdmin: IAdminServiceService_IGetAdmin;
    createAdmin: IAdminServiceService_ICreateAdmin;
    updateAdminRole: IAdminServiceService_IUpdateAdminRole;
    updateAdminPassword: IAdminServiceService_IUpdateAdminPassword;
    updateAdminProfile: IAdminServiceService_IUpdateAdminProfile;
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

export const AdminServiceService: IAdminServiceService;

export interface IAdminServiceServer extends grpc.UntypedServiceImplementation {
    listAdmin: grpc.handleUnaryCall<proto_user_apiv1_pb.ListAdminRequest, proto_user_apiv1_pb.AdminListResponse>;
    searchAdmin: grpc.handleUnaryCall<proto_user_apiv1_pb.SearchAdminRequest, proto_user_apiv1_pb.AdminListResponse>;
    getAdmin: grpc.handleUnaryCall<proto_user_apiv1_pb.GetAdminRequest, proto_user_apiv1_pb.AdminResponse>;
    createAdmin: grpc.handleUnaryCall<proto_user_apiv1_pb.CreateAdminRequest, proto_user_apiv1_pb.AdminResponse>;
    updateAdminRole: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAdminRoleRequest, proto_user_apiv1_pb.AdminResponse>;
    updateAdminPassword: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAdminPasswordRequest, proto_user_apiv1_pb.AdminResponse>;
    updateAdminProfile: grpc.handleUnaryCall<proto_user_apiv1_pb.UpdateAdminProfileRequest, proto_user_apiv1_pb.AdminResponse>;
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
}
