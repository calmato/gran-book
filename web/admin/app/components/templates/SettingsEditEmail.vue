<template>
  <v-container fill-height>
    <v-layout wrap>
      <v-row>
        <v-col cols="12">
          <v-card class="pa-4">
            <v-card-title>メールアドレス変更</v-card-title>
            <v-card-text>
              <settings-edit-email-form :form="form" @click="onClickSubmitButton" @cancel="onClickCancelButton" />
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { IAuthEditEmailForm } from '~/types/forms'
import SettingsEditEmailForm from '~/components/organisms/SettingsEditEmailForm.vue'

export default defineComponent({
  components: {
    SettingsEditEmailForm,
  },

  props: {
    form: {
      type: Object as PropType<IAuthEditEmailForm>,
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
