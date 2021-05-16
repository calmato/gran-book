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

export interface IAdminUpdateRequest {
  email: string
  phoneNumber: string
  role: number
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  thumbnailUrl: string
}
