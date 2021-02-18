import React, { ReactElement, useState } from 'react';
import { StyleSheet, View, ScrollView, Text } from 'react-native';
import { Input, Button } from 'react-native-elements';
import BookNameAuthorRegister from '~/components/organisms/BookNameAuthorRegister';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import ReadDate from '~/components/organisms/ReadDate';

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
});

const bookInfo = {
  title: '何者',
  image_url: 'https://thechara.xsrv.jp/wp-content/uploads/2020/06/200622%E3%80%90NARUTO%E3%80%91KV_02.jpg',
  author: '稲富',
};

const BookReadRegister = function BookReadRegister(): ReactElement {
  const [impreessionData, setState] = useState({
    date: new Date,
    impresstion: '',
    isDateUnknown: false,
  });

  return(
    <View style={styles.container}>
      <HeaderWithBackButton
        title='読んだ本登録'
        onPress={() => undefined}
      />
      <ScrollView>
        <BookNameAuthorRegister 
          title={bookInfo.title}
          image_url={bookInfo.image_url}
          author={bookInfo.author}
        />
        <ReadDate
          date={ impreessionData.date }
          handleSetDate={(date) => setState({...impreessionData, date: date})}
          isDateUnknown={ impreessionData.isDateUnknown }
          handleIsDateUnknown={(isDateUnknown) => setState({...impreessionData, isDateUnknown: isDateUnknown})}
        />
        <Text style={{fontSize: 16, marginStart: 20, marginTop: 20, marginBottom: 10, fontWeight: 'bold'}}>感想</Text>
        <Input
          onChangeText={(text) => setState({...impreessionData, impresstion: text})}
          value={impreessionData.impresstion}
          maxLength={1000}
          multiline={true}
        />
        <View style={{alignItems:'center', marginBottom: 20}}>
          <Button onPress={undefined} title='本を登録する'/>
        </View>
      </ScrollView>
    </View>
  );
};

export default BookReadRegister;
