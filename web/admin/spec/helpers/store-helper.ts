import Vue, { VueConstructor } from 'vue'
import Vuex from 'vuex'
import { createLocalVue } from '@vue/test-utils'
import { AdminStore, AuthStore, CommonStore, initialiseStores } from '~/store'
import AdminModule from '~/store/admin'
import AuthModule from '~/store/auth'
import CommonModule from '~/store/common'
import response from '~~/spec/helpers/response'

const localVue: VueConstructor<Vue> = createLocalVue()
localVue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    admin: AdminModule,
    auth: AuthModule,
    common: CommonModule,
  },
})

function setup(): void {
  initialiseStores(store)
}

function refresh(): void {
  AdminStore.factory()
  AuthStore.factory()
  CommonStore.factory()
}

// Error を返したいときだけ false にする
let isSafetyMode: boolean = true

function setSafetyMode(mode: boolean): void {
  isSafetyMode = mode
}

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

export { localVue, setup, refresh, setSafetyMode }
