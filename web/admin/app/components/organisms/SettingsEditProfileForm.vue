<template>
  <the-form-group>
    <the-text-field
      v-model="form.params.username"
      :label="form.options.username.label"
      :rules="form.options.username.rules"
      :autofocus="true"
    />
    <the-file-input
      v-model="form.params.thumbnail"
      :label="form.options.thumbnail.label"
      :rules="form.options.thumbnail.rules"
      :limit="form.options.thumbnail.rules.size"
      accept="image/*"
    />
    <the-text-area
      v-model="form.params.selfIntroduction"
      :label="form.options.selfIntroduction.label"
      :rules="form.options.selfIntroduction.rules"
    />
    <the-text-field
      v-model="form.params.lastName"
      :label="form.options.lastName.label"
      :rules="form.options.lastName.rules"
    />
    <the-text-field
      v-model="form.params.firstName"
      :label="form.options.firstName.label"
      :rules="form.options.firstName.rules"
    />
    <the-text-field
      v-model="form.params.lastNameKana"
      :label="form.options.lastNameKana.label"
      :rules="form.options.lastNameKana.rules"
    />
    <the-text-field
      v-model="form.params.firstNameKana"
      :label="form.options.firstNameKana.label"
      :rules="form.options.firstNameKana.rules"
    />
    <the-text-field
      v-model="form.params.phoneNumber"
      :label="form.options.phoneNumber.label"
      :rules="form.options.phoneNumber.rules"
    />
    <v-btn :loading="loading" :disabled="loading" color="primary" class="mt-4 mr-4" @click="onClick">変更する</v-btn>
    <v-btn class="mt-4" @click="onClickCancel">キャンセル</v-btn>
  </the-form-group>
</template>

<script lang="ts">
import { defineComponent, SetupContext, PropType } from '@nuxtjs/composition-api'
import { IAuthEditProfileForm } from '~/types/forms'
import TheFileInput from '~/components/atoms/TheFileInput.vue'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheTextArea from '~/components/atoms/TheTextArea.vue'
import TheTextField from '~/components/atoms/TheTextField.vue'

export default defineComponent({
  components: {
    TheFileInput,
    TheFormGroup,
    TheTextArea,
    TheTextField,
  },

  props: {
    form: {
      type: Object as PropType<IAuthEditProfileForm>,
      required: true,
    },
    loading: {
      type: Boolean,
      required: false,
      default: false,
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
