import { createStackNavigator } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { Bookshelf } from '~/containers';
import BookReadRegister from '~/containers/BookReadRegister';
import BookShow from '~/containers/BookShow';
import SearchResult from '~/screens/SearchResult';
import { BookshelfTabStackPramList } from '~/types/navigation';

const BookshelfTabStack = createStackNavigator<BookshelfTabStackPramList>();

const BookshelfTabRoute = function bookshelfTabRoute(): ReactElement {
  return (
    <BookshelfTabStack.Navigator
      screenOptions={{
        headerShown: false,
      }}
      initialRouteName="Bookshelf">
      <BookshelfTabStack.Screen name="Bookshelf" component={Bookshelf} />
      <BookshelfTabStack.Screen name="SearchResult" component={SearchResult} />
      <BookshelfTabStack.Screen name="SearchResultBookShow" component={BookShow} />
      <BookshelfTabStack.Screen name="BookShow" component={BookShow} />
      <BookshelfTabStack.Screen name="BookReadRegister" component={BookReadRegister} />
    </BookshelfTabStack.Navigator>
  );
};

export default BookshelfTabRoute;
