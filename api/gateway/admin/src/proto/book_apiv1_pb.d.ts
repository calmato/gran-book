// package: proto
// file: proto/book_apiv1.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class EmptyBook extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EmptyBook.AsObject;
    static toObject(includeInstance: boolean, msg: EmptyBook): EmptyBook.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EmptyBook, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EmptyBook;
    static deserializeBinaryFromReader(message: EmptyBook, reader: jspb.BinaryReader): EmptyBook;
}

export namespace EmptyBook {
    export type AsObject = {
    }
}

export class ListBookByBookIdsRequest extends jspb.Message { 
    clearBookIdsList(): void;
    getBookIdsList(): Array<number>;
    setBookIdsList(value: Array<number>): ListBookByBookIdsRequest;
    addBookIds(value: number, index?: number): number;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListBookByBookIdsRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListBookByBookIdsRequest): ListBookByBookIdsRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListBookByBookIdsRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListBookByBookIdsRequest;
    static deserializeBinaryFromReader(message: ListBookByBookIdsRequest, reader: jspb.BinaryReader): ListBookByBookIdsRequest;
}

export namespace ListBookByBookIdsRequest {
    export type AsObject = {
        bookIdsList: Array<number>,
    }
}

export class ListBookshelfRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): ListBookshelfRequest;
    getLimit(): number;
    setLimit(value: number): ListBookshelfRequest;
    getOffset(): number;
    setOffset(value: number): ListBookshelfRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListBookshelfRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListBookshelfRequest): ListBookshelfRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListBookshelfRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListBookshelfRequest;
    static deserializeBinaryFromReader(message: ListBookshelfRequest, reader: jspb.BinaryReader): ListBookshelfRequest;
}

export namespace ListBookshelfRequest {
    export type AsObject = {
        userId: string,
        limit: number,
        offset: number,
    }
}

export class ListBookReviewRequest extends jspb.Message { 
    getBookId(): number;
    setBookId(value: number): ListBookReviewRequest;
    getLimit(): number;
    setLimit(value: number): ListBookReviewRequest;
    getOffset(): number;
    setOffset(value: number): ListBookReviewRequest;

    hasOrder(): boolean;
    clearOrder(): void;
    getOrder(): ListBookReviewRequest.Order | undefined;
    setOrder(value?: ListBookReviewRequest.Order): ListBookReviewRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListBookReviewRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListBookReviewRequest): ListBookReviewRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListBookReviewRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListBookReviewRequest;
    static deserializeBinaryFromReader(message: ListBookReviewRequest, reader: jspb.BinaryReader): ListBookReviewRequest;
}

export namespace ListBookReviewRequest {
    export type AsObject = {
        bookId: number,
        limit: number,
        offset: number,
        order?: ListBookReviewRequest.Order.AsObject,
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

export class ListUserReviewRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): ListUserReviewRequest;
    getLimit(): number;
    setLimit(value: number): ListUserReviewRequest;
    getOffset(): number;
    setOffset(value: number): ListUserReviewRequest;

    hasOrder(): boolean;
    clearOrder(): void;
    getOrder(): ListUserReviewRequest.Order | undefined;
    setOrder(value?: ListUserReviewRequest.Order): ListUserReviewRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListUserReviewRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListUserReviewRequest): ListUserReviewRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListUserReviewRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListUserReviewRequest;
    static deserializeBinaryFromReader(message: ListUserReviewRequest, reader: jspb.BinaryReader): ListUserReviewRequest;
}

export namespace ListUserReviewRequest {
    export type AsObject = {
        userId: string,
        limit: number,
        offset: number,
        order?: ListUserReviewRequest.Order.AsObject,
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

export class GetBookRequest extends jspb.Message { 
    getIsbn(): string;
    setIsbn(value: string): GetBookRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetBookRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetBookRequest): GetBookRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetBookRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetBookRequest;
    static deserializeBinaryFromReader(message: GetBookRequest, reader: jspb.BinaryReader): GetBookRequest;
}

export namespace GetBookRequest {
    export type AsObject = {
        isbn: string,
    }
}

export class GetBookshelfRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): GetBookshelfRequest;
    getBookId(): number;
    setBookId(value: number): GetBookshelfRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetBookshelfRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetBookshelfRequest): GetBookshelfRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetBookshelfRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetBookshelfRequest;
    static deserializeBinaryFromReader(message: GetBookshelfRequest, reader: jspb.BinaryReader): GetBookshelfRequest;
}

