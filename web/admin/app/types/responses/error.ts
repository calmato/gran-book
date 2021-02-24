import { IErrorDetail } from '~/types/exception'

export interface IErrorData {
  status: number
  code: number
  message: string
  errors: IErrorDetail[]
}

export interface IErrorResponse {
  status: number
  data: IErrorData
}
