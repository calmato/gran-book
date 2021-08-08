import express, { NextFunction, Request, Response } from 'express'
import {
  createBook,
  getBook,
  getBookByIsbn,
  getBookshelf,
  getReview,
  getUser,
  listBookReview,
  multiGetUser,
  updateBook,
} from '~/api'
import { GrpcError } from '~/types/exception'
import {
  IBookInputAuthor,
  ICreateBookInput,
  IGetBookByIsbnInput,
  IGetBookInput,
  IGetBookshelfInput,
  IGetReviewInput,
  IGetUserInput,
  IListBookReviewInput,
  IMultiGetUserInput,
  IUpdateBookInput,
} from '~/types/input'
import {
  IBookOutput,
  IBookOutputAuthor,
  IBookshelfOutput,
  IReviewListOutput,
  IReviewListOutputReview,
  IReviewOutput,
  IUserMapOutput,
  IUserOutput,
} from '~/types/output'
import { ICreateBookRequest, IUpdateBookRequest } from '~/types/request'
import {
  IBookResponse,
  IBookResponseBookshelf,
  IBookReviewListResponse,
  IBookReviewListResponseUser,
  IReviewResponse,
  IReviewResponseBook,
  IReviewResponseUser,
} from '~/types/response'
import { BookStatus } from '~/util'

const router = express.Router()

