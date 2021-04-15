import { Request } from 'express'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import { userClient } from '~/plugins/grpc'
import { GetUserRequest, UserResponse } from '~/proto/user_apiv1_pb'
import { IGetUserInput } from '~/types/input'
import { IUserOutput } from '~/types/output'

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
