import { store, setSafetyMode } from '~~/spec/helpers/store-helper'
import { initialiseStores, AdminStore } from '~/store'
import { ApiError } from '~/types/exception'
import { IAdminListForm } from '~/types/forms'

describe('store/admin', () => {
  beforeEach(() => {
    initialiseStores(store)
  })

  describe('getters', () => {
    it('getUsers', () => {
      expect(AdminStore.getUsers).toEqual([])
    })

    it('getTotal', () => {
      expect(AdminStore.getTotal).toBe(0)
    })
  })

  describe('actions', () => {
    describe('indexAdmin', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('resolveが返されること', async () => {
          const payload: IAdminListForm = { limit: 20, offset: 0, order: { by: 'email', desc: false } }
          await expect(AdminStore.indexAdmin(payload)).resolves
        })

        it('stateが更新されていること', async () => {
          const payload: IAdminListForm = { limit: 20, offset: 0, order: { by: 'email', desc: false } }
          await AdminStore.indexAdmin(payload)
          expect(AdminStore.getUsers).toEqual([
            {
              id: '00000000-0000-0000-00000000',
              username: 'test-user',
              email: 'test@calmato.com',
              phoneNumber: '000-0000-0000',
              role: 0,
              thumbnailUrl: 'https://calmato.com/images/01',
              selfIntroduction: 'よろしくお願いします',
              lastName: 'テスト',
              firstName: 'ユーザ',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざ',
              createdAt: '2021-01-01 00:00:00',
              updatedAt: '2021-01-01 00:00:00',
            },
          ])
          expect(AdminStore.getTotal).toBe(1)
        })
      })

      describe('failure', () => {
        beforeEach(() => {
          setSafetyMode(false)
        })

        it('rejectが返されること', async () => {
          const payload: IAdminListForm = { limit: 20, offset: 0, order: { by: '', desc: false } }
          const err = new ApiError(400, 'api error', { status: 400, code: 0, message: 'api error', errors: [] })
          await expect(AdminStore.indexAdmin(payload)).rejects.toThrow(err)
        })
      })
    })
  })
})
