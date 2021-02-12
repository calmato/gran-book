import { Request, Response, NextFunction } from 'express'
import { auth } from '~/plugins/firebase'
import { unauthorized } from '~/lib/http-exception'

interface IRoute {
  path: string
  method: string
}

const excludeRoutes: IRoute[] = [
  { path: '/health', method: 'GET' },
  { path: '/v1/auth', method: 'POST' },
]

function isExcludeRoute(req: Request): boolean {
  if (req.method === 'OPTIONS') {
    return true
  }

  return excludeRoutes.some((r) => r.path === req.path && r.method === req.method)
}

function getToken(req: Request): string {
  const { authorization } = req.headers
  if (!authorization) {
    return ''
  }

  const arr: string[] = authorization.split(' ')
  if (arr.length === 2 && arr[0] === 'Bearer') {
    return arr[1]
  }

  return ''
}

export async function authentication(req: Request, res: Response, next: NextFunction): Promise<any> {
  if (isExcludeRoute(req)) {
    return next()
  }

  const token: string = getToken(req)
  if (token === '') {
    return next(unauthorized())
  }

  await auth.verifyIdToken(token)
    .then(() => next())
    .catch(() => next(unauthorized()))
}
