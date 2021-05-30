import { Request } from 'express'
import * as grpc from '@grpc/grpc-js'

export function getGrpcMetadata(req: Request): grpc.Metadata {
  const meta = new grpc.Metadata()

  if (req.headers.authorization) {
    meta.set('Authorization', req.headers.authorization)
  }

  const requestId: string | string[] = req.headers['x-request-id'] || 'unknown'
  if (Array.isArray(requestId)) {
    meta.set('x-request-id', requestId.join(','))
  } else {
    meta.set('x-request-id', requestId)
  }

  const userAgent: string = req.headers['user-agent'] || 'unknown'

  meta.set('user-agent', userAgent)
  meta.set('x-forwarded-for', req.ip)

  return meta
}
