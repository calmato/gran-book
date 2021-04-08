export interface IShowBookInput {
  isbn: string
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
  authors: Array<IBookAuthorInput>
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
  authors: Array<IBookAuthorInput>
}

export interface IBookAuthorInput {
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
