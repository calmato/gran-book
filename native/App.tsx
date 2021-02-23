import React, { ReactElement } from 'react';
import { Provider } from 'react-redux';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import { ThemeProvider } from 'react-native-elements';
import Onboarding from '~/screens/Onboarding';
import { RootStackParamList } from '~/types/navigation';
import { THEME } from '~~/constants/theme';
import AuthRoute from '~/routes/AuthRoute';
import ServiceRoute from '~/routes/ServiceRoute';
import store from '~/store';

const Stack = createStackNavigator<RootStackParamList>();

const App = function App(): ReactElement {

  // ログインの実装はまだなのでisLoggedInでOnBoardingとログイン後の画面の出し分けを一時的に行う
  // true -> ログイン後の画面を表示 / false -> OnBoardingの画面を表示
  const isLoggedIn = false;

  let navigator: ReactElement;

  if (isLoggedIn) {
    navigator = ServiceRoute();
  } else {
    navigator = <Stack.Navigator
      mode='modal'
      screenOptions={{
        headerShown: false
      }}
    >
      <Stack.Screen name="Onboarding" component={Onboarding} />
      <Stack.Screen name="SignInSelect" component={AuthRoute} />
    </Stack.Navigator>;
  }

  return (
    <Provider store={store}>
      <ThemeProvider theme={THEME}>
        <NavigationContainer>
          {navigator}
        </NavigationContainer>
      </ThemeProvider>
    </Provider>
  );
};

export default App;
