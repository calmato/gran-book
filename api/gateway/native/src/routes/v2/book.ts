import express, { NextFunction, Request, Response } from 'express'
import { getBook, getBookByIsbn, listBookReview, listUserWithUserIds } from '~/api'
import { IGetBookByIsbnInput, IGetBookInput, IListBookReviewInput, IListUserByUserIdsInput } from '~/types/input'
import {
  IBookOutput,
  IBookOutputAuthor,
  IReviewListOutput,
  IReviewListOutputReview,
  IUserHashOutput,
} from '~/types/output'
import { IBookV2Response, IBookV2ResponseReview, IBookV2ResponseUser } from '~/types/response'
import { LIST_DEFAULT_LIMIT, LIST_DEFAULT_OFFSET, LIST_DEFAULT_ORDER_BY, LIST_DEFAULT_ORDER_DIRECTION } from '~/util'

const router = express.Router()

router.get(
  '/v2/books/:bookId',
  async (req: Request, res: Response<IBookV2Response>, next: NextFunction): Promise<void> => {
    const { bookId } = req.params
    const { key } = req.query as { [key: string]: string }

    await bookByIdOrIsbn(req, key, bookId)
      .then(async (bookOutput: IBookOutput) => {
        const reviewListInput: IListBookReviewInput = {
          bookId: bookOutput.id,
          limit: LIST_DEFAULT_LIMIT,
          offset: LIST_DEFAULT_OFFSET,
          by: LIST_DEFAULT_ORDER_BY,
          direction: LIST_DEFAULT_ORDER_DIRECTION,
        }

        const reviewsOutput = await listBookReview(req, reviewListInput)

        const userIds = reviewsOutput.reviews.map((rv: IReviewListOutputReview): string => {
          return rv.userId
        })
        const userListInput: IListUserByUserIdsInput = {
          ids: Array.from(new Set(userIds)),
        }

        const usersOutput = await listUserWithUserIds(req, userListInput).catch(() => {
          return {}
        })

        return setBookResponse(bookOutput, reviewsOutput, usersOutput)
      })
      .then((response: IBookV2Response) => {
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

async function bookByIdOrIsbn(req: Request, key: string, value: string): Promise<IBookOutput> {
  switch (key) {
    case 'isbn': {
      const isbn: string = value
      return bookByIsbn(req, isbn)
    }
    case 'id':
    default: {
      const id: number = Number(value) || 0
      return bookById(req, id)
    }
  }
}

async function bookById(req: Request, bookId: number): Promise<IBookOutput> {
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

function setBookResponse(
  bookOutput: IBookOutput,
  reviewsOutput: IReviewListOutput,
  usersOutput: IUserHashOutput
): IBookV2Response {
  const authorNames = bookOutput.authors.map((item: IBookOutputAuthor): string => {
    return item.name
  })

  const authorNameKanas = bookOutput.authors.map((item: IBookOutputAuthor): string => {
    return item.nameKana
  })

  const reviews = reviewsOutput.reviews.map(
    (item: IReviewListOutputReview): IBookV2ResponseReview => {
      const user: IBookV2ResponseUser = {
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

  const response: IBookV2Response = {
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
    reviews,
    reviewLimit: reviewsOutput.limit,
    reviewOffset: reviewsOutput.offset,
    reviewTotal: reviewsOutput.total,
  }

  return response
}

export default router
