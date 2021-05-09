import React, { ReactElement } from 'react';
import { createStackNavigator } from '@react-navigation/stack';
import Home from '~/containers/Home';
import SearchResult from '~/screens/SearchResult';
import { HomeTabStackPramList } from '~/types/navigation';
import BookShow from '~/screens/BookShow';

const HomeTabStack = createStackNavigator<HomeTabStackPramList>();

const HomeTabRoute = function HomeTabRoute(): ReactElement {
  return (
    <HomeTabStack.Navigator
      screenOptions={{
        headerShown: false,
      }}
      initialRouteName="Home">
      <HomeTabStack.Screen name="Home" component={Home} />
      <HomeTabStack.Screen name="SearchResult" component={SearchResult} />
      <HomeTabStack.Screen name="SearchResultBookShow" component={BookShow} />
    </HomeTabStack.Navigator>
  );
};

export default HomeTabRoute;
