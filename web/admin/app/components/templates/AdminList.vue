<template>
  <v-container fill-height>
    <v-dialog v-model="newDialog" width="600px" scrollable @click:outside="$emit('update:new-dialog', false)">
      <v-card>
        <v-toolbar color="primary" dark>管理者ユーザー 追加</v-toolbar>
        <v-card-text>
          <admin-new-form :form="newForm" />
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="onClickNewClose"> Close </v-btn>
          <v-btn color="blue darken-1" text :loading="loading" :disabled="loading" @click="onClickCreateButton">
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-layout wrap>
      <v-row>
        <v-col cols="12">
          <v-card class="pa-4">
            <v-subheader>管理者一覧</v-subheader>
            <v-card-title>
              <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details />
              <v-spacer />
              <v-btn v-if="role === 1" color="primary" dark class="mb-2" @click="onClickNewButton">New Item</v-btn>
            </v-card-title>
            <admin-list-table
              :loading="loading"
              :page="page"
              :items-per-page="itemsPerPage"
              :sort-by="sortBy"
              :sort-desc="sortDesc"
              :search="search"
              :role="role"
              :users="users"
              :total="total"
              @show="onClickShowButton"
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
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheSelect from '~/components/atoms/TheSelect.vue'
import TheTextField from '~/components/atoms/TheTextField.vue'
import AdminEditForm from '~/components/organisms/AdminEditForm.vue'
import AdminListTable from '~/components/organisms/AdminListTable.vue'
import AdminNewForm from '~/components/organisms/AdminNewForm.vue'
import { IAdminNewForm } from '~/types/forms'
import { IAdminUser } from '~/types/store'

export default defineComponent({
  components: {
    AdminEditForm,
    AdminListTable,
    AdminNewForm,
    TheFormGroup,
    TheSelect,
    TheTextField,
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
    role: {
      type: Number,
      required: false,
      default: 0,
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
    newForm: {
      type: Object as PropType<IAdminNewForm>,
      required: false,
      default: () => ({}),
    },
    newDialog: {
      type: Boolean,
      required: false,
      default: false,
    },
  },

  setup(_, { emit }: SetupContext) {
    const onClickNewButton = (): void => {
      emit('new:open')
    }

    const onClickNewClose = (): void => {
      emit('new:close')
    }

    const onClickCreateButton = (): void => {
      emit('create')
    }

    const onClickShowButton = (userId: string): void => {
      emit('show', userId)
    }

    return {
      onClickNewButton,
      onClickNewClose,
      onClickCreateButton,
      onClickShowButton,
    }
  },
})
</script>
