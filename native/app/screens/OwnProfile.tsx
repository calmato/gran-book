import { useNavigation } from '@react-navigation/native';
import React, { ReactElement, useState } from 'react';
import { RefreshControl, StyleSheet, View } from 'react-native';
import { Text } from 'react-native-elements';
import { ScrollView } from 'react-native-gesture-handler';
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
  username: string;
  selfIntroduction: string;
  thumbnailUrl: string | undefined;
  gender: number;
  rating: number;
  reviewCount: number;
  saleCount: number;
  followerCount: number;
  followCount: number;
  actions: {
    getOwnProfile: () => Promise<void>;
  };
}

const OwnProfile = function OwnProfile(props: Props): ReactElement {
  const userInfo = {
    rating: props.rating,
    reviewCount: props.reviewCount,
    saleCount: props.saleCount,
    followerCount: props.followCount,
    followCount: props.followCount,
  };
  const navigation = useNavigation();

  const [hasGottonProfile, setHasGottonProfile] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const { getOwnProfile } = props.actions;

  const handleGetOwnProfile = () => {
    setIsLoading(true);
    getOwnProfile()
      .then(() => {
        setIsLoading(false);
      })
      .catch((err) => {
        console.log('debug', err);
      });
  };

  if (!hasGottonProfile) {
    setHasGottonProfile(true);
    handleGetOwnProfile();
  }

  return (
    <View style={styles.container}>
      <HeaderWithBackButton title={props.username} onPress={() => navigation.goBack()} />
      <ScrollView
        refreshControl={<RefreshControl refreshing={isLoading} onRefresh={handleGetOwnProfile} />}>
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
      </ScrollView>
    </View>
  );
};

export default OwnProfile;
