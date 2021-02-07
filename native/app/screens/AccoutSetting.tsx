import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement, useState } from 'react';
import { Switch } from 'react-native';
import { ListItem } from 'react-native-elements';
import { RootStackParamList } from '~/types/navigation';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { StyleSheet, View, Text, SafeAreaView, ScrollView } from 'react-native';
import { MaterialIcons } from '@expo/vector-icons';
import { COLOR } from '~~/constants/theme';

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
  },
  version: {
    fontSize: 20,
    color: COLOR.TEXT_GRAY,
  },
});

const version = '1.00'
type AccountSettingProp= StackNavigationProp<RootStackParamList, 'SignInSelect'>;

interface Props {
  navigation: AccountSettingProp
}

const AccountSetting  = function AccountSetting(props: Props): ReactElement {
  const navigation = props.navigation;
  const [isEnabledBook, setIsEnabledBook] = useState(false);
  const [isEnabledInformation, setIsEnabledInformation] = useState(false);
  const [isEnabledImpressions, setIsEnabledImpressions] = useState(false);
  const toggleSwitchBook = () => setIsEnabledBook(previousState => !previousState);
  const toggleSwitchInformation = () => setIsEnabledInformation(previousState => !previousState);
  const toggleSwitchImpressions = () => setIsEnabledImpressions(previousState => !previousState);

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
            <Switch
              trackColor={{ false: COLOR.GREY, true: COLOR.TEXT_SUCCESS }}
              thumbColor={isEnabledBook ? COLOR.BACKGROUND_WHITE : COLOR.BACKGROUND_WHITE}
              ios_backgroundColor= {COLOR.TEXT_GRAY}
              onValueChange={toggleSwitchBook}
              value={isEnabledBook}
            />
          </ListItem>
          <ListItem key={8} bottomDivider>
            <ListItem.Content>
              <ListItem.Title>{'新刊情報'}</ListItem.Title>
            </ListItem.Content>
            <Switch
              trackColor={{ false: COLOR.GREY, true: COLOR.TEXT_SUCCESS }}
              thumbColor={isEnabledInformation ? COLOR.BACKGROUND_WHITE : COLOR.BACKGROUND_WHITE}
              ios_backgroundColor= {COLOR.TEXT_GRAY}
              onValueChange={toggleSwitchInformation}
              value={isEnabledInformation}
            />
          </ListItem>
          <ListItem key={9} bottomDivider>
            <ListItem.Content>
              <ListItem.Title>{'感想'}</ListItem.Title>
            </ListItem.Content>
            <Switch
              trackColor={{ false: COLOR.GREY, true: COLOR.TEXT_SUCCESS }}
              thumbColor={isEnabledImpressions ? COLOR.BACKGROUND_WHITE : COLOR.BACKGROUND_WHITE}
              ios_backgroundColor= {COLOR.TEXT_GRAY}
              onValueChange={toggleSwitchImpressions}
              value={isEnabledImpressions}
            />
          </ListItem>
          <Text style={styles.subtilte}>情報</Text>
          <ListItem key={10} bottomDivider>
            <ListItem.Content>
              <ListItem.Title>{'ライセンス情報'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
          </ListItem>
          <ListItem key={11} bottomDivider>
            <ListItem.Content>
              <ListItem.Title>{'バージョン情報'}</ListItem.Title>
            </ListItem.Content>
          <Text style={styles.version}>{version}</Text>
          </ListItem>
        </ScrollView>
      </SafeAreaView>
    </View>
  );
};

export default AccountSetting;
