import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import firebase from '~/plugins/firebase'
import { AuthState } from '~/types/store'

@Module({
  name: 'auth',
  stateFactory: true,
  namespaced: true,
})
export default class AuthModule extends VuexModule implements AuthState {
  private id = ''
  private email = ''
  private emailVerified = false
  private token = ''

  public get getId() {
    return this.id
  }

  public get getEmail() {
    return this.email
  }

  public get getEmailVerified() {
    return this.emailVerified
  }

  public get getToken() {
    return this.token
  }

  @Mutation
  public setId(id: string) {
    this.id = id
  }

  @Mutation
  public setEmail(email: string) {
    this.email = email
  }

  @Mutation
  public setEmailVerified(emailVerified: boolean) {
    this.emailVerified = emailVerified
  }

  @Mutation
  public setToken(token: string) {
    this.token = token
  }

  @Action({ rawError: true })
  public authentication(): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      firebase.auth().onAuthStateChanged((res: any) => {
        if (res) {
          this.context.commit('setId', res.uid)
          this.context.commit('setEmail', res.email)
          this.context.commit('setEmailVerified', res.emailVerified)

          resolve()
        } else {
          reject(new Error())
        }
      })
    })
  }

  @Action({ rawError: true })
  public setIdToken(): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      firebase
        .auth()
        .currentUser?.getIdToken(true)
        .then((token: string) => {
          this.context.commit('setToken', token)
          resolve()
        })
        .catch((err: any) => {
          reject(new Error(err))
        })
    })
  }

  @Action({ rawError: true })
  public loginWithEmailAndPassword(payload): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      firebase
        .auth()
        .signInWithEmailAndPassword(payload.email, payload.password)
        .then(() => {
          this.context.dispatch('authentication')
          resolve()
        })
        .catch((error: any) => {
          reject(error)
        })
    })
  }
}
