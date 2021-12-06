import { NavigationContainer } from '@react-navigation/native';
import React, { ReactElement, useContext } from 'react';
import { Status, UiContext } from '~/context/ui';
import OnboadingRoute from '~/routes/OnboadingRoute';
import ServiceRoute from '~/routes/ServiceRoute';
import SplashScreen from '~/screens/SplashScreen';

// アプリ起動時、認証情報によってページ遷移
function SwitchingStatus(status: Status): ReactElement {
  switch (status) {
    // 認証済みならばサービスルートを表示
    case Status.AUTHORIZED:
      return ServiceRoute();
    // 未認証あるいはそれ以外ならば認証・登録ルートを表示
    case Status.UN_AUTHORIZED:
    default:
      return OnboadingRoute();
  }
}

export default function MainRoute(): ReactElement {
  const uiContext = useContext(UiContext);

  // ローディング中はSplashScreenを表示する
  return uiContext.applicationState === Status.LOADING ? (
    <SplashScreen />
  ) : (
    <NavigationContainer>{SwitchingStatus(uiContext.applicationState)}</NavigationContainer>
  );
}
