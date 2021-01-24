<template>
  <sign-in :form="form" @submit="handleSubmit" />
</template>

<script lang="ts">
import { defineComponent, reactive } from '@nuxtjs/composition-api'
import { AuthStore } from '~/store'
import { ISignInForm } from '~/types/forms'
import SignIn from '~/components/templates/SignIn.vue'

export default defineComponent({
  layout: 'auth',
  components: {
    SignIn,
  },

  setup() {
    const form = reactive<ISignInForm>({
      email: '',
      password: '',
    })

    const handleSubmit = async () => {
      // TODO: エラー処理
      await AuthStore.loginWithEmailAndPassword(form).catch((err: Error) => console.log('debug', err))
    }

    return {
      form,
      handleSubmit,
    }
  },
})
</script>
