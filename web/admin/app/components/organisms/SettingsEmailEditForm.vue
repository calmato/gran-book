<template>
  <v-form>
    <v-text-field v-model="form.email" label="email" class="mb-4" autofocus />
    <v-btn color="primary" class="mr-4" @click="onClick">確認メールを送信</v-btn>
    <v-btn @click="onClickCancel">キャンセル</v-btn>
  </v-form>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { ISettingsEmailEditForm } from '~/types/forms'

export default defineComponent({
  props: {
    form: {
      type: Object as PropType<ISettingsEmailEditForm>,
      required: true,
    },
  },

  setup(_, { emit }: SetupContext) {
    const onClick = () => {
      emit('click')
    }

    const onClickCancel = () => {
      emit('cancel')
    }

    return {
      onClick,
      onClickCancel,
    }
  },
})
</script>
