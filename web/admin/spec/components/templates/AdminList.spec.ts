import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import AdminList from '~/components/templates/AdminList.vue'

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

    describe('methods', () => {
      describe('onClickNewButton', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickNewButton()
          expect(wrapper.emitted('new')).toBeTruthy()
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
