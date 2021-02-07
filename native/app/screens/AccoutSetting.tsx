import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement, useState } from 'react';
import { Switch } from 'react-native';
import { ListItem } from 'react-native-elements';
import { RootStackParamList } from '~/types/navigation';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { StyleSheet, View, Text, SafeAreaView, ScrollView } from 'react-native';
import { MaterialIcons } from '@expo/vector-icons';

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
    color: '#6D4C41',
    fontWeight: '600',
  },
  scrollArea: {
    paddingBottom: 200,
  }
});

type AccountSettingProp= StackNavigationProp<RootStackParamList, 'SignInSelect'>;

interface Props {
  navigation: AccountSettingProp
}

const AccountSetting  = function AccountSetting(props: Props): ReactElement {
  const navigation = props.navigation;
  const [isEnabled, setIsEnabled] = useState(false);
  const toggleSwitch = () => setIsEnabled(previousState => !previousState);
  return (
    <View>
      <HeaderWithBackButton
        title="アカウント"
        onPress={() => navigation.goBack()}
      />
      <SafeAreaView>
        <ScrollView>
          <Text style={styles.subtilte}>アカウント設定</Text>
          <ListItem key={1} bottomDivider>
            <ListItem.Content>
              <ListItem.Title>{'プロフィール'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
          </ListItem>
          <ListItem key={2} bottomDivider>
            <ListItem.Content>
              <ListItem.Title>{'発送元・お届け先住所'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
          </ListItem>
          <ListItem key={3} bottomDivider>
            <ListItem.Content>
              <ListItem.Title>{'クレジットカード一覧'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
          </ListItem>
          <ListItem key={4} bottomDivider>
            <ListItem.Content>
              <ListItem.Title>{'メールアドレス・パスワード'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
          </ListItem>
          <ListItem key={5} bottomDivider>
            <ListItem.Content>
              <ListItem.Title>{'サインアウト'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
          </ListItem>
          <Text style={styles.subtilte}>プッシュ通知設定</Text>
          <ListItem key={6} bottomDivider>
            <ListItem.Content>
              <ListItem.Title>{'プッシュ通知設定'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
          </ListItem>
          <Text style={styles.subtilte}>ホーム画面表示項目設定</Text>
          <ListItem key={7} bottomDivider>
            <ListItem.Content>
              <ListItem.Title>{'おすすめの本'}</ListItem.Title>
            </ListItem.Content>
            <View >
              <Switch
                trackColor={{ false: '#767577', true: '#81b0ff' }}
                thumbColor={isEnabled ? '#f5dd4b' : '#f4f3f4'}
                ios_backgroundColor="#3e3e3e"
                onValueChange={toggleSwitch}
                value={isEnabled}
              />
            </View>
          </ListItem>
        </ScrollView>
      </SafeAreaView>
    </View>
  );
};

export default AccountSetting;
