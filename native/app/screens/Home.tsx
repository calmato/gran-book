import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import RecommendBooks from '~/components/organisms/RecommendBooks';
import { useRecommendBooks } from '~/hooks/useRecommendBooks';

const styles = StyleSheet.create({});

const Home = function Home(): ReactElement {
  const { recommendBooks } = useRecommendBooks();

  return (
    <View>
      <Header centerComponent={<HeaderText title="Gran Book" />} />
      <RecommendBooks books={recommendBooks} />
    </View>
  );
};

export default Home;
