import express, { Request, Response, NextFunction } from 'express'
import multer from '~/plugins/multer'
import {
  createAdmin,
  deleteAdmin,
  listAdmin,
  searchAdmin,
  updateAdminPassword,
  updateAdminProfile,
  updateAdminRole,
  uploadAdminThumbnail,
} from '~/api/admin'
import { badRequest } from '~/lib/http-exception'
import { GrpcError } from '~/types/exception'
import {
  ICreateAdminInput,
  IDeleteAdminInput,
  IListAdminInput,
  ISearchAdminInput,
  IUpdateAdminPasswordInput,
  IUpdateAdminProfileInput,
  IUpdateAdminRoleInput,
  IUploadAdminThumbnailInput,
} from '~/types/input'
import { IAdminListOutput, IAdminListOutputUser, IAdminOutput, IAdminThumbnailOutput } from '~/types/output'
import {
  ICreateAdminRequest,
  IUpdateAdminPasswordRequest,
  IUpdateAdminProfileRequest,
  IUpdateAdminRoleRequest,
} from '~/types/request'
import { IAdminListResponse, IAdminListResponseUser, IAdminResponse, IAdminThumbnailResponse } from '~/types/response'

const router = express.Router()

router.get(
  '/',
  async (req: Request, res: Response<IAdminListResponse>, next: NextFunction): Promise<void> => {
    const { limit, offset, by, direction, field, value } = req.query as { [key: string]: string }

    if (field && value) {
      const input: ISearchAdminInput = {
        limit: limit ? Number(limit) : 100,
        offset: offset ? Number(offset) : 0,
        by: by || '',
        direction: direction || '',
        field: field || '',
        value: value || '',
      }

      await searchAdmin(req, input)
        .then((output: IAdminListOutput) => {
          const response: IAdminListResponse = setAdminListResponse(output)
          res.status(200).json(response)
        })
        .catch((err: GrpcError) => next(err))
    } else {
      const input: IListAdminInput = {
        limit: limit ? Number(limit) : 100,
        offset: offset ? Number(offset) : 0,
        by: by || '',
        direction: direction || '',
      }

      await listAdmin(req, input)
        .then((output: IAdminListOutput) => {
          const response: IAdminListResponse = setAdminListResponse(output)
          res.status(200).json(response)
        })
        .catch((err: GrpcError) => next(err))
    }
  }
)

router.post(
  '/',
  async (req: Request, res: Response<IAdminResponse>, next: NextFunction): Promise<void> => {
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

router.delete(
  '/:userId',
  async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    const { userId } = req.params

    const input: IDeleteAdminInput = {
      userId,
    }

    await deleteAdmin(req, input)
      .then(() => {
        res.status(200).json({ status: 'ok' })
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.patch(
  '/:userId/role',
  async (req: Request, res: Response<IAdminResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params
    const { role } = req.body as IUpdateAdminRoleRequest

    const input: IUpdateAdminRoleInput = {
      id: userId,
      role,
    }

    await updateAdminRole(req, input)
      .then((output: IAdminOutput) => {
        const response: IAdminResponse = setAdminResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.patch(
  '/:userId/password',
  async (req: Request, res: Response<IAdminResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params
    const { password, passwordConfirmation } = req.body as IUpdateAdminPasswordRequest

    const input: IUpdateAdminPasswordInput = {
      id: userId,
      password,
      passwordConfirmation,
    }

    await updateAdminPassword(req, input)
      .then((output: IAdminOutput) => {
        const response: IAdminResponse = setAdminResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.patch(
  '/:userId/profile',
  async (req: Request, res: Response<IAdminResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params
    const { username, email, lastName, lastNameKana, firstName, firstNameKana } = req.body as IUpdateAdminProfileRequest

    const input: IUpdateAdminProfileInput = {
      id: userId,
      username,
      email,
      lastName,
      firstName,
      lastNameKana,
      firstNameKana,
    }

    await updateAdminProfile(req, input)
      .then((output: IAdminOutput) => {
        const response: IAdminResponse = setAdminResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.post(
  '/thumbnail',
  multer.single('thumbnail'),
  async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    const { userId } = req.params

    if (!req.file) {
      next(badRequest([{ message: 'thumbnail is not exists' }]))
      return
    }

    const input: IUploadAdminThumbnailInput = {
      userId,
      path: req.file.path,
    }

    await uploadAdminThumbnail(req, input)
      .then((output: IAdminThumbnailOutput) => {
        const response: IAdminThumbnailResponse = { thumbnailUrl: output.thumbnailUrl }
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

function setAdminListResponse(output: IAdminListOutput): IAdminListResponse {
  const users: IAdminListResponseUser[] = output.users.map(
    (u: IAdminListOutputUser): IAdminListResponseUser => ({
      id: u.id,
      username: u.username,
      email: u.email,
      phoneNumber: u.phoneNumber,
      role: u.role,
      thumbnailUrl: u.thumbnailUrl,
      selfIntroduction: u.selfIntroduction,
      lastName: u.lastName,
      firstName: u.firstName,
      lastNameKana: u.lastNameKana,
      firstNameKana: u.firstNameKana,
      createdAt: u.createdAt,
      updatedAt: u.updatedAt,
    })
  )

  const response: IAdminListResponse = {
    users,
    limit: output.limit,
    offset: output.offset,
    total: output.total,
  }

  if (output.order) {
    const { by, direction } = output.order
    response.order = {
      by,
      direction,
    }
  }

  return response
}
