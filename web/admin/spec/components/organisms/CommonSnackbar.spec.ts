import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import CommonSnackbar from '~/components/organisms/CommonSnackbar.vue'

describe('components/organisms/CommonSnackbar', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(CommonSnackbar, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('color', () => {
        it('初期値', () => {
          expect(wrapper.props().color).toBe('info')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ color: 'error' })
          expect(wrapper.props().color).toBe('error')
        })
      })

      describe('message', () => {
        it('初期値', () => {
          expect(wrapper.props().message).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ message: 'メッセージが変更されました' })
          expect(wrapper.props().message).toBe('メッセージが変更されました')
        })
      })

      describe('snackbar', () => {
        it('初期値', () => {
          expect(wrapper.props().snackbar).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ snackbar: true })
          expect(wrapper.props().snackbar).toBeTruthy()
        })
      })
    })

    describe('computed', () => {
      describe('isShow', () => {
        it('getter', () => {
          expect(wrapper.vm.isShow).toBeFalsy()
        })

        it('setter', async () => {
          await wrapper.setData({ isShow: true })
          expect(wrapper.emitted('update:snackbar')).toBeTruthy()
          expect(wrapper.emitted('update:snackbar')[0][0]).toBeTruthy()
        })
      })
    })
  })
})
