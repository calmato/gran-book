export interface IUserResponse {
  id: string
  username: string
  email: string
  phoneNumber: string
  thumbnailUrl: string
  selfIntroduction: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  createdAt: string
  updatedAt: string
}

export interface IUserListResponse {
  users: Array<IUserListResponseUser>
  limit: number
  offset: number
  total: number
  order?: IUserListResponseOrder
}

export interface IUserListResponseUser {
  id: string
  username: string
  email: string
  phoneNumber: string
  thumbnailUrl: string
  selfIntroduction: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  createdAt: string
  updatedAt: string
}

export interface IUserListResponseOrder {
  by: string
  direction: string
}
