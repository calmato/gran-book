import express, { NextFunction, Request, Response } from 'express'
import {
  getBook,
  getBookByIsbn,
  getBookshelf,
  getUser,
  listBookReview,
  listBookshelf,
  listUserWithUserIds,
} from '~/api'
import {
  IGetBookByIsbnInput,
  IGetBookInput,
  IGetBookshelfInput,
  IGetUserInput,
  IListBookReviewInput,
  IListBookshelfInput,
  IListUserByUserIdsInput,
} from '~/types/input'
import {
  IBookOutput,
  IBookshelfListOutput,
  IBookshelfListOutputAuthor,
  IBookshelfListOutputBookshelf,
  IBookshelfOutput,
  IBookshelfOutputAuthor,
  IReviewListOutput,
  IReviewListOutputReview,
  IUserHashOutput,
} from '~/types/output'
import {
  IBookshelfListV2Response,
  IBookshelfListV2ResponseBook,
  IBookshelfListV2ResponseBookshelf,
  IBookshelfV2Response,
  IBookshelfV2ResponseBookshelf,
  IBookshelfV2ResponseReview,
  IBookshelfV2ResponseUser,
} from '~/types/response'
import {
  BookStatus,
  LIST_DEFAULT_LIMIT,
  LIST_DEFAULT_OFFSET,
  LIST_DEFAULT_ORDER_BY,
  LIST_DEFAULT_ORDER_DIRECTION,
} from '~/util'

const router = express.Router()

