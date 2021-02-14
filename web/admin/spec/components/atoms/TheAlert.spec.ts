import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import TheAlert from '~/components/atoms/TheAlert.vue'

describe('components/atoms/TheAlert', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(TheAlert, {
      ...Options,
      propsData: {
        show: false,
      },
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('type', () => {
        it('初期値', () => {
          expect(wrapper.props().type).toBe('info')
        })

        it('値が代入されること', () => {
          wrapper.setProps({ type: 'warning' })
          expect(wrapper.props().type).toBe('warning')
        })
      })

      describe('show', () => {
        it('初期値', () => {
          expect(wrapper.props().show).toBeFalsy()
        })

        it('値が代入されること', () => {
          wrapper.setProps({ show: true })
          expect(wrapper.props().show).toBeTruthy()
        })
      })
    })

    describe('methods', () => {
      describe('hiddenAlert', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.hiddenAlert()
          expect(wrapper.emitted('update:show')).toBeTruthy()
          expect(wrapper.emitted('update:show')[0][0]).toBeFalsy()
        })
      })
    })
  })
})
