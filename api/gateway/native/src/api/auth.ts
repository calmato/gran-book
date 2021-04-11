import { Request } from 'express'
import { authClient } from '~/plugins/grpc'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import {
  EmptyUser,
  CreateAuthRequest,
  AuthResponse,
  UpdateAuthAddressRequest,
  UpdateAuthEmailRequest,
  UpdateAuthPasswordRequest,
  UpdateAuthProfileRequest,
} from '~/proto/user_apiv1_pb'
import {
  ICreateAuthInput,
  IUpdateAuthAddressInput,
  IUpdateAuthEmailInput,
  IUpdateAuthPasswordInput,
  IUpdateAuthProfileInput,
} from '~/types/input'
import { IAuthOutput } from '~/types/output'

export function getAuth(req: Request<any>): Promise<IAuthOutput> {
  const request = new EmptyUser()
  const metadata = getGrpcMetadata(req)

  return new Promise((resolve: (res: IAuthOutput) => void, reject: (reason: Error) => void) => {
    authClient.getAuth(request, metadata, (err: any, res: AuthResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setAuthOutput(res))
    })
  })
}

export function createAuth(_: Request<any>, input: ICreateAuthInput): Promise<IAuthOutput> {
  const request = new CreateAuthRequest()

  request.setUsername(input.username)
  request.setEmail(input.email)
  request.setPassword(input.password)
  request.setPasswordConfirmation(input.passwordConfirmation)

  return new Promise((resolve: (res: IAuthOutput) => void, reject: (reason: Error) => void) => {
    authClient.createAuth(request, (err: any, res: AuthResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setAuthOutput(res))
    })
  })
}

export function updateAuthEmail(req: Request<any>, input: IUpdateAuthEmailInput): Promise<IAuthOutput> {
  const request = new UpdateAuthEmailRequest()
  const metadata = getGrpcMetadata(req)

  request.setEmail(input.email)

  return new Promise((resolve: (res: IAuthOutput) => void, reject: (reason: Error) => void) => {
    authClient.updateAuthEmail(request, metadata, (err: any, res: AuthResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setAuthOutput(res))
    })
  })
}

export function UpdateAuthPassword(req: Request<any>, input: IUpdateAuthPasswordInput): Promise<IAuthOutput> {
  const request = new UpdateAuthPasswordRequest()
  const metadata = getGrpcMetadata(req)

  request.setPassword(input.password)
  request.setPasswordConfirmation(input.passwordConfirmation)

  return new Promise((resolve: (res: IAuthOutput) => void, reject: (reason: Error) => void) => {
    authClient.updateAuthPassword(request, metadata, (err: any, res: AuthResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setAuthOutput(res))
    })
  })
}

export function updateAuthProfile(req: Request<any>, input: IUpdateAuthProfileInput): Promise<IAuthOutput> {
  const request = new UpdateAuthProfileRequest()
  const metadata = getGrpcMetadata(req)

  request.setUsername(input.username)
  request.setGender(input.gender)
  request.setThumbnail(input.thumbnail)
  request.setSelfIntroduction(input.selfIntroduction)

  return new Promise((resolve: (res: IAuthOutput) => void, reject: (reason: Error) => void) => {
    authClient.updateAuthProfile(request, metadata, (err: any, res: AuthResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setAuthOutput(res))
    })
  })
}

export function updateAuthAddress(req: Request<any>, input: IUpdateAuthAddressInput): Promise<IAuthOutput> {
  const request = new UpdateAuthAddressRequest()
  const metadata = getGrpcMetadata(req)

  request.setLastName(input.lastName)
  request.setFirstName(input.firstName)
  request.setLastNameKana(input.lastNameKana)
  request.setFirstNameKana(input.firstNameKana)
  request.setPhoneNumber(input.phoneNumber)
  request.setPostalCode(input.postalCode)
  request.setPrefecture(input.prefecture)
  request.setCity(input.city)
  request.setAddressLine1(input.addressLine1)
  request.setAddressLine2(input.addressLine2)

  return new Promise((resolve: (res: IAuthOutput) => void, reject: (reason: Error) => void) => {
    authClient.updateAuthAddress(request, metadata, (err: any, res: AuthResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setAuthOutput(res))
    })
  })
}

export function deleteAuth(req: Request<any>): Promise<void> {
  const request = new EmptyUser()
  const metadata = getGrpcMetadata(req)

  return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
    authClient.deleteAuth(request, metadata, (err: any) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve()
    })
  })
}

function setAuthOutput(res: AuthResponse): IAuthOutput {
  const output: IAuthOutput = {
    id: res.getId(),
    username: res.getUsername(),
    gender: res.getGender(),
    email: res.getEmail(),
    phoneNumber: res.getPhoneNumber(),
    role: res.getRole(),
    thumbnailUrl: res.getThumbnailUrl(),
    selfIntroduction: res.getSelfIntroduction(),
    lastName: res.getLastName(),
    firstName: res.getFirstName(),
    lastNameKana: res.getLastNameKana(),
    firstNameKana: res.getFirstNameKana(),
    postalCode: res.getPostalCode(),
    prefecture: res.getPrefecture(),
    city: res.getCity(),
    addressLine1: res.getAddressLine1(),
    addressLine2: res.getAddressLine2(),
    activated: res.getActivated(),
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
  }

  return output
}
