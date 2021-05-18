<template>
  <v-form class="pa-4">
    <v-text-field v-model="form.params.password" type="password" label="パスワード" />
    <v-text-field v-model="form.params.passwordConfirmation" type="password" label="パスワード(確認用)" />
    <v-btn color="primary" class="mt-4 mr-4" @click="onSubmit">変更する</v-btn>
    <v-btn class="mt-4" @click="onCancel">キャンセル</v-btn>
  </v-form>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { IAdminEditSecurityForm } from '~/types/forms'

export default defineComponent({
  props: {
    form: {
      type: Object as PropType<IAdminEditSecurityForm>,
      required: false,
      default: () => ({}),
    },
  },

  setup(_, { emit }: SetupContext) {
    const onSubmit = (): void => {
      emit('submit')
    }

    const onCancel = (): void => {
      emit('cancel')
    }

    return {
      onSubmit,
      onCancel,
    }
  },
})
</script>
