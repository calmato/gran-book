import express, { NextFunction, Request, Response } from 'express'
import {
  deleteBookshelf,
  getBookshelf,
  getUser,
  listBookshelf,
  readBookshelf,
  readingBookshelf,
  releaseBookshelf,
  stackBookshelf,
  wantBookshelf,
} from '~/api'
import { BookStatus } from '~/types/book'
import { GrpcError } from '~/types/exception'
import {
  IDeleteBookshelfInput,
  IGetBookshelfInput,
  IGetUserInput,
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
  IBookshelfListResponseBookshelf,
  IBookshelfResponse,
  IBookshelfResponseBookshelf,
} from '~/types/response'

const router = express.Router()

router.get(
  '/v1/users/:userId/books',
  async (req: Request, res: Response<IBookshelfListResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params
    const { limit, offset } = req.query as { [key: string]: string }

    const userInput: IGetUserInput = {
      id: userId,
    }

    await getUser(req, userInput)
      .then(async () => {
        const bookshelfInput: IListBookshelfInput = {
          userId,
          limit: Number(limit) || 100,
          offset: Number(offset) || 0,
        }

        return listBookshelf(req, bookshelfInput)
      })
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

    const userInput: IGetUserInput = {
      id: userId,
    }

    await getUser(req, userInput)
      .then(async () => {
        const bookshelfInput: IGetBookshelfInput = {
          userId,
          bookId: Number(bookId) || 0,
        }

        return getBookshelf(req, bookshelfInput)
      })
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
    const { userId, bookId } = req.params
    const { readOn, impression } = req.body as IReadBookshelfRequest

    const userInput: IGetUserInput = {
      id: userId,
    }

    await getUser(req, userInput)
      .then(async () => {
        const bookshelfInput: IReadBookshelfInput = {
          userId,
          bookId: Number(bookId) || 0,
          impression,
          readOn,
        }

        return readBookshelf(req, bookshelfInput)
      })
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
    const { userId, bookId } = req.params

    const userInput: IGetUserInput = {
      id: userId,
    }

    await getUser(req, userInput)
      .then(async () => {
        const bookshelfInput: IReadingBookshelfInput = {
          userId,
          bookId: Number(bookId) || 0,
        }

        return readingBookshelf(req, bookshelfInput)
      })
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
    const { userId, bookId } = req.params

    const userInput: IGetUserInput = {
      id: userId,
    }

    await getUser(req, userInput)
      .then(async () => {
        const bookshelfInput: IStackBookshelfInput = {
          userId,
          bookId: Number(bookId) || 0,
        }

        return stackBookshelf(req, bookshelfInput)
      })
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
    const { userId, bookId } = req.params

    const userInput: IGetUserInput = {
      id: userId,
    }

    await getUser(req, userInput)
      .then(async () => {
        const bookshelfInput: IWantBookshelfInput = {
          userId,
          bookId: Number(bookId) || 0,
        }

        return wantBookshelf(req, bookshelfInput)
      })
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
    const { userId, bookId } = req.params

    const userInput: IGetUserInput = {
      id: userId,
    }

    await getUser(req, userInput)
      .then(async () => {
        const bookshelfInput: IReleaseBookshelfInput = {
          userId,
          bookId: Number(bookId) || 0,
        }

        return releaseBookshelf(req, bookshelfInput)
      })
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
    const { userId, bookId } = req.params

    const userInput: IGetUserInput = {
      id: userId,
    }

    await getUser(req, userInput)
      .then(async () => {
        const bookshelfInput: IDeleteBookshelfInput = {
          userId,
          bookId: Number(bookId) || 0,
        }

        return deleteBookshelf(req, bookshelfInput)
      })
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

  const bookshelf: IBookshelfResponseBookshelf = {
    id: output.id,
    status: BookStatus[output.status],
    impression: output.review?.impression || '',
    readOn: output.readOn,
    createdAt: output.createdAt,
    updatedAt: output.updatedAt,
  }

  const response: IBookshelfResponse = {
    id: output.book.id,
    title: output.book.title,
    titleKana: output.book.titleKana,
    description: output.book.description,
    isbn: output.book.isbn,
    publisher: output.book.publisher,
    publishedOn: output.book.publishedOn,
    thumbnailUrl: output.book.thumbnailUrl,
    rakutenUrl: output.book.rakutenUrl,
    size: output.book.rakutenSize,
    author: authorNames.join('/'),
    authorKana: authorNameKanas.join('/'),
    createdAt: output.book.createdAt,
    updatedAt: output.book.updatedAt,
    bookshelf,
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

    const bookshelf: IBookshelfListResponseBookshelf = {
      id: bs.id,
      status: BookStatus[bs.status],
      readOn: bs.readOn,
      createdAt: bs.createdAt,
      updatedAt: bs.updatedAt,
    }

    const book: IBookshelfListResponseBook = {
      id: bs.book.id,
      title: bs.book.title,
      titleKana: bs.book.titleKana,
      description: bs.book.description,
      isbn: bs.book.isbn,
      publisher: bs.book.publisher,
      publishedOn: bs.book.publishedOn,
      thumbnailUrl: bs.book.thumbnailUrl,
      rakutenUrl: bs.book.rakutenUrl,
      size: bs.book.rakutenSize,
      author: authorNames.join('/'),
      authorKana: authorNameKanas.join('/'),
      createdAt: bs.book.createdAt,
      updatedAt: bs.book.updatedAt,
      bookshelf,
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
