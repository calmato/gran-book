<template>
  <v-navigation-drawer v-model="navigationDrawer" app clipped mobile-breakpoint="960" mini-variant-width="70">
    <!-- profile -->
    <v-sheet class="py-1 px-4">
      <v-list>
        <v-list-item two-line class="px-0">
          <v-list-item-avatar color="grey lighten-2">
            <v-img :src="thumbnailUrl" />
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title>{{ username }}</v-list-item-title>
            <v-list-item-subtitle class="caption">{{ email }}</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-sheet>
    <v-divider />
    <!-- links -->
    <v-list dense nav shaped>
      <v-list-item-group v-model="selectedItem" color="primary">
        <common-sidebar-list-item v-for="item in commonItems" :key="item.text" :item="item" @click="onClick" />
        <v-divider class="py-1" />
        <common-sidebar-list-item v-for="item in maintenanceItems" :key="item.text" :item="item" @click="onClick" />
        <v-divider class="py-1" />
        <common-sidebar-list-item v-for="item in developerItems" :key="item.text" :item="item" @click="onClick" />
        <v-divider class="py-1" />
        <common-sidebar-list-item v-for="item in systemItems" :key="item.text" :item="item" @click="onClick" />
      </v-list-item-group>
    </v-list>
    <!-- footer -->
    <template v-slot:append>
      <v-divider />
      <div align="center" class="pa-2">@2021 - Calmato</div>
    </template>
  </v-navigation-drawer>
</template>

<script lang="ts">
import { defineComponent, computed, SetupContext } from '@nuxtjs/composition-api'
import CommonSidebarListItem from '~/components/organisms/CommonSidebarListItem.vue'
import { ISidebarListItem } from '~/types/props'

export default defineComponent({
  components: {
    CommonSidebarListItem,
  },

  props: {
    current: {
      type: String,
      required: true,
      default: '',
    },
    drawer: {
      type: Boolean,
      required: true,
      default: true,
    },
    username: {
      type: String,
      required: false,
      default: 'Calmato 管理者',
    },
    email: {
      type: String,
      required: false,
      default: 'support@calmato.com',
    },
    thumbnailUrl: {
      type: String,
      required: false,
      default: '',
    },
  },

  setup(props, { emit }: SetupContext) {
    const { current } = props

    const commonItems: ISidebarListItem[] = [{ icon: 'mdi-home', text: 'ホーム', to: '/' }]
    const maintenanceItems: ISidebarListItem[] = [
      { icon: 'mdi-cart', text: 'お取り引き管理', to: '/' },
      { icon: 'mdi-forum', text: 'お問い合わせ管理', to: '/' },
      { icon: 'mdi-bell-ring', text: 'お知らせ管理', to: '/' },
      { icon: 'mdi-cash-100', text: 'セール情報管理', to: '/' },
    ]
    const developerItems: ISidebarListItem[] = [
      { icon: 'mdi-account', text: '利用者管理', to: '/' },
      { icon: 'mdi-book', text: '書籍管理', to: '/' },
      { icon: 'mdi-store', text: 'ECサイト管理', to: '/' },
    ]
    const systemItems: ISidebarListItem[] = [
      { icon: 'mdi-shield-account', text: '管理者管理', to: '/admin' },
      { icon: 'mdi-cog', text: 'システム設定', to: '/system' },
    ]

    const items: ISidebarListItem[] = commonItems.concat(maintenanceItems, developerItems, systemItems)
    const target: ISidebarListItem | undefined = items.filter((item: ISidebarListItem) => item.to === current).shift()
    const selectedItem: number = target ? items.indexOf(target) : -1

    const navigationDrawer = computed({
      get: () => props.drawer,
      set: (val: boolean) => emit('update:drawer', val),
    })

    const onClick = (link: string) => {
      emit('click', link)
    }

    return {
      navigationDrawer,
      selectedItem,
      commonItems,
      maintenanceItems,
      developerItems,
      systemItems,
      onClick,
    }
  },
})
</script>
