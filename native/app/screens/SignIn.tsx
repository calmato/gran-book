import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement, useMemo, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import MailInput from '~/components/molecules/MailInput';
import PasswordInput from '~/components/molecules/PasswordInput';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { SignInForm } from '~/types/forms';
import { emailValidation } from '~/lib/validation';
import { AuthStackParamList } from '~/types/navigation';
import { Button } from 'react-native-elements';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  }
});

type SignInProp = StackNavigationProp<AuthStackParamList, 'SignIn'>

interface Props {
  navigation: SignInProp,
}

const SignIn = function SignIn(props: Props): ReactElement {
  const navigation = props.navigation;

  const [formData, setValue] = useState<SignInForm>({
    email: '',
    password: ''
  });

  const passwordError: boolean = useMemo((): boolean => {
    return formData.password.length < 6;
  }, [formData.password]);

  const emailError: boolean = useMemo((): boolean => {
    return !emailValidation(formData.email);
  }, [formData.email]);

  const canSubmit = useMemo(():boolean => {
    return !emailError && !passwordError;
  }, [emailError, passwordError]);

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
        errorMessage="パスワードは6文字以上でなければいけません．"
        hasError={passwordError}
      />
      <Button 
        disabled={!canSubmit}
        onPress={() => undefined}
        title="サインイン"
      />
    </View>
  );
};

export default SignIn;