export namespace GetBookshelfRequest {
    export type AsObject = {
        userId: string,
        bookId: number,
    }
}

export class GetReviewRequest extends jspb.Message { 
    getReviewId(): number;
    setReviewId(value: number): GetReviewRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetReviewRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetReviewRequest): GetReviewRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetReviewRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetReviewRequest;
    static deserializeBinaryFromReader(message: GetReviewRequest, reader: jspb.BinaryReader): GetReviewRequest;
}

export namespace GetReviewRequest {
    export type AsObject = {
        reviewId: number,
    }
}

export class CreateBookRequest extends jspb.Message { 
    getTitle(): string;
    setTitle(value: string): CreateBookRequest;
    getTitleKana(): string;
    setTitleKana(value: string): CreateBookRequest;
    getDescription(): string;
    setDescription(value: string): CreateBookRequest;
    getIsbn(): string;
    setIsbn(value: string): CreateBookRequest;
    getPublisher(): string;
    setPublisher(value: string): CreateBookRequest;
    getPublishedOn(): string;
    setPublishedOn(value: string): CreateBookRequest;
    getThumbnailUrl(): string;
    setThumbnailUrl(value: string): CreateBookRequest;
    getRakutenUrl(): string;
    setRakutenUrl(value: string): CreateBookRequest;
    getRakutenGenreId(): string;
    setRakutenGenreId(value: string): CreateBookRequest;
    clearAuthorsList(): void;
    getAuthorsList(): Array<CreateBookRequest.Author>;
    setAuthorsList(value: Array<CreateBookRequest.Author>): CreateBookRequest;
    addAuthors(value?: CreateBookRequest.Author, index?: number): CreateBookRequest.Author;

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
        titleKana: string,
        description: string,
        isbn: string,
        publisher: string,
        publishedOn: string,
        thumbnailUrl: string,
        rakutenUrl: string,
        rakutenGenreId: string,
        authorsList: Array<CreateBookRequest.Author.AsObject>,
    }


    export class Author extends jspb.Message { 
        getName(): string;
        setName(value: string): Author;
        getNameKana(): string;
        setNameKana(value: string): Author;

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
            nameKana: string,
        }
    }

}

export class UpdateBookRequest extends jspb.Message { 
    getTitle(): string;
    setTitle(value: string): UpdateBookRequest;
    getTitleKana(): string;
    setTitleKana(value: string): UpdateBookRequest;
    getDescription(): string;
    setDescription(value: string): UpdateBookRequest;
    getIsbn(): string;
    setIsbn(value: string): UpdateBookRequest;
    getPublisher(): string;
    setPublisher(value: string): UpdateBookRequest;
    getPublishedOn(): string;
    setPublishedOn(value: string): UpdateBookRequest;
    getThumbnailUrl(): string;
    setThumbnailUrl(value: string): UpdateBookRequest;
    getRakutenUrl(): string;
    setRakutenUrl(value: string): UpdateBookRequest;
    getRakutenGenreId(): string;
    setRakutenGenreId(value: string): UpdateBookRequest;
    clearAuthorsList(): void;
    getAuthorsList(): Array<UpdateBookRequest.Author>;
    setAuthorsList(value: Array<UpdateBookRequest.Author>): UpdateBookRequest;
    addAuthors(value?: UpdateBookRequest.Author, index?: number): UpdateBookRequest.Author;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateBookRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateBookRequest): UpdateBookRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateBookRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateBookRequest;
    static deserializeBinaryFromReader(message: UpdateBookRequest, reader: jspb.BinaryReader): UpdateBookRequest;
}

