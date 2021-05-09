import { IBook } from '~/types/response';
import { booksSampleData } from '~~/assets/sample/book';

// Model
export interface Model {
  readonly books: IBook[];
}

export const initialState: Model = {
  books: [],
};

// input
export interface BookValues {
  books: IBook[]
}

export function factory(): Model {
  return initialState;
}

export function setBooks(books: Model, values: BookValues): Model {
  return {
    ...books,
    ...values,
  };
}

export function readingBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.status == 0;
  });
}


/**
 * 読んだ本を返すgetter関数
 * @param books
 * @returns
 */
export function readBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.status == 1;
  });
}

export function stackBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.status == 2;
  });
}

export function releaseBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.status == 3;
  });
}

export function wantBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.status == 4;
  });
}
