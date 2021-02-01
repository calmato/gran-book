<template>
  <v-list class="pa-4">
    <div v-for="list in lists" :key="list.title">
      <v-list-item>
        <v-list-item-content>
          <v-list-item-subtitle>{{ list.title }}</v-list-item-subtitle>
        </v-list-item-content>
        <v-list-item-content v-if="list.contentType === 'image'">
          <v-img :src="list.content" max-height="120" max-width="120" />
        </v-list-item-content>
        <v-list-item-content v-else>
          {{ list.content }}
        </v-list-item-content>
      </v-list-item>
      <v-divider />
    </div>
    <div class="text-center">
      <v-btn color="primary mt-4" block @click="onClick">編集する</v-btn>
    </div>
  </v-list>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { IProfileList } from '~/types/props'

export default defineComponent({
  props: {
    lists: {
      type: Array as PropType<IProfileList[]>,
      required: false,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
    const onClick = () => {
      emit('click')
    }

    return {
      onClick,
    }
  }
})
</script>