export namespace UpdateBookRequest {
    export type AsObject = {
        title: string,
        titleKana: string,
        description: string,
        isbn: string,
        publisher: string,
        publishedOn: string,
        thumbnailUrl: string,
        rakutenUrl: string,
        rakutenGenreId: string,
        authorsList: Array<UpdateBookRequest.Author.AsObject>,
    }


    export class Author extends jspb.Message { 
        getName(): string;
        setName(value: string): Author;
        getNameKana(): string;
        setNameKana(value: string): Author;

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
            nameKana: string,
        }
    }

}

export class DeleteBookRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): DeleteBookRequest;
    getBookId(): number;
    setBookId(value: number): DeleteBookRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteBookRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteBookRequest): DeleteBookRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteBookRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteBookRequest;
    static deserializeBinaryFromReader(message: DeleteBookRequest, reader: jspb.BinaryReader): DeleteBookRequest;
}

export namespace DeleteBookRequest {
    export type AsObject = {
        userId: string,
        bookId: number,
    }
}

export class DeleteBookshelfRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): DeleteBookshelfRequest;
    getBookId(): number;
    setBookId(value: number): DeleteBookshelfRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteBookshelfRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteBookshelfRequest): DeleteBookshelfRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteBookshelfRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteBookshelfRequest;
    static deserializeBinaryFromReader(message: DeleteBookshelfRequest, reader: jspb.BinaryReader): DeleteBookshelfRequest;
}

export namespace DeleteBookshelfRequest {
    export type AsObject = {
        userId: string,
        bookId: number,
    }
}

export class ReadBookshelfRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): ReadBookshelfRequest;
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
        userId: string,
        bookId: number,
        impression: string,
        readOn: string,
    }
}

export class ReadingBookshelfRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): ReadingBookshelfRequest;
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
        userId: string,
        bookId: number,
    }
}

export class StackBookshelfRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): StackBookshelfRequest;
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
        userId: string,
        bookId: number,
    }
}

export class WantBookshelfRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): WantBookshelfRequest;
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
        userId: string,
        bookId: number,
    }
}

export class ReleaseBookshelfRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): ReleaseBookshelfRequest;
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
        userId: string,
        bookId: number,
    }
}

export class BookResponse extends jspb.Message { 
    getId(): number;
    setId(value: number): BookResponse;
    getTitle(): string;
    setTitle(value: string): BookResponse;
    getTitleKana(): string;
    setTitleKana(value: string): BookResponse;
    getDescription(): string;
    setDescription(value: string): BookResponse;
    getIsbn(): string;
    setIsbn(value: string): BookResponse;
    getPublisher(): string;
    setPublisher(value: string): BookResponse;
    getPublishedOn(): string;
    setPublishedOn(value: string): BookResponse;
    getThumbnailUrl(): string;
    setThumbnailUrl(value: string): BookResponse;
    getRakutenUrl(): string;
    setRakutenUrl(value: string): BookResponse;
    getRakutenGenreId(): string;
    setRakutenGenreId(value: string): BookResponse;
    getCreatedAt(): string;
    setCreatedAt(value: string): BookResponse;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): BookResponse;
    clearAuthorsList(): void;
    getAuthorsList(): Array<BookResponse.Author>;
    setAuthorsList(value: Array<BookResponse.Author>): BookResponse;
    addAuthors(value?: BookResponse.Author, index?: number): BookResponse.Author;

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
        titleKana: string,
        description: string,
        isbn: string,
        publisher: string,
        publishedOn: string,
        thumbnailUrl: string,
        rakutenUrl: string,
        rakutenGenreId: string,
        createdAt: string,
        updatedAt: string,
        authorsList: Array<BookResponse.Author.AsObject>,
    }


    export class Author extends jspb.Message { 
        getName(): string;
        setName(value: string): Author;
        getNameKana(): string;
        setNameKana(value: string): Author;

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
            nameKana: string,
        }
    }

}

