import { mount } from '@vue/test-utils'
import '~~/spec/helpers/component-helper'
import CommonHeader from '~/components/organisms/CommonHeader.vue'

describe('components/organisms/CommonHeader', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = mount(CommonHeader)
  })

  describe('script', () => {
    describe('props', () => {
      describe('thumbnailUrl', () => {
        test('初期値', () => {
          expect(wrapper.props().thumbnailUrl).toBe('')
        })

        test('代入', () => {
          wrapper.setProps({ thumbnailUrl: '/thumbnail.png' })
          expect(wrapper.props().thumbnailUrl).toBe('/thumbnail.png')
        })
      })
    })

    describe('data', () => {
      describe('items', () => {
        expect(wrapper.vm.items).toEqual([{ text: '設定', to: '/system' }])
      })
    })
  })
})
