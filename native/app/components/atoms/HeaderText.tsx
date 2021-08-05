import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { TouchableOpacity } from 'react-native-gesture-handler';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  textStyle: {
    fontSize: 16,
    fontWeight: 'bold',
    color: COLOR.TEXT_TITLE,
  },
  layout: {
    justifyContent: 'center',
  },
});

interface Props {
  title: string;
  onPress?: () => void;
}

const HeaderText = function HeaderText(props: Props): ReactElement {
  const headerText = (text: string) => {
    return (
      <Text style={styles.textStyle} numberOfLines={1}>
        {text}
      </Text>
    );
  };

  return (
    <View style={styles.layout}>
      {props.onPress ? (
        <TouchableOpacity onPress={props.onPress}>{headerText(props.title)}</TouchableOpacity>
      ) : (
        headerText(props.title)
      )}
    </View>
  );
};

export default HeaderText;
