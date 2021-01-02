import { createStackNavigator } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { StyleSheet } from 'react-native';
import SignInSelect from '~/screens/SignInSelect';
import SignUp from '~/screens/SignUp';

const styles = StyleSheet.create({});

// interface Props {}

const AuthStack = createStackNavigator();

const AuthRoute = function AuthRoute(): ReactElement {
  return (
    <AuthStack.Navigator
      screenOptions={{
        headerShown: false
      }}
    >
      <AuthStack.Screen name="SignInSelect" component={SignInSelect}/>
      <AuthStack.Screen name="SignUp" component={SignUp} />
    </AuthStack.Navigator>
  );
};

// AuthRoute.defaultProps={}

export default AuthRoute;
