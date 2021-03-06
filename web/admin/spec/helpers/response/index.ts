import * as AdminStore from './admin'
import * as AuthStore from './auth'

const err = {
  response: {
    status: 400,
    data: {
      status: 400,
      code: 0,
      message: 'api error',
      errors: [],
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
