import { RouteProp, useNavigation } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import { AxiosResponse } from 'axios';
import React, { ReactElement, useState, useEffect } from 'react';
import { RefreshControl, StyleSheet, View } from 'react-native';
import { Text } from 'react-native-elements';
import { ScrollView } from 'react-native-gesture-handler';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import ProfileViewGroup from '~/components/organisms/ProfileViewGroup';
import { UserProfileV1Response } from '~/types/api/user_apiv1_response_pb';
import { BookshelfTabStackParamList } from '~/types/navigation';
import { COLOR, FONT_SIZE } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  selfIntroduction: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
    color: COLOR.TEXT_DEFAULT,
    marginTop: 10,
    padding: 16,
    alignSelf: 'stretch',
    minHeight: 100,
    fontSize: FONT_SIZE.TEXT_DEFAULT,
  },
  title: {
    color: COLOR.TEXT_TITLE,
    padding: 10,
    fontSize: FONT_SIZE.TITLE_SUBHEADER,
  },
});

interface Props {
  route: RouteProp<BookshelfTabStackParamList, 'OtherProfile'>;
  navigation: StackNavigationProp<BookshelfTabStackParamList, 'OtherProfile'>;
  actions: {
    getOtherProfile: (id: string) => Promise<AxiosResponse<UserProfileV1Response.AsObject>>;
  };
}

const OtherProfile = function OtherProfile(props: Props): ReactElement {
  const id = props.route.params.id;
  const getOtherProfile = props.actions.getOtherProfile;

  const [userInfo, setUserInfo] = useState({
    username: '',
    thumbnailUrl: '',
    selfIntroduction: '',
    rating: 0,
    reviewCount: 0,
    saleCount: 0,
    followerCount: 0,
    followCount: 0,
  });

  const navigation = useNavigation();

  const [isLoading, setIsLoading] = useState(false);

  const handleGetOtherProfile = () => {
    setIsLoading(true);
    getOtherProfile(id)
      .then((res) =>
        setUserInfo({
          username: res.data.username,
          thumbnailUrl: res.data.thumbnailUrl,
          selfIntroduction: res.data.selfIntroduction,
          rating: res.data.rating,
          reviewCount: res.data.reviewCount,
          // TODO SalesCountの実装
          saleCount: 0,
          followerCount: res.data.followerCount,
          followCount: res.data.followCount,
        }),
      )
      .catch((err) => {
        console.log('debug', err);
      })
      .finally(() => {
        setIsLoading(false);
        console.log('refreshed');
      });
  };

  useEffect(() => {
    handleGetOtherProfile();
  }, []);

  return (
    <View style={styles.container}>
      <HeaderWithBackButton title={'プロフィール'} onPress={() => navigation.goBack()} />
      <ScrollView
        refreshControl={
          <RefreshControl refreshing={isLoading} onRefresh={handleGetOtherProfile} />
        }>
        <ProfileViewGroup
          name={userInfo.username}
          avatarUrl={userInfo.thumbnailUrl}
          rating={userInfo.rating}
          reviewCount={userInfo.reviewCount}
          saleCount={userInfo.saleCount}
          followerCount={userInfo.followerCount}
          followCount={userInfo.followCount}
          buttonTitle={'フォローする'}
          handleClick={() => undefined}
        />
        <Text style={styles.selfIntroduction}>{userInfo.selfIntroduction}</Text>
        <Text style={styles.title}>出品リスト</Text>
      </ScrollView>
    </View>
  );
};

export default OtherProfile;
