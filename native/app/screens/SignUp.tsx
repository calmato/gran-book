import { Ionicons, MaterialIcons } from '@expo/vector-icons';
import { useNavigation } from '@react-navigation/native';
import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { Button, CheckBox, colors, Input } from 'react-native-elements';
import MailInput from '~/components/molecules/MailInput';
import PasswordInput from '~/components/molecules/PasswordInput';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    // flexDirection:'row',
    // justifyContent: 'center',
    alignItems: 'center',
  },
});

interface Props {
}

const SignUp = function SignUp(props: Props): ReactElement {
  const navigatiopn = useNavigation();

  return (
    <View style={styles.container} >
      <HeaderWithBackButton
        title="ユーザー登録"
        onPress={() => navigatiopn.goBack() }
      />
      <MailInput />
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
      <Button title="登録する"/>
      <Text>利用規約</Text>
    </View>
  );
};

// SignUp.defaultProps={}

export default SignUp;
