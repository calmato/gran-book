import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { ThemeProvider } from 'react-native-elements';
import SignIn from '~/screens/SignIn';
import Onboarding from '~/screens/Onboarding';
import { StackParamList } from '~/types/navigation';
import { THEME } from '~~/constants/theme';

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
          <Stack.Screen name="SignIn" component={SignIn} />
        </Stack.Navigator>
      </NavigationContainer>
    </ThemeProvider>
  );
};

export default App;
