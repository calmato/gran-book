import Vue from 'vue'
import Vuetify from 'vuetify'
import { createLocalVue } from '@vue/test-utils'

const localVue: Vue.VueConstructor<Vue> = createLocalVue()
const vuetify = new Vuetify()

export { localVue, vuetify }
