import { Request } from 'express'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import { bookClient } from '~/plugins/grpc'
import {
  BookListResponse,
  BookResponse,
  BookshelfListResponse,
  BookshelfResponse,
  CreateBookRequest,
  DeleteBookshelfRequest,
  EmptyBook,
  GetBookByIsbnRequest,
  GetBookRequest,
  GetBookshelfRequest,
  GetReviewRequest,
  ListBookByBookIdsRequest,
  ListBookReviewRequest,
  ListBookshelfRequest,
  ListUserReviewRequest,
  ReadBookshelfRequest,
  ReadingBookshelfRequest,
  ReleaseBookshelfRequest,
  ReviewListResponse,
  ReviewResponse,
  StackBookshelfRequest,
  UpdateBookRequest,
  WantBookshelfRequest,
} from '~/proto/book_apiv1_pb'
import {
  IBookInputAuthor,
  ICreateBookInput,
  IDeleteBookshelfInput,
  IGetBookByIsbnInput,
  IGetBookInput,
  IGetBookshelfInput,
  IGetReviewInput,
  IListBookByBookIdsInput,
  IListBookReviewInput,
  IListBookshelfInput,
  IListUserReviewInput,
  IReadBookshelfInput,
  IReadingBookshelfInput,
  IReleaseBookshelfInput,
  IStackBookshelfInput,
  IUpdateBookInput,
  IWantBookshelfInput,
} from '~/types/input'
import {
  IBookHashOutput,
  IBookHashOutputAuthor,
  IBookHashOutputBook,
  IBookOutput,
  IBookOutputAuthor,
  IBookshelfListOutput,
  IBookshelfListOutputAuthor,
  IBookshelfListOutputBook,
  IBookshelfListOutputBookshelf,
  IBookshelfOutput,
  IBookshelfOutputAuthor,
  IBookshelfOutputBook,
  IBookshelfOutputReview,
  IReviewListOutput,
  IReviewListOutputOrder,
  IReviewListOutputReview,
  IReviewOutput,
} from '~/types/output'

export function listBookByBookIds(req: Request<any>, input: IListBookByBookIdsInput): Promise<IBookHashOutput> {
  const request = new ListBookByBookIdsRequest()
  const metadata = getGrpcMetadata(req)

  request.setBookIdsList(input.bookIds)

  return new Promise((resolve: (res: IBookHashOutput) => void, reject: (reason: Error) => void) => {
    bookClient.listBookByBookIds(request, metadata, (err: any, res: BookListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setBookHashOutput(res))
    })
  })
}

export function listBookshelf(req: Request<any>, input: IListBookshelfInput): Promise<IBookshelfListOutput> {
  const request = new ListBookshelfRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)
  request.setLimit(input.limit)
  request.setOffset(input.offset)

  return new Promise((resolve: (output: IBookshelfListOutput) => void, reject: (reason: Error) => void) => {
    bookClient.listBookshelf(request, metadata, (err: any, res: BookshelfListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IBookshelfListOutput = setBookshelfListOutput(res)
      resolve(output)
    })
  })
}

export function listUserReview(req: Request<any>, input: IListUserReviewInput): Promise<IReviewListOutput> {
  const request = new ListUserReviewRequest()
  const metadata = getGrpcMetadata(req)

  const order = new ListUserReviewRequest.Order()
  order.setBy(input.by)
  order.setDirection(input.direction)

  request.setUserId(input.userId)
  request.setLimit(input.limit)
  request.setOffset(input.offset)
  request.setOrder(order)

  return new Promise((resolve: (output: IReviewListOutput) => void, reject: (reason: Error) => void) => {
    bookClient.listUserReview(request, metadata, (err: any, res: ReviewListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IReviewListOutput = setReviewListOutput(res)
      resolve(output)
    })
  })
}

export function listBookReview(req: Request<any>, input: IListBookReviewInput): Promise<IReviewListOutput> {
  const request = new ListBookReviewRequest()
  const metadata = getGrpcMetadata(req)

  const order = new ListBookReviewRequest.Order()
  order.setBy(input.by)
  order.setDirection(input.direction)

  request.setBookId(input.bookId)
  request.setLimit(input.limit)
  request.setOffset(input.offset)
  request.setOrder(order)

  return new Promise((resolve: (output: IReviewListOutput) => void, reject: (reason: Error) => void) => {
    bookClient.listBookReview(request, metadata, (err: any, res: ReviewListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IReviewListOutput = setReviewListOutput(res)
      resolve(output)
    })
  })
}

