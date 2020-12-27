import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import AppleButton from '../atoms/AppleButton';
import FacebookButton from '../atoms/FacebookButton';
import GoogleButton from '../atoms/GoogleButton';
import MailSignInButton from '../atoms/MailSignInButton';
import RegisterMailButton from '../atoms/RegisterMailButton';
import TwitterButton from '../atoms/TwitterButton';

const styles = StyleSheet.create({
  buttonLayout: {
    marginTop: 10,
    marginBottom: 10,
  }
});

interface Props {}

const SignInButtonGroup = function SignInButtonGroup(props: Props): ReactElement {
  return (
    <View>
      <AppleButton style={styles.buttonLayout} />
      <TwitterButton style={styles.buttonLayout} />
      <FacebookButton style={styles.buttonLayout} />
      <GoogleButton style={styles.buttonLayout} />
      <MailSignInButton style={styles.buttonLayout} />
      <RegisterMailButton style={styles.buttonLayout} />
    </View>
  );
};

export default SignInButtonGroup;
