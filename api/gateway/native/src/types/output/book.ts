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
  rakutenSize: string
  rakutenGenreId: string
  createdAt: string
  updatedAt: string
  authors: Array<IBookOutputAuthor>
}

export interface IBookOutputAuthor {
  name: string
  nameKana: string
}

export interface IBookHashOutput {
  [key: number]: IBookHashOutputBook
}

export interface IBookHashOutputBook {
  id: number
  title: string
  titleKana: string
  description: string
  isbn: string
  publisher: string
  publishedOn: string
  thumbnailUrl: string
  rakutenUrl: string
  rakutenSize: string
  rakutenGenreId: string
  createdAt: string
  updatedAt: string
  authors: Array<IBookOutputAuthor>
}

export interface IBookHashOutputAuthor {
  name: string
  nameKana: string
}

export interface IBookshelfOutput {
  id: number
  bookId: number
  userId: string
  status: number
  readOn: string
  createdAt: string
  updatedAt: string
  book: IBookshelfOutputBook
  review?: IBookshelfOutputReview
}

export interface IBookshelfOutputBook {
  id: number
  title: string
  titleKana: string
  description: string
  isbn: string
  publisher: string
  publishedOn: string
  thumbnailUrl: string
  rakutenUrl: string
  rakutenSize: string
  rakutenGenreId: string
  createdAt: string
  updatedAt: string
  authors: Array<IBookshelfOutputAuthor>
}

export interface IBookshelfOutputAuthor {
  name: string
  nameKana: string
}

export interface IBookshelfOutputReview {
  score: number
  impression: string
}

export interface IBookshelfListOutput {
  bookshelves: Array<IBookshelfListOutputBookshelf>
  limit: number
  offset: number
  total: number
}

export interface IBookshelfListOutputAuthor {
  name: string
  nameKana: string
}

export interface IBookshelfListOutputBook {
  id: number
  title: string
  titleKana: string
  description: string
  isbn: string
  publisher: string
  publishedOn: string
  thumbnailUrl: string
  rakutenUrl: string
  rakutenSize: string
  rakutenGenreId: string
  createdAt: string
  updatedAt: string
  authors: Array<IBookshelfListOutputAuthor>
}

export interface IBookshelfListOutputBookshelf {
  id: number
  bookId: number
  userId: string
  status: number
  readOn: string
  createdAt: string
  updatedAt: string
  book: IBookshelfListOutputBook
}

export interface IReviewOutput {
  id: number
  bookId: number
  userId: string
  score: number
  impression: string
  createdAt: string
  updatedAt: string
}

export interface IReviewListOutput {
  reviews: Array<IReviewListOutputReview>
  limit: number
  offset: number
  total: number
  order?: IReviewListOutputOrder
}

export interface IReviewListOutputReview {
  id: number
  bookId: number
  userId: string
  score: number
  impression: string
  createdAt: string
  updatedAt: string
}

export interface IReviewListOutputOrder {
  by: string
  direction: string
}
