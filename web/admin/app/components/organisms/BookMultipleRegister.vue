<template>
  <v-col cols="12" md="4" class="my-2">
    <v-file-input accept=".csv" label="csvを選択" filled prepend-icon="mdi-file-excel"> </v-file-input>
    <v-btn outlined v-on="on" @click="handleSubmit"> 書籍を一括登録 </v-btn>
  </v-col>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  setup(_, { emit }: SetupContext) {
    const handleSubmit = () => {
      emit('register')
    }

    return {
      handleSubmit,
    }
  },
})
</script>
