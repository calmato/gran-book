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
  halfTextInput: {
    width: '50%',
    padding: 15,
    marginRight: 5,
    backgroundColor: COLOR.BACKGROUND_WHITE,
    fontSize: FONT_SIZE.INPUTAREA,
  },
});

const HalfTextInput = function HalfInput(props: Props): ReactElement {
  return (
    <TextInput
      style={styles.halfTextInput}
      onChangeText={(text) => props.onChangeText(text)}
      value={props.value}
      placeholder={props.placeholder}
      maxLength={props.length}
    />
  );
};

export default HalfTextInput;
