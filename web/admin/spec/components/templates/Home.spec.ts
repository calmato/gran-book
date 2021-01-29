import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import Home from '~/components/templates/Home.vue'

describe('components/templates/Home', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(Home, { ...Options })
  })

  describe('script', () => {
    describe('data', () => {
      it('cards', () => {
        expect(wrapper.vm.cards).toEqual(['Today', 'Yesterday'])
      })
    })
  })
})
