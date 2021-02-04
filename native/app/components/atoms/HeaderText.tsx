import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  textStyle: {
    fontSize: 16,
    fontWeight: 'bold',
    color: COLOR.TEXT_TITLE,
  },
  layout: {
    justifyContent: 'center'
  }
});

interface Props {
  title: string
}

const HeaderText = function HeaderText(props: Props): ReactElement {
  return (
    <View style={styles.layout}>
      <Text style={styles.textStyle}>{props.title}</Text>
    </View>
  );
};

export default HeaderText;
