export interface ICreateAdminRequest {
  username: string
  email: string
  password: string
  passwordConfirmation: string
  role: number
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
}

export interface IUpdateAdminContactRequest {
  email: string
  phoneNumber: string
}

export interface IUpdateAdminPasswordRequest {
  password: string
  passwordConfirmation: string
}

export interface IUpdateAdminProfileRequest {
  role: number
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  thumbnailUrl: string
}
