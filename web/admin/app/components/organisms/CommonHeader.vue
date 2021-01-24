<template>
  <v-app-bar app clipped-left clipped-right color="primary" dark class="px-2">
    <v-toolbar-title class="align-center d-flex">
      <v-img src="/logo.png" max-width="24" contain />
      <span class="ml-4">Gran Book</span>
      <v-icon class="ml-4">mdi-menu</v-icon>
      <v-icon class="ml-4">mdi-magnify</v-icon>
    </v-toolbar-title>
    <v-spacer />
    <v-icon class="mr-6">mdi-bell-ring</v-icon>
    <v-avatar size="40" color="grey lighten-2">
      <v-img :src="thumbnail ? thumbnail : '/thumbnail.png'" />
    </v-avatar>
  </v-app-bar>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'

export default defineComponent({
  setup(_) {
    const thumbnail: string = ''

    return {
      thumbnail,
    }
  },
})
</script>
