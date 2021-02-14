<template>
  <settings-edit-password :form="form" @click="handleSubmit" @cancel="handleCancel" />
</template>

<script lang="ts">
import { defineComponent, reactive, SetupContext } from '@nuxtjs/composition-api'
import { AuthStore, CommonStore } from '~/store'
import { IAuthEditPasswordForm, AuthEditPasswordOptions } from '~/types/forms'
import SettingsEditPassword from '~/components/templates/SettingsEditPassword.vue'

export default defineComponent({
  components: {
    SettingsEditPassword,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router

    const form = reactive<IAuthEditPasswordForm>({
      params: {
        password: '',
        passwordConfirmation: '',
      },
      options: {
        ...AuthEditPasswordOptions,
      },
    })

    const handleSubmit = async () => {
      await AuthStore.updatePassword(form)
        .then(() => {
          CommonStore.showSnackbar({ color: 'info', message: `パスワードを変更しました。` })
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
