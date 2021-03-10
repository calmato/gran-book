import { Request } from 'express'
import { adminClient } from '~/plugins/grpc'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import {
  AdminListResponse,
  AdminResponse,
  CreateAdminRequest,
  ListAdminRequest,
  SearchAdminRequest,
  UpdateAdminPasswordRequest,
  UpdateAdminProfileRequest,
  UpdateAdminRoleRequest,
} from '~/proto/user_apiv1_pb'
import {
  ICreateAdminInput,
  IListAdminInput,
  ISearchAdminInput,
  IUpdateAdminPasswordInput,
  IUpdateAdminProfileInput,
  IUpdateAdminRoleInput,
} from '~/types/input'
import { IAdminListOutput, IAdminListOutputOrder, IAdminListOutputUser, IAdminOutput } from '~/types/output'

export function listAdmin(req: Request<any>, input: IListAdminInput): Promise<IAdminListOutput> {
  const request = new ListAdminRequest()
  const metadata = getGrpcMetadata(req)

  request.setLimit(input.limit)
  request.setOffset(input.offset)

  if (input.by !== '') {
    const order = new ListAdminRequest.Order()
    order.setBy(input.by)
    order.setDirection(input.direction)

    request.setOrder(order)
  }

  return new Promise((resolve: (res: IAdminListOutput) => void, reject: (reason: Error) => void) => {
    adminClient.listAdmin(request, metadata, (err: any, res: AdminListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setAdminListOutput(res))
    })
  })
}

export function searchAdmin(req: Request<any>, input: ISearchAdminInput): Promise<IAdminListOutput> {
  const request = new SearchAdminRequest()
  const metadata = getGrpcMetadata(req)

  const search = new SearchAdminRequest.Search()
  search.setField(input.field)
  search.setValue(input.value)

  request.setLimit(input.limit)
  request.setOffset(input.offset)
  request.setSearch(search)

  if (input.by !== '') {
    const order = new SearchAdminRequest.Order()
    order.setBy(input.by)
    order.setDirection(input.direction)

    request.setOrder(order)
  }

  return new Promise((resolve: (res: IAdminListOutput) => void, reject: (reason: Error) => void) => {
    adminClient.searchAdmin(request, metadata, (err: any, res: AdminListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setAdminListOutput(res))
    })
  })
}

export function createAdmin(req: Request<any>, input: ICreateAdminInput): Promise<IAdminOutput> {
  const request = new CreateAdminRequest()
  const metadata = getGrpcMetadata(req)

  request.setUsername(input.username)
  request.setEmail(input.email)
  request.setPassword(input.password)
  request.setPasswordConfirmation(input.passwordConfirmation)
  request.setRole(input.role)
  request.setLastName(input.lastName)
  request.setFirstName(input.firstName)
  request.setLastNameKana(input.lastNameKana)
  request.setFirstNameKana(input.firstNameKana)

  return new Promise((resolve: (res: IAdminOutput) => void, reject: (reason: Error) => void) => {
    adminClient.createAdmin(request, metadata, (err: any, res: AdminResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setAdminOutput(res))
    })
  })
}

export function updateAdminRole(req: Request<any>, input: IUpdateAdminRoleInput): Promise<IAdminOutput> {
  const request = new UpdateAdminRoleRequest()
  const metadata = getGrpcMetadata(req)

  request.setId(input.id)
  request.setRole(input.role)

  return new Promise((resolve: (res: IAdminOutput) => void, reject: (reason: Error) => void) => {
    adminClient.updateAdminRole(request, metadata, (err: any, res: AdminResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setAdminOutput(res))
    })
  })
}

export function updateAdminPassword(req: Request<any>, input: IUpdateAdminPasswordInput): Promise<IAdminOutput> {
  const request = new UpdateAdminPasswordRequest()
  const metadata = getGrpcMetadata(req)

  request.setId(input.id)
  request.setPassword(input.password)
  request.setPasswordConfirmation(input.passwordConfirmation)

  return new Promise((resolve: (res: IAdminOutput) => void, reject: (reason: Error) => void) => {
    adminClient.updateAdminPassword(request, metadata, (err: any, res: AdminResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setAdminOutput(res))
    })
  })
}

export function updateAdminProfile(req: Request<any>, input: IUpdateAdminProfileInput): Promise<IAdminOutput> {
  const request = new UpdateAdminProfileRequest()
  const metadata = getGrpcMetadata(req)

  request.setId(input.id)
  request.setUsername(input.username)
  request.setEmail(input.email)
  request.setLastName(input.lastName)
  request.setFirstName(input.firstName)
  request.setLastNameKana(input.lastNameKana)
  request.setFirstNameKana(input.firstNameKana)

  return new Promise((resolve: (res: IAdminOutput) => void, reject: (reason: Error) => void) => {
    adminClient.updateAdminProfile(request, metadata, (err: any, res: AdminResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setAdminOutput(res))
    })
  })
}

function setAdminOutput(res: AdminResponse): IAdminOutput {
  const output: IAdminOutput = {
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
    activated: res.getActivated(),
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
  }

  return output
}

function setAdminListOutput(res: AdminListResponse): IAdminListOutput {
  const users: IAdminListOutputUser[] = res.getUsersList().map(
    (u: AdminListResponse.User): IAdminListOutputUser => ({
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
      activated: u.getActivated(),
      createdAt: u.getCreatedAt(),
      updatedAt: u.getUpdatedAt(),
    })
  )

  const output: IAdminListOutput = {
    users,
    limit: res.getLimit(),
    offset: res.getOffset(),
    total: res.getTotal(),
  }

  const orderRes: AdminListResponse.Order | undefined = res.getOrder()
  if (orderRes) {
    const order: IAdminListOutputOrder = {
      by: orderRes.getBy(),
      direction: orderRes.getDirection(),
    }

    output.order = order
  }

  return output
}
