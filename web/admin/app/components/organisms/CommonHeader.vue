<template>
  <v-app-bar app clipped-left dark color="primary" class="px-2">
    <!-- left -->
    <v-toolbar-title class="align-center d-flex">
      <v-img src="/logo.png" max-width="24" contain />
      <span class="ml-4">Gran Book</span>
      <v-app-bar-nav-icon class="ml-4" @click="onChange" />
      <v-icon class="ml-4">mdi-magnify</v-icon>
    </v-toolbar-title>
    <v-spacer />
    <!-- right -->
    <v-icon class="mr-6">mdi-bell-ring</v-icon>
    <v-menu rounded offset-y transition="scroll-x-reverse-transition">
      <template v-slot:activator="{ on, attrs }">
        <v-avatar v-bind="attrs" size="40" color="grey lighten-2" v-on="on">
          <v-img :src="thumbnailUrl ? thumbnailUrl : '/thumbnail.png'" />
        </v-avatar>
      </template>
      <v-list dense>
        <v-list-item v-for="(item, i) in items" :key="i" link>
          <v-list-item-title @click="onClick(item.to)">{{ item.text }}</v-list-item-title>
        </v-list-item>
        <v-list-item link>
          <v-list-item-title @click="onClickLogout">ログアウト</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-app-bar>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'
import { IHeaderListItem } from '~/types/props'

export default defineComponent({
  props: {
    thumbnailUrl: {
      type: String,
      required: false,
      default: '',
    },
  },

  setup(_, { emit }: SetupContext) {
    const items: IHeaderListItem[] = [{ text: '設定', to: '/settings' }]

    const onClick = (link: string): void => {
      emit('click', link)
    }

    const onClickLogout = (): void => {
      emit('logout')
    }

    const onChange = (): void => {
      emit('change')
    }

    return {
      items,
      onClick,
      onClickLogout,
      onChange,
    }
  },
})
</script>
