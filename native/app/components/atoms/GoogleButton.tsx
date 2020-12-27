import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { StyleSheet, ViewStyle } from 'react-native';
import { Button } from 'react-native-elements';
import { COLOR, SOCIAL_BUTTON } from '~~/constants/theme';

const color = '#939393';

const styles = StyleSheet.create({
  buttonStyle: {
    backgroundColor: COLOR.GOOGLE,
    ...SOCIAL_BUTTON
  },
  iconStyle: {
    marginRight: 10,
  },
  fontStyle: {
    color: color,
  }
});

interface Props {
  style?: ViewStyle
}

const GoogleButton = function GoogleButton(props: Props): ReactElement {
  return (
    <Button
      icon={
        <Ionicons
          name="logo-google"
          size={24}
          color={color}
          style={styles.iconStyle}
        />
      }
      buttonStyle={styles.buttonStyle}
      titleStyle={styles.fontStyle}
      containerStyle={props.style}
      title="Googleでサインイン" />
  );
};

// GoogleButton.defaultProps={}

export default GoogleButton;
