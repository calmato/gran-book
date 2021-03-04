export interface IUserProfileResponse {
  id: string
  username: string
  thumbnailUrl: string
  selfIntroduction: string
  isFollow: boolean
  isFollower: boolean
  followCount: number
  followerCount: number
  reviewCount: number
  rating: number
  products: Array<IUserProfileResponseProduct>
}

export interface IUserProfileResponseProduct {
  id: number
  name: string
  thumbnail: string
  authors: Array<IUserProfileResponseAuthor>
}

export interface IUserProfileResponseAuthor {
  id: number
  name: string
}
