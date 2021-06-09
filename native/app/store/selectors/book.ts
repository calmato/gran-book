import { AppState } from '../modules';
import { Book } from '~/store/models';
import { ViewBooks } from '~/types/models/book';

export const bookSelector = (state: AppState): ViewBooks => {
  return Book.filterBooks(state.book);
};
