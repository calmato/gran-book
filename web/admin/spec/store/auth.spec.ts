import { store } from '~~/spec/helpers/store-helper'
import { initialiseStores, AuthStore } from '~/store'

jest.mock('~/plugins/axios', () => ({
  $axios: {
    $get: jest.fn(() =>
      Promise.resolve({
        id: '00000000-0000-0000-00000000',
        username: 'test-user',
        gender: 0,
        phoneNumber: '000-0000-0000',
        role: 0,
        thumbnailUrl: 'https://calmato.com/images/01',
        selfIntroduction: 'よろしくお願いします',
        lastName: 'テスト',
        firstName: 'ユーザ',
        lastNameKana: 'てすと',
        firstNameKana: 'ゆーざ',
        postalCode: '100-0005',
        prefecture: '東京都',
        city: '千代田区',
        addressLine1: '丸の内１丁目',
        addressLine2: '',
        createdAt: '2021-01-01 00:00:00',
        updatedAt: '2021-01-01 00:00:00',
      })
    ),
  }
}))

describe('store/auth', () => {
  beforeEach(() => {
    initialiseStores(store)
  })

  describe('getters', () => {
    it('getEmail', () => {
      expect(AuthStore.getEmail).toBe('')
    })

    it('getToken', () => {
      expect(AuthStore.getToken).toBe('')
    })

    it('getUsername', () => {
      expect(AuthStore.getUsername).toBe('')
    })

    it('getThumbnailUrl', () => {
      expect(AuthStore.getThumbnailUrl).toBe('')
    })
  })

  describe('actions', () => {
    it('showAuth', async () => {
      await AuthStore.showAuth()
      expect(AuthStore.getUsername).toBe('test-user')
      expect(AuthStore.getThumbnailUrl).toBe('https://calmato.com/images/01')
    })
  })
})
