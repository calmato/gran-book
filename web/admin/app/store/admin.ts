import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators'
import { $axios } from '~/plugins/axios'
import {
  CreateAdminV1Request,
  UpdateAdminContactV1Request,
  UpdateAdminPasswordV1Request,
  UpdateAdminProfileV1Request,
} from '~/types/api/admin_apiv1_request_pb'
import { AdminListV1Response, AdminThumbnailV1Response, AdminV1Response } from '~/types/api/admin_apiv1_response_pb'
import { ApiError, IErrorResponse } from '~/types/exception'
import {
  IAdminEditContactForm,
  IAdminEditProfileForm,
  IAdminEditSecurityForm,
  IAdminListForm,
  IAdminNewForm,
} from '~/types/forms'
import { IAdminState, IAdminUser } from '~/types/store'

const initialState: IAdminState = {
  user: {
    id: '',
    username: '',
    email: '',
    phoneNumber: '',
    role: 0,
    thumbnailUrl: '',
    selfIntroduction: '',
    lastName: '',
    firstName: '',
    lastNameKana: '',
    firstNameKana: '',
    createdAt: '',
    updatedAt: '',
  },
  users: [],
  total: 0,
}

@Module({
  name: 'admin',
  stateFactory: true,
  namespaced: true,
})
export default class AdminModule extends VuexModule {
  private user: IAdminUser = initialState.user
  private users: IAdminUser[] = initialState.users
  private total: number = initialState.total

  public get getUser(): IAdminUser {
    return this.user
  }

  public get getUsers(): IAdminUser[] {
    return this.users
  }

  public get getTotal(): number {
    return this.total
  }

  @Mutation
  private setUser(user: IAdminUser): void {
    this.user = user
  }

  @Mutation
  private setUsers(users: IAdminUser[]): void {
    this.users = users
  }

  @Mutation
  private addUser(user: IAdminUser): void {
    this.users.push(user)
    this.total = this.total + 1
  }

  @Mutation
  private setTotal(total: number): void {
    this.total = total
  }

  @Action({})
  public factory(): void {
    this.setUsers(initialState.users)
    this.setTotal(initialState.total)
  }

  @Action({ rawError: true })
  public indexAdmin(payload: IAdminListForm): Promise<void> {
    const { limit, offset, order } = payload

    let query: string = `limit=${limit}&offset=${offset}`

    if (order?.by) {
      const direction: string = order.desc ? 'desc' : 'asc'
      query += `&by=${order.by}&direction=${direction}`
    }

    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      $axios
        .$get(`/v1/admin?${query}`)
        .then((res: AdminListV1Response.AsObject) => {
          const { usersList, total } = res
          const data: IAdminUser[] = usersList.map((user: AdminListV1Response.User.AsObject): IAdminUser => {
            return { ...user } as IAdminUser
          })
          this.setUsers(data)
          this.setTotal(total)
          resolve()
        })
        .catch((err: any) => {
          const { data, status }: IErrorResponse = err.response
          reject(new ApiError(status, data.message, data))
        })
    })
  }

  @Action({ rawError: true })
  public showAdmin(userId: string): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      $axios
        .$get(`/v1/admin/${userId}`)
        .then((res: AdminV1Response.AsObject) => {
          const data: IAdminUser = { ...res }
          this.setUser(data)
          resolve()
        })
        .catch((err: any) => {
          const { data, status }: IErrorResponse = err.response
          reject(new ApiError(status, data.message, data))
        })
    })
  }

  @Action({ rawError: true })
  public createAdmin(payload: IAdminNewForm): Promise<void> {
    const { email, password, passwordConfirmation, role, lastName, firstName, lastNameKana, firstNameKana } =
      payload.params

    const req: CreateAdminV1Request.AsObject = {
      username: `${lastName} ${firstName}`,
      email,
      password,
      passwordConfirmation,
      role,
      lastName,
      firstName,
      lastNameKana,
      firstNameKana,
    }

    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      $axios
        .$post('/v1/admin', req)
        .then((res: AdminV1Response.AsObject) => {
          const data: IAdminUser = { ...res }
          this.addUser(data)
          resolve()
        })
        .catch((err: any) => {
          const { data, status }: IErrorResponse = err.response
          reject(new ApiError(status, data.message, data))
        })
    })
  }

  @Action({ rawError: true })
  public updateProfile({ userId, form }: { userId: string; form: IAdminEditProfileForm }): Promise<void> {
    const { role, lastName, firstName, lastNameKana, firstNameKana, thumbnailUrl } = form.params

    const req: UpdateAdminProfileV1Request.AsObject = {
      role,
      username: '', // TODO: refactor
      lastName,
      firstName,
      lastNameKana,
      firstNameKana,
      thumbnailUrl: String(thumbnailUrl) || '',
    }

    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      $axios
        .$patch(`/v1/admin/${userId}/profile`, req)
        .then((res: AdminV1Response.AsObject) => {
          const data: IAdminUser = { ...res }
          this.addUser(data)
          resolve()
        })
        .catch((err: any) => {
          const { data, status }: IErrorResponse = err.response
          reject(new ApiError(status, data.message, data))
        })
    })
  }

  @Action({ rawError: true })
  public updateContact({ userId, form }: { userId: string; form: IAdminEditContactForm }): Promise<void> {
    const { email, phoneNumber } = form.params

    const req: UpdateAdminContactV1Request.AsObject = {
      email,
      phoneNumber,
    }

    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      $axios
        .$patch(`/v1/admin/${userId}/contact`, req)
        .then((res: AdminV1Response.AsObject) => {
          const data: IAdminUser = { ...res }
          this.addUser(data)
          resolve()
        })
        .catch((err: any) => {
          const { data, status }: IErrorResponse = err.response
          reject(new ApiError(status, data.message, data))
        })
    })
  }

  @Action({ rawError: true })
  public updatePassword({ userId, form }: { userId: string; form: IAdminEditSecurityForm }): Promise<void> {
    const { password, passwordConfirmation } = form.params

    const req: UpdateAdminPasswordV1Request.AsObject = {
      password,
      passwordConfirmation,
    }

    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      $axios
        .$patch(`/v1/admin/${userId}/password`, req)
        .then((res: AdminV1Response.AsObject) => {
          const data: IAdminUser = { ...res }
          this.addUser(data)
          resolve()
        })
        .catch((err: any) => {
          const { data, status }: IErrorResponse = err.response
          reject(new ApiError(status, data.message, data))
        })
    })
  }

  @Action({ rawError: true })
  public uploadThumbnail({ userId, file }: { userId: string; file: File }): Promise<string> {
    const config = {
      headers: {
        'content-type': 'multipart/form-data',
      },
    }

    const params = new FormData()
    params.append('thumbnail', file)

    return new Promise((resolve: (thumbnailUrl: string) => void, reject: (reason: Error) => void) => {
      $axios
        .$post(`/v1/admin/${userId}/thumbnail`, params, config)
        .then((res: AdminThumbnailV1Response.AsObject) => {
          resolve(res.thumbnailUrl)
        })
        .catch((err: any) => {
          const { data, status }: IErrorResponse = err.response
          reject(new ApiError(status, data.message, data))
        })
    })
  }

  @Action({ rawError: true })
  public deleteAdmin(userId: string): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      $axios
        .$delete(`/v1/admin/${userId}`)
        .then(() => {
          resolve()
        })
        .catch((err: any) => {
          const { data, status }: IErrorResponse = err.response
          reject(new ApiError(status, data.message, data))
        })
    })
  }
}
