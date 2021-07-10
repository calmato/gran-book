import express, { NextFunction, Request, Response } from 'express'
import { getBookshelf, getUser, listBookReview, listBookshelf, listUserWithUserIds } from '~/api'
import { IGetBookshelfInput, IGetUserInput, IListBookReviewInput, IListBookshelfInput, IListUserByUserIdsInput } from '~/types/input'
import { IBookshelfListOutput, IBookshelfListOutputAuthor, IBookshelfListOutputBookshelf, IBookshelfOutput, IBookshelfOutputAuthor, IReviewListOutput, IReviewListOutputReview, IUserHashOutput } from '~/types/output'
import { IBookshelfListV2Response, IBookshelfListV2ResponseBook, IBookshelfListV2ResponseBookshelf, IBookshelfV2Response, IBookshelfV2ResponseBookshelf, IBookshelfV2ResponseReview, IBookshelfV2ResponseUser } from '~/types/response'
import { BookStatus, LIST_DEFAULT_LIMIT, LIST_DEFAULT_OFFSET, LIST_DEFAULT_ORDER_BY, LIST_DEFAULT_ORDER_DIRECTION } from '~/util'

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
                return setBookshelfResponse(bookshelfOutput, reviewsOutput, usersOutput)
              })
              .catch(() => {
                return setBookshelfResponse(bookshelfOutput, reviewsOutput, {})
              })
          })
          .catch((err: Error) => {
            throw err
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

function setBookshelfResponse(bookshelfOutput: IBookshelfOutput, reviewsOutput: any, usersOutput: IUserHashOutput): IBookshelfV2Response {
  const authorNames = bookshelfOutput.book.authors.map((item: IBookshelfOutputAuthor): string => {
    return item.name
  })

  const authorNameKanas = bookshelfOutput.book.authors.map((item: IBookshelfOutputAuthor): string => {
    return item.nameKana
  })

  const bookshelf: IBookshelfV2ResponseBookshelf = {
    status: BookStatus[bookshelfOutput.status],
    readOn: bookshelfOutput.readOn,
    reviewId: bookshelfOutput.reviewId,
    createdAt: bookshelfOutput.createdAt,
    updatedAt: bookshelfOutput.updatedAt,
  }

  // TODO: refactor
  const reviews = reviewsOutput.map((item: any): IBookshelfV2ResponseReview => {
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
  })

  const response: IBookshelfV2Response = {
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
    reviewLimit: reviewsOutput.limit,
    reviewOffset: reviewsOutput.offset,
    reviewTotal: reviewsOutput.total,
  }

  return response
}

function setBookshelfListResponse(output: IBookshelfListOutput): IBookshelfListV2Response {
  const books = output.bookshelves.map((bs: IBookshelfListOutputBookshelf): IBookshelfListV2ResponseBook => {
    const authorNames = bs.book.authors.map((item: IBookshelfListOutputAuthor): string => {
      return item.name
    })

    const authorNameKanas = bs.book.authors.map((item: IBookshelfListOutputAuthor): string => {
      return item.nameKana
    })

    const bookshelf: IBookshelfListV2ResponseBookshelf = {
      status: BookStatus[bs.status],
      readOn: bs.readOn,
      reviewId: bs.reviewId,
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
  })

  const response: IBookshelfListV2Response = {
    books,
    limit: output.limit,
    offset: output.offset,
    total: output.total,
  }

  return response
}

export default router
