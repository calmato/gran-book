import express, { NextFunction, Request, Response } from 'express'
import multer from '~/plugins/multer'
import {
  getAuth,
  createAuth,
  updateAuthProfile,
  updateAuthAddress,
  updateAuthEmail,
  updateAuthPassword,
  uploadAuthThumbnail,
  deleteAuth,
  registerAuthDevice,
} from '~/api'
import {
  ICreateAuthRequest,
  IRegisterAuthDeviceRequest,
  IUpdateAuthAddressRequest,
  IUpdateAuthEmailRequest,
  IUpdateAuthPasswordRequest,
  IUpdateAuthProfileRequest,
} from '~/types/request'
import { IAuthResponse, IAuthThumbnailResponse } from '~/types/response'
import {
  ICreateAuthInput,
  IRegisterAuthDeviceInput,
  IUpdateAuthAddressInput,
  IUpdateAuthEmailInput,
  IUpdateAuthPasswordInput,
  IUpdateAuthProfileInput,
  IUploadAuthThumbnailInput,
} from '~/types/input'
import { IAuthOutput, IAuthThumbnailOutput } from '~/types/output'
import { GrpcError } from '~/types/exception'
import { badRequest } from '~/lib/http-exception'

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
      username,
      email,
      password,
      passwordConfirmation,
    }

    await createAuth(req, input)
      .then((output: IAuthOutput) => {
        const response: IAuthResponse = setAuthResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.delete(
  '/',
  async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    await deleteAuth(req)
      .then(() => {
        res.status(200).json({ message: 'ok' })
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

    await updateAuthPassword(req, input)
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
    const { username, gender, thumbnailUrl, selfIntroduction } = req.body as IUpdateAuthProfileRequest

    const input: IUpdateAuthProfileInput = {
      username,
      gender,
      thumbnailUrl,
      selfIntroduction,
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
    }

    await updateAuthAddress(req, input)
      .then((output: IAuthOutput) => {
        const response: IAuthResponse = setAuthResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.post(
  '/thumbnail',
  multer.single('thumbnail'),
  async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    if (!req.file) {
      next(badRequest([{ message: 'thumbnail is not exists' }]))
      return
    }

    const input: IUploadAuthThumbnailInput = {
      path: req.file.path,
    }

    await uploadAuthThumbnail(req, input)
      .then((output: IAuthThumbnailOutput) => {
        const response: IAuthThumbnailResponse = { thumbnailUrl: output.thumbnailUrl }
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.post(
  '/device',
  async (req: Request<IRegisterAuthDeviceRequest>, res: Response<IAuthResponse>, next: NextFunction): Promise<void> => {
    const { instanceId } = req.body as IRegisterAuthDeviceRequest

    const input: IRegisterAuthDeviceInput = {
      instanceId,
    }

    await registerAuthDevice(req, input)
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
