import { mount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import NotificationListTable from '~/components/organisms/NotificationListTable.vue'

describe('components/organisms/NotificationListTable', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = mount(NotificationListTable, { ...Options })
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
    })

    describe('data', () => {
      it('headers', () => {
        expect(wrapper.vm.headers).toEqual([
          { text: 'タイトル', value: 'title', sortable: false },
          { text: '説明', value: 'description', sortable: false },
          { text: '作成日', value: 'timestamp', sortable: true },
          { text: 'カテゴリー', value: 'category', sortable: false },
          { text: '重要度', value: 'importance', sortable: false },
          { text: 'Actions', value: 'actions', sortable: false },
        ])
      })

      it('footers', () => {
        expect(wrapper.vm.footers).toEqual({
          itemsPerPageOptions: [10, 20, 30, 50, 100],
        })
      })
    })

    describe('computed', () => {
      // TODO: itemsのテスト
      // describe('items', () => {})

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
      describe('getCategory', () => {
        it('category: 1', () => {
          expect(wrapper.vm.getCategory(1)).toBe('Maintenance')
        })

        it('category: 2', () => {
          expect(wrapper.vm.getCategory(2)).toBe('Event')
        })

        it('category: 3', () => {
          expect(wrapper.vm.getCategory(3)).toBe('Information')
        })

        it('category: other', () => {
          expect(wrapper.vm.getCategory(4)).toBe('Unknown')
        })
      })

      describe('getCategoryColor', () => {
        it('category: 1', () => {
          expect(wrapper.vm.getCategoryColor(1)).toBe('blue-grey')
        })

        it('category: 2', () => {
          expect(wrapper.vm.getCategoryColor(2)).toBe('purple lighten-3')
        })

        it('category: 3', () => {
          expect(wrapper.vm.getCategoryColor(3)).toBe('light-green')
        })

        it('category: other', () => {
          expect(wrapper.vm.getCategoryColor(4)).toBe('')
        })
      })

      describe('getImportance', () => {
        it('importance: 1', () => {
          expect(wrapper.vm.getImportance(1)).toBe('High')
        })

        it('importance: 2', () => {
          expect(wrapper.vm.getImportance(2)).toBe('Middle')
        })

        it('importance: 3', () => {
          expect(wrapper.vm.getImportance(3)).toBe('Low')
        })

        it('importance: other', () => {
          expect(wrapper.vm.getImportance(4)).toBe('Unknown')
        })
      })

      describe('getImportanceColor', () => {
        it('importance: 1', () => {
          expect(wrapper.vm.getImportanceColor(1)).toBe('red')
        })

        it('importance: 2', () => {
          expect(wrapper.vm.getImportanceColor(2)).toBe('amber')
        })

        it('importance: 3', () => {
          expect(wrapper.vm.getImportanceColor(3)).toBe('light-blue')
        })

        it('importance: other', () => {
          expect(wrapper.vm.getImportanceColor(4)).toBe('')
        })
      })

      describe('onClickEdit', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickEdit(wrapper.vm.items[0])
          expect(wrapper.emitted('edit')).toBeTruthy()
          expect(wrapper.emitted('edit')[0][0]).toBe(0)
        })
      })

      describe('onClickDelete', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickDelete(wrapper.vm.items[0])
          expect(wrapper.emitted('delete')).toBeTruthy()
          expect(wrapper.emitted('delete')[0][0]).toEqual(0)
        })
      })
    })
  })
})
