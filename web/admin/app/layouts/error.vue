<template>
  <v-app>
    <h1 v-if="error.statusCode === 404">
      {{ pageNotFound }}
    </h1>
    <h1 v-else>
      {{ otherError }}
    </h1>
    <NuxtLink to="/"> Home page </NuxtLink>
  </v-app>
</template>

<script lang="ts">
import { defineComponent, useMeta } from '@nuxtjs/composition-api'

export default defineComponent({
  layout: 'empty',
  head: {},

  props: {
    error: {
      type: Object,
      default: null,
    },
  },

  setup(props) {
    const { error } = props
    const pageNotFound: string = '404 Not Found'
    const otherError: string = 'An error occurred'

    useMeta(() => ({
      title: error.statusCode === 404 ? pageNotFound : otherError,
    }))

    return {
      pageNotFound,
      otherError,
    }
  },
})
</script>

<style scoped>
h1 {
  font-size: 20px;
}
</style>
