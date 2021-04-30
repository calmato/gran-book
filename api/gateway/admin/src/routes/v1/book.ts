import express, { Request, Response, NextFunction } from 'express'
import { deleteBook } from '~/api'
import { GrpcError } from '~/types/exception'
import { IDeleteBookInput } from '~/types/input'

const router = express.Router()

router.delete(
  '/v1/books/:bookId',
  async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    const { bookId } = req.params

    const input: IDeleteBookInput = {
      bookId: bookId ? Number(bookId) : 0,
    }

    await deleteBook(req, input)
      .then(() => {
        res.status(200).json({ status: 'ok' })
      })
      .catch((err: GrpcError) => next(err))
  }
)

export default router
