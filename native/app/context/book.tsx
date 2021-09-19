import React, { createContext, useCallback, useContext, useReducer } from 'react';
import { AuthContext } from './auth';
import { BookValues, initialState, Model } from '~/store/models/book';
import { getAllBookByUserId } from '~/store/usecases/v2/book';

interface BookContextProps {
  bookState: Model;
  dispatch: React.Dispatch<BookStateAction>;
}

const BookContext = createContext<BookContextProps>({
  bookState: initialState,
  dispatch: () => {
    return;
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

  return <BookContext.Provider value={{ bookState, dispatch }}>{children}</BookContext.Provider>;
};

export { BookContext, BookProvider };
