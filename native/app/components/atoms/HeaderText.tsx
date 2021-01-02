import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { colors } from 'react-native-elements';

const styles = StyleSheet.create({
  textStyle: {
    fontSize: 16,
    fontWeight: 'bold',
    color: colors.white,
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
