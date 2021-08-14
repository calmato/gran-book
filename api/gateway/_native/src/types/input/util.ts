import { OrderBy } from '~/util'

export interface ISearchInput {
  field: string
  value: string
}

export interface IOrderInput {
  field: string
  orderBy: OrderBy
}
