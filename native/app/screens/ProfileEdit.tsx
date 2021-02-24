import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import ChangeIconGroup from '~/components/organisms/ChangeIconGroup';
import ChangeNickname from '~/components/organisms/ChangeNickname';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';

const styles = StyleSheet.create({
  container: {

  },
});

interface Props {}

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

const ProfileEdit = function ProfileEdit(props: Props): ReactElement {
  return (
  <View>
    <HeaderWithBackButton 
      title='プロフィール編集'
      onPress={()=>undefined}
    />
    <ChangeIconGroup
      avatarUrl={userInfo.avatarUrl}
      handleOnPressed={()=>undefined}
    />
    <ChangeNickname
    defaultValue={userInfo.name}
    handelOnChangeText={()=>undefined}
    />
  </View>
  );
}

export default ProfileEdit
