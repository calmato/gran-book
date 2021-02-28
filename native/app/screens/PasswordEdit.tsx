import React, { ReactElement, useState, useMemo } from 'react';
import { StyleSheet, View, Text, Alert } from 'react-native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { COLOR } from '~~/constants/theme';
import { PasswordEditForm } from '~/types/forms';
import PasswordInput from '~/components/molecules/PasswordInput';
import { generateErrorMessage } from '~/lib/util/ErrorUtil';
import { Button } from 'react-native-elements';

const styles = StyleSheet.create({
  container: {
    alignItems: 'center',
  },
  subtilte: {
    marginTop: 12,
    marginLeft: 12,
    marginBottom: 6,
    fontSize: 15,
    color: COLOR.TEXT_TITLE,
    fontWeight: '600',
  },
});

interface Props {
  actions: {
    editPassword: (password: string, passwordConfirmation: string) => Promise<void>,
  },
}

const PasswordEdit = function PasswordEdit(props: Props): ReactElement {
  const { editPassword } = props.actions;
  const [formData, setValue] = useState<PasswordEditForm>({
    password: '',
    passwordConfirmation: '',
  });

const passwordError: boolean = useMemo((): boolean => {
    return formData.password.length < 6;
  }, [formData.password])

const passwordConfirmationError: boolean = useMemo(():boolean => {
    return formData.password !== formData.passwordConfirmation;
  }, [formData.password, formData.passwordConfirmation])

  const canSubmit = useMemo(():boolean => {
    return !passwordError && !passwordConfirmationError
  }, [passwordError, passwordConfirmationError]);


  const createAlertNotifyEditPasswordError= (code: number) =>
    Alert.alert(
      'パスワード編集に失敗',
      `${generateErrorMessage(code)}`,
      [
        {
          text: 'OK',
        }
      ],
    );

  const handleSubmit = React.useCallback(async () => {
    await editPassword(
      formData.password,
      formData.passwordConfirmation,
    )
      .then(() => {
      //  navigation.navigate('', { });
      })
      .catch((err) => {
        console.log('debug', err);
        createAlertNotifyEditPasswordError(err.code);
      });
  }, [formData.password, formData.passwordConfirmation]);


  return (
    <View>
      <HeaderWithBackButton
        title='パスワードの変更'
        onPress={() => undefined}
      />
      <Text style={styles.subtilte}>新しいパスワード</Text>
      <PasswordInput
        onChangeText={(text) => setValue({...formData, password: text})}
        value={formData.password}
        placeholder="新しいパスワード"
        errorMessage="パスワードは6文字以上でなければいけません．"
        hasError={passwordError}
      />
      <PasswordInput
        onChangeText={(text) => setValue({...formData, passwordConfirmation: text})}
        value={formData.passwordConfirmation}
        placeholder="新しいパスワード(確認用)"
        errorMessage="パスワードが一致しません．"
        hasError={passwordConfirmationError}
      />
      <Button
      containerStyle={styles.container}
      disabled={!canSubmit}
      onPress={handleSubmit}
      title="変更する"/>
    </View>
  );
};

export default PasswordEdit;
