import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import AppleButton from '~/components/molecules/AppleButton';
import FacebookButton from '~/components/molecules/FacebookButton';
import GoogleButton from '~/components/molecules/GoogleButton';
import MailSignInButton from '~/components/molecules/MailSignInButton';
import RegisterMailButton from '~/components/molecules/RegisterMailButton';
import TwitterButton from '~/components/molecules/TwitterButton';

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
