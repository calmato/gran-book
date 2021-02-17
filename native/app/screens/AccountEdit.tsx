import { StackNavigationProp } from '@react-navigation/stack';
import { RootStackParamList } from '~/types/navigation';
import React, { ReactElement, useState, useMemo } from 'react';
import { StyleSheet, View, Text, ScrollView, SafeAreaView, TextInput } from 'react-native';
import { COLOR } from '~~/constants/theme';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { AccountEditForm } from '~/types/forms';
import HalfInput from '~/components/molecules/HalfInput';

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
    paddingBottom: 200,
  },
  halfInputRow: {
    flexDirection: 'row',
  },
});

type AccountEditProp = StackNavigationProp<RootStackParamList, 'AccountEdit'>;

interface Props {
  navigation: AccountEditProp
}

const AccountEdit = function AccountEdit(props: Props): ReactElement {
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
            <HalfInput
              onChangeText={(text) => setValue({...formData, firstName: text})}
              value={formData.firstName}
              placeholder="田中"
            />
            <HalfInput
              onChangeText={(text) => setValue({...formData, lastName: text})}
              value={formData.lastName}
              placeholder="太郎"
            />
          </View>
          <Text style={styles.subtilte}>
            名前(かな)
          </Text>
        </ScrollView>
      </SafeAreaView>
    </View>
  );
};

export default AccountEdit;