export function getBook(req: Request<any>, input: IGetBookInput): Promise<IBookOutput> {
  const request = new GetBookRequest()
  const metadata = getGrpcMetadata(req)

  request.setId(input.bookId)

  return new Promise((resolve: (output: IBookOutput) => void, reject: (reason: Error) => void) => {
    bookClient.getBook(request, metadata, (err: any, res: BookResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IBookOutput = setBookOutput(res)
      resolve(output)
    })
  })
}

export function getBookByIsbn(req: Request<any>, input: IGetBookByIsbnInput): Promise<IBookOutput> {
  const request = new GetBookByIsbnRequest()
  const metadata = getGrpcMetadata(req)

  request.setIsbn(input.isbn)

  return new Promise((resolve: (output: IBookOutput) => void, reject: (reason: Error) => void) => {
    bookClient.getBookByIsbn(request, metadata, (err: any, res: BookResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IBookOutput = setBookOutput(res)
      resolve(output)
    })
  })
}

export function getBookshelf(req: Request<any>, input: IGetBookshelfInput): Promise<IBookshelfOutput> {
  const request = new GetBookshelfRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)
  request.setBookId(input.bookId)

  return new Promise((resolve: (output: IBookshelfOutput) => void, reject: (reason: Error) => void) => {
    bookClient.getBookshelf(request, metadata, (err: any, res: BookshelfResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IBookshelfOutput = setBookshelfOutput(res)
      resolve(output)
    })
  })
}

export function getReview(req: Request<any>, input: IGetReviewInput): Promise<IReviewOutput> {
  const request = new GetReviewRequest()
  const metadata = getGrpcMetadata(req)

  request.setReviewId(input.reviewId)

  return new Promise((resolve: (output: IReviewOutput) => void, reject: (reason: Error) => void) => {
    bookClient.getReview(request, metadata, (err: any, res: ReviewResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IReviewOutput = setReviewOutput(res)
      resolve(output)
    })
  })
}

export function createBook(req: Request<any>, input: ICreateBookInput): Promise<IBookOutput> {
  const request = new CreateBookRequest()
  const metadata = getGrpcMetadata(req)

  const authors = input.authors?.map((item: IBookInputAuthor) => {
    const author = new CreateBookRequest.Author()

    author.setName(item.name)
    author.setNameKana(item.nameKana)

    return author
  })

  request.setTitle(input.title)
  request.setTitleKana(input.titleKana)
  request.setDescription(input.description)
  request.setIsbn(input.isbn)
  request.setPublisher(input.publisher)
  request.setPublishedOn(input.publishedOn)
  request.setThumbnailUrl(input.thumbnailUrl)
  request.setRakutenUrl(input.rakutenUrl)
  request.setRakutenSize(input.rakutenSize)
  request.setRakutenGenreId(input.rakutenGenreId)
  request.setAuthorsList(authors)

  return new Promise((resolve: (output: IBookOutput) => void, reject: (reason: Error) => void) => {
    bookClient.createBook(request, metadata, (err: any, res: BookResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IBookOutput = setBookOutput(res)
      resolve(output)
    })
  })
}

export function updateBook(req: Request<any>, input: IUpdateBookInput): Promise<IBookOutput> {
  const request = new UpdateBookRequest()
  const metadata = getGrpcMetadata(req)

  const authors = input.authors?.map((item: IBookInputAuthor) => {
    const author = new UpdateBookRequest.Author()

    author.setName(item.name)
    author.setNameKana(item.nameKana)

    return author
  })

  request.setTitle(input.title)
  request.setTitleKana(input.titleKana)
  request.setDescription(input.description)
  request.setIsbn(input.isbn)
  request.setPublisher(input.publisher)
  request.setPublishedOn(input.publishedOn)
  request.setThumbnailUrl(input.thumbnailUrl)
  request.setRakutenUrl(input.rakutenUrl)
  request.setRakutenSize(input.rakutenSize)
  request.setRakutenGenreId(input.rakutenGenreId)
  request.setAuthorsList(authors)

  return new Promise((resolve: (output: IBookOutput) => void, reject: (reason: Error) => void) => {
    bookClient.updateBook(request, metadata, (err: any, res: BookResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IBookOutput = setBookOutput(res)
      resolve(output)
    })
  })
}

