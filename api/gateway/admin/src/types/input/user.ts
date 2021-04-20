export interface IListUserInput {
  limit: number
  offset: number
  by: string
  direction: string
}

export interface ISearchUserInput {
  limit: number
  offset: number
  by: string
  direction: string
  field: string
  value: string
}

export interface IGetUserInput {
  id: string
}
