import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import AdminShowContactList from '~/components/organisms/AdminShowContactList.vue'

describe('components/organisms/AdminShowContactList', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(AdminShowContactList, { ...Options })
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
  })
})
