import { Request } from 'express'
import fs from 'fs'
import { ClientWritableStream } from '@grpc/grpc-js'
import { authClient } from '~/plugins/grpc'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import {
  EmptyUser,
  AuthResponse,
  AuthThumbnailResponse,
  UpdateAuthAddressRequest,
  UpdateAuthEmailRequest,
  UpdateAuthPasswordRequest,
  UpdateAuthProfileRequest,
  UploadAuthThumbnailRequest,
} from '~/proto/user_apiv1_pb'
import {
  IUpdateAuthAddressInput,
  IUpdateAuthEmailInput,
  IUpdateAuthPasswordInput,
  IUpdateAuthProfileInput,
  IUploadAuthThumbnailInput,
} from '~/types/input'
import { IAuthOutput, IAuthThumbnailOutput } from '~/types/output'

export function getAuth(req: Request<any>): Promise<IAuthOutput> {
  const request = new EmptyUser()
  const metadata = getGrpcMetadata(req)

  return new Promise((resolve: (res: IAuthOutput) => void, reject: (reason: Error) => void) => {
    authClient.getAuth(request, metadata, (err: any, res: AuthResponse) => {
      if (err) {
        console.log('debug', err)
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

export function updateAuthPassword(req: Request<any>, input: IUpdateAuthPasswordInput): Promise<IAuthOutput> {
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
  request.setThumbnailUrl(input.thumbnailUrl)
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

export function uploadAuthThumbnail(
  req: Request<any>,
  input: IUploadAuthThumbnailInput
): Promise<IAuthThumbnailOutput> {
  const metadata = getGrpcMetadata(req)

  return new Promise((resolve: (res: IAuthThumbnailOutput) => void, reject: (reason: Error) => void) => {
    const call: ClientWritableStream<UploadAuthThumbnailRequest> = authClient.uploadAuthThumbnail(
      metadata,
      (err: any, res: AuthThumbnailResponse) => {
        if (err) {
          return reject(getGrpcError(err))
        }

        const output: IAuthThumbnailOutput = {
          thumbnailUrl: res.getThumbnailUrl(),
        }

        return resolve(output)
      }
    )

    const stream: fs.ReadStream = fs.createReadStream(input.path, { highWaterMark: 102400 })
    let count = 0 // 読み込み回数

    stream.on('data', (chunk: Buffer) => {
      const request = new UploadAuthThumbnailRequest()
      request.setThumbnail(chunk)
      request.setPosition(count)

      call.write(request)
      count += 1
    })

    stream.on('end', () => {
      call.end() // TODO: try-catchとかのエラー処理必要かも
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
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
  }

  return output
}
