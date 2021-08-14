import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { COLOR, FONT_SIZE } from '~~/constants/theme';

const styles = StyleSheet.create({
  boxLayout: {
    flex: 1,
    flexDirection: 'row',
    marginTop: 20,
    marginBottom: 20,
    paddingTop: 10,
    paddingBottom: 10,
    paddingLeft: 20,
    paddingRight: 20,
    alignSelf: 'stretch',
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
  text: {
    color: COLOR.TEXT_GRAY,
    fontSize: FONT_SIZE.TEXT_INFO,
  },
  name: {
    marginLeft: 32,
    color: COLOR.TEXT_WARNING,
    textDecorationLine: 'underline',
  },
});

interface Props {
  category: string;
}

const FlexBoxBookCategory = function FlexBoxBookCategory(props: Props): ReactElement {
  return (
    <View style={styles.boxLayout}>
      <Text style={styles.text}>カテゴリ</Text>
      <Text style={styles.name}>{props.category}</Text>
    </View>
  );
};

export default FlexBoxBookCategory;
