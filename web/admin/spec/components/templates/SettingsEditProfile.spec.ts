import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import SettingsEditProfile from '~/components/templates/SettingsEditProfile.vue'
import { IAuthEditProfileForm, IAuthEditProfileParams, AuthEditProfileOptions } from '~/types/forms'

describe('components/templates/SettingsEditProfile', () => {
  let wrapper: any

  beforeEach(() => {
    const params: IAuthEditProfileParams = {
      username: '',
      thumbnail: undefined,
      selfIntroduction: '',
      lastName: '',
      firstName: '',
      lastNameKana: '',
      firstNameKana: '',
      phoneNumber: '',
    }
    const form: IAuthEditProfileForm = { params, options: AuthEditProfileOptions }

    wrapper = shallowMount(SettingsEditProfile, {
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
              username: '',
              thumbnail: undefined,
              selfIntroduction: '',
              lastName: '',
              firstName: '',
              lastNameKana: '',
              firstNameKana: '',
              phoneNumber: '',
            },
            options: AuthEditProfileOptions,
          })
        })

        it('値が代入されること', () => {
          wrapper.setProps({
            form: {
              params: {
                username: 'test-user',
                thumbnail: undefined,
                selfIntroduction: 'よろしく',
                lastName: 'テスト',
                firstName: 'ユーザ',
                lastNameKana: 'てすと',
                firstNameKana: 'ゆーざ',
                phoneNumber: '000-0000-0000',
              },
              options: AuthEditProfileOptions,
            },
          })
          expect(wrapper.props().form).toEqual({
            params: {
              username: 'test-user',
              thumbnail: undefined,
              selfIntroduction: 'よろしく',
              lastName: 'テスト',
              firstName: 'ユーザ',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざ',
              phoneNumber: '000-0000-0000',
            },
            options: AuthEditProfileOptions,
          })
        })
      })

      describe('loading', () => {
        it('初期値', () => {
          expect(wrapper.props().loading).toBeFalsy()
        })

        it('値が代入されること', () => {
          wrapper.setProps({ loading: true })
          expect(wrapper.props().loading).toBeTruthy()
        })
      })
    })

    describe('methods', () => {
      describe('onClickSubmitButton', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickSubmitButton()
          expect(wrapper.emitted('click')).toBeTruthy()
        })
      })

      describe('onClickCancelButton', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickCancelButton()
          expect(wrapper.emitted('cancel')).toBeTruthy()
        })
      })
    })
  })
})
