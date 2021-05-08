import { ImageURISource } from "react-native";

export interface IBookResponse {
  limit: number;                                                                                                                           │
  offset: number;                                                                                                                            │
  total: number;
  books: Array<IBook>;
}

export interface IBook {
  id: number;
  readOn: string;
  status: number;
  createdAt: string; // TODO: 日付の扱いどうするか？
  updatedAt: string;
  detail: IBooDetail
}

interface IBooDetail {
  id: number;
  isbn: string;
  author: string;
  authorKana: string;
  title: string;
  description: string;
  titleKana: string;
  thumbnailUrl: string;
  rakutenGenreId: string;
  rakutenUrl: string;
  createdAt: string; // TODO: 日付の扱いどうするか？
  updatedAt: string;
  publishedOn: string
  publisher: string;
}
