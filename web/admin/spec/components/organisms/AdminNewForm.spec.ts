import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import AdminNewForm from '~/components/organisms/AdminNewForm.vue'
import { AdminNewOptions, IAdminNewForm, IAdminNewParams } from '~/types/forms'

describe('components/organisms/AdminNewForm', () => {
  let wrapper: any

  beforeEach(() => {
    const params: IAdminNewParams = {
      email: '',
      password: '',
      passwordConfirmation: '',
      role: 2,
      lastName: '',
      firstName: '',
      lastNameKana: '',
      firstNameKana: '',
    }
    const form: IAdminNewForm = { params, options: AdminNewOptions }

    wrapper = shallowMount(AdminNewForm, {
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
              password: '',
              passwordConfirmation: '',
              role: 2,
              lastName: '',
              firstName: '',
              lastNameKana: '',
              firstNameKana: '',
            },
            options: AdminNewOptions,
          })
        })

        it('値が代入されること', () => {
          wrapper.setProps({
            form: {
              params: {
                email: 'test@calmato.com',
                password: '12345678',
                passwordConfirmation: '12345678',
                role: 1,
                lastName: 'テスト',
                firstName: 'ユーザー',
                lastNameKana: 'てすと',
                firstNameKana: 'ゆーざー',
              },
              options: AdminNewOptions,
            },
          })
          expect(wrapper.props().form).toEqual({
            params: {
              email: 'test@calmato.com',
              password: '12345678',
              passwordConfirmation: '12345678',
              role: 1,
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
            },
            options: AdminNewOptions,
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
