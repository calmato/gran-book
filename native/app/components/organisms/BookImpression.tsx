import { Ionicons, MaterialIcons } from '@expo/vector-icons';
import dayjs from 'dayjs';
import React, { ReactElement } from 'react';
import { Text, View, StyleSheet } from 'react-native';
import { ListItem, Avatar, Divider, Image, Badge } from 'react-native-elements';
import { IBook, IImpressionResponse } from '~/types/response';
import { COLOR, FONT_SIZE } from '~~/constants/theme';

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
    paddingHorizontal: 10,
    alignSelf: 'stretch',
    backgroundColor: COLOR.TEXT_WHITE,
    marginVertical: 8,
  },
  titleStyle: {
    fontSize: FONT_SIZE.ITEM_TITLE,
    fontWeight: '500',
    color: COLOR.GREY,
    marginBottom: 8,
  },
  authorStyle: {
    fontSize: FONT_SIZE.ITEM_SUBTITLE,
    color: COLOR.GREY,
  },
  listTitleStyle: {
    color: COLOR.GREY,
    fontSize: FONT_SIZE.ITEM_TITLE,
    marginBottom: 4,
  },
  listSubTitleStyle: {
    color: COLOR.GREY,
    fontSize: FONT_SIZE.ITEM_SUBTITLE,
  },
  reviewStyle: {
    paddingHorizontal: 16,
    marginBottom: 8,
    color: COLOR.GREY,
    fontSize: FONT_SIZE.TEXT,
  },
  reactionContainerStyle: {
    marginVertical: 4,
    paddingHorizontal: 16,
    flexDirection: 'row',
    alignItems: 'center',
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
        <View
          style={{
            flexDirection: 'column',
            width: '90%',
            justifyContent: 'center',
            paddingHorizontal: 8,
          }}>
          <Text style={styles.titleStyle}>{book.title}</Text>
          <Text style={styles.authorStyle}>{book.author}</Text>
        </View>
      </View>

      <View>
        {reviews.map((review) => (
          <View style={{ backgroundColor: COLOR.TEXT_WHITE, marginBottom: 4 }} key={review.id}>
            <ListItem key={review.id}>
              {review.user.thumbnailUrl !== '' ? (
                <Avatar source={{ uri: review.user.thumbnailUrl }} rounded />
              ) : (
                <Avatar rounded>
                  <MaterialIcons name="person-outline" size={36} color={COLOR.GREY} />
                </Avatar>
              )}
              <ListItem.Content>
                <ListItem.Title style={styles.listTitleStyle}>
                  {review.user.username + ' が感想を投稿しました'}
                </ListItem.Title>
                <ListItem.Subtitle style={styles.listSubTitleStyle}>
                  {dayjs(review.createdAt).format('YYYY/MM/DD')}
                </ListItem.Subtitle>
              </ListItem.Content>
            </ListItem>
            <Text style={styles.reviewStyle}>{review.impression}</Text>
            <Divider style={{ height: 0.5 }} />
            <View style={styles.reactionContainerStyle}>
              <Ionicons name="heart-outline" size={20} color={COLOR.GREY} />
              <Text style={{ marginStart: 10 }}></Text>
            </View>
          </View>
        ))}
      </View>
    </View>
  );
};

export default BookImpression;
