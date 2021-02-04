import React, { ReactElement, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { Image, Text } from 'react-native-elements';
import ButtonGroupInfoImp from '~/components/organisms/ButtonGroupInfoImp';
import FlexBoxBookCategory from '~/components/organisms/FlexBoxBookCategory';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { COLOR } from '~~/constants/theme';
import ButtonGroupBookFooter from '~/components/organisms/ButtonGroupBookFooter';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  imageContainer: {
    marginTop: 45, 
    width: 200,
    height: 280,
  },
  titleContainer: {
    marginTop: 35,
    paddingTop: 10, 
    paddingLeft: 30,
    paddingRight: 30,
    fontSize: 20,
    alignSelf: 'stretch',
    color: COLOR.BLACK,
    backgroundColor: COLOR.WHITE,
  },
  autherContainer: {
    paddingTop: 15,
    paddingLeft: 30,
    paddingRight: 30,
    paddingBottom: 10,
    fontSize: 20,
    alignSelf: 'stretch',
    backgroundColor: COLOR.WHITE,
  },
  detailContainer: {
    marginTop: 40,
    paddingTop: 10, 
    paddingLeft: 20,
    paddingRight: 20,
    paddingBottom: 10,
    fontSize: 16,
    alignSelf: 'stretch',
    backgroundColor: COLOR.WHITE,
  }
});

const BookShow = function BookShow(): ReactElement {

  const [index, setValue] = useState(0);

  return (
    <View style={styles.container}>
      <HeaderWithBackButton
        // TODO 本の名前をタイトルに入れる
        title=''
        // TODO Navigation の変数できたらProps作って追加
        onPress={() => undefined}
      />
      <ButtonGroupInfoImp
        handleOnPressed={(selectedIndex) => setValue(selectedIndex)}
        selectedIndex={index}
      />
      <Image
        // TODO 本の情報を代入
        source={{uri:undefined}}
        style={styles.imageContainer}
        transition={true}
      />
      <Text style={styles.titleContainer}>題名</Text>
      <Text style={styles.autherContainer}>著者</Text>
      <Text style={styles.detailContainer}>本の詳細</Text>
      <FlexBoxBookCategory category={'コミック'}/>
      <ButtonGroupBookFooter
        handleNavigateToReadBoook={()=>undefined}
        handleNavigateToReadingBoook={()=>undefined}
        handleNavigateToTsundoku={()=>undefined}
        handleNavigateToSellBoook={()=>undefined}
        handleNavigateToWishList={()=>undefined}
      />
    </View>
  );
};

export default BookShow;
