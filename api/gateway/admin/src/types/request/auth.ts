export interface ICreateAuthRequest {
  username: string
  email: string
  password: string
  passwordConfirmation: string
}

export interface IUpdateAuthRequest {
  username: string
  gender: number
  thumbnail: string
  selfIntroduction: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  phoneNumber: string
}
