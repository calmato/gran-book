import { IErrorOutput } from '~/types/output'
import { IErrorResponse } from '~/types/response'

// TODO: unionに置き換え
function convertStatusGrpcToHttp(status: number): number {
  switch(status) {
    case 0:
      return 200
    case 2:
      return 500
    case 3:
      return 400
    case 5:
      return 404
    case 6:
      return 409
    case 7:
      return 403
    case 13:
      return 500
    case 15:
      return 503
    case 16:
      return 401
    default:
      return 500
  }
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
