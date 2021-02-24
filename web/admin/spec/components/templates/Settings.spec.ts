import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import Settings from '~/components/templates/Settings.vue'

describe('components/templates/Settings', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(Settings, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('username', () => {
        it('初期値', () => {
          expect(wrapper.props().username).toBe('')
        })

        it('値が代入されること', () => {
          wrapper.setProps({ username: 'test-user' })
          expect(wrapper.props().username).toBe('test-user')
        })
      })

      describe('name', () => {
        it('初期値', () => {
          expect(wrapper.props().name).toBe('')
        })

        it('値が代入されること', () => {
          wrapper.setProps({ name: 'Calmato 管理者' })
          expect(wrapper.props().name).toBe('Calmato 管理者')
        })
      })

      describe('nameKana', () => {
        it('初期値', () => {
          expect(wrapper.props().nameKana).toBe('')
        })

        it('値が代入されること', () => {
          wrapper.setProps({ nameKana: 'かるまーと かんりしゃ' })
          expect(wrapper.props().nameKana).toBe('かるまーと かんりしゃ')
        })
      })

      describe('thumbnailUrl', () => {
        it('初期値', () => {
          expect(wrapper.props().thumbnailUrl).toBe('')
        })

        it('値が代入されること', () => {
          wrapper.setProps({ thumbnailUrl: '/thumbnail.png' })
          expect(wrapper.props().thumbnailUrl).toBe('/thumbnail.png')
        })
      })

      describe('selfIntroduction', () => {
        it('初期値', () => {
          expect(wrapper.props().selfIntroduction).toBe('')
        })

        it('値が代入されること', () => {
          wrapper.setProps({ selfIntroduction: 'よろしくお願いします' })
          expect(wrapper.props().selfIntroduction).toBe('よろしくお願いします')
        })
      })

      describe('phoneNumber', () => {
        it('初期値', () => {
          expect(wrapper.props().phoneNumber).toBe('')
        })

        it('値が代入されること', () => {
          wrapper.setProps({ phoneNumber: '000-0000-0000' })
          expect(wrapper.props().phoneNumber).toBe('000-0000-0000')
        })
      })

      describe('email', () => {
        it('初期値', () => {
          expect(wrapper.props().email).toBe('')
        })

        it('値が代入されること', () => {
          wrapper.setProps({ email: 'test@calmato.com' })
          expect(wrapper.props().email).toBe('test@calmato.com')
        })
      })
    })

    describe('data', () => {
      it('profileEditPath', () => {
        expect(wrapper.vm.profileEditPath).toBe('/settings/profile')
      })

      it('profileLists', () => {
        expect(wrapper.vm.profileLists).toEqual([
          {
            title: '表示名',
            content: '',
            contentType: 'text',
          },
          {
            title: '氏名',
            content: '',
            contentType: 'text',
          },
          {
            title: 'アイコン',
            content: '',
            contentType: 'image',
          },
          {
            title: '自己紹介',
            content: '',
            contentType: 'text',
          },
          {
            title: '電話番号',
            content: '',
            contentType: 'text',
          },
        ])
      })

      it('accountLists', () => {
        expect(wrapper.vm.accountLists).toEqual([
          {
            title: 'メールアドレス',
            content: '',
            to: '/settings/email',
          },
          {
            title: 'パスワード',
            content: '************',
            to: '/settings/password',
          },
        ])
      })
    })

    describe('methods', () => {
      describe('onClickEditButton', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickEditButton('/')
          expect(wrapper.emitted('click')).toBeTruthy()
          expect(wrapper.emitted('click')[0][0]).toBe('/')
        })
      })
    })
  })
})
