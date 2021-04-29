export interface IListUserByUserIdsInput {
  ids: Array<string>
}

export interface IListFollowInput {
  id: string
  limit: number
  offset: number
}

export interface IListFollowerInput {
  id: string
  limit: number
  offset: number
}

export interface IGetUserProfileInput {
  id: string
}

export interface IRegisterFollowInput {
  id: string
}

export interface IUnregisterFollowInput {
  id: string
}
