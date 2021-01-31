import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import SignIn from '~/components/templates/SignIn.vue'
import { ISignInForm } from '~/types/forms'

describe('components/templates/Home', () => {
  let wrapper: any

  beforeEach(() => {
    const form: ISignInForm = {
      email: '',
      password: '',
    }

    wrapper = shallowMount(SignIn, {
      ...Options,
      propsData: { form },
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('form', () => {
        it('初期値', () => {
          expect(wrapper.props().form).toEqual({ email: '', password: '' })
        })

        it('値が代入されること', () => {
          wrapper.setProps({ form: { email: 'test@calmato.com', password: '12345678' } })
          expect(wrapper.props().form).toEqual({ email: 'test@calmato.com', password: '12345678' })
        })
      })

      describe('hasError', () => {
        it('初期値', () => {
          expect(wrapper.props().hasError).toBeFalsy()
        })

        it('値が代入されること', () => {
          wrapper.setProps({ hasError: true })
          expect(wrapper.props().hasError).toBeTruthy()
        })
      })
    })

    describe('computed', () => {
      describe('showAlert', () => {
        it('getter', () => {
          expect(wrapper.vm.showAlert).toBeFalsy()
        })

        it('setter', async () => {
          await wrapper.setData({ showAlert: true })
          expect(wrapper.emitted('update:hasError')).toBeTruthy()
          expect(wrapper.emitted('update:hasError')[0][0]).toBeTruthy()
        })
      })
    })

    describe('methods', () => {
      describe('onClickSubmitButton', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickSubmitButton()
          expect(wrapper.emitted().submit).toBeTruthy()
        })
      })
    })
  })
})
