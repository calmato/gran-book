import React, { useMemo } from 'react';
import Home from '~/screens/Home';
import { useSelector } from 'react-redux';
import { useReduxDispatch } from '~/store/modules';
import { getAllBookAsync } from '~/store/usecases';
import { bookSelector } from '~/store/selectors/book';
import { ViewBooks } from '~/types/models/book';
import { StackNavigationProp } from '@react-navigation/stack';
import { HomeTabStackPramList } from '~/types/navigation';

interface Props {
  navigation?: StackNavigationProp<HomeTabStackPramList, 'Home'>;
}

export default function ConnectedHome(props: Props) {
  const dispatch = useReduxDispatch();
  const books: ViewBooks = useSelector(bookSelector);

  const actions = useMemo(
    () => ({
      getAllBook() {
        return dispatch(getAllBookAsync());
      }
    }),[dispatch]);

  return <Home actions={actions} books={books} navigation={props.navigation} />;
}
