/* eslint-disable import/no-mutable-exports */
import { Store } from 'vuex'
import { getModule } from 'vuex-module-decorators'
import Auth from '~/store/auth'

let AuthStore: Auth

function initialiseStores(store: Store<any>): void {
  AuthStore = getModule(Auth, store)
}

export { initialiseStores, AuthStore }
