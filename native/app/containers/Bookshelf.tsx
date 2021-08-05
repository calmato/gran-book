import { StackNavigationProp } from '@react-navigation/stack';
import React, { useMemo } from 'react';
import { useSelector } from 'react-redux';
import Bookshelf from '~/screens/Bookshelf';
import { useReduxDispatch } from '~/store/modules';
import { bookSelector } from '~/store/selectors/book';
import { getAllBookAsync } from '~/store/usecases';
import { ViewBooks } from '~/types/models/book';
import { BookshelfTabStackParamList } from '~/types/navigation';

interface Props {
  navigation?: StackNavigationProp<BookshelfTabStackParamList, 'Bookshelf'>;
}

export default function ConnectedBookshelf(props: Props) {
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

  return <Bookshelf actions={actions} books={books} navigation={props.navigation} />;
}
