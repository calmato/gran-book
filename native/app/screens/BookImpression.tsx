import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';

const BookImpression = function BookImpression(): ReactElement {
  return (
    <View>
      <HeaderWithBackButton
        title='感想'
        onPress={() => undefined}
      />
    </View>
  );
};

export default BookImpression;
