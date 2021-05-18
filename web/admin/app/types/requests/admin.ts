export interface IAdminCreateRequest {
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

export interface IAdminUpdateProfileRequest {
  email: string
  phoneNumber: string
  role: number
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  thumbnailUrl: string
}

export interface IAdminUpdateContactRequest {
  email: string
  phoneNumber: string
}

export interface IAdminUpdatePasswordRequest {
  password: string
  passwordConfirmation: string
}
