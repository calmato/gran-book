import { createStackNavigator } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import SignInSelect from '~/screens/SignInSelect';
import SignUp from '~/screens/SignUp';
import SignUpCheckEmail from '~/screens/SignUpCheckEmail';
import { AuthStackParamList } from '~/types/navigation';

const AuthStack = createStackNavigator<AuthStackParamList>();

const AuthRoute = function AuthRoute(): ReactElement {
  return (
    <AuthStack.Navigator
      screenOptions={{
        headerShown: false
      }}
      initialRouteName="SignInSelect"
    >
      <AuthStack.Screen name="SignInSelect" component={SignInSelect}/>
      <AuthStack.Screen name="SignUp" component={SignUp} />
      <AuthStack.Screen name="SignUpCheckEmail" component={SignUpCheckEmail} />
    </AuthStack.Navigator>
  );
};

export default AuthRoute;
