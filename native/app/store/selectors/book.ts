import { Book } from '~/store/models';
import { ViewBooks } from '~/types/models/book';
import { AppState } from '../modules';

export const bookSelector = (state: AppState): ViewBooks => {
  return Book.filterBooks(state.book);
};
