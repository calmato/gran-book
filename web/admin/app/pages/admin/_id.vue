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
import { AdminStore } from '~/store'
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

    const handleClickUpdateProfile = (): void => {
      console.log('debug', 'update profile', editProfileForm)
      editProfile.value = false
    }

    const handleClickUpdateContact = (): void => {
      console.log('debug', 'update contact', editContactForm)
      editContact.value = false
    }

    const handleClickUpdateSecurity = (): void => {
      console.log('debug', 'update security', editSecurityForm)
      editSecurity.value = false
    }

    const handleClickDelete = (): void => {
      console.log('debug', 'delete')
      editProfile.value = false
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
