import { Ionicons, MaterialIcons } from '@expo/vector-icons';
import React, { ReactElement, useState } from 'react';
import { colors, Input } from 'react-native-elements';
import { TouchableOpacity } from 'react-native-gesture-handler';

interface Props {
  placeholder: string;
  value: string | undefined;
  onChangeText: (value: string) => void | undefined;
  hasError?: boolean;
  errorMessage?: string;
}

const color = colors.grey0;

const PasswordInput = function PasswordInput(props: Props): ReactElement {
  const [hidden, setValue] = useState(true);

  return (
    <Input
      leftIcon={<MaterialIcons name="lock" size={24} color={color} />}
      secureTextEntry={hidden}
      placeholder={props.placeholder}
      onChangeText={(text) => props.onChangeText(text)}
      value={props.value}
      rightIcon={
        <TouchableOpacity onPress={() => setValue(!hidden)}>
          <Ionicons name={hidden ? 'md-eye' : 'md-eye-off'} size={24} color={color} />
        </TouchableOpacity>
      }
      errorMessage={props.hasError ? props.errorMessage : ''}
      maxLength={32}
    />
  );
};

export default PasswordInput;
