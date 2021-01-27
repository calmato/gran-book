<template>
  <sign-in :form="form" @submit="handleSubmit" />
</template>

<script lang="ts">
import { defineComponent, reactive, SetupContext } from '@nuxtjs/composition-api'
import { AuthStore } from '~/store'
import { ISignInForm } from '~/types/forms'
import SignIn from '~/components/templates/SignIn.vue'

export default defineComponent({
  layout: 'auth',
  components: {
    SignIn,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router

    const form = reactive<ISignInForm>({
      email: '',
      password: '',
    })

    const handleSubmit = async () => {
      // TODO: エラー処理
      await AuthStore.loginWithEmailAndPassword(form)
        .then(async () => {
          await AuthStore.showAuth()
            .then(() => console.log('debug', 'success'))
            .catch((err: Error) => console.log('debug', 'failure', err))
          router.push('/')
        })
        .catch((err: Error) => console.log('debug', err))
    }

    return {
      form,
      handleSubmit,
    }
  },
})
</script>
