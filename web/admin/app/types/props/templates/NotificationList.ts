export interface INotificationTableHeader {
  text: string
  value: string
  sortable: boolean
}

export interface INotificationTableContent {
  title: string
  description: string
  timestamp: string
  category: number
  importance: number
}
