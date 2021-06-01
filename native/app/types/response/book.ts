export interface IBookResponse {
  limit: number;
  offset: number;
  total: number;
  books: Array<IBook>;
}

export interface IBook {
  id: number;
  isbn: string;
  author: string;
  authorKana: string;
  title: string;
  description: string;
  titleKana: string;
  thumbnailUrl: string;
  rakutenUrl: string;
  size: string;
  publishedOn: string;
  publisher: string;
  createdAt: string; // TODO: 日付の扱いどうするか？
  updatedAt: string;
  bookshelf?: IBookshelf
}

interface IBookshelf {
  id?: number;
  status: string;
  readOn?: string;
  impression?: string;
  createdAt?: string; // TODO: 日付の扱いどうするか？
  updatedAt?: string;
}
