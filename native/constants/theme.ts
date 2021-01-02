import { Platform, ViewStyle } from 'react-native';
import { colors, Theme } from 'react-native-elements';

const FONT_FAMILY = Platform.OS === 'ios' ? 'Arial' : 'Roboto';

export const COLOR = {
  MAIN: '#4ABBF3',
  GREY: '#3A3A3A',
  APPLE: '#333333',
  TWITTER: '#55ACEE',
  FACEBOOK: '#3B5998',
  GOOGLE: '#FFFFFF',
};

// react-native-elementsの共通の設定
export const THEME: Theme = {
  colors: {
    primary: COLOR.MAIN,
    grey0: COLOR.GREY,
  },
  Button: {
    raised: true,
    titleStyle: {
      fontWeight: 'bold',
    },
    buttonStyle: {
      width: 310,
      height: 42,
    }
  },
  Text: {
    h1Style: {
      fontWeight: '800',
      color: colors.black,
    },
    h2Style: {
      fontWeight: 'bold',
      fontSize: 18,
      color: colors.black,
    },
    style: {
      fontFamily: FONT_FAMILY,
      color: colors.grey2
    }
  },
  Input: {
    containerStyle: {
      paddingLeft: 0,
      paddingRight: 0,
    },
    inputContainerStyle: {
      backgroundColor: colors.white,
      borderColor: colors.white,
    },
    leftIconContainerStyle: {
      marginLeft: 8,
      marginRight: 8,
    },
    rightIconContainerStyle: {
      marginLeft: 8,
      marginRight: 8,
    },
    leftIcon: {}
  }
};

export const SOCIAL_BUTTON: ViewStyle ={
  width: 330,
  height: 42,
};