router.get(
  '/v2/users/:userId/books',
  async (req: Request, res: Response<IBookshelfListV2Response>, next: NextFunction): Promise<void> => {
    const { userId } = req.params
    const { limit, offset } = req.query as { [key: string]: string }

    const userInput: IGetUserInput = {
      id: userId,
    }

    await getUser(req, userInput)
      .then(async () => {
        const bookshelfInput: IListBookshelfInput = {
          userId,
          limit: Number(limit) || LIST_DEFAULT_LIMIT,
          offset: Number(offset) || LIST_DEFAULT_OFFSET,
        }

        return listBookshelf(req, bookshelfInput)
      })
      .then((output: IBookshelfListOutput) => {
        const response: IBookshelfListV2Response = setBookshelfListResponse(output)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

router.get(
  '/v2/users/:userId/books/:bookId',
  async (req: Request, res: Response<IBookshelfV2Response>, next: NextFunction): Promise<void> => {
    const { userId, bookId } = req.params
    const { key } = req.query as { [key: string]: string }

    const userInput: IGetUserInput = {
      id: userId,
    }

    await getUser(req, userInput)
      .then(() => {
        switch (key) {
          case 'isbn': {
            const isbn: string = bookId
            return bookByIsbn(req, isbn)
          }
          case 'id':
          default: {
            const id: number = Number(bookId) || 0
            return bookById(req, id)
          }
        }
      })
      .then(async (bookOutput: IBookOutput) => {
        const bookshelfInput: IGetBookshelfInput = {
          userId,
          bookId: bookOutput.id,
        }

        return getBookshelf(req, bookshelfInput)
          .then(async (bookshelfOutput: IBookshelfOutput) => {
            const reviewListInput: IListBookReviewInput = {
              bookId: Number(bookId) || 0,
              limit: LIST_DEFAULT_LIMIT,
              offset: LIST_DEFAULT_OFFSET,
              by: LIST_DEFAULT_ORDER_BY,
              direction: LIST_DEFAULT_ORDER_DIRECTION,
            }

            return listBookReview(req, reviewListInput)
              .then(async (reviewsOutput: IReviewListOutput) => {
                const userIds = reviewsOutput.reviews.map((rv: IReviewListOutputReview): string => {
                  return rv.userId
                })

                const userListInput: IListUserByUserIdsInput = {
                  ids: Array.from(new Set(userIds)),
                }

                return listUserWithUserIds(req, userListInput)
                  .then((usersOutput: IUserHashOutput) => {
                    return setBookshelfResponse(bookOutput, usersOutput, bookshelfOutput, reviewsOutput)
                  })
                  .catch(() => {
                    return setBookshelfResponse(bookOutput, {}, bookshelfOutput, reviewsOutput)
                  })
              })
              .catch((err: Error) => {
                throw err
              })
          })
          .catch(() => {
            return setBookshelfResponse(bookOutput, {})
          })
      })
      .then((response: IBookshelfV2Response) => {
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

function bookById(req: Request, bookId: number): Promise<IBookOutput> {
  const input: IGetBookInput = {
    bookId,
  }

  return getBook(req, input)
}

function bookByIsbn(req: Request, isbn: string): Promise<IBookOutput> {
  const input: IGetBookByIsbnInput = {
    isbn,
  }

  return getBookByIsbn(req, input)
}

function setBookshelfResponse(
  bookOutput: IBookOutput,
  usersOutput: IUserHashOutput,
  bookshelfOutput?: IBookshelfOutput,
  reviewsOutput?: IReviewListOutput
): IBookshelfV2Response {
  const authorNames = bookOutput.authors.map((item: IBookshelfOutputAuthor): string => {
    return item.name
  })

  const authorNameKanas = bookOutput.authors.map((item: IBookshelfOutputAuthor): string => {
    return item.nameKana
  })

  const reviews = reviewsOutput?.reviews.map(
    (item: IReviewListOutputReview): IBookshelfV2ResponseReview => {
      const user: IBookshelfV2ResponseUser = {
        id: '',
        username: 'unknown',
        thumbnailUrl: '',
      }

      if (usersOutput[item.userId]) {
        const { id, username, thumbnailUrl } = usersOutput[item.userId]

        user.id = id
        user.username = username
        user.thumbnailUrl = thumbnailUrl
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

  const response: IBookshelfV2Response = {
    id: bookOutput.id,
    title: bookOutput.title,
    titleKana: bookOutput.titleKana,
    description: bookOutput.description,
    isbn: bookOutput.isbn,
    publisher: bookOutput.publisher,
    publishedOn: bookOutput.publishedOn,
    thumbnailUrl: bookOutput.thumbnailUrl,
    rakutenUrl: bookOutput.rakutenUrl,
    size: bookOutput.rakutenSize,
    author: authorNames.join('/'),
    authorKana: authorNameKanas.join('/'),
    createdAt: bookOutput.createdAt,
    updatedAt: bookOutput.updatedAt,
    reviews: reviews || [],
    reviewLimit: reviewsOutput?.limit || 0,
    reviewOffset: reviewsOutput?.offset || 0,
    reviewTotal: reviewsOutput?.total || 0,
  }

  if (bookshelfOutput && bookshelfOutput.status !== BookStatus.none) {
    const bookshelf: IBookshelfV2ResponseBookshelf = {
      status: BookStatus[bookshelfOutput.status],
      readOn: bookshelfOutput.readOn,
      reviewId: bookshelfOutput.reviewId === 0 ? bookshelfOutput.reviewId : undefined,
      createdAt: bookshelfOutput.createdAt,
      updatedAt: bookshelfOutput.updatedAt,
    }

    response.bookshelf = bookshelf
  }

  return response
}

function setBookshelfListResponse(output: IBookshelfListOutput): IBookshelfListV2Response {
  const books = output.bookshelves.map(
    (bs: IBookshelfListOutputBookshelf): IBookshelfListV2ResponseBook => {
      const authorNames = bs.book.authors.map((item: IBookshelfListOutputAuthor): string => {
        return item.name
      })

      const authorNameKanas = bs.book.authors.map((item: IBookshelfListOutputAuthor): string => {
        return item.nameKana
      })

      const bookshelf: IBookshelfListV2ResponseBookshelf = {
        status: BookStatus[bs.status],
        readOn: bs.readOn,
        reviewId: bs.reviewId === 0 ? bs.reviewId : undefined,
        createdAt: bs.createdAt,
        updatedAt: bs.updatedAt,
      }

      const book: IBookshelfListV2ResponseBook = {
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
    }
  )

  const response: IBookshelfListV2Response = {
    books,
    limit: output.limit,
    offset: output.offset,
    total: output.total,
  }

  return response
}

export default router
