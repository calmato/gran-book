// package: proto
// file: proto/user_apiv1.proto

import * as proto_user_apiv1_pb from "../proto/user_apiv1_pb";
import {grpc} from "@improbable-eng/grpc-web";

type AuthServiceGetAuth = {
  readonly methodName: string;
  readonly service: typeof AuthService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_user_apiv1_pb.EmptyUser;
  readonly responseType: typeof proto_user_apiv1_pb.AuthResponse;
};

type AuthServiceCreateAuth = {
  readonly methodName: string;
  readonly service: typeof AuthService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_user_apiv1_pb.CreateAuthRequest;
  readonly responseType: typeof proto_user_apiv1_pb.AuthResponse;
};

type AuthServiceUpdateAuthEmail = {
  readonly methodName: string;
  readonly service: typeof AuthService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_user_apiv1_pb.UpdateAuthEmailRequest;
  readonly responseType: typeof proto_user_apiv1_pb.AuthResponse;
};

type AuthServiceUpdateAuthPassword = {
  readonly methodName: string;
  readonly service: typeof AuthService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_user_apiv1_pb.UpdateAuthPasswordRequest;
  readonly responseType: typeof proto_user_apiv1_pb.AuthResponse;
};

type AuthServiceUpdateAuthProfile = {
  readonly methodName: string;
  readonly service: typeof AuthService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_user_apiv1_pb.UpdateAuthProfileRequest;
  readonly responseType: typeof proto_user_apiv1_pb.AuthResponse;
};

type AuthServiceUpdateAuthAddress = {
  readonly methodName: string;
  readonly service: typeof AuthService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_user_apiv1_pb.UpdateAuthAddressRequest;
  readonly responseType: typeof proto_user_apiv1_pb.AuthResponse;
};

export class AuthService {
  static readonly serviceName: string;
  static readonly GetAuth: AuthServiceGetAuth;
  static readonly CreateAuth: AuthServiceCreateAuth;
  static readonly UpdateAuthEmail: AuthServiceUpdateAuthEmail;
  static readonly UpdateAuthPassword: AuthServiceUpdateAuthPassword;
  static readonly UpdateAuthProfile: AuthServiceUpdateAuthProfile;
  static readonly UpdateAuthAddress: AuthServiceUpdateAuthAddress;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class AuthServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getAuth(
    requestMessage: proto_user_apiv1_pb.EmptyUser,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_user_apiv1_pb.AuthResponse|null) => void
  ): UnaryResponse;
  getAuth(
    requestMessage: proto_user_apiv1_pb.EmptyUser,
    callback: (error: ServiceError|null, responseMessage: proto_user_apiv1_pb.AuthResponse|null) => void
  ): UnaryResponse;
  createAuth(
    requestMessage: proto_user_apiv1_pb.CreateAuthRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_user_apiv1_pb.AuthResponse|null) => void
  ): UnaryResponse;
  createAuth(
    requestMessage: proto_user_apiv1_pb.CreateAuthRequest,
    callback: (error: ServiceError|null, responseMessage: proto_user_apiv1_pb.AuthResponse|null) => void
  ): UnaryResponse;
  updateAuthEmail(
    requestMessage: proto_user_apiv1_pb.UpdateAuthEmailRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_user_apiv1_pb.AuthResponse|null) => void
  ): UnaryResponse;
  updateAuthEmail(
    requestMessage: proto_user_apiv1_pb.UpdateAuthEmailRequest,
    callback: (error: ServiceError|null, responseMessage: proto_user_apiv1_pb.AuthResponse|null) => void
  ): UnaryResponse;
  updateAuthPassword(
    requestMessage: proto_user_apiv1_pb.UpdateAuthPasswordRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_user_apiv1_pb.AuthResponse|null) => void
  ): UnaryResponse;
  updateAuthPassword(
    requestMessage: proto_user_apiv1_pb.UpdateAuthPasswordRequest,
    callback: (error: ServiceError|null, responseMessage: proto_user_apiv1_pb.AuthResponse|null) => void
  ): UnaryResponse;
  updateAuthProfile(
    requestMessage: proto_user_apiv1_pb.UpdateAuthProfileRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_user_apiv1_pb.AuthResponse|null) => void
  ): UnaryResponse;
  updateAuthProfile(
    requestMessage: proto_user_apiv1_pb.UpdateAuthProfileRequest,
    callback: (error: ServiceError|null, responseMessage: proto_user_apiv1_pb.AuthResponse|null) => void
  ): UnaryResponse;
  updateAuthAddress(
    requestMessage: proto_user_apiv1_pb.UpdateAuthAddressRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_user_apiv1_pb.AuthResponse|null) => void
  ): UnaryResponse;
  updateAuthAddress(
    requestMessage: proto_user_apiv1_pb.UpdateAuthAddressRequest,
    callback: (error: ServiceError|null, responseMessage: proto_user_apiv1_pb.AuthResponse|null) => void
  ): UnaryResponse;
}

