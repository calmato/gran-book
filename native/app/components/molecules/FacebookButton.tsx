import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { StyleSheet, ViewStyle } from 'react-native';
import { Button } from 'react-native-elements';
import { COLOR, SOCIAL_BUTTON } from '~~/constants/theme';

const styles = StyleSheet.create({
  buttonStyle: {
    backgroundColor: COLOR.FACEBOOK,
    ...SOCIAL_BUTTON
  },
  iconStyle: {
    marginRight: 10
  }
});

interface Props {
  style?: ViewStyle
}

const FacebookButton = function FacebookButton(props: Props): ReactElement {
  return (
    <Button
      icon={
        <Ionicons
          name="logo-facebook"
          size={24}
          color="white"
          style={styles.iconStyle}
        />
      }
      buttonStyle={styles.buttonStyle}
      containerStyle={props.style}
      title="Facebookでサインイン" />
  );
};

// FacebookButton.defaultProps={}

export default FacebookButton;
