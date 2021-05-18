<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <v-card>
          <v-dialog v-model="deleteDialog" max-width="320" @click:outside="closeDeleteDialog">
            <v-card>
              <v-card-title class="headline">管理者権限の削除</v-card-title>
              <v-card-text>本当に削除しますか?</v-card-text>
              <v-card-actions>
                <v-spacer />
                <v-btn class="mr-4" @click="closeDeleteDialog">キャンセル</v-btn>
                <v-btn color="warning" @click="onSubmitDelete">削除する</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
          <v-card-title>
            プロフィール
            <v-spacer />
            <v-icon v-if="role === 1" @click="$emit('update:edit-profile', !editProfile)">
              {{ editProfile ? 'mdi-close' : 'mdi-pencil' }}
            </v-icon>
          </v-card-title>
          <v-divider />
          <v-card-text>
            <admin-edit-profile-form
              v-if="editProfile"
              :form="editProfileForm"
              @submit="onSubmitProfile"
              @delete="openDeleteDialog"
              @cancel="$emit('update:edit-profile', false)"
            />
            <admin-show-profile-list v-else :user="user" />
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
            <admin-edit-contact-form
              v-if="editContact"
              :form="editContactForm"
              @submit="onSubmitContact"
              @cancel="$emit('update:edit-contact', false)"
            />
            <admin-show-contact-list v-else :user="user" />
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
            <admin-edit-security-form
              v-if="editSecurity"
              :form="editSecurityForm"
              @submit="onSubmitSecurity"
              @cancel="$emit('update:edit-security', false)"
            />
            <admin-show-security-list v-else />
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, ref, SetupContext } from '@nuxtjs/composition-api'
import AdminEditContactForm from '~/components/organisms/AdminEditContactForm.vue'
import AdminEditProfileForm from '~/components/organisms/AdminEditProfileForm.vue'
import AdminEditSecurityForm from '~/components/organisms/AdminEditSecurityForm.vue'
import AdminShowContactList from '~/components/organisms/AdminShowContactList.vue'
import AdminShowProfileList from '~/components/organisms/AdminShowProfileList.vue'
import AdminShowSecurityList from '~/components/organisms/AdminShowSecurityList.vue'
import { IAdminUser } from '~/types/store'

export default defineComponent({
  components: {
    AdminEditContactForm,
    AdminEditProfileForm,
    AdminEditSecurityForm,
    AdminShowContactList,
    AdminShowProfileList,
    AdminShowSecurityList,
  },

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

    const openDeleteDialog = (): void => {
      deleteDialog.value = true
    }

    const closeDeleteDialog = (): void => {
      deleteDialog.value = false
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

    const onSubmitDelete = (): void => {
      emit('delete')
      closeDeleteDialog()
    }

    return {
      roleItems,
      deleteDialog,
      openDeleteDialog,
      closeDeleteDialog,
      onSubmitProfile,
      onSubmitContact,
      onSubmitSecurity,
      onSubmitDelete,
    }
  },
})
</script>
