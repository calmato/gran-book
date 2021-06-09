import RadioButtonRN from 'radio-buttons-react-native';
import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { RadioGroupForm } from '~/types/forms';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    alignSelf: 'stretch',
    alignItems: 'center',
    backgroundColor: COLOR.BACKGROUND_WHITE,
    paddingStart: 10,
    paddingEnd: 10,
  },
  text: {
    fontSize: 16,
    color: COLOR.TEXT_DEFAULT,
    paddingStart: 10,
  },
  radioButton: {
    flex: 8,
    flexDirection: 'row',
    marginStart: 10,
    marginBottom: 10,
  },
  box: {
    flex: 1,
    borderColor: COLOR.TEXT_WHITE,
  },
});

interface Props {
  handleOnChange: (value: string) => void;
  data: Array<RadioGroupForm>;
  title: string;
  initial: number;
}

const GenderRadioGroup = function GenderRadioGroup(props: Props): ReactElement {
  return (
    <View style={styles.container}>
      <Text style={[styles.text, { flex: 1 }]}>{props.title}</Text>
      <RadioButtonRN
        style={styles.radioButton}
        boxStyle={styles.box}
        textStyle={styles.text}
        data={props.data}
        initial={props.initial}
        selectedBtn={(e) => props.handleOnChange(e.label)}
        activeColor={COLOR.PRIMARY}
      />
    </View>
  );
};

export default GenderRadioGroup;
