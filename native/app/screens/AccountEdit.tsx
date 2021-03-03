import { StackNavigationProp } from '@react-navigation/stack';
import { RootStackParamList } from '~/types/navigation';
import React, { ReactElement, useState, useMemo } from 'react';
import { StyleSheet, View, Text, TextInput } from 'react-native';
import { COLOR } from '~~/constants/theme';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { Button } from 'react-native-elements';
import HalfTextInput from '~/components/molecules/HalfTextInput';
import NumberTextInput from '~/components/molecules/NumberTextInput';
import FullTextInput from '~/components/molecules/FullTextInput';
import PrefecturePicker from '~/components/molecules/PrefecturePicker';
import { AccountEditForm } from '~/types/forms';

const maxNameLength7 = 7;
const maxNameLength16 = 16;
const maxNameLength32 = 32;
const maxNameLength64 = 64;

const styles = StyleSheet.create(
  {
    subtitle: {
      marginTop: 12,
      marginLeft: 12,
      marginBottom: 6,
      fontSize: 15,
      color: COLOR.TEXT_TITLE,
      fontWeight: '600',
    },
    halfInputRow: {
      flexDirection: 'row',
    },
    postalArea: {
      padding: 10,
      backgroundColor: COLOR.BACKGROUND_WHITE,
      flexDirection: 'row',
      alignItems: 'center',
    },
    prefectureArea: {
      padding: 15,
      marginTop: 12,
      backgroundColor: COLOR.BACKGROUND_WHITE,
    },
    searchButton: {
      flex: 1,
      backgroundColor: COLOR.PRIMARY,
    },
    saveButton: {
      alignSelf: 'center',
      marginTop: 20,
      backgroundColor: COLOR.PRIMARY,
    }
  });

type AccountEditProp = StackNavigationProp<RootStackParamList, 'AccountEdit'>;

interface Props {
  navigation: AccountEditProp,
  actions: {
    searchAdress: (postalCode: string) => Promise<void>,
  },
}

const AccountEdit = function AccountEdit(props: Props): ReactElement {
  const searchAdress = props.actions;
  const navigation = props.navigation;
  const [formData, setValue] = useState<AccountEditForm>({
    firstName: '',
    lastName: '',
    firstNameKana: '',
    lastNameKana: '',
    phoneNumber: '',
    postalCode: '',
    prefecture: '',
    city: '',
    addressLine1: '',
    addressLine2: '',
  });

  const postalCheck: boolean = useMemo((): boolean => {
    return formData.postalCode.length == 7;
  }, [formData.postalCode])

  const canSubmit = useMemo((): boolean => {
    return postalCheck;
  }, [postalCheck])

  //TODO: ボタンを押した時の処理を追加する
  const handlePostaSubmit = React.useCallback(async () => {
    await searchAdress(
      formData.postalCode
    );
    // .then(() => {
    // })
    // .catch(() => {
    // });
  }, [formData.postalCode]);

  const handleAccountEditSubmit = React.useCallback(async () => {
    await accountEdit(
      formData.firstName,
      formData.lastName,
      formData.firstNameKana,
      formData.lastNameKana,
      formData.phoneNumber,
      formData.postalCode,
      formData.prefecture,
      formData.city,
      formData.addressLine1,
      formData.addressLine2
    );
    // .then(() => {
    // })
    // .catch(() => {
    // });
  }, [formData.firstName, formData.lastName, formData.firstNameKana, formData.lastNameKana, formData.phoneNumber,
    formData.postalCode, formData.prefecture, formData.city, formData.addressLine1, formData.addressLine2]);

  return (
    <View>
      <HeaderWithBackButton
        title='発送元・お届け先住所'
        onPress={() => navigation.goBack()}
      />
      <Text style={styles.subtitle}>
            名前
      </Text>
      <View style={styles.halfInputRow}>
        <HalfTextInput
          onChangeText={(text) => setValue({...formData, firstName: text})}
          value={formData.firstName}
          placeholder='田中'
          length={maxNameLength16}
        />
        <HalfTextInput
          onChangeText={(text) => setValue({...formData, lastName: text})}
          value={formData.lastName}
          placeholder='太郎'
          length={maxNameLength16}
        />
      </View>
      <Text style={styles.subtitle}>
            名前(かな)
      </Text>
      <View style={styles.halfInputRow}>
        <HalfTextInput
          onChangeText={(text) => setValue({...formData, firstNameKana: text})}
          value={formData.firstNameKana}
          placeholder='たなか'
          length={maxNameLength16}
        />
        <HalfTextInput
          onChangeText={(text) => setValue({...formData, lastNameKana: text})}
          value={formData.lastNameKana}
          placeholder='たろう'
          length={maxNameLength16}
        />
      </View>
      <Text style={styles.subtitle}>
            電話番号
      </Text>
      <NumberTextInput
        onChangeText={(text) => setValue({...formData, phoneNumber: text})}
        value={formData.phoneNumber}
        placeholder=''
        length={maxNameLength16}
      />
      <Text style={styles.subtitle}>
            住所
      </Text>
      <View style={styles.postalArea}>
        <Text>〒 </Text>
        <TextInput
          onChangeText={(text) => setValue({...formData, postalCode: text})}
          value={formData.postalCode}
          maxLength={maxNameLength7}
          keyboardType='number-pad'
          style={{flex:3, alignSelf: 'stretch'}}
        >
        </TextInput>
        <Button
          buttonStyle={{width: '100%'}}
          containerStyle={styles.searchButton}
          disabled={!postalCheck}
          onPress={() => undefined}
          title='検索'
          titleStyle={{ color: COLOR.TEXT_TITLE}}
        />
      </View>
      <View style={styles.prefectureArea}>
        <PrefecturePicker />
      </View>
      <FullTextInput
        onChangeText={(text) => setValue({...formData, city: text})}
        value={formData.city}
        placeholder='市区町村'
        length={maxNameLength32}
      />
      <FullTextInput
        onChangeText={(text) => setValue({...formData, addressLine1: text})}
        value={formData.addressLine1}
        placeholder='地名・番地'
        length={maxNameLength64}
      />
      <FullTextInput
        onChangeText={(text) => setValue({...formData, addressLine2: text})}
        value={formData.addressLine2}
        placeholder='マンション・ビル名 部屋番号'
        length={maxNameLength64}
      />
      <Button
        containerStyle={styles.saveButton}
        disabled={false}
        onPress={() => undefined}
        title='保存する'
        titleStyle={{ color: COLOR.TEXT_TITLE}}
      />
    </View>
  );
};

export default AccountEdit;
