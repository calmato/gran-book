import { useNavigation } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { Text } from 'react-native-elements';
import { StackParamList } from '~/types/navigation';
import HeaderWithCloseButton from '~/components/organisms/HeaderWithCloseButton';
import SignInButtonGroup from '~/components/organisms/SingInButtonGroup';
import TitleLogoText from '~/components/atoms/TitleLogoText';

type SignInNavigationProp = StackNavigationProp<StackParamList, 'Onboarding'>

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  titleStyle: {
    margin: 24,
  },
  text: {
    fontSize: 12,
    lineHeight: 18.5,
  }
});

interface Props {
  navigation: SignInNavigationProp
}

const SignIn = function SignIn(props: Props): ReactElement {
  const navigation = useNavigation();
  
  return (
    <View style={styles.container}>
      <HeaderWithCloseButton
        title="登録/サインイン"
        onPress={() => navigation.goBack()}
      />
      <TitleLogoText
        style={styles.titleStyle}
        text="Gran Book"
      />
      <View style={{ width: 340, marginBottom: 12 }}>
        <Text style={styles.text}>サインイン及び新規登録について、Gran Bookの利用規約及び個人情報の取り扱いについて、アプリケーションプライバシーポリシーに同意するものとします。</Text>
      </View>
      <SignInButtonGroup />
    </View>
  );
};

export default SignIn;
