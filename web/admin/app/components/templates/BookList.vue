<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <v-card class="pa-4">
          <v-subheader>書籍検索</v-subheader>
          <v-card-text>
            <book-search-form :form="searchForm" @submit="onSubmitSearchForm" @clear="onClearSearchForm" />
          </v-card-text>
          <v-subheader>書籍一括登録</v-subheader>
          <book-multiple-register @submit="onRegisterMultilpleBooks" />
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col v-for="item in books" :key="item.isbn" cols="6" sm="4" md="3">
        <book-list-item :book="item" @click="onClickBookCard(item)" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import BookListItem from '~/components/organisms/BookListItem.vue'
import BookSearchForm from '~/components/organisms/BookSearchForm.vue'
import BookMultipleRegister from '~/components/organisms/BookMultipleRegister.vue'
import { IBookSearchForm } from '~/types/forms'
import { IBook } from '~/types/store'

export default defineComponent({
  components: {
    BookListItem,
    BookSearchForm,
    BookMultipleRegister,
  },

  props: {
    books: {
      type: Array as PropType<IBook[]>,
      required: false,
      default: () => [],
    },
    searchForm: {
      type: Object as PropType<IBookSearchForm>,
      required: false,
      default: () => ({}),
    },
  },

  setup(_, { emit }: SetupContext) {
    const onSubmitSearchForm = () => {
      emit('submit')
    }

    const onClickBookCard = (book: IBook) => {
      emit('click', book)
    }

    const onClearSearchForm = () => {
      emit('clear')
    }

    const onRegisterMultilpleBooks = () => {
      emit('register')
    }

    return {
      onClickBookCard,
      onClearSearchForm,
      onSubmitSearchForm,
      onRegisterMultilpleBooks,
    }
  },
})
</script>
