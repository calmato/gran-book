<template>
  <v-list class="px-4">
    <v-list-item>
      <v-list-item-content>
        <v-list-item-subtitle>メールアドレス</v-list-item-subtitle>
      </v-list-item-content>
      <v-list-item-content>{{ user.email }}</v-list-item-content>
    </v-list-item>
    <v-divider />
    <v-list-item>
      <v-list-item-content>
        <v-list-item-subtitle>電話番号</v-list-item-subtitle>
      </v-list-item-content>
      <v-list-item-content>{{ user.phoneNumber }}</v-list-item-content>
    </v-list-item>
    <v-divider />
  </v-list>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@nuxtjs/composition-api'
import { IAdminUser } from '~/types/store'

export default defineComponent({
  props: {
    user: {
      type: Object as PropType<IAdminUser>,
      required: false,
      default: () => ({}),
    },
  },
})
</script>
