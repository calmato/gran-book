export interface IBookshelfResponse {
  id: number
  title: string
  titleKana: string
  description: string
  isbn: string
  publisher: string
  publishedOn: string
  thumbnailUrl: string
  rakutenUrl: string
  rakutenGenreId: string
  author: string
  authorKana: string
  createdAt: string
  updatedAt: string
  bookshelf?: IBookshelfResponseBookshelf
}

export interface IBookshelfResponseBookshelf {
  id: number
  status: string
  readOn: string
  impression: string
  createdAt: string
  updatedAt: string
}

export interface IBookshelfListResponse {
  books: Array<IBookshelfListResponseBook>
  limit: number
  offset: number
  total: number
}

export interface IBookshelfListResponseBook {
  id: number
  title: string
  titleKana: string
  description: string
  isbn: string
  publisher: string
  publishedOn: string
  thumbnailUrl: string
  rakutenUrl: string
  rakutenGenreId: string
  author: string
  authorKana: string
  createdAt: string
  updatedAt: string
  bookshelf?: IBookshelfListResponseBookshelf
}

export interface IBookshelfListResponseBookshelf {
  id: number
  status: string
  readOn: string
  createdAt: string
  updatedAt: string
}
