import { IBook } from '~/types/response';

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

/**
 * 読んでいる本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 読んでいる本の一覧
 */
export function readingBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.status == 0;
  });
}

/**
 * 読んだ本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 読んだ本の一覧
 */
export function readBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.status == 1;
  });
}

/**
 * 積読本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 積読本の一覧
 */
export function stackBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.status == 2;
  });
}

/**
 * 手放したい本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 手放したい本の一覧
 */
export function releaseBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.status == 3;
  });
}

/**
 * 欲しい本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 欲しい本の一覧
 */
export function wantBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.status == 4;
  });
}
