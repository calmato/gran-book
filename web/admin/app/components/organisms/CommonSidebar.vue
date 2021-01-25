<template>
  <v-navigation-drawer id="main-sidebar" v-model="drawer" clipped mini-variant-width="70" app>
    <!-- profile -->
    <v-sheet class="py-1 px-4">
      <v-list>
        <v-list-item two-line class="px-0">
          <v-list-item-avatar color="grey lighten-2">
            <v-img src="/thumbnail.png" />
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title>Calmato 管理者</v-list-item-title>
            <v-list-item-subtitle class="caption">support@calmato.com</v-list-item-subtitle>
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
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'
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
  },

  setup(props, { emit }: SetupContext) {
    const { current } = props

    const commonItems: ISidebarListItem[] = [{ icon: 'mdi-home', text: 'ホーム', link: '/' }]
    const maintenanceItems: ISidebarListItem[] = [
      { icon: 'mdi-cart', text: 'お取り引き管理', link: '/' },
      { icon: 'mdi-forum', text: 'お問い合わせ管理', link: '/' },
      { icon: 'mdi-bell-ring', text: 'お知らせ管理', link: '/' },
      { icon: 'mdi-cash-100', text: 'セール情報管理', link: '/' },
    ]
    const developerItems: ISidebarListItem[] = [
      { icon: 'mdi-account', text: '利用者管理', link: '/' },
      { icon: 'mdi-book', text: '書籍管理', link: '/' },
      { icon: 'mdi-store', text: 'ECサイト管理', link: '/' },
    ]
    const systemItems: ISidebarListItem[] = [
      { icon: 'mdi-shield-account', text: '管理者管理', link: '/' },
      { icon: 'mdi-cog', text: 'システム設定', link: '/system' },
    ]

    // list item内でactiveになってる箇所をPathから判定
    const items: ISidebarListItem[] = commonItems.concat(maintenanceItems, developerItems, systemItems)
    const target: ISidebarListItem | undefined = items
      .filter((item: ISidebarListItem) => {
        return item.link === current
      })
      .shift()

    const selectedItem: number = target ? items.indexOf(target) : -1

    const onClick = (link: string) => {
      emit('click', link)
    }

    return {
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
