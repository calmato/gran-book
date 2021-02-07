import React, { ReactElement } from 'react';
import { Provider } from 'react-redux';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import { ThemeProvider } from 'react-native-elements';
import Onboarding from '~/screens/Onboarding';
import { RootStackParamList } from '~/types/navigation';
import { THEME } from '~~/constants/theme';
import AuthRoute from '~/routes/AuthRoute';
import store from '~/store';

const Stack = createStackNavigator<RootStackParamList>();

const App = function App(): ReactElement {
  return (
    <Provider store={store}>
      <ThemeProvider theme={THEME}>
        <NavigationContainer>
          <Stack.Navigator
            mode='modal'
            screenOptions={{
              headerShown: false
            }}
          >
            <Stack.Screen name="Onboarding" component={Onboarding} />
            <Stack.Screen name="SignInSelect" component={AuthRoute} />
          </Stack.Navigator>
        </NavigationContainer>
      </ThemeProvider>
    </Provider>
  );
};

export default App;
