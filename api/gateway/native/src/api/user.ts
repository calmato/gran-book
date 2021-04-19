import { Request } from 'express'
import { userClient } from '~/plugins/grpc'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import {
  FollowerListResponse,
  FollowListResponse,
  GetUserProfileRequest,
  ListFollowerRequest,
  ListFollowRequest,
  ListUserByUserIdsRequest,
  RegisterFollowRequest,
  UnregisterFollowRequest,
  UserListResponse,
  UserProfileResponse,
} from '~/proto/user_apiv1_pb'
import {
  IGetUserProfileInput,
  IListFollowerInput,
  IListFollowInput,
  IListUserByUserIdsInput,
  IRegisterFollowInput,
  IUnregisterFollowInput,
} from '~/types/input'
import {
  IFollowerListOutput,
  IFollowerListOutputUser,
  IFollowListOutput,
  IFollowListOutputUser,
  IUserListOutput,
  IUserListOutputOrder,
  IUserListOutputUser,
  IUserProfileOutput,
} from '~/types/output'

export function listUserWithUserIds(req: Request<any>, input: IListUserByUserIdsInput): Promise<IUserListOutput> {
  const request = new ListUserByUserIdsRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserIdsList(input.ids)

  return new Promise((resolve: (res: IUserListOutput) => void, reject: (reason: Error) => void) => {
    userClient.listUserByUserIds(request, metadata, (err: any, res: UserListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setUserListOutput(res))
    })
  })
}

export function listFollow(req: Request<any>, input: IListFollowInput): Promise<IFollowListOutput> {
  const request = new ListFollowRequest()
  const metadata = getGrpcMetadata(req)

  request.setId(input.id)
  request.setLimit(input.limit)
  request.setOffset(input.offset)

  return new Promise((resolve: (res: IFollowListOutput) => void, reject: (reason: Error) => void) => {
    userClient.listFollow(request, metadata, (err: any, res: FollowListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const users: IFollowListOutputUser[] = res.getUsersList().map(
        (user: FollowListResponse.User): IFollowListOutputUser => {
          return {
            id: user.getId(),
            username: user.getUsername(),
            thumbnailUrl: user.getThumbnailUrl(),
            selfIntroduction: user.getSelfIntroduction(),
            isFollow: user.getIsFollow(),
          }
        }
      )

      const output: IFollowListOutput = {
        users,
        limit: res.getLimit(),
        offset: res.getOffset(),
        total: res.getTotal(),
      }

      resolve(output)
    })
  })
}

export function listFollower(req: Request<any>, input: IListFollowerInput): Promise<IFollowerListOutput> {
  const request = new ListFollowerRequest()
  const metadata = getGrpcMetadata(req)

  request.setId(input.id)
  request.setLimit(input.limit)
  request.setOffset(input.offset)

  return new Promise((resolve: (res: IFollowerListOutput) => void, reject: (reason: Error) => void) => {
    userClient.listFollower(request, metadata, (err: any, res: FollowerListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      const users: IFollowerListOutputUser[] = res.getUsersList().map(
        (user: FollowerListResponse.User): IFollowerListOutputUser => {
          return {
            id: user.getId(),
            username: user.getUsername(),
            thumbnailUrl: user.getThumbnailUrl(),
            selfIntroduction: user.getSelfIntroduction(),
            isFollow: user.getIsFollow(),
          }
        }
      )

      const output: IFollowerListOutput = {
        users,
        limit: res.getLimit(),
        offset: res.getOffset(),
        total: res.getTotal(),
      }

      resolve(output)
    })
  })
}

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

function setUserListOutput(res: UserListResponse): IUserListOutput {
  const users: IUserListOutputUser[] = res.getUsersList().map(
    (u: UserListResponse.User): IUserListOutputUser => ({
      id: u.getId(),
      username: u.getUsername(),
      email: u.getEmail(),
      phoneNumber: u.getPhoneNumber(),
      role: u.getRole(),
      thumbnailUrl: u.getThumbnailUrl(),
      selfIntroduction: u.getSelfIntroduction(),
      lastName: u.getLastName(),
      firstName: u.getFirstName(),
      lastNameKana: u.getLastNameKana(),
      firstNameKana: u.getFirstNameKana(),
      createdAt: u.getCreatedAt(),
      updatedAt: u.getUpdatedAt(),
    })
  )

  const output: IUserListOutput = {
    users,
    limit: res.getLimit(),
    offset: res.getOffset(),
    total: res.getTotal(),
  }

  const orderRes: UserListResponse.Order | undefined = res.getOrder()
  if (orderRes) {
    const order: IUserListOutputOrder = {
      by: orderRes.getBy(),
      direction: orderRes.getDirection(),
    }

    output.order = order
  }

  return output
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
