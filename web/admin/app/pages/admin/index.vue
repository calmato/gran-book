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
    :new-form="newForm"
    :new-dialog.sync="newDialog"
    @new:open="handleClickNewItem"
    @new:close="handleClickCloseNewItem"
    @create="handleClickCreateItem"
    @edit="handleClickEditItem"
    @delete="handleClickDeleteItem"
  />
</template>

<script lang="ts">
import { defineComponent, SetupContext, ref, reactive, useAsync, computed, watch } from '@nuxtjs/composition-api'
import { AdminStore, CommonStore } from '~/store'
import { AdminNewOptions, IAdminListForm, IAdminNewForm, IAdminNewParams } from '~/types/forms'
import { PromiseState } from '~/types/store'
import AdminList from '~/components/templates/AdminList.vue'

export default defineComponent({
  components: {
    AdminList,
  },

  setup(_, { root }: SetupContext) {
    const store = root.$store

    const initializeNewForm: IAdminNewParams = {
      email: '',
      password: '',
      passwordConfirmation: '',
      role: 2,
      lastName: '',
      firstName: '',
      lastNameKana: '',
      firstNameKana: '',
    }

    const search = ref<string>()
    const page = ref<number>(1)
    const itemsPerPage = ref<number>(20)
    const sortBy = ref<string>()
    const sortDesc = ref<boolean>()
    const newDialog = ref<boolean>(false)
    const newForm = reactive<IAdminNewForm>({
      params: {
        ...initializeNewForm,
      },
      options: {
        ...AdminNewOptions,
      },
    })

    const users = computed(() => store.getters['admin/getUsers'])
    const total = computed(() => store.getters['admin/getTotal'])
    const loading = computed((): boolean => {
      const status = store.getters['common/getPromiseState']
      return status === PromiseState.LOADING
    })

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

    const handleClickNewItem = (): void => {
      newDialog.value = true
      newForm.params = { ...initializeNewForm }
    }

    const handleClickCloseNewItem = (): void => {
      newDialog.value = false
    }

    const handleClickCreateItem = async (): Promise<void> => {
      CommonStore.startConnection()
      await AdminStore.createAdmin(newForm)
        .then(() => {
          newDialog.value = false
          CommonStore.showSnackbar({ color: 'info', message: '管理者を新規登録しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickEditItem = (index: number): void => {
      console.log('debug', 'editItem', index)
    }

    const handleClickDeleteItem = (index: number): void => {
      console.log('debug', 'deleteItem', index)
    }

    async function indexAdmin(): Promise<void> {
      CommonStore.startConnection()

      const form: IAdminListForm = {
        limit: itemsPerPage.value,
        offset: itemsPerPage.value * (page.value - 1),
        order: {
          by: sortBy.value || '',
          desc: sortDesc.value || false,
        },
      }

      await AdminStore.indexAdmin(form).finally(() => {
        CommonStore.endConnection()
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
      newForm,
      newDialog,
      handleClickNewItem,
      handleClickCloseNewItem,
      handleClickCreateItem,
      handleClickEditItem,
      handleClickDeleteItem,
    }
  },
})
</script>
