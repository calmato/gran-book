import { IOrderInput, ISearchInput } from './util'

export interface IListUserInput {
  search?: ISearchInput
  order?: IOrderInput
  limit: number
  offset: number
}

export interface IListFollowInput {
  userId: string
  order?: IOrderInput
  limit: number
  offset: number
}

export interface IListFollowerInput {
  userId: string
  order?: IOrderInput
  limit: number
  offset: number
}

export interface IMultiGetUserInput {
  userIds: Array<string>
}

export interface IGetUserInput {
  userId: string
}

export interface IGetUserProfileInput {
  userId: string
}

export interface IFollowInput {
  userId: string
  followerId: string
}

export interface IUnfollowInput {
  userId: string
  followerId: string
}
