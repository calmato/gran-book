import express, { NextFunction, Request, Response } from 'express'
import { getBook, getReview, getUser, listBookByBookIds, listUserReview } from '~/api'
import { GrpcError } from '~/types/exception'
import {
  IGetBookInput,
  IGetReviewInput,
  IGetUserInput,
  IListBookByBookIdsInput,
  IListUserReviewInput,
} from '~/types/input'
import {
  IBookHashOutput,
  IBookOutput,
  IReviewListOutput,
  IReviewListOutputReview,
  IReviewOutput,
  IUserOutput,
} from '~/types/output'
import {} from '~/types/request'
import {
  IReviewResponse,
  IReviewResponseBook,
  IReviewResponseUser,
  IUserReviewListResponse,
  IUserReviewListResponseBook,
} from '~/types/response'

const router = express.Router()

router.get(
  '/v1/users/:userId/reviews',
  async (req: Request, res: Response<IUserReviewListResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params
    const { limit, offset } = req.query as { [key: string]: string }

    const userInput: IGetUserInput = {
      id: userId,
    }

    await getUser(req, userInput)
      .then(async () => {
        const reviewListInput: IListUserReviewInput = {
          userId,
          limit: Number(limit) || 100,
          offset: Number(offset) || 0,
          by: '',
          direction: '',
        }

        return listUserReview(req, reviewListInput)
      })
      .then(async (reviewsOutput: IReviewListOutput) => {
        const bookIds: number[] = reviewsOutput.reviews.map((rv: IReviewListOutputReview) => {
          return rv.bookId
        })

        const bookListInput: IListBookByBookIdsInput = {
          bookIds: Array.from(new Set(bookIds)),
        }

        return listBookByBookIds(req, bookListInput)
          .then((booksOutput: IBookHashOutput) => {
            return setUserReviewListResponse(reviewsOutput, booksOutput)
          })
          .catch(() => {
            return setUserReviewListResponse(reviewsOutput, {})
          })
      })
      .then((response: IUserReviewListResponse) => {
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.get(
  '/v1/users/:userId/reviews/:reviewId',
  async (req: Request, res: Response<IReviewResponse>, next: NextFunction): Promise<void> => {
    const { userId, reviewId } = req.params

    const userInput: IGetUserInput = {
      id: userId,
    }

    await getUser(req, userInput)
      .then(async (userOutput: IUserOutput) => {
        const reviewInput: IGetReviewInput = {
          reviewId: Number(reviewId) || 0,
        }

        await getReview(req, reviewInput)
          .then(async (reviewOutput: IReviewOutput) => {
            const bookInput: IGetBookInput = {
              bookId: reviewOutput.bookId,
            }

            await getBook(req, bookInput)
              .then((bookOutput: IBookOutput) => {
                const response: IReviewResponse = setReviewResponse(reviewOutput, bookOutput, userOutput)
                res.status(200).json(response)
              })
              .catch((err: GrpcError) => {
                throw err
              })
          })
          .catch((err: GrpcError) => {
            throw err
          })
      })
      .catch((err: GrpcError) => next(err))
  }
)

function setReviewResponse(
  reviewOutput: IReviewOutput,
  bookOutput: IBookOutput,
  userOutput: IUserOutput
): IReviewResponse {
  const book: IReviewResponseBook = {
    id: bookOutput.id,
    title: bookOutput.title,
    thumbnailUrl: bookOutput.thumbnailUrl,
  }

  const user: IReviewResponseUser = {
    id: userOutput.id,
    username: userOutput.username,
    thumbnailUrl: userOutput.thumbnailUrl,
  }

  const response: IReviewResponse = {
    id: reviewOutput.id,
    impression: reviewOutput.impression,
    createdAt: reviewOutput.createdAt,
    updatedAt: reviewOutput.updatedAt,
    book,
    user,
  }

  return response
}

function setUserReviewListResponse(
  reviewOutput: IReviewListOutput,
  booksOutput: IBookHashOutput
): IUserReviewListResponse {
  const reviews = reviewOutput.reviews.map((rv: IReviewListOutputReview) => {
    const book: IUserReviewListResponseBook = {
      id: 0,
      title: '',
      thumbnailUrl: '',
    }

    if (booksOutput[rv.bookId]) {
      const { id, title, thumbnailUrl } = booksOutput[rv.bookId]

      book.id = id
      book.title = title
      book.thumbnailUrl = thumbnailUrl
    }

    return {
      id: rv.id,
      impression: rv.impression,
      createdAt: rv.createdAt,
      updatedAt: rv.updatedAt,
      book,
    }
  })

  const response: IUserReviewListResponse = {
    reviews,
    limit: reviewOutput.limit,
    offset: reviewOutput.offset,
    total: reviewOutput.total,
  }

  if (reviewOutput.order) {
    response.order = {
      by: reviewOutput.order.by,
      direction: reviewOutput.order.direction,
    }
  }

  return response
}

export default router
