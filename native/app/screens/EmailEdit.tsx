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
  subtitle: {
    marginTop: 30,
    marginLeft: 12,
    marginBottom: 6,
    fontSize: 15,
    color: COLOR.TEXT_TITLE,
    fontWeight: '600',
    alignSelf: 'flex-start',
   },
   mailStatus: {
    padding: 15,
    fontSize: 16,
    textAlign: 'right',
    color: COLOR.TEXT_DEFAULT,
    backgroundColor: COLOR.BACKGROUND_WHITE,
    alignSelf: 'stretch',
  },
});

interface Props {
  email: string
}

const EmailEdit = function EmailEdit
(props: Props): ReactElement {
return (
  <View style={styles.container}>
    <HeaderWithBackButton 
    title='メールアドレスの変更' 
    onPress={()=>undefined}
    />
    <Text style={styles.textCard}>新しいメールアドレスを入力してください。{'\n'}確認メールが送信されます。</Text>
    <Text style={styles.subtitle}>現在のメールアドレス</Text>
    <Text style={styles.mailStatus}>{props.email || 'メールアドレス未登録'}</Text>
  </View>
);
}

export default EmailEdit
