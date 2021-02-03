import React, { ReactElement, useMemo, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { ButtonGroup } from 'react-native-elements';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  }
});

const BookShow = function BookShow(): ReactElement {
  
  const buttons = ['情報','感想']

  const [index, setValue] = useState(0);

  return (
    <View style={styles.container}>
      <HeaderWithBackButton
        // TODO 本の名前をタイトルに入れる
        title=''
        // TODO Navigation の変数できたらProps作って追加
        onPress={() => undefined}
      />
      <ButtonGroup
        buttons={buttons}
        onPress={(selectedIndex) => setValue(selectedIndex)}
        selectedIndex={index}
      />
    </View>
  );
};

export default BookShow;
