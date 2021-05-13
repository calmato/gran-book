<template>
  <validation-provider v-slot="{ errors, valid }" :name="label" :vid="name" :rules="rules">
    <v-file-input
      :accept="accept"
      :label="label"
      :error-messages="errors"
      :success="valid"
      show-size
      @change="selectedFile"
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
    limit: {
      type: Number,
      required: false,
      default: 10000000, // 10MB
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
    value: {
      type: File,
      required: false,
      default: undefined,
    },
  },

  setup(props, { emit }: SetupContext) {
    const selectedFile = (file: File): void => {
      if (checkFile(file)) {
        emit('input', file)
      }
    }

    function checkFile(file: File): boolean {
      if (!file) {
        return true
      }

      return file.size <= props.limit * 1000 // KB -> MB
    }

    return {
      selectedFile,
    }
  },
})
</script>
