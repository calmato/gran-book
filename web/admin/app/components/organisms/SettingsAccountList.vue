<template>
  <v-list class="pa-4">
    <div v-for="list in lists" :key="list.title">
      <v-list-item>
        <v-list-item-content class="col col-3">
          <v-list-item-subtitle>{{ list.title }}</v-list-item-subtitle>
        </v-list-item-content>
        <v-list-item-content>{{ list.content }}</v-list-item-content>
        <v-list-item-action>
          <v-btn color="primary" @click="onClick(list.to)">変更する</v-btn>
        </v-list-item-action>
      </v-list-item>
      <v-divider />
    </div>
  </v-list>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { IAccountList } from '~/types/props'

export default defineComponent({
  props: {
    lists: {
      type: Array as PropType<IAccountList[]>,
      required: false,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
    const onClick = (path: string) => {
      emit('click', path)
    }

    return {
      onClick,
    }
  },
})
</script>
