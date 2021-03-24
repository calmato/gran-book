export interface IBookOutput {
  id: number
  title: string
  description: string
  isbn: string
  thumbnailUrl: string
  version: string
  publisher: string
  publishedOn: string
  authors: Array<string>
  categories: Array<string>
  createdAt: string
  updatedAt: string
}

export interface IBookListOutput {
  books: Array<IBookOutput>
}
