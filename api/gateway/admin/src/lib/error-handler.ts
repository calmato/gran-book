import { Request, Response, NextFunction } from 'express'
import { getHttpError } from '~/lib/http-exception'
import { GrpcError, HttpError } from '~/types/exception'
import { IErrorResponse } from '~/types/response'

export function errorHandler(err: Error, _req: Request, res: Response, next: NextFunction) {
  if (!err) {
    return next()
  }

  // switchだと型エラーなるのでいったんifで書く
  if (err instanceof HttpError) {
    const response: IErrorResponse = {
      status: err.status,
      code: 0,
      message: err.message,
      errors: err.details || [],
    }

    return res.status(response.status).json(response)
  } else if (err instanceof GrpcError) {
    const httpError: HttpError = getHttpError(err)
    const response: IErrorResponse = {
      status: httpError.status,
      code: err.status,
      message: httpError.message,
      errors: httpError.details || [],
    }

    return res.status(response.status).json(response)
  } else {
    const httpError: HttpError = getHttpError(err)
    const response: IErrorResponse = {
      status: httpError.status,
      code: 0,
      message: httpError.message,
      errors: httpError.details || [],
    }

    return res.status(response.status).json(response)
  }
}
