import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import SettingsAccountList from '~/components/organisms/SettingsAccountList.vue'

describe('components/organisms/SettingsAccountList', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(SettingsAccountList, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('lists', () => {
        it('初期値', () => {
          expect(wrapper.props().lists).toEqual([])
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ lists: [{ title: 'パスワード', content: '************', to: '/settings/password' }] })
          expect(wrapper.props().lists).toEqual([
            { title: 'パスワード', content: '************', to: '/settings/password' },
          ])
        })
      })
    })

    describe('methods', () => {
      describe('onClick', async () => {
        it('emitが実行されること', async () => {
          await wrapper.setProps({ lists: [{ title: 'パスワード', content: '************', to: '/settings/password' }] })
          await wrapper.vm.onClick('/settings/password')
          expect(wrapper.emitted('click')).toBeTruthy()
          expect(wrapper.emitted('click')[0][0]).toBe('/settings/password')
        })
      })
    })
  })
})
