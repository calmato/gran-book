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
  fullTextInput: {
    width: '100%',
    padding: 15,
    marginRight: 5,
    marginTop: 12,
    backgroundColor: COLOR.BACKGROUND_WHITE,
    fontSize: FONT_SIZE.INPUTAREA,
  },
});

const FullTextInput = function FullInput(props: Props): ReactElement {
  return (
    <TextInput
      style={styles.fullTextInput}
      onChangeText={(text) => props.onChangeText(text)}
      value={props.value}
      placeholder={props.placeholder}
      maxLength={props.length}
    />
  );
};

export default FullTextInput;
