<template>
  <v-container fill-height>
    <v-layout wrap>
      <v-row>
        <v-col cols="12">
          <v-card class="pa-4">
            <v-card-title>プロフィール変更</v-card-title>
            <v-card-text>
              <settings-edit-profile-form :form="form" @click="onClickSubmitButton" @cancel="onClickCancelButton" />
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { IAuthEditProfileForm } from '~/types/forms'
import SettingsEditProfileForm from '~/components/organisms/SettingsEditProfileForm.vue'

export default defineComponent({
  components: {
    SettingsEditProfileForm,
  },

  props: {
    form: {
      type: Object as PropType<IAuthEditProfileForm>,
      required: true,
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
