import React, { ReactElement, useState, useMemo } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { COLOR } from '~~/constants/theme';
import MailInput from '~/components/molecules/MailInput';
import { emailValidation } from '~/lib/validation';
import { Button, Input } from 'react-native-elements';

const styles = StyleSheet.create({
  container:{
    alignItems: 'center',
  },
  textCard: {
    width: '90%',
    fontSize: 16,
    color: COLOR.TEXT_DEFAULT,
    backgroundColor: COLOR.BACKGROUND_YELLOW,
    borderColor: COLOR.PRIMARY,
    borderWidth: 2.5,
    marginTop:10,
    paddingStart: 30,
    paddingVertical: 10,
  },
  subtitle: {
    marginTop: 30,
    marginLeft: 12,
    marginBottom: 6,
    fontSize: 16,
    color: COLOR.TEXT_TITLE,
    fontWeight: '600',
    alignSelf: 'flex-start',
  },
  mailStatus: {
    padding: 16,
    fontSize: 16,
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
    backgroundColor:COLOR.BACKGROUND_WHITE,
  },
  buttonStyle: {
    marginTop:20,
  },
  buttonTitleStyle: {
    color: COLOR.TEXT_TITLE,
  },
});

interface Props {
  email: string
}

const EmailEdit = function EmailEdit
(props: Props): ReactElement {
  const [emailForm, setState] = useState('');

  const emailError: boolean = useMemo((): boolean => {
    return !emailValidation(emailForm);
  }, [emailForm]);
  return (
    <View style={styles.container}>
      <HeaderWithBackButton 
        title='メールアドレスの変更' 
        onPress={()=>undefined}
      />
      <Text style={styles.textCard}>新しいメールアドレスを入力してください。{'\n'}確認メールが送信されます。</Text>
      <Text style={styles.subtitle}>現在のメールアドレス</Text>
      <Text style={styles.mailStatus}>{props.email || 'メールアドレス未登録'}</Text>
      <Text style={styles.subtitle}>新しいメールアドレス</Text>
      <MailInput
        onChangeText={(text) => setState(text)}
        value={emailForm}
        hasError={emailError}
        sameEmailError={emailForm === 'A@f'}
      />
      <Button containerStyle={styles.buttonStyle} disabled={emailError} onPress={undefined} title='変更する' titleStyle={styles.buttonTitleStyle}/>
    </View>
  );
};

export default EmailEdit;
