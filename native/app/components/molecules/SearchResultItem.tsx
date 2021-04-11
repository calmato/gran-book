import React, { ReactElement } from 'react';
import { StyleSheet, Image } from 'react-native';
import { ListItem } from 'react-native-elements';
import { ISearchResultItem } from '~/types/response/external/rakuten-books';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  bookCoverStyle: {
    width: 80,
    height: 120,
    resizeMode: 'contain',
  },
  bookTitle: {
    color: COLOR.TEXT_DEFAULT,
    marginBottom: 8,
    fontWeight: '800',
    fontSize: 16,
  },
  authors: {
    color: COLOR.TEXT_GRAY,
    fontSize: 12,
  },
});

interface Props {
  book: Partial<ISearchResultItem>;
  onPress: () => void;
}

const SearchResultItem = function SearchResultItem(props: Props): ReactElement {
  const book = props.book;

  return (
    <ListItem bottomDivider onPress={props.onPress}>
      <Image
        source={
          book.largeImageUrl
            ? { uri: book.largeImageUrl }
            : require('assets/logo.png')
        }
        style={styles.bookCoverStyle}
        width={80}
        height={120}
      />
      <ListItem.Content>
        <ListItem.Title style={styles.bookTitle} allowFontScaling>
          {book.title}
        </ListItem.Title>
        <ListItem.Subtitle style={styles.authors}>
          {book.author ? book.author : '著者情報がありません'}
        </ListItem.Subtitle>
      </ListItem.Content>
    </ListItem>
  );
};

export default SearchResultItem;
