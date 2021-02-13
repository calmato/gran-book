<template>
  <settings-email-edit :form="form" @click="handleSubmit" />
</template>

<script lang="ts">
import { defineComponent, reactive, SetupContext } from '@nuxtjs/composition-api'
import { ISettingsEmailEditForm } from '~/types/forms'
import SettingsEmailEdit from '~/components/templates/SettingsEmailEdit.vue'

export default defineComponent({
  components: {
    SettingsEmailEdit,
  },

  setup(_, { root }: SetupContext) {
    const store = root.$store

    const form = reactive<ISettingsEmailEditForm>({
      email: store.getters['auth/getEmail'],
    })

    const handleSubmit = () => {
      console.log('debug', form)
    }

    return {
      form,
      handleSubmit,
    }
  },
})
</script>
