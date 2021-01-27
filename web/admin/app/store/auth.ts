import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { $axios } from '~/plugins/axios'
import firebase from '~/plugins/firebase'
import { ISignInForm } from '~/types/forms'
import { IShowAuthResponse } from '~/types/responses'

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
  private username: string = ''
  private gender: number = 0
  private phoneNumber: string = ''
  private role: number = 0
  private thumbnailUrl: string = ''
  private selfIntroduction: string = ''
  private lastName: string = ''
  private firstName: string = ''
  private lastNameKana: string = ''
  private firstNameKana: string = ''
  private postalCode: string = ''
  private prefecture: string = ''
  private city: string = ''
  private addressLine1: string = ''
  private addressLine2: string = ''
  private createdAt: string = ''
  private updatedAt: string = ''

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

  public get getUsername() {
    return this.username
  }

  public get getRole() {
    return this.role
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

  @Mutation
  public setUsername(username: string) {
    this.username = username
  }

  @Mutation
  public setRole(role: number) {
    this.role = role
  }

  @Action({ rawError: true })
  public authorization(): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      firebase.auth().onAuthStateChanged(async (res: any) => {
        if (res) {
          const { uid, email, emailVerified } = res
          if (emailVerified) {
            this.setId(uid)
            this.setEmail(email)
            await this.getIdToken()
          } else {
            this.sendEmailVerification()
          }

          this.setEmailVerified(res.emailVerified)
          resolve()
        } else {
          reject(new Error('unauthorized'))
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
        .catch((err: Error) => {
          reject(err)
        })
    })
  }

  @Action({ rawError: true })
  public sendEmailVerification(): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      firebase
        .auth()
        .currentUser?.sendEmailVerification()
        .then(() => resolve())
        .catch((err: Error) => reject(err))
    })
  }

  @Action({ rawError: true })
  public loginWithEmailAndPassword(payload: ISignInForm): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      firebase
        .auth()
        .signInWithEmailAndPassword(payload.email, payload.password)
        .then(async () => {
          await this.authorization()
          resolve()
        })
        .catch((err: Error) => {
          reject(err)
        })
    })
  }

  @Action({ rawError: true })
  public showAuth(): Promise<number> {
    return new Promise((resolve: (role: number) => void, reject: (reason: Error) => void) => {
      $axios
        .$get('/v1/auth')
        .then((res: IShowAuthResponse) => {
          const { username, role } = res
          this.setUsername(username)
          this.setRole(role)
          resolve(role)
        })
        .catch((err: Error) => {
          reject(err)
        })
    })
  }

  @Action({ rawError: true })
  public logout(): void {
    firebase
      .auth()
      .signOut()
      .finally(() => {
        this.setId('')
        this.setEmail('')
        this.setEmailVerified(false)
        this.setToken('')
        this.setUsername('')
        this.setRole(0)
      })
  }
}
