<template>
  <v-col cols="12" sm="8">
    <v-img src="/logo.png" max-height="96px" contain class="mb-8" />

    <v-card rounded="32">
      <v-card-title class="justify-center">
        <h3>管理者向けログイン</h3>
      </v-card-title>
      <v-divider />
      <v-card-text class="px-12">
        <v-form class="py-8">
          <v-text-field label="email" autofocus outlined />
          <v-text-field label="password" outlined />
          <v-btn :block="true" color="primary" class="mt-4">ログイン</v-btn>
        </v-form>
      </v-card-text>
    </v-card>
  </v-col>
</template>
