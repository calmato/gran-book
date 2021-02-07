import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { MaterialCommunityIcons, FontAwesome5 } from '@expo/vector-icons';
import { COLOR } from '~~/constants/theme';
import { TouchableOpacity } from 'react-native-gesture-handler';

const ICONSIZE = 36;

const styles = StyleSheet.create({
  parentStyle:{
    flex: 1,
    flexDirection: 'row',
    backgroundColor: COLOR.BACKGROUND_WHITE,
    alignSelf: 'stretch',
    justifyContent: 'space-between',
    paddingLeft: 20,
    paddingRight: 20,
    marginBottom:10,
  },
  childStyle: {
    alignItems: 'center',
    alignSelf: 'stretch',
    justifyContent: 'flex-end',
  }
});

interface Props {
  handleNavigateToReadBoook: () => void,
  handleNavigateToReadingBoook: () => void,
  handleNavigateToTsundoku: () => void,
  handleNavigateToSellBoook: () => void,
  handleNavigateToWishList: () => void,
}

const ButtonGroupBookFooter = function ButtonGroupBookFooter(props: Props): ReactElement {
  return (
    <View style={styles.parentStyle}>
      <TouchableOpacity style={styles.childStyle} onPress={props.handleNavigateToReadBoook}>
        <MaterialCommunityIcons 
          name='book-plus' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
          style={{marginBottom: 4}}
        />
        <Text>読んだ本</Text>
      </TouchableOpacity>
      <TouchableOpacity style={styles.childStyle} onPress={props.handleNavigateToReadingBoook}>
        <MaterialCommunityIcons 
          name='book-open-page-variant' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
          style={{marginBottom: 4}}
        />
        <Text>読んでる本</Text>
      </TouchableOpacity>
      <TouchableOpacity style={styles.childStyle} onPress={props.handleNavigateToTsundoku}>
        <MaterialCommunityIcons 
          name='bookshelf' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
          style={{marginBottom: 4}}
        />
        <Text>積読本</Text>
      </TouchableOpacity>
      <TouchableOpacity style={styles.childStyle} onPress={props.handleNavigateToSellBoook}>
        <FontAwesome5
          name='people-carry' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
          style={{marginBottom: 4}}
        />
        <Text>手放したい本</Text>
      </TouchableOpacity>
      <TouchableOpacity style={styles.childStyle} onPress={props.handleNavigateToWishList}>
        <MaterialCommunityIcons 
          name='bookmark-plus' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
          style={{marginBottom: 4}}
        />
        <Text>欲しい本</Text>
      </TouchableOpacity>
    </View>
  );
};

export default ButtonGroupBookFooter;
