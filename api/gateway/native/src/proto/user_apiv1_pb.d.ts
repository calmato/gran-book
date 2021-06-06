// package: proto
// file: proto/user_apiv1.proto

/* tslint:disable */
/* eslint-disable */

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
    setUsername(value: string): CreateAuthRequest;
    getEmail(): string;
    setEmail(value: string): CreateAuthRequest;
    getPassword(): string;
    setPassword(value: string): CreateAuthRequest;
    getPasswordConfirmation(): string;
    setPasswordConfirmation(value: string): CreateAuthRequest;

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
    setEmail(value: string): UpdateAuthEmailRequest;

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
    setPassword(value: string): UpdateAuthPasswordRequest;
    getPasswordConfirmation(): string;
    setPasswordConfirmation(value: string): UpdateAuthPasswordRequest;

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
    setUsername(value: string): UpdateAuthProfileRequest;
    getGender(): number;
    setGender(value: number): UpdateAuthProfileRequest;
    getThumbnailUrl(): string;
    setThumbnailUrl(value: string): UpdateAuthProfileRequest;
    getSelfIntroduction(): string;
    setSelfIntroduction(value: string): UpdateAuthProfileRequest;

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
        thumbnailUrl: string,
        selfIntroduction: string,
    }
}

export class UpdateAuthAddressRequest extends jspb.Message { 
    getLastName(): string;
    setLastName(value: string): UpdateAuthAddressRequest;
    getFirstName(): string;
    setFirstName(value: string): UpdateAuthAddressRequest;
    getLastNameKana(): string;
    setLastNameKana(value: string): UpdateAuthAddressRequest;
    getFirstNameKana(): string;
    setFirstNameKana(value: string): UpdateAuthAddressRequest;
    getPhoneNumber(): string;
    setPhoneNumber(value: string): UpdateAuthAddressRequest;
    getPostalCode(): string;
    setPostalCode(value: string): UpdateAuthAddressRequest;
    getPrefecture(): string;
    setPrefecture(value: string): UpdateAuthAddressRequest;
    getCity(): string;
    setCity(value: string): UpdateAuthAddressRequest;
    getAddressLine1(): string;
    setAddressLine1(value: string): UpdateAuthAddressRequest;
    getAddressLine2(): string;
    setAddressLine2(value: string): UpdateAuthAddressRequest;

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

export class UploadAuthThumbnailRequest extends jspb.Message { 
    getThumbnail(): Uint8Array | string;
    getThumbnail_asU8(): Uint8Array;
    getThumbnail_asB64(): string;
    setThumbnail(value: Uint8Array | string): UploadAuthThumbnailRequest;
    getPosition(): number;
    setPosition(value: number): UploadAuthThumbnailRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UploadAuthThumbnailRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UploadAuthThumbnailRequest): UploadAuthThumbnailRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UploadAuthThumbnailRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UploadAuthThumbnailRequest;
    static deserializeBinaryFromReader(message: UploadAuthThumbnailRequest, reader: jspb.BinaryReader): UploadAuthThumbnailRequest;
}

export namespace UploadAuthThumbnailRequest {
    export type AsObject = {
        thumbnail: Uint8Array | string,
        position: number,
    }
}

export class RegisterAuthDeviceRequest extends jspb.Message { 
    getInstanceId(): string;
    setInstanceId(value: string): RegisterAuthDeviceRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RegisterAuthDeviceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: RegisterAuthDeviceRequest): RegisterAuthDeviceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RegisterAuthDeviceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RegisterAuthDeviceRequest;
    static deserializeBinaryFromReader(message: RegisterAuthDeviceRequest, reader: jspb.BinaryReader): RegisterAuthDeviceRequest;
}

export namespace RegisterAuthDeviceRequest {
    export type AsObject = {
        instanceId: string,
    }
}

export class ListAdminRequest extends jspb.Message { 
    getLimit(): number;
    setLimit(value: number): ListAdminRequest;
    getOffset(): number;
    setOffset(value: number): ListAdminRequest;

    hasOrder(): boolean;
    clearOrder(): void;
    getOrder(): ListAdminRequest.Order | undefined;
    setOrder(value?: ListAdminRequest.Order): ListAdminRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListAdminRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListAdminRequest): ListAdminRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListAdminRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListAdminRequest;
    static deserializeBinaryFromReader(message: ListAdminRequest, reader: jspb.BinaryReader): ListAdminRequest;
}

export namespace ListAdminRequest {
    export type AsObject = {
        limit: number,
        offset: number,
        order?: ListAdminRequest.Order.AsObject,
    }


    export class Order extends jspb.Message { 
        getBy(): string;
        setBy(value: string): Order;
        getDirection(): string;
        setDirection(value: string): Order;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Order.AsObject;
        static toObject(includeInstance: boolean, msg: Order): Order.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Order, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Order;
        static deserializeBinaryFromReader(message: Order, reader: jspb.BinaryReader): Order;
    }

    export namespace Order {
        export type AsObject = {
            by: string,
            direction: string,
        }
    }

}

export class SearchAdminRequest extends jspb.Message { 
    getLimit(): number;
    setLimit(value: number): SearchAdminRequest;
    getOffset(): number;
    setOffset(value: number): SearchAdminRequest;

