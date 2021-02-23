import express, { Request, Response, NextFunction } from 'express'
import { getAuth, updateAuthProfile, updateAuthAddress, updateAuthEmail, UpdateAuthPassword } from '~/api'
import { IUpdateAuthRequest, IUpdateAuthEmailRequest, IUpdateAuthPasswordRequest } from '~/types/request'
import { IAuthResponse } from '~/types/response'
import {
  IUpdateAuthAddressInput,
  IUpdateAuthEmailInput,
  IUpdateAuthPasswordInput,
  IUpdateAuthProfileInput,
} from '~/types/input'
import { IAuthOutput } from '~/types/output'
import { GrpcError } from '~/types/exception'

const router = express.Router()

router.get(
  '/',
  async (req: Request, res: Response<IAuthResponse>, next: NextFunction): Promise<void> => {
    await getAuth(req)
      .then((output: IAuthOutput) => {
        const response: IAuthResponse = setAuthResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.patch(
  '/email',
  async (req: Request<IUpdateAuthEmailRequest>, res: Response<IAuthResponse>, next: NextFunction): Promise<void> => {
    const { email } = req.body as IUpdateAuthEmailRequest

    const input: IUpdateAuthEmailInput = {
      email,
    }

    await updateAuthEmail(req, input)
      .then((output: IAuthOutput) => {
        const response: IAuthResponse = setAuthResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.patch(
  '/password',
  async (req: Request<IUpdateAuthPasswordRequest>, res: Response<IAuthResponse>, next: NextFunction): Promise<void> => {
    const { password, passwordConfirmation } = req.body as IUpdateAuthPasswordRequest

    const input: IUpdateAuthPasswordInput = {
      password,
      passwordConfirmation,
    }

    await UpdateAuthPassword(req, input)
      .then((output: IAuthOutput) => {
        const response: IAuthResponse = setAuthResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.patch(
  '/profile',
  async (req: Request<IUpdateAuthRequest>, res: Response<IAuthResponse>, next: NextFunction): Promise<void> => {
    const {
      username,
      thumbnail,
      selfIntroduction,
      lastName,
      firstName,
      lastNameKana,
      firstNameKana,
      phoneNumber,
    } = req.body as IUpdateAuthRequest

    await getAuth(req)
      .then(
        (output: IAuthOutput): Promise<IAuthOutput> => {
          const input: IUpdateAuthProfileInput = {
            username,
            gender: output.gender,
            thumbnail,
            selfIntroduction,
          }

          return updateAuthProfile(req, input)
        }
      )
      .then(
        (output: IAuthOutput): Promise<IAuthOutput> => {
          const input: IUpdateAuthAddressInput = {
            lastName,
            firstName,
            lastNameKana,
            firstNameKana,
            phoneNumber,
            postalCode: output.postalCode || '000-0000',
            prefecture: output.prefecture || 'Unknown',
            city: output.city || 'Unknown',
            addressLine1: output.addressLine1 || 'Unknown',
            addressLine2: output.addressLine2,
          }

          return updateAuthAddress(req, input)
        }
      )
      .then((output: IAuthOutput) => {
        const response: IAuthResponse = setAuthResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

export default router

function setAuthResponse(output: IAuthOutput): IAuthResponse {
  const response: IAuthResponse = {
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
