<template>
  <v-app>
    <common-snackbar :snackbar.sync="snackbar" :color="snackbarColor" :message="snackbarMessage" />
    <common-header :thumbnail-url="thumbnailUrl" @click="handleClick" @change="handleChange" @logout="handleLogout" />
    <common-sidebar
      :drawer.sync="drawer"
      :username="username"
      :email="email"
      :thumbnail-url="thumbnailUrl"
      :current="current"
      @click="handleClick"
    />
    <v-main>
      <nuxt />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent, ref, computed, watch, SetupContext } from '@nuxtjs/composition-api'
import { AuthStore, CommonStore } from '~/store'
import CommonHeader from '~/components/organisms/CommonHeader.vue'
import CommonSidebar from '~/components/organisms/CommonSidebar.vue'
import CommonSnackbar from '~/components/organisms/CommonSnackbar.vue'

export default defineComponent({
  components: {
    CommonHeader,
    CommonSidebar,
    CommonSnackbar,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const current = root.$route.path
    const store = root.$store

    const drawer = ref<Boolean>(true)
    const snackbar = ref<Boolean>(false)

    const username = computed(() => store.getters['auth/getUsername'])
    const email = computed(() => store.getters['auth/getEmail'])
    const thumbnailUrl = computed(() => store.getters['auth/getThumbnailUrl'])
    const snackbarColor = computed(() => store.getters['common/getSnackbarColor'])
    const snackbarMessage = computed(() => store.getters['common/getSnackbarMessage'])

    watch(snackbarMessage, (): void => {
      snackbar.value = snackbarMessage.value !== ''
    })

    watch(snackbar, (): void => {
      if (!snackbar.value) {
        CommonStore.hiddenSnackbar()
      }
    })

    const handleClick = (link: string): void => {
      router.push(link)
    }

    const handleLogout = (): void => {
      AuthStore.logout()
      router.push('/signin')
    }

    const handleChange = (): void => {
      drawer.value = !drawer.value
    }

    return {
      current,
      drawer,
      snackbar,
      snackbarColor,
      snackbarMessage,
      username,
      email,
      thumbnailUrl,
      handleClick,
      handleLogout,
      handleChange,
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
