import express, { NextFunction, Request, Response } from 'express'
import { getUser } from '~/api'
import { GrpcError } from '~/types/exception'
import { IGetUserInput } from '~/types/input'
import { IUserOutput } from '~/types/output'
import { IUserResponse } from '~/types/response'

const router = express.Router()

router.get(
  '/:userId',
  async (req: Request, res: Response<IUserResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params

    const input: IGetUserInput = {
      id: userId,
    }

    await getUser(req, input)
      .then((output: IUserOutput) => {
        const response: IUserResponse = setUserResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

export default router

function setUserResponse(output: IUserOutput): IUserResponse {
  const response: IUserResponse = {
    id: output.id,
    username: output.username,
    email: output.email,
    phoneNumber: output.phoneNumber,
    thumbnailUrl: output.thumbnailUrl,
    selfIntroduction: output.selfIntroduction,
    lastName: output.lastName,
    firstName: output.firstName,
    lastNameKana: output.lastNameKana,
    firstNameKana: output.firstNameKana,
    createdAt: output.createdAt,
    updatedAt: output.updatedAt,
  }

  return response
}
