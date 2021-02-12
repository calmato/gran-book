import { Request } from 'express'
import { authClient } from '~/plugins/grpc'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import { EmptyUser, CreateAuthRequest, AuthResponse } from '~/proto/user_apiv1_pb'
import { ICreateAuthInput } from '~/types/input'
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
        createdAt: res.getCreatedAt(),
        updatedAt: res.getUpdatedAt(),
      }

      resolve(output)
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
        createdAt: res.getCreatedAt(),
        updatedAt: res.getUpdatedAt(),
      }

      resolve(output)
    })
  })
}
