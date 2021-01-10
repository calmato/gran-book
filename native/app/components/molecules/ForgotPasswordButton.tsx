import { Ionicons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { StyleSheet, ViewStyle } from 'react-native';
import { Button } from 'react-native-elements';

const styles = StyleSheet.create({
  buttonStyle: {
    justifyContent: 'flex-end'
  },
  titleStyle: {
    fontWeight: 'normal',
    fontSize: 14,
    color: 'black'
  },
  iconStyle: {
    marginLeft: 10
  },
});

interface Props {
  style?: ViewStyle
}

const ForgotPasswoedButton = function ForgotPasswoedButton(props:Props): ReactElement{
  return (
    <Button
      icon={
        <Ionicons
          name="chevron-forward-outline"
          size={24}
          color="black"
          style={styles.iconStyle}
        />
      }
      buttonStyle={styles.buttonStyle}
      iconRight={true}
      containerStyle={props.style}
      title="パスワードを忘れた方"
      titleStyle={styles.titleStyle}
      type={'clear'}
    />
  );
};

export default ForgotPasswoedButton;
