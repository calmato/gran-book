import { useNavigation } from '@react-navigation/native';
import * as Google from 'expo-auth-session/providers/google';
import * as WebBrowser from 'expo-web-browser';
import React, { ReactElement } from 'react';
import { Alert, StyleSheet, View } from 'react-native';
import { Text } from 'react-native-elements';
import TitleLogoText from '~/components/atoms/TitleLogoText';
import HeaderWithCloseButton from '~/components/organisms/HeaderWithCloseButton';
import SignInButtonGroup from '~/components/organisms/SingInButtonGroup';
import { UiContext, Status } from '~/context';
import firebase from '~/lib/firebase';
import * as LocalStorage from '~/lib/local-storage';
import { generateErrorMessage } from '~/lib/util/ErrorUtil';
import { Auth } from '~/store/models';
import { useReduxDispatch } from '~/store/modules';
import { setAuth } from '~/store/modules/auth';
import { FONT_SIZE } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  titleStyle: {
    margin: 24,
  },
  text: {
    fontSize: FONT_SIZE.TEXT_ALERT,
    lineHeight: 18.5,
  },
});

interface Props {
  actions: {
    getAuth: () => Promise<void>;
    registerForPushNotifications: () => Promise<void>;
  };
}

WebBrowser.maybeCompleteAuthSession();

const SignInSelect = function SignInSelect(props: Props): ReactElement {
  const navigation = useNavigation();
  const { setApplicationState } = React.useContext(UiContext);
  const { getAuth, registerForPushNotifications } = props.actions;

  const [_request, response, handleSignInWithGoogle] = Google.useIdTokenAuthRequest({
    clientId: process.env.CLIENT_ID_FOR_GOOGLE,
  });

  const createAlertNotifySignupError = (code: number) =>
    Alert.alert('サインインに失敗', `${generateErrorMessage(code)}`, [
      {
        text: 'OK',
      },
    ]);

  const dispatch = useReduxDispatch();

  React.useEffect(() => {
    if (response?.type === 'success') {
      const { id_token } = response.params;
      const credential = firebase.auth.GoogleAuthProvider.credential(id_token);
      firebase
        .auth()
        .signInWithCredential(credential)
        .then(async () => {
          await firebase
            .auth()
            .currentUser?.getIdToken(true)
            .then(async (token) => {
              const user = firebase.auth().currentUser!;

              const values: Auth.AuthValues = {
                id: user.uid,
                email: user.email || undefined,
                emailVerified: true,
                token: token,
              };

              const model: Auth.Model = {
                ...Auth.initialState,
                id: values.id,
                token: values.token,
                email: values.email || '',
                emailVerified: values.emailVerified || false,
              };
              dispatch(setAuth(values));
              await LocalStorage.AuthStorage.save(model);
            });
        })
        .then(() => {
          return registerForPushNotifications();
        })
        .then(async () => {
          await getAuth();
        })
        .then(() => {
          setApplicationState(Status.AUTHORIZED);
        })
        .catch((err) => {
          console.log(err);
          createAlertNotifySignupError(err);
        });
    }
  }, [response, dispatch, getAuth, registerForPushNotifications, setApplicationState]);

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
