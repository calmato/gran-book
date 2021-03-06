import { Request } from 'express'
import { userClient } from '~/plugins/grpc'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import {
  GetUserProfileRequest,
  RegisterFollowRequest,
  UnregisterFollowRequest,
  UserProfileResponse,
} from '~/proto/user_apiv1_pb'
import { IGetUserProfileInput, IRegisterFollowInput, IUnregisterFollowInput } from '~/types/input'
import { IUserProfileOutput } from '~/types/output'

export function getUserProfile(req: Request<any>, input: IGetUserProfileInput): Promise<IUserProfileOutput> {
  const request = new GetUserProfileRequest()
  const metadata = getGrpcMetadata(req)

  request.setId(input.id)

  return new Promise((resolve: (res: IUserProfileOutput) => void, reject: (reason: Error) => void) => {
    userClient.getUserProfile(request, metadata, (err: any, res: UserProfileResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IUserProfileOutput = setUserProfileOutput(res)
      resolve(output)
    })
  })
}

export function registerFollow(req: Request<any>, input: IRegisterFollowInput): Promise<IUserProfileOutput> {
  const request = new RegisterFollowRequest()
  const metadata = getGrpcMetadata(req)

  request.setId(input.id)

  return new Promise((resolve: (res: IUserProfileOutput) => void, reject: (reason: Error) => void) => {
    userClient.registerFollow(request, metadata, (err: any, res: UserProfileResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IUserProfileOutput = setUserProfileOutput(res)
      resolve(output)
    })
  })
}

export function unregisterFollow(req: Request<any>, input: IUnregisterFollowInput): Promise<IUserProfileOutput> {
  const request = new UnregisterFollowRequest()
  const metadata = getGrpcMetadata(req)

  request.setId(input.id)

  return new Promise((resolve: (res: IUserProfileOutput) => void, reject: (reason: Error) => void) => {
    userClient.unregisterFollow(request, metadata, (err: any, res: UserProfileResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const output: IUserProfileOutput = setUserProfileOutput(res)
      resolve(output)
    })
  })
}

function setUserProfileOutput(res: UserProfileResponse): IUserProfileOutput {
  const output: IUserProfileOutput = {
    id: res.getId(),
    username: res.getUsername(),
    thumbnailUrl: res.getThumbnailUrl(),
    selfIntroduction: res.getSelfIntroduction(),
    isFollow: res.getIsFollow(),
    isFollower: res.getIsFollower(),
    followCount: res.getFollowCount(),
    followerCount: res.getFollowerCount(),
  }

  return output
}
