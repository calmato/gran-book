import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'

@Module({
  name: 'common',
  stateFactory: true,
  namespaced: true,
})
export default class CommonModule extends VuexModule {
  private currentPath: string = ''

  public get getCurrentPath(): string {
    return this.currentPath
  }

  @Mutation
  private setCurrentPath(path: string) {
    this.currentPath = path
  }

  @Action({ rawError: true })
  public updateCurrentPath(path: string): void {
    this.setCurrentPath(path)
  }
}
