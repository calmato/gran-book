<template>
  <the-form-group>
    <v-row>
      <v-col cols="12" md="6">
        <the-text-field
          v-model="form.params.lastName"
          :label="form.options.lastName.label"
          :rules="form.options.lastName.rules"
          :autofocus="true"
        />
      </v-col>
      <v-col cols="12" md="6">
        <the-text-field
          v-model="form.params.firstName"
          :label="form.options.firstName.label"
          :rules="form.options.firstName.rules"
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" md="6">
        <the-text-field
          v-model="form.params.lastNameKana"
          :label="form.options.lastNameKana.label"
          :rules="form.options.lastNameKana.rules"
        />
      </v-col>
      <v-col cols="12" md="6">
        <the-text-field
          v-model="form.params.firstNameKana"
          :label="form.options.firstNameKana.label"
          :rules="form.options.firstNameKana.rules"
        />
      </v-col>
    </v-row>
    <the-text-field v-model="form.params.email" :label="form.options.email.label" :rules="form.options.email.rules" />
    <the-text-field
      v-model="form.params.phoneNumber"
      :label="form.options.phoneNumber.label"
      :rules="form.options.phoneNumber.rules"
    />
    <the-select
      v-model="form.params.role"
      :label="form.options.role.label"
      :rules="form.options.role.rules"
      :items="roleItems"
    />
    <v-img
      v-if="fileData || form.params.thumbnailUrl"
      :src="fileData || form.params.thumbnailUrl"
      width="auto"
      max-height="240"
      contain
    />
    <the-file-input
      v-model="form.params.thumbnail"
      :label="form.options.thumbnail.label"
      :rules="form.options.thumbnail.rules"
      :limit="form.options.thumbnail.rules.size"
      accept="image/*"
      @input="onImagePicked"
    />
  </the-form-group>
</template>

<script lang="ts">
import { defineComponent, PropType, ref } from '@nuxtjs/composition-api'
import TheFileInput from '~/components/atoms/TheFileInput.vue'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheSelect from '~/components/atoms/TheSelect.vue'
import TheTextField from '~/components/atoms/TheTextField.vue'
import { IAdminEditForm } from '~/types/forms'

export default defineComponent({
  components: {
    TheFileInput,
    TheFormGroup,
    TheSelect,
    TheTextField,
  },

  props: {
    form: {
      type: Object as PropType<IAdminEditForm>,
      required: true,
    },
  },

  setup(props) {
    const roleItems = [
      { text: '管理者', value: 1 },
      { text: '開発者', value: 2 },
      { text: '運用者', value: 3 },
    ]

    const fileData = ref<string | ArrayBuffer | null>()

    const onImagePicked = (file: File) => {
      if (!file) {
        props.form.params.thumbnailUrl = ''
        return
      }

      if (file.name.lastIndexOf('.') === -1) {
        return
      }

      const fr = new FileReader()

      fr.readAsDataURL(file)
      fr.addEventListener('load', () => {
        fileData.value = fr.result
      })
    }

    return {
      fileData,
      roleItems,
      onImagePicked,
    }
  },
})
</script>
