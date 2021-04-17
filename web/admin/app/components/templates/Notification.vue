<template>
  <v-container fill-height>
    <v-dialog v-model="newDialog" width="600px" scrollable @click:outside="$emit('update:new-dialog', false)">
      <v-card>
        <v-toolbar color="primary" dark>お知らせ 追加</v-toolbar>
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
            <v-subheader>お知らせ一覧</v-subheader>
            <v-card-title>
              <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details />
              <v-spacer />
              <v-btn color="primary" dark class="mb-2" @click="onClickNewButton">New Item</v-btn>
            </v-card-title>
            <notification-list-table
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

<script>
import { defineComponent } from '@vue/composition-api'
import NotificationListTable from '~/components/organisms/NotificationListTable.vue'

export default defineComponent({
  components: { NotificationListTable },
  setup() {},
})
</script>
