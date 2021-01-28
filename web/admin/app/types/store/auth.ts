export interface IAuthState {
  id: string
  email: string
  emailVerified: boolean
  token: string
  username: string
  gender: number
  phoneNumber: string
  role: number
  thumbnailUrl: string
  selfIntroduction: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  postalCode: string
  prefecture: string
  city: string
  addressLine1: string
  addressLine2: string
  createdAt: string
  updatedAt: string
}

export interface IAuthProfile {
  id: string
  username: string
  gender: number
  email: string
  phoneNumber: string
  role: number
  thumbnailUrl: string
  selfIntroduction: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  postalCode: string
  prefecture: string
  city: string
  addressLine1: string
  addressLine2: string
  createdAt: string
  updatedAt: string
}
