export interface IAdminState {
  users: IAdminUser[]
  total: number
}

export interface IAdminUser {
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
