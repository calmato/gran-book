/* eslint-disable import/no-mutable-exports */
import { Store } from 'vuex'
import { getModule } from 'vuex-module-decorators'
import AuthModule from '~/store/auth'

let AuthStore: AuthModule

function initialiseStores(store: Store<any>): void {
  AuthStore = getModule(AuthModule, store)
}

export { initialiseStores, AuthStore }
