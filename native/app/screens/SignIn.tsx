import { useNavigation } from '@react-navigation/native';
import React, { ReactElement, useMemo, useState } from 'react';
import { Alert, StyleSheet, View } from 'react-native';
import { Button } from 'react-native-elements';
import ForgotPasswordButton from '~/components/molecules/ForgotPasswordButton';
import MailInput from '~/components/molecules/MailInput';
import PasswordInput from '~/components/molecules/PasswordInput';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { UiContext } from '~/lib/context';
import { Status } from '~/lib/context/ui';
import { generateErrorMessage } from '~/lib/util/ErrorUtil';
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
    signInWithEmail: (email: string, password: string) => Promise<void>;
    getAuth: () => Promise<void>;
    registerForPushNotifications: () => Promise<void>;
  };
}

const SignIn = function SignIn(props: Props): ReactElement {
  const navigation = useNavigation();
  const { setApplicationState } = React.useContext(UiContext);
  const { signInWithEmail, getAuth, registerForPushNotifications } = props.actions;

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

  const createAlertNotifySignupError = (code: number) =>
    Alert.alert('サインインに失敗', `${generateErrorMessage(code)}`, [
      {
        text: 'OK',
      },
    ]);

  const handleSubmit = React.useCallback(async () => {
    await signInWithEmail(formData.email, formData.password)
      .then(() => {
        return registerForPushNotifications();
      })
      .then(() => {
        return getAuth();
      })
      .then(() => {
        setApplicationState(Status.AUTHORIZED);
      })
      .catch((err) => {
        createAlertNotifySignupError(err.code);
      });
  }, [
    formData.email,
    formData.password,
    registerForPushNotifications,
    signInWithEmail,
    getAuth,
    setApplicationState,
  ]);

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
