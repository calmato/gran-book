<template>
  <v-container fill-height>
    <v-layout wrap>
      <v-row justify="center" align="center">
        <v-col cols="12" sm="8">
          <v-img src="/logo.png" max-height="96px" contain class="mb-8" />
          <the-alert :show.sync="showAlert" type="error">
            メールアドレス もしくは パスワード が間違っています
          </the-alert>
          <sign-in-card>
            <sign-in-form :form="form" @click="onClickSubmitButton" />
          </sign-in-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, computed, SetupContext, PropType } from '@nuxtjs/composition-api'
import { ISignInForm } from '~/types/forms'
import TheAlert from '~/components/molecules/TheAlert.vue'
import SignInCard from '~/components/organisms/SignInCard.vue'
import SignInForm from '~/components/organisms/SignInForm.vue'

export default defineComponent({
  components: {
    TheAlert,
    SignInCard,
    SignInForm,
  },

  props: {
    form: {
      type: Object as PropType<ISignInForm>,
      required: true,
    },
    hasError: {
      type: Boolean,
      required: false,
      default: false,
    },
  },

  setup(props, { emit }: SetupContext) {
    const showAlert = computed({
      get: () => props.hasError,
      set: (val: Boolean) => emit('update:hasError', val),
    })

    const onClickSubmitButton = () => {
      emit('submit')
    }

    return {
      showAlert,
      onClickSubmitButton,
    }
  },
})
</script>
