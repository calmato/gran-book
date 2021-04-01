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
    getAuthorsList(): Array<string>;
    setAuthorsList(value: Array<string>): CreateBookRequest;
    addAuthors(value: string, index?: number): string;

    clearCategoriesList(): void;
    getCategoriesList(): Array<string>;
    setCategoriesList(value: Array<string>): CreateBookRequest;
    addCategories(value: string, index?: number): string;


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
        authorsList: Array<string>,
        categoriesList: Array<string>,
    }
}

export class CreateAndUpdateBooksRequest extends jspb.Message { 
    clearBooksList(): void;
    getBooksList(): Array<CreateAndUpdateBooksRequest.Book>;
    setBooksList(value: Array<CreateAndUpdateBooksRequest.Book>): CreateAndUpdateBooksRequest;
    addBooks(value?: CreateAndUpdateBooksRequest.Book, index?: number): CreateAndUpdateBooksRequest.Book;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateAndUpdateBooksRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateAndUpdateBooksRequest): CreateAndUpdateBooksRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateAndUpdateBooksRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateAndUpdateBooksRequest;
    static deserializeBinaryFromReader(message: CreateAndUpdateBooksRequest, reader: jspb.BinaryReader): CreateAndUpdateBooksRequest;
}

export namespace CreateAndUpdateBooksRequest {
    export type AsObject = {
        booksList: Array<CreateAndUpdateBooksRequest.Book.AsObject>,
    }


    export class Book extends jspb.Message { 
        getTitle(): string;
        setTitle(value: string): Book;

        getDescription(): string;
        setDescription(value: string): Book;

        getIsbn(): string;
        setIsbn(value: string): Book;

        getThumbnailUrl(): string;
        setThumbnailUrl(value: string): Book;

        getVersion(): string;
        setVersion(value: string): Book;

        getPublisher(): string;
        setPublisher(value: string): Book;

        getPublishedOn(): string;
        setPublishedOn(value: string): Book;

        clearAuthorsList(): void;
        getAuthorsList(): Array<string>;
        setAuthorsList(value: Array<string>): Book;
        addAuthors(value: string, index?: number): string;

        clearCategoriesList(): void;
        getCategoriesList(): Array<string>;
        setCategoriesList(value: Array<string>): Book;
        addCategories(value: string, index?: number): string;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Book.AsObject;
        static toObject(includeInstance: boolean, msg: Book): Book.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Book, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Book;
        static deserializeBinaryFromReader(message: Book, reader: jspb.BinaryReader): Book;
    }

    export namespace Book {
        export type AsObject = {
            title: string,
            description: string,
            isbn: string,
            thumbnailUrl: string,
            version: string,
            publisher: string,
            publishedOn: string,
            authorsList: Array<string>,
            categoriesList: Array<string>,
        }
    }

}

export class ReadBookshelfRequest extends jspb.Message { 
    getBookId(): number;
    setBookId(value: number): ReadBookshelfRequest;

    getImpression(): string;
    setImpression(value: string): ReadBookshelfRequest;

    getReadOn(): string;
    setReadOn(value: string): ReadBookshelfRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ReadBookshelfRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ReadBookshelfRequest): ReadBookshelfRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ReadBookshelfRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ReadBookshelfRequest;
    static deserializeBinaryFromReader(message: ReadBookshelfRequest, reader: jspb.BinaryReader): ReadBookshelfRequest;
}

export namespace ReadBookshelfRequest {
    export type AsObject = {
        bookId: number,
        impression: string,
        readOn: string,
    }
}

export class ReadingBookshelfRequest extends jspb.Message { 
    getBookId(): number;
    setBookId(value: number): ReadingBookshelfRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ReadingBookshelfRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ReadingBookshelfRequest): ReadingBookshelfRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ReadingBookshelfRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ReadingBookshelfRequest;
    static deserializeBinaryFromReader(message: ReadingBookshelfRequest, reader: jspb.BinaryReader): ReadingBookshelfRequest;
}

export namespace ReadingBookshelfRequest {
    export type AsObject = {
        bookId: number,
    }
}

