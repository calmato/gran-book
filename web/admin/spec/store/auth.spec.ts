import { store } from '~~/spec/helpers/store-helper'
import { initialiseStores, AuthStore } from '~/store'

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
