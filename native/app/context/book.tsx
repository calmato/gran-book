import React, { createContext, useCallback, useContext, useMemo, useReducer } from 'react';
import { AuthContext } from './auth';
import { BookValues, initialState, Model } from '~/store/models/book';
import { getAllBookByUserId } from '~/store/usecases/v2/book';
import { ViewBooks } from '~/types/models/book';

interface BookContextProps {
  bookState: Model;
  viewBooks: ViewBooks;
  fetchBooks: () => Promise<void>;
}

const BookContext = createContext<BookContextProps>({
  bookState: initialState,
  viewBooks: {
    reading: [],
    read: [],
    stack: [],
    release: [],
    want: [],
  },
  fetchBooks: () => {
    return Promise.resolve();
  },
});

type ActionType = 'SET_BOOKS';
type Payload = BookValues;

interface BookStateAction {
  type: ActionType;
  payload: Payload;
}

const reducer: React.Reducer<Model, BookStateAction> = function reducer(
  state: Model,
  action: BookStateAction,
): Model {
  switch (action.type) {
    case 'SET_BOOKS':
      return {
        ...state,
        ...action.payload,
      };
    default:
      return state;
  }
};

interface Props {
  children?: React.ReactNode;
}

const BookProvider = function BookProvider({ children }: Props) {
  const { authState } = useContext(AuthContext);

  const [bookState, dispatch] = useReducer(reducer, initialState);

  const fetchBooks = useCallback(async () => {
    const books = await getAllBookByUserId(authState.id, authState.token);
    dispatch({
      type: 'SET_BOOKS',
      payload: { books },
    });
  }, [authState.id, authState.token]);

  const viewBooks: ViewBooks = useMemo(() => {
    const value: ViewBooks = {
      reading: [],
      read: [],
      stack: [],
      release: [],
      want: [],
    };

    bookState.books.booksList.forEach((book) => {
      switch (book.bookshelf?.status) {
        case 'reading':
          value.reading.push(book);
          break;
        case 'read':
          value.read.push(book);
          break;
        case 'stack':
          value.stack.push(book);
          break;
        case 'release':
          value.release.push(book);
          break;
        case 'want':
          value.want.push(book);
          break;
        default:
          break;
      }
    });

    return value;
  }, [bookState]);

  return (
    <BookContext.Provider value={{ bookState, viewBooks, fetchBooks }}>
      {children}
    </BookContext.Provider>
  );
};

export { BookContext, BookProvider };
