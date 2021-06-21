import { createMaterialTopTabNavigator } from '@react-navigation/material-top-tabs';
import React, { ReactElement } from 'react';
import { View } from 'react-native';
import BookImpression from '~/components/organisms/BookImpression';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import BookShow from '~/screens/BookShow';
import { COLOR } from '~~/constants/theme';

const BookShowImpressionParent = function BookShowImpressionParent(): ReactElement {
  const Tab = createMaterialTopTabNavigator();

  return (
    <View
      style={{
        flex: 1,
      }}>
      <HeaderWithBackButton title="本のタイトル" onPress={() => undefined} />
      <Tab.Navigator tabBarOptions={{ indicatorStyle: { backgroundColor: COLOR.PRIMARY } }}>
        <Tab.Screen name="情報" component={BookShow} />
        <Tab.Screen name="感想" component={BookImpression} />
      </Tab.Navigator>
    </View>
  );
};

export default BookShowImpressionParent;
