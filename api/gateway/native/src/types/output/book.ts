export interface IBookOutput {
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
  createdAt: string
  updatedAt: string
  bookshelf: IBookBookshelfOutput
  authors: Array<IBookAuthorOutput>
  reviews: Array<IBookBookshelfOutput>
}

export interface IBookAuthorOutput {
  name: string
  nameKana: string
}

export interface IBookReviewOutput {
  id: number
  userId: string
  score: number
  impression: string
  createdAt: string
  updatedAt: string
}

export interface IBookBookshelfOutput {
  id: number
  status: number
  readOn: string
  createdAt: string
  updatedAt: string
}

export interface IBookshelfOutput {
  id: number
  bookId: number
  userId: string
  status: number
  impression: string
  readOn: string
  createdAt: string
  updatedAt: string
}
