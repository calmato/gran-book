import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { $axios } from '~/plugins/axios'
import firebase from '~/plugins/firebase'
import { ISignInForm } from '~/types/forms'
import { IAuthResponse } from '~/types/responses'

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
  public setAuth(auth: IAuthResponse) {
    this.username = auth.username
    this.gender = auth.gender
    this.phoneNumber = auth.phoneNumber
    this.role = auth.role
    this.thumbnailUrl = auth.thumbnailUrl
    this.selfIntroduction = auth.selfIntroduction
    this.lastName = auth.lastName
    this.firstName = auth.firstName
    this.lastNameKana = auth.lastNameKana
    this.firstNameKana = auth.firstNameKana
    this.postalCode = auth.postalCode
    this.prefecture = auth.prefecture
    this.city = auth.city
    this.addressLine1 = auth.addressLine1
    this.addressLine2 = auth.addressLine2
    this.createdAt = auth.createdAt
    this.updatedAt = auth.updatedAt
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
        .then((res: IAuthResponse) => {
          this.setAuth(res)
          resolve(res.role)
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
        // TODO: 初期化のいいやり方あったらリファクタする
        const res: IAuthResponse = {
          id: '',
          email: '',
          username: '',
          gender: 0,
          phoneNumber: '',
          role: 0,
          thumbnailUrl: '',
          selfIntroduction: '',
          lastName: '',
          firstName: '',
          lastNameKana: '',
          firstNameKana: '',
          postalCode: '',
          prefecture: '',
          city: '',
          addressLine1: '',
          addressLine2: '',
          createdAt: '',
          updatedAt: '',
        }

        this.setId('')
        this.setEmail('')
        this.setEmailVerified(false)
        this.setToken('')
        this.setAuth(res)
      })
  }
}
