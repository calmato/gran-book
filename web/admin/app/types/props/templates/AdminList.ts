export interface IAdminTableHeader {
  text: string
  value: string
  sortable: boolean
}

export interface IAdminTableFooter {
  itemsPerPageOptions: number[]
}

export interface IAdminTableContent {
  id: string
  name: string
  email: string
  phoneNumber: string
  thumbnailUrl: string
  role: number
}
