import { Request } from 'express'
import * as grpc from '@grpc/grpc-js'

export function getGrpcMetadata(req: Request): grpc.Metadata {
  const meta = new grpc.Metadata()

  if (req.headers.authorization) {
    meta.set('Authorization', req.headers.authorization)
  }

  const userAgent = req.headers['user-agent']

  meta.set('user-agent', userAgent || 'unknown')
  meta.set('x-forwarded-for', req.ip)

  return meta
}
