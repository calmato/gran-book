export interface IErrorResponse {
  status: number
  code: number
  message: string
  errors: Array<any>
}
