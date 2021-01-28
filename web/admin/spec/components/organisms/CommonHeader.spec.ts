import { mount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import CommonHeader from '~/components/organisms/CommonHeader.vue'

describe('components/organisms/CommonHeader', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = mount(CommonHeader, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('thumbnailUrl', () => {
        it('初期値', () => {
          expect(wrapper.props().thumbnailUrl).toBe('')
        })

        it('値の代入', () => {
          wrapper.setProps({ thumbnailUrl: '/thumbnail.png' })
          expect(wrapper.props().thumbnailUrl).toBe('/thumbnail.png')
        })
      })
    })

    describe('data', () => {
      it('items', () => {
        expect(wrapper.vm.items).toEqual([{ text: '設定', to: '/system' }])
      })
    })
  })
})
