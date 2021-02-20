import express, { Request, Response, NextFunction } from 'express'
import { createAdmin } from '~/api/admin'
import { GrpcError } from '~/types/exception'
import { ICreateAdminInput } from '~/types/input'
import { IAdminOutput } from '~/types/output'
import { ICreateAdminRequest } from '~/types/request'
import { IAdminResponse } from '~/types/response'

const router = express.Router()

router.post(
  '/',
  async (req: Request, res: Response<IAdminResponse>, next: NextFunction): Promise<IAdminResponse> => {
    const {
      username,
      email,
      password,
      passwordConfirmation,
      role,
      lastName,
      firstName,
      lastNameKana,
      firstNameKana,
    } = req.body as ICreateAdminRequest

    const input: ICreateAdminInput = {
      username,
      email,
      password,
      passwordConfirmation,
      role,
      lastName,
      firstName,
      lastNameKana,
      firstNameKana,
    }

    await createAdmin(req, input)
      .then((output: IAdminOutput) => {
        const response: IAdminResponse = setAdminResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

export default router

function setAdminResponse(output: IAdminOutput): IAdminResponse {
  const response: IAdminResponse = {
    id: output.id,
    username: output.username,
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

  return response
}
