/* eslint-disable import/no-mutable-exports */
import { Store } from 'vuex'
import { getModule } from 'vuex-module-decorators'
import AuthModule from '~/store/auth'
import CommonModule from '~/store/common'

let AuthStore: AuthModule
let CommonStore: CommonModule

function initialiseStores(store: Store<any>): void {
  AuthStore = getModule(AuthModule, store)
  CommonStore = getModule(CommonModule, store)
}

export { initialiseStores, AuthStore, CommonStore }
