export interface IBookshelfResponse {
  id: number
  bookId: number
  userId: string
  status: number
  impression: string
  readOn: string
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
  status: number
  impression: string
  readOn: string
  createdAt: string
  updatedAt: string
  detail: IBookshelfListResponseDetail
}

export interface IBookshelfListResponseDetail {
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
}
