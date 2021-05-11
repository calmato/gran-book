<template>
  <debug :user-id="userId" :token="token" />
</template>

<script lang="ts">
import { computed, defineComponent, SetupContext } from '@nuxtjs/composition-api'
import Debug from '~/components/templates/Debug.vue'

export default defineComponent({
  components: {
    Debug,
  },

  setup(_, { root }: SetupContext) {
    const store = root.$store

    const userId = computed(() => store.getters['auth/getId'])
    const token = computed(() => store.getters['auth/getToken'])

    return {
      userId,
      token,
    }
  },
})
</script>
