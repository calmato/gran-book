import { useNavigation } from '@react-navigation/native';
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
    minHeight: 100,
  },
  title: {
    color: COLOR.TEXT_TITLE,
    padding: 10,
  }
});

interface Props {
  username: string, 
  selfIntroduction: string | undefined, 
  thumbnailUrl: string | undefined, 
  gender: number
}

const OwnProfile = function OwnProfile( props : Props): ReactElement {
  // TODO 出品数・フォロワー数・フォロー数・星レート・レビュー数を実装
  const userInfo = 
{
  rating: 2.4,
  reviewNum: 20,
  saleNum: 3,
  followerNum: 20,
  followNum: 5,
};
  const navigation = useNavigation();
  return (
    <View style={styles.container}>
      <HeaderWithBackButton
        title="プロフィール"
        onPress={()=>navigation.goBack()}
      />
      <ProfileViewGroup
        name={props.username}
        avatarUrl={props.thumbnailUrl}
        rating={userInfo.rating}
        reviewNum={userInfo.reviewNum}
        saleNum={userInfo.saleNum}
        followerNum={userInfo.followerNum}
        followNum={userInfo.followNum}
        buttonTitle={'プロフィールを編集'}
        handleClick={() => navigation.navigate('ProfileEdit', {
          username: props.username, 
          selfIntroduction: props.selfIntroduction,
          thumbnailUrl: props.thumbnailUrl,
          gender: props.gender,
        })}
      />
      <Text style={styles.bio}>{props.selfIntroduction}</Text>
      <Text style={styles.title}>出品リスト</Text>
    </View>
  );
};

export default OwnProfile;
