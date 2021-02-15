import { store, setSafetyMode } from '~~/spec/helpers/store-helper'
import { initialiseStores, AuthStore } from '~/store'
import {
  IAuthEditEmailForm,
  AuthEditEmailOptions,
  IAuthEditEmailParams,
  IAuthEditPasswordForm,
  IAuthEditPasswordParams,
  AuthEditPasswordOptions,
} from '~/types/forms'
import { ApiError } from '~/types/exception'

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
    describe('showAuth', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('resolveが返されること', async () => {
          await expect(AuthStore.showAuth()).resolves.toBe(0)
        })

        it('stateが更新されていること', async () => {
          await AuthStore.showAuth()
          expect(AuthStore.getUsername).toBe('test-user')
          expect(AuthStore.getThumbnailUrl).toBe('https://calmato.com/images/01')
        })
      })

      describe('failure', () => {
        beforeEach(() => {
          setSafetyMode(false)
        })

        it('rejectが返されること', async () => {
          const err = new ApiError(400, 'api error', { status: 400, code: 0, message: 'api error', errors: [] })
          await expect(AuthStore.showAuth()).rejects.toThrow(err)
        })
      })
    })

    describe('updateEmail', () => {
      describe('success', () => {
        let form: IAuthEditEmailForm
        beforeEach(() => {
          setSafetyMode(true)
          const params: IAuthEditEmailParams = { email: 'test@calmato.com' }
          form = { params, options: AuthEditEmailOptions }
        })

        it('resolveが返されること', async () => {
          await expect(AuthStore.updateEmail(form)).resolves.toBeUndefined()
        })

        it('stateが更新されていること', async () => {
          await AuthStore.updateEmail(form)
          expect(AuthStore.getEmail).toBe('test@calmato.com')
        })
      })

      describe('failure', () => {
        let form: IAuthEditEmailForm
        beforeEach(() => {
          setSafetyMode(false)
          const params: IAuthEditEmailParams = { email: '' }
          form = { params, options: AuthEditEmailOptions }
        })

        it('rejectが返されること', async () => {
          const err = new ApiError(400, 'api error', { status: 400, code: 0, message: 'api error', errors: [] })
          await expect(AuthStore.updateEmail(form)).rejects.toThrow(err)
        })
      })
    })

    describe('updatePassword', () => {
      describe('success', () => {
        let form: IAuthEditPasswordForm
        beforeEach(() => {
          setSafetyMode(true)
          const params: IAuthEditPasswordParams = { password: '12345678', passwordConfirmation: '12345678' }
          form = { params, options: AuthEditPasswordOptions }
        })

        it('resolveが返されること', async () => {
          await expect(AuthStore.updatePassword(form)).resolves.toBeUndefined()
        })
      })

      describe('failure', () => {
        let form: IAuthEditPasswordForm
        beforeEach(() => {
          setSafetyMode(false)
          const params: IAuthEditPasswordParams = { password: '', passwordConfirmation: '' }
          form = { params, options: AuthEditPasswordOptions }
        })

        it('rejectが返されること', async () => {
          const err = new ApiError(400, 'api error', { status: 400, code: 0, message: 'api error', errors: [] })
          await expect(AuthStore.updatePassword(form)).rejects.toThrow(err)
        })
      })
    })
  })
})
