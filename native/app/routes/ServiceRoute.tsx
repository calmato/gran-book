import { MaterialCommunityIcons } from '@expo/vector-icons';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import React, { ReactElement } from 'react';
import { StyleSheet } from 'react-native';
import { colors } from 'react-native-elements';
import Home from '~/screens/Home';
import MyPage from '~/screens/MyPage';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({});

interface Props {}

function TabBarIcon(name, focused, size) {
  let iconName;

  switch (name) {
  case 'ホーム':
    iconName = 'home';
    break;
  case 'マイページ':
    iconName = 'account';
    break;
  default:
    iconName = 'home';
    break;
  }

  const color = focused ? COLOR.MAIN : colors.grey0;
        
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
    >
      <Tab.Screen name="ホーム" component={Home} />
      <Tab.Screen name="マイページ" component={MyPage} />
    </Tab.Navigator>
  );
};

// ServiceRoute.defaultProps={}

export default ServiceRoute;