    hasOrder(): boolean;
    clearOrder(): void;
    getOrder(): SearchAdminRequest.Order | undefined;
    setOrder(value?: SearchAdminRequest.Order): SearchAdminRequest;

    hasSearch(): boolean;
    clearSearch(): void;
    getSearch(): SearchAdminRequest.Search | undefined;
    setSearch(value?: SearchAdminRequest.Search): SearchAdminRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SearchAdminRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SearchAdminRequest): SearchAdminRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SearchAdminRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SearchAdminRequest;
    static deserializeBinaryFromReader(message: SearchAdminRequest, reader: jspb.BinaryReader): SearchAdminRequest;
}

export namespace SearchAdminRequest {
    export type AsObject = {
        limit: number,
        offset: number,
        order?: SearchAdminRequest.Order.AsObject,
        search?: SearchAdminRequest.Search.AsObject,
    }


    export class Order extends jspb.Message { 
        getBy(): string;
        setBy(value: string): Order;
        getDirection(): string;
        setDirection(value: string): Order;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Order.AsObject;
        static toObject(includeInstance: boolean, msg: Order): Order.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Order, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Order;
        static deserializeBinaryFromReader(message: Order, reader: jspb.BinaryReader): Order;
    }

    export namespace Order {
        export type AsObject = {
            by: string,
            direction: string,
        }
    }

    export class Search extends jspb.Message { 
        getField(): string;
        setField(value: string): Search;
        getValue(): string;
        setValue(value: string): Search;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Search.AsObject;
        static toObject(includeInstance: boolean, msg: Search): Search.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Search, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Search;
        static deserializeBinaryFromReader(message: Search, reader: jspb.BinaryReader): Search;
    }

    export namespace Search {
        export type AsObject = {
            field: string,
            value: string,
        }
    }

}

export class GetAdminRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): GetAdminRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetAdminRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetAdminRequest): GetAdminRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetAdminRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetAdminRequest;
    static deserializeBinaryFromReader(message: GetAdminRequest, reader: jspb.BinaryReader): GetAdminRequest;
}

export namespace GetAdminRequest {
    export type AsObject = {
        id: string,
    }
}

export class CreateAdminRequest extends jspb.Message { 
    getUsername(): string;
    setUsername(value: string): CreateAdminRequest;
    getEmail(): string;
    setEmail(value: string): CreateAdminRequest;
    getPassword(): string;
    setPassword(value: string): CreateAdminRequest;
    getPasswordConfirmation(): string;
    setPasswordConfirmation(value: string): CreateAdminRequest;
    getRole(): number;
    setRole(value: number): CreateAdminRequest;
    getLastName(): string;
    setLastName(value: string): CreateAdminRequest;
    getFirstName(): string;
    setFirstName(value: string): CreateAdminRequest;
    getLastNameKana(): string;
    setLastNameKana(value: string): CreateAdminRequest;
    getFirstNameKana(): string;
    setFirstNameKana(value: string): CreateAdminRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateAdminRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateAdminRequest): CreateAdminRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateAdminRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateAdminRequest;
    static deserializeBinaryFromReader(message: CreateAdminRequest, reader: jspb.BinaryReader): CreateAdminRequest;
}

export namespace CreateAdminRequest {
    export type AsObject = {
        username: string,
        email: string,
        password: string,
        passwordConfirmation: string,
        role: number,
        lastName: string,
        firstName: string,
        lastNameKana: string,
        firstNameKana: string,
    }
}

export class UpdateAdminContactRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): UpdateAdminContactRequest;
    getEmail(): string;
    setEmail(value: string): UpdateAdminContactRequest;
    getPhoneNumber(): string;
    setPhoneNumber(value: string): UpdateAdminContactRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateAdminContactRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateAdminContactRequest): UpdateAdminContactRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateAdminContactRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateAdminContactRequest;
    static deserializeBinaryFromReader(message: UpdateAdminContactRequest, reader: jspb.BinaryReader): UpdateAdminContactRequest;
}

export namespace UpdateAdminContactRequest {
    export type AsObject = {
        id: string,
        email: string,
        phoneNumber: string,
    }
}

export class UpdateAdminPasswordRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): UpdateAdminPasswordRequest;
    getPassword(): string;
    setPassword(value: string): UpdateAdminPasswordRequest;
    getPasswordConfirmation(): string;
    setPasswordConfirmation(value: string): UpdateAdminPasswordRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateAdminPasswordRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateAdminPasswordRequest): UpdateAdminPasswordRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateAdminPasswordRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateAdminPasswordRequest;
    static deserializeBinaryFromReader(message: UpdateAdminPasswordRequest, reader: jspb.BinaryReader): UpdateAdminPasswordRequest;
}

export namespace UpdateAdminPasswordRequest {
    export type AsObject = {
        id: string,
        password: string,
        passwordConfirmation: string,
    }
}

