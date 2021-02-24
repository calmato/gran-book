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
      type: String,
      required: false,
      default: '',
    },
  },

  setup(props, { emit }: SetupContext) {
    const selectedFile = async (file: File): Promise<void> => {
      if (checkFile(file)) {
        const picture = await getBase64(file)
        emit('input', picture)
      }
    }

    function checkFile(file: File): boolean {
      if (!file) {
        return false
      }

      return file.size <= props.limit * 1000 // KB -> MB
    }

    function getBase64(file: File): Promise<String | ArrayBuffer | null> {
      return new Promise((resolve, reject) => {
        const reader = new FileReader()
        reader.onload = (event: ProgressEvent<FileReader>) => resolve(event.target ? event.target.result : '')
        reader.onerror = (err: ProgressEvent<FileReader>) => reject(err)
        reader.readAsDataURL(file)
      })
    }

    return {
      selectedFile,
    }
  },
})
</script>
