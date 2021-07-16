export interface IFollowListResponse {
  users: Array<IFollowListResponseUser>
  limit: number
  offset: number
  total: number
}

export interface IFollowListResponseUser {
  id: string
  username: string
  thumbnailUrl: string
  selfIntroduction: string
  isFollow: boolean
}

export interface IFollowerListResponse {
  users: Array<IFollowerListResponseUser>
  limit: number
  offset: number
  total: number
}

export interface IFollowerListResponseUser {
  id: string
  username: string
  thumbnailUrl: string
  selfIntroduction: string
  isFollow: boolean
}

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
  authors: Array<string>
}
