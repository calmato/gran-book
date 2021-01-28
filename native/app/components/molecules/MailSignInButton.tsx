import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { StyleSheet, ViewStyle } from 'react-native';
import { Button } from 'react-native-elements';
import { SOCIAL_BUTTON } from '~~/constants/theme';

const styles = StyleSheet.create({
  buttonStyle: {
    ...SOCIAL_BUTTON,
  },
  iconStyle: {
    marginRight: 10
  }
});

interface Props {
  style?: ViewStyle,
  onPress: () => void,
}

const MailSignInButton = function MailSignInButton(props: Props): ReactElement {
  return (
    <Button
      icon={
        <Ionicons
          name="md-mail"
          size={24}
          color="white"
          style={styles.iconStyle}
        />
      }
      buttonStyle={styles.buttonStyle}
      containerStyle={props.style}
      onPress={props.onPress}
      title="メールアドレスでサインイン"
    />
  );
};

// MailSignInButton.defaultProps={}

export default MailSignInButton;