router.post(
  '/v1/books',
  async (req: Request, res: Response<IBookResponse>, next: NextFunction): Promise<void> => {
    const {
      title,
      titleKana,
      itemCaption,
      isbn,
      publisherName,
      salesDate,
      largeImageUrl,
      mediumImageUrl,
      smallImageUrl,
      itemUrl,
      booksGenreId,
      size,
      author,
      authorKana,
    } = req.body as ICreateBookRequest

    const authorNames: string[] = author.split('/')
    const authorNameKanas: string[] = authorKana.split('/')

    const authors: IBookInputAuthor[] = authorNames.map((val: string, i: number) => {
      const item: IBookInputAuthor = {
        name: val,
        nameKana: authorNameKanas[i],
      }

      return item
    })

    const input: ICreateBookInput = {
      title,
      titleKana,
      description: itemCaption,
      isbn,
      publisher: publisherName,
      publishedOn: salesDate,
      thumbnailUrl: largeImageUrl || mediumImageUrl || smallImageUrl,
      rakutenUrl: itemUrl,
      rakutenSize: size,
      rakutenGenreId: booksGenreId,
      authors,
    }

    await createBook(req, input)
      .then((output: IBookOutput) => {
        const response: IBookResponse = setBookResponse(output)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

router.patch(
  '/v1/books',
  async (req: Request, res: Response<IBookResponse>, next: NextFunction): Promise<void> => {
    const {
      title,
      titleKana,
      itemCaption,
      isbn,
      publisherName,
      salesDate,
      largeImageUrl,
      mediumImageUrl,
      smallImageUrl,
      itemUrl,
      booksGenreId,
      size,
      author,
      authorKana,
    } = req.body as IUpdateBookRequest

    const authorNames: string[] = author.split('/')
    const authorNameKanas: string[] = authorKana.split('/')

    const authors: IBookInputAuthor[] = authorNames.map((val: string, i: number) => {
      const item: IBookInputAuthor = {
        name: val,
        nameKana: authorNameKanas[i],
      }

      return item
    })

    const input: IUpdateBookInput = {
      title,
      titleKana,
      description: itemCaption,
      isbn,
      publisher: publisherName,
      publishedOn: salesDate,
      thumbnailUrl: largeImageUrl || mediumImageUrl || smallImageUrl,
      rakutenUrl: itemUrl,
      rakutenSize: size,
      rakutenGenreId: booksGenreId,
      authors,
    }

    await updateBook(req, input)
      .then((output: IBookOutput) => {
        const response: IBookResponse = setBookResponse(output)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

router.get(
  '/v1/books/:isbn',
  async (req: Request, res: Response<IBookResponse>, next: NextFunction): Promise<void> => {
    const { isbn } = req.params

    const bookInput: IGetBookByIsbnInput = {
      isbn,
    }

    await getBookByIsbn(req, bookInput)
      .then(async (bookOutput: IBookOutput) => {
        const bookshelfInput: IGetBookshelfInput = {
          userId: '',
          bookId: bookOutput.id,
        }

        const bookshelfOutput: IBookshelfOutput = await getBookshelf(req, bookshelfInput).catch(() => {
          return {} as IBookshelfOutput
        })

        const response: IBookResponse = setBookResponse(bookOutput, bookshelfOutput)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

router.get(
  '/v1/books/:bookId/reviews',
  async (req: Request, res: Response<IBookReviewListResponse>, next: NextFunction): Promise<void> => {
    const { bookId } = req.params
    const { limit, offset } = req.query as { [key: string]: string }

    const bookInput: IGetBookInput = {
      bookId: Number(bookId) || 0,
    }

    const reviewListInput: IListBookReviewInput = {
      bookId: Number(bookId) || 0,
      limit: Number(limit) || 100,
      offset: Number(offset) || 0,
      by: '',
      direction: '',
    }

    await getBook(req, bookInput)
      .then(async (_: IBookOutput) => {
        const reviewsOutput: IReviewListOutput = await listBookReview(req, reviewListInput)

        const usersIds: string[] = reviewsOutput.reviews.map((rv: IReviewListOutputReview) => {
          return rv.userId
        })
        const userListInput: IMultiGetUserInput = {
          userIds: Array.from(new Set(usersIds)),
        }

        const usersOutput = await multiGetUser(req, userListInput).catch(() => {
          return {} as IUserMapOutput
        })

        const response: IBookReviewListResponse = setBookReviewListResponse(reviewsOutput, usersOutput)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

router.get(
  '/v1/books/:bookId/reviews/:reviewId',
  async (req: Request, res: Response<IReviewResponse>, next: NextFunction): Promise<void> => {
    const { bookId, reviewId } = req.params

    const bookInput: IGetBookInput = {
      bookId: Number(bookId) || 0,
    }

    const reviewInput: IGetReviewInput = {
      reviewId: Number(reviewId) || 0,
    }

    await getBook(req, bookInput)
      .then(async (bookOutput: IBookOutput) => {
        const reviewOutput: IReviewOutput = await getReview(req, reviewInput)

        const userInput: IGetUserInput = {
          userId: reviewOutput.userId,
        }

        const userOutput: IUserOutput = await getUser(req, userInput)

        const response: IReviewResponse = setReviewResponse(reviewOutput, bookOutput, userOutput)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

function setBookResponse(bookOutput: IBookOutput, bookshelfOutput?: IBookshelfOutput): IBookResponse {
  const authorNames: string[] = bookOutput.authors.map((item: IBookOutputAuthor) => {
    return item.name
  })

  const authorNameKanas: string[] = bookOutput.authors.map((item: IBookOutputAuthor) => {
    return item.nameKana
  })

  const response: IBookResponse = {
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
  }

  if (bookshelfOutput) {
    const bookshelf: IBookResponseBookshelf = {
      id: bookshelfOutput.id,
      status: BookStatus[bookshelfOutput.status],
      impression: bookshelfOutput.review?.impression || '',
      readOn: bookshelfOutput.readOn,
      createdAt: bookshelfOutput.createdAt,
      updatedAt: bookshelfOutput.updatedAt,
    }

    response.bookshelf = bookshelf
  }

  return response
}

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

function setBookReviewListResponse(
  reviewOutput: IReviewListOutput,
  usersOutput: IUserMapOutput
): IBookReviewListResponse {
  const reviews = reviewOutput.reviews.map((rv: IReviewListOutputReview) => {
    const user: IBookReviewListResponseUser = {
      id: '',
      username: 'unknown',
      thumbnailUrl: '',
    }

    if (usersOutput[rv.userId]) {
      const { id, username, thumbnailUrl } = usersOutput[rv.userId]

      user.id = id
      user.username = username
      user.thumbnailUrl = thumbnailUrl
    }

    return {
      id: rv.id,
      impression: rv.impression,
      createdAt: rv.createdAt,
      updatedAt: rv.updatedAt,
      user,
    }
  })

  const response: IBookReviewListResponse = {
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
