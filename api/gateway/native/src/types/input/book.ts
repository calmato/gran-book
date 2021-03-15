export interface ICreateBookInput {
  title: string
  description: string
  isbn: string
  thumbnailURL: string
  version: string
  publisher: string
  publishedOn: string
  authors: Array<string>
  categories: Array<string>
}
