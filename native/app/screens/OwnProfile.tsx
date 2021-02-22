import React, { ReactElement } from 'react';
import { View } from 'react-native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import ProfileViewGroup from '~/components/organisms/ProfileViewGroup';

const userInfo = 
{
  name: 'hamachans',
  avatar_url: 'https://pbs.twimg.com/profile_images/1312909954148253696/Utr-sa_Y_400x400.jpg',
  rating: 2.4,
  numberOfReviews: 20,
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
    </View>
  );
};

export default OwnProfile;
