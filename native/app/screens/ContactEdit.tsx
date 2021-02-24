import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { RootStackParamList } from '~/types/navigation';
import { COLOR } from '~~/constants/theme';
import { ListItem } from 'react-native-elements';
import { MaterialIcons } from '@expo/vector-icons';

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
    textAlign: 'right',
    fontSize: 16,
    color: COLOR.TEXT_DEFAULT,
    backgroundColor: COLOR.BACKGROUND_WHITE
  },
  textStyle: {
    fontSize: 16,
    color: COLOR.TEXT_DEFAULT,
  },
});

type ContactEditProp = StackNavigationProp<RootStackParamList, 'ContactEdit'>;

interface Props {
  navigation: ContactEditProp
}

const ContactEdit = function ContactEdit(props: Props): ReactElement {
  const navigation = props.navigation;
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
    </View>
  );
};

// .defaultProps={}

export default ContactEdit;
