import Vue, { VueConstructor } from 'vue'
import Vuex, { Store } from 'vuex'
import { createLocalVue } from '@vue/test-utils'
import AuthStore from '~/store/auth'

const localVue: VueConstructor<Vue> = createLocalVue()
localVue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    auth: AuthStore,
  },
})

export { localVue, store }
