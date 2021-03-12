// package: proto
// file: proto/book_apiv1.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class CreateBookRequest extends jspb.Message { 
    getTitle(): string;
    setTitle(value: string): CreateBookRequest;

    getDescription(): string;
    setDescription(value: string): CreateBookRequest;

    getIsbn(): string;
    setIsbn(value: string): CreateBookRequest;

    getThumbnailUrl(): string;
    setThumbnailUrl(value: string): CreateBookRequest;

    getVersion(): string;
    setVersion(value: string): CreateBookRequest;

    getPublisher(): string;
    setPublisher(value: string): CreateBookRequest;

    getPublishedOn(): string;
    setPublishedOn(value: string): CreateBookRequest;

    clearAuthorsList(): void;
    getAuthorsList(): Array<CreateBookRequest.Author>;
    setAuthorsList(value: Array<CreateBookRequest.Author>): CreateBookRequest;
    addAuthors(value?: CreateBookRequest.Author, index?: number): CreateBookRequest.Author;

    clearCategoriesList(): void;
    getCategoriesList(): Array<CreateBookRequest.Category>;
    setCategoriesList(value: Array<CreateBookRequest.Category>): CreateBookRequest;
    addCategories(value?: CreateBookRequest.Category, index?: number): CreateBookRequest.Category;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateBookRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateBookRequest): CreateBookRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateBookRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateBookRequest;
    static deserializeBinaryFromReader(message: CreateBookRequest, reader: jspb.BinaryReader): CreateBookRequest;
}

export namespace CreateBookRequest {
    export type AsObject = {
        title: string,
        description: string,
        isbn: string,
        thumbnailUrl: string,
        version: string,
        publisher: string,
        publishedOn: string,
        authorsList: Array<CreateBookRequest.Author.AsObject>,
        categoriesList: Array<CreateBookRequest.Category.AsObject>,
    }


    export class Author extends jspb.Message { 
        getName(): string;
        setName(value: string): Author;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Author.AsObject;
        static toObject(includeInstance: boolean, msg: Author): Author.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Author, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Author;
        static deserializeBinaryFromReader(message: Author, reader: jspb.BinaryReader): Author;
    }

    export namespace Author {
        export type AsObject = {
            name: string,
        }
    }

    export class Category extends jspb.Message { 
        getName(): string;
        setName(value: string): Category;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Category.AsObject;
        static toObject(includeInstance: boolean, msg: Category): Category.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Category, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Category;
        static deserializeBinaryFromReader(message: Category, reader: jspb.BinaryReader): Category;
    }

    export namespace Category {
        export type AsObject = {
            name: string,
        }
    }

}

export class BookResponse extends jspb.Message { 
    getId(): number;
    setId(value: number): BookResponse;

    getPublisherId(): number;
    setPublisherId(value: number): BookResponse;

    getTitle(): string;
    setTitle(value: string): BookResponse;

    getDescription(): string;
    setDescription(value: string): BookResponse;

    getIsbn(): string;
    setIsbn(value: string): BookResponse;

    getThumbnailUrl(): string;
    setThumbnailUrl(value: string): BookResponse;

    getVersion(): string;
    setVersion(value: string): BookResponse;

    getPublishedOn(): string;
    setPublishedOn(value: string): BookResponse;

    getCreatedAt(): string;
    setCreatedAt(value: string): BookResponse;

    getUpdatedAt(): string;
    setUpdatedAt(value: string): BookResponse;

    clearAuthorsList(): void;
    getAuthorsList(): Array<BookResponse.Author>;
    setAuthorsList(value: Array<BookResponse.Author>): BookResponse;
    addAuthors(value?: BookResponse.Author, index?: number): BookResponse.Author;

    clearCategoriesList(): void;
    getCategoriesList(): Array<BookResponse.Category>;
    setCategoriesList(value: Array<BookResponse.Category>): BookResponse;
    addCategories(value?: BookResponse.Category, index?: number): BookResponse.Category;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BookResponse.AsObject;
    static toObject(includeInstance: boolean, msg: BookResponse): BookResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BookResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BookResponse;
    static deserializeBinaryFromReader(message: BookResponse, reader: jspb.BinaryReader): BookResponse;
}

export namespace BookResponse {
    export type AsObject = {
        id: number,
        publisherId: number,
        title: string,
        description: string,
        isbn: string,
        thumbnailUrl: string,
        version: string,
        publishedOn: string,
        createdAt: string,
        updatedAt: string,
        authorsList: Array<BookResponse.Author.AsObject>,
        categoriesList: Array<BookResponse.Category.AsObject>,
    }


    export class Author extends jspb.Message { 
        getId(): number;
        setId(value: number): Author;

        getName(): string;
        setName(value: string): Author;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Author.AsObject;
        static toObject(includeInstance: boolean, msg: Author): Author.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Author, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Author;
        static deserializeBinaryFromReader(message: Author, reader: jspb.BinaryReader): Author;
    }

    export namespace Author {
        export type AsObject = {
            id: number,
            name: string,
        }
    }

    export class Category extends jspb.Message { 
        getId(): number;
        setId(value: number): Category;

        getName(): string;
        setName(value: string): Category;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Category.AsObject;
        static toObject(includeInstance: boolean, msg: Category): Category.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Category, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Category;
        static deserializeBinaryFromReader(message: Category, reader: jspb.BinaryReader): Category;
    }

    export namespace Category {
        export type AsObject = {
            id: number,
            name: string,
        }
    }

}
