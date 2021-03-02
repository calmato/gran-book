<template>
  <v-data-table
    :search="search"
    :loading="loading"
    :page="page"
    :items-per-page="itemsPerPage"
    :sort-by.sync="sortByArray"
    :sort-desc.sync="sortDescArray"
    :headers="headers"
    :items="items"
    :server-items-length="total"
    :footer-props="footers"
    @update:page="$emit('update:page', $event)"
    @update:items-per-page="$emit('update:items-per-page', $event)"
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
      <v-chip :color="getRoleColor(item.role)" dark>
        {{ getRole(item.role) }}
      </v-chip>
    </template>
    <template v-slot:[`item.actions`]="{ item }">
      <v-icon small class="mr-2" @click="onClickEdit(item)">mdi-pencil</v-icon>
      <v-icon small @click="onClickDelete(item)">mdi-delete</v-icon>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import { defineComponent, computed, SetupContext, PropType } from '@nuxtjs/composition-api'
import { IAdminTableContent, IAdminTableFooter, IAdminTableHeader } from '~/types/props'
import { IAdminUser } from '~/types/store'

export default defineComponent({
  props: {
    loading: {
      type: Boolean,
      required: false,
      default: false,
    },
    page: {
      type: Number,
      required: false,
      default: 1,
    },
    itemsPerPage: {
      type: Number,
      required: false,
      default: 20,
    },
    sortBy: {
      type: String,
      required: false,
      default: undefined,
    },
    sortDesc: {
      type: Boolean,
      required: false,
      default: undefined,
    },
    search: {
      type: String,
      required: false,
      default: '',
    },
    users: {
      type: Array as PropType<IAdminUser[]>,
      required: false,
      default: () => [],
    },
    total: {
      type: Number,
      required: false,
      default: 0,
    },
  },

  setup(props, { emit }: SetupContext) {
    const headers: Array<IAdminTableHeader> = [
      { text: 'サムネ', value: 'thumbnailUrl', sortable: false },
      { text: '氏名', value: 'name', sortable: false },
      { text: 'メールアドレス', value: 'email', sortable: true },
      { text: '電話番号', value: 'phoneNumber', sortable: false },
      { text: '権限', value: 'role', sortable: true },
      { text: 'Actions', value: 'actions', sortable: false },
    ]
    const footers: IAdminTableFooter = {
      itemsPerPageOptions: [10, 20, 30, 50, 100],
    }
    const items = computed((): IAdminTableContent[] => {
      return props.users.map(
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

    const sortByArray = computed({
      get: (): string[] => [props.sortBy],
      set: (vals: string[]) => emit('update:sort-by', vals[0]),
    })

    const sortDescArray = computed({
      get: (): boolean[] => [props.sortDesc],
      set: (vals: boolean[]) => emit('update:sort-desc', vals[0]),
    })

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

    const getRoleColor = (role: number): string => {
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

    const onClickEdit = (item: IAdminTableContent): void => {
      const index: number = items.value.indexOf(item)
      emit('edit', index)
    }

    const onClickDelete = (item: IAdminTableContent): void => {
      const index: number = items.value.indexOf(item)
      emit('delete', index)
    }

    return {
      headers,
      footers,
      items,
      sortByArray,
      sortDescArray,
      getRole,
      getRoleColor,
      onClickEdit,
      onClickDelete,
    }
  },
})
</script>
