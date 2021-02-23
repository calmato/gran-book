import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { Divider } from 'react-native-elements';
import { COLOR } from '~~/constants/theme';
import ProfileNumberOfThings from '../atoms/ProfileNumberOfThings';

const styles = StyleSheet.create({
container: {
  // flex: 1 ,
  flexDirection: 'row',
  alignItems: 'center',
  padding: 5,
  backgroundColor: COLOR.BACKGROUND_WHITE,
},
divider: {
  borderLeftWidth: 1,
  borderLeftColor: COLOR.TEXT_GRAY,
  alignSelf: 'stretch',
}
});

interface Props {
  numberOfSales: number
  numberOfFollowers: number
  numberOfFollows: number
}

const ProfileFollowFollwer = function ProfileFollowFollwer(props: Props): ReactElement {
  return (
    <View style={styles.container}>
      <ProfileNumberOfThings
      name={'出品数'}
      numberOfThings={props.numberOfSales}
      />
      <View style={styles.divider}/>
      <ProfileNumberOfThings
      name={'フォロワー'}
      numberOfThings={props.numberOfFollowers}
      />
      <View style={styles.divider}/>
      <ProfileNumberOfThings
      name={'フォロー中'}
      numberOfThings={props.numberOfFollows}
      />
    </View>
  );
};

export default ProfileFollowFollwer;
