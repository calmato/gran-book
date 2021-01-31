import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import SignInForm from '~/components/organisms/SignInForm.vue'
import { ISignInForm } from '~/types/forms'

describe('components/organisms/SignInForm', () => {
  let wrapper: any

  beforeEach(() => {
    const form: ISignInForm = {
      email: '',
      password: '',
    }

    wrapper = shallowMount(SignInForm, {
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
    })

    describe('methods', () => {
      describe('onClick', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClick()
          expect(wrapper.emitted().click).toBeTruthy()
        })
      })
    })
  })
})
