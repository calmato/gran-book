import { authClient } from '~/plugins/grpc'
import { getGrpcError } from '~/lib/grpc-exception'
import { CreateAuthRequest, AuthResponse } from '~/proto/user_apiv1_pb'
import { ICreateAuthInput } from '~/types/input'
import { IAuthOutput } from '~/types/output'

export function createAuth(input: ICreateAuthInput): Promise<IAuthOutput> {
  const req = new CreateAuthRequest()

  req.setUsername(input.username)
  req.setEmail(input.email)
  req.setPassword(input.password)
  req.setPasswordConfirmation(input.passwordConfirmation)

  return new Promise((resolve: (res: IAuthOutput) => void, reject: (reason: Error) => void) => {
    authClient.createAuth(req, (err: any, res: AuthResponse) => {
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
