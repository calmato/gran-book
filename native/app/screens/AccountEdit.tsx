import { StackNavigationProp } from '@react-navigation/stack';
import { RootStackParamList } from '~/types/navigation';
import React, { ReactElement, useState, useMemo } from 'react';
import { StyleSheet, View, Text, ScrollView, SafeAreaView, TextInput, Button } from 'react-native';
import { COLOR } from '~~/constants/theme';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { AccountEditForm } from '~/types/forms';
import HalfTextInput from '~/components/molecules/HalfTextInput';
import NumberTextInput from '~/components/molecules/NumberTextInput';

const maxNameLength = 16;
const styles = StyleSheet.create({
  subtilte: {
    marginTop: 12,
    marginLeft: 12,
    marginBottom: 6,
    fontSize: 15,
    color: COLOR.TEXT_TITLE,
    fontWeight: '600',
  },
  scrollArea: {
    // paddingBottom: 200,
  },
  halfInputRow: {
    flexDirection: 'row',
  },
  postalArea: {
    padding: 15,
    backgroundColor: COLOR.BACKGROUND_WHITE,
    flexDirection: 'row',
    alignItems: 'center',
  },
  searchButton: {
    flex: 1,
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

  const nameError: boolean = useMemo((): boolean => {
    return formData.firstName.length < maxNameLength;
  }, [formData.firstName]);

  const handlePostaSubmit = React.useCallback(async () => {
    await searchAdress(
      formData.postalCode
    );
    // .then(() => {
    // })
    // .catch(() => {
    // });
  }, [formData.postalCode]);

  //TODO: かな入力・数字入力のvalidationを追加する
  return (
    <View style={styles.scrollArea}>
      <HeaderWithBackButton
        title='発送元・お届け先住所'
        onPress={() => navigation.goBack()}
      />
      <SafeAreaView>
        <ScrollView>
          <Text style={styles.subtilte}>
            名前
          </Text>
          <View style={styles.halfInputRow}>
            <HalfTextInput
              onChangeText={(text) => setValue({...formData, firstName: text})}
              value={formData.firstName}
              placeholder='田中'
              length={16}
            />
            <HalfTextInput
              onChangeText={(text) => setValue({...formData, lastName: text})}
              value={formData.lastName}
              placeholder='太郎'
              length={16}
            />
          </View>
          <Text style={styles.subtilte}>
            名前(かな)
          </Text>
          <View style={styles.halfInputRow}>
            <HalfTextInput
              onChangeText={(text) => setValue({...formData, firstNameKana: text})}
              value={formData.firstNameKana}
              placeholder='たなか'
              length={16}
            />
            <HalfTextInput
              onChangeText={(text) => setValue({...formData, lastNameKana: text})}
              value={formData.lastNameKana}
              placeholder='たろう'
              length={16}
            />
          </View>
          <Text style={styles.subtilte}>
            電話番号
          </Text>
          <NumberTextInput
            onChangeText={(text) => setValue({...formData, phoneNumber: text})}
            value={formData.phoneNumber}
            placeholder=''
            length={16}
          />
          <Text style={styles.subtilte}>
            住所
          </Text>
          <View style={styles.postalArea}>
            <Text>〒 </Text>
            <TextInput
              onChangeText={(text) => setValue({...formData, postalCode: text})}
              value={formData.postalCode}
              maxLength={7}
              keyboardType="number-pad"
              style={{flex:3, alignSelf: 'stretch'}}
            >
            </TextInput>
            <View style={styles.searchButton}>
              <Button onPress={handlePostaSubmit} title='検索する' color={COLOR.TEXT_TITLE}/>
            </View>
          </View>
        </ScrollView>
      </SafeAreaView>
    </View>
  );
};

export default AccountEdit;
