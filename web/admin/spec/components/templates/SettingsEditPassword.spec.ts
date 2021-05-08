import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import SettingsEditPassword from '~/components/templates/SettingsEditPassword.vue'
import { IAuthEditPasswordForm, IAuthEditPasswordParams, AuthEditPasswordOptions } from '~/types/forms'

describe('components/templates/SettingsEditEmail', () => {
  let wrapper: any

  beforeEach(() => {
    const params: IAuthEditPasswordParams = { password: '', passwordConfirmation: '' }
    const form: IAuthEditPasswordForm = { params, options: AuthEditPasswordOptions }

    wrapper = shallowMount(SettingsEditPassword, {
      ...Options,
      propsData: { form },
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('form', () => {
        it('初期値', () => {
          expect(wrapper.props().form).toEqual({
            params: { password: '', passwordConfirmation: '' },
            options: AuthEditPasswordOptions,
          })
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
            form: {
              params: { password: '12345678', passwordConfirmation: '12345678' },
              options: AuthEditPasswordOptions,
            },
          })
          expect(wrapper.props().form).toEqual({
            params: { password: '12345678', passwordConfirmation: '12345678' },
            options: AuthEditPasswordOptions,
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
