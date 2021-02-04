import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { MaterialCommunityIcons, Ionicons, FontAwesome5 } from '@expo/vector-icons';
import { COLOR } from '~~/constants/theme';

const ICONSIZE = 48;

const styles = StyleSheet.create({
  parentStyle:{
    flex: 1,
    flexDirection: 'row',
    backgroundColor: COLOR.WHITE,
    alignSelf: 'stretch',
    justifyContent: 'space-between',
    alignItems: 'flex-end',
    paddingLeft: 20,
    paddingRight: 20,
    paddingBottom:20,
  },
  childStyle: {
    alignItems: 'center',
  }
});

const ButtonGroupBookFooter = function ButtonGroupBookFooter(): ReactElement {
  return (
    <View style={styles.parentStyle}>
      <View style={styles.childStyle}>
        <MaterialCommunityIcons 
          name='book-plus' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
        />
        <Text>読んだ本</Text>
      </View>
      <View style={styles.childStyle}>
        <Ionicons 
          name='book' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
        />
        <Text>読んでる本</Text>
      </View>
      <View style={styles.childStyle}>
        <MaterialCommunityIcons 
          name='bookshelf' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
        />
        <Text>積読本</Text>
      </View>
      <View style={styles.childStyle}>
        <FontAwesome5
          name='people-carry' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
        />
        <Text>手放したい本</Text>
      </View>
      <View style={styles.childStyle}>
        <MaterialCommunityIcons 
          name='bookmark-plus' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
        />
        <Text>欲しい本</Text>
      </View>
    </View>
  );
};

export default ButtonGroupBookFooter;
