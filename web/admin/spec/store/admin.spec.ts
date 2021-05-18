import { setup, setSafetyMode, refresh } from '~~/spec/helpers/store-helper'
import { AdminStore } from '~/store'
import { ApiError } from '~/types/exception'
import { AdminNewOptions, IAdminListForm, IAdminNewForm, IAdminNewParams } from '~/types/forms'

describe('store/admin', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
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

        describe('order.by === ""のとき', () => {
          let form: IAdminListForm
          beforeEach(() => {
            form = { limit: 20, offset: 0, order: { by: '', desc: false } }
          })

          it('stateが更新されていること', async () => {
            await AdminStore.indexAdmin(form)
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
            expect(AdminStore.getTotal).toBe(2)
          })
        })

        describe('order.by !== ""のとき', () => {
          let form: IAdminListForm
          beforeEach(() => {
            form = { limit: 20, offset: 0, order: { by: 'email', desc: false } }
          })

          it('stateが更新されていること', async () => {
            await AdminStore.indexAdmin(form)
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
      })

      describe('failure', () => {
        let form: IAdminListForm
        beforeEach(() => {
          setSafetyMode(false)
          form = { limit: 20, offset: 0, order: { by: '', desc: false } }
        })

        it('rejectが返されること', async () => {
          const err = new ApiError(400, 'api error', { status: 400, code: 0, message: 'api error', errors: [] })
          await expect(AdminStore.indexAdmin(form)).rejects.toThrow(err)
        })
      })
    })

    describe('createAdmin', () => {
      describe('success', () => {
        let form: IAdminNewForm
        beforeEach(() => {
          setSafetyMode(true)
          const params: IAdminNewParams = {
            email: 'test@calmato.com',
            password: '12345678',
            passwordConfirmation: '12345678',
            role: 1,
            lastName: 'テスト',
            firstName: 'ユーザ',
            lastNameKana: 'てすと',
            firstNameKana: 'ゆーざ',
          }
          form = { params, options: AdminNewOptions }
        })

        it('stateが更新されていること', async () => {
          await AdminStore.createAdmin(form)
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
        let form: IAdminNewForm
        beforeEach(() => {
          setSafetyMode(false)
          const params: IAdminNewParams = {
            email: '',
            password: '',
            passwordConfirmation: '',
            role: 0,
            lastName: '',
            firstName: '',
            lastNameKana: '',
            firstNameKana: '',
          }
          form = { params, options: AdminNewOptions }
        })

        it('rejectが返されること', async () => {
          const err = new ApiError(400, 'api error', { status: 400, code: 0, message: 'api error', errors: [] })
          await expect(AdminStore.createAdmin(form)).rejects.toThrow(err)
        })
      })
    })
  })
})
