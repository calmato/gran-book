import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement, useState } from 'react';
import { colors, Input } from 'react-native-elements';
import { emailValidation } from '~/lib/validation';

const color = colors.grey0;

interface Props {
  onChangeText: (value: string) => void | undefined,
  value: string | undefined,
}

const MailInput = function MailInput(props: Props): ReactElement {

  const [hasError, setError] = useState<boolean>();

  const onChangeText = (value: string) => {
    !emailValidation(value) ? setError(true) : setError(false);
    props.onChangeText(value);
  };

  return (
    <Input
      leftIcon={
        <Ionicons name="md-mail" size={24} color={color} />
      }
      onChangeText={(text) => onChangeText(text)}
      value={props.value}
      placeholder="メールアドレス"
      errorMessage={hasError ? 'メールアドレスを入力してください．' : undefined }
    />
  );
};

export default MailInput;
