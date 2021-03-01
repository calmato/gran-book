import React, { ReactElement } from 'react';
import { ScrollView } from 'react-native';
import { View } from 'react-native';
import { Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import BookList from '~/components/molecules/BookList';
import SearchBar from '~/components/molecules/SearchBar';

const Home = function Home(): ReactElement {
  return (
    <View>
      <ScrollView
        stickyHeaderIndices={[0]}
      >
        <Header centerComponent={<HeaderText title="Gran Book"/>} />
        <SearchBar />
        <BookList />
      </ScrollView>
    </View>
  );
};

export default Home;
