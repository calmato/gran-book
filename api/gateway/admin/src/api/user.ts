import { credentials } from '@grpc/grpc-js'
import { AuthServiceClient, IAuthServiceClient } from '~/proto/user_apiv1_grpc_pb'
import { CreateAuthRequest, AuthResponse } from '~/proto/user_apiv1_pb'
import { ICreateAuthInput } from '~/types/input'
import { IAuthOutput } from '~/types/output'

// TODO: 環境変数から取得するように
const userAPIURL: string = 'user_api:8080'

export function createAuth(payload: ICreateAuthInput): Promise<IAuthOutput> {
  const req = new CreateAuthRequest()

  // TODO: 共通化
  const client: IAuthServiceClient = new AuthServiceClient(
    userAPIURL,
    credentials.createInsecure(),
  )

  req.setEmail(payload.email)
  req.setEmail(payload.username)
  req.setEmail(payload.password)
  req.setEmail(payload.passwordConfirmation)

  return new Promise((resolve: (res: IAuthOutput) => void, reject: (reason: Error) => void) => {
    client.createAuth(req, (err: any, res: AuthResponse) => {
      if (err) {
        reject(err)
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
