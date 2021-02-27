import React, { ReactElement } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import * as UiContext from '~/lib/context/ui';
import { UiContext as Context } from '~/lib/context';
import Onboarding from '~/screens/Onboarding';
import AuthRoute from '~/routes/AuthRoute';
import ServiceRoute from '~/routes/ServiceRoute';
import { RootStackParamList } from '~/types/navigation';
import { retrieve } from '~/lib/local-storage/auth-storage';
import { Status } from '~/lib/context/ui';
import { useState } from 'react';
import AppLoading from 'expo-app-loading';
const Stack = createStackNavigator<RootStackParamList>();

function OnboadingRoute(): ReactElement {
  return (
    <Stack.Navigator mode='modal' screenOptions={{ headerShown: false }}>
      <Stack.Screen name="Onboarding" component={Onboarding} />
      <Stack.Screen name="SignInSelect" component={AuthRoute} />
    </Stack.Navigator>
  );
}

// アプリ起動時、認証情報によってページ遷移
function SwitchingStatus(status: UiContext.Status): ReactElement {
  switch (status) {
  case UiContext.Status.AUTHORIZED:
    return ServiceRoute();
  default:
    return OnboadingRoute();
  }
}

export default function MainRoute(): ReactElement {
  const uiContext = React.useContext(UiContext.Context);
  const [isReady, setReady] = useState(false);
  const { setApplicationState } = React.useContext(Context);

  const _cacheResourcesAsync = async () => {
    const result = await retrieve();
    if (result !== null) {
      setApplicationState(Status.AUTHORIZED);
    }
    return result;
  };

  return (
    isReady ? (
      <NavigationContainer>
        {SwitchingStatus(uiContext.applicationState)}
      </NavigationContainer>
    ) : (
      <AppLoading
        startAsync={_cacheResourcesAsync}
        onFinish={() => setReady(true)}
        onError={console.warn}
      />
    ));
}
