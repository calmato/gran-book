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
            <v-data-table
              :search="search"
              :loading="loading"
              :page.sync="page"
              :items-per-page.sync="itemsPerPage"
              :sort-by.sync="sortBy"
              :sort-desc.sync="sortDesc"
              :headers="headers"
              :items="desserts"
              :server-items-length="total"
              :footer-props="footerProps"
            >
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
                <v-icon small @click="deleteItem(item.index)">mdi-delete</v-icon>
              </template>
            </v-data-table>
          </v-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, SetupContext, ref, useAsync, computed, watch } from '@nuxtjs/composition-api'
import { AdminStore } from '~/store'
import { IAdminListForm } from '~/types/forms'
import { IAdminTableHeader, IAdminTableContent } from '~/types/props'
import { IAdminUser } from '~/types/store'

export default defineComponent({
  setup(_, { root }: SetupContext) {
    const store = root.$store

    const headers: Array<IAdminTableHeader> = [
      { text: 'サムネ', value: 'thumbnailUrl', sortable: false },
      { text: '氏名', value: 'name', sortable: true },
      { text: 'メールアドレス', value: 'email', sortable: true },
      { text: '電話番号', value: 'phoneNumber', sortable: false },
      { text: '権限', value: 'role', sortable: true },
      { text: 'Actions', value: 'actions', sortable: false },
    ]
    const search = ''
    const loading = ref<boolean>(false)
    const footerProps = { itemsPerPageOptions: [20, 30, 50] }
    const page = ref<number>(1)
    const itemsPerPage = ref<number>(20)
    const sortBy = ref<string>('')
    const sortDesc = ref<boolean>(true)

    const total = computed(() => store.getters['admin/getTotal'])
    const desserts = computed((): IAdminTableContent[] => {
      const users: IAdminUser[] = store.getters['admin/getUsers']

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

    watch(page, (): void => {
      indexAdmin()
    })

    watch(itemsPerPage, (): void => {
      indexAdmin()
    })

    watch(sortBy, (): void => {
      indexAdmin()
    })

    watch(sortDesc, (): void => {
      indexAdmin()
    })

    useAsync(async () => {
      await indexAdmin()
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

    const editItem = (index: number): void => {
      console.log('debug', 'editItem', index)
    }

    const deleteItem = (index: number): void => {
      console.log('debug', 'deleteItem', index)
    }

    async function indexAdmin(): Promise<void> {
      loading.value = true

      const form: IAdminListForm = {
        limit: itemsPerPage.value,
        offset: itemsPerPage.value * (page.value - 1),
        order: {
          by: sortBy.value,
          desc: sortDesc.value,
        },
      }

      await AdminStore.indexAdmin(form).finally(() => {
        loading.value = false
      })
    }

    return {
      loading,
      search,
      footerProps,
      page,
      itemsPerPage,
      sortBy,
      sortDesc,
      headers,
      total,
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
