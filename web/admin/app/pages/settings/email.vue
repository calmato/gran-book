<template>
  <settings-edit-email :form="form" :rules="rules" @click="handleSubmit" @cancel="handleCancel" />
</template>

<script lang="ts">
import { defineComponent, reactive, SetupContext } from '@nuxtjs/composition-api'
import { AuthStore, CommonStore } from '~/store'
import { AuthEditEmailValidate, IAuthEditEmailForm, IAuthEditEmailValidate } from '~/types/forms'
import SettingsEditEmail from '~/components/templates/SettingsEditEmail.vue'

export default defineComponent({
  components: {
    SettingsEditEmail,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const form = reactive<IAuthEditEmailForm>({
      email: store.getters['auth/getEmail'],
    })
    const rules = reactive<IAuthEditEmailValidate>({
      ...AuthEditEmailValidate,
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
      rules,
      handleSubmit,
      handleCancel,
    }
  },
})
</script>
