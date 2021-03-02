export type RootStackParamList = {
  MyPage: undefined,
  Onboarding: undefined,
  SignIn:undefined,
  SignInSelect: undefined,
}

export type AuthStackParamList = {
  MyPage: undefined,
  SignIn: undefined,
  SignInSelect: undefined,
  SignUp: undefined,
  SignUpCheckEmail: { email: string | undefined },
  PasswordReset: undefined,
}

export type UserInfoStackParamList = {
  MyPage: undefined,
  OwnProfile: {username: string, selfIntroduction: string, thumbnailUrl: string, gender: number},
  AccountSetting: undefined,
  AccountEdit: undefined,
  ProfileEdit: {username: string, selfIntroduction: string, thumbnailUrl: string, gender: number},
  ContactEdit: undefined,
  EmailEdit: undefined,
  PasswordEmailEdit: undefined,
}
