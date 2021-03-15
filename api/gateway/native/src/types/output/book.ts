export interface IBookOutput {
  id: number
  title: string
  description: string
  isbn: string
  thumbnailUrl: string
  version: string
  publishedOn: string
  publisher: IBookOutputPublisher
  authors: Array<IBookOutputAuthor>
  categories: Array<IBookOutputCategory>
  createdAt: string
  updatedAt: string
}

export interface IBookOutputPublisher {
  id: number
  name: string
}

export interface IBookOutputAuthor {
  id: number
  name: string
}

export interface IBookOutputCategory {
  id: number
  name: string
}
