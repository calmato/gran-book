import React, { ReactElement } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { MaterialCommunityIcons, Ionicons, FontAwesome5 } from '@expo/vector-icons';
import { COLOR } from '~~/constants/theme';
import { TouchableOpacity } from 'react-native-gesture-handler';

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
        />
        <Text>読んだ本</Text>
      </TouchableOpacity>
      <TouchableOpacity style={styles.childStyle} onPress={props.handleNavigateToReadingBoook}>
        <Ionicons 
          name='book' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
        />
        <Text>読んでる本</Text>
      </TouchableOpacity>
      <TouchableOpacity style={styles.childStyle} onPress={props.handleNavigateToTsundoku}>
        <MaterialCommunityIcons 
          name='bookshelf' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
        />
        <Text>積読本</Text>
      </TouchableOpacity>
      <TouchableOpacity style={styles.childStyle} onPress={props.handleNavigateToSellBoook}>
        <FontAwesome5
          name='people-carry' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
        />
        <Text>手放したい本</Text>
      </TouchableOpacity>
      <TouchableOpacity style={styles.childStyle} onPress={props.handleNavigateWishList}>
        <MaterialCommunityIcons 
          name='bookmark-plus' 
          size={ICONSIZE}
          color={COLOR.PRIMARY}
        />
        <Text>欲しい本</Text>
      </TouchableOpacity>
    </View>
  );
};

export default ButtonGroupBookFooter;
