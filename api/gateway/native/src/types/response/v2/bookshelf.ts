/**
 * Bookshelf - 書籍情報の取得レスポンス
 */
export interface IBookshelfV2Response {
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
  bookshelf?: IBookshelfV2ResponseBookshelf
  reviews: Array<IBookshelfV2ResponseReview>
  reviewLimit: number
  reviewOffset: number
  reviewTotal: number
}

export interface IBookshelfV2ResponseBookshelf {
  status: string
  readOn: string
  reviewId?: number
  createdAt: string
  updatedAt: string
}

export interface IBookshelfV2ResponseReview {
  id: number
  impression: string
  createdAt: string
  updatedAt: string
  user: IBookshelfV2ResponseUser
}

export interface IBookshelfV2ResponseUser {
  id: string
  username: string
  thumbnailUrl: string
}

/**
 * BookshelfList - 書籍情報一覧の取得レスポンス
 */
export interface IBookshelfListV2Response {
  books: Array<IBookshelfListV2ResponseBook>
  limit: number
  offset: number
  total: number
}

export interface IBookshelfListV2ResponseBook {
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
  bookshelf?: IBookshelfListV2ResponseBookshelf
}

export interface IBookshelfListV2ResponseBookshelf {
  status: string
  readOn: string
  reviewId?: number
  createdAt: string
  updatedAt: string
}
