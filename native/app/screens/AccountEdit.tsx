import React, { ReactElement, useState, useMemo } from 'react';
import { StyleSheet, View, Text, TextInput, Alert } from 'react-native';
import { Button } from 'react-native-elements';
import FullTextInput from '~/components/molecules/FullTextInput';
import HalfTextInput from '~/components/molecules/HalfTextInput';
import NumberTextInput from '~/components/molecules/NumberTextInput';
import PrefecturePicker from '~/components/molecules/PrefecturePicker';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { external } from '~/lib/axios';
import { generateErrorMessage } from '~/lib/util/ErrorUtil';
import { AccountEditForm } from '~/types/forms';
import { COLOR } from '~~/constants/theme';

const maxNameLength7 = 7;
const maxNameLength16 = 16;
const maxNameLength32 = 32;
const maxNameLength64 = 64;

const re = /^[\u3040-\u309F]+$/;

const styles = StyleSheet.create({
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
  },
});

interface Props {
  actions: {
    accountEdit: (formData: AccountEditForm) => Promise<void>;
  };
}

async function searchAddress(postalCode: string) {
  return external
    .get('https://zipcoda.net/api', {
      params: {
        zipcode: postalCode,
      },
    })
    .then(function (r) {
      return r.data;
    })
    .catch((err: Error) => {
      throw err;
    });
}

const AccountEdit = function AccountEdit(props: Props): ReactElement {
  const { accountEdit } = props.actions;
  const [formData, setValue] = useState<AccountEditForm>({
    lastName: '',
    firstName: '',
    lastNameKana: '',
    firstNameKana: '',
    phoneNumber: '',
    postalCode: '',
    prefecture: '',
    city: '',
    addressLine1: '',
    addressLine2: '',
  });

  const postalCheck: boolean = useMemo((): boolean => {
    return formData.postalCode.length == 7;
  }, [formData.postalCode]);

  const canSubmit = useMemo((): boolean => {
    return (
      formData.firstName.length > 0 &&
      formData.lastName.length > 0 &&
      formData.firstNameKana.length > 0 &&
      formData.lastNameKana.length > 0 &&
      formData.phoneNumber.length > 0 &&
      formData.postalCode.length > 0 &&
      formData.prefecture !== undefined &&
      re.test(formData.firstNameKana) &&
      re.test(formData.lastNameKana) &&
      formData.city.length > 0 &&
      formData.addressLine1.length > 0
    );
  }, [
    formData.firstName,
    formData.lastName,
    formData.firstNameKana,
    formData.lastNameKana,
    formData.phoneNumber,
    formData.postalCode,
    formData.prefecture,
    formData.city,
    formData.addressLine1,
  ]);

  const createAlertNotifyEditPasswordError = (code: number) =>
    Alert.alert('アカウントの編集に失敗', `${generateErrorMessage(code)}`, [
      {
        text: 'OK',
      },
    ]);

  const handleSearch = React.useCallback(() => {
    (async () => {
      const address = await searchAddress(formData.postalCode);
      const jsonAddress = JSON.stringify(address);
      const parseAddress = JSON.parse(jsonAddress);
      if (parseAddress.status == '200') {
        setValue({
          ...formData,
          prefecture: parseAddress.items[0].components[0],
          city: parseAddress.items[0].components[1],
          addressLine1: parseAddress.items[0].components[2],
        });
      } else {
        throw 'Failed to get address';
      }
    })();
  }, [formData]);

  const handleAccountEditSubmit = React.useCallback(async () => {
    await accountEdit(formData)
      .then(() => {
        //  navigation.navigate('', { });
      })
      .catch((err) => {
        console.log('debug', err);
        createAlertNotifyEditPasswordError(err.code);
      });
  }, [formData, accountEdit]);

  return (
    <View>
      <HeaderWithBackButton title="発送元・お届け先住所" onPress={() => undefined} />
      <Text style={styles.subtitle}>名前</Text>
      <View style={styles.halfInputRow}>
        <HalfTextInput
          onChangeText={(text) => setValue({ ...formData, lastName: text })}
          value={formData.lastName}
          placeholder="田中"
          length={maxNameLength16}
        />
        <HalfTextInput
          onChangeText={(text) => setValue({ ...formData, firstName: text })}
          value={formData.firstName}
          placeholder="太郎"
          length={maxNameLength16}
        />
      </View>
      <Text style={styles.subtitle}>名前(かな)</Text>
      <View style={styles.halfInputRow}>
        <HalfTextInput
          onChangeText={(text) => setValue({ ...formData, lastNameKana: text })}
          value={formData.lastNameKana}
          placeholder="たなか"
          length={maxNameLength16}
        />
        <HalfTextInput
          onChangeText={(text) => setValue({ ...formData, firstNameKana: text })}
          value={formData.firstNameKana}
          placeholder="たろう"
          length={maxNameLength16}
        />
      </View>
      <Text style={styles.subtitle}>電話番号</Text>
      <NumberTextInput
        onChangeText={(text) => setValue({ ...formData, phoneNumber: text })}
        value={formData.phoneNumber}
        placeholder=""
        length={maxNameLength16}
      />
      <Text style={styles.subtitle}>住所</Text>
      <View style={styles.postalArea}>
        <Text>〒 </Text>
        <TextInput
          onChangeText={(text) => setValue({ ...formData, postalCode: text })}
          value={formData.postalCode}
          maxLength={maxNameLength7}
          keyboardType="number-pad"
          style={{ flex: 3, alignSelf: 'stretch' }}></TextInput>
        <Button
          buttonStyle={{ width: '100%' }}
          containerStyle={styles.searchButton}
          disabled={!postalCheck}
          onPress={handleSearch}
          title="検索"
          titleStyle={{ color: COLOR.TEXT_TITLE }}
        />
      </View>
      <View style={styles.prefectureArea}>
        <PrefecturePicker
          onValueChange={(text) => setValue({ ...formData, prefecture: text })}
          value={formData.prefecture}
        />
      </View>
      <FullTextInput
        onChangeText={(text) => setValue({ ...formData, city: text })}
        value={formData.city}
        placeholder="市区町村"
        length={maxNameLength32}
      />
      <FullTextInput
        onChangeText={(text) => setValue({ ...formData, addressLine1: text })}
        value={formData.addressLine1}
        placeholder="地名・番地"
        length={maxNameLength64}
      />
      <FullTextInput
        onChangeText={(text) => setValue({ ...formData, addressLine2: text })}
        value={formData.addressLine2}
        placeholder="マンション・ビル名 部屋番号"
        length={maxNameLength64}
      />
      <Button
        containerStyle={styles.saveButton}
        disabled={!canSubmit}
        onPress={handleAccountEditSubmit}
        title="保存する"
        titleStyle={{ color: COLOR.TEXT_TITLE }}
      />
    </View>
  );
};

export default AccountEdit;
