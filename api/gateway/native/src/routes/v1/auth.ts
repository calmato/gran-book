import express, { NextFunction, Request, Response } from 'express'
import { getAuth, createAuth, updateAuthProfile, updateAuthAddress, updateAuthEmail, UpdateAuthPassword } from '~/api'
import { ICreateAuthRequest, IUpdateAuthAddressRequest, IUpdateAuthEmailRequest, IUpdateAuthPasswordRequest, IUpdateAuthProfileRequest } from '~/types/request'
import { IAuthResponse } from '~/types/response'
import { ICreateAuthInput, IUpdateAuthAddressInput, IUpdateAuthEmailInput, IUpdateAuthPasswordInput, IUpdateAuthProfileInput } from '~/types/input'
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

router.post(
  '/',
  async (req: Request<ICreateAuthRequest>, res: Response<IAuthResponse>, next: NextFunction): Promise<void> => {
    const { username, email, password, passwordConfirmation } = req.body as ICreateAuthRequest

    const input: ICreateAuthInput = {
      username: username,
      email: email,
      password: password,
      passwordConfirmation: passwordConfirmation,
    }

    await createAuth(req, input)
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

router.patch(
  '/profile',
  async (req: Request<IUpdateAuthProfileRequest>, res: Response<IAuthResponse>, next: NextFunction): Promise<void> => {
    const { username, gender, thumbnail, selfIntroduction } = req.body as IUpdateAuthProfileRequest

    const input: IUpdateAuthProfileInput = {
      username: username,
      gender: gender,
      thumbnail: thumbnail,
      selfIntroduction: selfIntroduction,
    }

    await updateAuthProfile(req, input)
      .then((output: IAuthOutput) => {
        const response: IAuthResponse = setAuthResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.patch(
  '/address',
  async (req: Request<IUpdateAuthAddressRequest>, res: Response<IAuthResponse>, next: NextFunction): Promise<void> => {
    const {
      lastName,
      firstName,
      lastNameKana,
      firstNameKana,
      phoneNumber,
      postalCode,
      prefecture,
      city,
      addressLine1,
      addressLine2,
    } = req.body as IUpdateAuthAddressRequest

    const input: IUpdateAuthAddressInput = {
      lastName: lastName,
      firstName: firstName,
      lastNameKana: lastNameKana,
      firstNameKana: firstNameKana,
      phoneNumber: phoneNumber,
      postalCode: postalCode,
      prefecture: prefecture,
      city: city,
      addressLine1: addressLine1,
      addressLine2: addressLine2,
    }

    await updateAuthAddress(req, input)
      .then((output: IAuthOutput) => {
        const response: IAuthResponse = setAuthResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

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
    postalCode: output.postalCode,
    prefecture: output.prefecture,
    city: output.city,
    addressLine1: output.addressLine1,
    addressLine2: output.addressLine2,
    createdAt: output.createdAt,
    updatedAt: output.updatedAt,
  }

  return response
}

export default router
