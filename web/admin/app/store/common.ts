import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators'
import { MESSAGE } from '~/constants/error'
import { ApiError } from '~/types/exception'
import { ICommonState, ISnackbar, PromiseState } from '~/types/store'

const initialState: ICommonState = {
  snackbarColor: 'info',
  snackbarMessage: '',
  promiseState: PromiseState.NONE,
}

@Module({
  name: 'common',
  stateFactory: true,
  namespaced: true,
})
export default class CommonModule extends VuexModule {
  private snackbarColor: string = initialState.snackbarColor
  private snackbarMessage: string = initialState.snackbarMessage
  private promiseState: PromiseState = initialState.promiseState

  public get getSnackbarColor(): string {
    return this.snackbarColor
  }

  public get getSnackbarMessage(): string {
    return this.snackbarMessage
  }

  public get getPromiseState(): PromiseState {
    return this.promiseState
  }

  @Mutation
  private setSnackbarColor(color: string): void {
    this.snackbarColor = color
  }

  @Mutation
  private setSnackbarMessage(message: string): void {
    this.snackbarMessage = message
  }

  @Mutation
  private setPromiseState(promiseState: PromiseState): void {
    this.promiseState = promiseState
  }

  @Action({})
  public factory(): void {
    this.setSnackbarColor(initialState.snackbarColor)
    this.setSnackbarMessage(initialState.snackbarMessage)
    this.setPromiseState(initialState.promiseState)
  }

  @Action({ rawError: true })
  public showSnackbar(payload: ISnackbar): void {
    this.setSnackbarColor(payload.color)
    this.setSnackbarMessage(payload.message)
  }

  @Action({ rawError: true })
  public showSuccessInSnackbar(message: string): void {
    this.setSnackbarColor('success')
    this.setSnackbarMessage(message)
  }

  @Action({ rawError: true })
  public showErrorInSnackbar(err: Error): void {
    this.setSnackbarColor('error')

    if (err instanceof ApiError) {
      this.setSnackbarMessage(getApiErrorMessage(err))
    } else {
      this.setSnackbarMessage(MESSAGE.UNEXPEXTED_ERROR)
    }
  }

  @Action({ rawError: true })
  public hiddenSnackbar(): void {
    this.setSnackbarColor(initialState.snackbarColor)
    this.setSnackbarMessage(initialState.snackbarMessage)
  }

  @Action({ rawError: true })
  public startConnection(): void {
    this.setPromiseState(PromiseState.LOADING)
  }

  @Action({ rawError: true })
  public endConnection(): void {
    this.setPromiseState(PromiseState.NONE)
  }
}

function getApiErrorMessage(err: ApiError): string {
  switch (err.status) {
    case 400:
      return MESSAGE.BAD_REQUEST
    case 401:
      return MESSAGE.UNAUTHORIZED
    case 403:
      return MESSAGE.FORBIDDEN
    case 404:
      return MESSAGE.PROCESS_FAILED
    case 409:
      return MESSAGE.CONFLICT
    case 500:
    case 501:
    case 503:
      return MESSAGE.SERVER_ERROR
    case 504:
      return MESSAGE.TIMEOUT
    default:
      return MESSAGE.UNEXPEXTED_ERROR
  }
}
