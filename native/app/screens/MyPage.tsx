import {
  Ionicons,
  MaterialCommunityIcons,
  MaterialIcons,
  FontAwesome,
  FontAwesome5,
} from '@expo/vector-icons';

import { useNavigation } from '@react-navigation/native';
import React, { ReactElement } from 'react';
import { StyleSheet, View, ScrollView } from 'react-native';
import { ListItem, Text, Avatar, Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import { Auth } from '~/store/models';
import { COLOR, FONT_SIZE } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  title: {
    color: COLOR.TEXT_DEFAULT,
    fontSize: FONT_SIZE.ITEM_TITLE,
  },
  subtitle: {
    marginTop: 8,
    marginLeft: 12,
    marginBottom: 4,
    fontSize: FONT_SIZE.SUBHEAD,
    color: COLOR.TEXT_TITLE,
    fontWeight: '600',
  },
});

interface Props {
  auth: Auth.Model;
}

const MyPage = function MyPage(props: Props): ReactElement {
  const navigation = useNavigation();
  const { auth } = props;
  const iconColor = COLOR.GREY;

  // TODO: 型定義
  const avatar = {
    name: auth?.username || 'hamachans',
    thumbnailUrl:
      auth?.thumbnailUrl ||
      'https://pbs.twimg.com/profile_images/1312909954148253696/Utr-sa_Y_400x400.jpg',
  };

  return (
    <View>
      <Header centerComponent={<HeaderText title="Gran Book" />} />
      <ScrollView style={{ marginBottom: 'auto' }}>
        <View>
          <ListItem bottomDivider onPress={() => navigation.navigate('OwnProfile')}>
            <Avatar source={{ uri: avatar.thumbnailUrl }} rounded />
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{avatar.name}</ListItem.Title>
            </ListItem.Content>
          </ListItem>
        </View>
        <View>
          <Text style={styles.subtitle}>マイメニュー</Text>
          <ListItem key={1} bottomDivider>
            <Ionicons name="person" size={24} color={iconColor} />
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'プロフィール'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
          <ListItem key={2} bottomDivider>
            <MaterialCommunityIcons name="account-group" size={24} color={iconColor} />
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'友達一覧'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
          <ListItem key={3} bottomDivider>
            <MaterialCommunityIcons name="message" size={24} color={iconColor} />
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'メッセージボックス'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
        </View>
        <View>
          <Text style={styles.subtitle}>読書関連</Text>
          <ListItem key={4} bottomDivider>
            <MaterialCommunityIcons name="bookshelf" size={24} color={iconColor} />
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'本棚'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
          <ListItem key={5} bottomDivider>
            <MaterialCommunityIcons name="file-document-edit" size={24} color={iconColor} />
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'自分の感想'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
          <ListItem key={6} bottomDivider>
            <FontAwesome5 name="calendar-alt" size={24} color={iconColor} />
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'新刊チェック'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
        </View>
        <View>
          <Text style={styles.subtitle}>フリマ関連</Text>
          <ListItem key={7} bottomDivider>
            <MaterialCommunityIcons name="book-plus-multiple" size={24} color={iconColor} />
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'出品リスト'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
          <ListItem key={8} bottomDivider>
            <MaterialCommunityIcons name="cart" size={24} color={iconColor} />
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'購入リスト'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
          <ListItem key={9} bottomDivider>
            <MaterialCommunityIcons name="piggy-bank" size={24} color={iconColor} />
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'売り上げ申請'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
          <ListItem key={10} bottomDivider>
            <FontAwesome5 name="history" size={24} color={iconColor} />
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'最近見た商品'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
          <ListItem key={11} bottomDivider>
            <FontAwesome name="comments" size={24} color={iconColor} />
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'コメントした商品'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
        </View>
        <View>
          <Text style={styles.subtitle}>その他</Text>
          <ListItem key={12} bottomDivider>
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'お知らせ'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
          <ListItem key={13} bottomDivider>
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'お問い合わせ'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
          <ListItem key={14} bottomDivider>
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'ヘルプ'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
          <ListItem key={15} bottomDivider onPress={() => navigation.navigate('AccountSetting')}>
            <ListItem.Content>
              <ListItem.Title style={styles.title}>{'設定'}</ListItem.Title>
            </ListItem.Content>
            <MaterialIcons name="keyboard-arrow-right" size={24} color={iconColor} />
          </ListItem>
        </View>
      </ScrollView>
    </View>
  );
};

export default MyPage;
