<template>
  <sign-in :form="form" :has-error.sync="hasError" @submit="handleSubmit" />
</template>

<script lang="ts">
import { defineComponent, ref, reactive, SetupContext } from '@nuxtjs/composition-api'
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

    const hasError = ref<Boolean>(false)

    const handleSubmit = async () => {
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
        .catch(() => {
          hasError.value = true
        })
    }

    return {
      form,
      hasError,
      handleSubmit,
    }
  },
})
</script>
