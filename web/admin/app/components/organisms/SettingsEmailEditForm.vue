<template>
  <v-form class="py-8">
    <v-text-field v-model="form.email" label="email" autofocus />
    <v-btn :block="true" color="primary" classt="mt-4" @click="onClick">確認メールを送信</v-btn>
  </v-form>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { ISettingsEmailEditForm } from '~/types/forms'

export default defineComponent({
  props: {
    form: {
      type: Object as PropType<ISettingsEmailEditForm>,
      required: true,
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
