/*
 * 本棚の書籍詳細レスポンス
 */
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
  size: string
  author: string
  authorKana: string
  createdAt: string
  updatedAt: string
  bookshelf?: IBookshelfResponseBookshelf
  reviews: Array<IBookshelfResponseReview>
  reviewLimit: number
  reviewOffset: number
  reviewTotal: number
}

export interface IBookshelfResponseBookshelf {
  id: number
  status: string
  readOn: string
  impression: string
  createdAt: string
  updatedAt: string
}

export interface IBookshelfResponseReview {
  id: number
  impression: number
  createdAt: string
  updatedAt: string
  user: IBookshelfResponseUserOnReview
}

export interface IBookshelfResponseUserOnReview {
  id: string
  username: string
  thumbnailUrl: string
}

/*
 * 本棚の書籍一覧レスポンス
 */
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
  size: string
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
  impression: string
  createdAt: string
  updatedAt: string
}