export class UpdateAdminProfileRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): UpdateAdminProfileRequest;
    getUsername(): string;
    setUsername(value: string): UpdateAdminProfileRequest;
    getRole(): number;
    setRole(value: number): UpdateAdminProfileRequest;
    getLastName(): string;
    setLastName(value: string): UpdateAdminProfileRequest;
    getFirstName(): string;
    setFirstName(value: string): UpdateAdminProfileRequest;
    getLastNameKana(): string;
    setLastNameKana(value: string): UpdateAdminProfileRequest;
    getFirstNameKana(): string;
    setFirstNameKana(value: string): UpdateAdminProfileRequest;
    getThumbnailUrl(): string;
    setThumbnailUrl(value: string): UpdateAdminProfileRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateAdminProfileRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateAdminProfileRequest): UpdateAdminProfileRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateAdminProfileRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateAdminProfileRequest;
    static deserializeBinaryFromReader(message: UpdateAdminProfileRequest, reader: jspb.BinaryReader): UpdateAdminProfileRequest;
}

export namespace UpdateAdminProfileRequest {
    export type AsObject = {
        id: string,
        username: string,
        role: number,
        lastName: string,
        firstName: string,
        lastNameKana: string,
        firstNameKana: string,
        thumbnailUrl: string,
    }
}

export class UploadAdminThumbnailRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): UploadAdminThumbnailRequest;
    getThumbnail(): Uint8Array | string;
    getThumbnail_asU8(): Uint8Array;
    getThumbnail_asB64(): string;
    setThumbnail(value: Uint8Array | string): UploadAdminThumbnailRequest;
    getPosition(): number;
    setPosition(value: number): UploadAdminThumbnailRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UploadAdminThumbnailRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UploadAdminThumbnailRequest): UploadAdminThumbnailRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UploadAdminThumbnailRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UploadAdminThumbnailRequest;
    static deserializeBinaryFromReader(message: UploadAdminThumbnailRequest, reader: jspb.BinaryReader): UploadAdminThumbnailRequest;
}

export namespace UploadAdminThumbnailRequest {
    export type AsObject = {
        userId: string,
        thumbnail: Uint8Array | string,
        position: number,
    }
}

export class DeleteAdminRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): DeleteAdminRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteAdminRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteAdminRequest): DeleteAdminRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteAdminRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteAdminRequest;
    static deserializeBinaryFromReader(message: DeleteAdminRequest, reader: jspb.BinaryReader): DeleteAdminRequest;
}

export namespace DeleteAdminRequest {
    export type AsObject = {
        userId: string,
    }
}

export class ListUserRequest extends jspb.Message { 
    getLimit(): number;
    setLimit(value: number): ListUserRequest;
    getOffset(): number;
    setOffset(value: number): ListUserRequest;

    hasOrder(): boolean;
    clearOrder(): void;
    getOrder(): ListUserRequest.Order | undefined;
    setOrder(value?: ListUserRequest.Order): ListUserRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListUserRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListUserRequest): ListUserRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListUserRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListUserRequest;
    static deserializeBinaryFromReader(message: ListUserRequest, reader: jspb.BinaryReader): ListUserRequest;
}

export namespace ListUserRequest {
    export type AsObject = {
        limit: number,
        offset: number,
        order?: ListUserRequest.Order.AsObject,
    }


    export class Order extends jspb.Message { 
        getBy(): string;
        setBy(value: string): Order;
        getDirection(): string;
        setDirection(value: string): Order;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Order.AsObject;
        static toObject(includeInstance: boolean, msg: Order): Order.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Order, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Order;
        static deserializeBinaryFromReader(message: Order, reader: jspb.BinaryReader): Order;
    }

    export namespace Order {
        export type AsObject = {
            by: string,
            direction: string,
        }
    }

}

export class ListUserByUserIdsRequest extends jspb.Message { 
    clearUserIdsList(): void;
    getUserIdsList(): Array<string>;
    setUserIdsList(value: Array<string>): ListUserByUserIdsRequest;
    addUserIds(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListUserByUserIdsRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListUserByUserIdsRequest): ListUserByUserIdsRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListUserByUserIdsRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListUserByUserIdsRequest;
    static deserializeBinaryFromReader(message: ListUserByUserIdsRequest, reader: jspb.BinaryReader): ListUserByUserIdsRequest;
}

export namespace ListUserByUserIdsRequest {
    export type AsObject = {
        userIdsList: Array<string>,
    }
}

export class ListFollowRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): ListFollowRequest;
    getLimit(): number;
    setLimit(value: number): ListFollowRequest;
    getOffset(): number;
    setOffset(value: number): ListFollowRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListFollowRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListFollowRequest): ListFollowRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListFollowRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListFollowRequest;
    static deserializeBinaryFromReader(message: ListFollowRequest, reader: jspb.BinaryReader): ListFollowRequest;
}

export namespace ListFollowRequest {
    export type AsObject = {
        id: string,
        limit: number,
        offset: number,
    }
}

export class ListFollowerRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): ListFollowerRequest;
    getLimit(): number;
    setLimit(value: number): ListFollowerRequest;
    getOffset(): number;
    setOffset(value: number): ListFollowerRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListFollowerRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListFollowerRequest): ListFollowerRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListFollowerRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListFollowerRequest;
    static deserializeBinaryFromReader(message: ListFollowerRequest, reader: jspb.BinaryReader): ListFollowerRequest;
}

export namespace ListFollowerRequest {
    export type AsObject = {
        id: string,
        limit: number,
        offset: number,
    }
}

