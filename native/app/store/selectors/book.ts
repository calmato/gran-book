import { AppState } from '../modules';
import { Book } from '~/store/models';
import { ViewBooks } from '~/types/models/book';

export const bookSelector = (state: AppState): ViewBooks => {
  return Book.filterBooks(state.book);
};

export const recommendBookSelector = (state: AppState) => {
  return state.book.books;
};

export const monthlyBookSelector = (state: AppState) => {
  const readBooks = Book.filterBooks(state.book).read;

  const year = new Date().getFullYear();

  const labels = [
    '1月',
    '2月',
    '3月',
    '4月',
    '5月',
    '6月',
    '7月',
    '8月',
    '9月',
    '10月',
    '11月',
    '12月',
  ];

  const data = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

  const res = readBooks.filter((book) => {
    console.log(year);
    console.log(book.bookshelf?.updatedAt);
    return book.bookshelf?.readOn;
  });

  return { labels, data };
};
