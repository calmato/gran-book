import React, { ReactElement } from 'react';
import { View } from 'react-native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';

const OwnProfile = function OwnProfile(): ReactElement {
  return (
    <View>
      <HeaderWithBackButton
        title="プロフィール"
        onPress={()=>undefined}
      />
    </View>
  );
};

export default OwnProfile;
