import React, { ReactElement, useMemo, useState } from 'react';
import { StyleSheet, View, ScrollView } from 'react-native';
import { ListItem, Text, Input, Button } from 'react-native-elements';
import { MaterialIcons } from '@expo/vector-icons';
import { COLOR } from '~~/constants/theme';
import { useNavigation } from '@react-navigation/native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { Picker } from '@react-native-picker/picker';
import MailInput from '~/components/molecules/MailInput';
import { emailValidation } from '~/lib/validation';
import { Auth } from '~/store/models';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    paddingBottom: 12,
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
    marginBottom: 12,
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
  auth: Auth.Model;
}

const Support = function Support(props: Props): ReactElement {
  const navigation = useNavigation();
  const { auth } = props;
  const supportSubjectList = [
    '（お問い合わせの内容を選択）',
    '本人情報を変更できない',
    'アプリの不具合',
    '機能の要望',
    '取引中の商品について',
    '商品の削除・警告について',
    'その他',
  ];
  const [supportSubject, setsupportSubject] = useState(supportSubjectList[0]);
  const [detail, setdetail] = useState('');
  const [emailForm, setState] = useState(auth.email);

  const emailError: boolean = useMemo((): boolean => {
    return !emailValidation(emailForm);
  }, [emailForm]);

  const formError: boolean = useMemo((): boolean => {
    return !detail;
  }, [detail]);

  const canSubmit = useMemo((): boolean => {
    return !(emailError || formError);
  }, [emailError, formError]);

  const generateErrorMessage = () => {
    if (formError) {
      return 'お問い合わせ内容を入力してください';
    } else {
      return undefined;
    }
  };

  const handleSubmit = undefined;

  return (
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
      <Text style={[{ color: COLOR.TEXT_WARNING }, styles.defaultText]}>
        取引中の商品についてのお問い合わせは、取引画面下の「この取引についてのお問い合わせ」よりおねがいいたします。
      </Text>
      <Text style={styles.defaultText}>サポート返信目安：4営業日</Text>
      <Picker
        selectedValue={supportSubject}
        onValueChange={(itemVlue, _itemIndex) => {
          setsupportSubject(itemVlue);
        }}>
        {supportSubjectList.map((item) => {
          return <Picker.Item label={item} value={item} key={item} />;
        })}
      </Picker>
      {supportSubject != supportSubjectList[0] && (
        <View>
          <Text style={styles.defaultText}>返答を受信したいメールアドレス</Text>
          <MailInput
            onChangeText={(text) => setState(text)}
            value={emailForm}
            hasError={emailError}
            sameEmailError={false}
          />
          <Text style={styles.defaultText}>お問い合わせ内容</Text>
          <Input
            style={styles.questionDetail}
            placeholder={'お問い合わせ内容を入力してください'}
            multiline={true}
            onChangeText={(text) => setdetail(text)}
            value={detail}
            errorMessage={generateErrorMessage()}
          />
        </View>
      )}
      <Text style={styles.defaultText}>
        {'※問題のある行為や商品、コメントについては、商品詳細下部にある通報機能をご利用ください。\n\n' +
          '※サポート返信目安を過ぎても返信がない場合は、お手数ですが再度お問い合わせください。'}
      </Text>
      <View style={styles.container}>
        <Button disabled={!canSubmit} onPress={handleSubmit} title="お問い合わせを送信" />
      </View>
    </ScrollView>
  );
};

export default Support;
