<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <v-text-field label="タイトル*" required></v-text-field>
      </v-col>
      <v-col cols="12">
        <v-text-field label="説明*" required></v-text-field>
      </v-col>
      <v-col cols="12" sm="6">
        <v-select :items="categoryList" label="カテゴリー*" required></v-select>
      </v-col>
      <v-col cols="12" sm="6">
        <v-select :items="importanceList" label="重要度*" required></v-select>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'

export default defineComponent({
  setup(_) {
    const categoryList = ['Information', 'Event', 'Maintenance']
    const importanceList = ['High', 'Middle', 'Low']

    return {
      categoryList,
      importanceList,
    }
  },
})
</script>
