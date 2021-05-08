import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import SettingsEditEmailForm from '~/components/organisms/SettingsEditEmailForm.vue'
import { IAuthEditEmailForm, IAuthEditEmailParams, AuthEditEmailOptions } from '~/types/forms'

describe('components/organisms/SettingsEditEmailForm', () => {
  let wrapper: any

  beforeEach(() => {
    const params: IAuthEditEmailParams = { email: '' }
    const form: IAuthEditEmailForm = { params, options: AuthEditEmailOptions }

    wrapper = shallowMount(SettingsEditEmailForm, {
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
