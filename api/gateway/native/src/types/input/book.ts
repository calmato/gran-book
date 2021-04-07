export interface ICreateAndUpdateBooksInput {
  items: Array<IBookItemInput>
}

export interface IBookItemInput {
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
