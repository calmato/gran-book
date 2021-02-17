import React, { ReactElement } from 'react';
import { TextInput } from 'react-native';

interface Props{
  placeholder: string,
  value: string | undefined,
  onChangeText: (value: string) => void | undefined,
}

const HalfInput = function HalfInput(props: Props): ReactElement {

  return (
    <TextInput
      onChangeText={(text) => props.onChangeText(text)}
      value={props.value}
      placeholder={props.placeholder}
      maxLength={16}
    />
  );
};

export default HalfInput;
