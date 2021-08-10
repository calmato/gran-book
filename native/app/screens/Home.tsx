import React, { ReactElement } from 'react';
import { View, Text } from 'react-native';
import { Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';

const Home = function Home(): ReactElement {
  return (
    <View>
      <Header centerComponent={<HeaderText title="Gran Book" />} />
      <Text>ホーム画面</Text>
    </View>
  );
};

export default Home;
