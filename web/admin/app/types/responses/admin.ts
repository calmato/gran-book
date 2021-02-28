export interface IAdminListResponse {
  users: Array<IAdminListResponseUser>
  limit: number
  offset: number
  total: number
  order?: IAdminListResponseOrder
}

export interface IAdminListResponseUser {
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

export interface IAdminListResponseOrder {
  by: string
  direction: string
}
