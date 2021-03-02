import React from 'react';
import { createStackNavigator } from '@react-navigation/stack';
import { MyPage } from '~/containers';
import { OwnProfile } from '~/containers';
import { UserInfoStackParamList } from '~/types/navigation';
import AccountSetting from '~/screens/AccoutSetting';
import ProfileEdit from '~/screens/ProfileEdit';

const UserInfoStack = createStackNavigator<UserInfoStackParamList>();

const UserInfoRoute = function SettingRoute() {
  return (
    <UserInfoStack.Navigator
      initialRouteName='MyPage'
      headerMode='none'
    >
      <UserInfoStack.Screen name='MyPage' component={MyPage} />
      <UserInfoStack.Screen name='OwnProfile' component={OwnProfile}/>
      <UserInfoStack.Screen name='AccountSetting' component={AccountSetting}/>
      {/* <UserInfoStack.Screen name='AccountEdit' component={AccountEdit}/> */}
      <UserInfoStack.Screen name='ProfileEdit' component={ProfileEdit}/>
      {/* <UserInfoStack.Screen name='ContactEdit' component={ContactEdit}/> */}
      {/* <UserInfoStack.Screen name='EmailEdit' component={EmailEdit}/> */}
      {/* <UserInfoStack.Screen name='PasswordEmailEdit' componet={PasswordEmailEdit}/> */}
    </UserInfoStack.Navigator>
  );
};

export default UserInfoRoute;
