export interface ICreateAuthRequest {
  username: string
  email: string
  password: string
  passwordConfirmation: string
}

export interface IUpdateAuthEmailRequest {
  email: string
}

export interface IUpdateAuthPasswordRequest {
  password: string
  passwordConfirmation: string
}

export interface IUpdateAuthProfileRequest {
  username: string
  gender: number
  thumbnailUrl: string
  selfIntroduction: string
}

export interface IUpdateAuthAddressRequest {
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

export interface IRegisterAuthDeviceRequest {
  instanceId: string
}
