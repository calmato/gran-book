import * as AdminStore from './admin'
import * as AuthStore from './auth'
import { IErrorResponse } from '~/types/exception'

const err: { response: IErrorResponse } = {
  response: {
    status: 400,
    data: {
      status: 400,
      code: 0,
      message: 'api error',
      detail: 'some error',
    },
  },
}

export default {
  get: {
    ...AdminStore.get,
    ...AuthStore.get,
  },
  post: {
    ...AdminStore.post,
    ...AuthStore.post,
  },
  patch: {
    ...AdminStore.patch,
    ...AuthStore.patch,
  },
  put: {},
  delete: {
    ...AdminStore.destroy,
  },
  error: err,
}