export class StackBookshelfRequest extends jspb.Message { 
    getBookId(): number;
    setBookId(value: number): StackBookshelfRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): StackBookshelfRequest.AsObject;
    static toObject(includeInstance: boolean, msg: StackBookshelfRequest): StackBookshelfRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: StackBookshelfRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): StackBookshelfRequest;
    static deserializeBinaryFromReader(message: StackBookshelfRequest, reader: jspb.BinaryReader): StackBookshelfRequest;
}

export namespace StackBookshelfRequest {
    export type AsObject = {
        bookId: number,
    }
}

export class WantBookshelfRequest extends jspb.Message { 
    getBookId(): number;
    setBookId(value: number): WantBookshelfRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): WantBookshelfRequest.AsObject;
    static toObject(includeInstance: boolean, msg: WantBookshelfRequest): WantBookshelfRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: WantBookshelfRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): WantBookshelfRequest;
    static deserializeBinaryFromReader(message: WantBookshelfRequest, reader: jspb.BinaryReader): WantBookshelfRequest;
}

export namespace WantBookshelfRequest {
    export type AsObject = {
        bookId: number,
    }
}

export class ReleaseBookshelfRequest extends jspb.Message { 
    getBookId(): number;
    setBookId(value: number): ReleaseBookshelfRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ReleaseBookshelfRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ReleaseBookshelfRequest): ReleaseBookshelfRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ReleaseBookshelfRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ReleaseBookshelfRequest;
    static deserializeBinaryFromReader(message: ReleaseBookshelfRequest, reader: jspb.BinaryReader): ReleaseBookshelfRequest;
}

export namespace ReleaseBookshelfRequest {
    export type AsObject = {
        bookId: number,
    }
}

export class BookResponse extends jspb.Message { 
    getId(): number;
    setId(value: number): BookResponse;

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

    getPublisher(): string;
    setPublisher(value: string): BookResponse;

    getPublishedOn(): string;
    setPublishedOn(value: string): BookResponse;

    clearAuthorsList(): void;
    getAuthorsList(): Array<string>;
    setAuthorsList(value: Array<string>): BookResponse;
    addAuthors(value: string, index?: number): string;

    clearCategoriesList(): void;
    getCategoriesList(): Array<string>;
    setCategoriesList(value: Array<string>): BookResponse;
    addCategories(value: string, index?: number): string;

    getCreatedAt(): string;
    setCreatedAt(value: string): BookResponse;

    getUpdatedAt(): string;
    setUpdatedAt(value: string): BookResponse;


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
        title: string,
        description: string,
        isbn: string,
        thumbnailUrl: string,
        version: string,
        publisher: string,
        publishedOn: string,
        authorsList: Array<string>,
        categoriesList: Array<string>,
        createdAt: string,
        updatedAt: string,
    }
}

export class BookListResponse extends jspb.Message { 
    clearBooksList(): void;
    getBooksList(): Array<BookListResponse.Book>;
    setBooksList(value: Array<BookListResponse.Book>): BookListResponse;
    addBooks(value?: BookListResponse.Book, index?: number): BookListResponse.Book;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BookListResponse.AsObject;
    static toObject(includeInstance: boolean, msg: BookListResponse): BookListResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BookListResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BookListResponse;
    static deserializeBinaryFromReader(message: BookListResponse, reader: jspb.BinaryReader): BookListResponse;
}

export namespace BookListResponse {
    export type AsObject = {
        booksList: Array<BookListResponse.Book.AsObject>,
    }


    export class Book extends jspb.Message { 
        getId(): number;
        setId(value: number): Book;

        getTitle(): string;
        setTitle(value: string): Book;

        getDescription(): string;
        setDescription(value: string): Book;

        getIsbn(): string;
        setIsbn(value: string): Book;

        getThumbnailUrl(): string;
        setThumbnailUrl(value: string): Book;

        getVersion(): string;
        setVersion(value: string): Book;

        getPublisher(): string;
        setPublisher(value: string): Book;

        getPublishedOn(): string;
        setPublishedOn(value: string): Book;

