import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import AdminShow from '~/components/templates/AdminShow.vue'
import { AdminEditContactOptions, AdminEditProfileOptions, AdminEditSecurityOptions, AdminNewOptions } from '~/types/forms'

describe('components/templates/AdminShow', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = shallowMount(AdminShow, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('cuid', () => {
        it('初期値', () => {
          expect(wrapper.props().cuid).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ cuid: '00000000-0000-0000-000000000000' })
          expect(wrapper.props().cuid).toBe('00000000-0000-0000-000000000000')
        })
      })

      describe('role', () => {
        it('初期値', () => {
          expect(wrapper.props().role).toBe(0)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ role: 1 })
          expect(wrapper.props().role).toBe(1)
        })
      })

      describe('user', () => {
        it('初期値', () => {
          expect(wrapper.props().user).toEqual({})
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
            user: {
              id: '5',
              username: 'test-user',
              email: 'test@calmato.com',
              phoneNumber: '123-1234-1234',
              role: 3,
              thumbnailUrl: '',
              selfIntroduction: '',
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
              createdAt: '2020-01-01 00:00:00',
              updatedAt: '2020-01-01 00:00:00',
            },
          })
          expect(wrapper.props().user).toEqual({
            id: '5',
            username: 'test-user',
            email: 'test@calmato.com',
            phoneNumber: '123-1234-1234',
            role: 3,
            thumbnailUrl: '',
            selfIntroduction: '',
            lastName: 'テスト',
            firstName: 'ユーザー',
            lastNameKana: 'てすと',
            firstNameKana: 'ゆーざー',
            createdAt: '2020-01-01 00:00:00',
            updatedAt: '2020-01-01 00:00:00',
          })
        })
      })

      describe('editProfile', () => {
        it('初期値', () => {
          expect(wrapper.props().editProfile).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ editProfile: true })
          expect(wrapper.props().editProfile).toBeTruthy()
        })
      })

      describe('editContact', () => {
        it('初期値', () => {
          expect(wrapper.props().editContact).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ editContact: true })
          expect(wrapper.props().editContact).toBeTruthy()
        })
      })

      describe('editSecurity', () => {
        it('初期値', () => {
          expect(wrapper.props().editSecurity).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ editSecurity: true })
          expect(wrapper.props().editSecurity).toBeTruthy()
        })
      })

      describe('editProfileForm', () => {
        it('初期値', () => {
          expect(wrapper.props().editProfileForm).toEqual({})
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
            editProfileForm: {
              params: {
                role: 1,
                lastName: 'テスト',
                firstName: 'ユーザー',
                lastNameKana: 'てすと',
                firstNameKana: 'ゆーざー',
                thumbnail: null,
                thumbnailUrl: '',
              },
              options: AdminEditProfileOptions,
            },
          })
          expect(wrapper.props().editProfileForm).toEqual({
            params: {
              role: 1,
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
              thumbnail: null,
              thumbnailUrl: ''
            },
            options: AdminEditProfileOptions,
          })
        })
      })

      describe('editContactForm', () => {
        it('初期値', () => {
          expect(wrapper.props().editContactForm).toEqual({})
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
            editContactForm: {
              params: {
                email: 'test@calmato.com',
                phoneNumber: '123-1234-1234',
              },
              options: AdminEditContactOptions,
            },
          })
          expect(wrapper.props().editContactForm).toEqual({
            params: {
              email: 'test@calmato.com',
              phoneNumber: '123-1234-1234',
            },
            options: AdminEditContactOptions,
          })
        })
      })

      describe('editSecurityForm', () => {
        it('初期値', () => {
          expect(wrapper.props().editSecurityForm).toEqual({})
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
            editSecurityForm: {
              params: {
                password: '12345678',
                passwordConfirmation: '12345678',
              },
              options: AdminEditSecurityOptions,
            },
          })
          expect(wrapper.props().editSecurityForm).toEqual({
            params: {
              password: '12345678',
              passwordConfirmation: '12345678',
            },
            options: AdminEditSecurityOptions,
          })
        })
      })
    })

    describe('data', () => {
      it('roleItems', () => {
        expect(wrapper.vm.roleItems).toEqual([
          { text: '管理者', value: 1 },
          { text: '開発者', value: 2 },
          { text: '運用者', value: 3 },
        ])
      })

      it('deleteDialog', () => {
        expect(wrapper.vm.deleteDialog).toBeFalsy()
      })
    })

    describe('methods', () => {
      describe('openDeleteDialog', () => {
        it('deleteDialogがtrueになること', async () => {
          await wrapper.vm.openDeleteDialog()
          expect(wrapper.vm.deleteDialog).toBeTruthy()
        })
      })

      describe('closeDeleteDialog', () => {
        it('deleteDialogがfalseになること', async () => {
          await wrapper.vm.closeDeleteDialog()
          expect(wrapper.vm.deleteDialog).toBeFalsy()
        })
      })

      describe('canEdit', () => {
        it('role === 1 && cuid !== user.idのとき', async () => {
          await wrapper.setProps({
            cuid: '00000000-0000-0000-000000000000',
            role: 1,
            user: {
              id: '11111111-1111-1111-111111111111',
              username: 'test-user',
              email: 'test@calmato.com',
              phoneNumber: '123-1234-1234',
              role: 3,
              thumbnailUrl: '',
              selfIntroduction: '',
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
              createdAt: '2020-01-01 00:00:00',
              updatedAt: '2020-01-01 00:00:00',
            },
          })
          expect(wrapper.vm.canEdit()).toBeTruthy()
        })

        it('role !== 1 && cuid !== user.idのとき', async () => {
          await wrapper.setProps({
            cuid: '00000000-0000-0000-000000000000',
            role: 2,
            user: {
              id: '11111111-1111-1111-111111111111',
              username: 'test-user',
              email: 'test@calmato.com',
              phoneNumber: '123-1234-1234',
              role: 3,
              thumbnailUrl: '',
              selfIntroduction: '',
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
              createdAt: '2020-01-01 00:00:00',
              updatedAt: '2020-01-01 00:00:00',
            },
          })
          expect(wrapper.vm.canEdit()).toBeFalsy()
        })

        it('role === 1 && cuid === user.idのとき', async () => {
          await wrapper.setProps({
            cuid: '00000000-0000-0000-000000000000',
            role: 1,
            user: {
              id: '00000000-0000-0000-000000000000',
              username: 'test-user',
              email: 'test@calmato.com',
              phoneNumber: '123-1234-1234',
              role: 3,
              thumbnailUrl: '',
              selfIntroduction: '',
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
              createdAt: '2020-01-01 00:00:00',
              updatedAt: '2020-01-01 00:00:00',
            },
          })
          expect(wrapper.vm.canEdit()).toBeFalsy()
        })

        it('role !== 1 && cuid === user.idのとき', async () => {
          await wrapper.setProps({
            cuid: '00000000-0000-0000-000000000000',
            role: 2,
            user: {
              id: '00000000-0000-0000-000000000000',
              username: 'test-user',
              email: 'test@calmato.com',
              phoneNumber: '123-1234-1234',
              role: 3,
              thumbnailUrl: '',
              selfIntroduction: '',
              lastName: 'テスト',
              firstName: 'ユーザー',
              lastNameKana: 'てすと',
              firstNameKana: 'ゆーざー',
              createdAt: '2020-01-01 00:00:00',
              updatedAt: '2020-01-01 00:00:00',
            },
          })
          expect(wrapper.vm.canEdit()).toBeFalsy()
        })
      })

      describe('onSubmitProfile', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmitProfile()
          expect(wrapper.emitted('submit:profile')).toBeTruthy()
        })
      })

      describe('onSubmitContact', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmitContact()
          expect(wrapper.emitted('submit:contact')).toBeTruthy()
        })
      })

      describe('onSubmitSecurity', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmitSecurity()
          expect(wrapper.emitted('submit:security')).toBeTruthy()
        })
      })

      describe('onSubmitDelete', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmitDelete('1')
          expect(wrapper.emitted('delete')).toBeTruthy()
        })

        it('deleteDialogがfalseになること', async () => {
          await wrapper.setProps({ deleteDialog: true })
          await wrapper.vm.onSubmitDelete()
          expect(wrapper.vm.deleteDialog).toBeFalsy()
        })
      })
    })
  })
})
