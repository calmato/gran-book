import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import SettingsEditEmail from '~/components/templates/SettingsEditEmail.vue'
import { IAuthEditEmailForm, IAuthEditEmailParams, AuthEditEmailOptions } from '~/types/forms'

describe('components/templates/SettingsEditEmail', () => {
  let wrapper: any

  beforeEach(() => {
    const params: IAuthEditEmailParams = { email: '' }
    const form: IAuthEditEmailForm = { params, options: AuthEditEmailOptions }

    wrapper = shallowMount(SettingsEditEmail, {
      ...Options,
      propsData: { form },
    })
  })
  describe('script', () => {
    describe('props', () => {
      describe('form', () => {
        it('初期値', () => {
          expect(wrapper.props().form).toEqual({
            params: { email: '' },
            options: AuthEditEmailOptions,
          })
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
            form: {
              params: { email: 'test@calmato.com' },
              options: AuthEditEmailOptions,
            },
          })
          expect(wrapper.props().form).toEqual({
            params: { email: 'test@calmato.com' },
            options: AuthEditEmailOptions,
          })
        })
      })

      describe('loading', () => {
        it('初期値', () => {
          expect(wrapper.props().loading).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ loading: true })
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
