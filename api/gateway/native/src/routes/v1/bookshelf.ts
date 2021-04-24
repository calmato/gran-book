import express, { NextFunction, Request, Response } from 'express'
import {
  deleteBookshelf,
  readBookshelf,
  readingBookshelf,
  releaseBookshelf,
  stackBookshelf,
  wantBookshelf,
} from '~/api'
import { GrpcError } from '~/types/exception'
import {
  IDeleteBookshelfInput,
  IReadBookshelfInput,
  IReadingBookshelfInput,
  IReleaseBookshelfInput,
  IStackBookshelfInput,
  IWantBookshelfInput,
} from '~/types/input'
import { IBookshelfOutput } from '~/types/output'
import { IReadBookshelfRequest } from '~/types/request'
import { IBookshelfResponse } from '~/types/response'

const router = express.Router()

router.post(
  '/:bookId/read',
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
  '/:bookId/reading',
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
  '/:bookId/stack',
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
  '/:bookId/want',
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
  '/:bookId/release',
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
  '/:bookId',
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
  const response: IBookshelfResponse = {
    id: output.id,
    bookId: output.bookId,
    userId: output.userId,
    status: output.status,
    impression: output.impression,
    readOn: output.readOn,
    createdAt: output.createdAt,
    updatedAt: output.updatedAt,
  }

  return response
}

export default router
