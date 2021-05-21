<template>
  <book-list :books="books" :search-form="searchForm" />
</template>

<script lang="ts">
import { defineComponent, reactive } from '@nuxtjs/composition-api'
import BookList from '~/components/templates/BookList.vue'

export default defineComponent({
  components: {
    BookList,
  },

  setup() {
    const initializeForm = {
      title: '',
      author: '',
      publisher: '',
      isbn: '',
      size: 0,
    }

    const searchForm = reactive({
      params: {
        ...initializeForm,
      },
      options: {
        rules: {},
      },
    })
    const books = reactive([
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

    const handleSubmitSearch = () => {
      console.log('debug', 'submit', searchForm)
    }

    const handleClickShowBook = (book: any) => {
      console.log('debug', 'show', book)
    }

    const handleClearForm = () => {
      searchForm.params = { ...initializeForm }
    }

    return {
      books,
      searchForm,
      handleSubmitSearch,
      handleClickShowBook,
      handleClearForm,
    }
  },
})
</script>
