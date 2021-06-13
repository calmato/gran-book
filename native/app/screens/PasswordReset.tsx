import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement, useMemo, useState } from 'react';
import { Alert, StyleSheet, View } from 'react-native';
import { Button } from 'react-native-elements';
import MailInput from '~/components/molecules/MailInput';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { generateErrorMessage } from '~/lib/util/ErrorUtil';
import { emailValidation } from '~/lib/validation';
import { sendPasswordResetEmail } from '~/store/usecases';
import { PasswordResetForm } from '~/types/forms';
import { AuthStackParamList } from '~/types/navigation';

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

  const createAlertNotifyProfileEditError = (code: number) =>
    Alert.alert('パスワード更新に失敗', `${generateErrorMessage(code)}`, [
      {
        text: 'OK',
      },
    ]);

  const resetRequest = async () => {
    await sendPasswordResetEmail(formData.email)
      .then((): void => {
        navigation.navigate('SignUpCheckEmail', { email: formData.email });
      })
      .catch((err): void => {
        throw createAlertNotifyProfileEditError(err);
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
