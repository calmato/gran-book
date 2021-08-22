import { useNavigation } from '@react-navigation/native';
import React, { ReactElement, useState, useMemo } from 'react';
import { StyleSheet, View, Text, Alert, DevSettings } from 'react-native';
import { Button } from 'react-native-elements';
import MailInput from '~/components/molecules/MailInput';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { generateErrorMessage } from '~/lib/util/ErrorUtil';
import { emailValidation } from '~/lib/validation';
import { COLOR, FONT_SIZE } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    alignItems: 'center',
  },
  textCard: {
    width: '90%',
    fontSize: FONT_SIZE.TEXT_DEFAULT,
    color: COLOR.TEXT_DEFAULT,
    backgroundColor: COLOR.BACKGROUND_YELLOW,
    borderColor: COLOR.PRIMARY,
    borderWidth: 2.5,
    marginTop: 10,
    paddingStart: 30,
    paddingVertical: 10,
  },
  subtitle: {
    marginTop: 30,
    marginLeft: 12,
    marginBottom: 6,
    fontSize: FONT_SIZE.TITLE_SUBHEADER,
    color: COLOR.TEXT_TITLE,
    fontWeight: '600',
    alignSelf: 'flex-start',
  },
  mailStatus: {
    padding: 16,
    fontSize: FONT_SIZE.LISTITEM_TITLE,
    textAlign: 'right',
    color: COLOR.TEXT_DEFAULT,
    backgroundColor: COLOR.BACKGROUND_WHITE,
    alignSelf: 'stretch',
  },
  input: {
    fontSize: 16,
    alignSelf: 'stretch',
    paddingStart: 12,
    paddingVertical: 15,
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
  buttonStyle: {
    marginTop: 20,
  },
  buttonTitleStyle: {
    color: COLOR.TEXT_TITLE,
  },
});

interface Props {
  email: string;
  actions: {
    emailEdit: (email: string) => Promise<void>;
    signOut: () => Promise<void>;
  };
}

const EmailEdit = function EmailEdit(props: Props): ReactElement {
  const navigation = useNavigation();
  const { emailEdit, signOut } = props.actions;
  const [emailForm, setState] = useState('');

  const emailError: boolean = useMemo((): boolean => {
    return !emailValidation(emailForm);
  }, [emailForm]);

  const emailSameError: boolean = useMemo((): boolean => {
    return emailForm === props.email;
  }, [emailForm, props.email]);

  const createAlertNotifyEmailEditError = (code: number) =>
    Alert.alert('メールアドレス変更に失敗', `${generateErrorMessage(code)}`, [
      {
        text: 'OK',
      },
    ]);

  const handleSubmit = React.useCallback(async () => {
    await emailEdit(emailForm)
      .then(() => {
        signOut();
      })
      .then(() => {
        DevSettings.reload();
      })
      .catch((err) => {
        console.log('debug', err);
        createAlertNotifyEmailEditError(err.staus);
      });
  }, [emailForm, emailEdit, signOut]);

  return (
    <View style={styles.container}>
      <HeaderWithBackButton title="メールアドレスの変更" onPress={() => navigation.goBack()} />
      <Text style={styles.textCard}>
        新しいメールアドレスを入力してください。{'\n'}確認メールが送信されます。
      </Text>
      <Text style={styles.subtitle}>現在のメールアドレス</Text>
      <Text style={styles.mailStatus}>{props.email || 'メールアドレス未登録'}</Text>
      <Text style={styles.subtitle}>新しいメールアドレス</Text>
      <MailInput
        onChangeText={(text) => setState(text)}
        value={emailForm}
        hasError={emailError}
        sameEmailError={emailSameError}
      />
      <Button
        containerStyle={styles.buttonStyle}
        disabled={emailError || emailSameError}
        onPress={() => handleSubmit()}
        title="変更する"
        titleStyle={styles.buttonTitleStyle}
      />
    </View>
  );
};

export default EmailEdit;