export function readBookshelf(req: Request<any>, input: IReadBookshelfInput): Promise<IBookshelfOutput> {
  const request = new ReadBookshelfRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)
  request.setBookId(input.bookId)
  request.setImpression(input.impression)
  request.setReadOn(input.readOn)

  return new Promise((resolve: (output: IBookshelfOutput) => void, reject: (reason: Error) => void) => {
    bookClient.readBookshelf(request, metadata, (err: any, res: BookshelfResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IBookshelfOutput = setBookshelfOutput(res)
      resolve(output)
    })
  })
}

export function readingBookshelf(req: Request<any>, input: IReadingBookshelfInput): Promise<IBookshelfOutput> {
  const request = new ReadingBookshelfRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)
  request.setBookId(input.bookId)

  return new Promise((resolve: (output: IBookshelfOutput) => void, reject: (reason: Error) => void) => {
    bookClient.readingBookshelf(request, metadata, (err: any, res: BookshelfResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IBookshelfOutput = setBookshelfOutput(res)
      resolve(output)
    })
  })
}

export function stackBookshelf(req: Request<any>, input: IStackBookshelfInput): Promise<IBookshelfOutput> {
  const request = new StackBookshelfRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)
  request.setBookId(input.bookId)

  return new Promise((resolve: (output: IBookshelfOutput) => void, reject: (reason: Error) => void) => {
    bookClient.stackBookshelf(request, metadata, (err: any, res: BookshelfResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IBookshelfOutput = setBookshelfOutput(res)
      resolve(output)
    })
  })
}

export function wantBookshelf(req: Request<any>, input: IWantBookshelfInput): Promise<IBookshelfOutput> {
  const request = new WantBookshelfRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)
  request.setBookId(input.bookId)

  return new Promise((resolve: (output: IBookshelfOutput) => void, reject: (reason: Error) => void) => {
    bookClient.wantBookshelf(request, metadata, (err: any, res: BookshelfResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IBookshelfOutput = setBookshelfOutput(res)
      resolve(output)
    })
  })
}

export function releaseBookshelf(req: Request<any>, input: IReleaseBookshelfInput): Promise<IBookshelfOutput> {
  const request = new ReleaseBookshelfRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)
  request.setBookId(input.bookId)

  return new Promise((resolve: (output: IBookshelfOutput) => void, reject: (reason: Error) => void) => {
    bookClient.releaseBookshelf(request, metadata, (err: any, res: BookshelfResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IBookshelfOutput = setBookshelfOutput(res)
      resolve(output)
    })
  })
}

export function deleteBookshelf(req: Request<any>, input: IDeleteBookshelfInput): Promise<void> {
  const request = new DeleteBookshelfRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)
  request.setBookId(input.bookId)

  return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
    bookClient.deleteBookshelf(request, metadata, (err: any, _: EmptyBook) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve()
    })
  })
}

function setBookOutput(res: BookResponse): IBookOutput {
  const authors = res.getAuthorsList().map((item: BookResponse.Author) => {
    const author: IBookOutputAuthor = {
      name: item.getName(),
      nameKana: item.getNameKana(),
    }

    return author
  })

  const output: IBookOutput = {
    id: res.getId(),
    title: res.getTitle(),
    titleKana: res.getTitleKana(),
    description: res.getDescription(),
    isbn: res.getIsbn(),
    publisher: res.getPublisher(),
    publishedOn: res.getPublishedOn(),
    thumbnailUrl: res.getThumbnailUrl(),
    rakutenUrl: res.getRakutenUrl(),
    rakutenSize: res.getRakutenSize(),
    rakutenGenreId: res.getRakutenGenreId(),
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
    authors,
  }

  return output
}

function setBookHashOutput(res: BookListResponse): IBookHashOutput {
  const output: IBookHashOutput = {}

  res.getBooksList().forEach((b: BookListResponse.Book) => {
    const authors = b.getAuthorsList().map(
      (a: BookListResponse.Author): IBookHashOutputAuthor => ({
        name: a.getName(),
        nameKana: a.getNameKana(),
      })
    )

    const book: IBookHashOutputBook = {
      id: b.getId(),
      title: b.getTitle(),
      titleKana: b.getTitleKana(),
      description: b.getDescription(),
      isbn: b.getIsbn(),
      publisher: b.getPublisher(),
      publishedOn: b.getPublishedOn(),
      thumbnailUrl: b.getThumbnailUrl(),
      rakutenUrl: b.getRakutenUrl(),
      rakutenSize: b.getRakutenSize(),
      rakutenGenreId: b.getRakutenGenreId(),
      createdAt: b.getCreatedAt(),
      updatedAt: b.getUpdatedAt(),
      authors,
    }

    output[b.getId()] = book
  })

  return output
}

