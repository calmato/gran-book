import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';


const styles = StyleSheet.create({});

interface Props {}

const Home = function Home(props: Props): ReactElement {
  return (
    <View>
      <Header centerComponent={<HeaderText title="Gran Book"/>} />
      <Text>Home</Text>
    </View>
  );
};

// Home.defaultProps={}

export default Home;
