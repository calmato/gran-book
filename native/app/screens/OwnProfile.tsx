import React, { ReactElement } from 'react';
import { View } from 'react-native';
import { Divider } from 'react-native-elements';
import ProfileFollowFollwer from '~/components/molecules/ProfileFollowFollower';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import ProfileViewGroup from '~/components/organisms/ProfileViewGroup';

const userInfo = 
{
  name: 'hamachans',
  avatar_url: 'https://pbs.twimg.com/profile_images/1312909954148253696/Utr-sa_Y_400x400.jpg',
  rating: 2.4,
  numberOfReviews: 20,
  numberOfSales: 3,
  numberOfFollowers: 20,
  numberOfFollows: 5,
}

const OwnProfile = function OwnProfile(): ReactElement {
  return (
    <View>
      <HeaderWithBackButton
        title="プロフィール"
        onPress={()=>undefined}
      />
      <ProfileViewGroup
      name={userInfo.name}
      avatar_url={userInfo.avatar_url}
      rating={userInfo.rating}
      numberOfReviews={userInfo.numberOfReviews}
      />
      <ProfileFollowFollwer
      numberOfSales={userInfo.numberOfSales}
      numberOfFollowers={userInfo.numberOfFollowers}
      numberOfFollows={userInfo.numberOfFollows}
      />
    </View>
  );
};

export default OwnProfile;
