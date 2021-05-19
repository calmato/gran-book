import React, { useMemo } from 'react';
import Home from '~/screens/Home';
import { useSelector } from 'react-redux';
import { useReduxDispatch } from '~/store/modules';
import { getAllBookAsync } from '~/store/usecases';
import { bookSelector } from '~/store/selectors/book';
import { ViewBooks } from '~/types/models/book';

export default function ConnectedHome() {
  const dispatch = useReduxDispatch();
  const books: ViewBooks = useSelector(bookSelector);

  const actions = useMemo(
    () => ({
      getAllBook() {
        return dispatch(getAllBookAsync());
      }
    }),[dispatch]);

  return <Home actions={actions} books={books} />;
}
