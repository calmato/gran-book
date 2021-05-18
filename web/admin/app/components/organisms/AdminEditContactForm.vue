<template>
  <v-form class="pa-4">
    <the-text-field v-model="form.params.email" :label="form.options.email.label" :rules="form.options.email.rules" />
    <the-text-field v-model="form.params.phoneNumber" :label="form.options.phoneNumber.label" :rules="form.options.phoneNumber.rules" />
    <v-btn color="primary" class="mt-4 mr-4" @click="onSubmit">変更する</v-btn>
    <v-btn class="mt-4" @click="onCancel">キャンセル</v-btn>
  </v-form>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheTextField from '~/components/atoms/TheTextField.vue'
import { IAdminEditContactForm } from '~/types/forms'

export default defineComponent({
  components: {
    TheTextField,
  },

  props: {
    form: {
      type: Object as PropType<IAdminEditContactForm>,
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
