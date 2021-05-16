import React, { ReactElement, useMemo, useState } from 'react';
import { Alert, StyleSheet, View } from 'react-native';
import { AuthStackParamList } from '~/types/navigation';
import { StackNavigationProp } from '@react-navigation/stack';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import MailInput from '~/components/molecules/MailInput';
import { PasswordResetForm } from '~/types/forms';
import { emailValidation } from '~/lib/validation';
import { Button } from 'react-native-elements';
import { sendPasswordResetEmail } from '~/store/usecases';
import { generateErrorMessage } from '~/lib/util/ErrorUtil';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
});

type PasswordResetProp = StackNavigationProp<AuthStackParamList, 'PasswordReset'>;

interface Props {
  navigation: PasswordResetProp;
}

const PasswordReset = function PasswordReset(props: Props): ReactElement {
  const navigation = props.navigation;

  const [formData, setValue] = useState<PasswordResetForm>({
    email: '',
  });

  const createAlertNotifyProfileEditError = (errorMessage: string) =>
    Alert.alert('パスワード更新に失敗', `${generateErrorMessage(errorMessage)}`, [
      {
        text: 'OK',
      },
    ]);

  const resetRequest = async () => {
    await sendPasswordResetEmail(formData.email)
      .then((): void => {
        navigation.navigate('SignUpCheckEmail', { email: formData.email });
      })
      .catch((err: Error): void => {
        throw createAlertNotifyProfileEditError(err.message);
      });
  };

  const emailError: boolean = useMemo((): boolean => {
    return !emailValidation(formData.email);
  }, [formData.email]);

  const hasError = useMemo((): boolean => {
    return emailError;
  }, [emailError]);

  return (
    <View style={styles.container}>
      <HeaderWithBackButton title="パスワードリセット" onPress={() => navigation.goBack()} />
      <MailInput
        onChangeText={(text) => setValue({ ...formData, email: text })}
        hasError={emailError}
        value={formData?.email}
        sameEmailError={false}
      />
      <Button disabled={hasError} onPress={resetRequest} title="メールを送信する" />
    </View>
  );
};

export default PasswordReset;
