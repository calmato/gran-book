import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { Text } from 'react-native-elements';
import { AuthStackParamList } from '~/types/navigation';
import HeaderWithCloseButton from '~/components/organisms/HeaderWithCloseButton';
import SignInButtonGroup from '~/components/organisms/SingInButtonGroup';
import TitleLogoText from '~/components/atoms/TitleLogoText';
import * as WebBrowser from 'expo-web-browser';
import * as Google from 'expo-auth-session/providers/google';

type AuthSignInNavigationProp = StackNavigationProp<AuthStackParamList, 'SignUp'>;

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
  },
});

interface Props {
  navigation: AuthSignInNavigationProp;
}

WebBrowser.maybeCompleteAuthSession();

const SignInSelect = function SignInSelect(props: Props): ReactElement {
  const navigation = props.navigation;

  const [request, response, promptAsync] = Google.useAuthRequest({
    expoClientId: 'GOOGLE_GUID.apps.googleusercontent.com',
  });

  React.useEffect(() => {
    if (response?.type === 'success') {
      const { authentication } = response;
      }
  }, [response]);

  return (
    <View style={styles.container}>
      <HeaderWithCloseButton title="登録/サインイン" onPress={() => navigation.goBack()} />
      <TitleLogoText style={styles.titleStyle} text="Gran Book" />
      <View style={{ width: 340, marginBottom: 12 }}>
        <Text style={styles.text}>
          サインイン及び新規登録について、Gran
          Bookの利用規約及び個人情報の取り扱いについて、アプリケーションプライバシーポリシーに同意するものとします。
        </Text>
      </View>
      <SignInButtonGroup
        handleSignInWithGoogle={() => promptAsync()}
        handleRegisterWithMail={() => navigation.navigate('SignUp')}
        handleSignInWithMail={() => navigation.navigate('SignIn')}
      />
    </View>
  );
};

export default SignInSelect;
