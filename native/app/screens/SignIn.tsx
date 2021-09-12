import { useNavigation } from '@react-navigation/native';
import firebase from 'firebase';
import React, { ReactElement, useCallback, useMemo, useState } from 'react';
import { Alert, StyleSheet, View } from 'react-native';
import { Button } from 'react-native-elements';
import ForgotPasswordButton from '~/components/molecules/ForgotPasswordButton';
import MailInput from '~/components/molecules/MailInput';
import PasswordInput from '~/components/molecules/PasswordInput';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { emailValidation, passwordValidation } from '~/lib/validation';
import { SignInForm } from '~/types/forms';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
});

interface Props {
  actions: {
    signInWithEmailAndPassword: (payload: SignInForm) => Promise<firebase.User | null>;
  };
}

const SignIn = function SignIn(props: Props): ReactElement {
  const navigation = useNavigation();
  const { signInWithEmailAndPassword } = props.actions;

  const [formData, setValue] = useState<SignInForm>({
    email: '',
    password: '',
  });

  const passwordError: boolean = useMemo((): boolean => {
    return !passwordValidation(formData.password);
  }, [formData.password]);

  const emailError: boolean = useMemo((): boolean => {
    return !emailValidation(formData.email);
  }, [formData.email]);

  const canSubmit = useMemo((): boolean => {
    return !(emailError || passwordError);
  }, [emailError, passwordError]);

  const createAlertNotifySignInError = useCallback(
    (title: string, message: string) =>
      Alert.alert(title, message, [
        {
          text: 'OK',
          onPress: async () => {
            console.log(firebase.auth().currentUser);
            firebase
              .auth()
              .currentUser?.sendEmailVerification()
              .then(() => console.log('メールを送信しました'));
          },
        },
      ]),
    [],
  );

  const handleSubmit = React.useCallback(async () => {
    try {
      const user = await signInWithEmailAndPassword(formData);
      if (!user) {
        createAlertNotifySignInError('サインインに失敗', 'メールから登録を完了してください。');
      }
    } catch (e) {
      console.log(e);
      createAlertNotifySignInError(
        'サインインに失敗',
        'メールアドレスまたはパスワードが間違っています。',
      );
    }
  }, [formData, createAlertNotifySignInError, signInWithEmailAndPassword]);

  return (
    <View style={styles.container}>
      <HeaderWithBackButton title="サインイン" onPress={() => navigation.goBack()} />
      <MailInput
        onChangeText={(text) => setValue({ ...formData, email: text })}
        hasError={emailError}
        value={formData?.email}
        sameEmailError={false}
      />
      <PasswordInput
        onChangeText={(text) => setValue({ ...formData, password: text })}
        value={formData.password}
        placeholder="パスワード"
        errorMessage="パスワードは6文字以上32文字以下でなければいけません．"
        hasError={passwordError}
      />
      <Button disabled={!canSubmit} onPress={handleSubmit} title="サインイン" />
      <ForgotPasswordButton onPress={() => navigation.navigate('PasswordReset')} />
    </View>
  );
};

export default SignIn;
