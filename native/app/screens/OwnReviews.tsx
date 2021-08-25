import { MaterialIcons } from '@expo/vector-icons';
import { useNavigation } from '@react-navigation/native';
import React, { ReactElement, useEffect, useState } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { ListItem, Avatar } from 'react-native-elements';
import { Badge } from 'react-native-elements/dist/badge/Badge';
import { ScrollView } from 'react-native-gesture-handler';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { Auth } from '~/store/models';
import { getOwnReviews } from '~/store/usecases';
import { IReviewListResponse } from '~/types/response';
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
    marginBottom: 'auto',
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

const nullData: IReviewListResponse = {
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

const BookAvatar = (isAvatar: boolean, AvatarUri: string) => {
  if (isAvatar) {
    <Avatar source={{ uri: AvatarUri }} size="large" />;
  } else {
    <Avatar>
      <MaterialIcons name="person-outline" size={36} color={COLOR.GREY} />
    </Avatar>;
  }
};
export const OwnReviews = function OwnReviews(props: Props): ReactElement {
  const [reviewList, setReviewList] = useState<IReviewListResponse>(nullData);
  const navigation = useNavigation();

  useEffect(() => {
    const f = async () => {
      const res = await getOwnReviews(props.auth.id);
      setReviewList(res.data);
    };
    f();
  }, []);

  const reviews = reviewList.reviews;
  return (
    <View>
      <HeaderWithBackButton title="自分の感想" onPress={() => navigation.goBack()} />
      <Badge
        value={<Text style={{ fontSize: 16 }}>{`${reviewList.total}件`}</Text>}
        badgeStyle={styles.badgeStyle}
      />
      <ScrollView style={styles.reviewListStyle}>
        {reviews.map((review) => (
          <View key={review.id}>
            <ListItem key={review.id} bottomDivider={true} pad={20}>
              {BookAvatar(review.book.thumbnailUrl !== '', review.book.thumbnailUrl)}
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
