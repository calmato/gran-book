import React, { ReactElement, useMemo, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { AuthStackParamList } from '~/types/navigation';
import { StackNavigationProp } from '@react-navigation/stack';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import MailInput from '~/components/molecules/MailInput';
import { PasswordResetForm } from '~/types/forms';
import { emailValidation } from '~/lib/validation';
import { Button } from 'react-native-elements';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  }
});

type PasswordResetProp = StackNavigationProp<AuthStackParamList, 'PasswordReset'>

interface Props{
  navigation: PasswordResetProp,
}

const PasswordReset = function PasswordReset(props: Props): ReactElement {
  const navigation = props.navigation;

  const [formData, setValue] = useState<PasswordResetForm>({
    email: ''
  });

  const emailError: boolean = useMemo((): boolean => {
    return !emailValidation(formData.email);
  }, [formData.email]);

  const hasError = useMemo((): boolean => {
    return emailError;
  }, [emailError]);

  return (
    <View style={styles.container}>
      <HeaderWithBackButton
        title='パスワードリセット'
        onPress={() => navigation.goBack()}
      />
      <MailInput
        onChangeText={(text) => setValue({ ...formData, email: text})}
        hasError={emailError}
        value={formData?.email}
      />
      <Button
        disabled={hasError}
        onPress={() => undefined}
        title='メールを送信する'
      />
    </View>
  );
};

export default PasswordReset;
