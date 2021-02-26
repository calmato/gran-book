import React, { ReactElement } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import * as UiContext from '~/lib/context/ui';
import Onboarding from '~/screens/Onboarding';
import AuthRoute from '~/routes/AuthRoute';
import ServiceRoute from '~/routes/ServiceRoute';
import { RootStackParamList } from '~/types/navigation';

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

  return (
    <NavigationContainer>
      {SwitchingStatus(uiContext.applicationState)}
    </NavigationContainer>
  );
}
