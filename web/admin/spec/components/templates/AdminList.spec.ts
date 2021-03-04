import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import AdminList from '~/components/templates/AdminList.vue'
import { AdminNewOptions } from '~/types/forms'

describe('components/templates/AdminList', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(AdminList, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('loading', () => {
        it('初期値', () => {
          expect(wrapper.props().loading).toBeFalsy()
        })

        it('値が代入されること', () => {
          wrapper.setProps({ loading: true })
          expect(wrapper.props().loading).toBeTruthy()
        })
      })

      describe('page', () => {
        it('初期値', () => {
          expect(wrapper.props().page).toBe(1)
        })

        it('値が代入されること', () => {
          wrapper.setProps({ page: 2 })
          expect(wrapper.props().page).toBe(2)
        })
      })

      describe('itemsPerPage', () => {
        it('初期値', () => {
          expect(wrapper.props().itemsPerPage).toBe(20)
        })

        it('値が代入されること', () => {
          wrapper.setProps({ itemsPerPage: 30 })
          expect(wrapper.props().itemsPerPage).toBe(30)
        })
      })

      describe('sortBy', () => {
        it('初期値', () => {
          expect(wrapper.props().sortBy).toBeUndefined()
        })

        it('値が代入されること', () => {
          wrapper.setProps({ sortBy: 'email' })
          expect(wrapper.props().sortBy).toBe('email')
        })
      })

      describe('sortDesc', () => {
        it('初期値', () => {
          expect(wrapper.props().sortDesc).toBeUndefined()
        })

        it('値が代入されること', () => {
          wrapper.setProps({ sortDesc: true })
          expect(wrapper.props().sortDesc).toBeTruthy()
        })
      })

      describe('search', () => {
        it('初期値', () => {
          expect(wrapper.props().search).toBe('')
        })

        it('値が代入されること', () => {
          wrapper.setProps({ search: 'test' })
          expect(wrapper.props().search).toBe('test')
        })
      })

      describe('users', () => {
        it('初期値', () => {
          expect(wrapper.props().users).toEqual([])
        })

        it('値が代入されること', () => {
          wrapper.setProps({
            users: [
              {
                id: '5',
                username: 'test-user',
                email: 'test@calmato.com',
                phoneNumber: '123-1234-1234',
                role: 3,
                thumbnailUrl: '',
                selfIntroduction: '',
                lastName: 'テスト',
                firstName: 'ユーザー',
                lastNameKana: 'てすと',
                firstNameKana: 'ゆーざー',
                createdAt: '2020-01-01 00:00:00',
                updatedAt: '2020-01-01 00:00:00',
              },
            ],
          })
          expect(wrapper.props().users).toEqual([
            {
              id: '5',
              username: 'test-user',
              email: 'test@calmato.com',
              phoneNumber: '123-1234-1234',
              role: 3,
              thumbnailUrl: '',
              selfIntroduction: '',
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
              createdAt: '2020-01-01 00:00:00',
              updatedAt: '2020-01-01 00:00:00',
            },
          ])
        })
      })

      describe('newForm', () => {
        it('初期値', () => {
          expect(wrapper.props().newForm).toEqual({})
        })

        it('値が代入されること', () => {
          wrapper.setProps({
            newForm: {
              params: {
                email: 'test@calmato.com',
                password: '12345678',
                passwordConfirmation: '12345678',
                role: 1,
                lastName: 'テスト',
                firstName: 'ユーザー',
                lastNameKana: 'てすと',
                firstNameKana: 'ゆーざー',
              },
              options: AdminNewOptions,
            },
          })
          expect(wrapper.props().newForm).toEqual({
            params: {
              email: 'test@calmato.com',
              password: '12345678',
              passwordConfirmation: '12345678',
              role: 1,
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
            },
            options: AdminNewOptions,
          })
        })
      })

      describe('newDialog', () => {
        it('初期値', () => {
          expect(wrapper.props().newDialog).toBeFalsy()
        })

        it('値が代入されること', () => {
          wrapper.setProps({ newDialog: true })
          expect(wrapper.props().newDialog).toBeTruthy()
        })
      })

      describe('total', () => {
        it('初期値', () => {
          expect(wrapper.props().total).toBe(0)
        })

        it('値が代入されること', () => {
          wrapper.setProps({ total: 10 })
          expect(wrapper.props().total).toBe(10)
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