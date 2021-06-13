import { StackNavigationProp } from '@react-navigation/stack';
import React, { useMemo } from 'react';
import { useSelector } from 'react-redux';
import Home from '~/screens/Home';
import { useReduxDispatch } from '~/store/modules';
import { bookSelector } from '~/store/selectors/book';
import { getAllBookAsync } from '~/store/usecases';
import { ViewBooks } from '~/types/models/book';
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
      },
    }),
    [dispatch],
  );

  return <Home actions={actions} books={books} navigation={props.navigation} />;
}
