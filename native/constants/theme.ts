import { Platform, ViewStyle } from 'react-native';
import { colors, Theme } from 'react-native-elements';

const FONT_FAMILY = Platform.OS === 'ios' ? 'Arial' : 'Roboto';

export const COLOR = {
  MAIN: '#4ABBF3',
  GRAY: '#3A3A3A',
  APPLE: '#333333',
  TWITTER: '#55ACEE',
  FACEBOOK: '#3B5998',
  GOOGLE: '#FFFFFF',
};

export const THEME: Theme = {
  colors: {
    primary: COLOR.MAIN,
    greyOutline: COLOR.GRAY,
  },
  Button: {
    raised: true,
    titleStyle: {
      fontWeight: 'bold',
    },
  },
  Text: {
    h1Style: {
      fontWeight: '800',
      color: colors.black
    },
    style: {
      fontFamily: FONT_FAMILY,
      color: colors.grey2
    }
  }
};

export const SOCIAL_BUTTON: ViewStyle ={
  width: 330,
  height: 42,
};
