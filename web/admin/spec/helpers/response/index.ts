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
    ...AuthStore.get,
  },
  post: {},
  patch: {
    ...AuthStore.patch,
  },
  put: {},
  delete: {},
  error: err,
}
