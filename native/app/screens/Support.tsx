import React, { ReactElement } from 'react';
import { StyleSheet, View, ScrollView } from 'react-native';
import { ListItem, Text, Avatar, Header } from 'react-native-elements';
import { Ionicons } from '@expo/vector-icons';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { MaterialIcons } from '@expo/vector-icons';
import { FontAwesome } from '@expo/vector-icons';
import { FontAwesome5 } from '@expo/vector-icons';
import { COLOR } from '~~/constants/theme';
import { Auth } from '~/store/models';
import { useNavigation } from '@react-navigation/native';
import HeaderText from '~/components/atoms/HeaderText';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  subtilte: {
    marginTop: 12,
    marginLeft: 12,
    marginBottom: 6,
    fontSize: 15,
    color: COLOR.TEXT_TITLE,
    fontWeight: '600',
  },
});

interface Props {
}

const Support = function Support(props: Props): ReactElement{
  const navigation = useNavigation();
  return (
    <View>
      <HeaderWithBackButton title="お問い合わせ" onPress={() => navigation.goBack()} />
      <ListItem bottomDivider>
        <ListItem.Content>
          <ListItem.Title>{'お問い合わせ（メールアプリを起動）'}</ListItem.Title>
        </ListItem.Content>
        <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
      </ListItem>
    </View>
  );
};

export default Support;
