<template>
  <v-container fill-height>
    <v-layout wrap>
      <v-row>
        <v-col cols="12">
          <v-card class="pa-4 my-4">
            <v-tabs vertical flat>
              <v-tab>
                <v-icon left>mdi-account</v-icon>
                プロフィール
              </v-tab>
              <v-tab>
                <v-icon left>mdi-lock</v-icon>
                アカウント
              </v-tab>

              <v-tab-item>
                <settings-profile-list :lists="profileLists" @click="onClickEditButton(profileEditPath)" />
              </v-tab-item>

              <v-tab-item>
                <settings-account-list :lists="accountLists" @click="onClickEditButton" />
              </v-tab-item>
            </v-tabs>
          </v-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'
import { IAccountList, IProfileList } from '~/types/props'
import SettingsAccountList from '~/components/organisms/SettingsAccountList.vue'
import SettingsProfileList from '~/components/organisms/SettingsProfileList.vue'

export default defineComponent({
  components: {
    SettingsAccountList,
    SettingsProfileList,
  },

  props: {
    username: {
      type: String,
      required: false,
      default: '',
    },
    name: {
      type: String,
      required: false,
      default: '',
    },
    nameKana: {
      type: String,
      required: false,
      default: '',
    },
    thumbnailUrl: {
      type: String,
      required: false,
      default: '',
    },
    selfIntroduction: {
      type: String,
      required: false,
      default: '',
    },
    phoneNumber: {
      type: String,
      required: false,
      default: '',
    },
    email: {
      type: String,
      required: false,
      default: '',
    },
  },

  setup(props, { emit }: SetupContext) {
    const profileEditPath: string = '/settings/profile'

    const profileLists: IProfileList[] = [
      {
        title: '表示名',
        content: props.username,
        contentType: 'text',
      },
      {
        title: '氏名',
        content: props.nameKana ? `${props.name} (${props.nameKana})` : props.name,
        contentType: 'text',
      },
      {
        title: 'アイコン',
        content: props.thumbnailUrl,
        contentType: 'image',
      },
      {
        title: '自己紹介',
        content: props.selfIntroduction,
        contentType: 'text',
      },
      {
        title: '電話番号',
        content: props.phoneNumber,
        contentType: 'text',
      },
    ]

    const accountLists: IAccountList[] = [
      {
        title: 'メールアドレス',
        content: props.email,
        to: '/settings/email',
      },
      {
        title: 'パスワード',
        content: '************',
        to: '/settings/password',
      },
    ]

    const onClickEditButton = (path: string) => {
      emit('click', path)
    }

    return {
      profileEditPath,
      profileLists,
      accountLists,
      onClickEditButton,
    }
  },
})
</script>
