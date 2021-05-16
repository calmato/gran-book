<template>
  <admin-list
    :search.sync="search"
    :page.sync="page"
    :items-per-page.sync="itemsPerPage"
    :sort-by.sync="sortBy"
    :sort-desc.sync="sortDesc"
    :loading="loading"
    :role="role"
    :users="users"
    :total="total"
    :new-form="newForm"
    :new-dialog.sync="newDialog"
    :edit-form="editForm"
    :edit-dialog.sync="editDialog"
    @new:open="handleClickNewItem"
    @new:close="handleClickCloseNewItem"
    @edit="handleClickEditItem"
    @edit:close="handleClickCloseEditItem"
    @create="handleClickCreateItem"
    @update="handleClickUpdateItem"
    @delete="handleClickDeleteItem"
  />
</template>

<script lang="ts">
import { defineComponent, SetupContext, ref, reactive, useAsync, computed, watch } from '@nuxtjs/composition-api'
import { AdminStore, CommonStore } from '~/store'
import {
  AdminEditOptions,
  AdminNewOptions,
  IAdminEditForm,
  IAdminEditParams,
  IAdminListForm,
  IAdminNewForm,
  IAdminNewParams,
} from '~/types/forms'
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

    const initializeEditForm: IAdminEditParams = {
      email: '',
      phoneNumber: '',
      role: 2,
      lastName: '',
      firstName: '',
      lastNameKana: '',
      firstNameKana: '',
      thumbnail: null,
      thumbnailUrl: '',
    }

    const search = ref<string>()
    const page = ref<number>(1)
    const itemsPerPage = ref<number>(20)
    const sortBy = ref<string>()
    const sortDesc = ref<boolean>()
    const newDialog = ref<boolean>(false)
    const editDialog = ref<boolean>(false)
    const editUserId = ref<string>('')
    const newForm = reactive<IAdminNewForm>({
      params: {
        ...initializeNewForm,
      },
      options: {
        ...AdminNewOptions,
      },
    })
    const editForm = reactive<IAdminEditForm>({
      params: {
        ...initializeEditForm,
      },
      options: {
        ...AdminEditOptions,
      },
    })

    const role = computed(() => store.getters['auth/getRole'])
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

    const uploadThumbnail = (file: File | null): Promise<string> => {
      if (!file) {
        return Promise.resolve('')
      }

      return AdminStore.uploadThumbnail({ userId: editUserId.value, file })
        .then((res: string) => {
          return res
        })
        .catch((err: Error) => {
          throw err
        })
    }

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
      if (!users.value || users.value.length <= index) {
        // TODO: show alert
        return
      }

      const user = users.value[index]
      if (!user) {
        // TODO: show alert
        return
      }

      editDialog.value = true
      editUserId.value = user.id
      editForm.params = {
        ...initializeEditForm,
        email: user.email,
        phoneNumber: user.phoneNumber,
        role: user.role,
        lastName: user.lastName,
        firstName: user.firstName,
        lastNameKana: user.lastNameKana,
        firstNameKana: user.firstNameKana,
        thumbnailUrl: user.thumbnailUrl,
      }
    }

    const handleClickCloseEditItem = (): void => {
      editDialog.value = false
    }

    const handleClickUpdateItem = async (): Promise<void> => {
      CommonStore.startConnection()
      await uploadThumbnail(editForm.params.thumbnail)
        .then((res: string) => {
          if (res !== '') {
            editForm.params.thumbnailUrl = res
          }

          return AdminStore.updateAdmin({ userId: editUserId.value, form: editForm })
        })
        .then(() => {
          editDialog.value = false
          CommonStore.showSnackbar({ color: 'info', message: '管理者を更新しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
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
      role,
      users,
      total,
      newForm,
      newDialog,
      editForm,
      editDialog,
      handleClickNewItem,
      handleClickCloseNewItem,
      handleClickCreateItem,
      handleClickEditItem,
      handleClickCloseEditItem,
      handleClickUpdateItem,
      handleClickDeleteItem,
    }
  },
})
</script>
