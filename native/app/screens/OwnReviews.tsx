import { MaterialIcons } from '@expo/vector-icons';
import React, { ReactElement, useEffect, useState } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { Header, ListItem, Avatar } from 'react-native-elements';
import { Badge } from 'react-native-elements/dist/badge/Badge';
import { ScrollView } from 'react-native-gesture-handler';
import HeaderText from '~/components/atoms/HeaderText';
import { Auth } from '~/store/models';
import { getOwnReviews } from '~/store/usecases';
import { IReviewResponse } from '~/types/response';
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
  reviewListStyle: {
    marginTop: 8,
    height: '100%',
  },
  reviewStyle: {
    flexDirection: 'column',
    alignSelf: 'stretch',
  },
  titleStyle: {
    fontSize: FONT_SIZE.LISTITEM_TITLE,
    fontWeight: 'bold',
    paddingBottom: 4,
    marginBottom: 4,
  },
  subtitleStyle: {
    fontSize: FONT_SIZE.LISTITEM_SUBTITLE,
    paddingVertical: 4,
  },
});

interface Props {
  auth: Auth.Model;
}

const nullData: IReviewResponse = {
  total: 0,
  offset: 1,
  limit: 1,
  reviews: [
    {
      createdAt: '',
      book: {
        id: 0,
        title: '',
        thumbnailUrl: '',
      },
      impression: '',
      id: 0,
      updatedAt: '',
    },
  ],
};

export const OwnReviews = function OwnReviews(props: Props): ReactElement {
  const [reviewList, setReviewList] = useState<IReviewResponse>(nullData);

  useEffect(() => {
    const f = async () => {
      const res = await getOwnReviews(props.auth.id);
      setReviewList(res.data);
    };
    f();
  }, []);

  const review = reviewList.reviews;
  return (
    <View>
      <Header centerComponent={<HeaderText title="自分の感想" />} />
      <Badge
        value={<Text style={{ fontSize: 16 }}>{`${reviewList.total}件`}</Text>}
        badgeStyle={styles.badgeStyle}
      />
      <ScrollView style={styles.reviewListStyle}>
        {review.map((review) => (
          <View key={review.id}>
            <ListItem key={review.id} bottomDivider={true} pad={20}>
              {review.book.thumbnailUrl !== '' ? (
                <Avatar source={{ uri: review.book.thumbnailUrl }} size="large" />
              ) : (
                <Avatar>
                  <MaterialIcons name="person-outline" size={36} color={COLOR.GREY} />
                </Avatar>
              )}
              <ListItem.Content style={styles.reviewStyle}>
                <ListItem.Title style={styles.titleStyle}>{review.book.title}</ListItem.Title>
                <ListItem.Subtitle style={styles.subtitleStyle}>
                  {review.impression}
                </ListItem.Subtitle>
              </ListItem.Content>
              <ListItem.Chevron name="create" size={32} />
            </ListItem>
          </View>
        ))}
      </ScrollView>
    </View>
  );
};
