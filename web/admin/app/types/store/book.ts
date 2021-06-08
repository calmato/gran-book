export interface IBookState {
  books: Array<IBook>
}

export interface IBook {
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
}
