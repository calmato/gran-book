import express, { NextFunction, Request, Response } from 'express'
import { getUserProfile, registerFollow, unregisterFollow } from '~/api'
import { GrpcError } from '~/types/exception'
import { IGetUserProfileInput, IRegisterFollowInput, IUnregisterFollowInput } from '~/types/input'
import { IUserProfileOutput } from '~/types/output'
import { IUserProfileResponse } from '~/types/response'

const router = express.Router()

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
