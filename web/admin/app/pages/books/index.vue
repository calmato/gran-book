<template>
  <book-list />
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'
import BookList from '~/components/templates/BookList.vue'

export default defineComponent({
  components: {
    BookList,
  },

  setup() {},
})
</script>
