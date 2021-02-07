import { IErrorOutput } from '~/types/output'
import { IErrorResponse } from '~/types/response'

const statusConverter: { [key: number]: number } = {
  0: 200, // OK -> OK
  2: 500, // UNKNOWN -> Internal Server Error
  3: 400, // INVALID_ARGUMENT -> Bad Request
  5: 404, // NOT_FOUND -> Not Found
  6: 409, // ALREADY_EXISTS -> Conflict
  7: 403, // PREMISSION_DENIED -> Forbidden
  13: 500, // INTERNAL -> Internal Server Error
  14: 503, // UNAVAILABLE -> Service Unavailable
  16: 401, // UNAUTHENTICATED -> Unauthorized
}

function convertStatusGrpcToHttp(grpcStatus: number): number {
  const httpStatus = statusConverter[grpcStatus]
  if (httpStatus) {
    return httpStatus
  }

  return 500
}

export function getHttpError(output: IErrorOutput): IErrorResponse {
  const status: number = convertStatusGrpcToHttp(output.status)

  const res: IErrorResponse = {
    status: status,
    code: output.status,
    message: output.message,
    details: output.details,
  }

  return res
}
