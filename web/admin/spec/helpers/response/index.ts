import * as AuthStore from './auth'

const err = {
  code: 0,
  message: 'api error',
  errors: [],
}

export default {
  get: {
    ...AuthStore.get,
  },
  post: {},
  patch: {},
  put: {},
  delete: {},
  error: err,
}
