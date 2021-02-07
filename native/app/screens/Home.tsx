import React, { ReactElement } from 'react';
import { View, Text } from 'react-native';
import { Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import SearchBar from '~/components/molecules/SearchBar';

const Home = function Home(): ReactElement {
  return (
    <View>
      <Header centerComponent={<HeaderText title="Gran Book"/>} />
      <SearchBar />
      <Text>Home画面</Text>
    </View>
  );
};

export default Home;
