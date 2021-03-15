export interface IBookOutput {
  id: number
  publisherId: number
  title: string
  description: string
  isbn: string
  thumbnailUrl: string
  version: string
  publishedOn: string
  createdAt: string
  updatedAt: string
  authors: Array<IBookOutputAuthor>
  categories: Array<IBookOutputCategory>
}

export interface IBookOutputAuthor {
  id: number
  name: string
}

export interface IBookOutputCategory {
  id: number
  name: string
}
