import React, { ReactElement, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { ButtonGroup, Image, Text } from 'react-native-elements';
import ButtonGroupInfoImp from '~/components/organisms/ButtonGroupInfoImp';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  imageContainer: {
    marginTop: 45, 
    width: 200,
    height: 200,
  },
  titleContainer: {
    marginTop: 45, 
    marginLeft: 30,
    marginRight: 30,
    fontSize: 20,
    alignSelf: 'stretch',
    color: COLOR.BLACK
  },
  autherContainer: {
    marginTop: 15,
    marginLeft: 30,
    marginRight: 30,
    fontSize: 20,
    alignSelf: 'stretch'
  },
  detailContainer: {
    marginTop: 45, 
    marginLeft: 20,
    marginRight: 20,
    fontSize: 16,
    alignSelf: 'stretch'
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
        // TODO 画像のURIを代入
        source={{uri:undefined}}
        style={styles.imageContainer}
        transition={true}
      />
      <Text style={styles.titleContainer}>題名</Text>
      <Text style={styles.autherContainer}>著者</Text>
      <Text style={styles.detailContainer}>本の詳細</Text>

    </View>
  );
};

export default BookShow;
