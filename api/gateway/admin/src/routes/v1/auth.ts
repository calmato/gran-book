import express, { Response } from 'express'

const router = express.Router()

router.get('/', (_, res: Response): void => {
  res.status(200).json({ message: 'Hello World!!' })
})

export default router
