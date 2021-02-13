import { Request, Response, NextFunction } from 'express'
import { auth } from '~/plugins/firebase'
import { unauthorized, forbidden } from '~/lib/http-exception'
import { getAuth } from '~/api/user'
import { IAuthOutput } from '~/types/output'

interface IRoute {
  path: string
  method: string
}

const excludeRoutes: IRoute[] = [{ path: '/health', method: 'GET' }]
const developerForbiddenRoutes: IRoute[] = [{ path: '/v1/auth', method: 'POST' }]
const operatorForbiddenRoutes: IRoute[] = [{ path: '/v1/auth', method: 'POST' }]

function isExcludeRoute(req: Request): boolean {
  if (req.method === 'OPTIONS') {
    return true
  }

  return excludeRoutes.some((r: IRoute) => r.path === req.path && r.method === req.method)
}

function isForbiddenRoute(req: Request, role: number): boolean {
  switch (role) {
    case 1: // 1 -> Administrator
      return false
    case 2: // 2 -> Developer
      return developerForbiddenRoutes.some((r: IRoute) => r.path === req.path && r.method === req.method)
    case 3: // 3 -> Operator
      return operatorForbiddenRoutes.some((r: IRoute) => r.path === req.path && r.method === req.method)
    default:
      return true
  }
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

  await auth
    .verifyIdToken(token)
    .then(() => next())
    .catch(() => next(unauthorized()))
}

export async function authorization(req: Request, _: Response, next: NextFunction): Promise<any> {
  if (isExcludeRoute(req)) {
    return next()
  }

  const token: string = getToken(req)
  if (token === '') {
    return next(unauthorized())
  }

  await getAuth(req)
    .then((res: IAuthOutput) => {
      if (![1, 2, 3].includes(res.role)) {
        return next(forbidden())
      }

      if (isForbiddenRoute(req, res.role)) {
        return next(forbidden())
      }

      next()
    })
    .catch(() => next(unauthorized()))
}
