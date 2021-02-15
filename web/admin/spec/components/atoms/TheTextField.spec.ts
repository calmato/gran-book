import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import TheTextField from '~/components/atoms/TheTextField.vue'

describe('components/atoms/TheTextField', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(TheTextField, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('autofocus', () => {
        it('初期値', () => {
          expect(wrapper.props().autofocus).toBeFalsy()
        })

        it('値が代入されること', () => {
          wrapper.setProps({ autofocus: true })
          expect(wrapper.props().autofocus).toBeTruthy()
        })
      })

      describe('label', () => {
        it('初期値', () => {
          expect(wrapper.props().label).toBe('')
        })

        it('値が代入されること', () => {
          wrapper.setProps({ label: 'テスト' })
          expect(wrapper.props().label).toBe('テスト')
        })
      })

      describe('name', () => {
        it('初期値', () => {
          expect(wrapper.props().name).toBe('')
        })

        it('値が代入されること', () => {
          wrapper.setProps({ name: 'test' })
          expect(wrapper.props().name).toBe('test')
        })
      })

      describe('outlined', () => {
        it('初期値', () => {
          expect(wrapper.props().outlined).toBeFalsy()
        })

        it('値が代入されること', () => {
          wrapper.setProps({ outlined: true })
          expect(wrapper.props().outlined).toBeTruthy()
        })
      })

      describe('rules', () => {
        it('初期値', () => {
          expect(wrapper.props().rules).toEqual({})
        })

        it('値が代入されること', () => {
          wrapper.setProps({ rules: { required: true } })
          expect(wrapper.props().rules).toEqual({ required: true })
        })
      })

      describe('type', () => {
        it('初期値', () => {
          expect(wrapper.props().type).toBe('text')
        })

        it('値が代入されること', () => {
          wrapper.setProps({ type: 'password' })
          expect(wrapper.props().type).toBe('password')
        })
      })

      describe('value', () => {
        it('初期値', () => {
          expect(wrapper.props().value).toBe('')
        })

        it('値が代入されること', () => {
          wrapper.setProps({ value: 'test' })
          expect(wrapper.props().value).toBe('test')
        })
      })
    })

    describe('computed', () => {
      describe('formData', () => {
        it('getter', () => {
          expect(wrapper.vm.formData).toBe('')
        })

        it('setter', async () => {
          await wrapper.setData({ formData: 'test' })
          expect(wrapper.emitted('input')).toBeTruthy()
          expect(wrapper.emitted('input')[0][0]).toBe('test')
        })
      })
    })
  })
})
