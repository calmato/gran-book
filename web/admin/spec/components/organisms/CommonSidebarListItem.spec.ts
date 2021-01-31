import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import CommonSidebarListItem from '~/components/organisms/CommonSidebarListItem.vue'

describe('components/organisms/CommonSidebarListItem', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(CommonSidebarListItem, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('item', () => {
        it('初期値', () => {
          expect(wrapper.props().item).toEqual({ icon: '', text: '', to: '' })
        })

        it('値が代入されること', () => {
          wrapper.setProps({ item: { icon: 'mdi-home', text: 'ホーム画面', to: '/' } })
          expect(wrapper.props().item).toEqual({ icon: 'mdi-home', text: 'ホーム画面', to: '/' })
        })
      })
    })

    describe('methods', () => {
      describe('onClick', () => {
        it('emitが実行されること', async () => {
          wrapper.setData({ item: { icon: 'mdi-home', text: 'ホーム画面', to: '/' } })
          await wrapper.vm.onClick()
          expect(wrapper.emitted('click')).toBeTruthy()
          expect(wrapper.emitted('click')[0][0]).toBe('/')
        })
      })
    })
  })
})
