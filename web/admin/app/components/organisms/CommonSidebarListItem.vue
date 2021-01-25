<template>
  <v-list-item link @click="onClick">
    <v-list-item-icon>
      <v-icon>{{ item.icon }}</v-icon>
    </v-list-item-icon>
    <v-list-item-content>
      <v-list-item-title>{{ item.text }}</v-list-item-title>
    </v-list-item-content>
  </v-list-item>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { ISidebarListItem } from '~/types/props'

export default defineComponent({
  props: {
    item: {
      type: Object as PropType<ISidebarListItem>,
      required: false,
      default: () => ({ icon: '', text: '', link: '' }),
    },
  },

  setup(props, { emit }: SetupContext) {
    const { item } = props

    const onClick = () => {
      emit('click', item.link)
    }

    return {
      item,
      onClick,
    }
  },
})
</script>
