<template>
  <v-container>
    <v-dialog v-model="deleteDialog" max-width="320">
      <v-card>
        <v-card-title class="headline">管理者権限の削除</v-card-title>
        <v-card-text>本当に削除しますか?</v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn class="mr-4" @click="deleteDialog = false">キャンセル</v-btn>
          <v-btn color="warning" @click="onClickDeleteButton">削除する</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-row>
      <v-col cols="12">
        <v-card>
          <v-card-title>
            プロフィール
            <v-spacer />
            <v-icon v-if="role === 1" @click="$emit('update:edit-profile', !editProfile)">
              {{ editProfile ? 'mdi-close' : 'mdi-pencil' }}
            </v-icon>
          </v-card-title>
          <v-divider />
          <v-card-text>
            <v-form v-if="editProfile" class="px-4">
              <v-list-item>
                <v-list-item-content class="col col-3">
                  <v-list-item-subtitle>サムネイル</v-list-item-subtitle>
                </v-list-item-content>
                <v-list-item-content>
                  <v-img label="サムネイル" :src="editProfileForm.params.thumbnailUrl" max-width="240" contain />
                </v-list-item-content>
              </v-list-item>
              <v-file-input />
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field v-model="editProfileForm.params.lastName" label="姓" />
                </v-col>
                <v-col cols="12" md="6">
                  <v-text-field v-model="editProfileForm.params.firstName" label="名" />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field v-model="editProfileForm.params.lastNameKana" label="姓(かな)" />
                </v-col>
                <v-col cols="12" md="6">
                  <v-text-field v-model="editProfileForm.params.firstNameKana" label="名(かな)" />
                </v-col>
              </v-row>
              <v-select v-model="editProfileForm.params.role" :items="roleItems" />
              <v-btn color="primary" class="mt-4 mr-4" @click="onSubmitProfile">変更する</v-btn>
              <v-btn color="warning" class="mt-4 mr-4" @click.stop="deleteDialog = true">管理者権限を削除する</v-btn>
              <v-btn class="mt-4" @click="editProfile = false">キャンセル</v-btn>
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
                  {{ getNameKana() ? `${getName()} (${getNameKana()})` : getName(), }}
                </v-list-item-content>
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
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="6">
        <v-card>
          <v-card-title>
            連絡先
            <v-spacer />
            <v-icon v-if="role === 1" @click="$emit('update:edit-contact', !editContact)">
              {{ editContact ? 'mdi-close' : 'mdi-pencil' }}
            </v-icon>
          </v-card-title>
          <v-divider />
          <v-card-text>
            <v-form v-if="editContact" class="pa-4">
              <v-text-field v-model="editContactForm.params.email" label="メールアドレス" />
              <v-text-field v-model="editContactForm.params.phoneNumber" label="電話番号" />
              <v-btn color="primary" class="mt-4 mr-4" @click="onSubmitContact">変更する</v-btn>
              <v-btn class="mt-4" @click="editContact = false">キャンセル</v-btn>
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
            <v-icon v-if="role === 1" @click="$emit('update:edit-security', !editSecurity)">
              {{ editSecurity ? 'mdi-close' : 'mdi-pencil' }}
            </v-icon>
          </v-card-title>
          <v-divider />
          <v-card-text>
            <v-form v-if="editSecurity" class="pa-4">
              <v-text-field v-model="editSecurityForm.params.password" type="password" label="パスワード" />
              <v-text-field
                v-model="editSecurityForm.params.passwordConfirmation"
                type="password"
                label="パスワード(確認用)"
              />
              <v-btn color="primary" class="mt-4 mr-4" @click="onSubmitSecurity">変更する</v-btn>
              <v-btn class="mt-4" @click="editSecurity = false">キャンセル</v-btn>
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
import { defineComponent, PropType, ref, SetupContext } from '@nuxtjs/composition-api'
import { IAdminUser } from '~/types/store'

export default defineComponent({
  props: {
    role: {
      type: Number,
      required: false,
      default: 0,
    },
    user: {
      type: Object as PropType<IAdminUser>,
      required: false,
      default: () => ({}),
    },
    editProfile: {
      type: Boolean,
      required: false,
      default: false,
    },
    editContact: {
      type: Boolean,
      required: false,
      default: false,
    },
    editSecurity: {
      type: Boolean,
      required: false,
      default: false,
    },
    editProfileForm: {
      type: Object, // TODO: define type
      required: false,
      default: () => ({}),
    },
    editContactForm: {
      type: Object, // TODO: define type
      required: false,
      default: () => ({}),
    },
    editSecurityForm: {
      type: Object, // TODO: define type
      required: false,
      default: () => ({}),
    },
  },

  setup(props, { emit }: SetupContext) {
    const roleItems = [
      { text: '管理者', value: 1 },
      { text: '開発者', value: 2 },
      { text: '運用者', value: 3 },
    ]

    const deleteDialog = ref<boolean>(false)

    const getName = (): string => {
      if (!props.user) {
        return ''
      }

      const space: string = props.user.lastName && props.user.firstName ? ' ' : ''
      return props.user.lastName + space + props.user.firstName
    }

    const getNameKana = (): string => {
      if (!props.user) {
        return ''
      }

      const space: string = props.user.lastNameKana && props.user.firstNameKana ? ' ' : ''
      return props.user.lastNameKana + space + props.user.firstNameKana
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

    const onSubmitProfile = (): void => {
      emit('submit:profile')
    }

    const onSubmitContact = (): void => {
      emit('submit:contact')
    }

    const onSubmitSecurity = (): void => {
      emit('submit:security')
    }

    const onClickDeleteButton = (): void => {
      emit('delete')
    }

    return {
      roleItems,
      deleteDialog,
      getName,
      getNameKana,
      getRole,
      onSubmitProfile,
      onSubmitContact,
      onSubmitSecurity,
      onClickDeleteButton,
    }
  },
})
</script>
