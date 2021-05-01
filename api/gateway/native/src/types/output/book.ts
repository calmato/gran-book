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
  authors: Array<IBookOutputAuthor>
}

export interface IBookOutputAuthor {
  name: string
  nameKana: string
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
