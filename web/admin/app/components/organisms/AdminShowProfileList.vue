<template>
  <v-list class="px-4">
    <v-list-item>
      <v-list-item-content class="col col-3">
        <v-list-item-subtitle>サムネイル</v-list-item-subtitle>
      </v-list-item-content>
      <v-list-item-content>
        <v-img :src="user.thumbnailUrl" max-height="120" max-width="120" contain />
      </v-list-item-content>
    </v-list-item>
    <v-divider />
    <v-list-item>
      <v-list-item-content class="col col-3">
        <v-list-item-subtitle>氏名</v-list-item-subtitle>
      </v-list-item-content>
      <v-list-item-content>{{ getName() }}</v-list-item-content>
    </v-list-item>
    <v-divider />
    <v-list-item>
      <v-list-item-content class="col col-3">
        <v-list-item-subtitle>権限</v-list-item-subtitle>
      </v-list-item-content>
      <v-list-item-content>{{ getRole() }}</v-list-item-content>
    </v-list-item>
    <v-divider />
    <v-list-item>
      <v-list-item-content class="col col-3">
        <v-list-item-subtitle>自己紹介</v-list-item-subtitle>
      </v-list-item-content>
      <v-list-item-content>{{ user.selfIntroduction }}</v-list-item-content>
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

  setup(props) {
    const joinName = (lastName: string, firstName: string): string => {
      const space: string = lastName && firstName ? ' ' : ''
      return lastName + space + firstName
    }

    const getName = (): string => {
      if (!props.user) {
        return ''
      }

      const name: string = joinName(props.user.lastName, props.user.firstName)
      const nameKana: string = joinName(props.user.lastNameKana, props.user.firstNameKana)

      if (nameKana === '') {
        return name
      } else {
        return `${name} (${nameKana})`
      }
    }

    const getRole = (): string => {
      const role: number = props.user ? props.user.role : -1

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

    return {
      getName,
      getRole,
    }
  },
})
</script>
