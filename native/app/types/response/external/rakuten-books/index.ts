export interface ISearchResponse {
  Items: ISearchResultItem[];
  pageCount: number;
  hits: number;
  last: number;
  count: number;
  page: number;
  carrier: number;
  GenreInformation: any[];
  first: number;
}

export interface ISearchResultItem {
  limitedFlag: number;
  authorKana: string;
  author: string;
  subTitle: string;
  seriesNameKana: string;
  title: string;
  subTitleKana: string;
  itemCaption: string;
  publisherName: string;
  listPrice: number;
  isbn: string;
  largeImageUrl: string;
  mediumImageUrl: string;
  titleKana: string;
  availability: string;
  postageFlag: number;
  salesDate: string;
  contents: string;
  smallImageUrl: string;
  discountPrice: number;
  itemPrice: number;
  size: string;
  booksGenreId: string;
  affiliateUrl: string;
  seriesName: string;
  reviewCount: number;
  reviewAverage: string;
  discountRate: number;
  chirayomiUrl: string;
  itemUrl: string;
}

export interface IErrorResponse {
  error: string;
  error_description: string;
}
