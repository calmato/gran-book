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
                  <v-img :src="item.thumbnailUrl" />
                </v-avatar>
              </template>
              <template v-slot:[`item.name`]="{ item }">
                {{ item.name }}
              </template>
              <template v-slot:[`item.role`]="{ item }">
                <v-chip :color="getColor(item.role)" dark>
                  {{ getRole(item.role) }}
                </v-chip>
              </template>
              <template v-slot:[`item.actions`]="{ item }">
                <v-icon small class="mr-2" @click="editItem(item.index)">mdi-pencil</v-icon>
                <v-icon small @click="deleteItem(item.item)">mdi-delete</v-icon>
              </template>
            </v-data-table>
          </v-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, SetupContext, computed } from '@nuxtjs/composition-api'
import { AdminStore } from '~/store'
import { IAdminTableHeader, IAdminTableContent } from '~/types/props'
import { IAdminUser } from '~/types/store'

export default defineComponent({
  setup(_, { root }: SetupContext) {
    const store = root.$store

    const search = ''
    const headers: Array<IAdminTableHeader> = [
      { text: 'サムネ', value: 'thumbnailUrl', sortable: false },
      { text: '氏名', value: 'name', sortable: true },
      { text: 'メールアドレス', value: 'email', sortable: true },
      { text: '電話番号', value: 'phoneNumber', sortable: false },
      { text: '権限', value: 'role', sortable: true },
      { text: 'Actions', value: 'actions', sortable: false },
    ]

    const desserts = computed(() => {
      const users = store.getters['admin/getUsers']

      return users.map(
        (user: IAdminUser): IAdminTableContent => {
          const space: string = user.lastName && user.firstName ? ' ' : ''
          const name: string = user.lastName + space + user.firstName

          return {
            name,
            email: user.email,
            phoneNumber: user.phoneNumber,
            thumbnailUrl: user.thumbnailUrl || '/thumbnail.png',
            role: user.role,
          }
        }
      )
    })

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

    const newItem = async (): Promise<void> => {
      await AdminStore.createUser()
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
      getColor,
      getRole,
      newItem,
      editItem,
      deleteItem,
    }
  },
})
</script>
