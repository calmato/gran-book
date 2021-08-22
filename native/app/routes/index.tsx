import { NavigationContainer } from '@react-navigation/native';
import React, { ReactElement } from 'react';
import { usePrepare } from '~/hooks/usePrepare';
import * as UiContext from '~/lib/context/ui';
import OnboadingRoute from '~/routes/OnboadingRoute';
import ServiceRoute from '~/routes/ServiceRoute';
import SplashScreen from '~/screens/SplashScreen';

// アプリ起動時、認証情報によってページ遷移
function SwitchingStatus(status: UiContext.Status): ReactElement {
  switch (status) {
    // 認証済みならばサービスルートを表示
    case UiContext.Status.AUTHORIZED:
      return ServiceRoute();
    // 未認証あるいはそれ以外ならば認証・登録ルートを表示
    case UiContext.Status.UN_AUTHORIZED:
    default:
      return OnboadingRoute();
  }
}

export default function MainRoute(): ReactElement {
  const uiContext = usePrepare();

  // ローディング中はSplashScreenを表示する
  return uiContext.applicationState === UiContext.Status.LOADING ? (
    <SplashScreen />
  ) : (
    <NavigationContainer>{SwitchingStatus(uiContext.applicationState)}</NavigationContainer>
  );
}
