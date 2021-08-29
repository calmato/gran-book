import { BookshelfListV1Response } from '~/types/api/bookshelf_apiv1_response_pb';
import { ViewBooks } from '~/types/models/book';

// Model
export interface Model {
  readonly books: BookshelfListV1Response.Book.AsObject[];
}

export const initialState: Model = {
  books: [] as BookshelfListV1Response.Book.AsObject[],
};

// input
export interface BookValues {
  books: BookshelfListV1Response.Book.AsObject[];
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
function readingBooks(
  books: Array<BookshelfListV1Response.Book.AsObject>,
): BookshelfListV1Response.Book.AsObject[] {
  return books.filter((book: BookshelfListV1Response.Book.AsObject) => {
    return book.bookshelf?.status === 'reading';
  });
}

/**
 * 読んだ本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 読んだ本の一覧
 */
function readBooks(
  books: BookshelfListV1Response.Book.AsObject[],
): BookshelfListV1Response.Book.AsObject[] {
  return books.filter((book: BookshelfListV1Response.Book.AsObject) => {
    return book.bookshelf?.status === 'read';
  });
}

/**
 * 積読本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 積読本の一覧
 */
function stackBooks(
  books: BookshelfListV1Response.Book.AsObject[],
): BookshelfListV1Response.Book.AsObject[] {
  return books.filter((book: BookshelfListV1Response.Book.AsObject) => {
    return book.bookshelf?.status === 'stack';
  });
}

/**
 * 手放したい本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 手放したい本の一覧
 */
function releaseBooks(
  books: BookshelfListV1Response.Book.AsObject[],
): BookshelfListV1Response.Book.AsObject[] {
  return books.filter((book: BookshelfListV1Response.Book.AsObject) => {
    return book.bookshelf?.status === 'release';
  });
}

/**
 * 欲しい本の一覧を返すgetter関数
 * @param books 登録されている全ての本
 * @returns 欲しい本の一覧
 */
function wantBooks(
  books: BookshelfListV1Response.Book.AsObject[],
): BookshelfListV1Response.Book.AsObject[] {
  return books.filter((book: BookshelfListV1Response.Book.AsObject) => {
    return book.bookshelf?.status === 'want';
  });
}
