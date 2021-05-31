import { mount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import AdminListTable from '~/components/organisms/AdminListTable.vue'
import { IAdminUser } from '~/types/store'

describe('components/organisms/AdminListTable', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = mount(AdminListTable, { ...Options })
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

      describe('role', () => {
        it('初期値', () => {
          expect(wrapper.props().role).toBe(0)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ role: 1 })
          expect(wrapper.props().role).toBe(1)
        })
      })

      describe('users', () => {
        it('初期値', () => {
          expect(wrapper.props().users).toEqual([])
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
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

        it('値が代入されること', async () => {
          await wrapper.setProps({ total: 10 })
          expect(wrapper.props().total).toBe(10)
        })
      })
    })

    describe('data', () => {
      it('headers', () => {
        expect(wrapper.vm.headers).toEqual([
          { text: 'サムネ', value: 'thumbnailUrl', sortable: false },
          { text: '氏名', value: 'name', sortable: false },
          { text: 'メールアドレス', value: 'email', sortable: true },
          { text: '電話番号', value: 'phoneNumber', sortable: false },
          { text: '権限', value: 'role', sortable: true },
        ])
      })

      it('footers', () => {
        expect(wrapper.vm.footers).toEqual({
          itemsPerPageOptions: [10, 20, 30, 50, 100],
        })
      })
    })

    describe('computed', () => {
      describe('items', () => {
        it('change props', async () => {
          await wrapper.setProps({
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
          expect(wrapper.vm.items).toEqual([
            {
              id: '5',
              name: 'テスト ユーザー',
              email: 'test@calmato.com',
              phoneNumber: '123-1234-1234',
              thumbnailUrl: '/thumbnail.png',
              role: 3,
            },
          ])
        })
      })

      describe('sortByArray', () => {
        it('getter', () => {
          expect(wrapper.vm.sortByArray).toEqual([undefined])
        })

        it('setter', async () => {
          await wrapper.setData({ sortByArray: ['email'] })
          expect(wrapper.emitted('update:sort-by')).toBeTruthy()
          expect(wrapper.emitted('update:sort-by')[0][0]).toBe('email')
        })
      })

      describe('sortDescArray', () => {
        it('getter', () => {
          expect(wrapper.vm.sortDescArray).toEqual([undefined])
        })

        it('setter', async () => {
          await wrapper.setData({ sortDescArray: [true] })
          expect(wrapper.emitted('update:sort-desc')).toBeTruthy()
          expect(wrapper.emitted('update:sort-desc')[0][0]).toBeTruthy()
        })
      })
    })

    describe('methods', () => {
      describe('getRole', () => {
        it('role: 1', () => {
          expect(wrapper.vm.getRole(1)).toBe('Administrator')
        })

        it('role: 2', () => {
          expect(wrapper.vm.getRole(2)).toBe('Developer')
        })

        it('role: 3', () => {
          expect(wrapper.vm.getRole(3)).toBe('Operator')
        })

        it('role: other', () => {
          expect(wrapper.vm.getRole(4)).toBe('Unknown')
        })
      })

      describe('getRoleColor', () => {
        it('role: 1', () => {
          expect(wrapper.vm.getRoleColor(1)).toBe('red')
        })

        it('role: 2', () => {
          expect(wrapper.vm.getRoleColor(2)).toBe('orange')
        })

        it('role: 3', () => {
          expect(wrapper.vm.getRoleColor(3)).toBe('green')
        })

        it('role: other', () => {
          expect(wrapper.vm.getRoleColor(4)).toBe('')
        })
      })

      describe('onClick', () => {
        it('emitが実行されること', async () => {
          const user: IAdminUser = {
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
          }
          await wrapper.setProps({ users: [user] })
          await wrapper.vm.onClick(wrapper.vm.items[0])
          expect(wrapper.emitted('show')).toBeTruthy()
          expect(wrapper.emitted('show')[0][0]).toBe('5')
        })
      })
    })
  })
})
