export interface IAdminOutput {
  id: string
  username: string
  email: string
  phoneNumber: string
  role: number
  thumbnailUrl: string
  selfIntroduction: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  createdAt: string
  updatedAt: string
}

export interface IAdminListOutput {
  users: Array<IAdminListOutputUser>
  limit: number
  offset: number
  total: number
  order?: IAdminListOutputOrder
}

export interface IAdminListOutputUser {
  id: string
  username: string
  email: string
  phoneNumber: string
  role: number
  thumbnailUrl: string
  selfIntroduction: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  createdAt: string
  updatedAt: string
}

export interface IAdminListOutputOrder {
  by: string
  direction: string
}
