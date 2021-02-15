<template>
  <v-alert v-show="show" :type="type" transition="scale-transition">
    <slot />
  </v-alert>
</template>

<script lang="ts">
import { defineComponent, watch, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  props: {
    type: {
      type: String,
      required: false,
      default: 'info', // info, success, warning, error
    },
    show: {
      type: Boolean,
      required: true,
      default: false,
    },
  },

  setup(props, { emit }: SetupContext) {
    watch(
      () => props.show,
      () => {
        if (props.show) {
          setTimeout(hiddenAlert, 5000)
        }
      }
    )

    const hiddenAlert = () => {
      emit('update:show', false)
    }

    return {
      hiddenAlert,
    }
  },
})
</script>
