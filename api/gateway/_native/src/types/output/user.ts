export interface IUserOutput {
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
  createdAt: string
  updatedAt: string
}

export interface IUserListOutput {
  users: Array<IUserListOutputUser>
  limit: number
  offset: number
  total: number
}

export interface IUserListOutputUser {
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
  createdAt: string
  updatedAt: string
}

export interface IUserMapOutput {
  [key: string]: IUserMapOutputUser
}

export interface IUserMapOutputUser {
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

export interface IFollowListOutput {
  follows: Array<IFollowListOutputFollow>
  limit: number
  offset: number
  total: number
}

export interface IFollowListOutputFollow {
  id: string
  username: string
  thumbnailUrl: string
  selfIntroduction: string
  isFollow: boolean
}

export interface IFollowerListOutput {
  followers: Array<IFollowerListOutputFollower>
  limit: number
  offset: number
  total: number
}

export interface IFollowerListOutputFollower {
  id: string
  username: string
  thumbnailUrl: string
  selfIntroduction: string
  isFollow: boolean
}
