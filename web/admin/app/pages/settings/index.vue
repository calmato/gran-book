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
                <v-list class="pa-4">
                  <div v-for="list in profileLists" :key="list.title">
                    <v-list-item>
                      <v-list-item-content>
                        <v-list-item-subtitle>{{ list.title }}</v-list-item-subtitle>
                      </v-list-item-content>
                      <v-list-item-content v-if="list.contentType === 'image'">
                        <v-img :src="list.content" max-height="120" max-width="120" />
                      </v-list-item-content>
                      <v-list-item-content v-else>
                        {{ list.content }}
                      </v-list-item-content>
                    </v-list-item>
                    <v-divider />
                  </div>
                  <div class="text-center">
                    <v-btn color="primary mt-4" block @click="onClick(profileEditPath)">編集する</v-btn>
                  </div>
                </v-list>
              </v-tab-item>

              <v-tab-item>
                <v-list>
                  <div v-for="list in accountLists" :key="list.title">
                    <v-list-item>
                      <v-list-item-content>
                        <v-list-item-subtitle>{{ list.title }}</v-list-item-subtitle>
                      </v-list-item-content>
                      <v-list-item-content>{{ list.content }}</v-list-item-content>
                      <v-list-item-action>
                        <v-btn color="primary" @click="onClick(list.to)">変更する</v-btn>
                      </v-list-item-action>
                    </v-list-item>
                    <v-divider />
                  </div>
                </v-list>
              </v-tab-item>
            </v-tabs>
          </v-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'

export default defineComponent({
  setup(_) {
    const profileEditPath: string = '/settings/profile'

    const profileLists: any[] = [
      {
        title: '表示名',
        content: 'Calmato 管理者',
        contentType: 'text',
      },
      {
        title: '氏名',
        content: 'Calmato 管理者',
        contentType: 'text',
      },
      {
        title: 'アイコン',
        content: '/thumbnail.png',
        contentType: 'image',
      },
      {
        title: '自己紹介',
        content: 'よろしくお願いします',
        contentType: 'text',
      },
      {
        title: '電話番号',
        content: '000-0000-0000',
        contentType: 'text',
      },
    ]

    const accountLists: any[] = [
      {
        title: 'メールアドレス',
        content: 'support@calmato.com',
        to: '/settings/email',
      },
      {
        title: 'パスワード',
        content: '************',
        to: '/settings/password',
      },
    ]

    const onClick = (path: string) => {
      console.log('debug', path)
    }

    return {
      profileEditPath,
      profileLists,
      accountLists,
      onClick,
    }
  },
})
</script>
