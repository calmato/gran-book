import express, { Response } from 'express'

const router = express.Router()

router.get('/health', (_, res: Response): void => {
  res.status(200).json({ status: 'ok' })
})

router.options('/*', (_, res: Response): void => {
  res.status(200).json({ status: 'ok' })
})

export default router
