import React, { ReactElement, useMemo, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import MailInput from '~/components/molecules/MailInput';
import PasswordInput from '~/components/molecules/PasswordInput';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { SignInForm } from '~/types/forms';
import { emailValidation, passwordValidation } from '~/lib/validation';
import { Button } from 'react-native-elements';
import ForgotPasswordButton from '~/components/molecules/ForgotPasswordButton';
import { useNavigation } from '@react-navigation/native';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  }
});

interface Props {
  actions: {
    signInWithEmail: (email: string, password: string) => Promise<void>,
  },
}

const SignIn = function SignIn(props: Props): ReactElement {
  const navigation = useNavigation();
  const { signInWithEmail } = props.actions;

  const [formData, setValue] = useState<SignInForm>({
    email: '',
    password: ''
  });

  const passwordError: boolean = useMemo((): boolean => {
    return !passwordValidation(formData.password);
  }, [formData.password]);

  const emailError: boolean = useMemo((): boolean => {
    return !emailValidation(formData.email);
  }, [formData.email]);

  const canSubmit = useMemo(():boolean => {
    return !(emailError || passwordError);
  }, [emailError, passwordError]);

  const handleSubmit = React.useCallback(async () => {
    await signInWithEmail(
      formData.email,
      formData.password,
    )
      .then(() => {
        console.log('debug', 'success');
      })
      .catch((err: Error) => {
        console.log('debug', 'failure', err);
      });
  }, [formData.email, formData.password, signInWithEmail]);

  return (
    <View style={styles.container}>
      <HeaderWithBackButton
        title='サインイン'
        onPress={() => navigation.goBack()}
      />
      <MailInput
        onChangeText={(text) => setValue({ ...formData, email: text})}
        hasError={emailError}
        value={formData?.email}
      />
      <PasswordInput
        onChangeText={(text) => setValue({...formData, password: text})}
        value={formData.password}
        placeholder="パスワード"
        errorMessage="パスワードは6文字以上32文字以下でなければいけません．"
        hasError={passwordError}
      />
      <Button
        disabled={!canSubmit}
        onPress={handleSubmit}
        title="サインイン"
      />
      <ForgotPasswordButton
        onPress={() => navigation.navigate('PasswordReset')}
      />
    </View>
  );
};

export default SignIn;
