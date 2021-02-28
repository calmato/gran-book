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
      v-model="form.params.password"
      :label="form.options.password.label"
      :rules="form.options.password.rules"
      type="password"
    />
    <the-text-field
      v-model="form.params.passwordConfirmation"
      :label="form.options.passwordConfirmation.label"
      :rules="form.options.passwordConfirmation.rules"
      type="password"
    />
    <the-select
      v-model="form.params.role"
      :label="form.options.role.label"
      :rules="form.options.role.rules"
      :items="roleItems"
    />
  </the-form-group>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheSelect from '~/components/atoms/TheSelect.vue'
import TheTextField from '~/components/atoms/TheTextField.vue'

export default defineComponent({
  components: {
    TheFormGroup,
    TheSelect,
    TheTextField,
  },

  props: {
    form: {
      type: Object,
      required: true,
    },
  },

  setup(_) {
    const roleItems = [
      { text: '未選択', value: 0 },
      { text: '男性', value: 1 },
      { text: '女性', value: 2 },
    ]

    return {
      roleItems,
    }
  },
})
</script>
