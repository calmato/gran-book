import express, { NextFunction, Request, Response } from 'express'
import {
  deleteBookshelf,
  getBookshelf,
  listBookshelf,
  readBookshelf,
  readingBookshelf,
  releaseBookshelf,
  stackBookshelf,
  wantBookshelf,
} from '~/api'
import { GrpcError } from '~/types/exception'
import {
  IDeleteBookshelfInput,
  IGetBookshelfInput,
  IListBookshelfInput,
  IReadBookshelfInput,
  IReadingBookshelfInput,
  IReleaseBookshelfInput,
  IStackBookshelfInput,
  IWantBookshelfInput,
} from '~/types/input'
import {
  IBookshelfListOutput,
  IBookshelfListOutputAuthor,
  IBookshelfListOutputBookshelf,
  IBookshelfOutput,
  IBookshelfOutputAuthor,
} from '~/types/output'
import { IReadBookshelfRequest } from '~/types/request'
import {
  IBookshelfListResponse,
  IBookshelfListResponseBook,
  IBookshelfListResponseDetail,
  IBookshelfResponse,
  IBookshelfResponseDetail,
} from '~/types/response'

const router = express.Router()

router.get(
  '/v1/users/:userId/books',
  async (req: Request, res: Response<IBookshelfListResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params
    const { limit, offset } = req.query as { [key: string]: string }

    const input: IListBookshelfInput = {
      userId,
      limit: Number(limit) || 100,
      offset: Number(offset) || 0,
    }

    await listBookshelf(req, input)
      .then((output: IBookshelfListOutput) => {
        const response: IBookshelfListResponse = setBookshelfListResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.get(
  '/v1/users/:userId/books/:bookId',
  async (req: Request, res: Response<IBookshelfResponse>, next: NextFunction): Promise<void> => {
    const { userId, bookId } = req.params

    const input: IGetBookshelfInput = {
      userId,
      bookId: Number(bookId) || 0,
    }

    await getBookshelf(req, input)
      .then((output: IBookshelfOutput) => {
        const response: IBookshelfResponse = setBookshelfResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.post(
  '/v1/users/:userId/books/:bookId/read',
  async (req: Request, res: Response<IBookshelfResponse>, next: NextFunction): Promise<void> => {
    const { bookId } = req.params
    const { readOn, impression } = req.body as IReadBookshelfRequest

    const input: IReadBookshelfInput = {
      bookId: Number(bookId) || 0,
      impression,
      readOn,
    }

    await readBookshelf(req, input)
      .then((output: IBookshelfOutput) => {
        const response: IBookshelfResponse = setBookshelfResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.post(
  '/v1/users/:userId/books/:bookId/reading',
  async (req: Request, res: Response<IBookshelfResponse>, next: NextFunction): Promise<void> => {
    const { bookId } = req.params

    const input: IReadingBookshelfInput = {
      bookId: Number(bookId) || 0,
    }

    await readingBookshelf(req, input)
      .then((output: IBookshelfOutput) => {
        const response: IBookshelfResponse = setBookshelfResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.post(
  '/v1/users/:userId/books/:bookId/stack',
  async (req: Request, res: Response<IBookshelfResponse>, next: NextFunction): Promise<void> => {
    const { bookId } = req.params

    const input: IStackBookshelfInput = {
      bookId: Number(bookId) || 0,
    }

    await stackBookshelf(req, input)
      .then((output: IBookshelfOutput) => {
        const response: IBookshelfResponse = setBookshelfResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.post(
  '/v1/users/:userId/books/:bookId/want',
  async (req: Request, res: Response<IBookshelfResponse>, next: NextFunction): Promise<void> => {
    const { bookId } = req.params

    const input: IWantBookshelfInput = {
      bookId: Number(bookId) || 0,
    }

    await wantBookshelf(req, input)
      .then((output: IBookshelfOutput) => {
        const response: IBookshelfResponse = setBookshelfResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.post(
  '/v1/users/:userId/books/:bookId/release',
  async (req: Request, res: Response<IBookshelfResponse>, next: NextFunction): Promise<void> => {
    const { bookId } = req.params

    const input: IReleaseBookshelfInput = {
      bookId: Number(bookId) || 0,
    }

    await releaseBookshelf(req, input)
      .then((output: IBookshelfOutput) => {
        const response: IBookshelfResponse = setBookshelfResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.delete(
  '/v1/users/:userId/books/:bookId',
  async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    const { bookId } = req.params

    const input: IDeleteBookshelfInput = {
      bookId: Number(bookId) || 0,
    }

    await deleteBookshelf(req, input)
      .then(() => {
        res.status(200).json({ status: 'ok' })
      })
      .catch((err: GrpcError) => next(err))
  }
)

function setBookshelfResponse(output: IBookshelfOutput): IBookshelfResponse {
  const authorNames: string[] = output.book.authors.map((item: IBookshelfOutputAuthor) => {
    return item.name
  })

  const authorNameKanas: string[] = output.book.authors.map((item: IBookshelfOutputAuthor) => {
    return item.nameKana
  })

  const detail: IBookshelfResponseDetail = {
    id: output.book.id,
    title: output.book.title,
    titleKana: output.book.titleKana,
    description: output.book.description,
    isbn: output.book.isbn,
    publisher: output.book.publisher,
    publishedOn: output.book.publishedOn,
    thumbnailUrl: output.book.thumbnailUrl,
    rakutenUrl: output.book.rakutenUrl,
    rakutenGenreId: output.book.rakutenGenreId,
    author: authorNames.join('/'),
    authorKana: authorNameKanas.join('/'),
    createdAt: output.book.createdAt,
    updatedAt: output.book.updatedAt,
  }

  const response: IBookshelfResponse = {
    id: output.id,
    userId: output.userId,
    status: output.status,
    impression: '',
    readOn: output.readOn,
    createdAt: output.createdAt,
    updatedAt: output.updatedAt,
    detail,
  }

  if (output.review) {
    response.impression = output.review.impression
  }

  return response
}

function setBookshelfListResponse(output: IBookshelfListOutput): IBookshelfListResponse {
  const books: IBookshelfListResponseBook[] = output.bookshelves.map((bs: IBookshelfListOutputBookshelf) => {
    const authorNames: string[] = bs.book.authors.map((item: IBookshelfListOutputAuthor) => {
      return item.name
    })

    const authorNameKanas: string[] = bs.book.authors.map((item: IBookshelfListOutputAuthor) => {
      return item.nameKana
    })

    const detail: IBookshelfListResponseDetail = {
      id: bs.book.id,
      title: bs.book.title,
      titleKana: bs.book.titleKana,
      description: bs.book.description,
      isbn: bs.book.isbn,
      publisher: bs.book.publisher,
      publishedOn: bs.book.publishedOn,
      thumbnailUrl: bs.book.thumbnailUrl,
      rakutenUrl: bs.book.rakutenUrl,
      rakutenGenreId: bs.book.rakutenGenreId,
      author: authorNames.join('/'),
      authorKana: authorNameKanas.join('/'),
      createdAt: bs.book.createdAt,
      updatedAt: bs.book.updatedAt,
    }

    const book: IBookshelfListResponseBook = {
      id: bs.id,
      status: bs.status,
      readOn: bs.readOn,
      createdAt: bs.createdAt,
      updatedAt: bs.updatedAt,
      detail,
    }

    return book
  })

  const response: IBookshelfListResponse = {
    books,
    limit: output.limit,
    offset: output.offset,
    total: output.total,
  }

  return response
}

export default router
