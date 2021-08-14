import { MaterialIcons } from '@expo/vector-icons';
import React, { ReactElement, useEffect, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { Header, ListItem, Avatar } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import { Auth } from '~/store/models';
import { getOwnReviews } from '~/store/usecases';
import { IReviewResponse } from '~/types/response';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({});

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
  const [reviews, setReviews] = useState<IReviewResponse>(nullData);

  useEffect(() => {
    const f = async () => {
      const res = await getOwnReviews(props.auth.id);
      setReviews(res.data);
    };
    f();
  }, []);

  const review = reviews.reviews;
  return (
    <View>
      <Header centerComponent={<HeaderText title="自分の感想" />} />
      {review.map((review) => (
        <View key={review.id}>
          <ListItem key={review.id}>
            {review.book.thumbnailUrl !== '' ? (
              <Avatar source={{ uri: review.book.thumbnailUrl }} />
            ) : (
              <Avatar>
                <MaterialIcons name="person-outline" size={36} color={COLOR.GREY} />
              </Avatar>
            )}
            <ListItem.Content>
              <ListItem.Title>{review.book.title}</ListItem.Title>
              <ListItem.Subtitle>{review.impression}</ListItem.Subtitle>
            </ListItem.Content>
          </ListItem>
        </View>
      ))}
    </View>
  );
};
