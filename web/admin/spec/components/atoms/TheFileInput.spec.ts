import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import TheFileInput from '~/components/atoms/TheFileInput.vue'

describe('components/atoms/TheFileInput', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(TheFileInput, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('accept', () => {
        it('初期値', () => {
          expect(wrapper.props().accept).toBe('image/*')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ accept: 'video/*' })
          expect(wrapper.props().accept).toBe('video/*')
        })
      })

      describe('label', () => {
        it('初期値', () => {
          expect(wrapper.props().label).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ label: 'テスト' })
          expect(wrapper.props().label).toBe('テスト')
        })
      })

      describe('name', () => {
        it('初期値', () => {
          expect(wrapper.props().name).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ name: 'test' })
          expect(wrapper.props().name).toBe('test')
        })
      })

      describe('rules', () => {
        it('初期値', () => {
          expect(wrapper.props().rules).toEqual({})
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ rules: { required: true } })
          expect(wrapper.props().rules).toEqual({ required: true })
        })
      })

      describe('file', () => {
        it('初期値', () => {
          expect(wrapper.props().file).toBeNull()
        })

        it('値が代入されること', async () => {
          const file = new File(['thumbnail'], 'thumbnail.png', { lastModified: Date.now(), type: 'image/png' })
          await wrapper.setProps({ file })
          expect(wrapper.props().file).toBe(file)
        })
      })
    })

    describe('methods', () => {
      describe('onSelect', () => {
        describe('fileが入力されたとき', () => {
          it('emitが実行されること', async () => {
            const file = new File(['thumbnail'], 'thumbnail.png', { lastModified: Date.now(), type: 'image/png' })
            await wrapper.vm.onSelect(file)
            expect(wrapper.emitted('change')).toBeTruthy()
            expect(wrapper.emitted('change')[0][0]).toBe(file)
          })
        })
      })
    })
  })
})
