export interface IAdminListForm {
  limit: number
  offset: number
  order: IAdminListFormOrder
}

export interface IAdminListFormOrder {
  by: string
  desc: boolean
}
