import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import AppleButton from '~/components/atoms/AppleButton';
import FacebookButton from '~/components/atoms/FacebookButton';
import GoogleButton from '~/components/atoms/GoogleButton';
import MailSignInButton from '~/components/atoms/MailSignInButton';
import RegisterMailButton from '~/components/atoms/RegisterMailButton';
import TwitterButton from '~/components/atoms/TwitterButton';

const styles = StyleSheet.create({
  buttonLayout: {
    marginTop: 10,
    marginBottom: 10,
  }
});

// interface Props {}

const SignInButtonGroup = function SignInButtonGroup(): ReactElement {
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
