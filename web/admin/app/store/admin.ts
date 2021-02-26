import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators'
import { IAdminState, IAdminUser } from '~/types/store'

const initialState: IAdminState = {
  users: [
    {
      id: '1',
      username: 'k.hamada',
      email: 'k.hamada@calmato.com',
      phoneNumber: '000-0000-0000',
      role: 1,
      thumbnailUrl: '',
      selfIntroduction: '',
      lastName: '濵田',
      firstName: '広大',
      lastNameKana: 'はまだ',
      firstNameKana: 'こうだい',
      createdAt: '2020-01-01 00:00:00',
      updatedAt: '2020-01-01 00:00:00',
    },
    {
      id: '2',
      username: 't.nishikawa',
      email: 't.nishikawa@calmato.com',
      phoneNumber: '111-1111-1111',
      role: 1,
      thumbnailUrl: '',
      selfIntroduction: '',
      lastName: '西川',
      firstName: '西川',
      lastNameKana: 'にしかわ',
      firstNameKana: 'ただし',
      createdAt: '2020-01-01 00:00:00',
      updatedAt: '2020-01-01 00:00:00',
    },
    {
      id: '3',
      username: 'y.yamada',
      email: 'y.yamada@calmato.com',
      phoneNumber: '000-1111-1111',
      role: 2,
      thumbnailUrl: '',
      selfIntroduction: '',
      lastName: '山田',
      firstName: '侑樹',
      lastNameKana: 'やまだ',
      firstNameKana: 'ゆうき',
      createdAt: '2020-01-01 00:00:00',
      updatedAt: '2020-01-01 00:00:00',
    },
    {
      id: '4',
      username: 'a.inatomi',
      email: 'a.inatomi@calmato.com',
      phoneNumber: '111-0000-0000',
      role: 3,
      thumbnailUrl: '',
      selfIntroduction: '',
      lastName: 'Inatomi',
      firstName: 'Atsuhide',
      lastNameKana: 'いなとみ',
      firstNameKana: 'あつひで',
      createdAt: '2020-01-01 00:00:00',
      updatedAt: '2020-01-01 00:00:00',
    },
  ],
}

@Module({
  name: 'admin',
  stateFactory: true,
  namespaced: true,
})
export default class AdminModule extends VuexModule {
  private users: IAdminUser[] = initialState.users

  public get getUsers(): IAdminUser[] {
    return this.users
  }

  @Mutation
  private addUser(user: IAdminUser): void {
    this.users.push(user)
  }

  @Action({ rawError: true })
  public createUser(): Promise<void> {
    return new Promise((resolve: () => void) => {
      const user: IAdminUser = {
        id: '5',
        username: 'a.inatomi',
        email: 'u.new@calmato.com',
        phoneNumber: '123-1234-1234',
        role: 3,
        thumbnailUrl: '',
        selfIntroduction: '',
        lastName: 'テスト',
        firstName: 'ユーザー',
        lastNameKana: 'てすと',
        firstNameKana: 'ゆーざー',
        createdAt: '2020-01-01 00:00:00',
        updatedAt: '2020-01-01 00:00:00',
      }
      this.addUser(user)
      console.log('debug', this.users)
      resolve()
    })
  }
}
