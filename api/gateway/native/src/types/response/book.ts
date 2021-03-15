export interface IBookResponse {
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
  authors: Array<IBookResponseAuthor>
  categories: Array<IBookResponseCategory>
}

export interface IBookResponseAuthor {
  id: number
  name: string
}

export interface IBookResponseCategory {
  id: number
  name: string
}
