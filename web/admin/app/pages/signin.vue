<template>
  <sign-in :form="form" :is-loading="isLoading" :has-error.sync="hasError" @submit="handleSubmit" />
</template>

<script lang="ts">
import { defineComponent, computed, reactive, ref, SetupContext } from '@nuxtjs/composition-api'
import { AuthStore, CommonStore } from '~/store'
import { ISignInForm } from '~/types/forms'
import SignIn from '~/components/templates/SignIn.vue'
import { PromiseState } from '~/types/store'

export default defineComponent({
  layout: 'auth',
  components: {
    SignIn,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const hasError = ref<Boolean>(false)
    const form = reactive<ISignInForm>({
      email: '',
      password: '',
    })

    const isLoading = computed((): boolean => {
      const status = store.getters['common/getPromiseState']
      return status === PromiseState.LOADING
    })

    const handleSubmit = async () => {
      CommonStore.startConnection()
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
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    return {
      form,
      hasError,
      isLoading,
      handleSubmit,
    }
  },
})
</script>
