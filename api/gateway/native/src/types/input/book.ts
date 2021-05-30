export interface IListBookByBookIdsInput {
  bookIds: Array<number>
}

export interface IListBookshelfInput {
  userId: string
  limit: number
  offset: number
}

export interface IListBookReviewInput {
  bookId: number
  limit: number
  offset: number
  by: string
  direction: string
}

export interface IListUserReviewInput {
  userId: string
  limit: number
  offset: number
  by: string
  direction: string
}

export interface IGetBookInput {
  bookId: number
}

export interface IGetBookByIsbnInput {
  isbn: string
}

export interface IGetBookshelfInput {
  userId: string
  bookId: number
}

export interface IGetReviewInput {
  reviewId: number
}

export interface ICreateBookInput {
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
  authors: Array<IBookInputAuthor>
}

export interface IUpdateBookInput {
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
  authors: Array<IBookInputAuthor>
}

export interface IBookInputAuthor {
  name: string
  nameKana: string
}

export interface IReadBookshelfInput {
  userId: string
  bookId: number
  impression: string
  readOn: string
}

export interface IReadingBookshelfInput {
  userId: string
  bookId: number
}

export interface IStackBookshelfInput {
  userId: string
  bookId: number
}

export interface IWantBookshelfInput {
  userId: string
  bookId: number
}

export interface IReleaseBookshelfInput {
  userId: string
  bookId: number
}

export interface IDeleteBookshelfInput {
  userId: string
  bookId: number
}
