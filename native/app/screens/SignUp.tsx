import { Ionicons } from '@expo/vector-icons';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement, useMemo, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { Button, colors, Input } from 'react-native-elements';
import CheckBox from '~/components/molecules/CheckBox';
import MailInput from '~/components/molecules/MailInput';
import PasswordInput from '~/components/molecules/PasswordInput';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { emailValidation } from '~/lib/validation';
import { SingUpForm } from '~/types/forms';
import { AuthStackParamList } from '~/types/navigation';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  checkBox: {
    marginBottom: 24,
  },
});

type SignUpProp = StackNavigationProp<AuthStackParamList, 'SignUp'>

interface Props {
  navigation: SignUpProp,
}

const SignUp = function SignUp(props: Props): ReactElement {
  const navigation = props.navigation;

  const [formData, setValue] = useState<SingUpForm>({
    email: '',
    password: '',
    passwordConfirmation: '',
    username: '',
    agreement: false,
  });

  const emailError: boolean = useMemo((): boolean => {
    return !emailValidation(formData.email);
  }, [formData.email]);

  const passwordError: boolean = useMemo((): boolean => {
    return formData.password.length < 6;
  }, [formData.password]);

  const passwordConfirmationError: boolean = useMemo(():boolean => {
    return formData.password !== formData.passwordConfirmation;
  }, [formData.password, formData.passwordConfirmation]);

  const canSubmit = useMemo(():boolean => {
    return !emailError && !passwordError && !passwordConfirmationError && formData.agreement;
  }, [ emailError, passwordError, passwordConfirmationError, formData.agreement]);

  return (
    <View style={styles.container} >
      <HeaderWithBackButton
        title="ユーザー登録"
        onPress={() => navigation.goBack() }
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
      <PasswordInput
        onChangeText={(text) => setValue({...formData, passwordConfirmation: text})}
        value={formData.passwordConfirmation}
        placeholder="パスワード(確認用)"
        errorMessage="パスワードが一致しません．"
        hasError={passwordConfirmationError}
      />
      <Input
        leftIcon={
          <Ionicons name="md-person" size={24} color={colors.grey0} />
        }
        onChangeText={(text) => setValue({...formData, username: text})}
        value={formData.username}
        placeholder="ニックネーム"
      />
      <CheckBox
        styles={styles.checkBox}
        onPress={() => setValue({...formData, agreement: !formData.agreement})}
        checked={formData.agreement}
        title="利用規約に同意しました．"
      />
      <Button disabled={!canSubmit} onPress={() => navigation.navigate('SignUpCheckEmail', {email: formData.email})} title="登録する"/>
    </View>
  );
};

export default SignUp;
