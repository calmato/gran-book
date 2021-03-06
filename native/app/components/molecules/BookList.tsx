﻿import React, { ReactElement } from 'react';
import { StyleSheet, View, Dimensions } from 'react-native';
import Book from '../atoms/Book';
import { IBook } from '~/types/response';

const styles = StyleSheet.create({
  containerStyle: {
    flex: 1,
    flexDirection: 'row',
    flexWrap: 'wrap',
    alignSelf: 'flex-start',
  },
  childStyle: {
    marginTop: 4,
    marginLeft: 4,
  },
});

interface Props {
  books: IBook[];
  handleClick: (book: IBook) => void;
}

const BookList = function BookList(props: Props): ReactElement {
  const width = Dimensions.get('window').width;
  const bookList = props.books;
  const rowItem = 3;

  return (
    <View style={styles.containerStyle}>
      {bookList.map((book: IBook, i: number) => (
        <Book
          style={styles.childStyle}
          title={book.title}
          image={book.thumbnailUrl}
          author={book.author}
          height={250}
          width={width / rowItem - 5}
          onPress={() => props.handleClick(book)}
          key={i}
        />
      ))}
    </View>
  );
};

export default BookList;
