import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { colors, Input } from 'react-native-elements';

const color = colors.grey0;

interface Props {
  onChangeText: (value: string) => void | undefined,
  value: string,
  hasError: boolean,
}

const MailInput = function MailInput(props: Props): ReactElement {

  return (
    <Input
      leftIcon={
        <Ionicons name="md-mail" size={24} color={color} />
      }
      onChangeText={(text) => props.onChangeText(text)}
      value={props.value}
      keyboardType="email-address"
      placeholder="メールアドレス"
      errorMessage={props.hasError ? 'メールアドレスを入力してください．' : undefined }
    />
  );
};

export default MailInput;
