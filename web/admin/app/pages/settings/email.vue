<template>
  <settings-edit-email :form="form" :loading="loading" @click="handleSubmit" @cancel="handleCancel" />
</template>

<script lang="ts">
import { defineComponent, computed, reactive, SetupContext } from '@nuxtjs/composition-api'
import { AuthStore, CommonStore } from '~/store'
import { IAuthEditEmailForm, AuthEditEmailOptions } from '~/types/forms'
import { PromiseState } from '~/types/store'
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

    const loading = computed((): boolean => {
      const status = store.getters['common/getPromiseState']
      return status === PromiseState.LOADING
    })

    const handleSubmit = async () => {
      CommonStore.startConnection()
      await AuthStore.updateEmail(form)
        .then(async () => {
          await AuthStore.getIdToken()
          await AuthStore.sendEmailVerification()

          // TODO: ログイン画面にメッセージが出るように
          AuthStore.logout()
          router.push('/signin')
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
