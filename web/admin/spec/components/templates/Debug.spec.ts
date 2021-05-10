import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import Debug from '~/components/templates/Debug.vue'

describe('components/templates/Debug', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(Debug, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('userId', () => {
        it('初期値', () => {
          expect(wrapper.props().userId).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ userId: '1e234631-a38f-492c-9008-b558aa0e59e5' })
          expect(wrapper.props().userId).toBe('1e234631-a38f-492c-9008-b558aa0e59e5')
        })
      })

      describe('token', () => {
        it('初期値', () => {
          expect(wrapper.props().token).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ token: 'aiueo12345678' })
          expect(wrapper.props().token).toBe('aiueo12345678')
        })
      })
    })

    describe('data', () => {
      it('showTooltip', () => {
        expect(wrapper.vm.showTooltip).toBeFalsy()
      })
    })

    describe('computed', () => {
      describe('omitToken', () => {
        it('初期値', () => {
          expect(wrapper.vm.omitToken).toBe('')
        })

        it('Tokenの文字数が64未満のとき', async () => {
          await wrapper.setProps({ token: 'a'.repeat(32) })
          expect(wrapper.vm.omitToken).toBe('a'.repeat(32))
        })

        it('Tokenの文字数が64以上のとき', async () => {
          await wrapper.setProps({ token: 'a'.repeat(80) })
          expect(wrapper.vm.omitToken).toBe('a'.repeat(64) + '...')
        })
      })
    })
  })
})
