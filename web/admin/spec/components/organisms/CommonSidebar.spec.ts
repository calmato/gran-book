import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import CommonSidebar from '~/components/organisms/CommonSidebar.vue'

describe('components/ogranisms/CommonSidebar', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(CommonSidebar, {
      ...Options,
      propsData: {
        current: '',
        drawer: true,
      },
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('current', () => {
        it('初期値', () => {
          expect(wrapper.props().current).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ current: '/' })
          expect(wrapper.props().current).toBe('/')
        })
      })

      describe('drawer', () => {
        it('初期値', () => {
          expect(wrapper.props().drawer).toBeTruthy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ drawer: false })
          expect(wrapper.props().current).toBeFalsy()
        })
      })

      describe('username', () => {
        it('初期値', () => {
          expect(wrapper.props().username).toBe('Calmato 管理者')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ username: 'テストユーザ' })
          expect(wrapper.props().username).toBe('テストユーザ')
        })
      })

      describe('email', () => {
        it('初期値', () => {
          expect(wrapper.props().email).toBe('support@calmato.com')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ email: 'test@calmato.com' })
          expect(wrapper.props().email).toBe('test@calmato.com')
        })
      })

      describe('thumbnailUrl', () => {
        it('初期値', () => {
          expect(wrapper.props().thumbnailUrl).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ thumbnailUrl: '/thumbnail.png' })
          expect(wrapper.props().thumbnailUrl).toBe('/thumbnail.png')
        })
      })
    })

    describe('data', () => {
      it('selectedItem', () => {
        expect(wrapper.vm.selectedItem).toBe(-1)
      })

      it('commonItems', () => {
        expect(wrapper.vm.commonItems).toEqual([{ icon: 'mdi-home', text: 'ホーム', to: '/' }])
      })

      it('maintenanceItems', () => {
        expect(wrapper.vm.maintenanceItems).toEqual([
          { icon: 'mdi-cart', text: 'お取り引き管理', to: '/' },
          { icon: 'mdi-forum', text: 'お問い合わせ管理', to: '/' },
          { icon: 'mdi-bell-ring', text: 'お知らせ管理', to: '/' },
          { icon: 'mdi-cash-100', text: 'セール情報管理', to: '/' },
        ])
      })

      it('developerItems', () => {
        expect(wrapper.vm.developerItems).toEqual([
          { icon: 'mdi-account', text: '利用者管理', to: '/' },
          { icon: 'mdi-book', text: '書籍管理', to: '/' },
          { icon: 'mdi-store', text: 'ECサイト管理', to: '/' },
        ])
      })

      it('systemItems', () => {
        expect(wrapper.vm.systemItems).toEqual([
          { icon: 'mdi-shield-account', text: '管理者管理', to: '/admin' },
          { icon: 'mdi-cog', text: 'システム設定', to: '/system' },
          { icon: 'mdi-developer-board', text: 'デバッグ用', to: '/debug' },
        ])
      })
    })

    describe('computed', () => {
      describe('navigationDrawer', () => {
        it('getter', () => {
          expect(wrapper.vm.navigationDrawer).toBeTruthy()
        })

        it('setter', async () => {
          await wrapper.setData({ navigationDrawer: false })
          expect(wrapper.emitted('update:drawer')).toBeTruthy()
          expect(wrapper.emitted('update:drawer')[0][0]).toBeFalsy()
        })
      })
    })

    describe('methods', () => {
      describe('onClick', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClick('/')
          expect(wrapper.emitted('click')).toBeTruthy()
          expect(wrapper.emitted('click')[0][0]).toBe('/')
        })
      })
    })
  })
})
