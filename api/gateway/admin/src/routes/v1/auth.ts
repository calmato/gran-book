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
  '/',
  async (req: Request<IUpdateAuthRequest>, res: Response<IAuthResponse>, next: NextFunction): Promise<void> => {
    const {
      username,
      gender,
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
        (): Promise<IAuthOutput> => {
          const input: IUpdateAuthProfileInput = {
            username: username,
            gender: gender,
            thumbnail: thumbnail,
            selfIntroduction: selfIntroduction,
          }

          return updateAuthProfile(req, input)
        }
      )
      .then(
        (output: IAuthOutput): Promise<IAuthOutput> => {
          const input: IUpdateAuthAddressInput = {
            lastName: lastName,
            firstName: firstName,
            lastNameKana: lastNameKana,
            firstNameKana: firstNameKana,
            phoneNumber: phoneNumber,
            postalCode: output.postalCode,
            prefecture: output.prefecture,
            city: output.city,
            addressLine1: output.addressLine1,
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

router.patch(
  '/email',
  async (req: Request<IUpdateAuthEmailRequest>, res: Response<IAuthResponse>, next: NextFunction): Promise<void> => {
    const { email } = req.body as IUpdateAuthEmailRequest

    const input: IUpdateAuthEmailInput = {
      email: email,
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
      password: password,
      passwordConfirmation: passwordConfirmation,
    }

    await UpdateAuthPassword(req, input)
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

  return response
}
