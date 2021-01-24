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
import { ISidebarLink } from '~/types/props'

export default defineComponent({
  components: {
    CommonSidebarListItem,
  },

  setup(_, { emit }: SetupContext) {
    const selectedItem: any = undefined
    const commonItems: ISidebarLink[] = [{ icon: 'mdi-home', text: 'ホーム', link: '/' }]
    const maintenanceItems: ISidebarLink[] = [
      { icon: 'mdi-cart', text: 'お取り引き管理', link: '/' },
      { icon: 'mdi-forum', text: 'お問い合わせ管理', link: '/' },
      { icon: 'mdi-bell-ring', text: 'お知らせ管理', link: '/' },
      { icon: 'mdi-cash-100', text: 'セール情報管理', link: '/' },
    ]
    const developerItems: ISidebarLink[] = [
      { icon: 'mdi-account', text: '利用者管理', link: '/' },
      { icon: 'mdi-book', text: '書籍管理', link: '/' },
      { icon: 'mdi-store', text: 'ECサイト管理', link: '/' },
    ]
    const systemItems: ISidebarLink[] = [
      { icon: 'mdi-shield-account', text: '管理者管理', link: '/' },
      { icon: 'mdi-cog', text: 'システム設定', link: '/system' },
    ]

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
