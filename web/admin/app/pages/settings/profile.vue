<template>
  <settings-edit-profile :form="form" :loading="loading" @click="handleSubmit" @cancel="handleCancel" />
</template>

<script lang="ts">
import { defineComponent, computed, reactive, SetupContext } from '@nuxtjs/composition-api'
import { AuthStore, CommonStore } from '~/store'
import { IAuthEditProfileForm, AuthEditProfileOptions } from '~/types/forms'
import { PromiseState } from '~/types/store'
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
        thumbnail: null,
        thumbnailUrl: store.getters['auth/getThumbnailUrl'],
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

    const loading = computed((): boolean => {
      const status = store.getters['common/getPromiseState']
      return status === PromiseState.LOADING
    })

    const uploadThumbnail = (file: File | null): Promise<string> => {
      if (!file) {
        return Promise.resolve('')
      }

      return AuthStore.uploadThumbnail(file)
        .then((res: string) => {
          return res
        })
        .catch((err: Error) => {
          throw err
        })
    }

    const handleSubmit = async () => {
      CommonStore.startConnection()
      await uploadThumbnail(form.params.thumbnail)
        .then((res: string) => {
          if (res !== '') {
            form.params.thumbnailUrl = res
          }

          return AuthStore.updateProfile(form)
        })
        .then(() => {
          CommonStore.showSnackbar({ color: 'info', message: 'プロフィールを変更しました。' })
          router.push('/settings')
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleCancel = () => {
      router.back()
    }

    return {
      form,
      loading,
      handleSubmit,
      handleCancel,
    }
  },
})
</script>
