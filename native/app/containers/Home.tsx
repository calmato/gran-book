import React, { useMemo } from 'react';
import Home from '~/screens/Home';
import { Book } from '~/store/models';
import { useSelector } from 'react-redux';
import { useReduxDispatch } from '~/store/modules';
import { getAllBookAsync } from '~/store/usecases';
import { bookSelector } from '~/store/selectors/book';

export default function ConnectedHome() {
  const dispatch = useReduxDispatch();
  const books: Book.Model = useSelector(bookSelector);

  const actions = useMemo(
    () => ({
      getAllBook() {
        return dispatch(getAllBookAsync());
      }
    }),[dispatch]);

  return <Home actions={actions} books={books.books} />;
}
