<template>
  <v-btn @click="handleClick">ログインページへ</v-btn>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  setup(_props, { root }: SetupContext) {
    const router = root.$router

    const handleClick = (): void => {
      router.push('/signin')
    }

    return {
      handleClick,
    }
  },
})
</script>
