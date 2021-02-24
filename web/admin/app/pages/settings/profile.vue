<template>
  <settings-edit-profile :form="form" @click="handleSubmit" @cancel="handleCancel" />
</template>

<script lang="ts">
import { defineComponent, reactive, SetupContext } from '@nuxtjs/composition-api'
import { AuthStore, CommonStore } from '~/store'
import { IAuthEditProfileForm, AuthEditProfileOptions } from '~/types/forms'
import SettingsEditProfile from '~/components/templates/SettingsEditProfile.vue'

export default defineComponent({
  components: {
    SettingsEditProfile,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const form = reactive<IAuthEditProfileForm>({
      params: {
        username: store.getters['auth/getUsername'],
        thumbnail: undefined,
        selfIntroduction: store.getters['auth/getSelfIntroduction'],
        lastName: store.getters['auth/getLastName'],
        firstName: store.getters['auth/getFirstName'],
        lastNameKana: store.getters['auth/getLastNameKana'],
        firstNameKana: store.getters['auth/getFirstNameKana'],
        phoneNumber: store.getters['auth/getPhoneNumber'],
      },
      options: {
        ...AuthEditProfileOptions,
      },
    })

    const handleSubmit = async () => {
      await AuthStore.updateProfile(form)
        .then(() => {
          CommonStore.showSnackbar({ color: 'info', message: 'プロフィールを変更しました。' })
          router.push('/settings')
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
    }

    const handleCancel = () => {
      router.back()
    }

    return {
      form,
      handleSubmit,
      handleCancel,
    }
  },
})
</script>
