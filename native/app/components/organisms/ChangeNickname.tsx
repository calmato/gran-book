import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { TextInput } from 'react-native-gesture-handler';
import { COLOR, FONT_SIZE } from '~~/constants/theme';

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
    fontSize: FONT_SIZE.LISTITEM_TITLE,
    flex: 1,
  },
  input: {
    flex: 2,
    fontSize: FONT_SIZE.TEXT_INPUT,
  },
});

interface Props {
  value: string;
  handelOnChangeText: (text) => void;
}

const ChangeNickname = function ChangeNickname(props: Props): ReactElement {
  return (
    <View style={styles.container}>
      <Text style={styles.text}>ニックネーム</Text>
      <TextInput
        style={styles.input}
        maxLength={32}
        textAlign="right"
        value={props.value}
        onChangeText={(text) => props.handelOnChangeText(text)}
        placeholder={'ニックネームを入力してください'}
      />
    </View>
  );
};

export default ChangeNickname;
