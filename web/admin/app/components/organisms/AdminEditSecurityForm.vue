<template>
  <v-form class="pa-4">
    <the-text-field
      v-model="form.params.password"
      :label="form.options.password.label"
      :rules="form.options.password.rules"
      type="password"
    />
    <the-text-field
      v-model="form.params.passwordConfirmation"
      :label="form.options.passwordConfirmation.label"
      :rules="form.options.passwordConfirmation.rules"
      type="password"
    />
    <v-btn color="primary" class="mt-4 mr-4" @click="onSubmit">変更する</v-btn>
    <v-btn class="mt-4" @click="onCancel">キャンセル</v-btn>
  </v-form>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheTextField from '~/components/atoms/TheTextField.vue'
import { IAdminEditSecurityForm } from '~/types/forms'

export default defineComponent({
  components: {
    TheTextField,
  },

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
