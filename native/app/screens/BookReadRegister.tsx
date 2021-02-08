import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import BookNameAuthorRegister from '~/components/organisms/BookNameAuthorRegister';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import ReadDate from '~/components/organisms/ReadDate';

const styles = StyleSheet.create({
});

const bookInfo = {
  title: '何者',
  image_url: 'https://thechara.xsrv.jp/wp-content/uploads/2020/06/200622%E3%80%90NARUTO%E3%80%91KV_02.jpg',
  author: '稲富',
};

interface Props {}
const BookReadRegister = function BookReadRegister(props: Props): ReactElement {
  return(
    <View>
      <HeaderWithBackButton
        title='読んだ本登録'
        onPress={() => undefined}
      />
    <BookNameAuthorRegister 
      title={bookInfo.title}
      image_url={bookInfo.image_url}
      author={bookInfo.author}
    />
    <ReadDate/>
    </View>
  );
};

export default BookReadRegister;
