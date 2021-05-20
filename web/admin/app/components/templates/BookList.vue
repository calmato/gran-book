<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <v-card class="pa-4">
          <v-subheader>書籍管理トップ</v-subheader>
          <v-card-title>
            <form @submit.prevent="onSubmit">
              <v-text-field v-model="form.title" prepend-icon="mdi-magnify" label="書籍名を入力" hide-details />
            </form>
            <v-spacer />
            <v-menu v-model="menu" :close-on-content-click="false" bottom left>
              <template v-slot:activator="{ on, attrs }">
                <v-btn outlined v-bind="attrs" v-on="on">検索条件を追加</v-btn>
              </template>
              <v-card>
                <v-card-text>
                  <v-list>
                    <v-list-item>
                      <v-text-field v-model="form.author" label="著者名" />
                    </v-list-item>
                    <v-list-item>
                      <v-text-field v-model="form.publisher" label="出版社名" />
                    </v-list-item>
                    <v-list-item>
                      <v-text-field v-model="form.isbn" label="ISBN" />
                    </v-list-item>
                    <v-list-item>
                      <v-select v-model="form.size" :items="sizes" label="書籍サイズ" />
                    </v-list-item>
                  </v-list>
                </v-card-text>
                <v-card-actions>
                  <v-spacer />
                  <v-btn text @click="onClear">Clear</v-btn>
                  <v-btn color="primary" text @click="menu = false">Close</v-btn>
                </v-card-actions>
              </v-card>
            </v-menu>
          </v-card-title>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col v-for="item in items" :key="item.isbn" cols="6" sm="4" md="3">
        <v-card outlined shaped @click="onClick(item)">
          <v-img :src="item.thumbnailUrl" height="200px" contain />
          <v-card-title>{{ item.title }}</v-card-title>
          <v-card-subtitle>{{ item.author }}</v-card-subtitle>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from '@nuxtjs/composition-api'

export default defineComponent({
  setup() {
    const initializeForm = {
      title: '',
      author: '',
      publisher: '',
      isbn: '',
      size: 0,
    }

    const menu = ref<boolean>(false)
    const items = reactive([
      {
        title: '小説　ちはやふる　上の句',
        isbn: '9784062938426',
        thumbnailUrl: 'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120',
        author: '有沢 ゆう希/末次 由紀',
      },
      {
        title: '小説　ちはやふる　下の句',
        isbn: '9784062938471',
        thumbnailUrl: 'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8471/9784062938471.jpg?_ex=120x120',
        author: '有沢 ゆう希/末次 由紀',
      },
      {
        title: '小説　ちはやふる　結び',
        isbn: '9784062938556',
        thumbnailUrl: 'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8556/9784062938556.jpg?_ex=120x120',
        author: '有沢 ゆう希/末次 由紀',
      },
      {
        title: '呪術廻戦 1',
        isbn: '9784088815169',
        thumbnailUrl: 'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/5169/9784088815169.jpg?_ex=120x120',
        author: '芥見 下々',
      },
      {
        title: '呪術廻戦 2',
        isbn: '9784088816081',
        thumbnailUrl: 'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/6081/9784088816081.jpg?_ex=120x120',
        author: '芥見 下々',
      },
      {
        title: 'ネットワークはなぜつながるのか第2版',
        isbn: '9784822283117',
        thumbnailUrl: 'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8222/82228311.jpg?_ex=120x120',
        author: '戸根勤/日経network編集部',
      },
      {
        title: 'MySQL徹底入門 第4版 MySQL 8.0対応',
        isbn: '9784798161488',
        thumbnailUrl: 'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/1488/9784798161488.jpg?_ex=120x120',
        author: 'yoku0825/坂井 恵',
      },
    ])
    const form = ref({
      ...initializeForm,
    })

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

    const onSubmit = () => {
      console.log('debug', 'submit', form)
    }

    const onClick = (item: any) => {
      console.log('debug', 'show', item)
    }

    const onClear = () => {
      form.value = { ...initializeForm }
    }

    return {
      menu,
      form,
      items,
      sizes,
      onSubmit,
      onClick,
      onClear,
    }
  },
})
</script>
