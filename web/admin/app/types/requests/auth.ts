export interface IAuthUpdateEmailRequest {
  email: string
}

export interface IAuthUpdatePasswordRequest {
  password: string
  passwordConfirmation: string
}

export interface IAuthUpdateProfileRequest {
  username: string
  thumbnail: string
  selfIntroduction: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  phoneNumber: string
}
