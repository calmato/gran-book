export interface IUserOutput {
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

export interface IUserListOutput {
  users: Array<IUserListOutputUser>
  limit: number
  offset: number
  total: number
  order?: IUserListOutputOrder
}

export interface IUserListOutputUser {
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

export interface IUserListOutputOrder {
  by: string
  direction: string
}

export interface IUserThumbnailOutput {
  thumbnailUrl: string
}
