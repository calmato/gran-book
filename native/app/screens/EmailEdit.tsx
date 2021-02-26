import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';

const styles = StyleSheet.create({
  textCard: {

  },
});

interface Props {}

const EmailEdit = function EmailEdit
(props: Props): ReactElement {
return (
  <View>
    <HeaderWithBackButton title='メールアドレスの変更' onPress={()=>undefined}/>
    <Text style={styles.textCard}>新しいメールアドレスを入力してください。確認メールが送信されます。</Text>
  </View>
);
}

export default EmailEdit

