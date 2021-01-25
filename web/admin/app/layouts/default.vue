<template>
  <v-app>
    <common-header />
    <common-sidebar :current="current" @click="handleClick" />
    <v-main>
      <nuxt />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'
import CommonHeader from '~/components/organisms/CommonHeader.vue'
import CommonSidebar from '~/components/organisms/CommonSidebar.vue'

export default defineComponent({
  components: {
    CommonHeader,
    CommonSidebar,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const current = root.$route.path

    const handleClick = (link: string): void => {
      router.push(link)
    }

    return {
      current,
      handleClick,
    }
  },
})
</script>

<style scoped>
.application {
  font-family: Roboto, sans-serif;
}

.theme--light .v-main {
  background-color: #eef5f9;
}
</style>
