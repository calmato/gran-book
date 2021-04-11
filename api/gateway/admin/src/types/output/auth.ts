export interface IAuthOutput {
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
  activated: boolean
  createdAt: string
  updatedAt: string
}

export interface IAuthThumbnailOutput {
  thumbnailUrl: string
}
