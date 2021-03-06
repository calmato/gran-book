<template>
  <v-form class="px-4">
    <v-list-item>
      <v-list-item-content class="col col-3">
        <v-list-item-subtitle>サムネイル</v-list-item-subtitle>
      </v-list-item-content>
      <v-list-item-content>
        <v-img v-if="form.params.thumbnailUrl" :src="form.params.thumbnailUrl" max-width="240" contain />
      </v-list-item-content>
    </v-list-item>
    <the-file-input
      :file="form.params.thumbnail"
      :label="form.options.thumbnail.label"
      :rules="form.options.thumbnail.rules"
      accept="image/*"
      @change="onImagePicked"
    />
    <v-row>
      <v-col cols="12" md="6">
        <the-text-field
          v-model="form.params.lastName"
          :label="form.options.lastName.label"
          :rules="form.options.lastName.rules"
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
    <the-select
      v-model="form.params.role"
      :label="form.options.role.label"
      :rules="form.options.role.rules"
      :items="roleItems"
    />
    <v-btn color="primary" class="mt-4 mr-4" @click="onSubmit">変更する</v-btn>
    <v-btn color="warning" class="mt-4 mr-4" @click="onDelete">管理者権限を削除する</v-btn>
    <v-btn class="mt-4" @click="onCancel">キャンセル</v-btn>
  </v-form>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheFileInput from '~/components/atoms/TheFileInput.vue'
import TheSelect from '~/components/atoms/TheSelect.vue'
import TheTextField from '~/components/atoms/TheTextField.vue'
import { IAdminEditProfileForm } from '~/types/forms'

export default defineComponent({
  components: {
    TheFileInput,
    TheSelect,
    TheTextField,
  },

  props: {
    form: {
      type: Object as PropType<IAdminEditProfileForm>,
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
        props.form.params.thumbnail = file
        props.form.params.thumbnailUrl = fr.result
      })
    }

    const onSubmit = (): void => {
      emit('submit')
    }

    const onCancel = (): void => {
      emit('cancel')
    }

    const onDelete = (): void => {
      emit('delete')
    }

    return {
      roleItems,
      onImagePicked,
      onSubmit,
      onCancel,
      onDelete,
    }
  },
})
</script>
