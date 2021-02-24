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

export interface IUpdateAdminRoleRequest {
  role: number
}

export interface IUpdateAdminPasswordRequest {
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
