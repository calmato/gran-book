import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { StyleSheet, ViewStyle } from 'react-native';
import { Button } from 'react-native-elements';
import { COLOR, SOCIAL_BUTTON } from '~~/constants/theme';

const styles = StyleSheet.create({
  buttonStyle: {
    backgroundColor: COLOR.APPLE,
    ...SOCIAL_BUTTON
  },
  iconStyle: {
    marginRight: 10
  }
});

interface Props {
  style?: ViewStyle
}

const AppleButton = function AppleButton(props: Props): ReactElement {
  return (
    <Button
      icon={
        <Ionicons
          name="logo-apple"
          size={24}
          color="white"
          style={styles.iconStyle}
        />
      }
      buttonStyle={styles.buttonStyle}
      containerStyle={props.style}
      title="Appleでサインイン" />
  );
};

// AppleButton.defaultProps={}

export default AppleButton;
