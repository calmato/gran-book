import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { TextInput } from 'react-native-gesture-handler';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    marginTop: 20,
    marginBottom: 20,
    padding: 20,
    flexDirection: 'row',
    backgroundColor: COLOR.BACKGROUND_WHITE,
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  text: {
    color: COLOR.TEXT_DEFAULT,
    fontSize: 16,
    flex: 1,
  },
  input: {
    flex: 2,
    fontSize: 16,
  },
});

interface Props {
  defaultValue: string,
  handelOnChangeText: (text) => void,
}

const ChangeNickname = function ChangeNickname(props: Props): ReactElement {
  return (
    <View style={styles.container}>
      <Text style={styles.text}>ニックネーム</Text>
      <TextInput 
        style={styles.input} 
        maxLength={32} 
        textAlign='right' 
        multiline={true} 
        value={props.defaultValue} 
        onChangeText={(text)=>props.handelOnChangeText(text)}
        placeholder={'ニックネームを入力してください'}
        placeholderTextColor={COLOR.TEXT_ALERT}/>
    </View>
  );
};

export default ChangeNickname;
