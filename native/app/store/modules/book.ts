import { Book } from '~/store/models';

export function createInitialState(): Book.Model {
  return Book.factory();
}

export type State = ReturnType<typeof createInitialState>;

export const SET_BOOKS = 'gran-book/books/SET_BOOKS';

export function setBooks(books: Book.BookValues) {
  return {
    type: SET_BOOKS,
    payload: { books },
  };
}

export type Action =
  | Readonly<ReturnType<typeof setBooks>>

export default function reducer(state: State = createInitialState(), action: Action): Book.Model {
  const { payload } = action;

  switch (action.type) {
    case SET_BOOKS:
      return Book.setBooks(state, payload.books);
    default:
      return state;
  }
}
