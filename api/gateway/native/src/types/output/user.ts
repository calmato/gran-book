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
