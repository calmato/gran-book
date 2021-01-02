import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { ThemeProvider } from 'react-native-elements';
import Onboarding from '~/screens/Onboarding';
import { StackParamList } from '~/types/navigation';
import { THEME } from '~~/constants/theme';
import AuthRoute from '~/routes/AuthRoute';

const Stack = createStackNavigator<StackParamList>();

const App = function App(): ReactElement {
  return (
    <ThemeProvider theme={THEME}>
      <NavigationContainer>
        <Stack.Navigator
          mode='modal'
          screenOptions={{
            headerShown: false
          }}
        >
          <Stack.Screen name="Onboarding" component={Onboarding} />
          <Stack.Screen name="SignIn" component={AuthRoute} />
        </Stack.Navigator>
      </NavigationContainer>
    </ThemeProvider>
  );
};

export default App;
