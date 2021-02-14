<template>
  <settings-email-edit :form="form" @click="handleSubmit" @cancel="handleCancel" />
</template>

<script lang="ts">
import { defineComponent, reactive, SetupContext } from '@nuxtjs/composition-api'
import { AuthStore, CommonStore } from '~/store'
import { ISettingsEmailEditForm } from '~/types/forms'
import SettingsEmailEdit from '~/components/templates/SettingsEmailEdit.vue'

export default defineComponent({
  components: {
    SettingsEmailEdit,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const form = reactive<ISettingsEmailEditForm>({
      email: store.getters['auth/getEmail'],
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
