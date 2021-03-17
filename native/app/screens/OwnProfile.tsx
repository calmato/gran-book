import { useNavigation } from '@react-navigation/native';
import React, { ReactElement, useState } from 'react';
import { StyleSheet, View } from 'react-native';
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
  }
});

interface Props {
  id: string,
  username: string, 
  selfIntroduction: string, 
  thumbnailUrl: string | undefined, 
  gender: number,
  followCount: number,
  followerCount: number,
  reviewCount: number,
  rating: number,
  products: Array<
    {id:number, 
    name: string, 
    thumbnailUrl: string, 
    authors: Array<string>
    }>,
  actions: { getOwnProfile: (id: string) => Promise<void>, },
}

const OwnProfile = function OwnProfile( props : Props): ReactElement {
  const userInfo = 
  {
    username: props.username, 
    selfIntroduction: props.selfIntroduction, 
    thumbnailUrl: props.thumbnailUrl, 
    gender: props.gender,
    followCount: props.followCount,
    followerCount: props.followerCount,
    rating: props.rating,
    reviewCount: props.reviewCount,
    saleCount: props.products.length,
  };
  const navigation = useNavigation();

  const [hasGottonProfile, setHasGottonProfile] = useState(false);
  const { getOwnProfile } = props.actions;

  // OwnProfileをGetするとAuthを書き換えるので再レンダリングされる
  // =>またゲットが呼ばれるの無限ループになるので、
  // Booleanを入れて制御。他にいい案あれば募集
  if(!hasGottonProfile){
    setHasGottonProfile(true);
    getOwnProfile(props.id)
    .catch((err) => {
      console.log('debug', err);
    });
  }

  return (
    <View style={styles.container}>
      <HeaderWithBackButton
        title="プロフィール"
        onPress={()=>navigation.goBack()}
      />
      <ScrollView>
        <ProfileViewGroup
          name={userInfo.username}
          avatarUrl={props.thumbnailUrl}
          rating={props.rating}
          reviewCount={userInfo.reviewCount}
          saleCount={userInfo.saleCount}
          followerCount={props.followerCount}
          followCount={props.followCount}
          buttonTitle={'プロフィールを編集'}
          handleClick={() => navigation.navigate('ProfileEdit')}
        />
        <Text style={styles.selfIntroduction}>{props.selfIntroduction}</Text>
        <Text style={styles.title}>出品リスト</Text>
        <BookList />
      </ScrollView>
    </View>
  );
};

export default OwnProfile;
