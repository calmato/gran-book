import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  container:{
    alignItems: 'center',
  },
  textCard: {
    width: '90%',
    fontSize: 16,
    color: COLOR.TEXT_DEFAULT,
    backgroundColor: COLOR.BACKGROUND_YELLOW,
    borderColor: COLOR.PRIMARY,
    borderWidth: 2.5,
    marginTop:10,
    paddingStart: 30,
    paddingVertical: 10,
  },
});

interface Props {}

const EmailEdit = function EmailEdit
(props: Props): ReactElement {
  const statusDefault = 'メールアドレス未登録';
return (
  <View style={styles.container}>
    <HeaderWithBackButton 
    title='メールアドレスの変更' 
    onPress={()=>undefined}
    />
    <Text style={styles.textCard}>新しいメールアドレスを入力してください。{'\n'}確認メールが送信されます。</Text>
  </View>
);
}

export default EmailEdit
