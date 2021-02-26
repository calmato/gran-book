import React, { ReactElement } from 'react';
import { createStackNavigator } from '@react-navigation/stack';
import { MyPage } from '~/containers';
import OwnProfile from '~/screens/OwnProfile';
import { SettingStackParamList } from '~/types/navigation';

const SettingStack = createStackNavigator();

const SettingRoute = function SettingRoute() {
  return (
    <SettingStack.Navigator
      initialRouteName='MyPage'
      headerMode='none'
    >
      <SettingStack.Screen name='MyPage' component={MyPage} />
      <SettingStack.Screen name='OwnProfile' component={OwnProfile}/>
    </SettingStack.Navigator>
  );
};

export default SettingRoute;
