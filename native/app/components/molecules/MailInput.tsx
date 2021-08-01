import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { colors, Input } from 'react-native-elements';

const color = colors.grey0;

interface Props {
  onChangeText: (value: string) => void | undefined;
  value: string;
  hasError: boolean;
  sameEmailError: boolean;
}

const MailInput = function MailInput(props: Props): ReactElement {
  const generateErrorMessage = () => {
    if (props.hasError) {
      return 'メールアドレスを入力してください．';
    } else if (props.sameEmailError) {
      return '現在登録されたものとは異なるメールアドレスを入力してください．';
    } else {
      return undefined;
    }
  };

  return (
    <Input
      leftIcon={<Ionicons name="md-mail" size={24} color={color} />}
      onChangeText={(text) => props.onChangeText(text)}
      value={props.value}
      keyboardType="email-address"
      placeholder="メールアドレス"
      errorMessage={generateErrorMessage()}
      inputStyle={{ fontSize: 18 }}
    />
  );
};

export default MailInput;
