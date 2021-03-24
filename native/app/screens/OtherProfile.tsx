import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { Text } from 'react-native-elements';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import ProfileViewGroup from '~/components/organisms/ProfileViewGroup';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  selfIntroduction: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
    color: COLOR.TEXT_DEFAULT,
    marginTop: 10,
    padding: 10,
    alignSelf: 'stretch',
    minHeight: 100,
  },
  title: {
    color: COLOR.TEXT_TITLE,
    padding: 10,
  },
});

const userInfo = {
  name: 'にしくん',
  avatarUrl: 'https://pbs.twimg.com/profile_images/1312909954148253696/Utr-sa_Y_400x400.jpg',
  rating: 4.8,
  reviewNum: 20,
  saleNum: 3,
  followerNum: 1,
  followNum: 1,
  selfIntroduction: 'たくさん買うぞー。',
};

const OtherProfile = function OtherProfile(): ReactElement {
  return (
    <View style={styles.container}>
      <HeaderWithBackButton title="プロフィール" onPress={() => undefined} />
      <ProfileViewGroup
        name={userInfo.name}
        avatarUrl={userInfo.avatarUrl}
        rating={userInfo.rating}
        reviewNum={userInfo.reviewNum}
        saleNum={userInfo.saleNum}
        followerNum={userInfo.followerNum}
        followNum={userInfo.followNum}
        buttonTitle={'フォローする'}
        handleClick={() => undefined}
      />
      <Text style={styles.selfIntroduction}>{userInfo.selfIntroduction}</Text>
      <Text style={styles.title}>出品リスト</Text>
    </View>
  );
};

export default OtherProfile;