function setBookshelfOutput(res: BookshelfResponse): IBookshelfOutput {
  const b: BookshelfResponse.Book = res.getBook() || ({} as BookshelfResponse.Book)

  const authors = b.getAuthorsList().map(
    (a: BookshelfResponse.Author): IBookshelfOutputAuthor => ({
      name: a.getName(),
      nameKana: a.getNameKana(),
    })
  )

  const book: IBookshelfOutputBook = {
    id: b.getId(),
    title: b.getTitle(),
    titleKana: b.getTitleKana(),
    description: b.getDescription(),
    isbn: b.getIsbn(),
    publisher: b.getPublisher(),
    publishedOn: b.getPublishedOn(),
    thumbnailUrl: b.getThumbnailUrl(),
    rakutenUrl: b.getRakutenUrl(),
    rakutenSize: b.getRakutenSize(),
    rakutenGenreId: b.getRakutenGenreId(),
    createdAt: b.getCreatedAt(),
    updatedAt: b.getUpdatedAt(),
    authors,
  }

  const output: IBookshelfOutput = {
    id: res.getId(),
    bookId: res.getBookId(),
    userId: res.getUserId(),
    reviewId: res.getReviewId(),
    status: res.getStatus(),
    readOn: res.getReadOn(),
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
    book,
  }

  const rv: BookshelfResponse.Review | undefined = res.getReview()
  if (rv) {
    const review: IBookshelfOutputReview = {
      score: rv.getScore(),
      impression: rv.getImpression(),
    }

    output.review = review
  }

  return output
}

function setBookshelfListOutput(res: BookshelfListResponse): IBookshelfListOutput {
  const bookshelves: Array<IBookshelfListOutputBookshelf> = res
    .getBookshelvesList()
    .map((bs: BookshelfListResponse.Bookshelf): IBookshelfListOutputBookshelf | undefined => {
      const b = bs.getBook()
      if (!b) {
        return
      }

      const authors = b.getAuthorsList().map(
        (a: BookshelfListResponse.Author): IBookshelfListOutputAuthor => ({
          name: a.getName(),
          nameKana: a.getNameKana(),
        })
      )

      const book: IBookshelfListOutputBook = {
        id: b.getId(),
        title: b.getTitle(),
        titleKana: b.getTitleKana(),
        description: b.getDescription(),
        isbn: b.getIsbn(),
        publisher: b.getPublisher(),
        publishedOn: b.getPublishedOn(),
        thumbnailUrl: b.getThumbnailUrl(),
        rakutenUrl: b.getRakutenUrl(),
        rakutenSize: b.getRakutenSize(),
        rakutenGenreId: b.getRakutenGenreId(),
        createdAt: b.getCreatedAt(),
        updatedAt: b.getUpdatedAt(),
        authors,
      }

      return {
        id: bs.getId(),
        bookId: bs.getBookId(),
        userId: bs.getUserId(),
        reviewId: bs.getReviewId(),
        status: bs.getStatus(),
        readOn: bs.getReadOn(),
        createdAt: bs.getCreatedAt(),
        updatedAt: bs.getUpdatedAt(),
        book,
      }
    })
    .filter((item): item is NonNullable<typeof item> => !!item)

  const output: IBookshelfListOutput = {
    bookshelves,
    limit: res.getLimit(),
    offset: res.getOffset(),
    total: res.getTotal(),
  }

  return output
}

function setReviewOutput(res: ReviewResponse): IReviewOutput {
  const output: IReviewOutput = {
    id: res.getId(),
    bookId: res.getBookId(),
    userId: res.getUserId(),
    score: res.getScore(),
    impression: res.getImpression(),
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
  }

  return output
}

function setReviewListOutput(res: ReviewListResponse): IReviewListOutput {
  const reviews = res.getReviewsList().map(
    (rv: ReviewListResponse.Review): IReviewListOutputReview => ({
      id: rv.getId(),
      bookId: rv.getBookId(),
      userId: rv.getUserId(),
      score: rv.getScore(),
      impression: rv.getImpression(),
      createdAt: rv.getCreatedAt(),
      updatedAt: rv.getUpdatedAt(),
    })
  )

  const output: IReviewListOutput = {
    reviews,
    limit: res.getLimit(),
    offset: res.getOffset(),
    total: res.getTotal(),
  }

  const order: ReviewListResponse.Order | undefined = res.getOrder()
  if (order) {
    const o: IReviewListOutputOrder = {
      by: order.getBy(),
      direction: order.getDirection(),
    }

    output.order = o
  }

  return output
}
