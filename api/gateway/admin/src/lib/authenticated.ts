import { Request, Response, NextFunction } from 'express'
import { auth } from '~/plugins/firebase'

interface IRoute {
  path: string
  method: string
}

const openRoutes: IRoute[] = [
  { path: '/health', method: 'GET' },
  { path: '/v1/auth', method: 'POST' },
]

function isOpenRoute(req: Request): boolean {
  return openRoutes.some((r) => r.path === req.path && r.method === req.method)
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
  if (isOpenRoute(req)) {
    return next()
  }

  const token: string = getToken(req)
  if (token === '') {
    return res.status(401).json({ message: 'debug: token' })
  }

  await auth.verifyIdToken(token)
    .then(() => next())
    .catch((err: Error) => {
      // TODO: refactor
      console.log('debug', err)
      return res.status(401).json({ message: 'debug: request' })
    })
}
