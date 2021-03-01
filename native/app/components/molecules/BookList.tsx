import React, { ReactElement } from 'react';
import { StyleSheet, View, Dimensions } from 'react-native';
import { booksSampleData } from '~~/assets/sample/book';
import Book from '../atoms/Book';

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
  }
});

interface Props {
  books?: Array<any>
}

const bookList = booksSampleData;

const BookList = function BookList(props: Props): ReactElement {
  const width = Dimensions.get('window').width;
  const rowItem = 3;

  return (
    <View style={styles.containerStyle}>
      {
        bookList.map((book: any, i: number) => (
          <Book
            style={styles.childStyle}
            title={book.title}
            image={book.image}
            author={book.author}
            height={250}
            width={width/rowItem - 5}
            onPress={()=>console.log(book.title)}
            key={i}
          />
        ))
      }
    </View>
  );
};

export default BookList;
