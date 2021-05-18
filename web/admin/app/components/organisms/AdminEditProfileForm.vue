<template>
  <v-form class="px-4">
    <v-list-item>
      <v-list-item-content class="col col-3">
        <v-list-item-subtitle>サムネイル</v-list-item-subtitle>
      </v-list-item-content>
      <v-list-item-content>
        <v-img label="サムネイル" :src="form.params.thumbnailUrl" max-width="240" contain />
      </v-list-item-content>
    </v-list-item>
    <v-file-input @change="onImagePicked" />
    <v-row>
      <v-col cols="12" md="6">
        <v-text-field v-model="form.params.lastName" label="姓" />
      </v-col>
      <v-col cols="12" md="6">
        <v-text-field v-model="form.params.firstName" label="名" />
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" md="6">
        <v-text-field v-model="form.params.lastNameKana" label="姓(かな)" />
      </v-col>
      <v-col cols="12" md="6">
        <v-text-field v-model="form.params.firstNameKana" label="名(かな)" />
      </v-col>
    </v-row>
    <v-select v-model="form.params.role" :items="roleItems" />
    <v-btn color="primary" class="mt-4 mr-4" @click="onSubmit">変更する</v-btn>
    <v-btn color="warning" class="mt-4 mr-4" @click="onDelete">管理者権限を削除する</v-btn>
    <v-btn class="mt-4" @click="onCancel">キャンセル</v-btn>
  </v-form>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  props: {
    form: {
      type: Object, // TODO: define type
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
