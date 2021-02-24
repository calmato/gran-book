import { Request } from 'express'
import { adminClient } from '~/plugins/grpc'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import {
  AdminResponse,
  CreateAdminRequest,
  UpdateAdminPasswordRequest,
  UpdateAdminProfileRequest,
  UpdateAdminRoleRequest,
} from '~/proto/user_apiv1_pb'
import {
  ICreateAdminInput,
  IUpdateAdminPasswordInput,
  IUpdateAdminProfileInput,
  IUpdateAdminRoleInput,
} from '~/types/input'
import { IAdminOutput } from '~/types/output'

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