export class SearchUserRequest extends jspb.Message { 
    getLimit(): number;
    setLimit(value: number): SearchUserRequest;
    getOffset(): number;
    setOffset(value: number): SearchUserRequest;

    hasOrder(): boolean;
    clearOrder(): void;
    getOrder(): SearchUserRequest.Order | undefined;
    setOrder(value?: SearchUserRequest.Order): SearchUserRequest;

    hasSearch(): boolean;
    clearSearch(): void;
    getSearch(): SearchUserRequest.Search | undefined;
    setSearch(value?: SearchUserRequest.Search): SearchUserRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SearchUserRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SearchUserRequest): SearchUserRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SearchUserRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SearchUserRequest;
    static deserializeBinaryFromReader(message: SearchUserRequest, reader: jspb.BinaryReader): SearchUserRequest;
}

export namespace SearchUserRequest {
    export type AsObject = {
        limit: number,
        offset: number,
        order?: SearchUserRequest.Order.AsObject,
        search?: SearchUserRequest.Search.AsObject,
    }


    export class Order extends jspb.Message { 
        getBy(): string;
        setBy(value: string): Order;
        getDirection(): string;
        setDirection(value: string): Order;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Order.AsObject;
        static toObject(includeInstance: boolean, msg: Order): Order.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Order, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Order;
        static deserializeBinaryFromReader(message: Order, reader: jspb.BinaryReader): Order;
    }

    export namespace Order {
        export type AsObject = {
            by: string,
            direction: string,
        }
    }

    export class Search extends jspb.Message { 
        getField(): string;
        setField(value: string): Search;
        getValue(): string;
        setValue(value: string): Search;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Search.AsObject;
        static toObject(includeInstance: boolean, msg: Search): Search.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Search, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Search;
        static deserializeBinaryFromReader(message: Search, reader: jspb.BinaryReader): Search;
    }

    export namespace Search {
        export type AsObject = {
            field: string,
            value: string,
        }
    }

}

export class GetUserRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): GetUserRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetUserRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetUserRequest): GetUserRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetUserRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetUserRequest;
    static deserializeBinaryFromReader(message: GetUserRequest, reader: jspb.BinaryReader): GetUserRequest;
}

export namespace GetUserRequest {
    export type AsObject = {
        id: string,
    }
}

export class GetUserProfileRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): GetUserProfileRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetUserProfileRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetUserProfileRequest): GetUserProfileRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetUserProfileRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetUserProfileRequest;
    static deserializeBinaryFromReader(message: GetUserProfileRequest, reader: jspb.BinaryReader): GetUserProfileRequest;
}

export namespace GetUserProfileRequest {
    export type AsObject = {
        id: string,
    }
}

export class RegisterFollowRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): RegisterFollowRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RegisterFollowRequest.AsObject;
    static toObject(includeInstance: boolean, msg: RegisterFollowRequest): RegisterFollowRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RegisterFollowRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RegisterFollowRequest;
    static deserializeBinaryFromReader(message: RegisterFollowRequest, reader: jspb.BinaryReader): RegisterFollowRequest;
}

export namespace RegisterFollowRequest {
    export type AsObject = {
        id: string,
    }
}

export class UnregisterFollowRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): UnregisterFollowRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UnregisterFollowRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UnregisterFollowRequest): UnregisterFollowRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UnregisterFollowRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UnregisterFollowRequest;
    static deserializeBinaryFromReader(message: UnregisterFollowRequest, reader: jspb.BinaryReader): UnregisterFollowRequest;
}

export namespace UnregisterFollowRequest {
    export type AsObject = {
        id: string,
    }
}

export class ListChatRoomRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): ListChatRoomRequest;
    getLimit(): number;
    setLimit(value: number): ListChatRoomRequest;
    getOffset(): number;
    setOffset(value: number): ListChatRoomRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListChatRoomRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListChatRoomRequest): ListChatRoomRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListChatRoomRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListChatRoomRequest;
    static deserializeBinaryFromReader(message: ListChatRoomRequest, reader: jspb.BinaryReader): ListChatRoomRequest;
}

export namespace ListChatRoomRequest {
    export type AsObject = {
        userId: string,
        limit: number,
        offset: number,
    }
}

export class CreateChatRoomRequest extends jspb.Message { 
    clearUserIdsList(): void;
    getUserIdsList(): Array<string>;
    setUserIdsList(value: Array<string>): CreateChatRoomRequest;
    addUserIds(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateChatRoomRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateChatRoomRequest): CreateChatRoomRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateChatRoomRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateChatRoomRequest;
    static deserializeBinaryFromReader(message: CreateChatRoomRequest, reader: jspb.BinaryReader): CreateChatRoomRequest;
}

export namespace CreateChatRoomRequest {
    export type AsObject = {
        userIdsList: Array<string>,
    }
}

export class CreateChatMessageRequest extends jspb.Message { 
    getText(): string;
    setText(value: string): CreateChatMessageRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateChatMessageRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateChatMessageRequest): CreateChatMessageRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateChatMessageRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateChatMessageRequest;
    static deserializeBinaryFromReader(message: CreateChatMessageRequest, reader: jspb.BinaryReader): CreateChatMessageRequest;
}

export namespace CreateChatMessageRequest {
    export type AsObject = {
        text: string,
    }
}

