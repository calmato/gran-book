<template>
  <admin-show
    :role="role"
    :user="user"
    :edit-profile.sync="editProfile"
    :edit-contact.sync="editContact"
    :edit-security.sync="editSecurity"
    :edit-profile-form="editProfileForm"
    :edit-contact-form="editContactForm"
    :edit-security-form="editSecurityForm"
    @submit:profile="handleClickUpdateProfile"
    @submit:contact="handleClickUpdateContact"
    @submit:security="handleClickUpdateSecurity"
    @delete="handleClickDelete"
  />
</template>

<script lang="ts">
import { computed, defineComponent, reactive, ref, SetupContext, useAsync } from '@nuxtjs/composition-api'
import AdminShow from '~/components/templates/AdminShow.vue'
import { AdminStore, CommonStore } from '~/store'
import {
  AdminEditContactOptions,
  AdminEditProfileOptions,
  AdminEditSecurityOptions,
  IAdminEditContactForm,
  IAdminEditContactParams,
  IAdminEditProfileForm,
  IAdminEditProfileParams,
  IAdminEditSecurityForm,
  IAdminEditSecurityParams,
} from '~/types/forms'

export default defineComponent({
  components: {
    AdminShow,
  },

  setup(_, { root }: SetupContext) {
    const route = root.$route
    const router = root.$router
    const store = root.$store

    const initializeEditProfileForm: IAdminEditProfileParams = {
      role: 2,
      lastName: '',
      firstName: '',
      lastNameKana: '',
      firstNameKana: '',
      thumbnail: null,
      thumbnailUrl: '',
    }

    const initializeEditContactForm: IAdminEditContactParams = {
      email: '',
      phoneNumber: '',
    }

    const initializeEditSecurityForm: IAdminEditSecurityParams = {
      password: '',
      passwordConfirmation: '',
    }

    const editProfile = ref<boolean>(false)
    const editContact = ref<boolean>(false)
    const editSecurity = ref<boolean>(false)

    const editProfileForm = reactive<IAdminEditProfileForm>({
      params: {
        ...initializeEditProfileForm,
      },
      options: {
        ...AdminEditProfileOptions,
      },
    })

    const editContactForm = reactive<IAdminEditContactForm>({
      params: {
        ...initializeEditContactForm,
      },
      options: {
        ...AdminEditContactOptions,
      },
    })

    const editSecurityForm = reactive<IAdminEditSecurityForm>({
      params: {
        ...initializeEditSecurityForm,
      },
      options: {
        ...AdminEditSecurityOptions,
      },
    })

    const role = computed(() => store.getters['auth/getRole'])
    const user = computed(() => store.getters['admin/getUser'])

    useAsync(async () => {
      const { id } = route.params
      await AdminStore.showAdmin(id).then(() => {
        const { email, phoneNumber, role, lastName, firstName, lastNameKana, firstNameKana, thumbnailUrl } = user.value

        editProfileForm.params = {
          ...initializeEditProfileForm,
          role,
          lastName,
          firstName,
          lastNameKana,
          firstNameKana,
          thumbnailUrl,
        }

        editContactForm.params = {
          ...initializeEditContactForm,
          email,
          phoneNumber,
        }
      })
    })

    const handleClickUpdateProfile = async (): Promise<void> => {
      CommonStore.startConnection()
      await AdminStore.updateProfile({ userId: user.value.id, form: editProfileForm })
        .then(() => {
          editProfile.value = false
          CommonStore.showSnackbar({ color: 'info', message: '管理者情報を更新しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickUpdateContact = async (): Promise<void> => {
      CommonStore.startConnection()
      await AdminStore.updateContact({ userId: user.value.id, form: editContactForm })
        .then(() => {
          editContact.value = false
          CommonStore.showSnackbar({ color: 'info', message: '管理者情報を更新しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickUpdateSecurity = async (): Promise<void> => {
      CommonStore.startConnection()
      await AdminStore.updatePassword({ userId: user.value.id, form: editSecurityForm })
        .then(() => {
          editSecurity.value = false
          CommonStore.showSnackbar({ color: 'info', message: '管理者情報を更新しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickDelete = async (): Promise<void> => {
      CommonStore.startConnection()
      await AdminStore.deleteAdmin(user.value.id)
        .then(() => {
          CommonStore.showSnackbar({ color: 'info', message: '管理者権限を削除しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          router.push('/admin')
          CommonStore.endConnection()
        })
    }

    return {
      role,
      user,
      editProfile,
      editContact,
      editSecurity,
      editProfileForm,
      editContactForm,
      editSecurityForm,
      handleClickUpdateProfile,
      handleClickUpdateContact,
      handleClickUpdateSecurity,
      handleClickDelete,
    }
  },
})
</script>
