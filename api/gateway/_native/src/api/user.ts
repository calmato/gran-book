import { Request } from 'express'
import { userClient } from '~/plugins/grpc'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import {
  IFollowInput,
  IGetUserInput,
  IGetUserProfileInput,
  IListFollowerInput,
  IListFollowInput,
  IListUserInput,
  IMultiGetUserInput,
  IUnfollowInput,
} from '~/types/input'
import {
  IFollowerListOutput,
  IFollowerListOutputFollower,
  IFollowListOutput,
  IFollowListOutputFollow,
  IUserListOutput,
  IUserListOutputUser,
  IUserMapOutput,
  IUserMapOutputUser,
  IUserOutput,
  IUserProfileOutput,
} from '~/types/output'
import {
  FollowerListResponse,
  FollowListResponse,
  FollowRequest,
  GetUserProfileRequest,
  GetUserRequest,
  ListFollowerRequest,
  ListFollowRequest,
  ListUserRequest,
  MultiGetUserRequest,
  UnfollowRequest,
  UserListResponse,
  UserMapResponse,
  UserProfileResponse,
  UserResponse,
} from '~/proto/user_service_pb'
import { Order, Search } from '~/proto/common_pb'

export function listUser(req: Request<any>, input: IListUserInput): Promise<IUserListOutput> {
  const request = new ListUserRequest()
  const metadata = getGrpcMetadata(req)

  request.setLimit(input.limit)
  request.setOffset(input.offset)

  if (input.search) {
    const search = new Search()
    search.setField(input.search.field)
    search.setValue(input.search.value)

    request.setSearch(search)
  }

  if (input.order) {
    const order = new Order()
    order.setField(input.order.field)
    order.setOrderBy(input.order.orderBy)

    request.setOrder(order)
  }

  return new Promise((resolve: (res: IUserListOutput) => void, reject: (reason: Error) => void) => {
    userClient.listUser(request, metadata, (err: any, res: UserListResponse) => {
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

  request.setUserId(input.userId)
  request.setLimit(input.limit)
  request.setOffset(input.offset)

  if (input.order) {
    const order = new Order()
    order.setField(input.order.field)
    order.setOrderBy(input.order.orderBy)

    request.setOrder(order)
  }

  return new Promise((resolve: (res: IFollowListOutput) => void, reject: (reason: Error) => void) => {
    userClient.listFollow(request, metadata, (err: any, res: FollowListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setFollowListOutput(res))
    })
  })
}

export function listFollower(req: Request<any>, input: IListFollowerInput): Promise<IFollowerListOutput> {
  const request = new ListFollowerRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)
  request.setLimit(input.limit)
  request.setOffset(input.offset)

  if (input.order) {
    const order = new Order()
    order.setField(input.order.field)
    order.setOrderBy(input.order.orderBy)

    request.setOrder(order)
  }

  return new Promise((resolve: (res: IFollowerListOutput) => void, reject: (reason: Error) => void) => {
    userClient.listFollower(request, metadata, (err: any, res: FollowerListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setFollowerListOutput(res))
    })
  })
}

export function multiGetUser(req: Request<any>, input: IMultiGetUserInput): Promise<IUserMapOutput> {
  const request = new MultiGetUserRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserIdsList(input.userIds)

  return new Promise((resolve: (res: IUserMapOutput) => void, reject: (reason: Error) => void) => {
    userClient.multiGetUser(request, metadata, (err: any, res: UserMapResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setUserMapOutput(res))
    })
  })
}

export function getUser(req: Request<any>, input: IGetUserInput): Promise<IUserOutput> {
  const request = new GetUserRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)

  return new Promise((resolve: (res: IUserOutput) => void, reject: (reason: Error) => void) => {
    userClient.getUser(request, metadata, (err: any, res: UserResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setUserOutput(res))
    })
  })
}

export function getUserProfile(req: Request<any>, input: IGetUserProfileInput): Promise<IUserProfileOutput> {
  const request = new GetUserProfileRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)

  return new Promise((resolve: (res: IUserProfileOutput) => void, reject: (reason: Error) => void) => {
    userClient.getUserProfile(request, metadata, (err: any, res: UserProfileResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setUserProfileOutput(res))
    })
  })
}

export function follow(req: Request<any>, input: IFollowInput): Promise<IUserProfileOutput> {
  const request = new FollowRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)

  return new Promise((resolve: (res: IUserProfileOutput) => void, reject: (reason: Error) => void) => {
    userClient.follow(request, metadata, (err: any, res: UserProfileResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setUserProfileOutput(res))
    })
  })
}

export function unfollow(req: Request<any>, input: IUnfollowInput): Promise<IUserProfileOutput> {
  const request = new UnfollowRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)

  return new Promise((resolve: (res: IUserProfileOutput) => void, reject: (reason: Error) => void) => {
    userClient.unfollow(request, metadata, (err: any, res: UserProfileResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setUserProfileOutput(res))
    })
  })
}

function setUserOutput(res: UserResponse): IUserOutput {
  const user: UserResponse.AsObject = res.toObject()

  const output: IUserOutput = { ...user }
  return output
}

function setUserListOutput(res: UserListResponse): IUserListOutput {
  const users: Array<IUserListOutputUser> = res.getUsersList().map(
    (value: UserListResponse.User): IUserListOutputUser => {
      const user: UserListResponse.User.AsObject = value.toObject()
      return { ...user }
    }
  )

  const output: IUserListOutput = {
    users,
    limit: res.getLimit(),
    offset: res.getOffset(),
    total: res.getTotal(),
  }

  return output
}

function setUserMapOutput(res: UserMapResponse): IUserMapOutput {
  const output: IUserMapOutput = {}

  res.toObject().usersMap.map((value: [string, UserMapResponse.User.AsObject]) => {
    const key: string = value[0]
    const res: UserMapResponse.User.AsObject = value[1]

    const user: IUserMapOutputUser = { ...res }
    output[key] = user
  })

  return output
}

function setUserProfileOutput(res: UserProfileResponse): IUserProfileOutput {
  const profile: UserProfileResponse.AsObject = res.toObject()

  const output: IUserProfileOutput = { ...profile }
  return output
}

function setFollowListOutput(res: FollowListResponse): IFollowListOutput {
  const follows: Array<IFollowListOutputFollow> = res.getFollowsList().map(
    (value: FollowListResponse.Follow): IFollowListOutputFollow => {
      const follow: FollowListResponse.Follow.AsObject = value.toObject()
      return { ...follow }
    }
  )

  const output: IFollowListOutput = {
    follows,
    limit: res.getLimit(),
    offset: res.getOffset(),
    total: res.getTotal(),
  }

  return output
}

function setFollowerListOutput(res: FollowerListResponse): IFollowerListOutput {
  const followers: Array<IFollowerListOutputFollower> = res.getFollowersList().map(
    (value: FollowerListResponse.Follower): IFollowerListOutputFollower => {
      const follower: FollowerListResponse.Follower.AsObject = value.toObject()
      return { ...follower }
    }
  )

  const output: IFollowerListOutput = {
    followers,
    limit: res.getLimit(),
    offset: res.getOffset(),
    total: res.getTotal(),
  }

  return output
}
