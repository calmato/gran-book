<template>
  <v-btn @click="handleClick">ログインページへ</v-btn>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  methods: {
    handleClick() {
      this.$router.push('/signin')
    },
  },
})
</script>
