import React, { ReactElement } from 'react';
import { View, Text } from 'react-native';
import { Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';

const Bookshelf = function Bookshelf(): ReactElement {
  return (
    <View>
      <Header centerComponent={<HeaderText title="Gran Book"/>} />
      <Text>本棚画面</Text>
    </View>
  );
};

export default Bookshelf;
