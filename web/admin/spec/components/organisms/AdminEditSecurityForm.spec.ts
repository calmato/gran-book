import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import AdminEditSecurityForm from '~/components/organisms/AdminEditSecurityForm.vue'
import { AdminEditSecurityOptions, IAdminEditSecurityForm, IAdminEditSecurityParams } from '~/types/forms'

describe('components/organisms/AdminEditSecurityForm', () => {
  let wrapper: any

  beforeEach(() => {
    const params: IAdminEditSecurityParams = {
      password: '',
      passwordConfirmation: '',
    }
    const form: IAdminEditSecurityForm = { params, options: AdminEditSecurityOptions }

    wrapper = shallowMount(AdminEditSecurityForm, {
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
              password: '',
              passwordConfirmation: '',
            },
            options: AdminEditSecurityOptions,
          })
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
            form: {
              params: {
                password: '12345678',
                passwordConfirmation: '12345678',
              },
              options: AdminEditSecurityOptions,
            },
          })
          expect(wrapper.props().form).toEqual({
            params: {
              password: '12345678',
              passwordConfirmation: '12345678',
            },
            options: AdminEditSecurityOptions,
          })
        })
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
    })
  })
})
