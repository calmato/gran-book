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
            .then((role: number) => {
              if (role === 0) {
                throw new Error('forbidden user role')
              }

              router.push('/')
            })
            .catch(() => {
              AuthStore.logout()
            })
        })
        .catch((err: Error) => {
          // TODO: show alert
          console.log('debug', err)
        })
    }

    return {
      form,
      handleSubmit,
    }
  },
})
</script>
