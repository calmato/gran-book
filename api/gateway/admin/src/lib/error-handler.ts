import { Request, Response, NextFunction } from 'express'
import { errorLogHandler } from '~/lib/log-handler'
import { getHttpError, notFound } from '~/lib/http-exception'
import { GrpcError, HttpError } from '~/types/exception'
import { IErrorResponse } from '~/types/response'

export function notFoundErrorHandler(req: Request, res: Response, _: NextFunction): void {
  const err: HttpError = notFound()
  const response: IErrorResponse = setErrorResponse(err)

  errorLogHandler(req, err)
  res.status(response.status).json(response)
}

export function otherErrorHandler(err: Error, req: Request, res: Response, _: NextFunction): void {
  let grpcError: GrpcError | undefined = undefined
  if (err instanceof GrpcError) {
    grpcError = err
  }

  const httpError: HttpError = getHttpError(err)
  const response: IErrorResponse = setErrorResponse(httpError, grpcError)

  errorLogHandler(req, httpError)
  res.status(response.status).json(response)
}

function setErrorResponse(httpError: HttpError, grpcError?: GrpcError): IErrorResponse {
  const response: IErrorResponse = {
    status: httpError.status,
    code: grpcError?.status || 0,
    message: httpError.message,
    errors: httpError.details || [],
  }

  return response
}
