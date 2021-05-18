<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <v-card>
          <v-card-title>
            プロフィール
            <v-spacer />
            <v-btn @click="editProfile = !editProfile">Toggle</v-btn>
          </v-card-title>
          <v-divider />
          <v-card-text>
            <v-form v-if="editProfile" class="px-4">
              <v-list-item>
                <v-list-item-content class="col col-3">
                  <v-list-item-subtitle>サムネイル</v-list-item-subtitle>
                </v-list-item-content>
                <v-list-item-content>
                  <v-img label="サムネイル" :src="editForm.params.thumbnailUrl" max-width="240" contain />
                </v-list-item-content>
              </v-list-item>
              <v-file-input />
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field v-model="editForm.params.lastName" label="姓" />
                </v-col>
                <v-col cols="12" md="6">
                  <v-text-field v-model="editForm.params.firstName" label="名" />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field v-model="editForm.params.lastNameKana" label="姓(かな)" />
                </v-col>
                <v-col cols="12" md="6">
                  <v-text-field v-model="editForm.params.firstNameKana" label="名(かな)" />
                </v-col>
              </v-row>
              <v-select v-model="editForm.params.role" :items="roleItems" />
            </v-form>
            <v-list v-else class="px-4">
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
                <v-list-item-content>
                  {{ `${user.lastName} ${user.firstName} (${user.lastNameKana} ${user.firstNameKana})` }}
                </v-list-item-content>
              </v-list-item>
              <v-divider />
              <v-list-item>
                <v-list-item-content class="col col-3">
                  <v-list-item-subtitle>権限</v-list-item-subtitle>
                </v-list-item-content>
                <v-list-item-content>{{ getRole(user.role) }}</v-list-item-content>
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
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="6">
        <v-card>
          <v-card-title>
            連絡先
            <v-spacer />
            <v-btn @click="editContact = !editContact">Toggle</v-btn>
          </v-card-title>
          <v-divider />
          <v-card-text>
            <v-form v-if="editContact" class="pa-4">
              <v-text-field v-model="editForm.params.email" label="メールアドレス" />
              <v-text-field v-model="editForm.params.phoneNumber" label="電話番号" />
            </v-form>
            <v-list v-else class="px-4">
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
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="6">
        <v-card>
          <v-card-title>
            セキュリティ
            <v-spacer />
            <v-btn @click="editSecurity = !editSecurity">Toggle</v-btn>
          </v-card-title>
          <v-divider />
          <v-card-text>
            <v-form v-if="editSecurity" class="pa-4">
              <v-text-field v-model="editForm.params.password" label="パスワード" />
              <v-text-field v-model="editForm.params.passwordConfirmation" label="パスワード(確認用)" />
            </v-form>
            <v-list v-else class="px-4">
              <v-list-item>
                <v-list-item-content>
                  <v-list-item-subtitle>パスワード</v-list-item-subtitle>
                </v-list-item-content>
                <v-list-item-content>************</v-list-item-content>
              </v-list-item>
              <v-divider />
            </v-list>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from '@vue/composition-api'

export default defineComponent({
  setup() {
    const roleItems = [
      { text: '管理者', value: 1 },
      { text: '開発者', value: 2 },
      { text: '運用者', value: 3 },
    ]
    const initializeEditForm = {
      email: '',
      phoneNumber: '',
      role: 2,
      lastName: '',
      firstName: '',
      lastNameKana: '',
      firstNameKana: '',
      thumbnail: null,
      thumbnailUrl: '',
      password: '',
      passwordConfirmation: '',
    }

    const editProfile = ref<boolean>(false)
    const editContact = ref<boolean>(false)
    const editSecurity = ref<boolean>(false)

    const user = reactive({
      email: 'test-user@calmato.jp',
      phoneNumber: '000-0000-0000',
      lastName: 'Calmato',
      firstName: '管理者',
      lastNameKana: 'かるまーと',
      firstNameKana: 'かんりしゃ',
      role: 2,
      thumbnail: null,
      thumbnailUrl: '/thumbnail.png',
      selfIntroduction: 'よろしくお願いします。',
    })
    const editForm = reactive({
      params: {
        ...initializeEditForm,
      },
      options: {},
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

    return {
      roleItems,
      user,
      editProfile,
      editContact,
      editSecurity,
      editForm,
      getRole,
    }
  },
})
</script>
