<template>
  <v-container fill-height>
    <v-layout wrap>
      <v-row>
        <v-col cols="12">
          <v-card class="pa-4">
            <v-card-title>パスワード変更</v-card-title>
            <v-card-text>
              <settings-edit-password-form :form="form" :loading="loading" @click="onClickSubmitButton" @cancel="onClickCancelButton" />
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { IAuthEditPasswordForm } from '~/types/forms'
import SettingsEditPasswordForm from '~/components/organisms/SettingsEditPasswordForm.vue'

export default defineComponent({
  components: {
    SettingsEditPasswordForm,
  },

  props: {
    form: {
      type: Object as PropType<IAuthEditPasswordForm>,
      required: true,
    },
    loading: {
      type: Boolean,
      required: false,
      default: false,
    },
  },

  setup(_, { emit }: SetupContext) {
    const onClickSubmitButton = () => {
      emit('click')
    }

    const onClickCancelButton = () => {
      emit('cancel')
    }

    return {
      onClickSubmitButton,
      onClickCancelButton,
    }
  },
})
</script>
