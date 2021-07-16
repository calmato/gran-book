/**
 * Book - 書籍情報のレスポンス
 */
export interface IBookV2Response {
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
  reviews: Array<IBookV2ResponseReview>
  reviewLimit: number
  reviewOffset: number
  reviewTotal: number
}

export interface IBookV2ResponseReview {
  id: number
  impression: string
  createdAt: string
  updatedAt: string
  user: IBookV2ResponseUser
}

export interface IBookV2ResponseUser {
  id: string
  username: string
  thumbnailUrl: string
}
