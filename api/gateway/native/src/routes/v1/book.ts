import express, { NextFunction, Request, Response } from 'express'
import { createBook } from '~/api'
import { GrpcError } from '~/types/exception'

const router = express.Router()

router.post(
  '/',
  async (req: Request, res: Response<void>, next: NextFunction): Promise<void> => {
    await createBook(req)
      .then(() => res.status(200).json())
      .catch((err: GrpcError) => next(err))
  }
)

export default router
