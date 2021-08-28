import { setup, setSafetyMode, refresh } from '~~/spec/helpers/store-helper'
import { AdminStore } from '~/store'
import { ApiError } from '~/types/exception'
import {
  AdminEditContactOptions,
  AdminEditProfileOptions,
  AdminEditSecurityOptions,
  AdminNewOptions,
  IAdminEditContactForm,
  IAdminEditContactParams,
  IAdminEditProfileForm,
  IAdminEditProfileParams,
  IAdminEditSecurityForm,
  IAdminEditSecurityParams,
  IAdminListForm,
  IAdminNewForm,
  IAdminNewParams,
} from '~/types/forms'

describe('store/admin', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
  })

  describe('getters', () => {
    it('getUser', () => {
      expect(AdminStore.getUser).toEqual({
        id: '',
        username: '',
        email: '',
        phoneNumber: '',
        role: 0,
        thumbnailUrl: '',
        selfIntroduction: '',
        lastName: '',
        firstName: '',
        lastNameKana: '',
        firstNameKana: '',
        createdAt: '',
        updatedAt: '',
      })
    })

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
          const err = new ApiError(400, 'api error', {
            status: 400,
            code: 0,
            message: 'api error',
            detail: 'some error',
          })
          await expect(AdminStore.indexAdmin(form)).rejects.toThrow(err)
        })
      })
    })

    describe('showAdmin', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('stateが更新されていること', async () => {
          await AdminStore.showAdmin('00000000-0000-0000-00000000')
          expect(AdminStore.getUser).toEqual({
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
          })
        })
      })

      describe('failure', () => {
        beforeEach(() => {
          setSafetyMode(false)
        })

        it('rejectが返されること', async () => {
          const err = new ApiError(400, 'api error', {
            status: 400,
            code: 0,
            message: 'api error',
            detail: 'some error',
          })
          await expect(AdminStore.showAdmin('')).rejects.toThrow(err)
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
          const err = new ApiError(400, 'api error', {
            status: 400,
            code: 0,
            message: 'api error',
            detail: 'some error',
          })
          await expect(AdminStore.createAdmin(form)).rejects.toThrow(err)
        })
      })
    })

    describe('updateProfile', () => {
      describe('success', () => {
        let form: IAdminEditProfileForm
        let payload: { userId: string; form: IAdminEditProfileForm }
        beforeEach(() => {
          setSafetyMode(true)
          const params: IAdminEditProfileParams = {
            role: 1,
            lastName: 'テスト',
            firstName: 'ユーザ',
            lastNameKana: 'てすと',
            firstNameKana: 'ゆーざ',
            thumbnail: null,
            thumbnailUrl: 'https://calmato.com/images/01',
          }
          form = { params, options: AdminEditProfileOptions }
          payload = { userId: '00000000-0000-0000-00000000', form }
        })

        it('stateが更新されていること', async () => {
          await AdminStore.updateProfile(payload)
          expect(AdminStore.getUser).toEqual({
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
          })
        })
      })

      describe('failure', () => {
        let form: IAdminEditProfileForm
        let payload: { userId: string; form: IAdminEditProfileForm }
        beforeEach(() => {
          setSafetyMode(false)
          const params: IAdminEditProfileParams = {
            role: 0,
            lastName: '',
            firstName: '',
            lastNameKana: '',
            firstNameKana: '',
            thumbnail: null,
            thumbnailUrl: '',
          }
          form = { params, options: AdminEditProfileOptions }
          payload = { userId: '00000000-0000-0000-00000000', form }
        })

        it('rejectが返されること', async () => {
          const err = new ApiError(400, 'api error', {
            status: 400,
            code: 0,
            message: 'api error',
            detail: 'some error',
          })
          await expect(AdminStore.updateProfile(payload)).rejects.toThrow(err)
        })
      })
    })

    describe('updateContact', () => {
      describe('success', () => {
        let form: IAdminEditContactForm
        let payload: { userId: string; form: IAdminEditContactForm }
        beforeEach(() => {
          setSafetyMode(true)
          const params: IAdminEditContactParams = {
            email: 'test@calmato.com',
            phoneNumber: '000-0000-0000',
          }
          form = { params, options: AdminEditContactOptions }
          payload = { userId: '00000000-0000-0000-00000000', form }
        })

        it('stateが更新されていること', async () => {
          await AdminStore.updateContact(payload)
          expect(AdminStore.getUser).toEqual({
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
          })
        })
      })

      describe('failure', () => {
        let form: IAdminEditContactForm
        let payload: { userId: string; form: IAdminEditContactForm }
        beforeEach(() => {
          setSafetyMode(false)
          const params: IAdminEditContactParams = {
            email: '',
            phoneNumber: '000-0000-0000',
          }
          form = { params, options: AdminEditContactOptions }
          payload = { userId: '00000000-0000-0000-00000000', form }
        })

        it('rejectが返されること', async () => {
          const err = new ApiError(400, 'api error', {
            status: 400,
            code: 0,
            message: 'api error',
            detail: 'some error',
          })
          await expect(AdminStore.updateContact(payload)).rejects.toThrow(err)
        })
      })
    })

    describe('updatePassword', () => {
      describe('success', () => {
        let form: IAdminEditSecurityForm
        let payload: { userId: string; form: IAdminEditSecurityForm }
        beforeEach(() => {
          setSafetyMode(true)
          const params: IAdminEditSecurityParams = {
            password: '12345678',
            passwordConfirmation: '12345678',
          }
          form = { params, options: AdminEditSecurityOptions }
          payload = { userId: '00000000-0000-0000-00000000', form }
        })

        it('stateが更新されていること', async () => {
          await AdminStore.updatePassword(payload)
          expect(AdminStore.getUser).toEqual({
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
          })
        })
      })

      describe('failure', () => {
        let form: IAdminEditSecurityForm
        let payload: { userId: string; form: IAdminEditSecurityForm }
        beforeEach(() => {
          setSafetyMode(false)
          const params: IAdminEditSecurityParams = {
            password: '',
            passwordConfirmation: '',
          }
          form = { params, options: AdminEditSecurityOptions }
          payload = { userId: '00000000-0000-0000-00000000', form }
        })

        it('rejectが返されること', async () => {
          const err = new ApiError(400, 'api error', {
            status: 400,
            code: 0,
            message: 'api error',
            detail: 'some error',
          })
          await expect(AdminStore.updatePassword(payload)).rejects.toThrow(err)
        })
      })
    })

    describe('uploadThumbnail', () => {
      describe('success', () => {
        let file: File
        let payload: { userId: string; file: File }
        beforeEach(() => {
          setSafetyMode(true)
          file = new File(['thumbnail'], 'thumbnail.png', { lastModified: Date.now(), type: 'image/png' })
          payload = { userId: '00000000-0000-0000-00000000', file }
        })

        it('resolveが返されること', async () => {
          await expect(AdminStore.uploadThumbnail(payload)).resolves.toBe('https://calmato.com/images/01')
        })
      })

      describe('failure', () => {
        let file: File
        let payload: { userId: string; file: File }
        beforeEach(() => {
          setSafetyMode(false)
          file = new File(['thumbnail'], 'thumbnail.png', { lastModified: Date.now(), type: 'image/png' })
          payload = { userId: '00000000-0000-0000-00000000', file }
        })

        it('rejectが返されること', async () => {
          const err = new ApiError(400, 'api error', {
            status: 400,
            code: 0,
            message: 'api error',
            detail: 'some error',
          })
          await expect(AdminStore.uploadThumbnail(payload)).rejects.toThrow(err)
        })
      })
    })

    describe('deleteAdmin', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('resolveが返されること', async () => {
          await expect(AdminStore.deleteAdmin('00000000-0000-0000-00000000')).resolves.toBeUndefined()
        })
      })

      describe('failure', () => {
        beforeEach(() => {
          setSafetyMode(false)
        })

        it('rejectが返されること', async () => {
          const err = new ApiError(400, 'api error', {
            status: 400,
            code: 0,
            message: 'api error',
            detail: 'some error',
          })
          await expect(AdminStore.deleteAdmin('')).rejects.toThrow(err)
        })
      })
    })
  })
})
