import React, { ReactElement, useState } from 'react';
import { StyleSheet, View, ScrollView, Linking } from 'react-native';
import { ListItem, Text, Avatar, Header, Input } from 'react-native-elements';
import { MaterialIcons } from '@expo/vector-icons';
import { COLOR } from '~~/constants/theme';
import { Auth } from '~/store/models';
import { useNavigation } from '@react-navigation/native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import {Picker} from '@react-native-picker/picker';
import { DEFAULT_ACTION_IDENTIFIER } from 'expo-notifications';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  subtilte: {
    paddingTop: 12,
    paddingLeft: 12,
    paddingBottom: 12,
    fontSize: 15,
    color: COLOR.TEXT_WHITE,
    backgroundColor: COLOR.SECONDARY,
    fontWeight: '600',
  },
  defaultText: {
    marginTop: 12,
    marginBottom:12,
    marginLeft: 12,
    marginRight: 12,
  },
  responseBox: {
    marginTop: 12,
    marginBottom: 12,
    marginLeft: 12,
    flexDirection: 'row',
  },
  questionDetail: {
    minHeight: 200,
    textAlignVertical: 'top',
    paddingRight: 20,
  },
});

interface Props {
}

const Support = function Support(props: Props): ReactElement{
  const navigation = useNavigation();
  const supportSubjectList = [
      "（お問い合わせの内容を選択）",
      "本人情報を変更できない",
      "アプリの不具合",
      "機能の要望",
      "取引中の商品について",
      "商品の削除・警告について",
      "その他",
  ];
  const [supportSubject, setsupportSubject] = useState(supportSubjectList[0])
  let tomorrow = new Date();
  tomorrow.setDate(tomorrow.getDate() + 2);
  const mm = String(tomorrow.getMonth() + 1).padStart(2, '0');
  const dd = String(tomorrow.getDate()).padStart(2, '0');
  const dayOfWeek = tomorrow.getDay() ;	// 曜日(数値)
  const dayOfWeekStr = [ "日", "月", "火", "水", "木", "金", "土" ][dayOfWeek] ;
  const [detail, setdetail] = useState("")

  return (
    <View>
      <ScrollView stickyHeaderIndices={[0]}>
      <HeaderWithBackButton title="お問い合わせ" onPress={() => navigation.goBack()} />
      <Text style={styles.subtilte}>アカウント設定</Text>
      <Text style={styles.defaultText}>お問い合わせ前に、ヘルプをご確認ください。</Text>
      <ListItem bottomDivider onPress={() => undefined}>
        <ListItem.Content>
          <ListItem.Title>{'ヘルプページへ'}</ListItem.Title>
        </ListItem.Content>
        <MaterialIcons name="keyboard-arrow-right" size={24} color="black" />
      </ListItem>
      <Text style={styles.subtilte}>お問い合わせ</Text>
      <Text style={[{color: COLOR.TEXT_WARNING},styles.defaultText]}>
        取引中の商品についてのお問い合わせは、取引画面下の「この取引についてのお問い合わせ」よりおねがいいたします。
      </Text>
      <View style={styles.responseBox}>
        <Text>サポート返信目安：</Text>
        <Text style={{color:COLOR.SECONDARY}}>{mm+"/"+dd+"("+dayOfWeekStr+")"}</Text>
      </View>
      <Picker 
        selectedValue={supportSubject}
        onValueChange={(itemVlue, _itemIndex) => 
          {
            setsupportSubject(itemVlue)
          }
        }
      >
        {supportSubjectList.map((item) => {
          return (
            <Picker.Item label={item} value={item} key={item}/>
          );
        })}
      </Picker>
      {supportSubject!=supportSubjectList[0] && 
      <Input
        style={styles.questionDetail}
        placeholder={'お問い合わせ内容を入力してください'}
        multiline={true}
        onChangeText={(text) => setdetail(text)}
        value={detail}
      />
      }
      <Text style={styles.defaultText}>
        {"※問題のある行為や商品、コメントについては、商品詳細下部にある通報機能をご利用ください。\n\n"+
        "※サポート返信目安を過ぎても返信がない場合は、お手数ですが再度お問い合わせください。"}
      </Text>
      </ScrollView>
    </View>
  );
};

export default Support;