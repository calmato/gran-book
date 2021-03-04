import express, { NextFunction, Request, Response } from 'express'
import { getUserProfile } from '~/api'
import { GrpcError } from '~/types/exception'
import { IGetUserProfileInput } from '~/types/input'
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
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

export default router
