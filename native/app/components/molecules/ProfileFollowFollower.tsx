import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { COLOR } from '~~/constants/theme';
import ProfileNumberOfThings from '../atoms/ProfileNumberOfThings';

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    alignItems: 'center',
    padding: 5,
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
  divider: {
    borderLeftWidth: 1,
    borderLeftColor: COLOR.TEXT_GRAY,
    alignSelf: 'stretch',
  },
});

interface Props {
  saleNum: number
  followerNum: number
  followNum: number
}

const ProfileFollowFollwer = function ProfileFollowFollwer(props: Props): ReactElement {
  return (
    <View style={styles.container}>
      <ProfileNumberOfThings
        name={'出品数'}
        numberOfThings={props.saleNum}
      />
      <View style={styles.divider}/>
      <ProfileNumberOfThings
        name={'フォロワー'}
        numberOfThings={props.followerNum}
      />
      <View style={styles.divider}/>
      <ProfileNumberOfThings
        name={'フォロー中'}
        numberOfThings={props.followNum}
      />
    </View>
  );
};

export default ProfileFollowFollwer;