export class UploadChatImageRequest extends jspb.Message { 
    getImage(): Uint8Array | string;
    getImage_asU8(): Uint8Array;
    getImage_asB64(): string;
    setImage(value: Uint8Array | string): UploadChatImageRequest;
    getPosition(): number;
    setPosition(value: number): UploadChatImageRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UploadChatImageRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UploadChatImageRequest): UploadChatImageRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UploadChatImageRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UploadChatImageRequest;
    static deserializeBinaryFromReader(message: UploadChatImageRequest, reader: jspb.BinaryReader): UploadChatImageRequest;
}

export namespace UploadChatImageRequest {
    export type AsObject = {
        image: Uint8Array | string,
        position: number,
    }
}

export class AuthResponse extends jspb.Message { 
    getId(): string;
    setId(value: string): AuthResponse;
    getUsername(): string;
    setUsername(value: string): AuthResponse;
    getGender(): number;
    setGender(value: number): AuthResponse;
    getEmail(): string;
    setEmail(value: string): AuthResponse;
    getPhoneNumber(): string;
    setPhoneNumber(value: string): AuthResponse;
    getRole(): number;
    setRole(value: number): AuthResponse;
    getThumbnailUrl(): string;
    setThumbnailUrl(value: string): AuthResponse;
    getSelfIntroduction(): string;
    setSelfIntroduction(value: string): AuthResponse;
    getLastName(): string;
    setLastName(value: string): AuthResponse;
    getFirstName(): string;
    setFirstName(value: string): AuthResponse;
    getLastNameKana(): string;
    setLastNameKana(value: string): AuthResponse;
    getFirstNameKana(): string;
    setFirstNameKana(value: string): AuthResponse;
    getPostalCode(): string;
    setPostalCode(value: string): AuthResponse;
    getPrefecture(): string;
    setPrefecture(value: string): AuthResponse;
    getCity(): string;
    setCity(value: string): AuthResponse;
    getAddressLine1(): string;
    setAddressLine1(value: string): AuthResponse;
    getAddressLine2(): string;
    setAddressLine2(value: string): AuthResponse;
    getCreatedAt(): string;
    setCreatedAt(value: string): AuthResponse;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): AuthResponse;

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

export class AuthThumbnailResponse extends jspb.Message { 
    getThumbnailUrl(): string;
    setThumbnailUrl(value: string): AuthThumbnailResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AuthThumbnailResponse.AsObject;
    static toObject(includeInstance: boolean, msg: AuthThumbnailResponse): AuthThumbnailResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AuthThumbnailResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AuthThumbnailResponse;
    static deserializeBinaryFromReader(message: AuthThumbnailResponse, reader: jspb.BinaryReader): AuthThumbnailResponse;
}

export namespace AuthThumbnailResponse {
    export type AsObject = {
        thumbnailUrl: string,
    }
}

export class AdminResponse extends jspb.Message { 
    getId(): string;
    setId(value: string): AdminResponse;
    getUsername(): string;
    setUsername(value: string): AdminResponse;
    getEmail(): string;
    setEmail(value: string): AdminResponse;
    getPhoneNumber(): string;
    setPhoneNumber(value: string): AdminResponse;
    getRole(): number;
    setRole(value: number): AdminResponse;
    getThumbnailUrl(): string;
    setThumbnailUrl(value: string): AdminResponse;
    getSelfIntroduction(): string;
    setSelfIntroduction(value: string): AdminResponse;
    getLastName(): string;
    setLastName(value: string): AdminResponse;
    getFirstName(): string;
    setFirstName(value: string): AdminResponse;
    getLastNameKana(): string;
    setLastNameKana(value: string): AdminResponse;
    getFirstNameKana(): string;
    setFirstNameKana(value: string): AdminResponse;
    getCreatedAt(): string;
    setCreatedAt(value: string): AdminResponse;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): AdminResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminResponse.AsObject;
    static toObject(includeInstance: boolean, msg: AdminResponse): AdminResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminResponse;
    static deserializeBinaryFromReader(message: AdminResponse, reader: jspb.BinaryReader): AdminResponse;
}

export namespace AdminResponse {
    export type AsObject = {
        id: string,
        username: string,
        email: string,
        phoneNumber: string,
        role: number,
        thumbnailUrl: string,
        selfIntroduction: string,
        lastName: string,
        firstName: string,
        lastNameKana: string,
        firstNameKana: string,
        createdAt: string,
        updatedAt: string,
    }
}

export class AdminListResponse extends jspb.Message { 
    clearUsersList(): void;
    getUsersList(): Array<AdminListResponse.User>;
    setUsersList(value: Array<AdminListResponse.User>): AdminListResponse;
    addUsers(value?: AdminListResponse.User, index?: number): AdminListResponse.User;
    getLimit(): number;
    setLimit(value: number): AdminListResponse;
    getOffset(): number;
    setOffset(value: number): AdminListResponse;
    getTotal(): number;
    setTotal(value: number): AdminListResponse;

    hasOrder(): boolean;
    clearOrder(): void;
    getOrder(): AdminListResponse.Order | undefined;
    setOrder(value?: AdminListResponse.Order): AdminListResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminListResponse.AsObject;
    static toObject(includeInstance: boolean, msg: AdminListResponse): AdminListResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminListResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminListResponse;
    static deserializeBinaryFromReader(message: AdminListResponse, reader: jspb.BinaryReader): AdminListResponse;
}

