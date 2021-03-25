import { useNavigation } from '@react-navigation/native';
import React, { ReactElement, useState } from 'react';
import { RefreshControl, StyleSheet, View } from 'react-native';
import { Text } from 'react-native-elements';
import { ScrollView } from 'react-native-gesture-handler';
import BookList from '~/components/molecules/BookList';
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

interface Props {
  id: string;
  username: string;
  selfIntroduction: string;
  thumbnailUrl: string | undefined;
  gender: number;
  actions: any;
}

const OwnProfile = function OwnProfile(props: Props): ReactElement {
  // TODO 出品数・フォロワー数・フォロー数・星レート・レビュー数を実装
  const userInfo = {
    rating: 2.4,
    reviewCount: 20,
    saleCount: 3,
    followerCount: 20,
    followCount: 5,
  };
  const navigation = useNavigation();

  const [hasGottonProfile, setHasGottonProfile] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const { getOwnProfile } = props.actions;

  const handleGetOwnProfile = () => {
    setIsLoading(true);
    getOwnProfile(props.id)
      .then(() => {
        setIsLoading(false);
      })
      .catch((err) => {
        console.log('debug', err);
      });
  };

  // OwnProfileをGetするとAuthを書き換えるので再レンダリングされる
  // =>またゲットが呼ばれるの無限ループになるので、
  // Booleanを入れて制御。他にいい案あれば募集
  if (!hasGottonProfile) {
    setHasGottonProfile(true);
    handleGetOwnProfile();
  }

  return (
    <View style={styles.container}>
      <HeaderWithBackButton title="プロフィール" onPress={() => navigation.goBack()} />
      <ProfileViewGroup
        name={props.username}
        avatarUrl={props.thumbnailUrl}
        rating={userInfo.rating}
        reviewCount={userInfo.reviewCount}
        saleCount={userInfo.saleCount}
        followerCount={userInfo.followerCount}
        followCount={userInfo.followCount}
        buttonTitle={'プロフィールを編集'}
        handleClick={() => navigation.navigate('ProfileEdit')}
      />
      <Text style={styles.selfIntroduction}>{props.selfIntroduction}</Text>
      <Text style={styles.title}>出品リスト</Text>
    </View>
  );
};

export default OwnProfile;
