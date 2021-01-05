<template>
  <v-btn @click="handleClick">ログインページへ</v-btn>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@vue/composition-api'

export default defineComponent({
  setup(_props, ctx: SetupContext) {
    const router = ctx.root.$router

    const handleClick = (): void => {
      router.push('/signin')
    }

    return {
      handleClick,
    }
  },
})
</script>
