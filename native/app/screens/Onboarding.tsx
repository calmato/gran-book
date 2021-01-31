import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { StyleSheet, View, Image } from 'react-native';
import { Button, colors, Text } from 'react-native-elements';
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
    margin: 12
  },
  registerButton: {
    width: 300,
    height: 50,
  },
  signInButton: {
    borderColor: colors.grey0,
    width: 120,
    height: 30,
  },
  signInButtonFont: {
    fontSize: 14,
    color: colors.grey0
  }
});

type OnboardingNavigationProp = StackNavigationProp<RootStackParamList, 'SignInSelect'>;

interface Props {
  navigation: OnboardingNavigationProp
}

const Onboarding = function Onboarding(props: Props): ReactElement {
  const navigation = props.navigation;

  return (
    <View style={styles.container}>
      <TitleLogoText style={styles.titleStyle} text="Gran Book" />
      <Text style={styles.subTitleStyle}>読書管理・本専用のフリマアプリ</Text>
      <Image style={styles.logo} source={logo}/>
      <Button
        onPress={()=> navigation.navigate('SignInSelect')}
        title="新規登録する"
        buttonStyle={styles.registerButton}
      />
      <Text style={{marginTop: 30, marginBottom: 12}}>既にIDをお持ちの方はこちら</Text>
      <Button title="サインイン" buttonStyle={styles.signInButton} titleStyle={styles.signInButtonFont} type="outline" />
      <Text style={{margin:30}}>利用規約</Text>
    </View>
  );
};

export default Onboarding;
