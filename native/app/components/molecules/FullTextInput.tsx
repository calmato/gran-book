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
  halfTextInput: {
    width: '100%',
    padding: 15,
    marginRight: 5,
    marginTop: 12,
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
});

const FullTextInput = function FullInput(props: Props): ReactElement {

  return (
    <View style={styles.halfTextInput}>
      <TextInput
        onChangeText={(text) => props.onChangeText(text)}
        value={props.value}
        placeholder={props.placeholder}
        maxLength={props.length}
      />
    </View>
  );
};

export default FullTextInput;
