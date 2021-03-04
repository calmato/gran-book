import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { COLOR } from '~~/constants/theme';
import { ListItem } from 'react-native-elements';
import { MaterialIcons } from '@expo/vector-icons';
import { useNavigation } from '@react-navigation/native';

const styles = StyleSheet.create({
  subtitle: {
    marginTop: 12,
    marginLeft: 12,
    marginBottom: 6,
    fontSize: 15,
    color: COLOR.TEXT_TITLE,
    fontWeight: '600',
  },
  mailStatus: {
    padding: 15,
    fontSize: 16,
    textAlign: 'right',
    color: COLOR.TEXT_DEFAULT,
    backgroundColor: COLOR.BACKGROUND_WHITE
  },
  textStyle: {
    fontSize: 16,
    color: COLOR.TEXT_DEFAULT,
  },
});

interface Props {
  email: string
}

const ContactEdit = function ContactEdit(props: Props): ReactElement {
  const navigation = useNavigation();
  const statusDefault = 'メールアドレス未登録';
  return (
    <View>
      <HeaderWithBackButton
        title='メールアドレス・パスワード設定'
        onPress={() => navigation.goBack()}
      />
      <Text style={styles.subtitle}>現在のメールアドレス</Text>
      <Text style={styles.mailStatus}>{statusDefault}</Text>
      <ListItem style={{ marginTop: 5 }} bottomDivider >
        <ListItem.Content>
          <Text style={styles.textStyle}>メールアドレスの変更</Text>
        </ListItem.Content>
        <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
      </ListItem>
      <Text style={styles.subtitle}>パスワード</Text>
      <ListItem style={{ marginTop: 5 }} bottomDivider >
        <ListItem.Content>
          <Text style={styles.textStyle}>パスワードの変更</Text>
        </ListItem.Content>
        <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
      </ListItem>
    </View>
  );
};

export default ContactEdit;
