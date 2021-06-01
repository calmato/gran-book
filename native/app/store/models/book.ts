import { ViewBooks } from '~/types/models/book';
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
  books: IBook[];
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
 * stateの本の情報を整形して取得するfilter関数
 * @param model state
 * @returns コンポーネントで表示に使用する書籍情報
 */
export function filterBooks(model: Model): ViewBooks {
  return {
    reading: readingBooks(model.books),
    read: readBooks(model.books),
    stack: stackBooks(model.books),
    release: releaseBooks(model.books),
    want: wantBooks(model.books),
  };
}

/**
 * 読んでいる本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 読んでいる本の一覧
 */
function readingBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.bookshelf?.status === 'read';
  });
}

/**
 * 読んだ本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 読んだ本の一覧
 */
function readBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.bookshelf?.status === 'reading';
  });
}

/**
 * 積読本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 積読本の一覧
 */
function stackBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.bookshelf?.status === 'stack';
  });
}

/**
 * 手放したい本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 手放したい本の一覧
 */
function releaseBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.bookshelf?.status === 'release';
  });
}

/**
 * 欲しい本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 欲しい本の一覧
 */
function wantBooks(books: IBook[]): IBook[] {
  return books.filter((book: IBook) => {
    return book.bookshelf?.status === 'want';
  });
}
