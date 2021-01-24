<template>
  <v-app>
    <v-app-bar app clipped-left clipped-right color="primary" dark class="px-2">
      <v-toolbar-title class="align-center d-flex">
        <v-icon>mdi-menu</v-icon>
        <span class="ml-4">Gran Book</span>
      </v-toolbar-title>
      <v-spacer />
      <v-icon class="mr-6">mdi-bell-ring</v-icon>
      <v-avatar size="32" color="grey lighten-2">
        <v-img src="/thumbnail.png" />
      </v-avatar>
    </v-app-bar>

    <v-navigation-drawer id="main-sidebar" v-model="drawer" clipped mini-variant-width="70" app>
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
      <v-list dense nav>
        <v-list-item v-for="[icon, text] in commonLinks" :key="icon" link>
          <v-list-item-icon>
            <v-icon>{{ icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{ text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-divider class="py-1" />
        <v-list-item v-for="[icon, text] in maintenanceLinks" :key="icon" link>
          <v-list-item-icon>
            <v-icon>{{ icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{ text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-divider class="py-1" />
        <v-list-item v-for="[icon, text] in developerLinks" :key="icon" link>
          <v-list-item-icon>
            <v-icon>{{ icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{ text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-divider class="py-1" />
        <v-list-item v-for="[icon, text] in systemLinks" :key="icon" link>
          <v-list-item-icon>
            <v-icon>{{ icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{ text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <template v-slot:append>
        <v-divider />
        <div align="center" class="pa-2">@2021 - Calmato</div>
      </template>
    </v-navigation-drawer>

    <v-main>
      <nuxt />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'

export default defineComponent({
  setup(_) {
    const commonLinks: string[][] = [['mdi-home', 'ホーム']]
    const maintenanceLinks: string[][] = [
      ['mdi-cart', 'お取り引き管理'],
      ['mdi-forum', 'お問い合わせ管理'],
      ['mdi-bell-ring', 'お知らせ管理'],
      ['mdi-cash-100', 'セール情報管理'],
    ]
    const developerLinks: string[][] = [
      ['mdi-account', '利用者管理'],
      ['mdi-book', '書籍管理'],
      ['mdi-store', 'ECサイト管理'],
    ]
    const systemLinks: string[][] = [
      ['mdi-shield-account', '管理者管理'],
      ['mdi-cog', 'システム設定'],
    ]

    return {
      commonLinks,
      maintenanceLinks,
      developerLinks,
      systemLinks,
    }
  },
})
</script>

<style scoped>
.theme--light .v-main {
  background-color: #eef5f9;
}
</style>