export namespace AdminListResponse {
    export type AsObject = {
        usersList: Array<AdminListResponse.User.AsObject>,
        limit: number,
        offset: number,
        total: number,
        order?: AdminListResponse.Order.AsObject,
    }


    export class User extends jspb.Message { 
        getId(): string;
        setId(value: string): User;
        getUsername(): string;
        setUsername(value: string): User;
        getEmail(): string;
        setEmail(value: string): User;
        getPhoneNumber(): string;
        setPhoneNumber(value: string): User;
        getRole(): number;
        setRole(value: number): User;
        getThumbnailUrl(): string;
        setThumbnailUrl(value: string): User;
        getSelfIntroduction(): string;
        setSelfIntroduction(value: string): User;
        getLastName(): string;
        setLastName(value: string): User;
        getFirstName(): string;
        setFirstName(value: string): User;
        getLastNameKana(): string;
        setLastNameKana(value: string): User;
        getFirstNameKana(): string;
        setFirstNameKana(value: string): User;
        getCreatedAt(): string;
        setCreatedAt(value: string): User;
        getUpdatedAt(): string;
        setUpdatedAt(value: string): User;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): User.AsObject;
        static toObject(includeInstance: boolean, msg: User): User.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): User;
        static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
    }

    export namespace User {
        export type AsObject = {
            id: string,
            username: string,
            email: string,
            phoneNumber: string,
            role: number,
            thumbnailUrl: string,
            selfIntroduction: string,
            lastName: string,
            firstName: string,
            lastNameKana: string,
            firstNameKana: string,
            createdAt: string,
            updatedAt: string,
        }
    }

    export class Order extends jspb.Message { 
        getBy(): string;
        setBy(value: string): Order;
        getDirection(): string;
        setDirection(value: string): Order;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Order.AsObject;
        static toObject(includeInstance: boolean, msg: Order): Order.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Order, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Order;
        static deserializeBinaryFromReader(message: Order, reader: jspb.BinaryReader): Order;
    }

    export namespace Order {
        export type AsObject = {
            by: string,
            direction: string,
        }
    }

}

export class AdminThumbnailResponse extends jspb.Message { 
    getThumbnailUrl(): string;
    setThumbnailUrl(value: string): AdminThumbnailResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminThumbnailResponse.AsObject;
    static toObject(includeInstance: boolean, msg: AdminThumbnailResponse): AdminThumbnailResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminThumbnailResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminThumbnailResponse;
    static deserializeBinaryFromReader(message: AdminThumbnailResponse, reader: jspb.BinaryReader): AdminThumbnailResponse;
}

export namespace AdminThumbnailResponse {
    export type AsObject = {
        thumbnailUrl: string,
    }
}

export class UserResponse extends jspb.Message { 
    getId(): string;
    setId(value: string): UserResponse;
    getUsername(): string;
    setUsername(value: string): UserResponse;
    getEmail(): string;
    setEmail(value: string): UserResponse;
    getPhoneNumber(): string;
    setPhoneNumber(value: string): UserResponse;
    getRole(): number;
    setRole(value: number): UserResponse;
    getThumbnailUrl(): string;
    setThumbnailUrl(value: string): UserResponse;
    getSelfIntroduction(): string;
    setSelfIntroduction(value: string): UserResponse;
    getLastName(): string;
    setLastName(value: string): UserResponse;
    getFirstName(): string;
    setFirstName(value: string): UserResponse;
    getLastNameKana(): string;
    setLastNameKana(value: string): UserResponse;
    getFirstNameKana(): string;
    setFirstNameKana(value: string): UserResponse;
    getCreatedAt(): string;
    setCreatedAt(value: string): UserResponse;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): UserResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UserResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UserResponse): UserResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UserResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UserResponse;
    static deserializeBinaryFromReader(message: UserResponse, reader: jspb.BinaryReader): UserResponse;
}

export namespace UserResponse {
    export type AsObject = {
        id: string,
        username: string,
        email: string,
        phoneNumber: string,
        role: number,
        thumbnailUrl: string,
        selfIntroduction: string,
        lastName: string,
        firstName: string,
        lastNameKana: string,
        firstNameKana: string,
        createdAt: string,
        updatedAt: string,
    }
}

export class UserListResponse extends jspb.Message { 
    clearUsersList(): void;
    getUsersList(): Array<UserListResponse.User>;
    setUsersList(value: Array<UserListResponse.User>): UserListResponse;
    addUsers(value?: UserListResponse.User, index?: number): UserListResponse.User;
    getLimit(): number;
    setLimit(value: number): UserListResponse;
    getOffset(): number;
    setOffset(value: number): UserListResponse;
    getTotal(): number;
    setTotal(value: number): UserListResponse;

    hasOrder(): boolean;
    clearOrder(): void;
    getOrder(): UserListResponse.Order | undefined;
    setOrder(value?: UserListResponse.Order): UserListResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UserListResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UserListResponse): UserListResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UserListResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UserListResponse;
    static deserializeBinaryFromReader(message: UserListResponse, reader: jspb.BinaryReader): UserListResponse;
}

