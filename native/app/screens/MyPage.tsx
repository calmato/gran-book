import { StackNavigationProp } from '@react-navigation/stack';
import { StyleSheet, View, ScrollView, SafeAreaView } from 'react-native';
import React, { ReactElement } from 'react';
import { RootStackParamList } from '~/types/navigation';
import { ListItem, Text, colors, Avatar } from 'react-native-elements';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { Ionicons } from '@expo/vector-icons';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { MaterialIcons } from '@expo/vector-icons';
import { FontAwesome } from '@expo/vector-icons';
import { FontAwesome5 } from '@expo/vector-icons';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  subtilte: {
    marginTop: 12,
    marginLeft: 12,
    fontSize: 15,
    color: colors.black,
    fontWeight: '600'
  },
  scrollArea: {
    paddingBottom: 200
  }
});

type MyPageProp = StackNavigationProp<RootStackParamList, 'SignInSelect'>;

interface Props {
  navigaton: MyPageProp
}

const avatarList =
  {
    name: 'hamachans',
    avatar_url: 'https://storage.cloud.google.com/presto-pay-dev.appspot.com/user_thumbnails/80d01b6c-566f-43fa-89e1-7b54cfcb6558',
  };

const MyPage = function MyPage(props: Props): ReactElement {
  const navigaton = props.navigaton;

  return (
    <View　style={styles.scrollArea}>
      <HeaderWithBackButton
        title="マイページ"
        onPress={() => navigaton.goBack()}
      />
      <SafeAreaView>
        <ScrollView>
          <View>
            <ListItem bottomDivider>
              <Avatar source={{uri: avatarList.avatar_url}} />
              <ListItem.Content>
                <ListItem.Title>{avatarList.name}</ListItem.Title>
              </ListItem.Content>
            </ListItem>
          </View>
          <View>
            <Text style={styles.subtilte}>マイメニュー</Text>
            <ListItem key={1} bottomDivider>
              <Ionicons name="person" size={24} color="black" />
              <ListItem.Content>
                <ListItem.Title>{'プロフィール'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
            <ListItem key={2} bottomDivider>
              <MaterialCommunityIcons name="account-group" size={24} color="black" />
              <ListItem.Content>
                <ListItem.Title>{'友達一覧'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
            <ListItem key={3} bottomDivider>
              <MaterialCommunityIcons name="message" size={24} color="black" />
              <ListItem.Content>
                <ListItem.Title>{'メッセージボックス'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
          </View>
          <View>
            <Text style={styles.subtilte}>読書関連</Text>
            <ListItem key={4} bottomDivider>
              <MaterialCommunityIcons name="bookshelf" size={24} color="black" />
              <ListItem.Content>
                <ListItem.Title>{'本棚'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
            <ListItem key={5} bottomDivider>
              <MaterialCommunityIcons name="file-document-edit" size={24} color="black" />
              <ListItem.Content>
                <ListItem.Title>{'自分の感想'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
            <ListItem key={6} bottomDivider>
              <FontAwesome5 name="calendar-alt" size={24} color="black" />
              <ListItem.Content>
                <ListItem.Title>{'新刊チェック'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
          </View>
          <View>
            <Text style={styles.subtilte}>フリマ関連</Text>
            <ListItem key={7} bottomDivider>
              <MaterialCommunityIcons name="book-plus-multiple" size={24} color="black" />
              <ListItem.Content>
                <ListItem.Title>{'出品リスト'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
            <ListItem key={8} bottomDivider>
              <MaterialCommunityIcons name="cart" size={24} color="black" />
              <ListItem.Content>
                <ListItem.Title>{'購入リスト'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
            <ListItem key={9} bottomDivider>
              <MaterialCommunityIcons name="piggy-bank" size={24} color="black" />
              <ListItem.Content>
                <ListItem.Title>{'売り上げ申請'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
            <ListItem key={10} bottomDivider>
              <FontAwesome5 name="history" size={24} color="black" />
              <ListItem.Content>
                <ListItem.Title>{'最近見た商品'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
            <ListItem key={11} bottomDivider>
              <FontAwesome name="comments" size={24} color="black" />
              <ListItem.Content>
                <ListItem.Title>{'コメントした商品'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
          </View>
          <View>
            <Text style={styles.subtilte}>その他</Text>
            <ListItem key={12} bottomDivider>
              <ListItem.Content>
                <ListItem.Title>{'お知らせ'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
            <ListItem key={13} bottomDivider>
              <ListItem.Content>
                <ListItem.Title>{'お問い合わせ'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
            <ListItem key={14} bottomDivider>
              <ListItem.Content>
                <ListItem.Title>{'ヘルプ'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
            <ListItem key={15} bottomDivider>
              <ListItem.Content>
                <ListItem.Title>{'設定'}</ListItem.Title>
              </ListItem.Content>
              <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
            </ListItem>
          </View>
        </ScrollView>
      </SafeAreaView>
    </View>
  );
};

export default MyPage;
