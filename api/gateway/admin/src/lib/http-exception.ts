import { GrpcError, HttpError } from '~/types/exception'

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

export function badRequest(details: any): HttpError {
  const message: string = 'Bad Request'
  return new HttpError(400, message, details)
}

export function unauthorized(): HttpError {
  const message: string = 'Unauthorized'
  return new HttpError(401, message)
}

export function forbidden(): HttpError {
  const message: string = 'Forbidden'
  return new HttpError(403, message)
}

export function notFound(): HttpError {
  const message: string = 'Not Found'
  return new HttpError(403, message)
}

export function alreadyExists(details: any): HttpError {
  const message: string = 'Conflict'
  return new HttpError(409, message, details)
}

export function serverError(details: any): HttpError {
  const message: string = 'Internal Server Error'
  return new HttpError(500, message, details)
}

export function getHttpError(err: Error): HttpError {
  if (err instanceof GrpcError) {
    const status: number = convertStatusGrpcToHttp(err.status)
    switch (status) {
    case 400:
      return badRequest(err.details)
    case 401:
      return unauthorized()
    case 403:
      return forbidden()
    case 404:
      return notFound()
    case 409:
      return alreadyExists(err.details)
    default:
      return serverError(err.details)
    }
  } else {
    // TODO: refactor
    return serverError(err)
  }
}
