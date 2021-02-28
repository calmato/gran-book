<template>
  <admin-list
    :search.sync="search"
    :page.sync="page"
    :items-per-page.sync="itemsPerPage"
    :sort-by.sync="sortBy"
    :sort-desc.sync="sortDesc"
    :loading="loading"
    :users="users"
    :total="total"
    @new="handleClickNewItem"
    @edit="handleClickEditItem"
    @delete="handleClickDeleteItem"
  />
</template>

<script lang="ts">
import { defineComponent, SetupContext, ref, useAsync, computed, watch } from '@nuxtjs/composition-api'
import AdminList from '~/components/templates/AdminList.vue'
import { AdminStore } from '~/store'
import { IAdminListForm } from '~/types/forms'

export default defineComponent({
  components: {
    AdminList,
  },

  setup(_, { root }: SetupContext) {
    const store = root.$store

    const loading = ref<boolean>(false)
    const search = ref<string>()
    const page = ref<number>(1)
    const itemsPerPage = ref<number>(20)
    const sortBy = ref<string>()
    const sortDesc = ref<boolean>()

    const users = computed(() => store.getters['admin/getUsers'])
    const total = computed(() => store.getters['admin/getTotal'])

    watch(page, (): void => {
      indexAdmin()
    })

    watch(itemsPerPage, (): void => {
      indexAdmin()
    })

    watch(sortBy, (): void => {
      indexAdmin()
    })

    watch(sortDesc, (): void => {
      indexAdmin()
    })

    useAsync(async () => {
      await indexAdmin()
    })

    const handleClickNewItem = async (): Promise<void> => {
      await AdminStore.createUser()
    }

    const handleClickEditItem = (index: number): void => {
      console.log('debug', 'editItem', index)
    }

    const handleClickDeleteItem = (index: number): void => {
      console.log('debug', 'deleteItem', index)
    }

    async function indexAdmin(): Promise<void> {
      loading.value = true

      const form: IAdminListForm = {
        limit: itemsPerPage.value,
        offset: itemsPerPage.value * (page.value - 1),
        order: {
          by: sortBy.value || '',
          desc: sortDesc.value || false,
        },
      }

      await AdminStore.indexAdmin(form).finally(() => {
        loading.value = false
      })
    }

    return {
      loading,
      search,
      page,
      itemsPerPage,
      sortBy,
      sortDesc,
      users,
      total,
      handleClickNewItem,
      handleClickEditItem,
      handleClickDeleteItem,
    }
  },
})
</script>