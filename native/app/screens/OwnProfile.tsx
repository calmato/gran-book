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
  bio: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
    color: COLOR.TEXT_DEFAULT,
    marginTop: 10,
    padding: 10,
    alignSelf: 'stretch',
  },
  title: {
    color: COLOR.TEXT_TITLE,
    padding: 10,
  }
});

const userInfo = 
{
  name: 'hamachans',
  avatarUrl: 'https://pbs.twimg.com/profile_images/1312909954148253696/Utr-sa_Y_400x400.jpg',
  rating: 2.4,
  reviewNum: 20,
  saleNum: 3,
  followerNum: 20,
  followNum: 5,
  bio: 'よろしくお願いします。',
};

const OwnProfile = function OwnProfile(): ReactElement {
  return (
    <View style={styles.container}>
      <HeaderWithBackButton
        title="プロフィール"
        onPress={()=>undefined}
      />
      <ProfileViewGroup
        name={userInfo.name}
        avatarUrl={userInfo.avatarUrl}
        rating={userInfo.rating}
        reviewNum={userInfo.reviewNum}
        saleNum={userInfo.saleNum}
        followerNum={userInfo.followerNum}
        followNum={userInfo.followNum}
        buttonTitle={'フォローする'}
        handleButtonPress={() => undefined}
      />
      <Text style={styles.bio}>{userInfo.bio}</Text>
      <Text style={styles.title}>出品リスト</Text>
    </View>
  );
};

export default OwnProfile;
