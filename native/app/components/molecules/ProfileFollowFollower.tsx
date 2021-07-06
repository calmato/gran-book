import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import ProfileNumberOfThings from '../atoms/ProfileNumberOfThings';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    alignItems: 'center',
    padding: 5,
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
  divider: {
    borderLeftWidth: 0.5,
    borderLeftColor: COLOR.TEXT_GRAY,
    alignSelf: 'stretch',
  },
});

interface Props {
  saleCount: number;
  followerCount: number;
  followCount: number;
}

const ProfileFollowFollwer = function ProfileFollowFollwer(props: Props): ReactElement {
  return (
    <View style={styles.container}>
      <ProfileNumberOfThings name={'出品数'} numberOfThings={props.saleCount} />
      <View style={styles.divider} />
      <ProfileNumberOfThings name={'フォロワー'} numberOfThings={props.followerCount} />
      <View style={styles.divider} />
      <ProfileNumberOfThings name={'フォロー中'} numberOfThings={props.followCount} />
    </View>
  );
};

export default ProfileFollowFollwer;
