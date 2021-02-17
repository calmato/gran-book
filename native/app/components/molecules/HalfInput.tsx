import React, { ReactElement } from 'react';
import { StyleSheet, TextInput, View } from 'react-native';
import { COLOR } from '~~/constants/theme';

interface Props{
  placeholder: string,
  value: string | undefined,
  onChangeText: (value: string) => void | undefined,
}

const styles = StyleSheet.create({
  halfInput: {
    width: '50%',
    padding: 15,
    marginRight: 5,
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
});

const HalfInput = function HalfInput(props: Props): ReactElement {

  return (
    <View style={styles.halfInput}>
      <TextInput
        onChangeText={(text) => props.onChangeText(text)}
        value={props.value}
        placeholder={props.placeholder}
        maxLength={16}
      />
    </View>
  );
};

export default HalfInput;
