import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { colors, Input } from 'react-native-elements';

const color = colors.grey0;

const MailInput = function MailInput(): ReactElement {
  return (
    <Input
      leftIcon={
        <Ionicons name="md-mail" size={24} color={color} />
      }
      placeholder="メールアドレス"
    />
  );
};

export default MailInput;
