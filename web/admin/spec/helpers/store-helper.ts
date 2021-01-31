import Vue, { VueConstructor } from 'vue'
import Vuex from 'vuex'
import { createLocalVue } from '@vue/test-utils'
import AuthStore from '~/store/auth'
import response from '~~/spec/helpers/response'

const localVue: VueConstructor<Vue> = createLocalVue()
localVue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    auth: AuthStore,
  },
})

// Error を返したいときだけ false にする
const isSafetyMode: boolean = true

jest.mock('~/plugins/axios', () => ({
  $axios: {
    $get: (key: string) => (isSafetyMode ? Promise.resolve(response['get'][key]) : Promise.reject(response['error'])),
    $post: (key: string) => (isSafetyMode ? Promise.resolve(response['post'][key]) : Promise.reject(response['error'])),
    $patch: (key: string) =>
      isSafetyMode ? Promise.resolve(response['patch'][key]) : Promise.reject(response['error']),
    $put: (key: string) => (isSafetyMode ? Promise.resolve(response['put'][key]) : Promise.reject(response['error'])),
    $delete: (key: string) =>
      isSafetyMode ? Promise.resolve(response['delete'][key]) : Promise.reject(response['error']),
  },
}))

export { localVue, store, isSafetyMode }
