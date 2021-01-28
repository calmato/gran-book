import createPersistedState from 'vuex-persistedstate'
import * as Cookies from 'js-cookie'

export default ({ store, isDev }) => {
  createPersistedState({
    paths: ['auth'],
    storage: {
      getItem: (key: string) => Cookies.get(key),
      setItem: (key: string, value: any) => Cookies.set(key, value, { expires: 7, secure: !isDev }),
      removeItem: (key: string) => Cookies.remove(key),
    },
  })(store)
}
