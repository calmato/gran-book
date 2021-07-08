import { createStackNavigator } from '@react-navigation/stack';
import React from 'react';
import { MyPage , OwnProfile , EmailEdit , ContactEdit , ProfileEdit , AccountSetting } from '~/containers';

import { UserInfoStackParamList } from '~/types/navigation';
<<<<<<< HEAD
import { EmailEdit } from '~/containers';
import { ContactEdit } from '~/containers';
import { ProfileEdit } from '~/containers';
import { AccountSetting } from '~/containers';
import { Support } from '~/containers';
=======




>>>>>>> 3c9a5a31619a7e0e50130070a2bb56746e750f03

const UserInfoStack = createStackNavigator<UserInfoStackParamList>();

const UserInfoRoute = function SettingRoute() {
  return (
    <UserInfoStack.Navigator initialRouteName="MyPage" headerMode="none">
      <UserInfoStack.Screen name="MyPage" component={MyPage} />
      <UserInfoStack.Screen name="OwnProfile" component={OwnProfile} />
      <UserInfoStack.Screen name="AccountSetting" component={AccountSetting} />
      {/* <UserInfoStack.Screen name='AccountEdit' component={AccountEdit}/> */}
      <UserInfoStack.Screen name="ContactEdit" component={ContactEdit} />
      <UserInfoStack.Screen name="EmailEdit" component={EmailEdit} />
      <UserInfoStack.Screen name="ProfileEdit" component={ProfileEdit} />
      <UserInfoStack.Screen name="Support" component = {Support} />
      {/* <UserInfoStack.Screen name='PasswordEmailEdit' componet={PasswordEmailEdit}/> */}
    </UserInfoStack.Navigator>
  );
};

export default UserInfoRoute;