export namespace UserListResponse {
    export type AsObject = {
        usersList: Array<UserListResponse.User.AsObject>,
        limit: number,
        offset: number,
        total: number,
        order?: UserListResponse.Order.AsObject,
    }


    export class User extends jspb.Message { 
        getId(): string;
        setId(value: string): User;
        getUsername(): string;
        setUsername(value: string): User;
        getEmail(): string;
        setEmail(value: string): User;
        getPhoneNumber(): string;
        setPhoneNumber(value: string): User;
        getRole(): number;
        setRole(value: number): User;
        getThumbnailUrl(): string;
        setThumbnailUrl(value: string): User;
        getSelfIntroduction(): string;
        setSelfIntroduction(value: string): User;
        getLastName(): string;
        setLastName(value: string): User;
        getFirstName(): string;
        setFirstName(value: string): User;
        getLastNameKana(): string;
        setLastNameKana(value: string): User;
        getFirstNameKana(): string;
        setFirstNameKana(value: string): User;
        getCreatedAt(): string;
        setCreatedAt(value: string): User;
        getUpdatedAt(): string;
        setUpdatedAt(value: string): User;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): User.AsObject;
        static toObject(includeInstance: boolean, msg: User): User.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): User;
        static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
    }

    export namespace User {
        export type AsObject = {
            id: string,
            username: string,
            email: string,
            phoneNumber: string,
            role: number,
            thumbnailUrl: string,
            selfIntroduction: string,
            lastName: string,
            firstName: string,
            lastNameKana: string,
            firstNameKana: string,
            createdAt: string,
            updatedAt: string,
        }
    }

    export class Order extends jspb.Message { 
        getBy(): string;
        setBy(value: string): Order;
        getDirection(): string;
        setDirection(value: string): Order;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Order.AsObject;
        static toObject(includeInstance: boolean, msg: Order): Order.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Order, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Order;
        static deserializeBinaryFromReader(message: Order, reader: jspb.BinaryReader): Order;
    }

    export namespace Order {
        export type AsObject = {
            by: string,
            direction: string,
        }
    }

}

export class UserProfileResponse extends jspb.Message { 
    getId(): string;
    setId(value: string): UserProfileResponse;
    getUsername(): string;
    setUsername(value: string): UserProfileResponse;
    getThumbnailUrl(): string;
    setThumbnailUrl(value: string): UserProfileResponse;
    getSelfIntroduction(): string;
    setSelfIntroduction(value: string): UserProfileResponse;
    getIsFollow(): boolean;
    setIsFollow(value: boolean): UserProfileResponse;
    getIsFollower(): boolean;
    setIsFollower(value: boolean): UserProfileResponse;
    getFollowCount(): number;
    setFollowCount(value: number): UserProfileResponse;
    getFollowerCount(): number;
    setFollowerCount(value: number): UserProfileResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UserProfileResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UserProfileResponse): UserProfileResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UserProfileResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UserProfileResponse;
    static deserializeBinaryFromReader(message: UserProfileResponse, reader: jspb.BinaryReader): UserProfileResponse;
}

export namespace UserProfileResponse {
    export type AsObject = {
        id: string,
        username: string,
        thumbnailUrl: string,
        selfIntroduction: string,
        isFollow: boolean,
        isFollower: boolean,
        followCount: number,
        followerCount: number,
    }
}

export class FollowListResponse extends jspb.Message { 
    clearUsersList(): void;
    getUsersList(): Array<FollowListResponse.User>;
    setUsersList(value: Array<FollowListResponse.User>): FollowListResponse;
    addUsers(value?: FollowListResponse.User, index?: number): FollowListResponse.User;
    getLimit(): number;
    setLimit(value: number): FollowListResponse;
    getOffset(): number;
    setOffset(value: number): FollowListResponse;
    getTotal(): number;
    setTotal(value: number): FollowListResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FollowListResponse.AsObject;
    static toObject(includeInstance: boolean, msg: FollowListResponse): FollowListResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FollowListResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FollowListResponse;
    static deserializeBinaryFromReader(message: FollowListResponse, reader: jspb.BinaryReader): FollowListResponse;
}

export namespace FollowListResponse {
    export type AsObject = {
        usersList: Array<FollowListResponse.User.AsObject>,
        limit: number,
        offset: number,
        total: number,
    }


    export class User extends jspb.Message { 
        getId(): string;
        setId(value: string): User;
        getUsername(): string;
        setUsername(value: string): User;
        getThumbnailUrl(): string;
        setThumbnailUrl(value: string): User;
        getSelfIntroduction(): string;
        setSelfIntroduction(value: string): User;
        getIsFollow(): boolean;
        setIsFollow(value: boolean): User;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): User.AsObject;
        static toObject(includeInstance: boolean, msg: User): User.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): User;
        static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
    }

    export namespace User {
        export type AsObject = {
            id: string,
            username: string,
            thumbnailUrl: string,
            selfIntroduction: string,
            isFollow: boolean,
        }
    }

}

