export interface ICreateBookRequest {
  title: string
  titleKana: string
  itemCaption: string
  isbn: string
  publisherName: string
  salesDate: string
  largeImageUrl: string
  mediumImageUrl: string
  smallImageUrl: string
  itemUrl: string
  booksGenreId: string
  author: string
  authorKana: string
}

export interface IReadBookshelfRequest {
  impression: string
  readOn: string
}
