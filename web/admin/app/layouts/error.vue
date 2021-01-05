<template>
  <v-app dark>
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
import { defineComponent, SetupContext } from '@vue/composition-api'

export default defineComponent({
  layout: 'empty',

  props: {
    error: {
      type: Object,
      default: null,
    },
  },

  head() {
    const title = this.error.statusCode === 404 ? this.pageNotFound : this.otherError
    return {
      title,
    }
  },

  setup(props, ctx: SetupContext) {
    const pageNotFound: string = '404 Not Found'
    const otherError: string = 'An error occurred'

    return {
      pageNotFound,
      otherError
    }
  }
})
</script>

<style scoped>
h1 {
  font-size: 20px;
}
</style>
