<template>
  <validation-provider v-slot="{ errors, valid }" :name="label" :vid="name" :rules="rules">
    <v-textarea
      v-model="formData"
      :label="label"
      :error-messages="errors"
      :success="valid"
      :autofocus="autofocus"
      :outlined="outlined"
    />
  </validation-provider>
</template>

<script lang="ts">
import { defineComponent, computed, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  props: {
    autofocus: {
      type: Boolean,
      required: false,
      default: false,
    },
    label: {
      type: String,
      required: false,
      default: '',
    },
    name: {
      type: String,
      required: false,
      default: '',
    },
    outlined: {
      type: Boolean,
      required: false,
      default: false,
    },
    rules: {
      type: Object,
      required: false,
      default: () => ({}),
    },
    value: {
      type: String,
      required: false,
      default: '',
    },
  },

  setup(props, { emit }: SetupContext) {
    const formData = computed({
      get: () => props.value,
      set: (val: string) => emit('input', val),
    })

    return {
      formData,
    }
  },
})
</script>
