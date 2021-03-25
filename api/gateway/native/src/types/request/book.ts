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

export interface ICreateAndUpdateBooksRequest {
  items: Array<IBookItem>
}

export interface IBookItem {
  volumeInfo: IBookVolumeInfo
}

export interface IBookVolumeInfo {
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
