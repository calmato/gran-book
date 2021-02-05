import React, { ReactElement, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { Button, Image, Text } from 'react-native-elements';
import ButtonGroupInfoImp from '~/components/organisms/ButtonGroupInfoImp';
import FlexBoxBookCategory from '~/components/organisms/FlexBoxBookCategory';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { COLOR } from '~~/constants/theme';
import ButtonGroupBookFooter from '~/components/organisms/ButtonGroupBookFooter';
import { ScrollView } from 'react-native-gesture-handler';
import { MaterialCommunityIcons } from '@expo/vector-icons'; 

const styles = StyleSheet.create({
  container: {
    // flex: 1,
    alignItems: 'center',
  },
  imageContainer: {
    marginTop: 30, 
    marginBottom: 30,
    width: 200,
    height: 280,
  },
  titleContainer: {
    paddingTop: 10, 
    paddingLeft: 30,
    paddingRight: 30,
    fontSize: 16,
    alignSelf: 'stretch',
    color: COLOR.BLACK,
    backgroundColor: COLOR.WHITE,
  },
  autherContainer: {
    paddingTop: 10,
    paddingLeft: 30,
    paddingRight: 30,
    paddingBottom: 10,
    fontSize: 16,
    alignSelf: 'stretch',
    backgroundColor: COLOR.WHITE,
  },
  detailContainer: {
    marginTop: 20,
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
    <ScrollView contentContainerStyle={styles.container}>
      <HeaderWithBackButton
        // TODO 本の名前をタイトルに入れる
        title='タイトル'
        // TODO Navigation の変数できたらProps作って追加
        onPress={() => undefined}
      />
      <ButtonGroupInfoImp
        handleOnPressed={(selectedIndex) => setValue(selectedIndex)}
        selectedIndex={index}
      />
      <View style={{flexDirection:'row',alignSelf: 'stretch' , justifyContent: 'space-around', alignItems: 'center'}}>
        <MaterialCommunityIcons name="chevron-left-circle" size={36} color={COLOR.TEXT_GRAY} />
        <Image
          // TODO 本の情報を代入
          source={{uri:undefined}}
          style={styles.imageContainer}
          transition={true}
        />
        <MaterialCommunityIcons name="chevron-right-circle" size={36} color={COLOR.TEXT_GRAY} />
      </View>
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
      <Button 
        onPress={() => undefined}
        title="Amazonで見る"
        containerStyle={{marginTop: 10, marginBottom: 10}}
        buttonStyle={{backgroundColor: COLOR.PRIMARY_DARK }}
        
      />
      <Button
        onPress={() => undefined}
        title="本を買う"
        containerStyle={{marginTop: 10, marginBottom: 10}}
      />
    </ScrollView>
  );
};

export default BookShow;
