// package: proto
// file: proto/information_apiv1.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class HelloRequest extends jspb.Message { 
    getName(): string;
    setName(value: string): HelloRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): HelloRequest.AsObject;
    static toObject(includeInstance: boolean, msg: HelloRequest): HelloRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: HelloRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): HelloRequest;
    static deserializeBinaryFromReader(message: HelloRequest, reader: jspb.BinaryReader): HelloRequest;
}

export namespace HelloRequest {
    export type AsObject = {
        name: string,
    }
}

export class HelloResponse extends jspb.Message { 
    getMessage(): string;
    setMessage(value: string): HelloResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): HelloResponse.AsObject;
    static toObject(includeInstance: boolean, msg: HelloResponse): HelloResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: HelloResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): HelloResponse;
    static deserializeBinaryFromReader(message: HelloResponse, reader: jspb.BinaryReader): HelloResponse;
}

export namespace HelloResponse {
    export type AsObject = {
        message: string,
    }
}

export class CreateInquiryRequest extends jspb.Message { 
    getSenderId(): string;
    setSenderId(value: string): CreateInquiryRequest;
    getSubject(): string;
    setSubject(value: string): CreateInquiryRequest;
    getDescription(): string;
    setDescription(value: string): CreateInquiryRequest;
    getEmail(): string;
    setEmail(value: string): CreateInquiryRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateInquiryRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateInquiryRequest): CreateInquiryRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateInquiryRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateInquiryRequest;
    static deserializeBinaryFromReader(message: CreateInquiryRequest, reader: jspb.BinaryReader): CreateInquiryRequest;
}

export namespace CreateInquiryRequest {
    export type AsObject = {
        senderId: string,
        subject: string,
        description: string,
        email: string,
    }
}

export class InquiryResponse extends jspb.Message { 
    getInquiryId(): string;
    setInquiryId(value: string): InquiryResponse;
    getSenderId(): string;
    setSenderId(value: string): InquiryResponse;
    getAdminId(): string;
    setAdminId(value: string): InquiryResponse;
    getSubject(): string;
    setSubject(value: string): InquiryResponse;
    getDescription(): string;
    setDescription(value: string): InquiryResponse;
    getEmail(): string;
    setEmail(value: string): InquiryResponse;
    getIsReplied(): string;
    setIsReplied(value: string): InquiryResponse;
    getCreatedAt(): string;
    setCreatedAt(value: string): InquiryResponse;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): InquiryResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): InquiryResponse.AsObject;
    static toObject(includeInstance: boolean, msg: InquiryResponse): InquiryResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: InquiryResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): InquiryResponse;
    static deserializeBinaryFromReader(message: InquiryResponse, reader: jspb.BinaryReader): InquiryResponse;
}

export namespace InquiryResponse {
    export type AsObject = {
        inquiryId: string,
        senderId: string,
        adminId: string,
        subject: string,
        description: string,
        email: string,
        isReplied: string,
        createdAt: string,
        updatedAt: string,
    }
}
