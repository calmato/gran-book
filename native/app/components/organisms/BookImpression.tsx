import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { Text, View, StyleSheet } from 'react-native';
import { ListItem, Avatar, Divider, Image, Badge } from 'react-native-elements';
import { IBook, IImpressionResponse } from '~/types/response';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  badgeStyle: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
    alignSelf: 'flex-start',
    marginStart: 10,
    marginTop: 10,
    height: 30,
    width: 150,
    borderRadius: 75,
  },
  bookInfoStyle: {
    flexDirection: 'row',
    marginStart: 10,
    marginTop: 10,
    backgroundColor: COLOR.TEXT_WHITE,
  },
});

interface Props {
  impressionResponse: IImpressionResponse;
  book: IBook;
}

const BookImpression = function BookImpression(props: Props): ReactElement {
  const book = props.book;
  const reviews = props.impressionResponse.reviews;
  const total = props.impressionResponse.total;

  return (
    <View>
      <Badge
        value={<Text style={{ fontSize: 16 }}>{`${total}件`}</Text>}
        badgeStyle={styles.badgeStyle}
      />
      <View style={styles.bookInfoStyle}>
        <Image source={{ uri: book.thumbnailUrl }} style={{ width: 50, height: 70 }} />
        <View style={{ justifyContent: 'space-around', marginStart: 20 }}>
          <Text style={{ fontSize: 16 }}>{book.title}</Text>
          <Text style={{ fontSize: 16, color: COLOR.GREY }}>{book.author}</Text>
        </View>
      </View>
      <View style={{ marginTop: 10 }}>
        {reviews.map((review) => (
          <View style={{ backgroundColor: COLOR.TEXT_WHITE }} key={review.id}>
            <ListItem key={review.id}>
              <Avatar source={{ uri: review.user.thumbnailUrl }} rounded />
              <ListItem.Content>
                <ListItem.Title>{review.user.username + 'が感想を投稿しました'}</ListItem.Title>
                <ListItem.Subtitle>{review.createdAt}</ListItem.Subtitle>
              </ListItem.Content>
            </ListItem>
            <Text style={{ fontSize: 16, marginStart: 15, marginEnd: 15 }}>
              {review.impression}
            </Text>
            <View style={{ marginStart: 15, flexDirection: 'row', alignItems: 'center' }}>
              <Ionicons name="heart-outline" size={36} color={COLOR.GREY} />
              <Text style={{ marginStart: 10 }}></Text>
            </View>
            <Divider style={{ height: 2 }} />
          </View>
        ))}
      </View>
    </View>
  );
};

export default BookImpression;
