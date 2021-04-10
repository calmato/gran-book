import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { Alert, StyleSheet, View } from 'react-native';
import { Text } from 'react-native-elements';
import { AuthStackParamList } from '~/types/navigation';
import HeaderWithCloseButton from '~/components/organisms/HeaderWithCloseButton';
import SignInButtonGroup from '~/components/organisms/SingInButtonGroup';
import TitleLogoText from '~/components/atoms/TitleLogoText';
import * as WebBrowser from 'expo-web-browser';
import * as Google from 'expo-auth-session/providers/google';
import Firebase from 'firebase';
import { useNavigation } from '@react-navigation/native';
import { UiContext } from '~/lib/context';
import { Status } from '~/lib/context/ui';
import { generateErrorMessage } from '~/lib/util/ErrorUtil';

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
  actions: {
    getAuth: () => Promise<void>;
  };
}

WebBrowser.maybeCompleteAuthSession();

const SignInSelect = function SignInSelect(props: Props): ReactElement {
  const navigation = useNavigation();
  const { setApplicationState } = React.useContext(UiContext);
  const { getAuth } = props.actions;

  const [_request, response, promptAsync] = Google.useIdTokenAuthRequest(
    {
      clientId: '711103859602-pl5m005fp0bhhum9lm99fgoinneaar7m.apps.googleusercontent.com',
    },
  );

  const createAlertNotifySignupError = (code: number) =>
    Alert.alert('サインインに失敗', `${generateErrorMessage(code)}`, [
      {
        text: 'OK',
      },
    ]);

  const handleSignInWithGoogle = () => {
    promptAsync()
      .then(() => {
        console.log('loggedIn');
        return getAuth();
      })
      .then(() => {
        console.log('gotAuth');
        setApplicationState(Status.AUTHORIZED);
      })
      .catch((err) => {
        console.log(err);
        createAlertNotifySignupError(err.code);
      });
  };

  React.useEffect(() => {
    if (response?.type === 'success') {
      const { id_token } = response.params;
      
      const credential = Firebase.auth.GoogleAuthProvider.credential(id_token);
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
        handleSignInWithGoogle={() => handleSignInWithGoogle()}
        handleRegisterWithMail={() => navigation.navigate('SignUp')}
        handleSignInWithMail={() => navigation.navigate('SignIn')}
      />
    </View>
  );
};

export default SignInSelect;
