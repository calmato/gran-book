import { createStackNavigator } from '@react-navigation/stack';
import React from 'react';
import {
  MyPage,
  OwnProfile,
  EmailEdit,
  ContactEdit,
  ProfileEdit,
  AccountSetting,
  OwnReviews,
} from '~/containers';

import { UserInfoStackParamList } from '~/types/navigation';

const UserInfoStack = createStackNavigator<UserInfoStackParamList>();

const UserInfoRoute = function SettingRoute() {
  return (
    <UserInfoStack.Navigator screenOptions={{ headerShown: false }} initialRouteName="MyPage">
      <UserInfoStack.Screen name="MyPage" component={MyPage} />
      <UserInfoStack.Screen name="OwnProfile" component={OwnProfile} />
      <UserInfoStack.Screen name="OwnReviews" component={OwnReviews} />
      <UserInfoStack.Screen name="AccountSetting" component={AccountSetting} />
      <UserInfoStack.Screen name="ContactEdit" component={ContactEdit} />
      <UserInfoStack.Screen name="EmailEdit" component={EmailEdit} />
      <UserInfoStack.Screen name="ProfileEdit" component={ProfileEdit} />
    </UserInfoStack.Navigator>
  );
};

export default UserInfoRoute;
