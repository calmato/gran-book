/* eslint-disable @typescript-eslint/no-unsafe-assignment */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */
/* eslint-disable @typescript-eslint/no-unsafe-return */
import { Request, Response, NextFunction } from 'express'
import logger from '~/plugins/logger'
import { HttpError } from '~/types/exception'

const filterWords: string[] = ['password', 'passwordConfirmation', 'authorization']

function logFilter(body: any): any {
  const obj: typeof body[keyof string] = {}

  Object.keys(body).forEach((key: string) => {
    obj[key] = filterWords.includes(key) ? '<FILTERED>' : body[key]
  })

  return obj
}

export function accessLogHandler(req: Request, _: Response, next: NextFunction): void {
  try {
    const logs: any = {
      direction: 'request',
      path: req.path,
      method: req.method,
      remoteIp: req.ip,
      requestHeader: logFilter(req.headers),
      requestQuery: logFilter(req.query),
      requestBody: logFilter(req.body),
    }

    logger.default.info(JSON.stringify(logs))
    logger.access.info(JSON.stringify(logs))
    console.log('debug: request', 'passed')
  } catch (err) {
    console.log('debug: request', err)
  }
  next()
}

export function errorLogHandler(req: Request, err: HttpError): void {
  console.log('debug: response', err)

  const logs: any = {
    direction: 'response',
    path: req.path,
    method: req.method,
    remoteIp: req.ip,
    status: err.status,
    message: err.message,
    errors: err.details,
  }

  logger.default.info(JSON.stringify(logs))
  logger.access.info(JSON.stringify(logs))
}
/* eslint-enable */
