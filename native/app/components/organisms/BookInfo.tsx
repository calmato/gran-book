import React, { ReactElement } from 'react';
import { StyleSheet, View, Text, ScrollView } from 'react-native';
import { Image, Button } from 'react-native-elements';
import ButtonGroupBookFooter from './ButtonGroupBookFooter';
import FlexBoxBookCategory from './FlexBoxBookCategory';
import { fullWidth2halfWidth } from '~/lib/util';
import { IBook } from '~/types/response';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    alignItems: 'center',
  },
  imageContainer: {
    marginVertical: 24,
    width: 200,
    height: 280,
    resizeMode: 'contain',
  },
  titleContainer: {
    paddingTop: 10,
    paddingHorizontal: 16,
    fontSize: 16,
    alignSelf: 'stretch',
    color: COLOR.GREY,
    fontWeight: '500',
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
  authorContainer: {
    paddingTop: 10,
    paddingHorizontal: 16,
    paddingBottom: 10,
    fontSize: 12,
    alignSelf: 'stretch',
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
  detailContainer: {
    paddingHorizontal: 16,
    paddingBottom: 10,
    fontSize: 14,
    lineHeight: 16,
    color: COLOR.GREY,
    alignSelf: 'stretch',
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
});

interface Props {
  book: IBook;
  isRegister: boolean;
  handleBookStatusButton: (status: string) => void;
  handleOpenRakutenPageButton: (url: string) => Promise<void>;
  handleAddButton: () => Promise<void>;
}

const BookInfo = function BookInfo(props: Props): ReactElement {
  const book = props.book;
  const isRegister = props.isRegister;

  return (
    <ScrollView
      contentContainerStyle={styles.container}
      style={{ marginBottom: 'auto', height: '100%' }}>
      <View
        style={{
          alignSelf: 'stretch',
          justifyContent: 'space-around',
          alignItems: 'center',
        }}>
        <Image
          source={book.thumbnailUrl ? { uri: book.thumbnailUrl } : require('assets/logo.png')}
          style={styles.imageContainer}
          transition={true}
        />
      </View>
      <Text style={styles.titleContainer}>{book.title}</Text>
      <Text style={styles.authorContainer}>
        {book.author ? book.author : '著者情報がありません'}
      </Text>
      {book.description !== '' ? (
        <Text style={styles.detailContainer}>{fullWidth2halfWidth(book.description)}</Text>
      ) : null}
      <FlexBoxBookCategory category={book.size ? book.size : 'カテゴリ情報がありません'} />

      {isRegister ? (
        <ButtonGroupBookFooter
          status={book.bookshelf ? book.bookshelf.status : ''}
          onPress={props.handleBookStatusButton}
        />
      ) : (
        <Button title="本を登録する" onPress={props.handleAddButton} />
      )}

      <Button
        title="楽天で見る"
        containerStyle={{ marginTop: 10, marginBottom: 10 }}
        buttonStyle={{ backgroundColor: COLOR.PRIMARY_DARK }}
        onPress={() => props.handleOpenRakutenPageButton(book.rakutenUrl)}
      />
    </ScrollView>
  );
};

export default BookInfo;
