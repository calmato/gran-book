import React, { ReactElement } from 'react';
import { View, Text } from 'react-native';
import { Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';

const Store = function Store(): ReactElement {
  return (
    <View>
      <Header centerComponent={<HeaderText title="Gran Book" />} />
      <Text>Store</Text>
    </View>
  );
};

export default Store;
