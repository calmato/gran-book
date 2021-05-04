export interface IBookResponse {
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
  bookshelf?: IBookResponseBookshelf
}

export interface IBookResponseBookshelf {
  id: number
  status: number
  impression: string
  readOn: string
  createdAt: string
  updatedAt: string
}

export interface IReviewResponse {
  id: number
  impression: string
  createdAt: string
  updatedAt: string
  book: IReviewResponseBook
  user: IReviewResponseUser
}

export interface IReviewResponseBook {
  id: number
  title: string
  thumbnailUrl: string
}

export interface IReviewResponseUser {
  id: number
  username: string
  thumbnailUrl: string
}

export interface IUserReviewListResponse {
  reviews: Array<IUserReviewListResponseReview>
  limit: number
  offset: number
  total: number
  order?: IUserReviewListResponseOrder
}

export interface IUserReviewListResponseReview {
  id: number
  impression: string
  createdAt: string
  updatedAt: string
  book: IUserReviewListResponseBook
}

export interface IUserReviewListResponseBook {
  id: number
  title: string
  thumbnailUrl: string
}

export interface IUserReviewListResponseOrder {
  by: string
  direction: string
}

export interface IBookReviewListResponse {
  reviews: Array<IBookReviewListResponseReview>
  limit: number
  offset: number
  total: number
  order?: IBookReviewListResponseOrder
}

export interface IBookReviewListResponseReview {
  id: number
  impression: string
  createdAt: string
  updatedAt: string
  user: IBookReviewListResponseUser
}

export interface IBookReviewListResponseUser {
  id: number
  username: string
  thumbnailUrl: string
}

export interface IBookReviewListResponseOrder {
  by: string
  direction: string
}
