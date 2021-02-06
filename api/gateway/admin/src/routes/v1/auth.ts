import express, { Response } from 'express'
import { createAuth } from '~/api'
import { ICreateAuthInput } from '~/types/input'
import { IAuthOutput } from '~/types/output'

const router = express.Router()

router.get('/', (_, res: Response): void => {
  res.status(200).json({ message: 'Hello World!!' })
})

router.post('/', async (_, res: Response): Promise<void> => {
  const input: ICreateAuthInput = {
    username: '',
    email: '',
    password: '',
    passwordConfirmation: '',
  }

  await createAuth(input).then((output: IAuthOutput) => {
    // TODO: responseの型定義
    res.status(200).json(output)
  })
  .catch((err: Error) => {
    // TODO: status制御
    // TODO: responseの型定義
    res.status(400).json(err)
  })
})

export default router
