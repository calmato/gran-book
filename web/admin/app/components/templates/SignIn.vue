<template>
  <v-container fill-height>
    <v-layout wrap>
      <v-row justify="center" align="center">
        <v-col cols="12" sm="8">
          <v-img src="/logo.png" max-height="96px" contain class="mb-8" />

          <sign-in-card>
            <sign-in-form :form="form" @click="onClickSubmitButton" />
          </sign-in-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { ISignInForm } from '~/types/forms'
import SignInCard from '~/components/organisms/SignInCard.vue'
import SignInForm from '~/components/organisms/SignInForm.vue'

export default defineComponent({
  components: {
    SignInCard,
    SignInForm,
  },

  props: {
    form: {
      type: Object as PropType<ISignInForm>,
      required: true,
    },
  },

  setup(props, { emit }: SetupContext) {
    const { form } = props

    const onClickSubmitButton = () => {
      emit('submit')
    }

    return {
      form,
      onClickSubmitButton,
    }
  },
})
</script>
