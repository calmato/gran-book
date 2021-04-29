import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { Alert, StyleSheet, View } from 'react-native';
import { Text } from 'react-native-elements';
import { AuthStackParamList } from '~/types/navigation';
import HeaderWithCloseButton from '~/components/organisms/HeaderWithCloseButton';
import SignInButtonGroup from '~/components/organisms/SingInButtonGroup';
import TitleLogoText from '~/components/atoms/TitleLogoText';
import { useNavigation } from '@react-navigation/native';
import { UiContext } from '~/lib/context';
import { generateErrorMessage } from '~/lib/util/ErrorUtil';
import { Status } from '~/lib/context/ui';
import Firebase from 'firebase';

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
    signInWithFacebook: () => Promise<void>;
  };
}

const SignInSelect = function SignInSelect(props: Props): ReactElement {
  const navigation = useNavigation();
  const { setApplicationState } = React.useContext(UiContext);
  const { getAuth, signInWithFacebook } = props.actions;

  const createAlertNotifySignupError = (code) => {
    Alert.alert('サインインに失敗', `${code}`, [
      {
        text: 'OK',
      },
    ]);
  }

  Firebase.auth().onAuthStateChanged(user => {
    if (user != null) {
      console.log('We are authenticated now!');
    }
    // Do other things
  });

  const handleSignInWithFacebook = React.useCallback(async () => {
    await signInWithFacebook()
      .then(() => {
        return getAuth();
      })
      .then(() => {
        setApplicationState(Status.AUTHORIZED);
      })
      .catch((err) => {
        console.log(err);
        createAlertNotifySignupError(err);
      });
  },[getAuth, signInWithFacebook,setApplicationState]);

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
        handleSignInWithFacebook={() => handleSignInWithFacebook()}
        handleRegisterWithMail={() => navigation.navigate('SignUp')}
        handleSignInWithMail={() => navigation.navigate('SignIn')}
      />
    </View>
  );
};

export default SignInSelect;
