/* eslint-disable import/no-mutable-exports */
import { Store } from 'vuex'
import { getModule } from 'vuex-module-decorators'
import AdminModule from '~/store/admin'
import AuthModule from '~/store/auth'
import CommonModule from '~/store/common'

let AdminStore: AdminModule
let AuthStore: AuthModule
let CommonStore: CommonModule

function initialiseStores(store: Store<any>): void {
  AdminStore = getModule(AdminModule, store)
  AuthStore = getModule(AuthModule, store)
  CommonStore = getModule(CommonModule, store)
}

export { initialiseStores, AdminStore, AuthStore, CommonStore }
