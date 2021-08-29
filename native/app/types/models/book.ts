import { BookshelfListV1Response } from '../api/bookshelf_apiv1_response_pb';

/**
 * コンポーネントで表示に使用する書籍のデータ構造
 */
export interface ViewBooks {
  reading: BookshelfListV1Response.Book.AsObject[];
  read: BookshelfListV1Response.Book.AsObject[];
  stack: BookshelfListV1Response.Book.AsObject[];
  release: BookshelfListV1Response.Book.AsObject[];
  want: BookshelfListV1Response.Book.AsObject[];
}
