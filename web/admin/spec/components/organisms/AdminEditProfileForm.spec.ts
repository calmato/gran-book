import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import AdminEditProfileForm from '~/components/organisms/AdminEditProfileForm.vue'
import { AdminEditProfileOptions, IAdminEditProfileForm, IAdminEditProfileParams } from '~/types/forms'

describe('components/organisms/AdminEditProfileForm', () => {
  let wrapper: any

  beforeEach(() => {
    const params: IAdminEditProfileParams = {
      role: 2,
      lastName: '',
      firstName: '',
      lastNameKana: '',
      firstNameKana: '',
      thumbnail: null,
      thumbnailUrl: '',
    }
    const form: IAdminEditProfileForm = { params, options: AdminEditProfileOptions }

    wrapper = shallowMount(AdminEditProfileForm, {
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
              role: 2,
              lastName: '',
              firstName: '',
              lastNameKana: '',
              firstNameKana: '',
              thumbnail: null,
              thumbnailUrl: '',
            },
            options: AdminEditProfileOptions,
          })
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
            form: {
              params: {
                role: 1,
                lastName: 'テスト',
                firstName: 'ユーザー',
                lastNameKana: 'てすと',
                firstNameKana: 'ゆーざー',
                thumbnail: null,
                thumbnailUrl: 'https://calmato.com/images/01',
              },
              options: AdminEditProfileOptions,
            },
          })
          expect(wrapper.props().form).toEqual({
            params: {
              role: 1,
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
              thumbnail: null,
              thumbnailUrl: 'https://calmato.com/images/01',
            },
            options: AdminEditProfileOptions,
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

    describe('methods', () => {
      describe('onSubmit', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmit()
          expect(wrapper.emitted('submit')).toBeTruthy()
        })
      })

      describe('onCancel', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onCancel()
          expect(wrapper.emitted('cancel')).toBeTruthy()
        })
      })

      describe('onDelete', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onDelete()
          expect(wrapper.emitted('delete')).toBeTruthy()
        })
      })
    })
  })
})
