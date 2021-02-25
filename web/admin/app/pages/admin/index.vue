<template>
  <v-container fill-height>
    <v-layout wrap>
      <v-row>
        <v-col cols="12">
          <v-card class="pa-4">
            <v-subheader>管理者一覧</v-subheader>
            <v-card-title>
              <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details />
              <v-spacer />
              <v-btn color="primary" dark class="mb-2" @click="newItem">New Item</v-btn>
            </v-card-title>
            <v-data-table :headers="headers" :items="desserts" :search="search">
              <template v-slot:[`item.thumbnailUrl`]="{ item }">
                <v-avatar>
                  <v-img :src="getThumbnailUrl(item.thumbnailUrl)" />
                </v-avatar>
              </template>
              <template v-slot:[`item.role`]="{ item }">
                <v-chip :color="getColor(item.role)" dark>
                  {{ getRole(item.role) }}
                </v-chip>
              </template>
              <template v-slot:[`item.actions`]="{ item }">
                <v-icon small class="mr-2" @click="editItem(item)">mdi-pencil</v-icon>
                <v-icon small @click="deleteItem(item)">mdi-delete</v-icon>
              </template>
            </v-data-table>
          </v-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, reactive } from '@nuxtjs/composition-api'

export default defineComponent({
  setup(_) {
    const search = ''
    const headers: Array<{
      text: string
      value: string
      sortable: boolean
    }> = [
      { text: 'サムネ', value: 'thumbnailUrl', sortable: false },
      { text: '氏名', value: 'name', sortable: true },
      { text: 'メールアドレス', value: 'email', sortable: true },
      { text: '電話番号', value: 'phoneNumber', sortable: false },
      { text: '権限', value: 'role', sortable: true },
      { text: 'Actions', value: 'actions', sortable: false },
    ]
    const desserts = reactive<
      Array<{
        name: string
        email: string
        phoneNumber: string
        role: number
      }>
    >([
      {
        name: '濵田 広大',
        email: 'k.hamada@calmato.com',
        phoneNumber: '000-0000-0000',
        role: 1,
      },
      {
        name: '西川 直志',
        email: 't.nishikawa@calmato.com',
        phoneNumber: '111-1111-1111',
        role: 1,
      },
      {
        name: '山田 侑樹',
        email: 'y.yamada@calmato.com',
        phoneNumber: '000-1111-1111',
        role: 2,
      },
      {
        name: 'Inatomi Atsuhide',
        email: 'a.inatomi@calmato.com',
        phoneNumber: '111-0000-0000',
        role: 3,
      },
    ])

    const getThumbnailUrl = (thumbnailUrl: string): string => {
      return thumbnailUrl || '/thumbnail.png'
    }

    const getColor = (role: number): string => {
      switch (role) {
        case 1:
          return 'red'
        case 2:
          return 'orange'
        case 3:
          return 'green'
        default:
          return ''
      }
    }

    const getRole = (role: number): string => {
      switch (role) {
        case 1:
          return 'Administrator'
        case 2:
          return 'Developer'
        case 3:
          return 'Operator'
        default:
          return 'Unknown'
      }
    }

    const newItem = (): void => {
      console.log('debug', 'newItem')
      desserts.push({
        name: 'new user',
        email: 'u.new@calmato.com',
        phoneNumber: '123-1234-1234',
        role: 3,
      })
    }

    const editItem = (item: { name: string; email: string; phoneNumber: string; role: number }): void => {
      console.log('debug', 'editItem', item)
    }

    const deleteItem = (item: { name: string; email: string; phoneNumber: string; role: number }): void => {
      console.log('debug', 'deleteItem', item)
    }

    return {
      search,
      headers,
      desserts,
      getThumbnailUrl,
      getColor,
      getRole,
      newItem,
      editItem,
      deleteItem,
    }
  },
})
</script>
