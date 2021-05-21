<template>
  <v-form dense @submit.prevent="onSubmit">
    <v-text-field v-model="form.params.title" prepend-icon="mdi-magnify" label="書籍名を入力" hide-details>
      <template v-slot:append-outer>
        <v-menu v-model="menu" :close-on-content-click="false" bottom left>
          <template v-slot:activator="{ on, attrs }">
            <v-btn outlined v-bind="attrs" v-on="on">検索条件を追加</v-btn>
          </template>
          <v-card>
            <v-card-text>
              <v-form class="px-4">
                <v-text-field v-model="form.params.author" label="著者名" />
                <v-text-field v-model="form.params.publisher" label="出版社名" />
                <v-text-field v-model="form.params.isbn" label="ISBN" />
                <v-select v-model="form.params.size" :items="sizes" label="書籍サイズ" />
              </v-form>
            </v-card-text>
            <v-card-actions>
              <v-spacer />
              <v-btn text color="error" @click="onClear">Clear</v-btn>
              <v-btn text color="primary" @click="menu = false">Close</v-btn>
            </v-card-actions>
          </v-card>
        </v-menu>
      </template>
    </v-text-field>
  </v-form>
</template>

<script lang="ts">
import { defineComponent, ref, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  props: {
    form: {
      type: Object, // TODO: define type
      required: false,
      default: () => ({}),
    },
  },

  setup(_, { emit }: SetupContext) {
    const sizes = [
      { text: '全て', value: 0 },
      { text: '単行本', value: 1 },
      { text: '文庫', value: 2 },
      { text: '新書', value: 3 },
      { text: '全集・双書', value: 4 },
      { text: '事・辞典', value: 5 },
      { text: '図鑑', value: 6 },
      { text: '絵本', value: 7 },
      { text: 'カセット,CDなど', value: 8 },
      { text: 'コミック', value: 9 },
      { text: 'ムックその他', value: 10 },
    ]

    const menu = ref<boolean>(false)

    const onSubmit = () => {
      emit('submit')
    }

    const onClear = () => {
      emit('clear')
    }

    return {
      menu,
      sizes,
      onSubmit,
      onClear,
    }
  },
})
</script>
