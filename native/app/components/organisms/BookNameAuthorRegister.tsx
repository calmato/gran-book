import React, { ReactElement } from 'react';
import { StyleSheet, View, Text, Image } from 'react-native';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  bookInfoStyle: {
    flexDirection: 'row',
    paddingStart: 20,
    paddingTop: 20,
    paddingBottom: 20,
    backgroundColor: COLOR.TEXT_WHITE,
  },
});

interface Props {
  title: string;
  imageUrl: string;
  author: string;
}

const BookNameAuthorRegister = function BookNameAuthorRegister(props: Props): ReactElement {
  return (
    <View style={styles.bookInfoStyle}>
      <Image source={{ uri: props.imageUrl }} style={{ width: 75, height: 105 }} />
      <View style={{ justifyContent: 'space-around', marginStart: 20 }}>
        <Text style={{ fontSize: 16 }}>{props.title}</Text>
        <Text style={{ fontSize: 16, color: COLOR.GREY }}>{props.author}</Text>
      </View>
    </View>
  );
};

export default BookNameAuthorRegister;
