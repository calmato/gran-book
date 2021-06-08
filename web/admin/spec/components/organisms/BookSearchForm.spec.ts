import { mount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import BookSearchForm from '~/components/organisms/BookSearchForm.vue'
import { BookSearchOptions, IBookSearchForm, IBookSearchParams } from '~/types/forms'

describe('components/organisms/BookSearchForm', () => {
  let wrapper: any

  beforeEach(() => {
    const params: IBookSearchParams = {
      title: '',
      author: '',
      publisher: '',
      isbn: '',
      size: 2,
    }
    const form: IBookSearchForm = { params, options: BookSearchOptions }

    wrapper = mount(BookSearchForm, {
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
              title: '',
              author: '',
              publisher: '',
              isbn: '',
              size: 2,
            },
            options: BookSearchOptions,
          })
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
            form: {
              params: {
                title: '小説　ちはやふる　上の句',
                author: '末次 由紀',
                publisher: '講談社',
                isbn: '9784062938426',
                size: 0,
              },
              options: BookSearchOptions,
            },
          })
          expect(wrapper.props().form).toEqual({
            params: {
              title: '小説　ちはやふる　上の句',
              author: '末次 由紀',
              publisher: '講談社',
              isbn: '9784062938426',
              size: 0,
            },
            options: BookSearchOptions,
          })
        })
      })
    })

    describe('data', () => {
      it('menu', () => {
        expect(wrapper.vm.menu).toBeFalsy()
      })
    })

    describe('methods', () => {
      describe('onSubmit', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmit()
          expect(wrapper.emitted('submit')).toBeTruthy()
        })
      })

      describe('onClear', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClear()
          expect(wrapper.emitted('clear')).toBeTruthy()
        })
      })
    })
  })
})
