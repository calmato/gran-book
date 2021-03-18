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

export class CreateMultipleBooksRequest extends jspb.Message { 
    clearItemsList(): void;
    getItemsList(): Array<CreateMultipleBooksRequest.Item>;
    setItemsList(value: Array<CreateMultipleBooksRequest.Item>): CreateMultipleBooksRequest;
    addItems(value?: CreateMultipleBooksRequest.Item, index?: number): CreateMultipleBooksRequest.Item;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateMultipleBooksRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateMultipleBooksRequest): CreateMultipleBooksRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateMultipleBooksRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateMultipleBooksRequest;
    static deserializeBinaryFromReader(message: CreateMultipleBooksRequest, reader: jspb.BinaryReader): CreateMultipleBooksRequest;
}

export namespace CreateMultipleBooksRequest {
    export type AsObject = {
        itemsList: Array<CreateMultipleBooksRequest.Item.AsObject>,
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

    export class Item extends jspb.Message { 
        getTitle(): string;
        setTitle(value: string): Item;

        getDescription(): string;
        setDescription(value: string): Item;

        getIsbn(): string;
        setIsbn(value: string): Item;

        getThumbnailUrl(): string;
        setThumbnailUrl(value: string): Item;

        getVersion(): string;
        setVersion(value: string): Item;

        getPublisher(): string;
        setPublisher(value: string): Item;

        getPublishedOn(): string;
        setPublishedOn(value: string): Item;

        clearAuthorsList(): void;
        getAuthorsList(): Array<CreateMultipleBooksRequest.Author>;
        setAuthorsList(value: Array<CreateMultipleBooksRequest.Author>): Item;
        addAuthors(value?: CreateMultipleBooksRequest.Author, index?: number): CreateMultipleBooksRequest.Author;

        clearCategoriesList(): void;
        getCategoriesList(): Array<CreateMultipleBooksRequest.Category>;
        setCategoriesList(value: Array<CreateMultipleBooksRequest.Category>): Item;
        addCategories(value?: CreateMultipleBooksRequest.Category, index?: number): CreateMultipleBooksRequest.Category;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Item.AsObject;
        static toObject(includeInstance: boolean, msg: Item): Item.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Item, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Item;
        static deserializeBinaryFromReader(message: Item, reader: jspb.BinaryReader): Item;
    }

    export namespace Item {
        export type AsObject = {
            title: string,
            description: string,
            isbn: string,
            thumbnailUrl: string,
            version: string,
            publisher: string,
            publishedOn: string,
            authorsList: Array<CreateMultipleBooksRequest.Author.AsObject>,
            categoriesList: Array<CreateMultipleBooksRequest.Category.AsObject>,
        }
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

    getPublishedOn(): string;
    setPublishedOn(value: string): BookResponse;


    hasPublisher(): boolean;
    clearPublisher(): void;
    getPublisher(): BookResponse.Publisher | undefined;
    setPublisher(value?: BookResponse.Publisher): BookResponse;

    clearAuthorsList(): void;
    getAuthorsList(): Array<BookResponse.Author>;
    setAuthorsList(value: Array<BookResponse.Author>): BookResponse;
    addAuthors(value?: BookResponse.Author, index?: number): BookResponse.Author;

    clearCategoriesList(): void;
    getCategoriesList(): Array<BookResponse.Category>;
    setCategoriesList(value: Array<BookResponse.Category>): BookResponse;
    addCategories(value?: BookResponse.Category, index?: number): BookResponse.Category;

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
        publishedOn: string,
        publisher?: BookResponse.Publisher.AsObject,
        authorsList: Array<BookResponse.Author.AsObject>,
        categoriesList: Array<BookResponse.Category.AsObject>,
        createdAt: string,
        updatedAt: string,
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

    export class Publisher extends jspb.Message { 
        getId(): number;
        setId(value: number): Publisher;

        getName(): string;
        setName(value: string): Publisher;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Publisher.AsObject;
        static toObject(includeInstance: boolean, msg: Publisher): Publisher.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Publisher, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Publisher;
        static deserializeBinaryFromReader(message: Publisher, reader: jspb.BinaryReader): Publisher;
    }

    export namespace Publisher {
        export type AsObject = {
            id: number,
            name: string,
        }
    }

}

export class BookListResponse extends jspb.Message { 
    clearItemsList(): void;
    getItemsList(): Array<BookListResponse.Item>;
    setItemsList(value: Array<BookListResponse.Item>): BookListResponse;
    addItems(value?: BookListResponse.Item, index?: number): BookListResponse.Item;


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
        itemsList: Array<BookListResponse.Item.AsObject>,
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

    export class Publisher extends jspb.Message { 
        getId(): number;
        setId(value: number): Publisher;

        getName(): string;
        setName(value: string): Publisher;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Publisher.AsObject;
        static toObject(includeInstance: boolean, msg: Publisher): Publisher.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Publisher, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Publisher;
        static deserializeBinaryFromReader(message: Publisher, reader: jspb.BinaryReader): Publisher;
    }

    export namespace Publisher {
        export type AsObject = {
            id: number,
            name: string,
        }
    }

    export class Item extends jspb.Message { 
        getId(): number;
        setId(value: number): Item;

        getTitle(): string;
        setTitle(value: string): Item;

        getDescription(): string;
        setDescription(value: string): Item;

        getIsbn(): string;
        setIsbn(value: string): Item;

        getThumbnailUrl(): string;
        setThumbnailUrl(value: string): Item;

        getVersion(): string;
        setVersion(value: string): Item;

        getPublishedOn(): string;
        setPublishedOn(value: string): Item;


        hasPublisher(): boolean;
        clearPublisher(): void;
        getPublisher(): BookListResponse.Publisher | undefined;
        setPublisher(value?: BookListResponse.Publisher): Item;

        clearAuthorsList(): void;
        getAuthorsList(): Array<BookListResponse.Author>;
        setAuthorsList(value: Array<BookListResponse.Author>): Item;
        addAuthors(value?: BookListResponse.Author, index?: number): BookListResponse.Author;

        clearCategoriesList(): void;
        getCategoriesList(): Array<BookListResponse.Category>;
        setCategoriesList(value: Array<BookListResponse.Category>): Item;
        addCategories(value?: BookListResponse.Category, index?: number): BookListResponse.Category;

        getCreatedAt(): string;
        setCreatedAt(value: string): Item;

        getUpdatedAt(): string;
        setUpdatedAt(value: string): Item;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Item.AsObject;
        static toObject(includeInstance: boolean, msg: Item): Item.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Item, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Item;
        static deserializeBinaryFromReader(message: Item, reader: jspb.BinaryReader): Item;
    }

    export namespace Item {
        export type AsObject = {
            id: number,
            title: string,
            description: string,
            isbn: string,
            thumbnailUrl: string,
            version: string,
            publishedOn: string,
            publisher?: BookListResponse.Publisher.AsObject,
            authorsList: Array<BookListResponse.Author.AsObject>,
            categoriesList: Array<BookListResponse.Category.AsObject>,
            createdAt: string,
            updatedAt: string,
        }
    }

}
