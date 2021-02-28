export interface IListAdminInput {
  limit: number
  offset: number
  by: string
  direction: string
}

export interface ICreateAdminInput {
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
  id: string
  role: number
}

export interface IUpdateAdminPasswordInput {
  id: string
  password: string
  passwordConfirmation: string
}

export interface IUpdateAdminProfileInput {
  id: string
  username: string
  email: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
}
