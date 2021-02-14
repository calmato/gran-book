<template>
  <v-snackbar v-model="isShow" :color="color" top>
    {{ message }}
    <template v-slot:action="{ attrs }">
      <v-btn v-bind="attrs" icon @click="isShow = false">
        <v-icon>mdi-close</v-icon>
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script lang="ts">
import { defineComponent, computed, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  props: {
    color: {
      type: String,
      required: false,
      default: 'info',
    },
    message: {
      type: String,
      required: false,
      default: '',
    },
    snackbar: {
      type: Boolean,
      required: false,
      default: false,
    },
  },

  setup(props, { emit }: SetupContext) {
    const isShow = computed({
      get: () => props.snackbar,
      set: (val: Boolean) => emit('update:snackbar', val),
    })

    return {
      isShow,
    }
  },
})
</script>
