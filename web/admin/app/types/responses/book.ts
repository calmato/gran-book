/*
 * URL: https://webservice.rakuten.co.jp/api/booksbooksearch/
 */
export interface IRakutenBookResponse {
  count: number
  page: number
  first: number
  last: number
  hits: number
  carrier: string
  pageCount: number
  Items: Array<IRakutenBookItem>
}

export interface IRakutenBookItem {
  title: string
  titleKana: string
  subTitle?: string
  subTitleKana?: string
  seriesName?: string
  seriesNameKana?: string
  contents?: string
  contentsKana?: string
  author: string
  authorKana: string
  publisherName: string
  size: number
  isbn: string
  itemCaption: string
  salesDate: string
  itemPrice: number
  itemUrl: string
  smallImageUrl: string
  mediumImageUrl: string
  largeImageUrl: string
  chirayomiUrl: string
  availability: number
  postageFlag: number
  limitedFlag: number
  reviewCount: number
  reviewAverage: number
  booksGenreId: string
}
