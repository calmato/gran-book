import express, { NextFunction, Request, Response } from 'express'
import { getUserProfile, listFollow, listFollower, registerFollow, unregisterFollow } from '~/api'
import { GrpcError } from '~/types/exception'
import {
  IGetUserProfileInput,
  IListFollowerInput,
  IListFollowInput,
  IRegisterFollowInput,
  IUnregisterFollowInput,
} from '~/types/input'
import {
  IFollowerListOutput,
  IFollowerListOutputUser,
  IFollowListOutput,
  IFollowListOutputUser,
  IUserProfileOutput,
} from '~/types/output'
import {
  IFollowerListResponse,
  IFollowerListResponseUser,
  IFollowListResponse,
  IFollowListResponseUser,
  IUserProfileResponse,
} from '~/types/response'

const router = express.Router()

router.get(
  '/:userId/follow',
  async (req: Request, res: Response<IFollowListResponse>, next: NextFunction): Promise<void> => {
    const { limit, offset } = req.query
    const { userId } = req.params

    const input: IListFollowInput = {
      id: userId || '',
      limit: Number(limit) || 100,
      offset: Number(offset) || 0,
    }

    await listFollow(req, input)
      .then((output: IFollowListOutput) => {
        const users: IFollowListResponseUser[] = output.users.map(
          (user: IFollowListOutputUser): IFollowListResponseUser => {
            return {
              id: user.id,
              username: user.username,
              thumbnailUrl: user.thumbnailUrl,
              selfIntroduction: user.selfIntroduction,
            }
          }
        )

        const response: IFollowListResponse = {
          users,
          limit: output.limit,
          offset: output.offset,
          total: output.total,
        }

        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.post(
  '/:userId/follow',
  async (req: Request, res: Response<IUserProfileResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params

    const input: IRegisterFollowInput = {
      id: userId || '',
    }

    await registerFollow(req, input)
      .then((output: IUserProfileOutput) => {
        const response: IUserProfileResponse = setUserProfileResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.delete(
  '/:userId/follow',
  async (req: Request, res: Response<IUserProfileResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params

    const input: IUnregisterFollowInput = {
      id: userId || '',
    }

    await unregisterFollow(req, input)
      .then((output: IUserProfileOutput) => {
        const response: IUserProfileResponse = setUserProfileResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.get(
  '/:userId/follower',
  async (req: Request, res: Response<IFollowerListResponse>, next: NextFunction): Promise<void> => {
    const { limit, offset } = req.query
    const { userId } = req.params

    const input: IListFollowerInput = {
      id: userId || '',
      limit: Number(limit) || 100,
      offset: Number(offset) || 0,
    }

    await listFollower(req, input)
      .then((output: IFollowerListOutput) => {
        const users: IFollowerListResponseUser[] = output.users.map(
          (user: IFollowerListOutputUser): IFollowerListResponseUser => {
            return {
              id: user.id,
              username: user.username,
              thumbnailUrl: user.thumbnailUrl,
              selfIntroduction: user.selfIntroduction,
            }
          }
        )

        const response: IFollowerListResponse = {
          users,
          limit: output.limit,
          offset: output.offset,
          total: output.total,
        }

        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.get(
  '/:userId/profile',
  async (req: Request, res: Response<IUserProfileResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params

    const input: IGetUserProfileInput = {
      id: userId || '',
    }

    await getUserProfile(req, input)
      .then((output: IUserProfileOutput) => {
        const response: IUserProfileResponse = setUserProfileResponse(output)
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

function setUserProfileResponse(output: IUserProfileOutput): IUserProfileResponse {
  const response: IUserProfileResponse = {
    id: output.id,
    username: output.username,
    thumbnailUrl: output.thumbnailUrl,
    selfIntroduction: output.selfIntroduction,
    isFollow: output.isFollow,
    isFollower: output.isFollower,
    followCount: output.followCount,
    followerCount: output.followerCount,
    reviewCount: 0,
    rating: 0,
    products: [],
  }

  return response
}

export default router
