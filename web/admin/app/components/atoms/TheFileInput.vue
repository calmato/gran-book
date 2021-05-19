<template>
  <validation-provider v-slot="{ errors, valid }" :name="label" :vid="name" :rules="rules">
    <v-file-input
      :accept="accept"
      :label="label"
      :error-messages="errors"
      :success="valid"
      :value="file"
      show-size
      @change="onSelect"
    />
  </validation-provider>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  props: {
    accept: {
      type: String,
      required: false,
      default: 'image/*',
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
    rules: {
      type: Object,
      required: false,
      default: () => ({}),
    },
    file: {
      type: File,
      required: false,
      default: null,
    },
  },

  setup(_, { emit }: SetupContext) {
    const onSelect = (file: File) => {
      emit('change', file)
    }

    return {
      onSelect,
    }
  },
})
</script>
