import React, { ReactElement } from 'react';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { StyleSheet } from 'react-native';
import Home from '~/screens/Home';
import MyPage from '~/screens/MyPage';
import { COLOR } from '~~/constants/theme';
import Store from '~/screens/Store';
import Bookshelf from '~/screens/Bookshelf';
import Sale from '~/screens/Sale';

const styles = StyleSheet.create({
  labelStyle: {
    fontWeight: 'bold',
  }
});

function TabBarIcon(name: string, focused: boolean, size: number) {
  let iconName: any;

  switch (name) {
  case 'ホーム':
    iconName = 'home';
    break;
  case 'マイページ':
    iconName = 'account';
    break;
  case '本棚':
    iconName= 'bookshelf';
    break;
  case '本を出品':
    iconName = 'plus-circle-outline';
    break;
  case '本を買う':
    iconName = 'cart';
    break;
  default:
    iconName = 'home';
    break;
  }

  const color = focused ? COLOR.PRIMARY : COLOR.TEXT_GRAY;

  return <MaterialCommunityIcons name={iconName} color={color} size={size} />;
}

const ServiceRoute = function ServiceRoute(): ReactElement {
  const Tab = createBottomTabNavigator();

  return (
    <Tab.Navigator
      screenOptions={({ route }) => ({
        tabBarIcon: ({ focused, size }) => {
          return TabBarIcon(route.name, focused, size);
        },
      })}
      tabBarOptions={{
        activeTintColor: COLOR.TEXT_TITLE,
        labelStyle: styles.labelStyle
      }}
    >
      <Tab.Screen name="ホーム" component={Home} />
      <Tab.Screen name="本棚" component={Bookshelf} />
      <Tab.Screen name="本を出品" component={Sale} />
      <Tab.Screen name="本を買う" component={Store} />
      <Tab.Screen name="マイページ" component={MyPage} />
    </Tab.Navigator>
  );
};

export default ServiceRoute;