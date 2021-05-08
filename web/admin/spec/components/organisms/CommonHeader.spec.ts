import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import CommonHeader from '~/components/organisms/CommonHeader.vue'

describe('components/organisms/CommonHeader', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(CommonHeader, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('thumbnailUrl', () => {
        it('初期値', () => {
          expect(wrapper.props().thumbnailUrl).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ thumbnailUrl: '/thumbnail.png' })
          expect(wrapper.props().thumbnailUrl).toBe('/thumbnail.png')
        })
      })
    })

    describe('data', () => {
      it('items', () => {
        expect(wrapper.vm.items).toEqual([{ text: '設定', to: '/settings' }])
      })
    })

    describe('methods', () => {
      describe('onClick', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClick('/')
          expect(wrapper.emitted('click')).toBeTruthy()
          expect(wrapper.emitted('click')[0][0]).toBe('/')
        })
      })

      describe('onClickLogout', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickLogout()
          expect(wrapper.emitted('logout')).toBeTruthy()
        })
      })

      describe('onChange', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onChange()
          expect(wrapper.emitted('change')).toBeTruthy()
        })
      })
    })
  })
})
