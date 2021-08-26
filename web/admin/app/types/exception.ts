import { ErrorResponse } from '~/types/api/common_pb'

export interface IErrorResponse {
  data: ErrorResponse.AsObject
  status: number
}

export class ApiError extends Error {
  constructor(public status: number, public message: string, public data?: ErrorResponse.AsObject) {
    super(message)
  }
}
