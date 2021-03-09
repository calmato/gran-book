export interface ICommonState {
  snackbarColor: string
  snackbarMessage: string
  promiseState: PromiseState
}

export interface ISnackbar {
  color: string
  message: string
}

export enum PromiseState {
  NONE,
  LOADING,
}