export class BookListResponse extends jspb.Message { 
    clearBooksList(): void;
    getBooksList(): Array<BookListResponse.Book>;
    setBooksList(value: Array<BookListResponse.Book>): BookListResponse;
    addBooks(value?: BookListResponse.Book, index?: number): BookListResponse.Book;
    getLimit(): number;
    setLimit(value: number): BookListResponse;
    getOffset(): number;
    setOffset(value: number): BookListResponse;
    getTotal(): number;
    setTotal(value: number): BookListResponse;

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
        limit: number,
        offset: number,
        total: number,
    }


    export class Author extends jspb.Message { 
        getName(): string;
        setName(value: string): Author;
        getNameKana(): string;
        setNameKana(value: string): Author;

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
            nameKana: string,
        }
    }

    export class Book extends jspb.Message { 
        getId(): number;
        setId(value: number): Book;
        getTitle(): string;
        setTitle(value: string): Book;
        getTitleKana(): string;
        setTitleKana(value: string): Book;
        getDescription(): string;
        setDescription(value: string): Book;
        getIsbn(): string;
        setIsbn(value: string): Book;
        getPublisher(): string;
        setPublisher(value: string): Book;
        getPublishedOn(): string;
        setPublishedOn(value: string): Book;
        getThumbnailUrl(): string;
        setThumbnailUrl(value: string): Book;
        getRakutenUrl(): string;
        setRakutenUrl(value: string): Book;
        getRakutenGenreId(): string;
        setRakutenGenreId(value: string): Book;
        getCreatedAt(): string;
        setCreatedAt(value: string): Book;
        getUpdatedAt(): string;
        setUpdatedAt(value: string): Book;
        clearAuthorsList(): void;
        getAuthorsList(): Array<BookListResponse.Author>;
        setAuthorsList(value: Array<BookListResponse.Author>): Book;
        addAuthors(value?: BookListResponse.Author, index?: number): BookListResponse.Author;

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
            titleKana: string,
            description: string,
            isbn: string,
            publisher: string,
            publishedOn: string,
            thumbnailUrl: string,
            rakutenUrl: string,
            rakutenGenreId: string,
            createdAt: string,
            updatedAt: string,
            authorsList: Array<BookListResponse.Author.AsObject>,
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

    hasReview(): boolean;
    clearReview(): void;
    getReview(): BookshelfResponse.Review | undefined;
    setReview(value?: BookshelfResponse.Review): BookshelfResponse;

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
        readOn: string,
        createdAt: string,
        updatedAt: string,
        book?: BookshelfResponse.Book.AsObject,
        review?: BookshelfResponse.Review.AsObject,
    }


    export class Author extends jspb.Message { 
        getName(): string;
        setName(value: string): Author;
        getNameKana(): string;
        setNameKana(value: string): Author;

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
            nameKana: string,
        }
    }

    export class Book extends jspb.Message { 
        getId(): number;
        setId(value: number): Book;
        getTitle(): string;
        setTitle(value: string): Book;
        getTitleKana(): string;
        setTitleKana(value: string): Book;
        getDescription(): string;
        setDescription(value: string): Book;
        getIsbn(): string;
        setIsbn(value: string): Book;
        getPublisher(): string;
        setPublisher(value: string): Book;
        getPublishedOn(): string;
        setPublishedOn(value: string): Book;
        getThumbnailUrl(): string;
        setThumbnailUrl(value: string): Book;
        getRakutenUrl(): string;
        setRakutenUrl(value: string): Book;
        getRakutenGenreId(): string;
        setRakutenGenreId(value: string): Book;
        getCreatedAt(): string;
        setCreatedAt(value: string): Book;
        getUpdatedAt(): string;
        setUpdatedAt(value: string): Book;
        clearAuthorsList(): void;
        getAuthorsList(): Array<BookshelfResponse.Author>;
        setAuthorsList(value: Array<BookshelfResponse.Author>): Book;
        addAuthors(value?: BookshelfResponse.Author, index?: number): BookshelfResponse.Author;

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
            titleKana: string,
            description: string,
            isbn: string,
            publisher: string,
            publishedOn: string,
            thumbnailUrl: string,
            rakutenUrl: string,
            rakutenGenreId: string,
            createdAt: string,
            updatedAt: string,
            authorsList: Array<BookshelfResponse.Author.AsObject>,
        }
    }

    export class Review extends jspb.Message { 
        getScore(): number;
        setScore(value: number): Review;
        getImpression(): string;
        setImpression(value: string): Review;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Review.AsObject;
        static toObject(includeInstance: boolean, msg: Review): Review.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Review, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Review;
        static deserializeBinaryFromReader(message: Review, reader: jspb.BinaryReader): Review;
    }

    export namespace Review {
        export type AsObject = {
            score: number,
            impression: string,
        }
    }

}

