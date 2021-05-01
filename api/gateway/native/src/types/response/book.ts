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
