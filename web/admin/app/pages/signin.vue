<template>
  <v-container fill-height>
    <v-layout wrap>
      <v-row justify="center" align="center">
        <v-col cols="12" sm="8">
          <v-img src="/logo.png" max-height="96px" contain class="mb-8" />

          <v-card rounded="32">
            <v-card-title class="justify-center">
              <h3>管理者向けログイン</h3>
            </v-card-title>
            <v-divider />
            <v-card-text class="px-12">
              <v-form class="py-8">
                <v-text-field v-model="form.email" label="email" autofocus outlined />
                <v-text-field v-model="form.password" label="password" type="password" outlined />
                <v-btn :block="true" color="primary" class="mt-4" @click="handleClick">ログイン</v-btn>
              </v-form>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, reactive } from '@nuxtjs/composition-api'
import { AuthStore } from '~/store'
import { SignInForm } from '~/types/forms'

export default defineComponent({
  setup() {
    const form = reactive<SignInForm>({
      email: '',
      password: '',
    })

    const handleClick = async () => {
      // TODO: エラー処理
      await AuthStore.loginWithEmailAndPassword(form).catch((err: Error) => console.log('debug', err))
    }

    return {
      form,
      handleClick,
    }
  },
})
</script>