export class BookshelfListResponse extends jspb.Message { 
    clearBookshelvesList(): void;
    getBookshelvesList(): Array<BookshelfListResponse.Bookshelf>;
    setBookshelvesList(value: Array<BookshelfListResponse.Bookshelf>): BookshelfListResponse;
    addBookshelves(value?: BookshelfListResponse.Bookshelf, index?: number): BookshelfListResponse.Bookshelf;
    getLimit(): number;
    setLimit(value: number): BookshelfListResponse;
    getOffset(): number;
    setOffset(value: number): BookshelfListResponse;
    getTotal(): number;
    setTotal(value: number): BookshelfListResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BookshelfListResponse.AsObject;
    static toObject(includeInstance: boolean, msg: BookshelfListResponse): BookshelfListResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BookshelfListResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BookshelfListResponse;
    static deserializeBinaryFromReader(message: BookshelfListResponse, reader: jspb.BinaryReader): BookshelfListResponse;
}

export namespace BookshelfListResponse {
    export type AsObject = {
        bookshelvesList: Array<BookshelfListResponse.Bookshelf.AsObject>,
        limit: number,
        offset: number,
        total: number,
    }


    export class Author extends jspb.Message { 
        getName(): string;
        setName(value: string): Author;
        getNameKana(): string;
        setNameKana(value: string): Author;

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
            nameKana: string,
        }
    }

    export class Book extends jspb.Message { 
        getId(): number;
        setId(value: number): Book;
        getTitle(): string;
        setTitle(value: string): Book;
        getTitleKana(): string;
        setTitleKana(value: string): Book;
        getDescription(): string;
        setDescription(value: string): Book;
        getIsbn(): string;
        setIsbn(value: string): Book;
        getPublisher(): string;
        setPublisher(value: string): Book;
        getPublishedOn(): string;
        setPublishedOn(value: string): Book;
        getThumbnailUrl(): string;
        setThumbnailUrl(value: string): Book;
        getRakutenUrl(): string;
        setRakutenUrl(value: string): Book;
        getRakutenGenreId(): string;
        setRakutenGenreId(value: string): Book;
        getCreatedAt(): string;
        setCreatedAt(value: string): Book;
        getUpdatedAt(): string;
        setUpdatedAt(value: string): Book;
        clearAuthorsList(): void;
        getAuthorsList(): Array<BookshelfListResponse.Author>;
        setAuthorsList(value: Array<BookshelfListResponse.Author>): Book;
        addAuthors(value?: BookshelfListResponse.Author, index?: number): BookshelfListResponse.Author;

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
            titleKana: string,
            description: string,
            isbn: string,
            publisher: string,
            publishedOn: string,
            thumbnailUrl: string,
            rakutenUrl: string,
            rakutenGenreId: string,
            createdAt: string,
            updatedAt: string,
            authorsList: Array<BookshelfListResponse.Author.AsObject>,
        }
    }

    export class Bookshelf extends jspb.Message { 
        getId(): number;
        setId(value: number): Bookshelf;
        getBookId(): number;
        setBookId(value: number): Bookshelf;
        getUserId(): string;
        setUserId(value: string): Bookshelf;
        getStatus(): number;
        setStatus(value: number): Bookshelf;
        getReadOn(): string;
        setReadOn(value: string): Bookshelf;
        getCreatedAt(): string;
        setCreatedAt(value: string): Bookshelf;
        getUpdatedAt(): string;
        setUpdatedAt(value: string): Bookshelf;

        hasBook(): boolean;
        clearBook(): void;
        getBook(): BookshelfListResponse.Book | undefined;
        setBook(value?: BookshelfListResponse.Book): Bookshelf;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Bookshelf.AsObject;
        static toObject(includeInstance: boolean, msg: Bookshelf): Bookshelf.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Bookshelf, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Bookshelf;
        static deserializeBinaryFromReader(message: Bookshelf, reader: jspb.BinaryReader): Bookshelf;
    }

    export namespace Bookshelf {
        export type AsObject = {
            id: number,
            bookId: number,
            userId: string,
            status: number,
            readOn: string,
            createdAt: string,
            updatedAt: string,
            book?: BookshelfListResponse.Book.AsObject,
        }
    }

}

