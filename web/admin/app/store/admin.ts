import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators'
import { $axios } from '~/plugins/axios'
import { ApiError } from '~/types/exception'
import { IAdminListForm, IAdminNewForm } from '~/types/forms'
import { IAdminNewRequest } from '~/types/requests'
import { IAdminListResponse, IAdminListResponseUser, IAdminResponse, IErrorResponse } from '~/types/responses'
import { IAdminState, IAdminUser } from '~/types/store'

const initialState: IAdminState = {
  users: [],
  total: 0,
}

@Module({
  name: 'admin',
  stateFactory: true,
  namespaced: true,
})
export default class AdminModule extends VuexModule {
  private users: IAdminUser[] = initialState.users
  private total: number = initialState.total

  public get getUsers(): IAdminUser[] {
    return this.users
  }

  public get getTotal(): number {
    return this.total
  }

  @Mutation
  private setUsers(users: IAdminUser[]): void {
    this.users = users
  }

  @Mutation
  private addUser(user: IAdminUser): void {
    this.users.push(user)
  }

  @Mutation
  private setTotal(total: number): void {
    this.total = total
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
        .then((res: IAdminListResponse) => {
          const { users, total } = res
          const data: IAdminUser[] = users.map(
            (user: IAdminListResponseUser): IAdminUser => {
              return { ...user } as IAdminUser
            }
          )
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
  public createUser(payload: IAdminNewForm): Promise<void> {
    const {
      username,
      email,
      password,
      passwordConfirmation,
      role,
      lastName,
      firstName,
      lastNameKana,
      firstNameKana,
    } = payload.params

    const req: IAdminNewRequest = {
      username,
      email,
      password,
      passwordConfirmation,
      role,
      lastName,
      firstName,
      lastNameKana,
      firstNameKana,
    }

    return new Promise((resolve: () => void, reject: (reason: ApiError) => void) => {
      $axios
        .$post('/v1/admin', req)
        .then((res: IAdminResponse) => {
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
}