export class FollowerListResponse extends jspb.Message { 
    clearUsersList(): void;
    getUsersList(): Array<FollowerListResponse.User>;
    setUsersList(value: Array<FollowerListResponse.User>): FollowerListResponse;
    addUsers(value?: FollowerListResponse.User, index?: number): FollowerListResponse.User;
    getLimit(): number;
    setLimit(value: number): FollowerListResponse;
    getOffset(): number;
    setOffset(value: number): FollowerListResponse;
    getTotal(): number;
    setTotal(value: number): FollowerListResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FollowerListResponse.AsObject;
    static toObject(includeInstance: boolean, msg: FollowerListResponse): FollowerListResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FollowerListResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FollowerListResponse;
    static deserializeBinaryFromReader(message: FollowerListResponse, reader: jspb.BinaryReader): FollowerListResponse;
}

export namespace FollowerListResponse {
    export type AsObject = {
        usersList: Array<FollowerListResponse.User.AsObject>,
        limit: number,
        offset: number,
        total: number,
    }


    export class User extends jspb.Message { 
        getId(): string;
        setId(value: string): User;
        getUsername(): string;
        setUsername(value: string): User;
        getThumbnailUrl(): string;
        setThumbnailUrl(value: string): User;
        getSelfIntroduction(): string;
        setSelfIntroduction(value: string): User;
        getIsFollow(): boolean;
        setIsFollow(value: boolean): User;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): User.AsObject;
        static toObject(includeInstance: boolean, msg: User): User.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): User;
        static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
    }

    export namespace User {
        export type AsObject = {
            id: string,
            username: string,
            thumbnailUrl: string,
            selfIntroduction: string,
            isFollow: boolean,
        }
    }

}

export class ChatRoomResponse extends jspb.Message { 
    getId(): string;
    setId(value: string): ChatRoomResponse;
    clearUserIdsList(): void;
    getUserIdsList(): Array<string>;
    setUserIdsList(value: Array<string>): ChatRoomResponse;
    addUserIds(value: string, index?: number): string;
    getCreatedAt(): string;
    setCreatedAt(value: string): ChatRoomResponse;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): ChatRoomResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ChatRoomResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ChatRoomResponse): ChatRoomResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ChatRoomResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ChatRoomResponse;
    static deserializeBinaryFromReader(message: ChatRoomResponse, reader: jspb.BinaryReader): ChatRoomResponse;
}

export namespace ChatRoomResponse {
    export type AsObject = {
        id: string,
        userIdsList: Array<string>,
        createdAt: string,
        updatedAt: string,
    }
}

export class ChatRoomListResponse extends jspb.Message { 
    clearRoomsList(): void;
    getRoomsList(): Array<ChatRoomListResponse.Room>;
    setRoomsList(value: Array<ChatRoomListResponse.Room>): ChatRoomListResponse;
    addRooms(value?: ChatRoomListResponse.Room, index?: number): ChatRoomListResponse.Room;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ChatRoomListResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ChatRoomListResponse): ChatRoomListResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ChatRoomListResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ChatRoomListResponse;
    static deserializeBinaryFromReader(message: ChatRoomListResponse, reader: jspb.BinaryReader): ChatRoomListResponse;
}

export namespace ChatRoomListResponse {
    export type AsObject = {
        roomsList: Array<ChatRoomListResponse.Room.AsObject>,
    }


    export class Message extends jspb.Message { 
        getUserId(): string;
        setUserId(value: string): Message;
        getText(): string;
        setText(value: string): Message;
        getImage(): string;
        setImage(value: string): Message;
        getCreatedAt(): string;
        setCreatedAt(value: string): Message;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Message.AsObject;
        static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Message;
        static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
    }

    export namespace Message {
        export type AsObject = {
            userId: string,
            text: string,
            image: string,
            createdAt: string,
        }
    }

    export class Room extends jspb.Message { 
        getId(): string;
        setId(value: string): Room;
        clearUserIdsList(): void;
        getUserIdsList(): Array<string>;
        setUserIdsList(value: Array<string>): Room;
        addUserIds(value: string, index?: number): string;

        hasLatestmessage(): boolean;
        clearLatestmessage(): void;
        getLatestmessage(): ChatRoomListResponse.Message | undefined;
        setLatestmessage(value?: ChatRoomListResponse.Message): Room;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Room.AsObject;
        static toObject(includeInstance: boolean, msg: Room): Room.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Room, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Room;
        static deserializeBinaryFromReader(message: Room, reader: jspb.BinaryReader): Room;
    }

    export namespace Room {
        export type AsObject = {
            id: string,
            userIdsList: Array<string>,
            latestmessage?: ChatRoomListResponse.Message.AsObject,
        }
    }

}

export class ChatMessageResponse extends jspb.Message { 
    getId(): string;
    setId(value: string): ChatMessageResponse;
    getUserId(): string;
    setUserId(value: string): ChatMessageResponse;
    getText(): string;
    setText(value: string): ChatMessageResponse;
    getImage(): string;
    setImage(value: string): ChatMessageResponse;
    getCreatedAt(): string;
    setCreatedAt(value: string): ChatMessageResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ChatMessageResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ChatMessageResponse): ChatMessageResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ChatMessageResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ChatMessageResponse;
    static deserializeBinaryFromReader(message: ChatMessageResponse, reader: jspb.BinaryReader): ChatMessageResponse;
}

export namespace ChatMessageResponse {
    export type AsObject = {
        id: string,
        userId: string,
        text: string,
        image: string,
        createdAt: string,
    }
}
