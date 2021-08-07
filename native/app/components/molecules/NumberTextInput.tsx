import React, { ReactElement } from 'react';
import { StyleSheet, TextInput } from 'react-native';
import { COLOR, FONT_SIZE } from '~~/constants/theme';

interface Props {
  placeholder: string;
  value: string | undefined;
  length: number | undefined;
  onChangeText: (value: string) => void | undefined;
}

const styles = StyleSheet.create({
  phoneTextInput: {
    width: '100%',
    padding: 15,
    backgroundColor: COLOR.BACKGROUND_WHITE,
    fontSize: FONT_SIZE.TEXT_INPUT,
  },
});

const NumberTextInput = function PhoneTextInput(props: Props): ReactElement {
  return (
    <TextInput
      style={styles.phoneTextInput}
      onChangeText={(text) => props.onChangeText(text)}
      value={props.value}
      placeholder={props.placeholder}
      maxLength={props.length}
      keyboardType="number-pad"
    />
  );
};

export default NumberTextInput;
