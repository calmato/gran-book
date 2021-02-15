import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import SettingsEditPasswordForm from '~/components/organisms/SettingsEditPasswordForm.vue'
import { IAuthEditPasswordForm, IAuthEditPasswordParams, AuthEditPasswordOptions } from '~/types/forms'

describe('components/organisms/SignInForm', () => {
  let wrapper: any

  beforeEach(() => {
    const params: IAuthEditPasswordParams = { password: '', passwordConfirmation: '' }
    const form: IAuthEditPasswordForm = { params, options: AuthEditPasswordOptions }

    wrapper = shallowMount(SettingsEditPasswordForm, {
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

        it('値が代入されること', () => {
          wrapper.setProps({
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
    })

    describe('methods', () => {
      describe('onClick', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClick()
          expect(wrapper.emitted('click')).toBeTruthy()
        })
      })

      describe('onCancel', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickCancel()
          expect(wrapper.emitted('cancel')).toBeTruthy()
        })
      })
    })
  })
})
