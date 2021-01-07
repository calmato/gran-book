import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import firebase from '~/plugins/firebase'
import { SignInForm } from '~/types/forms'

@Module({
  name: 'auth',
  stateFactory: true,
  namespaced: true,
})
export default class AuthModule extends VuexModule {
  private id: string = ''
  private email: string = ''
  private emailVerified: boolean = false
  private token: string = ''

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
  public authorization(): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      firebase.auth().onAuthStateChanged((res: any) => {
        if (res) {
          this.setId(res.uid)
          this.setEmail(res.email)
          this.setEmailVerified(res.emailVerified)
          this.getIdToken()

          resolve()
        } else {
          reject(new Error())
        }
      })
    })
  }

  @Action({ rawError: true })
  public getIdToken(): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      firebase
        .auth()
        .currentUser?.getIdToken(true)
        .then((token: string) => {
          this.setToken(token)
          resolve()
        })
        .catch((err: any) => {
          reject(new Error(err))
        })
    })
  }

  @Action({ rawError: true })
  public loginWithEmailAndPassword(payload: SignInForm): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      firebase
        .auth()
        .signInWithEmailAndPassword(payload.email, payload.password)
        .then(() => {
          this.authorization()
          resolve()
        })
        .catch((error: any) => {
          reject(error)
        })
    })
  }
}
