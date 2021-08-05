import { createStackNavigator } from '@react-navigation/stack';
import React from 'react';
import AuthRoute from './AuthRoute';
import Onboarding from '~/screens/Onboarding';
import { RootStackParamList } from '~/types/navigation';

export default function OnboadingRoute() {
  const Stack = createStackNavigator<RootStackParamList>();

  return (
    <Stack.Navigator mode="modal" screenOptions={{ headerShown: false }}>
      <Stack.Screen name="Onboarding" component={Onboarding} />
      <Stack.Screen name="SignInSelect" component={AuthRoute} />
    </Stack.Navigator>
  );
}
