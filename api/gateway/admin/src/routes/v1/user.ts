import express, { NextFunction, Request, Response } from 'express'
import { getUser, listUser, searchUser } from '~/api'
import { GrpcError } from '~/types/exception'
import { IGetUserInput, IListUserInput, ISearchUserInput } from '~/types/input'
import { IUserListOutput, IUserListOutputUser, IUserOutput } from '~/types/output'
import { IUserListResponse, IUserListResponseUser, IUserResponse } from '~/types/response'

const router = express.Router()

router.get(
  '/',
  async (req: Request, res: Response<IUserListResponse>, next: NextFunction): Promise<void> => {
    const { limit, offset, by, direction, field, value } = req.query as { [key: string]: string }

    if (field && value) {
      const input: ISearchUserInput = {
        limit: limit ? Number(limit) : 100,
        offset: offset ? Number(offset) : 0,
        by: by || '',
        direction: direction || '',
        field: field || '',
        value: value || '',
      }

      await searchUser(req, input)
        .then((output: IUserListOutput) => {
          const response: IUserListResponse = setUserListResponse(output)
          res.status(200).json(response)
        })
        .catch((err: GrpcError) => next(err))
    } else {
      const input: IListUserInput = {
        limit: limit ? Number(limit) : 100,
        offset: offset ? Number(offset) : 0,
        by: by || '',
        direction: direction || '',
      }

      await listUser(req, input)
        .then((output: IUserListOutput) => {
          const response: IUserListResponse = setUserListResponse(output)
          res.status(200).json(response)
        })
        .catch((err: GrpcError) => next(err))
    }
  }
)

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

function setUserListResponse(output: IUserListOutput): IUserListResponse {
  const users: IUserListResponseUser[] = output.users.map(
    (u: IUserListOutputUser): IUserListResponseUser => ({
      id: u.id,
      username: u.username,
      email: u.email,
      phoneNumber: u.phoneNumber,
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

  const response: IUserListResponse = {
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
