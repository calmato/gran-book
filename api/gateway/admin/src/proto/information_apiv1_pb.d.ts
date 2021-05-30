// package: proto
// file: proto/information_apiv1.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class EmptyNotification extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EmptyNotification.AsObject;
    static toObject(includeInstance: boolean, msg: EmptyNotification): EmptyNotification.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EmptyNotification, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EmptyNotification;
    static deserializeBinaryFromReader(message: EmptyNotification, reader: jspb.BinaryReader): EmptyNotification;
}

export namespace EmptyNotification {
    export type AsObject = {
    }
}

export class GetNotificationRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): GetNotificationRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetNotificationRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetNotificationRequest): GetNotificationRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetNotificationRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetNotificationRequest;
    static deserializeBinaryFromReader(message: GetNotificationRequest, reader: jspb.BinaryReader): GetNotificationRequest;
}

export namespace GetNotificationRequest {
    export type AsObject = {
        id: string,
    }
}

export class CreateNotificationRequest extends jspb.Message { 
    getTitle(): string;
    setTitle(value: string): CreateNotificationRequest;
    getDescription(): string;
    setDescription(value: string): CreateNotificationRequest;
    getImportance(): string;
    setImportance(value: string): CreateNotificationRequest;
    getCategory(): string;
    setCategory(value: string): CreateNotificationRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateNotificationRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateNotificationRequest): CreateNotificationRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateNotificationRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateNotificationRequest;
    static deserializeBinaryFromReader(message: CreateNotificationRequest, reader: jspb.BinaryReader): CreateNotificationRequest;
}

export namespace CreateNotificationRequest {
    export type AsObject = {
        title: string,
        description: string,
        importance: string,
        category: string,
    }
}

export class UpdateNotificationRequest extends jspb.Message { 
    getTitle(): string;
    setTitle(value: string): UpdateNotificationRequest;
    getDescription(): string;
    setDescription(value: string): UpdateNotificationRequest;
    getImportance(): string;
    setImportance(value: string): UpdateNotificationRequest;
    getCategory(): string;
    setCategory(value: string): UpdateNotificationRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateNotificationRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateNotificationRequest): UpdateNotificationRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateNotificationRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateNotificationRequest;
    static deserializeBinaryFromReader(message: UpdateNotificationRequest, reader: jspb.BinaryReader): UpdateNotificationRequest;
}

export namespace UpdateNotificationRequest {
    export type AsObject = {
        title: string,
        description: string,
        importance: string,
        category: string,
    }
}

export class DeleteNotificationRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): DeleteNotificationRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteNotificationRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteNotificationRequest): DeleteNotificationRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteNotificationRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteNotificationRequest;
    static deserializeBinaryFromReader(message: DeleteNotificationRequest, reader: jspb.BinaryReader): DeleteNotificationRequest;
}

export namespace DeleteNotificationRequest {
    export type AsObject = {
        id: string,
    }
}

export class SearchNotificationRequest extends jspb.Message { 
    getLimit(): number;
    setLimit(value: number): SearchNotificationRequest;
    getOffset(): number;
    setOffset(value: number): SearchNotificationRequest;

    hasOrder(): boolean;
    clearOrder(): void;
    getOrder(): SearchNotificationRequest.Order | undefined;
    setOrder(value?: SearchNotificationRequest.Order): SearchNotificationRequest;

    hasSearch(): boolean;
    clearSearch(): void;
    getSearch(): SearchNotificationRequest.Search | undefined;
    setSearch(value?: SearchNotificationRequest.Search): SearchNotificationRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SearchNotificationRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SearchNotificationRequest): SearchNotificationRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SearchNotificationRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SearchNotificationRequest;
    static deserializeBinaryFromReader(message: SearchNotificationRequest, reader: jspb.BinaryReader): SearchNotificationRequest;
}

export namespace SearchNotificationRequest {
    export type AsObject = {
        limit: number,
        offset: number,
        order?: SearchNotificationRequest.Order.AsObject,
        search?: SearchNotificationRequest.Search.AsObject,
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

export class NotificationResponse extends jspb.Message { 
    getId(): string;
    setId(value: string): NotificationResponse;
    getAuthorId(): string;
    setAuthorId(value: string): NotificationResponse;
    getEditorId(): string;
    setEditorId(value: string): NotificationResponse;
    getTitle(): string;
    setTitle(value: string): NotificationResponse;
    getDescription(): string;
    setDescription(value: string): NotificationResponse;
    getImportance(): string;
    setImportance(value: string): NotificationResponse;
    getCategory(): string;
    setCategory(value: string): NotificationResponse;
    getCreatedAt(): string;
    setCreatedAt(value: string): NotificationResponse;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): NotificationResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NotificationResponse.AsObject;
    static toObject(includeInstance: boolean, msg: NotificationResponse): NotificationResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NotificationResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NotificationResponse;
    static deserializeBinaryFromReader(message: NotificationResponse, reader: jspb.BinaryReader): NotificationResponse;
}

export namespace NotificationResponse {
    export type AsObject = {
        id: string,
        authorId: string,
        editorId: string,
        title: string,
        description: string,
        importance: string,
        category: string,
        createdAt: string,
        updatedAt: string,
    }
}
