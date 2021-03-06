import React, { ReactElement } from 'react';
import { StyleSheet, ViewStyle, Image } from 'react-native';
import { Button } from 'react-native-elements';
import google from '~~/assets/g-logo.png';
import { COLOR, SOCIAL_BUTTON } from '~~/constants/theme';

const color = 'rgba(0,0,0,0.54)';

const styles = StyleSheet.create({
  buttonStyle: {
    backgroundColor: COLOR.GOOGLE,
    ...SOCIAL_BUTTON,
  },
  iconStyle: {
    marginRight: 10,
    width: 24,
    height: 24,
  },
  fontStyle: {
    color: color,
  },
});

interface Props {
  style?: ViewStyle;
  onPress: () => void;
}

const GoogleButton = function GoogleButton(props: Props): ReactElement {
  return (
    <Button
      icon={<Image source={google} style={styles.iconStyle} />}
      buttonStyle={styles.buttonStyle}
      titleStyle={styles.fontStyle}
      containerStyle={props.style}
      onPress={props.onPress}
      title="Googleでサインイン"
    />
  );
};

export default GoogleButton;
