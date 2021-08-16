import { Request } from 'express'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import { userClient } from '~/plugins/grpc'
import {
  GetUserRequest,
  ListUserRequest,
  SearchUserRequest,
  UserListResponse,
  UserResponse,
} from '~/proto/user_apiv1_pb'
import { IGetUserInput, IListUserInput, ISearchUserInput } from '~/types/input'
import { IUserListOutput, IUserListOutputOrder, IUserListOutputUser, IUserOutput } from '~/types/output'

export function listUser(req: Request<any>, input: IListUserInput): Promise<IUserListOutput> {
  const request = new ListUserRequest()
  const metadata = getGrpcMetadata(req)

  const order = new ListUserRequest.Order()
  order.setBy(input.by)
  order.setDirection(input.direction)

  request.setLimit(input.limit)
  request.setOffset(input.offset)
  request.setOrder(order)

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

export function searchUser(req: Request<any>, input: ISearchUserInput): Promise<IUserListOutput> {
  const request = new SearchUserRequest()
  const metadata = getGrpcMetadata(req)

  const search = new SearchUserRequest.Search()
  search.setField(input.field)
  search.setValue(input.value)

  request.setLimit(input.limit)
  request.setOffset(input.offset)
  request.setSearch(search)

  if (input.by !== '') {
    const order = new SearchUserRequest.Order()
    order.setBy(input.by)
    order.setDirection(input.direction)

    request.setOrder(order)
  }

  return new Promise((resolve: (res: IUserListOutput) => void, reject: (reason: Error) => void) => {
    userClient.searchUser(request, metadata, (err: any, res: UserListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setUserListOutput(res))
    })
  })
}

export function getUser(req: Request<any>, input: IGetUserInput): Promise<IUserOutput> {
  const request = new GetUserRequest()
  const metadata = getGrpcMetadata(req)

  request.setId(input.id)

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

function setUserOutput(res: UserResponse): IUserOutput {
  const output: IUserOutput = {
    id: res.getId(),
    username: res.getUsername(),
    email: res.getEmail(),
    phoneNumber: res.getPhoneNumber(),
    role: res.getRole(),
    thumbnailUrl: res.getThumbnailUrl(),
    selfIntroduction: res.getSelfIntroduction(),
    lastName: res.getLastName(),
    firstName: res.getFirstName(),
    lastNameKana: res.getLastNameKana(),
    firstNameKana: res.getFirstNameKana(),
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
  }

  return output
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
