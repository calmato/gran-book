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
  avatarUrl: string
  rating: number
  reviewNum: number
}

const ProfileViewGroup = function ProfileViewGroup(props:Props): ReactElement {
  return (
    <View style={styles.container}>
      <ProfileBasicInfoGroup
        name={props.name}
        avatarUrl={props.avatarUrl}
        rating={props.rating}
        reviewNum={props.reviewNum}
      />
    </View>
  );
};

export default ProfileViewGroup;
