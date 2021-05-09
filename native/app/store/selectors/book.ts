import { Book } from '~/store/models';
import { AppState } from '../modules';

export const bookSelector = (state: AppState): Book.Model => {
  return state.book;
};
