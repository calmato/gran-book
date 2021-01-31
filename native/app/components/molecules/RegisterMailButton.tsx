import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { StyleSheet, ViewStyle } from 'react-native';
import { Button, colors } from 'react-native-elements';
import { SOCIAL_BUTTON } from '~~/constants/theme';

const color = colors.grey0;

const styles = StyleSheet.create({
  buttonStyle: {
    borderColor: color,
    backgroundColor: '#00000000',
    ...SOCIAL_BUTTON
  },
  iconStyle: {
    marginRight: 10
  },
  fontStyle: {
    color: color
  }
});

interface Props {
  style?: ViewStyle,
  onPress: () => void,
}

const RegisterMailButton = function RegisterMailButton(props: Props): ReactElement {
  return (
    <Button
      icon={
        <Ionicons
          name="md-mail"
          size={24}
          color={color}
          style={styles.iconStyle}
        />
      }
      containerStyle={props.style}
      buttonStyle={styles.buttonStyle}
      titleStyle={styles.fontStyle}
      onPress={props.onPress}
      type="outline"
      title="メールアドレスで新規登録"
    />
  );
};

export default RegisterMailButton;
