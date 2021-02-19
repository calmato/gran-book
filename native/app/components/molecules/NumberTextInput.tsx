import React, { ReactElement } from 'react';
import { StyleSheet, TextInput, View } from 'react-native';
import { COLOR } from '~~/constants/theme';

interface Props{
  placeholder: string,
  value: string | undefined,
  length: number | undefined,
  onChangeText: (value: string) => void | undefined,
}

const styles = StyleSheet.create({
  phoneTextInput: {
    width: '100%',
    padding: 15,
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
});

const NumberTextInput = function PhoneTextInput(props: Props): ReactElement {

  return (
    <View style={styles.phoneTextInput}>
      <TextInput
        onChangeText={(text) => props.onChangeText(text)}
        value={props.value}
        placeholder={props.placeholder}
        maxLength={props.length}
        keyboardType="number-pad"
      />
    </View>
  );
};

export default NumberTextInput;
