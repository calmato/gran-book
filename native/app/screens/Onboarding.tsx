import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { StyleSheet, View, Image } from 'react-native';
import { Button, Text } from 'react-native-elements';
import TitleLogoText from '~/components/atoms/TitleLogoText';
import { RootStackParamList } from '~/types/navigation';
import logo from '~~/assets/logo.png';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  logo: {
    width: 218.75,
    height: 250,
    margin: 30,
  },
  titleStyle: {
    marginTop: 12,
  },
  subTitleStyle: {
    margin: 12,
  },
  signInButton: {
    width: 300,
    height: 50,
  },
});

type OnboardingNavigationProp = StackNavigationProp<RootStackParamList, 'SignInSelect'>;

interface Props {
  navigation: OnboardingNavigationProp;
}

const Onboarding = function Onboarding(props: Props): ReactElement {
  const navigation = props.navigation;

  return (
    <View style={styles.container}>
      <TitleLogoText style={styles.titleStyle} text="Gran Book" />
      <Text style={styles.subTitleStyle}>読書管理・本専用のフリマアプリ</Text>
      <Image style={styles.logo} source={logo} />
      <Button
        onPress={() => navigation.navigate('SignInSelect')}
        title="サインイン"
        buttonStyle={styles.signInButton}
      />
      <Text style={{ margin: 30 }}>利用規約</Text>
    </View>
  );
};

export default Onboarding;
