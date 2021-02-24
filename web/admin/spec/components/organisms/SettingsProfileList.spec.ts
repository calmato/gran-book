import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import SettingsProfileList from '~/components/organisms/SettingsProfileList.vue'

describe('components/organisms/SettingsProfileList', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(SettingsProfileList, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('lists', () => {
        it('初期値', () => {
          expect(wrapper.props().lists).toEqual([])
        })

        it('値が代入されること', () => {
          wrapper.setProps({ lists: [{ title: '表示名', content: 'calmato', contentType: 'text' }] })
          expect(wrapper.props().lists).toEqual([{ title: '表示名', content: 'calmato', contentType: 'text' }])
        })
      })
    })

    describe('methods', () => {
      describe('onClick', () => {
        it('emitが実行されること', async () => {
          wrapper.setProps({ lists: [{ title: '表示名', content: 'calmato', contentType: 'text' }] })
          await wrapper.vm.onClick()
          expect(wrapper.emitted('click')).toBeTruthy()
        })
      })
    })
  })
})
