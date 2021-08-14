import express, { NextFunction, Request, Response } from 'express'
import { follow, getUserProfile, listFollow, listFollower, unfollow } from '~/api'
import { GrpcError } from '~/types/exception'
import { IFollowInput, IGetUserProfileInput, IListFollowerInput, IListFollowInput, IUnfollowInput } from '~/types/input'
import {
  IFollowerListOutput,
  IFollowerListOutputFollower,
  IFollowListOutput,
  IFollowListOutputFollow,
  IUserProfileOutput,
} from '~/types/output'
import {
  IFollowerListResponse,
  IFollowerListResponseUser,
  IFollowListResponse,
  IFollowListResponseUser,
  IUserProfileResponse,
} from '~/types/response'
import { LIST_DEFAULT_LIMIT, LIST_DEFAULT_OFFSET } from '~/util'

const router = express.Router()

router.post(
  '/v1/users/:userId/follow/:followerId',
  async (req: Request, res: Response<IUserProfileResponse>, next: NextFunction): Promise<void> => {
    const { userId, followerId } = req.params

    const input: IFollowInput = {
      userId: userId || '',
      followerId: followerId || '',
    }

    await follow(req, input)
      .then((output: IUserProfileOutput) => {
        const response: IUserProfileResponse = setUserProfileResponse(output)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

router.delete(
  '/v1/users/:userId/follow/:followerId',
  async (req: Request, res: Response<IUserProfileResponse>, next: NextFunction): Promise<void> => {
    const { userId, followerId } = req.params

    const input: IUnfollowInput = {
      userId: userId || '',
      followerId: followerId || '',
    }

    await unfollow(req, input)
      .then((output: IUserProfileOutput) => {
        const response: IUserProfileResponse = setUserProfileResponse(output)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

router.get(
  '/v1/users/:userId/follows',
  async (req: Request, res: Response<IFollowListResponse>, next: NextFunction): Promise<void> => {
    const { limit, offset } = req.query
    const { userId } = req.params

    const input: IListFollowInput = {
      userId: userId || '',
      limit: Number(limit) || LIST_DEFAULT_LIMIT,
      offset: Number(offset) || LIST_DEFAULT_OFFSET,
    }

    await listFollow(req, input)
      .then((output: IFollowListOutput) => {
        const response: IFollowListResponse = setFollowListResponse(output)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

router.get(
  '/v1/users/:userId/followers',
  async (req: Request, res: Response<IFollowerListResponse>, next: NextFunction): Promise<void> => {
    const { limit, offset } = req.query
    const { userId } = req.params

    const input: IListFollowerInput = {
      userId: userId || '',
      limit: Number(limit) || 100,
      offset: Number(offset) || 0,
    }

    await listFollower(req, input)
      .then((output: IFollowerListOutput) => {
        const response: IFollowerListResponse = setFollowerListResponse(output)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

router.get(
  '/v1/users/:userId/profile',
  async (req: Request, res: Response<IUserProfileResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params

    const input: IGetUserProfileInput = {
      userId: userId || '',
    }

    await getUserProfile(req, input)
      .then((output: IUserProfileOutput) => {
        const response: IUserProfileResponse = setUserProfileResponse(output)
        res.status(200).json(response)
      })
      .catch((err: Error) => {
        next(err)
      })
  }
)

function setFollowListResponse(output: IFollowListOutput): IFollowListResponse {
  const users: Array<IFollowListResponseUser> = output.follows.map(
    (value: IFollowListOutputFollow): IFollowListResponseUser => {
      const user: IFollowListResponseUser = {
        id: value.id,
        username: value.username,
        thumbnailUrl: value.thumbnailUrl,
        selfIntroduction: value.selfIntroduction,
        isFollow: value.isFollow,
      }

      return user
    }
  )

  const response: IFollowListResponse = {
    users,
    limit: output.limit,
    offset: output.offset,
    total: output.total,
  }

  return response
}

function setFollowerListResponse(output: IFollowerListOutput): IFollowerListResponse {
  const users: Array<IFollowerListResponseUser> = output.followers.map(
    (value: IFollowerListOutputFollower): IFollowerListResponseUser => {
      const user: IFollowerListResponseUser = {
        id: value.id,
        username: value.username,
        thumbnailUrl: value.thumbnailUrl,
        selfIntroduction: value.selfIntroduction,
        isFollow: value.isFollow,
      }

      return user
    }
  )

  const response: IFollowerListResponse = {
    users,
    limit: output.limit,
    offset: output.offset,
    total: output.total,
  }

  return response
}

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
