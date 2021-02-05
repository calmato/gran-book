// package: proto
// file: proto/user_apiv1.proto

import * as jspb from "google-protobuf";

export class EmptyUser extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EmptyUser.AsObject;
  static toObject(includeInstance: boolean, msg: EmptyUser): EmptyUser.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EmptyUser, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EmptyUser;
  static deserializeBinaryFromReader(message: EmptyUser, reader: jspb.BinaryReader): EmptyUser;
}

export namespace EmptyUser {
  export type AsObject = {
  }
}

export class CreateAuthRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): void;

  getEmail(): string;
  setEmail(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  getPasswordConfirmation(): string;
  setPasswordConfirmation(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAuthRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAuthRequest): CreateAuthRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateAuthRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAuthRequest;
  static deserializeBinaryFromReader(message: CreateAuthRequest, reader: jspb.BinaryReader): CreateAuthRequest;
}

export namespace CreateAuthRequest {
  export type AsObject = {
    username: string,
    email: string,
    password: string,
    passwordConfirmation: string,
  }
}

export class UpdateAuthEmailRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateAuthEmailRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateAuthEmailRequest): UpdateAuthEmailRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateAuthEmailRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateAuthEmailRequest;
  static deserializeBinaryFromReader(message: UpdateAuthEmailRequest, reader: jspb.BinaryReader): UpdateAuthEmailRequest;
}

export namespace UpdateAuthEmailRequest {
  export type AsObject = {
    email: string,
  }
}

export class UpdateAuthPasswordRequest extends jspb.Message {
  getPassword(): string;
  setPassword(value: string): void;

  getPasswordConfirmation(): string;
  setPasswordConfirmation(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateAuthPasswordRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateAuthPasswordRequest): UpdateAuthPasswordRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateAuthPasswordRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateAuthPasswordRequest;
  static deserializeBinaryFromReader(message: UpdateAuthPasswordRequest, reader: jspb.BinaryReader): UpdateAuthPasswordRequest;
}

export namespace UpdateAuthPasswordRequest {
  export type AsObject = {
    password: string,
    passwordConfirmation: string,
  }
}

export class UpdateAuthProfileRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): void;

  getGender(): number;
  setGender(value: number): void;

  getThumbnail(): string;
  setThumbnail(value: string): void;

  getSelfIntroduction(): string;
  setSelfIntroduction(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateAuthProfileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateAuthProfileRequest): UpdateAuthProfileRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateAuthProfileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateAuthProfileRequest;
  static deserializeBinaryFromReader(message: UpdateAuthProfileRequest, reader: jspb.BinaryReader): UpdateAuthProfileRequest;
}

export namespace UpdateAuthProfileRequest {
  export type AsObject = {
    username: string,
    gender: number,
    thumbnail: string,
    selfIntroduction: string,
  }
}

export class UpdateAuthAddressRequest extends jspb.Message {
  getLastName(): string;
  setLastName(value: string): void;

  getFirstName(): string;
  setFirstName(value: string): void;

  getLastNameKana(): string;
  setLastNameKana(value: string): void;

  getFirstNameKana(): string;
  setFirstNameKana(value: string): void;

  getPhoneNumber(): string;
  setPhoneNumber(value: string): void;

  getPostalCode(): string;
  setPostalCode(value: string): void;

  getPrefecture(): string;
  setPrefecture(value: string): void;

  getCity(): string;
  setCity(value: string): void;

  getAddressLine1(): string;
  setAddressLine1(value: string): void;

  getAddressLine2(): string;
  setAddressLine2(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateAuthAddressRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateAuthAddressRequest): UpdateAuthAddressRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateAuthAddressRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateAuthAddressRequest;
  static deserializeBinaryFromReader(message: UpdateAuthAddressRequest, reader: jspb.BinaryReader): UpdateAuthAddressRequest;
}

export namespace UpdateAuthAddressRequest {
  export type AsObject = {
    lastName: string,
    firstName: string,
    lastNameKana: string,
    firstNameKana: string,
    phoneNumber: string,
    postalCode: string,
    prefecture: string,
    city: string,
    addressLine1: string,
    addressLine2: string,
  }
}

export class AuthResponse extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getUsername(): string;
  setUsername(value: string): void;

  getGender(): number;
  setGender(value: number): void;

  getEmail(): string;
  setEmail(value: string): void;

  getPhoneNumber(): string;
  setPhoneNumber(value: string): void;

  getRole(): number;
  setRole(value: number): void;

  getThumbnailUrl(): string;
  setThumbnailUrl(value: string): void;

  getSelfIntroduction(): string;
  setSelfIntroduction(value: string): void;

  getLastName(): string;
  setLastName(value: string): void;

  getFirstName(): string;
  setFirstName(value: string): void;

  getLastNameKana(): string;
  setLastNameKana(value: string): void;

  getFirstNameKana(): string;
  setFirstNameKana(value: string): void;

  getPostalCode(): string;
  setPostalCode(value: string): void;

  getPrefecture(): string;
  setPrefecture(value: string): void;

  getCity(): string;
  setCity(value: string): void;

  getAddressLine1(): string;
  setAddressLine1(value: string): void;

  getAddressLine2(): string;
  setAddressLine2(value: string): void;

  getCreatedAt(): string;
  setCreatedAt(value: string): void;

  getUpdatedAt(): string;
  setUpdatedAt(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AuthResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AuthResponse): AuthResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AuthResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AuthResponse;
  static deserializeBinaryFromReader(message: AuthResponse, reader: jspb.BinaryReader): AuthResponse;
}

export namespace AuthResponse {
  export type AsObject = {
    id: string,
    username: string,
    gender: number,
    email: string,
    phoneNumber: string,
    role: number,
    thumbnailUrl: string,
    selfIntroduction: string,
    lastName: string,
    firstName: string,
    lastNameKana: string,
    firstNameKana: string,
    postalCode: string,
    prefecture: string,
    city: string,
    addressLine1: string,
    addressLine2: string,
    createdAt: string,
    updatedAt: string,
  }
}

