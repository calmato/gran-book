<template>
  <book-list
    :books="books"
    :search-form="searchForm"
    @submit="handleSubmitSearch"
    @clear="handleClearForm"
    @click="handleClickShowBook"
  />
</template>

<script lang="ts">
import { computed, defineComponent, reactive, SetupContext } from '@nuxtjs/composition-api'
import BookList from '~/components/templates/BookList.vue'
import { BookStore } from '~/store'
import { BookSearchOptions, IBookSearchForm, IBookSearchParams } from '~/types/forms'
import { IBook } from '~/types/store'

export default defineComponent({
  components: {
    BookList,
  },

  setup(_, { root }: SetupContext) {
    const store = root.$store

    const initializeForm: IBookSearchParams = {
      title: '',
      author: '',
      publisher: '',
      isbn: '',
      size: 0,
    }

    const searchForm = reactive<IBookSearchForm>({
      params: {
        ...initializeForm,
      },
      options: {
        ...BookSearchOptions,
      },
    })

    const books = computed(() => store.getters['book/getBooks'])

    const handleSubmitSearch = async () => {
      await BookStore.searchBookFromRakutenBooksAPI(searchForm).catch((err: Error) => {
        console.log('debug', 'err', err)
      })
    }

    const handleClickShowBook = (book: IBook) => {
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
