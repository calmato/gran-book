import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import ProfileBasicInfoGroup from '../molecules/ProfileBasicInfoGroup';

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    alignItems: 'center',
  },
});

interface Props {
  name: string
  avatar_url: string
  rating: number
  numberOfReviews: number
}

const ProfileViewGroup = function ProfileViewGroup(props:Props): ReactElement {
  return (
    <View style={styles.container}>
      <ProfileBasicInfoGroup
        name={props.name}
        avatar_url={props.avatar_url}
        rating={props.rating}
        numberOfReviews={props.numberOfReviews}
      />
    </View>
  );
};

export default ProfileViewGroup;
