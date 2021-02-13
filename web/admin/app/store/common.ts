import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators'
import { ISnackbar } from '~/types/forms'
import { ICommonState } from '~/types/store'

const initialState: ICommonState = {
  snackbarColor: 'info',
  snackbarMessage: '',
}

@Module({
  name: 'common',
  stateFactory: true,
  namespaced: true,
})
export default class CommonModule extends VuexModule {
  private snackbarColor: string = initialState.snackbarColor
  private snackbarMessage: string = initialState.snackbarMessage

  public get getSnackbarColor(): string {
    return this.snackbarColor
  }

  public get getSnackbarMessage(): string {
    return this.snackbarMessage
  }

  @Mutation
  private setSnackbarColor(color: string): void {
    this.snackbarColor = color
  }

  @Mutation
  private setSnackbarMessage(message: string): void {
    this.snackbarMessage = message
  }

  @Action({ rawError: true })
  public showSnackbar(payload: ISnackbar): void {
    this.setSnackbarColor(payload.color)
    this.setSnackbarMessage(payload.message)
  }

  @Action({ rawError: true })
  public hiddenSnackbar(): void {
    this.setSnackbarColor(initialState.snackbarColor)
    this.setSnackbarMessage(initialState.snackbarMessage)
  }
}
