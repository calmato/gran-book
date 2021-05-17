import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import AdminEditForm from '~/components/organisms/AdminEditForm.vue'
import { AdminEditOptions, IAdminEditForm, IAdminEditParams } from '~/types/forms'

describe('components/organisms/AdminEditForm', () => {
  let wrapper: any

  beforeEach(() => {
    const params: IAdminEditParams = {
      email: '',
      phoneNumber: '',
      role: 2,
      lastName: '',
      firstName: '',
      lastNameKana: '',
      firstNameKana: '',
      thumbnail: null,
      thumbnailUrl: '',
    }
    const form: IAdminEditForm = { params, options: AdminEditOptions }

    wrapper = shallowMount(AdminEditForm, {
      ...Options,
      propsData: { form },
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('form', () => {
        it('初期値', () => {
          expect(wrapper.props().form).toEqual({
            params: {
              email: '',
              phoneNumber: '',
              role: 2,
              lastName: '',
              firstName: '',
              lastNameKana: '',
              firstNameKana: '',
              thumbnail: null,
              thumbnailUrl: '',
            },
            options: AdminEditOptions,
          })
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
            form: {
              params: {
                email: 'test@calmato.com',
                phoneNumber: '000-0000-0000',
                role: 1,
                lastName: 'テスト',
                firstName: 'ユーザー',
                lastNameKana: 'てすと',
                firstNameKana: 'ゆーざー',
                thumbnail: null,
                thumbnailUrl: 'https://calmato.com/images/01',
              },
              options: AdminEditOptions,
            },
          })
          expect(wrapper.props().form).toEqual({
            params: {
              email: 'test@calmato.com',
              phoneNumber: '000-0000-0000',
              role: 1,
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
              thumbnail: null,
              thumbnailUrl: 'https://calmato.com/images/01',
            },
            options: AdminEditOptions,
          })
        })
      })
    })

    describe('data', () => {
      it('roleItems', () => {
        expect(wrapper.vm.roleItems).toEqual([
          { text: '管理者', value: 1 },
          { text: '開発者', value: 2 },
          { text: '運用者', value: 3 },
        ])
      })
    })
  })
})
