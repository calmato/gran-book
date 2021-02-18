<template>
  <home />
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'
import Home from '~/components/templates/Home.vue'

export default defineComponent({
  components: {
    Home,
  },
})
</script>
