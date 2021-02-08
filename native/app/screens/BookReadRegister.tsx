import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';

const styles = StyleSheet.create({});
interface Props {}
const BookReadRegister = function BookReadRegister(props: Props): ReactElement {
  return(
    <View>
      <HeaderWithBackButton
        title='読んだ本登録'
        onPress={() => undefined}
      />
    </View>
  );
};

export default BookReadRegister;
