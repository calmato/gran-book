import React, { ReactElement, useMemo, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { AuthStackParamList } from '~/types/navigation';
import { StackNavigationProp } from '@react-navigation/stack';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';

const styles = StyleSheet.create({
    container: {
      flex: 1,
      alignItems: 'center',
    }
  });

  type PasswordResetProp = StackNavigationProp<AuthStackParamList, 'PasswordReset'>

  interface Props{
      navigation: PasswordResetProp,
  }

  const PasswordReset = function PasswordReset(props: Props): ReactElement {
      const navigation = props.navigation;

      return (
          <View style={styles.container}>
              <HeaderWithBackButton
                title='パスワードリセット'
                onPress={() => navigation.goBack()}
              />
          </View>
      );
  };

  export default PasswordReset;
