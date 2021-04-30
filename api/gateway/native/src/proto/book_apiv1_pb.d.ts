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

export class ShowBookRequest extends jspb.Message { 
    getIsbn(): string;
    setIsbn(value: string): ShowBookRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ShowBookRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ShowBookRequest): ShowBookRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ShowBookRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ShowBookRequest;
    static deserializeBinaryFromReader(message: ShowBookRequest, reader: jspb.BinaryReader): ShowBookRequest;
}

export namespace ShowBookRequest {
    export type AsObject = {
        isbn: string,
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
        bookId: number,
    }
}

export class DeleteBookshelfRequest extends jspb.Message { 
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
        bookId: number,
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

    hasBookshelf(): boolean;
    clearBookshelf(): void;
    getBookshelf(): BookResponse.Bookshelf | undefined;
    setBookshelf(value?: BookResponse.Bookshelf): BookResponse;
    clearAuthorsList(): void;
    getAuthorsList(): Array<BookResponse.Author>;
    setAuthorsList(value: Array<BookResponse.Author>): BookResponse;
    addAuthors(value?: BookResponse.Author, index?: number): BookResponse.Author;
    clearReviewsList(): void;
    getReviewsList(): Array<BookResponse.Review>;
    setReviewsList(value: Array<BookResponse.Review>): BookResponse;
    addReviews(value?: BookResponse.Review, index?: number): BookResponse.Review;

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
        bookshelf?: BookResponse.Bookshelf.AsObject,
        authorsList: Array<BookResponse.Author.AsObject>,
        reviewsList: Array<BookResponse.Review.AsObject>,
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

    export class Review extends jspb.Message { 
        getId(): number;
        setId(value: number): Review;
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
            userId: string,
            score: number,
            impression: string,
            createdAt: string,
            updatedAt: string,
        }
    }

    export class Bookshelf extends jspb.Message { 
        getId(): number;
        setId(value: number): Bookshelf;
        getStatus(): number;
        setStatus(value: number): Bookshelf;
        getReadOn(): string;
        setReadOn(value: string): Bookshelf;
        getCreatedAt(): string;
        setCreatedAt(value: string): Bookshelf;
        getUpdatedAt(): string;
        setUpdatedAt(value: string): Bookshelf;

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
            status: number,
            readOn: string,
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
        getImpression(): string;
        setImpression(value: string): Bookshelf;
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
            impression: string,
            readOn: string,
            createdAt: string,
            updatedAt: string,
            book?: BookshelfListResponse.Book.AsObject,
        }
    }

}
