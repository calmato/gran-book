import { Request, Response, NextFunction } from 'express'
import { auth } from '~/plugins/firebase'
import { unauthorized, forbidden } from '~/lib/http-exception'
import { getAuth } from '~/api/auth'
import { IAuthOutput } from '~/types/output'

interface IRoute {
  pathRegex: RegExp
  method: string
}

const excludeRoutes: IRoute[] = [
  { pathRegex: /^\/health$/g, method: 'GET' }, // GET - /v/health
]
const developerForbiddenRoutes: IRoute[] = [
  { pathRegex: /^\/v1\/admin$/g, method: 'POST' }, // POST - /v1/admin
  { pathRegex: /^\/v1\/admin\/[^/].*\/role$/g, method: 'PATCH' }, // PATCH - /v1/admin/{userId}/role
  { pathRegex: /^\/v1\/admin\/[^/].*\/password$/g, method: 'PATCH' }, // PATCH - /v1/admin/{userId}/password
  { pathRegex: /^\/v1\/admin\/[^/].*\/profile$/g, method: 'PATCH' }, // PATCH - /v1/admin/{userId}/profile
]
const operatorForbiddenRoutes: IRoute[] = [
  { pathRegex: /^\/v1\/admin$/g, method: 'POST' }, // POST - /v1/admin
  { pathRegex: /^\/v1\/admin\/[^/].*\/role$/g, method: 'PATCH' }, // PATCH - /v1/admin/{userId}/role
  { pathRegex: /^\/v1\/admin\/[^/].*\/password$/g, method: 'PATCH' }, // PATCH - /v1/admin/{userId}/password
  { pathRegex: /^\/v1\/admin\/[^/].*\/profile$/g, method: 'PATCH' }, // PATCH - /v1/admin/{userId}/profile
]

function isMatchRoute(req: Request, routes: IRoute[]): boolean {
  return routes.some((r: IRoute) => {
    if (r.method !== req.method) {
      return false
    }

    return r.pathRegex.test(req.path)
  })
}

function isExcludeRoute(req: Request): boolean {
  if (req.method === 'OPTIONS') {
    return true
  }

  return isMatchRoute(req, excludeRoutes)
}

function isForbiddenRoute(req: Request, role: number): boolean {
  switch (role) {
    case 1: // 1 -> Administrator
      return false
    case 2: // 2 -> Developer
      return isMatchRoute(req, developerForbiddenRoutes)
    case 3: // 3 -> Operator
      return isMatchRoute(req, operatorForbiddenRoutes)
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
