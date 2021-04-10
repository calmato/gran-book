import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { Text } from 'react-native-elements';
import { AuthStackParamList } from '~/types/navigation';
import HeaderWithCloseButton from '~/components/organisms/HeaderWithCloseButton';
import SignInButtonGroup from '~/components/organisms/SingInButtonGroup';
import TitleLogoText from '~/components/atoms/TitleLogoText';
import Firebase from 'firebase';
import * as WebBrowser from 'expo-web-browser';
import {useAuthRequest} from 'expo-auth-session/providers/facebook';
import { ResponseType } from 'expo-auth-session';
import { useNavigation } from '@react-navigation/native';

WebBrowser.maybeCompleteAuthSession();

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
  actions: {
    getAuth: () => Promise<void>;
  };
}

const SignInSelect = function SignInSelect(props: Props): ReactElement {
  const navigation = useNavigation();

  const [_request, response, promptAsync] = useAuthRequest({
    responseType: ResponseType.Token,
    clientId: '254032153080907',
  });

  React.useEffect(() => {
    if (response?.type === 'success') {
      const { access_token } = response.params;
      
      const credential = Firebase.auth.FacebookAuthProvider.credential(access_token);
      // Sign in with the credential from the Facebook user.
      Firebase.auth().signInWithCredential(credential);
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
        handleSignInWithFacebook={() => promptAsync()}
        handleRegisterWithMail={() => navigation.navigate('SignUp')}
        handleSignInWithMail={() => navigation.navigate('SignIn')}
      />
    </View>
  );
};

export default SignInSelect;
