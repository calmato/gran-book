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
              <v-btn color="primary" dark class="mb-2" @click="onClickNewButton">New Item</v-btn>
            </v-card-title>
            <admin-list-table
              :loading="loading"
              :page="page"
              :items-per-page="itemsPerPage"
              :sort-by="sortBy"
              :sort-desc="sortDesc"
              :search="search"
              :users="users"
              :total="total"
              @edit="onClickEditButton"
              @delete="onClickDeleteButton"
              @update:page="$emit('update:page', $event)"
              @update:items-per-page="$emit('update:items-per-page', $event)"
              @update:sort-by="$emit('update:sort-by', $event)"
              @update:sort-desc="$emit('update:sort-desc', $event)"
            />
          </v-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import AdminListTable from '~/components/organisms/AdminListTable.vue'
import { IAdminUser } from '~/types/store'

export default defineComponent({
  components: {
    AdminListTable,
  },

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

  setup(_, { emit }: SetupContext) {
    const onClickNewButton = (): void => {
      emit('new')
    }

    const onClickEditButton = (index: number): void => {
      emit('edit', index)
    }

    const onClickDeleteButton = (index: number): void => {
      emit('delete', index)
    }

    return {
      onClickNewButton,
      onClickEditButton,
      onClickDeleteButton,
    }
  },
})
</script>
