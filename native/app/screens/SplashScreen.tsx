import LottieView from 'lottie-react-native';
import React, { ReactElement, useRef, useEffect } from 'react';

import { StyleSheet, View, ActivityIndicator } from 'react-native';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
  logo: {
    width: 218.75 / 2,
    height: 250 / 2,
    margin: 30,
    borderRadius: 16,
  },
});

const options = {
  loop: true,
  autoplay: true,
};

const SplashScreen = function SplashScreen(): ReactElement {
  const animation = useRef(null);

  useEffect(() => {
    animation.current.play();
  }, []);

  return (
    <View style={styles.container}>
      {/* <Image style={styles.logo} source={logo} /> */}
      <LottieView
        ref={animation}
        source={require('~~/assets/animations/book')}
        style={{
          width: 300,
          height: 300,
        }}
      />
      <ActivityIndicator size="large" />
    </View>
  );
};

export default SplashScreen;
