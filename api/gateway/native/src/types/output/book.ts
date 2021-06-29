/*
 * 書籍詳細アウトプット
 */
export interface IBookOutput {
  id: number
  title: string
  titleKana: string
  description: string
  isbn: string
  publisher: string
  publishedOn: string
  thumbnailUrl: string
  rakutenUrl: string
  rakutenSize: string
  rakutenGenreId: string
  createdAt: string
  updatedAt: string
  authors: Array<IBookOutputAuthor>
}

export interface IBookOutputAuthor {
  name: string
  nameKana: string
}

/*
 * 書籍一覧(ハッシュ)アウトプット
 */
export interface IBookHashOutput {
  [key: number]: IBookHashOutputBook
}

export interface IBookHashOutputBook {
  id: number
  title: string
  titleKana: string
  description: string
  isbn: string
  publisher: string
  publishedOn: string
  thumbnailUrl: string
  rakutenUrl: string
  rakutenSize: string
  rakutenGenreId: string
  createdAt: string
  updatedAt: string
  authors: Array<IBookOutputAuthor>
}

export interface IBookHashOutputAuthor {
  name: string
  nameKana: string
}

/*
 * 本棚の書籍詳細アウトプット
 */
export interface IBookshelfOutput {
  id: number
  bookId: number
  userId: string
  status: number
  readOn: string
  createdAt: string
  updatedAt: string
  book: IBookshelfOutputBook
  myReview: IBookshelfOutputReview
}

export interface IBookshelfOutputBook {
  id: number
  title: string
  titleKana: string
  description: string
  isbn: string
  publisher: string
  publishedOn: string
  thumbnailUrl: string
  rakutenUrl: string
  rakutenSize: string
  rakutenGenreId: string
  createdAt: string
  updatedAt: string
  authors: Array<IBookshelfOutputAuthor>
  reviews: Array<IBookshelfOutputReview>
}

export interface IBookshelfOutputAuthor {
  name: string
  nameKana: string
}

export interface IBookshelfOutputReview {
  id: number
  userId: string
  score: number
  impression: string
  createdAt: string
  updatedAt: string
}

/*
 * 本棚の書籍一覧アウトプット
 */
export interface IBookshelfListOutput {
  bookshelves: Array<IBookshelfListOutputBookshelf>
  limit: number
  offset: number
  total: number
}

export interface IBookshelfListOutputAuthor {
  name: string
  nameKana: string
}

export interface IBookshelfListOutputBook {
  id: number
  title: string
  titleKana: string
  description: string
  isbn: string
  publisher: string
  publishedOn: string
  thumbnailUrl: string
  rakutenUrl: string
  rakutenSize: string
  rakutenGenreId: string
  createdAt: string
  updatedAt: string
  authors: Array<IBookshelfListOutputAuthor>
}

export interface IBookshelfListOutputBookshelf {
  id: number
  bookId: number
  userId: string
  status: number
  readOn: string
  createdAt: string
  updatedAt: string
  book: IBookshelfListOutputBook
  myReview: IBookshelfListOutputReview
}

export interface IBookshelfListOutputReview {
  id: number
  userId: string
  score: number
  impression: string
  createdAt: string
  updatedAt: string
}

/*
 * 書籍レビュー詳細アウトプット
 */
export interface IReviewOutput {
  id: number
  bookId: number
  userId: string
  score: number
  impression: string
  createdAt: string
  updatedAt: string
}

/*
 * 書籍レビュー一覧アウトプット
 */
export interface IReviewListOutput {
  reviews: Array<IReviewListOutputReview>
  limit: number
  offset: number
  total: number
  order?: IReviewListOutputOrder
}

export interface IReviewListOutputReview {
  id: number
  bookId: number
  userId: string
  score: number
  impression: string
  createdAt: string
  updatedAt: string
}

export interface IReviewListOutputOrder {
  by: string
  direction: string
}
