import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import AdminEditContactForm from '~/components/organisms/AdminEditContactForm.vue'
import { AdminEditContactOptions, IAdminEditContactForm, IAdminEditContactParams } from '~/types/forms'

describe('components/organisms/AdminEditContactForm', () => {
  let wrapper: any

  beforeEach(() => {
    const params: IAdminEditContactParams = {
      email: '',
      phoneNumber: '',
    }
    const form: IAdminEditContactForm = { params, options: AdminEditContactOptions }

    wrapper = shallowMount(AdminEditContactForm, {
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
            },
            options: AdminEditContactOptions,
          })
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
            form: {
              params: {
                email: 'test@calmato.com',
                phoneNumber: '000-0000-0000',
              },
              options: AdminEditContactOptions,
            },
          })
          expect(wrapper.props().form).toEqual({
            params: {
              email: 'test@calmato.com',
              phoneNumber: '000-0000-0000',
            },
            options: AdminEditContactOptions,
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
