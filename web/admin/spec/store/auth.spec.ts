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

    it('getPhoneNumber', () => {
      expect(AuthStore.getPhoneNumber).toBe('')
    })

    it('getThumbnailUrl', () => {
      expect(AuthStore.getThumbnailUrl).toBe('/thumbnail.png')
    })

    it('getSelfIntroduction', () => {
      expect(AuthStore.getSelfIntroduction).toBe('')
    })

    it('getName', () => {
      expect(AuthStore.getName).toBe('')
    })

    it('getNameKana', () => {
      expect(AuthStore.getNameKana).toBe('')
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
