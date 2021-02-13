export interface ICreateAuthInput {
  username: string
  email: string
  password: string
  passwordConfirmation: string
}

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
  thumbnail: string
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
