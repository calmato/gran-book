import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import NotificationNewForm from '~/components/organisms/NotificationNewForm.vue'

describe('components/organisms/NotificationNewForm', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(NotificationNewForm, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('data', () => {
      it('categoryList', () => {
        expect(wrapper.vm.categoryList).toEqual(['Information', 'Event', 'Maintenance'])
      })

      it('importanceList', () => {
        expect(wrapper.vm.importanceList).toEqual(['High', 'Middle', 'Low'])
      })
    })
  })
})
