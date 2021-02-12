import { Ionicons } from '@expo/vector-icons';
import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { Button, Card, Text, colors } from 'react-native-elements';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { AuthStackParamList } from '~/types/navigation';

type SignUpCheckEmailProp = StackNavigationProp<AuthStackParamList, 'SignUpCheckEmail'>
type SignUpCheckEmailRouteProp = RouteProp<AuthStackParamList, 'SignUpCheckEmail'>

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  card: {
    textAlign: 'center',
    justifyContent: 'center',
    alignItems: 'center'
  },
  cardText: {
    textAlign: 'center',
    margin: 20,
  }
});

interface Props {
  route: SignUpCheckEmailRouteProp,
  navigation: SignUpCheckEmailProp
}

const SignUpCheckEmail = function SignUpCheckEmail(props: Props): ReactElement {

  const email = props.route.params.email;
  const navigation = props.navigation;

  return (
    <View style={styles.container}>
      <HeaderWithBackButton title="メール送信" onPress={() => navigation.goBack()} />
      <Card>
        <View style={styles.card}>
          <Ionicons name="md-mail" size={55} color={colors.grey0} />
          <Text h2>メールを送信しました</Text>
          <Text style={styles.cardText}>
            {email}
            {'\n'}
            にメールを送信しました．
            {'\n'}
            メール内のリンクをクリックして
            {'\n'}
            ユーザー登録を完了してください．
          </Text>
        </View>
        <Button onPress={() => navigation.popToTop()} title="登録/サインイン画面へ" />
      </Card>
    </View>
  );
};

export default SignUpCheckEmail;
