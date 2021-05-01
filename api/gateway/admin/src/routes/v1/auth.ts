import express, { Request, Response, NextFunction } from 'express'
import multer from '~/plugins/multer'
import {
  getAuth,
  updateAuthProfile,
  updateAuthAddress,
  updateAuthEmail,
  updateAuthPassword,
  uploadAuthThumbnail,
} from '~/api'
import { IUpdateAuthRequest, IUpdateAuthEmailRequest, IUpdateAuthPasswordRequest } from '~/types/request'
import { IAuthResponse, IAuthThumbnailResponse } from '~/types/response'
import {
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
  '/v1/auth',
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
  '/v1/auth/email',
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
  '/v1/auth/password',
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
  '/v1/auth/profile',
  async (req: Request<IUpdateAuthRequest>, res: Response<IAuthResponse>, next: NextFunction): Promise<void> => {
    const {
      username,
      thumbnailUrl,
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
            thumbnailUrl,
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

router.post(
  '/v1/auth/thumbnail',
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
