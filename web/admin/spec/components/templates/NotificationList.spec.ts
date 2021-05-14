import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import NotificationList from '~/components/templates/NotificationList.vue'

describe('components/templates/NotificationList', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(NotificationList, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('loading', () => {
        it('初期値', () => {
          expect(wrapper.props().loading).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ loading: true })
          expect(wrapper.props().loading).toBeTruthy()
        })
      })

      describe('page', () => {
        it('初期値', () => {
          expect(wrapper.props().page).toBe(1)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ page: 2 })
          expect(wrapper.props().page).toBe(2)
        })
      })

      describe('itemsPerPage', () => {
        it('初期値', () => {
          expect(wrapper.props().itemsPerPage).toBe(20)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ itemsPerPage: 30 })
          expect(wrapper.props().itemsPerPage).toBe(30)
        })
      })

      describe('sortBy', () => {
        it('初期値', () => {
          expect(wrapper.props().sortBy).toBeUndefined()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ sortBy: 'email' })
          expect(wrapper.props().sortBy).toBe('email')
        })
      })

      describe('sortDesc', () => {
        it('初期値', () => {
          expect(wrapper.props().sortDesc).toBeUndefined()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ sortDesc: true })
          expect(wrapper.props().sortDesc).toBeTruthy()
        })
      })

      describe('search', () => {
        it('初期値', () => {
          expect(wrapper.props().search).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ search: 'test' })
          expect(wrapper.props().search).toBe('test')
        })
      })

      describe('total', () => {
        it('初期値', () => {
          expect(wrapper.props().total).toBe(0)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ total: 10 })
          expect(wrapper.props().total).toBe(10)
        })
      })

      describe('newDialog', () => {
        it('初期値', () => {
          expect(wrapper.props().newDialog).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ newDialog: true })
          expect(wrapper.props().newDialog).toBeTruthy()
        })
      })
    })

    describe('data', () => {
      it('dialog', () => {
        expect(wrapper.vm.dialog).toBeFalsy()
      })
    })

    describe('methods', () => {
      describe('onClickNewButton', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickNewButton()
          expect(wrapper.emitted('new:open')).toBeTruthy()
        })
      })

      describe('onClickNewClose', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickNewClose()
          expect(wrapper.emitted('new:close')).toBeTruthy()
        })
      })

      describe('onClickCreateButton', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickCreateButton()
          expect(wrapper.emitted('create')).toBeTruthy()
        })
      })

      describe('onClickEditButton', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickEditButton(1)
          expect(wrapper.emitted('edit')).toBeTruthy()
          expect(wrapper.emitted('edit')[0][0]).toBe(1)
        })
      })

      describe('onClickDeleteButton', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickDeleteButton(1)
          expect(wrapper.emitted('delete')).toBeTruthy()
          expect(wrapper.emitted('delete')[0][0]).toBe(1)
        })
      })
    })
  })
})
