export interface IErrorData {
  status: number
  code: number
  message: string
  errors: IErrorDetail[]
}

export interface IErrorDetail {
  field: string
  reason: string
}

export class ApiError extends Error {
  constructor(public status: number, public message: string, public data?: IErrorData) {
    super(message)
  }
}
