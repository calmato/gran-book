export interface IListBookshelfInput {
  userId: string
  limit: number
  offset: number
}

export interface IGetBookInput {
  isbn: string
}

export interface IGetBookshelfInput {
  userId: string
  bookId: number
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
  rakutenGenreId: string
  authors: Array<IBookInputAuthor>
}

export interface IBookInputAuthor {
  name: string
  nameKana: string
}

export interface IReadBookshelfInput {
  bookId: number
  impression: string
  readOn: string
}

export interface IReadingBookshelfInput {
  bookId: number
}

export interface IStackBookshelfInput {
  bookId: number
}

export interface IWantBookshelfInput {
  bookId: number
}

export interface IReleaseBookshelfInput {
  bookId: number
}

export interface IDeleteBookshelfInput {
  bookId: number
}
