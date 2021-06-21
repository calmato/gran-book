import { Ionicons } from '@expo/vector-icons';
import { useNavigation } from '@react-navigation/native';
import React, { ReactElement, useMemo, useState } from 'react';
import { StyleSheet, View, Alert } from 'react-native';
import { Button, colors, Input } from 'react-native-elements';
import CheckBox from '~/components/molecules/CheckBox';
import MailInput from '~/components/molecules/MailInput';
import PasswordInput from '~/components/molecules/PasswordInput';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { generateErrorMessage } from '~/lib/util/ErrorUtil';
import { emailValidation } from '~/lib/validation';
import { SingUpForm } from '~/types/forms';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  checkBox: {
    marginBottom: 24,
  },
});

interface Props {
  actions: {
    signUpWithEmail: (
      email: string,
      password: string,
      passwordConfirmation: string,
      username: string,
    ) => Promise<void>;
  };
}

const SignUp = function SignUp(props: Props): ReactElement {
  const navigation = useNavigation();
  const { signUpWithEmail } = props.actions;

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

  const passwordConfirmationError: boolean = useMemo((): boolean => {
    return formData.password !== formData.passwordConfirmation;
  }, [formData.password, formData.passwordConfirmation]);

  const canSubmit = useMemo((): boolean => {
    return !emailError && !passwordError && !passwordConfirmationError && formData.agreement;
  }, [emailError, passwordError, passwordConfirmationError, formData.agreement]);

  const createAlertNotifySignupError = (code: number) =>
    Alert.alert('ユーザー登録に失敗', `${generateErrorMessage(code)}`, [
      {
        text: 'OK',
      },
    ]);

  const handleSubmit = React.useCallback(async () => {
    await signUpWithEmail(
      formData.email,
      formData.password,
      formData.passwordConfirmation,
      formData.username,
    )
      .then(() => {
        navigation.navigate('SignUpCheckEmail', { email: formData.email });
      })
      .catch((err) => {
        console.log('debug', err);
        createAlertNotifySignupError(err.code);
      });
  }, [
    formData.email,
    formData.password,
    formData.passwordConfirmation,
    formData.username,
    signUpWithEmail,
    navigation,
  ]);

  return (
    <View style={styles.container}>
      <HeaderWithBackButton title="ユーザー登録" onPress={() => navigation.goBack()} />
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
        errorMessage="パスワードは6文字以上でなければいけません．"
        hasError={passwordError}
      />
      <PasswordInput
        onChangeText={(text) => setValue({ ...formData, passwordConfirmation: text })}
        value={formData.passwordConfirmation}
        placeholder="パスワード(確認用)"
        errorMessage="パスワードが一致しません．"
        hasError={passwordConfirmationError}
      />
      <Input
        leftIcon={<Ionicons name="md-person" size={24} color={colors.grey0} />}
        onChangeText={(text) => setValue({ ...formData, username: text })}
        value={formData.username}
        placeholder="ニックネーム"
      />
      <CheckBox
        styles={styles.checkBox}
        onPress={() => setValue({ ...formData, agreement: !formData.agreement })}
        checked={formData.agreement}
        title="利用規約に同意しました．"
      />
      <Button disabled={!canSubmit} onPress={handleSubmit} title="登録する" />
    </View>
  );
};

export default SignUp;
