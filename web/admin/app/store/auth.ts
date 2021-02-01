import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { $axios } from '~/plugins/axios'
import firebase from '~/plugins/firebase'
import { ISignInForm } from '~/types/forms'
import { IAuthResponse } from '~/types/responses'
import { IAuthState, IAuthProfile } from '~/types/store'

const initialState: IAuthState = {
  id: '',
  email: '',
  emailVerified: false,
  token: '',
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

@Module({
  name: 'auth',
  stateFactory: true,
  namespaced: true,
})
export default class AuthModule extends VuexModule {
  private id: string = initialState.id
  private email: string = initialState.email
  private emailVerified: boolean = initialState.emailVerified
  private token: string = initialState.token
  private username: string = initialState.username
  private gender: number = initialState.gender
  private phoneNumber: string = initialState.phoneNumber
  private role: number = initialState.role
  private thumbnailUrl: string = initialState.thumbnailUrl
  private selfIntroduction: string = initialState.selfIntroduction
  private lastName: string = initialState.lastName
  private firstName: string = initialState.firstName
  private lastNameKana: string = initialState.lastNameKana
  private firstNameKana: string = initialState.firstNameKana
  private postalCode: string = initialState.postalCode
  private prefecture: string = initialState.prefecture
  private city: string = initialState.city
  private addressLine1: string = initialState.addressLine1
  private addressLine2: string = initialState.addressLine2
  private createdAt: string = initialState.createdAt
  private updatedAt: string = initialState.updatedAt

  public get getEmail(): string {
    return this.email
  }

  public get getToken(): string {
    return this.token
  }

  public get getUsername(): string {
    return this.username
  }

  public get getPhoneNumber(): string {
    return this.phoneNumber
  }

  public get getThumbnailUrl(): string {
    return this.thumbnailUrl ? this.thumbnailUrl : '/thumbnail.png'
  }

  public get getSelfIntroduction(): string {
    return this.selfIntroduction
  }

  public get getName(): string {
    return `${this.lastName} ${this.firstName}`
  }

  public get getNameKana(): string {
    return `${this.lastNameKana} ${this.firstNameKana}`
  }

  @Mutation
  private setId(id: string): void {
    this.id = id
  }

  @Mutation
  private setEmail(email: string): void {
    this.email = email
  }

  @Mutation
  private setEmailVerified(emailVerified: boolean): void {
    this.emailVerified = emailVerified
  }

  @Mutation
  private setToken(token: string): void {
    this.token = token
  }

  @Mutation
  private setProfile(auth: IAuthProfile): void {
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
          const data: IAuthProfile = { ...res }
          this.setProfile(data)
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
        const profile: IAuthProfile = { ...initialState }

        this.setId(initialState.id)
        this.setEmail(initialState.email)
        this.setEmailVerified(initialState.emailVerified)
        this.setToken(initialState.token)
        this.setProfile(profile)
      })
  }
}
