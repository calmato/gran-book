<template>
  <settings-edit-email :form="form" @click="handleSubmit" @cancel="handleCancel" />
</template>

<script lang="ts">
import { defineComponent, reactive, SetupContext } from '@nuxtjs/composition-api'
import { AuthStore, CommonStore } from '~/store'
import { IAuthEditEmailForm, AuthEditEmailOptions } from '~/types/forms'
import SettingsEditEmail from '~/components/templates/SettingsEditEmail.vue'

export default defineComponent({
  components: {
    SettingsEditEmail,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const form = reactive<IAuthEditEmailForm>({
      params: {
        email: store.getters['auth/getEmail'],
      },
      options: {
        ...AuthEditEmailOptions,
      },
    })

    const handleSubmit = async () => {
      await AuthStore.updateEmail(form)
        .then(() => {
          CommonStore.showSnackbar({ color: 'info', message: `確認用のメールを送信しました。` })
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
