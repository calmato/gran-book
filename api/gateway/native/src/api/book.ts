import { Request } from 'express'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import { bookClient } from '~/plugins/grpc'
import {
  BookResponse,
  BookshelfResponse,
  CreateBookRequest,
  DeleteBookshelfRequest,
  EmptyBook,
  ReadBookshelfRequest,
  ReadingBookshelfRequest,
  ReleaseBookshelfRequest,
  ShowBookRequest,
  StackBookshelfRequest,
  UpdateBookRequest,
  WantBookshelfRequest,
} from '~/proto/book_apiv1_pb'
import {
  IBookInputAuthor,
  ICreateBookInput,
  IDeleteBookshelfInput,
  IReadBookshelfInput,
  IReadingBookshelfInput,
  IReleaseBookshelfInput,
  IShowBookInput,
  IStackBookshelfInput,
  IUpdateBookInput,
  IWantBookshelfInput,
} from '~/types/input'
import {
  IBookOutput,
  IBookOutputAuthor,
  IBookOutputBookshelf,
  IBookOutputReview,
  IBookshelfOutput,
} from '~/types/output'

export function showBook(req: Request<any>, input: IShowBookInput): Promise<IBookOutput> {
  const request = new ShowBookRequest()
  const metadata = getGrpcMetadata(req)

  request.setIsbn(input.isbn)

  return new Promise((resolve: (output: IBookOutput) => void, reject: (reason: Error) => void) => {
    bookClient.showBook(request, metadata, (err: any, res: BookResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IBookOutput = setBookOutput(res)
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
  const bookshelf: IBookOutputBookshelf = {
    id: res.getBookshelf()?.getId() || 0,
    status: res.getBookshelf()?.getStatus() || 0,
    readOn: res.getBookshelf()?.getReadOn() || '',
    createdAt: res.getBookshelf()?.getCreatedAt() || '',
    updatedAt: res.getBookshelf()?.getUpdatedAt() || '',
  }

  const authors = res.getAuthorsList().map((item: BookResponse.Author) => {
    const author: IBookOutputAuthor = {
      name: item.getName(),
      nameKana: item.getNameKana(),
    }

    return author
  })

  const reviews = res.getReviewsList().map((item: BookResponse.Review) => {
    const reviews: IBookOutputReview = {
      id: item.getId(),
      userId: item.getUserId(),
      score: item.getScore(),
      impression: item.getImpression(),
      createdAt: item.getCreatedAt(),
      updatedAt: item.getUpdatedAt(),
    }

    return reviews
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
    rakutenGenreId: res.getRakutenGenreId(),
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
    bookshelf,
    authors,
    reviews,
  }

  return output
}

function setBookshelfOutput(res: BookshelfResponse): IBookshelfOutput {
  const output: IBookshelfOutput = {
    id: res.getId(),
    bookId: res.getBookId(),
    userId: res.getUserId(),
    status: res.getStatus(),
    impression: res.getImpression(),
    readOn: res.getReadOn(),
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
  }

  return output
}
