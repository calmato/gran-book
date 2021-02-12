import React, { ReactElement } from 'react';
import { View, ViewStyle } from 'react-native';
import { Text } from 'react-native-elements';

interface Props {
  style?: ViewStyle,
  text: string,
}

const TitleLogoText = function TitleLogoText(props: Props): ReactElement {
  return (
    <View style={props.style}>
      <Text h1>{props.text}</Text>
    </View>
  );
};

export default TitleLogoText;
