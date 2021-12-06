import { MaterialCommunityIcons } from '@expo/vector-icons';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import React, { ReactElement } from 'react';
import { StyleSheet } from 'react-native';
import { BookProvider } from '~/context/book';
import { UserProvider } from '~/context/user';
import BookshelfTabRoute from '~/routes/BookshelfTabRoute';
import UserInfoRoute from '~/routes/UserInfoRoute';
import Home from '~/screens/Home';
import Sale from '~/screens/Sale';
import Store from '~/screens/Store';
import { COLOR, FONT_SIZE } from '~~/constants/theme';

const styles = StyleSheet.create({
  labelStyle: {
    fontWeight: 'bold',
    fontSize: FONT_SIZE.BUTTON_WITH_ICON,
  },
});

function TabBarIcon(name: string, focused: boolean, size: number) {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  let iconName: any;

  switch (name) {
    case 'ホーム':
      iconName = 'home';
      break;
    case 'マイページ':
      iconName = 'account';
      break;
    case '本棚':
      iconName = 'bookshelf';
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
    <UserProvider>
      <BookProvider>
        <Tab.Navigator
          screenOptions={({ route }) => ({
            tabBarIcon: ({ focused, size }) => {
              return TabBarIcon(route.name, focused, size);
            },
          })}
          tabBarOptions={{
            activeTintColor: COLOR.TEXT_TITLE,
            labelStyle: styles.labelStyle,
          }}>
          <Tab.Screen name="ホーム" component={Home} />
          <Tab.Screen name="本棚" component={BookshelfTabRoute} />
          <Tab.Screen name="本を出品" component={Sale} />
          <Tab.Screen name="本を買う" component={Store} />
          <Tab.Screen name="マイページ" component={UserInfoRoute} />
        </Tab.Navigator>
      </BookProvider>
    </UserProvider>
  );
};

export default ServiceRoute;
