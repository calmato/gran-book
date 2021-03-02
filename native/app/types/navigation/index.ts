// todo: 相対パスでのインポートを直す
import { ISearchResponse, ISearchResultItem } from '../response/search';

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

export type HomeTabStackPramList = {
  Home: undefined,
  SearchResult: { keyword: string, results: ISearchResponse },
  SearchResultBookShow: { book: ISearchResultItem }
  // SearchResultBookShow: undefined
}
