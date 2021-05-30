import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import AdminShowProfileList from '~/components/organisms/AdminShowProfileList.vue'

describe('components/organisms/AdminShowProfileList', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(AdminShowProfileList, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('user', () => {
        it('初期値', () => {
          expect(wrapper.props().user).toEqual({})
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
            user: {
              id: '5',
              username: 'test-user',
              email: 'test@calmato.com',
              phoneNumber: '123-1234-1234',
              role: 3,
              thumbnailUrl: '',
              selfIntroduction: '',
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
              createdAt: '2020-01-01 00:00:00',
              updatedAt: '2020-01-01 00:00:00',
            },
          })
          expect(wrapper.props().user).toEqual({
            id: '5',
            username: 'test-user',
            email: 'test@calmato.com',
            phoneNumber: '123-1234-1234',
            role: 3,
            thumbnailUrl: '',
            selfIntroduction: '',
            lastName: 'テスト',
            firstName: 'ユーザー',
            lastNameKana: 'てすと',
            firstNameKana: 'ゆーざー',
            createdAt: '2020-01-01 00:00:00',
            updatedAt: '2020-01-01 00:00:00',
          })
        })
      })
    })

    describe('methods', () => {
      describe('getName', () => {
        it('値が返されること', async () => {
          await wrapper.setProps({
            user: {
              id: '5',
              username: 'test-user',
              email: 'test@calmato.com',
              phoneNumber: '123-1234-1234',
              role: 3,
              thumbnailUrl: '',
              selfIntroduction: '',
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
              createdAt: '2020-01-01 00:00:00',
              updatedAt: '2020-01-01 00:00:00',
            },
          })
          expect(wrapper.vm.getName()).toBe('テスト ユーザー (てすと ゆーざー)')
        })
      })

      describe('getRole', () => {
        it('値が返されること', async () => {
          await wrapper.setProps({
            user: {
              id: '5',
              username: 'test-user',
              email: 'test@calmato.com',
              phoneNumber: '123-1234-1234',
              role: 3,
              thumbnailUrl: '',
              selfIntroduction: '',
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
              createdAt: '2020-01-01 00:00:00',
              updatedAt: '2020-01-01 00:00:00',
            },
          })
          expect(wrapper.vm.getRole()).toBe('Operator')
        })
      })
    })
  })
})
