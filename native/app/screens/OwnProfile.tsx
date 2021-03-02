import { RouteProp, useNavigation } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { Text } from 'react-native-elements';
import { Route } from 'react-native-tab-view';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import ProfileViewGroup from '~/components/organisms/ProfileViewGroup';
import { UserInfoStackParamList } from '~/types/navigation';
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
    minHeight: 100,
  },
  title: {
    color: COLOR.TEXT_TITLE,
    padding: 10,
  }
});

type OwnProfileNavigationProp = StackNavigationProp<
  UserInfoStackParamList,
  'OwnProfile'
>

type OwnProfileRouteProp = RouteProp<
  UserInfoStackParamList, 
  'OwnProfile'
>

interface Props {
  route: OwnProfileRouteProp,
  navigation: OwnProfileNavigationProp,
}

const OwnProfile = function OwnProfile({ route, navigation }: Props): ReactElement {
  const {username, selfIntroduction, thumbnailUrl, gender} = route.params;
  // TODO 出品数・フォロワー数・フォロー数・星レート・レビュー数を実装
  const userInfo = 
{
  rating: 2.4,
  reviewNum: 20,
  saleNum: 3,
  followerNum: 20,
  followNum: 5,
};
  return (
    <View style={styles.container}>
      <HeaderWithBackButton
        title="プロフィール"
        onPress={()=>navigation.goBack()}
      />
      <ProfileViewGroup
        name={username}
        avatarUrl={thumbnailUrl}
        rating={userInfo.rating}
        reviewNum={userInfo.reviewNum}
        saleNum={userInfo.saleNum}
        followerNum={userInfo.followerNum}
        followNum={userInfo.followNum}
        buttonTitle={'プロフィールを編集'}
        handleClick={() => navigation.navigate('ProfileEdit', {
          username: username, 
          selfIntroduction: selfIntroduction,
          thumbnailUrl: thumbnailUrl,
          gender: gender,
        })}
      />
      <Text style={styles.bio}>{selfIntroduction}</Text>
      <Text style={styles.title}>出品リスト</Text>
    </View>
  );
};

export default OwnProfile;