export class ReviewResponse extends jspb.Message { 
    getId(): number;
    setId(value: number): ReviewResponse;
    getBookId(): number;
    setBookId(value: number): ReviewResponse;
    getUserId(): string;
    setUserId(value: string): ReviewResponse;
    getScore(): number;
    setScore(value: number): ReviewResponse;
    getImpression(): string;
    setImpression(value: string): ReviewResponse;
    getCreatedAt(): string;
    setCreatedAt(value: string): ReviewResponse;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): ReviewResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ReviewResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ReviewResponse): ReviewResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ReviewResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ReviewResponse;
    static deserializeBinaryFromReader(message: ReviewResponse, reader: jspb.BinaryReader): ReviewResponse;
}

export namespace ReviewResponse {
    export type AsObject = {
        id: number,
        bookId: number,
        userId: string,
        score: number,
        impression: string,
        createdAt: string,
        updatedAt: string,
    }
}

export class ReviewListResponse extends jspb.Message { 
    clearReviewsList(): void;
    getReviewsList(): Array<ReviewListResponse.Review>;
    setReviewsList(value: Array<ReviewListResponse.Review>): ReviewListResponse;
    addReviews(value?: ReviewListResponse.Review, index?: number): ReviewListResponse.Review;
    getLimit(): number;
    setLimit(value: number): ReviewListResponse;
    getOffset(): number;
    setOffset(value: number): ReviewListResponse;
    getTotal(): number;
    setTotal(value: number): ReviewListResponse;

    hasOrder(): boolean;
    clearOrder(): void;
    getOrder(): ReviewListResponse.Order | undefined;
    setOrder(value?: ReviewListResponse.Order): ReviewListResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ReviewListResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ReviewListResponse): ReviewListResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ReviewListResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ReviewListResponse;
    static deserializeBinaryFromReader(message: ReviewListResponse, reader: jspb.BinaryReader): ReviewListResponse;
}

export namespace ReviewListResponse {
    export type AsObject = {
        reviewsList: Array<ReviewListResponse.Review.AsObject>,
        limit: number,
        offset: number,
        total: number,
        order?: ReviewListResponse.Order.AsObject,
    }


    export class Review extends jspb.Message { 
        getId(): number;
        setId(value: number): Review;
        getBookId(): number;
        setBookId(value: number): Review;
        getUserId(): string;
        setUserId(value: string): Review;
        getScore(): number;
        setScore(value: number): Review;
        getImpression(): string;
        setImpression(value: string): Review;
        getCreatedAt(): string;
        setCreatedAt(value: string): Review;
        getUpdatedAt(): string;
        setUpdatedAt(value: string): Review;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Review.AsObject;
        static toObject(includeInstance: boolean, msg: Review): Review.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Review, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Review;
        static deserializeBinaryFromReader(message: Review, reader: jspb.BinaryReader): Review;
    }

    export namespace Review {
        export type AsObject = {
            id: number,
            bookId: number,
            userId: string,
            score: number,
            impression: string,
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
