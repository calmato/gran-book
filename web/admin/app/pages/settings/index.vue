<template>
  <settings
    :username="username"
    :name="name"
    :name-kana="nameKana"
    :thumbnail-url="thumbnailUrl"
    :self-introduction="selfIntroduction"
    :phone-number="phoneNumber"
    :email="email"
    @click="handleClick"
  />
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'
import Settings from '~/components/templates/Settings.vue'

export default defineComponent({
  components: {
    Settings,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const username = store.getters['auth/getUsername']
    const name = store.getters['auth/getName']
    const nameKana = store.getters['auth/getNameKana']
    const thumbnailUrl = store.getters['auth/getThumbnailUrl']
    const selfIntroduction = store.getters['auth/getSelfIntroduction']
    const phoneNumber = store.getters['auth/getPhoneNumber']
    const email = store.getters['auth/getEmail']

    const handleClick = (path: string) => {
      router.push(path)
    }

    return {
      username,
      name,
      nameKana,
      thumbnailUrl,
      selfIntroduction,
      phoneNumber,
      email,
      handleClick,
    }
  },
})
</script>
