import { createStackNavigator } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import SignInSelect from '~/screens/SignInSelect';
import SignUpCheckEmail from '~/screens/SignUpCheckEmail';
import SingIn from '~/screens/SignIn';
import { SignUp } from '~/containers';
import { AuthStackParamList } from '~/types/navigation';
import PasswordReset from '~/screens/PasswordReset';

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
      <AuthStack.Screen name="SignIn" component={SingIn} />
      <AuthStack.Screen name="PasswordReset" component={PasswordReset} />
    </AuthStack.Navigator>
  );
};

export default AuthRoute;
