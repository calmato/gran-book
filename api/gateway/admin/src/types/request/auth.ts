export interface IUpdateAuthEmailRequest {
  email: string
}

export interface IUpdateAuthPasswordRequest {
  password: string
  passwordConfirmation: string
}

export interface IUpdateAuthRequest {
  username: string
  thumbnailUrl: string
  selfIntroduction: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  phoneNumber: string
}
