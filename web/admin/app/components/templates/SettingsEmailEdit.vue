<template>
  <v-container fill-height>
    <v-layout wrap>
      <v-row>
        <v-col cols="12">
          <v-card class="pa-4">
            <v-card-title>メールアドレス変更</v-card-title>
            <v-card-text>
              <settings-email-edit-form :form="form" @click="onClickSubmitButton" @cancel="onClickCancelButton" />
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { ISettingsEmailEditForm } from '~/types/forms'
import SettingsEmailEditForm from '~/components/organisms/SettingsEmailEditForm.vue'

export default defineComponent({
  components: {
    SettingsEmailEditForm,
  },

  props: {
    form: {
      type: Object as PropType<ISettingsEmailEditForm>,
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
