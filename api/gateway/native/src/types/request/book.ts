export interface ICreateBookRequest {
  title: string
  authors: Array<string>
  publisher: string
  publishedDate: string
  description: string
  industryIdentifiers: Array<IBookIdentifier>
  categories: Array<string>
  contentVersion: string
  imageLinks: IBookImage
}

export interface IBookIdentifier {
  type: string
  identifier: string
}

export interface IBookImage {
  smallThumbnail?: string
  thumbnail: string
}
