export interface IBookResponse {
  id: number
  title: string
  description: string
  isbn: string
  thumbnailUrl: string
  version: string
  publishedOn: string
  publisher: string
  authors: Array<string>
  categories: Array<string>
  createdAt: string
  updatedAt: string
}
