import LottieView from 'lottie-react-native';
import React, { ReactElement, useRef, useEffect } from 'react';

import { ActivityIndicator, StyleSheet, View } from 'react-native';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
});

const SplashScreen = function SplashScreen(): ReactElement {
  const animation = useRef<LottieView>(null);

  useEffect(() => {
    if (animation) {
      animation.current?.play();
    }
  }, []);

  return (
    <View style={styles.container}>
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
