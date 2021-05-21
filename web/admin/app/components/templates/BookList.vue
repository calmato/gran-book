<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <book-list-form-card :form="searchForm" @submit="onSubmitSearchForm" @clear="onClearSearchForm" />
      </v-col>
    </v-row>
    <v-row>
      <v-col v-for="item in books" :key="item.isbn" cols="6" sm="4" md="3">
        <book-list-item :book="item" @click="onClickBookCard" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@nuxtjs/composition-api'
import BookListFormCard from '~/components/organisms/BookListFormCard.vue'
import BookListItem from '~/components/organisms/BookListItem.vue'

export default defineComponent({
  components: {
    BookListFormCard,
    BookListItem,
  },

  props: {
    books: {
      type: Array, // TODO: define type
      required: false,
      default: () => [],
    },
    searchForm: {
      type: Object, // TODO: define type
      required: false,
      default: () => ({}),
    },
  },

  setup(_, { emit }: SetupContext) {
    const onSubmitSearchForm = () => {
      emit('submit')
    }

    // TODO: define type
    const onClickBookCard = (book: any) => {
      emit('click', book)
    }

    const onClearSearchForm = () => {
      emit('clear')
    }

    return {
      onClickBookCard,
      onClearSearchForm,
      onSubmitSearchForm,
    }
  },
})
</script>
