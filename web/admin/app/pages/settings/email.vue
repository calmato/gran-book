<template>
  <settings-email-edit :form="form" @click="handleSubmit" @cancel="handleCancel" />
</template>

<script lang="ts">
import { defineComponent, reactive, SetupContext } from '@nuxtjs/composition-api'
import { AuthStore } from '~/store'
import { ISettingsEmailEditForm } from '~/types/forms'
import SettingsEmailEdit from '~/components/templates/SettingsEmailEdit.vue'

export default defineComponent({
  components: {
    SettingsEmailEdit,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const form = reactive<ISettingsEmailEditForm>({
      email: store.getters['auth/getEmail'],
    })

    const handleSubmit = async () => {
      await AuthStore.updateEmail(form)
        .then(() => {
          // TODO: トースト表示処理追加
          router.push('/settings')
        })
        .catch(() => console.log('debug', 'failure'))
    }

    const handleCancel = () => {
      router.back()
    }

    return {
      form,
      handleSubmit,
      handleCancel,
    }
  },
})
</script>
