import React, { ReactElement } from 'react';
import { Dimensions, StyleSheet, View, Text } from 'react-native';

import { Button, Image } from 'react-native-elements';
import { FlatList } from 'react-native-gesture-handler';
import { IBook } from '~/types/response';
import { COLOR } from '~~/constants/theme';

const { width } = Dimensions.get('screen');

const styles = StyleSheet.create({
  root: {
    backgroundColor: '#FFD996',
    opacity: 0.8,
  },
  contentArea: {
    display: 'flex',
    alignItems: 'center',
    paddingVertical: 16,
  },
  carousel: {
    width: width,
    height: 200,
    justifyContent: 'center',
    alignItems: 'center',
    paddingHorizontal: 12,
  },
  recommendBookCard: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
    marginHorizontal: 16,
    height: 200,
    width: '100%',
    borderRadius: 16,
  },
  imageContainer: {
    marginVertical: 24,
    width: '100%',
    height: '100%',
    resizeMode: 'contain',
  },
});

interface Props {
  books: IBook[];
}

const RecommendBooks = function RecommendBooks(props: Props): ReactElement {
  const recommendBooks = props.books;

  return (
    <View style={styles.root}>
      <Text>あなたへのおすすめ</Text>
      <View style={styles.contentArea}>
        <FlatList
          data={recommendBooks}
          horizontal
          pagingEnabled
          keyExtractor={(_, index) => index.toString()}
          renderItem={({ item }) => (
            <View style={[styles.carousel]}>
              <View style={styles.recommendBookCard}>
                <Text>{item.title}</Text>
                <Image source={{ uri: item.thumbnailUrl }} style={[styles.imageContainer]} />
              </View>
            </View>
          )}
        />
        <Button title="おすすめ一覧を見る" />
      </View>
    </View>
  );
};

export default RecommendBooks;
