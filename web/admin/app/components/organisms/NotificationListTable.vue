<template>
  <v-data-table
    :search="search"
    :loading="loading"
    :page="page"
    :items-per-page="itemsPerPage"
    :sort-by.sync="sortByArray"
    :sort-desc.sync="sortDescArray"
    :headers="headers"
    :items="item"
    :footer-props="footers"
    @update:page="$emit('update:page', $event)"
    @update:items-per-page="$emit('update:items-per-page', $event)"
  >
    <template v-slot:[`item.title`]="{ item }">
      {{ item.title }}
    </template>
    <template v-slot:[`item.description`]="{ item }">
      {{ item.description }}
    </template>
    <template v-slot:[`item.timestamp`]="{ item }">
      {{ item.timestamp }}
    </template>
    <template v-slot:[`item.category`]="{ item }">
      <v-chip :color="getColor(item.category)" dark>
        {{ getCategory(item.category) }}
      </v-chip>
    </template>
    <template v-slot:[`item.importance`]="{ item }">
      <v-chip :color="getColor(item.importance)" dark>
        {{ getImportance(item.importance) }}
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
import { INotificationTableContent, IAdminTableFooter, INotificationTableHeader } from '~/types/props'
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
    const headers: Array<INotificationTableHeader> = [
      { text: 'タイトル', value: 'title', sortable: false },
      { text: '説明', value: 'description', sortable: false },
      { text: '作成日', value: 'timestamp', sortable: true },
      { text: 'カテゴリー', value: 'category', sortable: false },
      { text: '重要度', value: 'importance', sortable: false },
      { text: 'Actions', value: 'actions', sortable: false },
    ]

    const item: Array<INotificationTableContent> = [
      {
        title: 'Gran book公開日',
        description: 'Gran book公開に伴ってキャンペーンの開催',
        timestamp: '2022/01/01',
        category: 3,
        importance: 2,
      },
      {
        title: 'おすすめ本情報',
        description: '呪術廻戦１３巻販売開始',
        timestamp: '2022/01/03',
        category: 2,
        importance: 3,
      },
      {
        title: 'メンテナンス情報',
        description: 'メンテナンスに伴うサービスの一時停止',
        timestamp: '2022/01/02',
        category: 1,
        importance: 1,
      },
    ]

    const footers: IAdminTableFooter = {
      itemsPerPageOptions: [10, 20, 30, 50, 100],
    }
    const items = computed((): INotificationTableContent[] => {
      return props.users.map(
        (user: IAdminUser): INotificationTableContent => {
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

    const getCategory = (role: number): string => {
      switch (role) {
        case 1:
          return 'Maintenance'
        case 2:
          return 'Event'
        case 3:
          return 'Information'
        default:
          return 'Unknown'
      }
    }

    const getImportance = (role: number): string => {
      switch (role) {
        case 1:
          return 'High'
        case 2:
          return 'Middle'
        case 3:
          return 'Low'
        default:
          return 'Unknown'
      }
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

    const onClickEdit = (item: INotificationTableContent): void => {
      const index: number = items.value.indexOf(item)
      emit('edit', index)
    }

    const onClickDelete = (item: INotificationTableContent): void => {
      const index: number = items.value.indexOf(item)
      emit('delete', index)
    }

    return {
      headers,
      footers,
      items,
      item,
      sortByArray,
      sortDescArray,
      getCategory,
      getImportance,
      getColor,
      onClickEdit,
      onClickDelete,
    }
  },
})
</script>
