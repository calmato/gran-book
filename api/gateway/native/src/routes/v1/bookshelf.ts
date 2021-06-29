import express, { NextFunction, Request, Response } from 'express'
import {
  deleteBookshelf,
  getBookshelf,
  getUser,
  listBookshelf,
  listUserWithUserIds,
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
  IListUserByUserIdsInput,
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
  IBookshelfOutputReview,
  IUserHashOutput,
} from '~/types/output'
import { IReadBookshelfRequest } from '~/types/request'
import {
  IBookshelfListResponse,
  IBookshelfListResponseBook,
  IBookshelfListResponseBookshelf,
  IBookshelfResponse,
  IBookshelfResponseBookshelf,
  IBookshelfResponseReview,
  IBookshelfResponseUserOnReview,
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
      .then(async (bookshelfOutput: IBookshelfOutput) => {
        let userIds: string[] = []
        bookshelfOutput.book.reviews.forEach((rv: IBookshelfOutputReview) => {
          userIds = userIds.concat(rv.userId)
        })

        const userListInput: IListUserByUserIdsInput = {
          ids: userIds,
        }

        return listUserWithUserIds(req, userListInput)
          .then((usersOutput: IUserHashOutput) => {
            return setBookshelfResponse(bookshelfOutput, usersOutput)
          })
          .catch(() => {
            return setBookshelfResponse(bookshelfOutput, {})
          })
      })
      .then((response: IBookshelfResponse) => {
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
        const response: IBookshelfResponse = setBookshelfResponse(output, {})
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
        const response: IBookshelfResponse = setBookshelfResponse(output, {})
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
        const response: IBookshelfResponse = setBookshelfResponse(output, {})
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
        const response: IBookshelfResponse = setBookshelfResponse(output, {})
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
        const response: IBookshelfResponse = setBookshelfResponse(output, {})
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

function setBookshelfResponse(bookshelfOutput: IBookshelfOutput, usersOutput: IUserHashOutput): IBookshelfResponse {
  const authorNames: string[] = bookshelfOutput.book.authors.map((item: IBookshelfOutputAuthor) => {
    return item.name
  })

  const authorNameKanas: string[] = bookshelfOutput.book.authors.map((item: IBookshelfOutputAuthor) => {
    return item.nameKana
  })

  const reviews = bookshelfOutput.book.reviews.map(
    (item: IBookshelfOutputReview): IBookshelfResponseReview => {
      const user: IBookshelfResponseUserOnReview = {
        id: item.userId,
        username: '',
        thumbnailUrl: '',
      }

      if (usersOutput[item.userId]) {
        user.username = usersOutput[item.userId].username
        user.thumbnailUrl = usersOutput[item.userId].thumbnailUrl
      }

      return {
        id: item.id,
        impression: item.impression,
        createdAt: item.createdAt,
        updatedAt: item.updatedAt,
        user,
      }
    }
  )

  const bookshelf: IBookshelfResponseBookshelf = {
    id: bookshelfOutput.id,
    status: BookStatus[bookshelfOutput.status],
    readOn: bookshelfOutput.readOn,
    impression: bookshelfOutput.myReview?.impression || '',
    createdAt: bookshelfOutput.createdAt,
    updatedAt: bookshelfOutput.updatedAt,
  }

  const response: IBookshelfResponse = {
    id: bookshelfOutput.book.id,
    title: bookshelfOutput.book.title,
    titleKana: bookshelfOutput.book.titleKana,
    description: bookshelfOutput.book.description,
    isbn: bookshelfOutput.book.isbn,
    publisher: bookshelfOutput.book.publisher,
    publishedOn: bookshelfOutput.book.publishedOn,
    thumbnailUrl: bookshelfOutput.book.thumbnailUrl,
    rakutenUrl: bookshelfOutput.book.rakutenUrl,
    size: bookshelfOutput.book.rakutenSize,
    author: authorNames.join('/'),
    authorKana: authorNameKanas.join('/'),
    createdAt: bookshelfOutput.book.createdAt,
    updatedAt: bookshelfOutput.book.updatedAt,
    bookshelf,
    reviews,
    reviewLimit: bookshelfOutput.book.reviewLimit,
    reviewOffset: bookshelfOutput.book.reviewOffset,
    reviewTotal: bookshelfOutput.book.reviewTotal,
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
      impression: bs.myReview?.impression || '',
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
