import React, { ReactElement, useEffect } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import * as UiContext from '~/lib/context/ui';
import { UiContext as Context } from '~/lib/context';
import Onboarding from '~/screens/Onboarding';
import AuthRoute from '~/routes/AuthRoute';
import ServiceRoute from '~/routes/ServiceRoute';
import { RootStackParamList } from '~/types/navigation';
import { authenticationAsync, getAuthAsync } from '~/store/usecases';
import { useReduxDispatch } from '~/store/modules';
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
    case UiContext.Status.AUTHORIZED:
      return ServiceRoute();
    case UiContext.Status.UN_AUTHORIZED:
    default:
      return OnboadingRoute();
  }
}

export default function MainRoute(): ReactElement {
  const uiContext = React.useContext(UiContext.Context);
  const { setApplicationState } = React.useContext(Context);
  const dispatch = useReduxDispatch();

  useEffect(() => {
    async function prepare() {
      await actions.authentication()
        .then(() => {
          return actions.getAuth();
        })
        .then(() => {
          setApplicationState(UiContext.Status.AUTHORIZED);
        })
        .catch(() => {
          setApplicationState(UiContext.Status.UN_AUTHORIZED);
        });
    }

    prepare();
  }, []);

  const actions = React.useMemo(
    () => ({
      authentication(): Promise<void> {
        return dispatch(authenticationAsync());
      },
      getAuth(): Promise<void> {
        return dispatch(getAuthAsync());
      }
    }),
    [dispatch],
  );

  return (
    <NavigationContainer>
      {SwitchingStatus(uiContext.applicationState)}
    </NavigationContainer>
  );
}
