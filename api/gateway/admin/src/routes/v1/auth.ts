import express, { Request, Response } from 'express'
import { createAuth } from '~/api'
import { ICreateAuthRequest } from '~/types/request'
import { IAuthResponse } from '~/types/response'
import { ICreateAuthInput } from '~/types/input'
import { IAuthOutput } from '~/types/output'

const router = express.Router()

router.get('/', (_, res: Response): void => {
  res.status(200).json({ message: 'Hello World!!' })
})

router.post('/', async ({ body }: Request<ICreateAuthRequest>, res: Response<IAuthResponse|Error>): Promise<void> => {
  const input: ICreateAuthInput = {
    username: body.username,
    email: body.email,
    password: body.password,
    passwordConfirmation: body.passwordConfirmation,
  }

  await createAuth(input).then((output: IAuthOutput) => {
    const response: IAuthResponse = {
      id: output.id,
      username: output.username,
      gender: output.gender,
      email: output.email,
      phoneNumber: output.phoneNumber,
      role: output.role,
      thumbnailUrl: output.thumbnailUrl,
      selfIntroduction: output.selfIntroduction,
      lastName: output.lastName,
      firstName: output.firstName,
      lastNameKana: output.lastNameKana,
      firstNameKana: output.firstNameKana,
      createdAt: output.createdAt,
      updatedAt: output.updatedAt,
    }

    res.status(200).json(response)
  })
  .catch((err: Error) => {
    // TODO: status制御
    // TODO: responseの型定義
    res.status(400).json(err)
  })
})

export default router
