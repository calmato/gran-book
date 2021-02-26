import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { COLOR } from '~~/constants/theme';
import RadioButtonRN from 'radio-buttons-react-native';


const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    alignSelf: 'stretch',
    alignItems: 'center',
    backgroundColor: COLOR.BACKGROUND_WHITE,
    marginStart: 10,
    marginEnd: 10,
  },
  text: {
    fontSize: 16,
    color: COLOR.TEXT_DEFAULT,
    paddingStart: 10,
  },
  radioButton: {
    flex:8,
    flexDirection: 'row', 
    marginStart: 10,
    marginBottom: 10,
  },
  box: {
    flex: 1,
  },
});

interface Props {
  gender: number,
}

const data = [
  {
    label: '男性',
    code: 1,
   },
   {
    label: '女性',
    code: 2,
   },
   {
    label: '未選択',
    code: 3,
   },
  ];

const GenderRadioGroup = function GenderRadioGroup(props: Props): ReactElement {
  const [gender, setValue] = React.useState(props.gender || 3);
  return (
    <View style={styles.container}>
      <Text style={[styles.text, {flex: 1}]}>性別</Text>
      <RadioButtonRN
      style={styles.radioButton}
      boxStyle={styles.box}
      textStyle={styles.text}
      data={data}
      initial={gender}
      selectedBtn={(e) => setValue(e.code)}
      activeColor={COLOR.PRIMARY}
      />
    </View>
  );
}

export default GenderRadioGroup
