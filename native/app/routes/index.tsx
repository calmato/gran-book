import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { usePrepare } from '~/hooks/usePrepare';
import * as UiContext from '~/lib/context/ui';
import AuthRoute from '~/routes/AuthRoute';
import ServiceRoute from '~/routes/ServiceRoute';
import Onboarding from '~/screens/Onboarding';
import SplashScreen from '~/screens/SplashScreen';
import { RootStackParamList } from '~/types/navigation';
const Stack = createStackNavigator<RootStackParamList>();

function OnboadingRoute(): ReactElement {
  return (
    <Stack.Navigator mode="modal" screenOptions={{ headerShown: false }}>
      <Stack.Screen name="Onboarding" component={Onboarding} />
      <Stack.Screen name="SignInSelect" component={AuthRoute} />
    </Stack.Navigator>
  );
}

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
