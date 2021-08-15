export interface IUpdateAuthEmailInput {
  email: string
}

export interface IUpdateAuthPasswordInput {
  password: string
  passwordConfirmation: string
}

export interface IUpdateAuthProfileInput {
  username: string
  gender: number
  thumbnailUrl: string
  selfIntroduction: string
}

export interface IUpdateAuthAddressInput {
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  phoneNumber: string
  postalCode: string
  prefecture: string
  city: string
  addressLine1: string
  addressLine2: string
}

export interface IUploadAuthThumbnailInput {
  path: string
}
