import express, { Request, Response, NextFunction } from 'express'
import { createAuth } from '~/api'
import { ICreateAuthRequest } from '~/types/request'
import { IAuthResponse } from '~/types/response'
import { ICreateAuthInput } from '~/types/input'
import { IAuthOutput } from '~/types/output'
import { GrpcError } from '~/types/exception'

const router = express.Router()

router.get('/', (_: Request , res: Response): void => {
  res.status(200).json({ message: 'Hello World!!' })
})

router.post('/', async (req: Request<ICreateAuthRequest>, res: Response<IAuthResponse>, next: NextFunction): Promise<void> => {
  const { username, email, password, passwordConfirmation } = req.body as ICreateAuthRequest

  const input: ICreateAuthInput = {
    username: username,
    email: email,
    password: password,
    passwordConfirmation: passwordConfirmation,
  }

  await createAuth(input)
    .then((output: IAuthOutput) => {
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
    .catch((err: GrpcError) => next(err))
})

export default router
