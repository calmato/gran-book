import { MaterialIcons } from '@expo/vector-icons';
import { useNavigation } from '@react-navigation/core';
import React, { ReactElement, useState } from 'react';
import { Switch, StyleSheet, View, Text, SafeAreaView, ScrollView } from 'react-native';
import { ListItem } from 'react-native-elements';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { UiContext } from '~/lib/context';
import { Status } from '~/lib/context/ui';
import { COLOR, FONT_SIZE } from '~~/constants/theme';
const styles = StyleSheet.create({
  subtilte: {
    marginTop: 12,
    marginLeft: 12,
    marginBottom: 6,
    fontSize: FONT_SIZE.TITLE_SUBHEADER,
    color: COLOR.TEXT_TITLE,
    fontWeight: '600',
  },
  scrollArea: {
    paddingBottom: 200,
  },
  textStyle: {
    fontSize: FONT_SIZE.LISTITEM_TITLE,
    color: COLOR.TEXT_DEFAULT,
  },
});

const version = '1.00';

interface Props {
  actions: {
    signOut: () => Promise<void>;
  };
}

const AccountSetting = function AccountSetting(props: Props): ReactElement {
  const { signOut } = props.actions;
  const navigation = useNavigation();
  const { setApplicationState } = React.useContext(UiContext);
  const [isEnabledBook, setIsEnabledBook] = useState(false);
  const [isEnabledInformation, setIsEnabledInformation] = useState(false);
  const [isEnabledImpressions, setIsEnabledImpressions] = useState(false);
  const toggleSwitchBook = () => setIsEnabledBook((previousState) => !previousState);
  const toggleSwitchInformation = () => setIsEnabledInformation((previousState) => !previousState);
  const toggleSwitchImpressions = () => setIsEnabledImpressions((previousState) => !previousState);

  const accountSettingList = [
    {
      title: 'プロフィール',
      onClicked: undefined,
    },
    {
      title: '発送元・お届け先住所',
      onClicked: undefined,
    },
    {
      title: 'クレジットコード一覧',
      onClicked: undefined,
    },
    {
      title: 'メールアドレス・パスワード',
      onClicked: () => navigation.navigate('ContactEdit'),
    },
    {
      title: 'サインアウト',
      onClicked: () => handleSignOut(),
    },
  ];

  const handleSignOut = React.useCallback(() => {
    signOut().finally(() => {
      setApplicationState(Status.UN_AUTHORIZED);
    });
  }, [signOut, setApplicationState]);

  return (
    <View>
      <HeaderWithBackButton title="アカウント" onPress={() => navigation.goBack()} />
      <SafeAreaView>
        <ScrollView>
          <Text style={styles.subtilte}>アカウント設定</Text>
          <View>
            {accountSettingList.map((item, i) => (
              <ListItem key={i} bottomDivider onPress={item.onClicked}>
                <ListItem.Content>
                  <Text style={styles.textStyle}>{item.title}</Text>
                </ListItem.Content>
                <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
              </ListItem>
            ))}
          </View>
          <Text style={styles.subtilte}>プッシュ通知設定</Text>
          <ListItem key={6} bottomDivider>
            <ListItem.Content>
              <Text style={styles.textStyle}>プッシュ通知</Text>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={COLOR.GREY} />
          </ListItem>
          <Text style={styles.subtilte}>ホーム画面表示項目設定</Text>
          <ListItem key={7} bottomDivider>
            <ListItem.Content>
              <Text style={styles.textStyle}>おすすめの本</Text>
            </ListItem.Content>
            <Switch
              trackColor={{ false: COLOR.GREY, true: COLOR.TEXT_SUCCESS }}
              thumbColor={COLOR.BACKGROUND_WHITE}
              ios_backgroundColor={COLOR.TEXT_GRAY}
              onValueChange={toggleSwitchBook}
              value={isEnabledBook}
            />
          </ListItem>
          <ListItem key={8} bottomDivider>
            <ListItem.Content>
              <Text style={styles.textStyle}>新刊情報</Text>
            </ListItem.Content>
            <Switch
              trackColor={{ false: COLOR.GREY, true: COLOR.TEXT_SUCCESS }}
              thumbColor={COLOR.BACKGROUND_WHITE}
              ios_backgroundColor={COLOR.TEXT_GRAY}
              onValueChange={toggleSwitchInformation}
              value={isEnabledInformation}
            />
          </ListItem>
          <ListItem key={9} bottomDivider>
            <ListItem.Content>
              <Text style={styles.textStyle}>感想</Text>
            </ListItem.Content>
            <Switch
              trackColor={{ false: COLOR.GREY, true: COLOR.TEXT_SUCCESS }}
              thumbColor={COLOR.BACKGROUND_WHITE}
              ios_backgroundColor={COLOR.TEXT_GRAY}
              onValueChange={toggleSwitchImpressions}
              value={isEnabledImpressions}
            />
          </ListItem>
          <Text style={styles.subtilte}>情報</Text>
          <ListItem key={10} bottomDivider>
            <ListItem.Content>
              <Text style={styles.textStyle}>ライセンス情報</Text>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
          </ListItem>
          <ListItem key={11} bottomDivider>
            <ListItem.Content>
              <Text style={styles.textStyle}>バージョン情報</Text>
            </ListItem.Content>
            <Text style={styles.textStyle}>{version}</Text>
          </ListItem>
        </ScrollView>
      </SafeAreaView>
    </View>
  );
};

export default AccountSetting;
