<template>
  <the-form-group>
    <the-text-field
      v-model="form.params.password"
      :label="form.options.password.label"
      :rules="form.options.password.rules"
      :autofocus="true"
      type="password"
    />
    <the-text-field
      v-model="form.params.passwordConfirmation"
      :label="form.options.passwordConfirmation.label"
      :rules="form.options.passwordConfirmation.rules"
      type="password"
    />
    <v-btn color="primary" class="mt-4 mr-4" @click="onClick">変更する</v-btn>
    <v-btn class="mt-4" @click="onClickCancel">キャンセル</v-btn>
  </the-form-group>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { IAuthEditPasswordForm } from '~/types/forms'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheTextField from '~/components/atoms/TheTextField.vue'

export default defineComponent({
  components: {
    TheFormGroup,
    TheTextField,
  },

  props: {
    form: {
      type: Object as PropType<IAuthEditPasswordForm>,
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
