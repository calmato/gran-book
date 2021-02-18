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

export interface IUpdateAdminRoleInput {
  role: number
}

export interface IUpdateAdminPasswordInput {
  password: string
  passwordConfirmation: string
}

export interface IUpdateAdminProfileRequest {
  username: string
  email: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
}
