import { ISearchResponse, ISearchResultItem } from '~/types/response/external/rakuten-books';

export type RootStackParamList = {
  MyPage: undefined;
  Onboarding: undefined;
  SignIn: undefined;
  SignInSelect: undefined;
};

export type AuthStackParamList = {
  MyPage: undefined;
  SignIn: undefined;
  SignInSelect: undefined;
  SignUp: undefined;
  SignUpCheckEmail: { email: string | undefined };
  PasswordReset: undefined;
};

export type HomeTabStackPramList = {
  Home: undefined;
  SearchResult: { keyword: string; results: ISearchResponse };
  SearchResultBookShow: { book: ISearchResultItem };
};

export type UserInfoStackParamList = {
  MyPage: undefined;
  OwnProfile: undefined;
  AccountSetting: undefined;
  AccountEdit: undefined;
  ProfileEdit: undefined;
  ContactEdit: undefined;
  EmailEdit: undefined;
  PasswordEmailEdit: undefined;
};
