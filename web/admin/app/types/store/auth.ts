export interface IAuthState {
  id: string
  email: string
  emailVerified: boolean
  token: string
  username: string
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

export interface IAuthProfile {
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
