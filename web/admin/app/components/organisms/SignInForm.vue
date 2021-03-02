<template>
  <v-form class="py-8">
    <!-- TODO: バリデーションの設定 -->
    <v-text-field v-model="form.email" label="email" autofocus outlined />
    <v-text-field v-model="form.password" label="password" type="password" outlined />
    <v-btn :block="true" :loading="loading" color="primary" class="mt-4" @click="onClick">ログイン</v-btn>
  </v-form>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { ISignInForm } from '~/types/forms'

export default defineComponent({
  props: {
    form: {
      type: Object as PropType<ISignInForm>,
      required: true,
    },
    loading: {
      type: Boolean,
      required: false,
      default: false,
    },
  },

  setup(_, { emit }: SetupContext) {
    const onClick = () => {
      emit('click')
    }

    return {
      onClick,
    }
  },
})
</script>
