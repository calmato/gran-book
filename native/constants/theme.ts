import { Platform, ViewStyle } from 'react-native';
import { colors, Theme } from 'react-native-elements';

const FONT_FAMILY = Platform.OS === 'ios' ? 'Arial' : 'Roboto';

export const COLOR = {
  PRIMARY: '#FFC107',
  PRIMARY_LIGHT: '#FFF350',
  PRIMARY_DARK: '#F57C00',
  SECONDARY: '#F57C00',
  ACCENT: '#03A9F4',
  TEXT_DEFAULT: '#3A3A3A',
  TEXT_TITLE: '#6D4C41',
  TEXT_WHITE: '#FFFFFF',
  TEXT_GRAY: '#939393',
  TEXT_INFO: '#03A9F4',
  TEXT_SUCCESS: '#00E676',
  TEXT_ALERT: '#FF5252',
  TEXT_WARNING: '#FF5252',
  BACKGROUND_WHITE: '#FFFFFF',
  BACKGROUND_GREY: '#F9F9F9',
  BACKGROUND_YELLOW: '#FEEEC3',
  GREY: '#3A3A3A',
  LIGHT_GREY: '#BDBDBD',
  APPLE: '#333333',
  TWITTER: '#55ACEE',
  FACEBOOK: '#3B5998',
  GOOGLE: '#FFFFFF',
  MESSAGE_BACKGROUND: '#1da1f2',
};

export const FONT_SIZE = {
  TAB_TITLE: 14,
  SUBHEAD: 15,
  TEXT: 14,
  INPUTAREA: 18,
  ALERT: 12,
  ITEM_TITLE: 16,
  ITEM_SUBTITLE: 14,
  BUTTON: 18,
  BUTTON_ICON: 10,
};

// react-native-elementsの共通の設定
export const THEME: Theme = {
  colors: {
    primary: COLOR.PRIMARY,
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
    errorStyle: {
      marginLeft: 12,
    },
    leftIcon: {}
  }
};

export const SOCIAL_BUTTON: ViewStyle ={
  width: 330,
  height: 42,
};
