<template>
  <notification
    :search.sync="search"
    :page.sync="page"
    :items-per-page.sync="itemsPerPage"
    :sort-by.sync="sortBy"
    :new-form="newForm"
    :new-dialog.sync="newDialog"
    @new:open="handleClickNewItem"
    @new:close="handleClickCloseNewItem"
    @create="handleClickCreateItem"
  />
</template>

<script lang="ts">
import { defineComponent, SetupContext, ref, reactive, useAsync, computed, watch } from '@nuxtjs/composition-api'
import { AdminStore, CommonStore } from '~/store'
import { AdminNewOptions, IAdminListForm, IAdminNewForm, IAdminNewParams } from '~/types/forms'
import { PromiseState } from '~/types/store'
import Notification from '~/components/templates/Notification.vue'

export default defineComponent({
  components: {
    Notification,
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
      indexNotification()
    })

    watch(itemsPerPage, (): void => {
      indexNotification()
    })

    watch(sortBy, (): void => {
      indexNotification()
    })

    watch(sortDesc, (): void => {
      indexNotification()
    })

    useAsync(async () => {
      await indexNotification()
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
          CommonStore.showSnackbar({ color: 'info', message: 'お知らせを新規登録しました。' })
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

    async function indexNotification(): Promise<void> {
      CommonStore.startConnection()

      const form: IAdminListForm = {
        limit: itemsPerPage.value,
        offset: itemsPerPage.value * (page.value - 1),
        order: {
          by: sortBy.value || '',
          desc: sortDesc.value || false,
        },
      }

      await AdminStore.indexNotification(form).finally(() => {
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
