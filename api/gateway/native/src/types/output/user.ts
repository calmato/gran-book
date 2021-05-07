export interface IUserHashOutput {
  [key: string]: IUserHashOutputUser
}

export interface IUserHashOutputUser {
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

export interface IUserThumbnailOutput {
  thumbnailUrl: string
}

export interface IFollowListOutput {
  users: Array<IFollowListOutputUser>
  limit: number
  offset: number
  total: number
}

export interface IFollowListOutputUser {
  id: string
  username: string
  thumbnailUrl: string
  selfIntroduction: string
  isFollow: boolean
}

export interface IFollowerListOutput {
  users: Array<IFollowerListOutputUser>
  limit: number
  offset: number
  total: number
}

export interface IFollowerListOutputUser {
  id: string
  username: string
  thumbnailUrl: string
  selfIntroduction: string
  isFollow: boolean
}

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

export interface IUserProfileOutput {
  id: string
  username: string
  thumbnailUrl: string
  selfIntroduction: string
  isFollow: boolean
  isFollower: boolean
  followCount: number
  followerCount: number
}