        clearAuthorsList(): void;
        getAuthorsList(): Array<string>;
        setAuthorsList(value: Array<string>): Book;
        addAuthors(value: string, index?: number): string;

        clearCategoriesList(): void;
        getCategoriesList(): Array<string>;
        setCategoriesList(value: Array<string>): Book;
        addCategories(value: string, index?: number): string;

        getCreatedAt(): string;
        setCreatedAt(value: string): Book;

        getUpdatedAt(): string;
        setUpdatedAt(value: string): Book;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Book.AsObject;
        static toObject(includeInstance: boolean, msg: Book): Book.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Book, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Book;
        static deserializeBinaryFromReader(message: Book, reader: jspb.BinaryReader): Book;
    }

    export namespace Book {
        export type AsObject = {
            id: number,
            title: string,
            description: string,
            isbn: string,
            thumbnailUrl: string,
            version: string,
            publisher: string,
            publishedOn: string,
            authorsList: Array<string>,
            categoriesList: Array<string>,
            createdAt: string,
            updatedAt: string,
        }
    }

}

export class BookshelfResponse extends jspb.Message { 
    getId(): number;
    setId(value: number): BookshelfResponse;

    getBookId(): number;
    setBookId(value: number): BookshelfResponse;

    getUserId(): string;
    setUserId(value: string): BookshelfResponse;

    getStatus(): number;
    setStatus(value: number): BookshelfResponse;

    getImpression(): string;
    setImpression(value: string): BookshelfResponse;

    getReadOn(): string;
    setReadOn(value: string): BookshelfResponse;

    getCreatedAt(): string;
    setCreatedAt(value: string): BookshelfResponse;

    getUpdatedAt(): string;
    setUpdatedAt(value: string): BookshelfResponse;


    hasBook(): boolean;
    clearBook(): void;
    getBook(): BookshelfResponse.Book | undefined;
    setBook(value?: BookshelfResponse.Book): BookshelfResponse;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BookshelfResponse.AsObject;
    static toObject(includeInstance: boolean, msg: BookshelfResponse): BookshelfResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BookshelfResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BookshelfResponse;
    static deserializeBinaryFromReader(message: BookshelfResponse, reader: jspb.BinaryReader): BookshelfResponse;
}

export namespace BookshelfResponse {
    export type AsObject = {
        id: number,
        bookId: number,
        userId: string,
        status: number,
        impression: string,
        readOn: string,
        createdAt: string,
        updatedAt: string,
        book?: BookshelfResponse.Book.AsObject,
    }


    export class Book extends jspb.Message { 
        getId(): number;
        setId(value: number): Book;

        getTitle(): string;
        setTitle(value: string): Book;

        getDescription(): string;
        setDescription(value: string): Book;

        getIsbn(): string;
        setIsbn(value: string): Book;

        getThumbnailUrl(): string;
        setThumbnailUrl(value: string): Book;

        getVersion(): string;
        setVersion(value: string): Book;

        getPublisher(): string;
        setPublisher(value: string): Book;

        getPublishedOn(): string;
        setPublishedOn(value: string): Book;

        clearAuthorsList(): void;
        getAuthorsList(): Array<string>;
        setAuthorsList(value: Array<string>): Book;
        addAuthors(value: string, index?: number): string;

        clearCategoriesList(): void;
        getCategoriesList(): Array<string>;
        setCategoriesList(value: Array<string>): Book;
        addCategories(value: string, index?: number): string;

        getCreatedAt(): string;
        setCreatedAt(value: string): Book;

        getUpdatedAt(): string;
        setUpdatedAt(value: string): Book;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Book.AsObject;
        static toObject(includeInstance: boolean, msg: Book): Book.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Book, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Book;
        static deserializeBinaryFromReader(message: Book, reader: jspb.BinaryReader): Book;
    }

    export namespace Book {
        export type AsObject = {
            id: number,
            title: string,
            description: string,
            isbn: string,
            thumbnailUrl: string,
            version: string,
            publisher: string,
            publishedOn: string,
            authorsList: Array<string>,
            categoriesList: Array<string>,
            createdAt: string,
            updatedAt: string,
        }
    }

}
