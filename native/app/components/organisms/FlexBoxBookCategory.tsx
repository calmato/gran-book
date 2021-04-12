import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { COLOR } from '~~/constants/theme';

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
});

interface Props {
  category: string;
}

const FlexBoxBookCategory = function FlexBoxBookCategory(props: Props): ReactElement {
  return (
    <View style={styles.boxLayout}>
      <Text>カテゴリ</Text>
      <Text style={{ marginLeft: 45 }}>{props.category}</Text>
    </View>
  );
};

export default FlexBoxBookCategory;
