import React, { ReactElement, useCallback, useRef } from 'react';
import {
  Dimensions,
  StyleSheet,
  View,
  Text,
  NativeScrollEvent,
  NativeSyntheticEvent,
} from 'react-native';

import { Image } from 'react-native-elements';
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
    width: 0.85 * width,
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
    padding: 8,
  },
  recommendBookCardContent: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    width: '100%',
    padding: 8,
  },
  bookTitle: {
    fontWeight: 'bold',
  },
  imageContainer: {
    marginRight: 16,
    width: 80,
    height: 150,
    resizeMode: 'contain',
  },
  bookDescription: {
    width: 0.85 * 0.55 * width,
  },
});

interface Props {
  books: IBook[];
}

const RecommendBooks = function RecommendBooks(props: Props): ReactElement {
  const recommendBooks = props.books;

  const carouselRef = useRef<FlatList<IBook>>(null);

  const onResponderReleaseHandler = useCallback(
    (event: NativeSyntheticEvent<NativeScrollEvent>) => {
      if (!event.nativeEvent.velocity) {
        return;
      }
      const offset = event.nativeEvent.contentOffset.x;
      let index: number;
      const rate = width * 0.85;
      const velocity = event.nativeEvent.velocity.x;
      if (velocity > 0) {
        index = Math.round(offset / rate);
      } else {
        index = Math.ceil(offset / rate);
      }
      carouselRef.current?.scrollToOffset({
        animated: true,
        offset: event.nativeEvent.velocity.x > 0 ? (index + 1) * rate : (index - 1) * rate,
      });
    },
    [],
  );

  return (
    <View style={styles.root}>
      <View style={styles.contentArea}>
        <FlatList
          initialNumToRender={5}
          ref={carouselRef}
          data={recommendBooks}
          horizontal
          onScrollEndDrag={onResponderReleaseHandler}
          showsHorizontalScrollIndicator={false}
          keyExtractor={(_, index) => index.toString()}
          renderItem={({ item }) => (
            <View style={[styles.carousel]}>
              <View style={styles.recommendBookCard}>
                <Text style={styles.bookTitle}>{item.title}</Text>
                <View style={styles.recommendBookCardContent}>
                  <Image source={{ uri: item.thumbnailUrl }} style={[styles.imageContainer]} />
                  <Text style={styles.bookDescription} numberOfLines={10}>
                    {item.description}
                  </Text>
                </View>
              </View>
            </View>
          )}
        />
      </View>
    </View>
  );
};

export default RecommendBooks;
