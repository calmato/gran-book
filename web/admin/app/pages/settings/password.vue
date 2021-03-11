<template>
  <settings-edit-password :form="form" :loading="loading" @click="handleSubmit" @cancel="handleCancel" />
</template>

<script lang="ts">
import { defineComponent, computed, reactive, SetupContext } from '@nuxtjs/composition-api'
import { AuthStore, CommonStore } from '~/store'
import { IAuthEditPasswordForm, AuthEditPasswordOptions, ISignInForm } from '~/types/forms'
import { PromiseState } from '~/types/store'
import SettingsEditPassword from '~/components/templates/SettingsEditPassword.vue'

export default defineComponent({
  components: {
    SettingsEditPassword,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const form = reactive<IAuthEditPasswordForm>({
      params: {
        password: '',
        passwordConfirmation: '',
      },
      options: {
        ...AuthEditPasswordOptions,
      },
    })

    const loading = computed((): boolean => {
      const status = store.getters['common/getPromiseState']
      return status === PromiseState.LOADING
    })

    const handleSubmit = async () => {
      CommonStore.startConnection()
      await AuthStore.updatePassword(form)
        .then(async () => {
          const formData: ISignInForm = {
            email: store.getters['auth/getEmail'],
            password: form.params.password,
          }

          await AuthStore.loginWithEmailAndPassword(formData)

          CommonStore.showSnackbar({ color: 'info', message: `パスワードを変更しました。` })
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
