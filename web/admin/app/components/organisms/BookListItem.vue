<template>
  <v-card outlined shaped @click="onClick">
    <v-img :src="book.thumbnailUrl" height="200px" contain />
    <v-card-title>{{ book.title }}</v-card-title>
    <v-card-subtitle>{{ book.author }}</v-card-subtitle>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { IBook } from '~/types/store'

export default defineComponent({
  props: {
    book: {
      type: Object as PropType<IBook>,
      required: false,
      default: () => ({}),
    },
  },

  setup(_, { emit }: SetupContext) {
    const onClick = () => {
      emit('click')
    }

    return {
      onClick,
    }
  },
})
</script>
