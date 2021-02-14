<template>
  <the-form-group>
    <the-text-field
      v-model="form.params.email"
      :label="form.options.email.label"
      :rules="form.options.email.rules"
      :autofocus="true"
    />
    <v-btn color="primary" class="mt-4 mr-4" @click="onClick">確認メールを送信</v-btn>
    <v-btn class="mt-4" @click="onClickCancel">キャンセル</v-btn>
  </the-form-group>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { IAuthEditEmailForm } from '~/types/forms'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheTextField from '~/components/atoms/TheTextField.vue'

export default defineComponent({
  components: {
    TheFormGroup,
    TheTextField,
  },

  props: {
    form: {
      type: Object as PropType<IAuthEditEmailForm>,
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
