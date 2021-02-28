import React, { ReactElement } from 'react';
import { createStackNavigator } from '@react-navigation/stack';
import { MyPage } from '~/containers';
import OwnProfile from '~/screens/OwnProfile';
import { UserInfoStackParamList } from '~/types/navigation';

const UserInfoStack = createStackNavigator();

const UserInfoRoute = function SettingRoute() {
  return (
    <UserInfoStack.Navigator
      initialRouteName='MyPage'
      headerMode='none'
    >
      <UserInfoStack.Screen name='MyPage' component={MyPage} />
      <UserInfoStack.Screen name='OwnProfile' component={OwnProfile}/>
    </UserInfoStack.Navigator>
  );
};

export default UserInfoRoute;
