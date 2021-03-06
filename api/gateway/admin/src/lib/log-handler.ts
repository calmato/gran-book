/* eslint-disable @typescript-eslint/no-unsafe-assignment */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */
/* eslint-disable @typescript-eslint/no-unsafe-return */
import { Request, Response, NextFunction } from 'express'
import dayjs from 'dayjs'
import logger from '~/plugins/logger'
import { HttpError } from '~/types/exception'

const filterWords: string[] = ['password', 'passwordConfirmation', 'authorization', 'thumbnail']

function logFilter(body: any): any {
  const obj: typeof body[keyof string] = {}

  Object.keys(body).forEach((key: string) => {
    obj[key] = filterWords.includes(key) ? '<FILTERED>' : body[key]
  })

  return obj
}

export function accessLogHandler(req: Request, _: Response, next: NextFunction): void {
  const logs: any = {
    direction: 'request',
    path: req.path,
    method: req.method,
    remoteIp: req.ip,
    requestHeader: logFilter(req.headers),
    requestQuery: logFilter(req.query),
    requestBody: logFilter(req.body),
    time: dayjs().format(),
  }

  logger.default.info(JSON.stringify(logs))
  logger.access.info(JSON.stringify(logs))
  next()
}

export function errorLogHandler(req: Request, err: HttpError): void {
  const logs: any = {
    direction: 'response',
    path: req.path,
    method: req.method,
    remoteIp: req.ip,
    status: err.status,
    message: err.message,
    errors: err.details,
    time: dayjs().format(),
  }

  logger.default.info(JSON.stringify(logs))
  logger.access.info(JSON.stringify(logs))
}
/* eslint-enable */
