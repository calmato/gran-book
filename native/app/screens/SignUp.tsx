import { Ionicons } from '@expo/vector-icons';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { Button, CheckBox, colors, Input, Text } from 'react-native-elements';
import MailInput from '~/components/molecules/MailInput';
import PasswordInput from '~/components/molecules/PasswordInput';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { AuthStackParamList } from '~/types/navigation';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
});

type SignUpProp = StackNavigationProp<AuthStackParamList, 'SignUp'>

interface Props {
  navigation: SignUpProp,
}

const SignUp = function SignUp(props: Props): ReactElement {
  const navigation = props.navigation;

  const [email, setEmail] = useState<string>();

  return (
    <View style={styles.container} >
      <HeaderWithBackButton
        title="ユーザー登録"
        onPress={() => navigation.goBack() }
      />
      <MailInput
        onChangeText={(text) => setEmail(text)}
        value={email}
      />
      <PasswordInput
        placeholder="パスワード"
      />
      <PasswordInput
        placeholder="パスワード(確認用)"
      />
      <Input
        leftIcon={
          <Ionicons name="md-person" size={24} color={colors.grey0} />
        }
        placeholder="ニックネーム"
      />
      {/* <CheckBox checked={false} title="利用規約に同意しました"/> */}
      <Button onPress={() => navigation.navigate('SignUpCheckEmail', {email: email})} title="登録する"/>
      <Text>利用規約</Text>
    </View>
  );
};

export default SignUp;
