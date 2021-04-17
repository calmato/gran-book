export interface INotificationTableHeader {
  text: string
  value: string
  sortable: boolean
}

export interface INotificationTableContent {
  title: string
  description: string
  timestmap: string
  category: number
  importance: number
}
